package command

import (
	"bytes"
	"strconv"
	"strings"
	"unicode/utf8"

	gloo "github.com/gloo-foo/framework"
	"github.com/gloo-foo/framework/patterns"
)

// counts holds the tallies wc computes in a single pass over the input lines.
type counts struct {
	lines  int
	words  int
	bytes  int
	chars  int
	maxLen int
}

// tally folds the input lines into a counts value. Each line is the content
// between newlines, with the newline already stripped by the stream, so byte
// and char counts exclude line terminators (see COMPATIBILITY.md).
func tally(lines [][]byte) counts {
	var c counts
	for _, line := range lines {
		c.lines++
		c.bytes += len(line)
		c.words += len(bytes.Fields(line))
		c.chars += utf8.RuneCount(line)
		c.maxLen = max(c.maxLen, len(line))
	}
	return c
}

// column pairs a count value with the flag that selects it for display.
type column struct {
	enabled bool
	value   int
}

// render joins the values of the enabled columns with single spaces, matching
// GNU wc's left-to-right field order (lines, words, bytes, chars, max length).
func render(cols []column) []byte {
	fields := make([]string, 0, len(cols))
	for _, col := range cols {
		if col.enabled {
			fields = append(fields, strconv.Itoa(col.value))
		}
	}
	return []byte(strings.Join(fields, " "))
}

// columns builds the ordered column set for the given flags. With no count
// flag set, the GNU default (lines, words, bytes) is shown.
func columns(f flags, c counts) []column {
	showAll := !bool(f.lines) && !bool(f.words) && !bool(f.bytes) &&
		!bool(f.chars) && !bool(f.maxLineLength)
	return []column{
		{enabled: showAll || bool(f.lines), value: c.lines},
		{enabled: showAll || bool(f.words), value: c.words},
		{enabled: showAll || bool(f.bytes), value: c.bytes},
		{enabled: bool(f.chars), value: c.chars},
		{enabled: bool(f.maxLineLength), value: c.maxLen},
	}
}

// Wc counts lines, words, and bytes from the input stream.
// With no flags, all three counts are shown space-separated.
// Pass WcLines, WcWords, or WcBytes to show individual counts.
// Pass WcChars (-m) to count characters (runes).
// Pass WcMaxLineLength (-L) to report the maximum line length in bytes.
func Wc(opts ...any) gloo.Command[[]byte, []byte] {
	f := gloo.NewParameters[gloo.File, flags](opts...).Flags
	return patterns.Aggregate(func(lines [][]byte) ([]byte, error) {
		return render(columns(f, tally(lines))), nil
	})
}
