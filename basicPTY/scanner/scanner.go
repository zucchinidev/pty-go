package scanner

import (
	"errors"
	"unicode"
	"unicode/utf8"
)

var ErrClosingQuote = errors.New("Missing closing quote")

func isQuote(r rune) bool {
	return r == '"' || r == '\''
}

func ScanArgs(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start, first := 0, rune(0)

	// searching the beginning of the parameter, repeat for each one
	for width := 0; start < len(data); start += width {
		first, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(first) {
			break
		}
	}

	// if the first rune is a quote, add one to the beginning index
	if isQuote(first) {
		start++
	}

	// searching the end of the parameter, repeat for each one
	width := 0 // width in bytes of the utf-8 found element
	i := start // the index to start the searching

	firstElementIsQuote := isQuote(first)
	for ; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		theSecondOneIsAlsoAnotherQuote := firstElementIsQuote && r == first
		if !firstElementIsQuote && unicode.IsSpace(r) || theSecondOneIsAlsoAnotherQuote {
			return i + width, data[start:i], nil
		}
	}

	// is the end and we have not reached the last part of the token
	if atEOF && len(data) > start {
		if firstElementIsQuote {
			err = ErrClosingQuote
		}
		return len(data), data[start:], err
	}

	// the first element is a quote but there is nothing in between
	if firstElementIsQuote {
		start--
	}
	return start, nil, nil
}
