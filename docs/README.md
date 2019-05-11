## Generating Source Code

### Get `protoc` compiler
Download the `protoc` compiler from Protocol Buffers [release page](https://github.com/protocolbuffers/protobuf/releases).


### Install `gRPC` tooling

```
go get -u google.golang.org/grpc
```

### Install Protocol Buffers tooling 

```
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go

```

### Compile PB Messages -> Go Source Code
```
protoc 
    -I ./pb-messages
    ./pb-messages/*.proto
    --go_out=plugins=grpc:./pb
```
