# sentry-raven-grpc: Sentry integration for gRPC server
Do you want to integrate a gPRC server with Sentry?  
Do you think `panic` should not be used without proper reason?  

There you go. `sentry-raven-grpc` provide a clean way to integrate with Sentry for gRPC servers.

## Installation
A small setup is required so that all errors raised by each gRPC service method to the Sentry.  
[Take a look at this commit.](https://github.com/sat0yu/sentry-raven-grpc/commit/fcd93038e151952f94a88b10ab29c61efe8187e3), there's all what you need to do.  
Don't forget install the package with the below command. 

```
go get https://github.com/sat0yu/sentry-raven-grpc 
```

## Tips
### Collect full stacktrace
Use ["github.com/pkg/errors"](https://godoc.org/github.com/pkg/errors), not the default "errors".  
The default "errors" package does not give sufficient information for error tracking, while "github.com/pkg/errors" provides all stacktrace from the line raising the original error.

```
## When using "github.com/pkg/errors"

*errors.fundamental: too long
  File "/Users/yusuke_sato/.ghq/github.com/sat0yu/sentry-raven-grpc/example/server/main.go", line 26, in Echo
    return nil, errors.New("too long")
  File "/Users/yusuke_sato/.ghq/github.com/sat0yu/sentry-raven-grpc/example/mirror/mirror.pb.go", line 180, in func1
    return srv.(MirrorServer).Echo(ctx, req.(*EchoRequest))
  File "/Users/yusuke_sato/.ghq/github.com/sat0yu/sentry-raven-grpc/interceptor.go", line 17, in func1
    resp, err := handler(ctx, req)
  File "/Users/yusuke_sato/.ghq/github.com/sat0yu/sentry-raven-grpc/example/mirror/mirror.pb.go", line 182, in _Mirror_Echo_Handler
    return interceptor(ctx, in, info, handler)
  File "/Users/yusuke_sato/go/pkg/mod/google.golang.org/grpc@v1.21.1/server.go", line 998, in processUnaryRPC
    reply, appErr := md.Handler(srv.server, ctx, df, s.opts.unaryInt)
  File "/Users/yusuke_sato/go/pkg/mod/google.golang.org/grpc@v1.21.1/server.go", line 1278, in handleStream
    s.processUnaryRPC(t, stream, srv, md, trInfo)
  File "/Users/yusuke_sato/go/pkg/mod/google.golang.org/grpc@v1.21.1/server.go", line 717, in 1
    s.handleStream(st, stream, s.traceInfo(st, stream))
```
