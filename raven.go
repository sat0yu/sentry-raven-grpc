package interceptor

import (
	"github.com/getsentry/raven-go"
)

// ravenClient wraps the original raven client
type ravenClient struct {
	*raven.Client
}

// CaptureError calls the original function which provided by the raven package
func (rc ravenClient) CaptureError(err error, tags map[string]string) string {
	return rc.Client.CaptureError(err, tags)
}

// NewClient injects the raven package
func NewClient(dsn string) ravenClient {
	c := raven.DefaultClient
	c.SetDSN(dsn)
	return ravenClient{c}
}
