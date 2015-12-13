package termboxui

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Wraps the text into multiple lines at most lim chars long.
// Returns a list of all the lines.
// Supports unicode.
func WrapText(text string, lim int) []string {
	slice := []byte(text)
	lines := make([]string, 0, 2)

	curLine := ""

	for len(slice) > 0 {
		length := lenWord(slice)
		if len(curLine)+length <= lim {
			curLine += string(slice[:length])
		} else if len(curLine) > 0 {
			lines = append(lines, strings.TrimSpace(curLine))
			curLine = string(slice[:length])
		} else {
			tmp := breakWord(string(slice[:length]), lim)
			curLine = tmp[len(tmp)-1]
			lines = append(lines, tmp[:len(tmp)-1]...)
		}
		slice = slice[length:]
	}
	if len(curLine) > 0 {
		lines = append(lines, strings.TrimSpace(curLine))
	}

	return lines
}

func lenWord(text []byte) int {
	var size, length int
	var r rune

	for i := 0; i < len(text); i += size {
		r, size = utf8.DecodeRune(text[i:])
		if unicode.IsSpace(r) {
			return length + size
		}
		length += size
	}
	return length
}

func wordAt(text []byte, at int) (start, end int) {
	start, end = -1, -1

	// Iterate backwords to find the start of the word
	for i := at - 1; i > 0; i-- {
		if !utf8.RuneStart(text[i]) {
			continue
		}
		r, size := utf8.DecodeRune(text[i:])
		if unicode.IsSpace(r) {
			start = i + size
		}
	}

	//Iterate forwards to find end of word
	r, size := utf8.DecodeRune(text[at:])
	last := at

	for i := at + size; i < len(text); i += size {
		r, size = utf8.DecodeRune(text[i:])
		if unicode.IsSpace(r) {
			end = last
		}
		last = i + size
	}

	return
}

// Returns a string representing the word broken onto multiple lines
// and the number of characters on the last line
func breakWord(word string, lim int) []string {
	var lines []string

	for len(word) > lim {
		lines = append(lines, word[:lim-1]+"-")
		word = word[lim-1:]
	}
	if len(word) > 0 {
		lines = append(lines, word)
	}
	return lines
}
