package hexadecimal

import (
	"errors"
	"math"
)

var ErrRange = errors.New("value out of range")
var ErrSyntax = errors.New("invalid syntax")

type ParseError struct {
	Num string
	Err error
}

func (e *ParseError) Error() string {
	return "hexadecimal.ParseHex: parsing \"" + e.Num + "\": " + e.Err.Error()
}

func ParseHex(s string) (n int64, err error) {
	if len(s) < 1 {
		err = ErrSyntax
		goto Error
	}

	for i := 0; i < len(s); i++ {
		d := s[i]
		var v byte
		switch {
		case '0' <= d && d <= '9':
			v = d - '0'
		case 'a' <= d && d <= 'f':
			v = d - 'a' + 10
		case 'A' <= d && d <= 'F':
			v = d - 'A' + 10
		default:
			n = 0
			err = ErrSyntax
			goto Error
		}

		if n >= math.MaxInt64/16+1 {
			n = math.MaxInt64
			err = ErrRange
			goto Error
		}

		n *= 16
		n1 := n + int64(v)

		if n1 < n {
			n = math.MaxInt64
			err = ErrRange
			goto Error
		}
		n = n1
	}
	return n, nil

Error:
	return n, &ParseError{s, err}
}

func HandleErrors(tests []string) []string {
	e := make([]string, len(tests))
	for i, s := range tests {
		_, err := ParseHex(s)
		switch pe, ok := err.(*ParseError); {
		case err == nil:
			e[i] = "none"
		case ok && pe.Err == ErrSyntax:
			e[i] = "syntax"
		case ok && pe.Err == ErrRange:
			e[i] = "range"
		}
	}
	return e
}
