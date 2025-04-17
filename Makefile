PROTOC = protoc
PROTOC_GEN_GO = $(shell go env GOPATH)/bin/protoc-gen-go

proto:
	$(PROTOC) --go_out=. --go_opt=paths=source_relative pkg/record/record.proto
