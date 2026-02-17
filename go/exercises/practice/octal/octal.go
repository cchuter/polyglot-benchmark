package octal

import "fmt"

func ParseOctal(input string) (int64, error) {
	num := int64(0)
	for _, r := range input {
		if r < '0' || r > '7' {
			return 0, fmt.Errorf("invalid octal digit: %c", r)
		}
		num = num<<3 + int64(r-'0')
	}
	return num, nil
}
