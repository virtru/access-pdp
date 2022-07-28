#!/usr/bin/env bash

# We have to get wonky here because
#
# 1. gRPC python codegen sucks a bit
# 2. Python import path resolution sucks a bit
# 3. I don't want to make proto definitions messier just to "fix" Python codegen issues caused by #1 and #2
buf generate --template buf.gen.python.accesspdp.yaml --path accesspdp
buf generate --template buf.gen.python.attributes.yaml --path attributes
