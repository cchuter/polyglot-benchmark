package forth

import (
	"errors"
	"strconv"
	"strings"
)

var (
	errInsufficientOperands = errors.New("insufficient operands")
	errDivisionByZero       = errors.New("division by zero")
	errUndefinedWord        = errors.New("undefined word")
	errIllegalOperation     = errors.New("illegal operation")
)

// Forth evaluates a sequence of Forth phrases and returns the resulting stack.
func Forth(input []string) ([]int, error) {
	stack := []int{}
	defs := map[string][]string{}

	for _, phrase := range input {
		tokens := strings.Fields(phrase)
		if err := eval(tokens, &stack, defs); err != nil {
			return nil, err
		}
	}
	return stack, nil
}

func eval(tokens []string, stack *[]int, defs map[string][]string) error {
	for i := 0; i < len(tokens); i++ {
		token := strings.ToUpper(tokens[i])

		if token == ":" {
			// Parse definition: : word body ;
			i++
			if i >= len(tokens) {
				return errIllegalOperation
			}
			word := strings.ToUpper(tokens[i])

			// Cannot redefine numbers
			if _, err := strconv.Atoi(word); err == nil {
				return errIllegalOperation
			}

			// Find the closing ;
			i++
			start := i
			for i < len(tokens) && strings.ToUpper(tokens[i]) != ";" {
				i++
			}
			if i >= len(tokens) {
				return errIllegalOperation
			}

			// Expand body tokens using current defs (snapshot semantics)
			var body []string
			for _, t := range tokens[start:i] {
				t = strings.ToUpper(t)
				if expanded, ok := defs[t]; ok {
					body = append(body, expanded...)
				} else {
					body = append(body, t)
				}
			}
			defs[word] = body

		} else if expanded, ok := defs[token]; ok {
			if err := eval(expanded, stack, defs); err != nil {
				return err
			}
		} else if n, err := strconv.Atoi(token); err == nil {
			*stack = append(*stack, n)
		} else {
			if err := execBuiltin(token, stack); err != nil {
				return err
			}
		}
	}
	return nil
}

func execBuiltin(op string, stack *[]int) error {
	switch op {
	case "+", "-", "*", "/":
		if len(*stack) < 2 {
			return errInsufficientOperands
		}
		b := pop(stack)
		a := pop(stack)
		switch op {
		case "+":
			push(stack, a+b)
		case "-":
			push(stack, a-b)
		case "*":
			push(stack, a*b)
		case "/":
			if b == 0 {
				return errDivisionByZero
			}
			push(stack, a/b)
		}
	case "DUP":
		if len(*stack) < 1 {
			return errInsufficientOperands
		}
		push(stack, (*stack)[len(*stack)-1])
	case "DROP":
		if len(*stack) < 1 {
			return errInsufficientOperands
		}
		pop(stack)
	case "SWAP":
		if len(*stack) < 2 {
			return errInsufficientOperands
		}
		n := len(*stack)
		(*stack)[n-1], (*stack)[n-2] = (*stack)[n-2], (*stack)[n-1]
	case "OVER":
		if len(*stack) < 2 {
			return errInsufficientOperands
		}
		push(stack, (*stack)[len(*stack)-2])
	default:
		return errUndefinedWord
	}
	return nil
}

func push(stack *[]int, v int) {
	*stack = append(*stack, v)
}

func pop(stack *[]int) int {
	n := len(*stack)
	v := (*stack)[n-1]
	*stack = (*stack)[:n-1]
	return v
}
