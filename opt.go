package command

type wcLinesFlag bool

const (
	WcLines   wcLinesFlag = true
	WcNoLines wcLinesFlag = false
)

type wcWordsFlag bool

const (
	WcWords   wcWordsFlag = true
	WcNoWords wcWordsFlag = false
)

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
	lines         wcLinesFlag
	words         wcWordsFlag
	bytes         wcBytesFlag
	chars         wcCharsFlag
	maxLineLength wcMaxLineLengthFlag
}

func (f wcLinesFlag) Configure(flags *flags)         { flags.lines = f }
func (f wcWordsFlag) Configure(flags *flags)         { flags.words = f }
func (f wcBytesFlag) Configure(flags *flags)         { flags.bytes = f }
func (f wcCharsFlag) Configure(flags *flags)         { flags.chars = f }
func (f wcMaxLineLengthFlag) Configure(flags *flags) { flags.maxLineLength = f }
