# go_grpc


init commands:
```
brew install protobuf
brew install protoc-gen-go
```

```
go get -u github.com/golang/protobuf/protoc-gen
```

```
protoc --go_out=plugins=grpc:chat chat.proto

```

