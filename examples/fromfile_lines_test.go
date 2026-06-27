package wc_test

import (
	"fmt"

	. "github.com/gloo-foo/cmd-wc"
	"github.com/gloo-foo/testable"
)

func ExampleWc_fromFile_lines() {
	// cat testdata/sample.txt | wc -l
	output, _ := testable.Test(Wc(WcLines), "Hello World\nThis is a test file\nWith three lines\n")
	fmt.Print(output)
	// Output:
	// 3
}
