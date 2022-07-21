FROM golang:1.18-alpine AS builder

ARG GOLANGCI_VERSION=v1.47.2
ARG COVERAGE_THRESH_PCT=81

ENV GO111MODULE=on \
    CGO_ENABLED=0

# Get git and other tools needed
RUN apk add --no-cache git=~2 wget=~1

# Get test coverage tool and protobuf codegen
RUN go install github.com/klmitch/overcover@v1.2.1 \
    && go install github.com/bufbuild/buf/cmd/buf@v1.6.0 \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

WORKDIR /build

# Copy the code necessary to build the application
# Hoovering in everything here doesn't matter -
# we're going to discard this intermediate image later anyway
# and just copy over the resulting binary
COPY . .

RUN go mod tidy

# Vendor modules here so that subsequent steps don't
# re-fetch, and just use the vendored versions
RUN go mod vendor

# Let's create a /dist folder containing just the files necessary for runtime.
# Later, it will be copied as the / (root) of the output image.
RUN mkdir /dist

# Run linters

#Lint/gen protobuf code
WORKDIR /build/proto
RUN buf lint && buf generate
WORKDIR /build

SHELL ["/bin/ash", "-eo", "pipefail", "-c"]
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./ ${GOLANGCI_VERSION}

# TODO this is very slow on `arm64` - like 5x slower
RUN ./golangci-lint --version && ./golangci-lint run --timeout 20m

# Run tests
RUN go test --coverprofile cover.out ./attributes ./pdp

# Test coverage
RUN overcover --coverprofile cover.out ./attributes ./pdp --threshold ${COVERAGE_THRESH_PCT}

# Build the application
RUN go build -o /dist/access-pdp-grpc-server

# Create the minimal runtime image
FROM scratch AS emptyfinal

COPY --chown=0:0 --from=builder /dist/access-pdp-grpc-server /access-pdp-grpc-server

ENTRYPOINT ["/access-pdp-grpc-server"]
