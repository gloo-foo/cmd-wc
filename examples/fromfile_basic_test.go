package wc_test

import (
	"fmt"

	. "github.com/gloo-foo/cmd-wc"
	"github.com/gloo-foo/testable"
)

func ExampleWc_fromFile_basic() {
	// cat testdata/sample.txt | wc
	output, _ := testable.Test(Wc(), "Hello World\nThis is a test file\nWith three lines\n")
	fmt.Print(output)
	// Output:
	// 3 10 46
}
