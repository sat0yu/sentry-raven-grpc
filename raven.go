package interceptor

import (
	"github.com/getsentry/raven-go"
)

// RavenClient wraps the original raven client
type RavenClient struct {
	*raven.Client
}

// CaptureError calls the original function which provided by the raven package
func (rc RavenClient) CaptureError(err error, tags map[string]string) string {
	return rc.Client.CaptureError(err, tags)
}

// NewClient injects the raven package
func NewClient(dsn string) RavenClient {
	c := raven.DefaultClient
	c.SetDSN(dsn)
	return RavenClient{c}
}
