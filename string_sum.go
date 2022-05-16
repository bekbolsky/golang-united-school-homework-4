package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// ParseOperands parses the input string and returns the first and second operands
func ParseOperands(input, operator string) (int, int, error) {
	operands := strings.Split(input, operator)
	if len(operands) != 2 {
		return 0, 0, fmt.Errorf("%w", errorNotTwoOperands)
	}

	firstOperand, err := strconv.Atoi(operands[0])
	if err != nil {
		return 0, 0, fmt.Errorf("%w", err)
	}

	secondOperand, err := strconv.Atoi(operands[1])
	if err != nil {
		return 0, 0, fmt.Errorf("%w", err)
	}

	return firstOperand, secondOperand, nil
}

// StringSum computes the sum of two int numbers written as a string.
// For example, having an input string "3+5", it returns output string "8" and nil error
// Considers cases, when operands are negative ("-3+5" or "-3-5") and
// when input string contains whitespace (" 3 + 5 ")
//
// For the cases, when the input expression is not valid(contains characters,
// that are not numbers, +, - or whitespace) the function returns an empty string and
// an appropriate error from strconv package wrapped into own error
// with fmt.Errorf function
//
// Uses the errors defined above as described, again wrapping into fmt.Errorf
func StringSum(input string) (output string, err error) {
	if len(input) == 0 || strings.ReplaceAll(input, " ", "") == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, " ", "")

	if strings.HasPrefix(input, "-") {
		if strings.Contains(input[1:], "+") {
			firstOperand, secondOperand, err := ParseOperands(input[1:], "+")
			if err != nil {
				return "", err
			}
			output = strconv.Itoa(-1*firstOperand + secondOperand)
			return output, nil
		} else if strings.Contains(input[1:], "-") {
			firstOperand, secondOperand, err := ParseOperands(input[1:], "-")
			if err != nil {
				return "", err
			}
			output = strconv.Itoa(-1*firstOperand - secondOperand)
			return output, nil
		}
	} else {
		if strings.Contains(input, "+") {
			firstOperand, secondOperand, err := ParseOperands(input, "+")
			if err != nil {
				return "", err
			}
			output = strconv.Itoa(firstOperand + secondOperand)
			return output, nil
		} else if strings.Contains(input, "-") {
			firstOperand, secondOperand, err := ParseOperands(input, "-")
			if err != nil {
				return "", err
			}
			output = strconv.Itoa(firstOperand - secondOperand)
			return output, nil
		}
	}
	return "", fmt.Errorf("%w", errorNotTwoOperands)
}
