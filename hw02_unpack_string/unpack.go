package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")
var ErrinvalidSimbols = errors.New("invalid simbols")

func Unpack(stroka string) (string, error) {
	var NewLine strings.Builder
	massive := []rune(stroka)
	maxLen := len(massive)
	if len(massive) == 0 {
		return "", nil
	}
	if unicode.IsDigit(massive[0]) {
		return "", ErrInvalidString
	}
	var con bool = false
	for i, k := range massive {
		if !unicode.IsDigit(k) && !unicode.IsLetter(k) && string(massive[i]) != "\\" {
			return "", ErrinvalidSimbols
		}
		if con {
			con = false
			continue
		}
		if string(massive[i]) == "\\" {
			NewLine.WriteString(string(massive[i+1]))
			con = true
			continue
		}
		if unicode.IsDigit(k) && int(k-'0') != 0 {
			if i < maxLen-1 && unicode.IsDigit(massive[i+1]) {
				return "", ErrInvalidString
			}
			letter := string(massive[i-1])
			repeatLetters := strings.Repeat(letter, int(k-'0'-1))
			NewLine.WriteString(repeatLetters)
			continue
		}
		if i != maxLen-1 && string(massive[i+1]) == "0" {
			con = true
			continue
		} else {
			NewLine.WriteString(string(massive[i]))
		}
	}
	ResNewLine := NewLine.String()
	return ResNewLine, nil
}
