// Package alias provides unprefixed type aliases for wc command flags.
//
//	import wc "github.com/gloo-foo/cmd-wc/alias"
//	wc.Wc(wc.Lines)
package alias

import command "github.com/gloo-foo/cmd-wc"

// Wc re-exports the constructor.
var Wc = command.Wc

// -l flag: count lines
const Lines = command.WcLines

// default: don't count lines
const NoLines = command.WcNoLines

// -w flag: count words
const Words = command.WcWords

// default: don't count words
const NoWords = command.WcNoWords

// -c flag: count bytes
const Bytes = command.WcBytes

// default: don't count bytes
const NoBytes = command.WcNoBytes

// -m flag: count characters (runes)
const Chars = command.WcChars

// default: don't count characters
const NoChars = command.WcNoChars

// -L flag: report the maximum line length
const MaxLineLength = command.WcMaxLineLength

// default: don't report the maximum line length
const NoMaxLineLength = command.WcNoMaxLineLength
