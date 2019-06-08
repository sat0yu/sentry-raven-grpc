package interceptor_test

import (
	"context"
	"errors"
	"testing"
	. "github.com/sat0yu/sentry-raven-grpc"
)

type ErrorCapturerMock struct {
	CaughtError error
}

func (ec *ErrorCapturerMock) CaptureError(err error, _tags map[string]string) string {
	ec.CaughtError = err
	return err.Error()
}

func Test_SentryRavenInterceptor(t *testing.T) {
	expectedRes := "result"
	expectedErr := errors.New("error")
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return expectedRes, expectedErr
	}
	mock := ErrorCapturerMock{}
	interceptor := SentryRavenInterceptor(&mock)
	actualRes, actualErr := interceptor(context.Background(), nil, nil, handler)
	if mock.CaughtError != expectedErr {
		t.Errorf("got: %v\nwant: %v", mock.CaughtError, expectedErr)
	}
	if actualRes != expectedRes {
		t.Errorf("got: %v\nwant: %v", actualRes, expectedRes)
	}

	if actualErr != expectedErr {
		t.Errorf("got: %v\nwant: %v", actualErr, expectedErr)
	}
}
