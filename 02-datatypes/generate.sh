#!/bin/bash

# protoc --proto_path=proto --go_out=plugins=grpc:proto datatypes.proto
protoc --proto_path=proto --go_out=plugins=grpc:proto person.proto