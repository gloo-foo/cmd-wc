package alias_test

import (
	"slices"
	"testing"

	wc "github.com/gloo-foo/cmd-wc/alias"
	"github.com/gloo-foo/testable"
)

// The alias package re-exports the constructor and flag constants under
// unprefixed names. A mis-wired re-export (say, Words bound to the disabled
// constant, or Wc bound to the wrong function) compiles cleanly, so only
// behavior can prove the wiring. Each test exercises one re-export and asserts
// the wc output it must produce.

// countInput has 2 lines, 4 words, 21 bytes (newlines stripped by the stream),
// and a longest line of 11 bytes. "日本" adds multi-byte runes so the char
// count (17) differs from the byte count (21).
const countInput = "hello world\nfoo 日本\n"

func assertLines(t *testing.T, got, want []string) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestAlias_DefaultShowsLinesWordsBytes(t *testing.T) {
	lines, err := testable.TestLines(wc.Wc(), countInput)
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"2 4 21"})
}

func TestAlias_LinesCountsLines(t *testing.T) {
	lines, err := testable.TestLines(wc.Wc(wc.Lines), countInput)
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"2"})
}

func TestAlias_WordsCountsWords(t *testing.T) {
	lines, err := testable.TestLines(wc.Wc(wc.Words), countInput)
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"4"})
}

func TestAlias_BytesCountsBytes(t *testing.T) {
	lines, err := testable.TestLines(wc.Wc(wc.Bytes), countInput)
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"21"})
}

func TestAlias_CharsCountsRunes(t *testing.T) {
	// "日本" is 6 bytes but 2 runes, so chars (17) < bytes (21).
	lines, err := testable.TestLines(wc.Wc(wc.Chars), countInput)
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"17"})
}

func TestAlias_MaxLineLengthReportsLongest(t *testing.T) {
	// "hello world" is 11 bytes; "foo 日本" is 10 bytes => 11.
	lines, err := testable.TestLines(wc.Wc(wc.MaxLineLength), countInput)
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"11"})
}

func TestAlias_DisabledFlagsMatchDefault(t *testing.T) {
	// The No* constants are the disabled forms: passing every one must behave
	// exactly like passing no flag at all (the lines/words/bytes default).
	lines, err := testable.TestLines(wc.Wc(
		wc.NoLines, wc.NoWords, wc.NoBytes, wc.NoChars, wc.NoMaxLineLength,
	), countInput)
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"2 4 21"})
}
