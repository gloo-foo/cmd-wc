package wc_test

import (
	"fmt"

	"github.com/gloo-foo/testable"

	wc "github.com/gloo-foo/cmd-wc"
)

func ExampleWc_words() {
	// echo "one two three four" | wc -w
	output, _ := testable.Test(wc.Wc(wc.WcWords), "one two three four\n")
	fmt.Print(output)
	// Output:
	// 4
}
