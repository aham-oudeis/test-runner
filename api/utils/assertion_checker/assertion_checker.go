package assertion_checker

import (
	"strconv"
	"strings"
	"test-runner/api/entities"
	"test-runner/api/utils/value_interpolator"
)

func IsString(b any) bool {
	_, ok := b.(string)
	return ok
}

func IsInt(b any) bool {
	_, ok := b.(int)
	return ok
}

//the basic function is to check if the two values are as specified by the operator
//T comparable makes sure that the type of the two values are comparable
//the function expects both the actual and expected to be of the same type
func testEquality[T comparable](actual, expected T, operator string) bool {
	switch operator {
	case "is equal to":
		return actual == expected
	case "is not equal to":
		return actual != expected
	default:
		return false
	}
}

func ValidateNumbericalComparison(actual, expected int, operator string) bool {
	switch operator {
	case "is greater than":
		return actual > expected
	case "is not greater than":
		return actual <= expected
	case "is less than":
		return actual < expected
	case "is not less than":
		return actual >= expected
	default:
		return false
	}
}

func IsAssertionOfNumbersPassing(actual, expected int, operator string) bool {
	if strings.Contains(operator, "equal") {
		return testEquality(actual, expected, operator)
	} else if strings.Contains(operator, "greater") || strings.Contains(operator, "less") {
		return ValidateNumbericalComparison(actual, expected, operator)
	} else {
		return false
	}
}

func IsAssertionOfStringsPassing(actual, expected, operator string) bool {
	if strings.Contains(operator, "equal") {
		return testEquality(actual, expected, operator)
	} else if strings.Contains(operator, "contains") {
		return strings.Contains(actual, expected)
	} else {
		return false
	}
}

func IsAssertionPassing(actual, expected any, operator string) bool {
	if IsString(actual) && IsString(expected) {
		return IsAssertionOfStringsPassing(actual.(string), expected.(string), operator)
	} else if IsInt(actual) && IsString(expected) {
		rightOperand, err := strconv.Atoi(expected.(string))
		if err != nil {
			return false
		}
		return IsAssertionOfNumbersPassing(actual.(int), rightOperand, operator)
	} else if IsInt(actual) && IsInt(expected) {
		return IsAssertionOfNumbersPassing(actual.(int), expected.(int), operator)
	} else {
		return false
	}
}

func IsAssertionValid(assertion entities.Assertion, response map[string]interface{}) (bool, any, error) {
	actual, err := value_interpolator.InterpolateValue(assertion.Property, response)

	if err != nil {
		return false, actual, err
	}

	return IsAssertionPassing(actual, assertion.Expected, assertion.Comparison), actual, nil
}
