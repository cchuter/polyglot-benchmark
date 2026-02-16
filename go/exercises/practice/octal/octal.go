package octal

import "fmt"

func ParseOctal(input string) (int64, error) {
	num := int64(0)
	for _, ch := range input {
		if ch < '0' || ch > '7' {
			return 0, fmt.Errorf("invalid octal digit: %c", ch)
		}
		num = num*8 + int64(ch-'0')
	}
	return num, nil
}
