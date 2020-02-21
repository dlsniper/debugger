// Debugger package allows an application to be compiled for production without
// incurring any penalty due to instrumentation.
//
// When the time comes to use a debugger and inspect your code, use the
// -tags=debugger build tag to switch the implementation to one that will set
// labels that can be displayed in the debugger.
//
// The current implementation relies on pprof labels to be created and then
// the debugger can read them and display them.
package debugger

import "net/http"

// MiddlewareLabels generates labels that are aware of the request properties
type MiddlewareLabels func(r *http.Request) []string

// Labels generates the labels for a function/method
type Labels func() []string
