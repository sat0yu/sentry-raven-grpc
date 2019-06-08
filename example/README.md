# Mirror: sample gRPC server & client

"Mirror" returns the mirrored string of the given one as the request body.  
For example, given "Mirror", the server answers "rorriM" 🔄

## Server
Set the [Sentry DSN](https://docs.sentry.io/clients/go/#configure) in `server.go` beforehand.

```
go run server.go
```

## Client
`client.go` requires one argument which represents the request body.

```
# The response from the server is "CPRg" as mirroring the argument.
go run client.go gRPC

# The server reports an error to Sentry.
go run client.go golang 
```

## Update proto 
```
protoc mirror/mirror.proto --go_out=plugins=grpc:.
```
