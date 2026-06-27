package wc_test

import (
	"fmt"

	. "github.com/gloo-foo/cmd-wc"
	"github.com/gloo-foo/testable"
)

func ExampleWc_basic() {
	// echo "Hello World\nSecond line" | wc
	output, _ := testable.Test(Wc(), "Hello World\nSecond line\n")
	fmt.Print(output)
	// Output:
	// 2 4 22
}
