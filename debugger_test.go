// Important! Run tests/benchmarks with -tags=debugger
// to include the actual labeling code

package debugger_test

import (
	"strconv"
	"testing"

	"github.com/dlsniper/debugger"
)

var result int

func workerWithout(n int) int {
	if n != 1 {
		return n + workerWithout(n-1)
	}

	return 1
}

func workerWithOne(n int) int {
	debugger.SetLabels(func() []string {
		return []string{
			"label1", "label1value",
		}
	})

	if n != 1 {
		return n + workerWithOne(n-1)
	}

	return 1
}

func workerWithThree(n int) int {
	debugger.SetLabels(func() []string {
		return []string{
			"label1", "label1value",
			"label2", "label2value",
			"label3", "label3value",
		}
	})

	if n != 1 {
		return n + workerWithThree(n-1)
	}

	return 1
}

func workerWithTen(n int) int {
	debugger.SetLabels(func() []string {
		return []string{
			"label1", "label1value",
			"label2", "label2value",
			"label3", "label3value",
			"label4", "label41value",
			"label5", "label5value",
			"label6", "label6value",
			"label7", "label7value",
			"label8", "label8value",
			"label9", "label9value",
			"label10", "label10value",
		}
	})

	if n != 1 {
		return n + workerWithTen(n-1)
	}

	return 1
}

func workerWithConv(n int) int {
	debugger.SetLabels(func() []string {
		return []string{
			"label1", "label1value",
			"label2", "label2value",
			"label3", strconv.Itoa(n),
		}
	})

	if n != 1 {
		return n + workerWithConv(n-1)
	}

	return 1
}

func BenchmarkWorkerWithout(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = workerWithout(100)
	}
	result = res
}

func BenchmarkWorkerWithOne(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = workerWithOne(100)
	}
	result = res
}

func BenchmarkWorkerWithThree(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = workerWithThree(100)
	}
	result = res
}

func BenchmarkWorkerWithTen(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = workerWithTen(100)
	}
	result = res
}

func BenchmarkWorkerWithConv(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = workerWithConv(100)
	}
	result = res
}

func init() {
	// force Go to include our variable in the results and not optimize any code based on it
	println(result)
}
