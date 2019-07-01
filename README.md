# grpc-examples

Clone into your GO_PATH
```bash
$ git clone https://github.com/BenchLord/grpc-examples.git
$ cd grpc-examples/
```
Compile proto

`$ protoc -I proto/ people.proto --go_out=plugins=grpc:proto`

Start server

`$ go run server/server.go`

Start client

`$ go run client/client.go`
