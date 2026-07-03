package command

// wcLinesFlag enables line counting (-l flag).
type wcLinesFlag bool

const (
	WcLines   wcLinesFlag = true
	WcNoLines wcLinesFlag = false
)

// wcWordsFlag enables word counting (-w flag).
type wcWordsFlag bool

const (
	WcWords   wcWordsFlag = true
	WcNoWords wcWordsFlag = false
)

// wcBytesFlag enables byte counting (-c flag).
type wcBytesFlag bool

const (
	WcBytes   wcBytesFlag = true
	WcNoBytes wcBytesFlag = false
)

// wcCharsFlag enables character (rune) counting (-m flag).
type wcCharsFlag bool

const (
	WcChars   wcCharsFlag = true
	WcNoChars wcCharsFlag = false
)

// wcMaxLineLengthFlag enables max line length reporting (-L flag).
type wcMaxLineLengthFlag bool

const (
	WcMaxLineLength   wcMaxLineLengthFlag = true
	WcNoMaxLineLength wcMaxLineLengthFlag = false
)

type flags struct {
	linesEnabled         wcLinesFlag
	wordsEnabled         wcWordsFlag
	bytesEnabled         wcBytesFlag
	charsEnabled         wcCharsFlag
	maxLineLengthEnabled wcMaxLineLengthFlag
}

// fold partitions opts: wc's own option values are folded into the flag set,
// and every other argument is passed through unchanged for the framework's
// positional classifier.
func fold(opts []any) (flags, []any) {
	var f flags
	rest := make([]any, 0, len(opts))
	for _, o := range opts {
		switch v := o.(type) {
		case wcLinesFlag:
			f.linesEnabled = v
		case wcWordsFlag:
			f.wordsEnabled = v
		case wcBytesFlag:
			f.bytesEnabled = v
		case wcCharsFlag:
			f.charsEnabled = v
		case wcMaxLineLengthFlag:
			f.maxLineLengthEnabled = v
		default:
			rest = append(rest, o)
		}
	}
	return f, rest
}
