# Build the Source Code

## Prerequisites
You must have:
- installed [Git](https://git-scm.com/) 
- installed [Go](https://golang.org/dl/) and setup the GOPATH
- [configured Docker hosts](https://github.com/gauravgahlot/dockerdoodle/blob/master/docs/host-configuration.md) to listen for `TCP` requests at port `2375`

## Getting Required Packages

To start with, you first need to clone the repository. I you would suggest, you select the latest `release` branch. Now let's start installing required packages:

```
# explicitly get GRPC and Protocol Buffers
$ go get -u google.golang.org/grpc
$ go get -u github.com/golang/protobuf/proto

# to get remaining packages (ensure that you are in the reposity directory)
$ go get ./...
```

## Adding Docker Hosts

At the root of the repository, you must see a `config.json` file with a list of hosts. All you need is to add your [pre-configured](https://github.com/gauravgahlot/dockerdoodle/blob/master/docs/host-configuration.md) Docker hosts with their `name` and `IP` address to this list. Since `name` is for representation purpose only, you may name the host as it suits your needs. 

## Running the GRPC Server and Client

Being in the root of your repository, execute the following commands to spin up the server and client:
```
$ go run server/main.go
$ go run client/main.go
```
By default the client and server start listen at port `:7585` and `:7584` respectively.

## Additional Steps for Contributors

### Setup `protoc` compiler

 - In case you plan to change `GRPC` service of `messages` you will need the `protoc` compiler. You can download the same from Protocol Buffers [release page](https://github.com/protocolbuffers/protobuf/releases). 
 - You also need the `protoc-gen-go` plugin for the protocol buffer compiler to generate Go code. You can get it with the following command:
    ```
    $ go get -u github.com/golang/protobuf/protoc-gen-go
    ```

### Compile PB Messages to Go Source Code

Once you have made the changes in PB messages or RPC service, it's time to compile them and generate Go code. Use the following command for the same:
```
protoc -I ./pb-messages ./pb-messages/<changed-file>.proto --go_out=plugins=grpc:./pb
```

### Change the GRPC Client and Server Ports

The default ports for the client and server are defined as `ClientPort` and `ServerPort` in `constants/constants.go`. You just change the port here and it will work. 

If you change the `ClientPort`, it's is **very** important to update the port used for building `Web Socket` connection. This web socket connection is used to *push* container stats from server to the client. Therefore, it's very critical. To change the port, go to `client/templates/content/host.html` and at the end you must see:
```
<snip>
    window.onload = () => {
      if (window["WebSocket"]) {
        conn = new WebSocket("ws://localhost:7585/ws");
<snip>
```
Change the port to `:7585` to the new value of `ClientPort` and you are good.

