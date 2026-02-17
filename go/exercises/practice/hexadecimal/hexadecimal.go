package hexadecimal

import (
	"fmt"
	"strings"
)

func ParseHex(s string) (int64, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("syntax error: empty string")
	}

	if len(s) > 16 {
		return 0, fmt.Errorf("range error: too many digits")
	}

	var result int64
	for i, ch := range s {
		var d int64
		switch {
		case ch >= '0' && ch <= '9':
			d = int64(ch - '0')
		case ch >= 'a' && ch <= 'f':
			d = int64(ch-'a') + 10
		case ch >= 'A' && ch <= 'F':
			d = int64(ch-'A') + 10
		default:
			return 0, fmt.Errorf("syntax error: invalid character '%c'", ch)
		}
		if i == 0 && len(s) == 16 && d >= 8 {
			return 0, fmt.Errorf("range error: overflow")
		}
		result = result*16 + d
	}
	return result, nil
}

func HandleErrors(inputs []string) []string {
	results := make([]string, len(inputs))
	for i, s := range inputs {
		_, err := ParseHex(s)
		if err == nil {
			results[i] = "none"
		} else if strings.Contains(err.Error(), "range") {
			results[i] = "range"
		} else {
			results[i] = "syntax"
		}
	}
	return results
}
