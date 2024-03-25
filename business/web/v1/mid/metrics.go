package mid

import (
	"context"
	"github.com/draculaas/deploy/business/web/v1/metrics"
	"github.com/draculaas/deploy/core/web"
	"net/http"
)

// Metrics ... Updates program counters
func Metrics() web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			// set metrics into the context, make the call
			ctx = metrics.Set(ctx)

			err := handler(ctx, w, r)

			n := metrics.AddRequests(ctx)
			if n%1000 == 0 {
				metrics.AddGoroutines(ctx)
			}

			if err != nil {
				metrics.AddErrors(ctx)
			}

			return err
		}

		return h
	}

	return m
}
