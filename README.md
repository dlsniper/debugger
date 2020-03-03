# Debugger Middleware

This package provides a debugging middleware for Go applications to enable
better display of goroutines in the debugger.

It has nearly-zero performance penalty in production code when not actively used.

## How this looks like in the IDE

Below you can see how this feature looks like in GoLand IDE:

![Debugger labels in GoLand](debugger-labels.png "Debugger labels in GoLand")

## How to use

Include it in your application using one of the patterns below.

Then, compile the application with `-tags debugger`, e.g.

```shell script
go build -tags debugger
```

More details on how to use this can be found in this blog post:
https://blog.jetbrains.com/go/2020/03/03/how-to-find-goroutines-during-debugging/

### HTTP handlers

In your code, replace the HTTP handler with the `Middleware` function call.

Original:
```go
router.HandleFunc("/", homeHandler)
```

Replacement:
```go
router.HandleFunc("/", debugger.Middleware(homeHandler, func(r *http.Request) []string {
    return []string{
        "path", r.RequestURI,
    }
}))
``` 

### Non-HTTP handlers

For normal functions/methods, you can use the `SetLabels` / `SetLabelsCtx` functions
to set the debugger labels. 

Original:
```go
func sum(a, b int) int {
    return a+b
}
```

Replacement:
```go
func sum(a, b int) int {
    debugger.SetLabels(func() []string {
        return []string{
            "a", strconv.Itoa(a),
            "b", strconv.Itoa(b),
        }
    })

    return a+b
}
```

## Performance

You can asses the performance of this library by running the included benchmarks
in your environment.

Here are the results from my own machine (Intel Core i7 6700HQ, 32GB RAM, Windows 10),
when running with a `-count=5` pass.

Go 1.13.8

Without labels:

| Name | Execution count  | Time |
|---|---|---|
|BenchmarkWorkerWithout-8   |  	 3133654|	       366 ns/op|
|BenchmarkWorkerWithout-8   |  	 3288013|	       373 ns/op|
|BenchmarkWorkerWithout-8   |  	 3306148|	       376 ns/op|
|BenchmarkWorkerWithout-8   |  	 3225550|	       371 ns/op|
|BenchmarkWorkerWithout-8   |  	 3324428|	       366 ns/op|

With a single label:

| Name | Execution count  | Time |
|---|---|---|
|BenchmarkWorkerWithOne-8   |  	   25917|	     44604 ns/op|
|BenchmarkWorkerWithOne-8   |  	   26604|	     45220 ns/op|
|BenchmarkWorkerWithOne-8   |  	   24740|	     47211 ns/op|
|BenchmarkWorkerWithOne-8   |  	   26842|	     46196 ns/op|
|BenchmarkWorkerWithOne-8   |  	   25586|	     45260 ns/op|

With three labels:

| Name | Execution count  | Time |
|---|---|---|
|BenchmarkWorkerWithThree-8 |  	   15584|	     76809 ns/op|
|BenchmarkWorkerWithThree-8 |  	   15769|	     85167 ns/op|
|BenchmarkWorkerWithThree-8 |  	   15936|	     76244 ns/op|
|BenchmarkWorkerWithThree-8 |  	   15325|	     76933 ns/op|
|BenchmarkWorkerWithThree-8 |  	   15666|	     75960 ns/op|

With ten labels:

| Name | Execution count  | Time |
|---|---|---|
|BenchmarkWorkerWithTen-8   |  	    6009|	    218503 ns/op|
|BenchmarkWorkerWithTen-8   |  	    6000|	    211000 ns/op|
|BenchmarkWorkerWithTen-8   |  	    5714|	    228909 ns/op|
|BenchmarkWorkerWithTen-8   |  	    5445|	    210107 ns/op|
|BenchmarkWorkerWithTen-8   |  	    5713|	    215472 ns/op|

With three labels, and an int to string conversion:

| Name | Execution count  | Time |
|---|---|---|
|BenchmarkWorkerWithConv-8  |  	   15403|	     77515 ns/op|
|BenchmarkWorkerWithConv-8  |  	   15055|	     76717 ns/op|
|BenchmarkWorkerWithConv-8  |  	   15644|	     78752 ns/op|
|BenchmarkWorkerWithConv-8  |  	   15344|	     85701 ns/op|
|BenchmarkWorkerWithConv-8  |  	   15423|	     79362 ns/op|


Go 1.14 RC 1

Without labels:

| Name | Execution count  | Time |
|---|---|---|
|BenchmarkWorkerWithout-8   |  	 3108807|	       366 ns/op|
|BenchmarkWorkerWithout-8   |  	 3260863|	       372 ns/op|
|BenchmarkWorkerWithout-8   |  	 3234224|	       365 ns/op|
|BenchmarkWorkerWithout-8   |  	 3157615|	       372 ns/op|
|BenchmarkWorkerWithout-8   |  	 3158044|	       366 ns/op|

With a single label:

| Name | Execution count  | Time |
|---|---|---|
|BenchmarkWorkerWithOne-8   |  	   25476|	     46004 ns/op|
|BenchmarkWorkerWithOne-8   |  	   25750|	     46487 ns/op|
|BenchmarkWorkerWithOne-8   |  	   25104|	     45969 ns/op|
|BenchmarkWorkerWithOne-8   |  	   26084|	     46772 ns/op|
|BenchmarkWorkerWithOne-8   |  	   25803|	     47165 ns/op|

With three labels:

| Name | Execution count  | Time |
|---|---|---|
|BenchmarkWorkerWithThree-8 |  	   19323|	     60291 ns/op|
|BenchmarkWorkerWithThree-8 |  	   19606|	     59774 ns/op|
|BenchmarkWorkerWithThree-8 |  	   19639|	     60339 ns/op|
|BenchmarkWorkerWithThree-8 |  	   20338|	     59593 ns/op|
|BenchmarkWorkerWithThree-8 |  	   19866|	     60658 ns/op|

With ten labels:

| Name | Execution count  | Time |
|---|---|---|
|BenchmarkWorkerWithTen-8   |  	    7064|	    172565 ns/op|
|BenchmarkWorkerWithTen-8   |  	    7507|	    169574 ns/op|
|BenchmarkWorkerWithTen-8   |  	    7504|	    174311 ns/op|
|BenchmarkWorkerWithTen-8   |  	    7476|	    169342 ns/op|
|BenchmarkWorkerWithTen-8   |  	    7058|	    173420 ns/op|

With three labels, and an int to string conversion:

| Name | Execution count  | Time |
|---|---|---|
|BenchmarkWorkerWithConv-8  |  	   19353|	     64589 ns/op|
|BenchmarkWorkerWithConv-8  |  	   17340|	     66205 ns/op|
|BenchmarkWorkerWithConv-8  |  	   19016|	     64314 ns/op|
|BenchmarkWorkerWithConv-8  |  	   19670|	     61669 ns/op|
|BenchmarkWorkerWithConv-8  |  	   19866|	     62316 ns/op|


## License

This project is provided under the [MIT license](LICENSE).