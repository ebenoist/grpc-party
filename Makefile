PROTOC := protoc
RUBY_PROTOC := grpc_tools_ruby_protoc
PYTHON_PROTOC := python -m grpc_tools.protoc
SCHEMAS := party.proto

all: ruby go python

deps:
	go get -u github.com/golang/protobuf/protoc-gen-go
	gem install grpc-tools
	pip install grpcio-tools

.PHONY: ruby
ruby:
	$(RUBY_PROTOC) --ruby_out=./ruby --grpc_out=./ruby  $(SCHEMAS)

.PHONY: go
go:
	mkdir -p go/party
	$(PROTOC)  --go_out=plugins=grpc:go/party $(SCHEMAS)


.PHONY: python
python:
	$(PYTHON_PROTOC) -I . --python_out=./python --grpc_python_out=./python $(SCHEMAS)
