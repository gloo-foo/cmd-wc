package command_test

import (
	"testing"

	command "github.com/gloo-foo/cmd-wc"

	"github.com/gloo-foo/testable"
	"github.com/gloo-foo/testable/assertion"
)

func TestWc_Default_AllCounts(t *testing.T) {
	// "hello world" = 11 bytes, 2 words
	// "foo bar"     = 7 bytes, 2 words
	// 2 lines, 4 words, 18 bytes
	lines, err := testable.TestLines(command.Wc(), "hello world\nfoo bar\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"2 4 18"})
}

func TestWc_Default_SingleLine(t *testing.T) {
	// "hello" = 5 bytes, 1 word, 1 line
	lines, err := testable.TestLines(command.Wc(), "hello\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"1 1 5"})
}

func TestWc_Default_EmptyInput(t *testing.T) {
	lines, err := testable.TestLines(command.Wc(), "")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"0 0 0"})
}

func TestWc_LinesOnly(t *testing.T) {
	lines, err := testable.TestLines(command.Wc(command.WcLines), "a\nb\nc\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"3"})
}

func TestWc_LinesOnly_Empty(t *testing.T) {
	lines, err := testable.TestLines(command.Wc(command.WcLines), "")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"0"})
}

func TestWc_WordsOnly(t *testing.T) {
	lines, err := testable.TestLines(command.Wc(command.WcWords), "hello world\nfoo bar baz\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"5"})
}

func TestWc_WordsOnly_WhitespaceOnly(t *testing.T) {
	lines, err := testable.TestLines(command.Wc(command.WcWords), "   \n\t\t\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"0"})
}

func TestWc_WordsOnly_ExtraSpaces(t *testing.T) {
	lines, err := testable.TestLines(command.Wc(command.WcWords), "  hello   world  \n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"2"})
}

func TestWc_BytesOnly(t *testing.T) {
	// "hello" = 5 bytes (newline stripped by stream)
	lines, err := testable.TestLines(command.Wc(command.WcBytes), "hello\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"5"})
}

func TestWc_BytesOnly_MultiLine(t *testing.T) {
	// "abc" = 3, "def" = 3 => 6 bytes total
	lines, err := testable.TestLines(command.Wc(command.WcBytes), "abc\ndef\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"6"})
}

func TestWc_BytesOnly_Empty(t *testing.T) {
	lines, err := testable.TestLines(command.Wc(command.WcBytes), "")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"0"})
}

func TestWc_MultiLine(t *testing.T) {
	input := "one two three\nfour five\nsix\n"
	// "one two three" = 13, "four five" = 9, "six" = 3 => 25 bytes
	// 3 lines, 6 words, 25 bytes
	lines, err := testable.TestLines(command.Wc(), input)
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"3 6 25"})
}

func TestWc_EmptyLines(t *testing.T) {
	// Three empty lines: each is "" after newline stripping
	lines, err := testable.TestLines(command.Wc(command.WcLines), "\n\n\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"3"})
}

func TestWc_EmptyLines_Bytes(t *testing.T) {
	// Three empty lines: each line content is "" => 0 bytes total
	lines, err := testable.TestLines(command.Wc(command.WcBytes), "\n\n\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"0"})
}

func TestWc_Unicode_Bytes(t *testing.T) {
	// "日本語" in UTF-8 = 9 bytes
	lines, err := testable.TestLines(command.Wc(command.WcBytes), "日本語\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"9"})
}

func TestWc_Unicode_Words(t *testing.T) {
	// "こんにちは 世界" = 2 words (space-separated)
	lines, err := testable.TestLines(command.Wc(command.WcWords), "こんにちは 世界\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"2"})
}

func TestWc_TabSeparated(t *testing.T) {
	lines, err := testable.TestLines(command.Wc(command.WcWords), "hello\tworld\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"2"})
}

func TestWc_CharsOnly(t *testing.T) {
	// "日本語" = 3 runes, 9 bytes
	lines, err := testable.TestLines(command.Wc(command.WcChars), "日本語\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"3"})
}

func TestWc_CharsASCII(t *testing.T) {
	// ASCII: rune count == byte count
	lines, err := testable.TestLines(command.Wc(command.WcChars), "hello\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"5"})
}

func TestWc_CharsMultiLine(t *testing.T) {
	// "abc" = 3 runes, "日本" = 2 runes => 5 total
	lines, err := testable.TestLines(command.Wc(command.WcChars), "abc\n日本\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"5"})
}

func TestWc_CharsEmpty(t *testing.T) {
	lines, err := testable.TestLines(command.Wc(command.WcChars), "")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"0"})
}

func TestWc_MaxLineLength(t *testing.T) {
	// "hello world" = 11 bytes, "foo" = 3 bytes => max 11
	lines, err := testable.TestLines(command.Wc(command.WcMaxLineLength), "hello world\nfoo\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"11"})
}

func TestWc_MaxLineLengthSingleLine(t *testing.T) {
	lines, err := testable.TestLines(command.Wc(command.WcMaxLineLength), "abcde\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"5"})
}

func TestWc_MaxLineLengthEmpty(t *testing.T) {
	lines, err := testable.TestLines(command.Wc(command.WcMaxLineLength), "")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"0"})
}

func TestWc_MaxLineLengthWithEmptyLines(t *testing.T) {
	// "abc" = 3, "" = 0, "defgh" = 5 => max 5
	lines, err := testable.TestLines(command.Wc(command.WcMaxLineLength), "abc\n\ndefgh\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"5"})
}

func TestWc_CharsAndBytes(t *testing.T) {
	// "日本語" = 3 runes, 9 bytes
	lines, err := testable.TestLines(command.Wc(command.WcBytes, command.WcChars), "日本語\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"9 3"})
}

func TestWc_LinesAndMaxLineLength(t *testing.T) {
	lines, err := testable.TestLines(command.Wc(command.WcLines, command.WcMaxLineLength), "abc\ndefgh\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"2 5"})
}

func TestWc_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		opts     []any
		input    string
		expected string
	}{
		{"default single word", nil, "hello\n", "1 1 5"},
		{"lines three", []any{command.WcLines}, "a\nb\nc\n", "3"},
		{"words five", []any{command.WcWords}, "one two\nthree four five\n", "5"},
		{"bytes abc", []any{command.WcBytes}, "abc\n", "3"},
		{"default empty", nil, "", "0 0 0"},
		{"lines empty", []any{command.WcLines}, "", "0"},
		{"words empty", []any{command.WcWords}, "", "0"},
		{"bytes empty", []any{command.WcBytes}, "", "0"},
		{"chars unicode", []any{command.WcChars}, "cafe\u0301\n", "5"},
		{"max line length", []any{command.WcMaxLineLength}, "short\nlonger line\n", "11"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := testable.TestLines(command.Wc(tt.opts...), tt.input)
			assertion.NoError(t, err)
			assertion.Lines(t, lines, []string{tt.expected})
		})
	}
}
