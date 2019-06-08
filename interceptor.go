package interceptor

import (
	"context"
	"google.golang.org/grpc"
)

// ErrorCapturer specifies the implementation of a method to capture the given error
type ErrorCapturer interface {
	CaptureError(err error, tags map[string]string) string
}

// SentryRavenInterceptor creates an interceptor which catches the errors from each service method and reports them to Sentry
func SentryRavenInterceptor(ec ErrorCapturer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, _info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			ec.CaptureError(err, nil)
		}
		return resp, err
	}
}

// SentryRavenInterceptorOption
func SentryRavenInterceptorOption(ec ErrorCapturer) grpc.ServerOption {
	return grpc.UnaryInterceptor(SentryRavenInterceptor(ec))
}
