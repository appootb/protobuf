## Shipping package

ifeq ($(shell uname), Darwin)
PROTO_ROOT_DIR := $(shell brew --prefix)/Cellar/protobuf/*
else
PROTO_ROOT_DIR := /usr/local
endif

SUBDIRS = $(shell ls ./appootb)
CLEANDIRS = $(SUBDIRS:%=go/%)

all: $(CLEANDIRS) $(SUBDIRS) fix

$(CLEANDIRS):
	@rm -rf $@/*

$(SUBDIRS):
	@protoc -Iappootb/$@ -I. \
		-I${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:go/$@ \
		--grpc-gateway_out=logtostderr=true:go/$@ \
		--validate_out="lang=go:go/$@" \
		--auth_out="lang=go:go/$@" \
		appootb/$@/*.proto

fix:
	@rm -rf go/permission/*.pb.*.go

clean: $(CLEANDIRS)

.PHONY: $(SUBDIRS)
.PHONY: $(CLEANDIRS)