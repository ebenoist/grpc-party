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

serve: py-serve ruby-serve go-serve
	@echo "make sure to start with make -j3 serve"

py-serve:
	python python/server.py

go-serve:
	go run go/server.go

ruby-serve:
	ruby ruby/server.rb

party:
	@go run go/client.go
	@ruby ruby/client.rb
	@python python/client.py
