.PHONY: localprep
localprep:
	@echo "Making sure Go protobuf generators/helpers are installed"
	@go install github.com/bufbuild/buf/cmd/buf@v1.6.0
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

.PHONY: protogen-python
protogen-python: localprep
    #We have to get wonky here because \
	1. gRPC python codegen sucks a bit \
	2. Python import path resolution sucks a bit \
	3. I don't want to make proto definitions messier just to "fix" Python codegen issues caused by #1 and #2
	@buf generate --template buf.gen.python.accesspdp.yaml --path accesspdp
	@buf generate --template buf.gen.python.attributes.yaml --path attributes

.PHONY: protogen-go
protogen-go: localprep
	@buf generate --template buf.gen.go.yaml

#List targets in makefile
.PHONY: list
list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make database/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'
