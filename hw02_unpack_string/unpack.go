package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var (
	ErrInvalidString  = errors.New("invalid string")
	ErrinvalidSimbols = errors.New("invalid simbols")
)

func Unpack(stroka string) (string, error) {
	massive := []rune(stroka)
	if len(massive) == 0 {
		return "", nil
	}
	if unicode.IsDigit(massive[0]) {
		return "", ErrInvalidString
	}
	ResNewLine, err := Iteration(massive)
	if err != nil {
		return "", err
	}
	return ResNewLine, nil
}

func Iteration(massive []rune) (string, error) {
	var NewLine strings.Builder
	maxLen := len(massive)
	var con bool
	for i, k := range massive {
		if TestSimbols(k) {
			return "", ErrinvalidSimbols
		}
		if con {
			con = false
			continue
		}
		if string(massive[i]) == "\\" {
			switch i < maxLen-1 {
			case true:
				if unicode.IsLetter(massive[i+1]) {
					return "", ErrInvalidString
				}
				NewLine.WriteString(string(massive[i+1]))
				con = true
				continue
			default:
				return "", ErrInvalidString
			}
		}
		if unicode.IsDigit(k) {
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

func TestSimbols(letter rune) bool {
	if !unicode.IsDigit(letter) && !unicode.IsLetter(letter) && string(letter) != "\\" && string(letter) != "\n" {
		return true
	}
	return false
}
