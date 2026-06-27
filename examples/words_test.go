package wc_test

import (
	"fmt"

	. "github.com/gloo-foo/cmd-wc"
	"github.com/gloo-foo/testable"
)

func ExampleWc_words() {
	// echo "one two three four" | wc -w
	output, _ := testable.Test(Wc(WcWords), "one two three four\n")
	fmt.Print(output)
	// Output:
	// 4
}
