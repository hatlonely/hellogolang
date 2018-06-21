package kit

import (
	"context"
	"testing"
	"time"

	"github.com/go-kit/kit/ratelimit"
	"golang.org/x/time/rate"
)

var nopEndpoint = func(context.Context, interface{}) (interface{}, error) { return struct{}{}, nil }

func TestRateLimit(t *testing.T) {
	{
		limit := rate.NewLimiter(rate.Every(time.Minute), 1)
		e := ratelimit.NewErroringLimiter(limit)(nopEndpoint)

		ctx, cxl := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cxl()
		if _, err := e(ctx, struct{}{}); err != nil {
			t.Error(err)
		}

		if _, err := e(ctx, struct{}{}); err != nil {
			t.Error(err)
		}
	}

	{
		limit := rate.NewLimiter(rate.Every(time.Minute), 1)
		e := ratelimit.NewDelayingLimiter(limit)(nopEndpoint)

		ctx, cxl := context.WithTimeout(context.Background(), 50000*time.Millisecond)
		defer cxl()
		if _, err := e(ctx, struct{}{}); err != nil {
			t.Error(err)
		}

		if _, err := e(ctx, struct{}{}); err != nil {
			t.Error(err)
		}
	}
}
