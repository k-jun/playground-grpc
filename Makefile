PROTO_FILE:=proto/playground-grpc.proto

.PHONY: build
build:
	mkdir -p grpc/
	protoc --proto_path=proto --go_out=plugins=grpc:grpc --go_opt=paths=source_relative $(PROTO_FILE)

.PHONY: format
format:
	clang-format -i $(PROTO_FILE)
