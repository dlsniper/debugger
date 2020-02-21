// +build debugger

package debugger

import (
	"context"
	"net/http"
	"runtime/pprof"
)

// Middleware allows HTTP middleware to receive debugger labels
func Middleware(f http.HandlerFunc, l MiddlewareLabels) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := pprof.Labels(l(r)...)

		pprof.Do(r.Context(), l, func(ctx context.Context) {
			c := r.Context()
			defer r.WithContext(c)

			f(w, r.WithContext(ctx))
		})
	}
}

var bgCtx = context.Background()

// SetLabels will set debugger labels for any function/method call
func SetLabels(l Labels) {
	ctx := pprof.WithLabels(bgCtx, pprof.Labels(l()...))
	pprof.SetGoroutineLabels(ctx)
}

// SetLabelsWithCtx will set debugger labels for any function/method call using a custom context
func SetLabelsWithCtx(ctx context.Context, l Labels) {
	pprof.SetGoroutineLabels(pprof.WithLabels(ctx, pprof.Labels(l()...)))
}
