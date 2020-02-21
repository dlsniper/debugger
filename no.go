// +build !debugger

package debugger

import "net/http"

// Middleware allows HTTP middleware to receive debugger labels
func Middleware(f http.HandlerFunc, l MiddlewareLabels) http.HandlerFunc {
	_ = l
	return f
}

// SetLabels will set debugger labels for any function/method call
func SetLabels(l Labels) {
	_ = l
}
