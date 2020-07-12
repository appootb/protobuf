## Shipping package
ifeq ($(shell uname), Darwin)
PROTO_ROOT_DIR := $(shell brew --prefix)/Cellar/protobuf/*
else
PROTO_ROOT_DIR := /usr/local
endif

SUBDIRS = $(shell ls ./appootb)
CLEANDIRS = $(SUBDIRS:%=go/%)

all: $(SUBDIRS) todo

$(SUBDIRS):
	# generating go/$@
	@rm -rf go/$@
	@mkdir -p go/$@
	@protoc -Iappootb/$@ -I. \
		-I${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--auth_out=lang=go,paths=source_relative:go/$@ \
		--go_out=plugins=grpc,paths=source_relative:go/$@ \
		--grpc-gateway_out=logtostderr=true,paths=source_relative:go/$@ \
		--validate_out=lang=go,paths=source_relative:go/$@ \
		appootb/$@/*.proto

	# generating dart/$@
	@rm -rf dart/$@/lib/src/generated
	@mkdir -p dart/$@/lib/src/generated/appootb/common
	@mkdir -p dart/$@/lib/src/generated/appootb/permission
	@protoc -Iappootb/$@ -I. \
		-I${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--dart_out=grpc:dart/$@/lib/src/generated \
		--dart-export_out=file=$@,paths=source_relative:dart/$@/lib \
		appootb/$@/*.proto
	@protoc --dart_out=dart/$@/lib/src/generated $(PROTO_ROOT_DIR)/include/google/protobuf/*.proto
	@protoc -I. --dart_out=dart/$@/lib/src/generated appootb/*/include.proto
	@protoc -Iappootb/common --dart_out=dart/$@/lib/src/generated/appootb/common appootb/common/*.proto
	@protoc -Iappootb/permission --dart_out=dart/$@/lib/src/generated/appootb/permission appootb/permission/*.proto

	# generating doc/$@, swagger/$@
	@rm -rf doc/$@ swagger/$@
	@mkdir -p doc/$@ swagger/$@
	@protoc -Iappootb/$@ -I. \
		-I${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:swagger/$@ \
		--markdown_out=paths=source_relative:doc/$@ \
		appootb/$@/*.proto

todo:
	@rm -rf go/permission/*.pb.*.go
	@rm -rf go/secret/*.pb.*.go
	@rm -rf doc/common doc/code doc/permission doc/secret
	@rm -rf swagger/common swagger/code swagger/permission swagger/secret
	@rm -rf dart/common dart/code dart/permission dart/secret

.PHONY: $(SUBDIRS)