package wc_test

import (
	"fmt"

	"github.com/gloo-foo/testable"

	wc "github.com/gloo-foo/cmd-wc"
)

func ExampleWc_lines() {
	// echo "line1\nline2\nline3" | wc -l
	output, _ := testable.Test(wc.Wc(wc.WcLines), "line1\nline2\nline3\n")
	fmt.Print(output)
	// Output:
	// 3
}
