package wc_test

import (
	"fmt"

	. "github.com/gloo-foo/cmd-wc"
	"github.com/gloo-foo/testable"
)

func ExampleWc_lines() {
	// echo "line1\nline2\nline3" | wc -l
	output, _ := testable.Test(Wc(WcLines), "line1\nline2\nline3\n")
	fmt.Print(output)
	// Output:
	// 3
}
