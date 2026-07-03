package wc_test

import (
	"fmt"

	"github.com/gloo-foo/testable"

	wc "github.com/gloo-foo/cmd-wc"
)

func ExampleWc_basic() {
	// echo "Hello World\nSecond line" | wc
	output, _ := testable.Test(wc.Wc(), "Hello World\nSecond line\n")
	fmt.Print(output)
	// Output:
	// 2 4 22
}
