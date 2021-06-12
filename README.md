# playground-grpc

## setup

```bash
brew install protobuf
go get -u github.com/golang/protobuf/protoc-gen-go

# document generate
go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
protoc --doc_out=html,index.html:./ proto/*.proto
```

## build 

```bash
make build
```

## format

```bash
make format
```

## run

```bash
go run main.go
```

## reference

* https://developers.google.com/protocol-buffers/docs/proto3

