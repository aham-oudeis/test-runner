package assertion_checker

import (
	"testing"
)

func TestEqualityCheck(t *testing.T) {
	t.Run("test if string values are equal", func(t *testing.T) {
		got := testEquality("hello", "hello", "is equal to")
		want := true
		assertCorrectMessage(t, got, want)
	})

	t.Run("test if strings are not equal", func(t *testing.T) {
		got := testEquality("hello", "hello", "is not equal to")
		want := false
		assertCorrectMessage(t, got, want)
	})

	t.Run("test if numbers are equal", func(t *testing.T) {
		got := testEquality(1, 1, "is equal to")
		want := true
		assertCorrectMessage(t, got, want)
	})

	t.Run("test if numbers are not equal", func(t *testing.T) {
		got := testEquality(1, 1, "is not equal to")
		want := false
		assertCorrectMessage(t, got, want)
	})
}

func TestValidateNumericalComparison(t *testing.T) {
	t.Run("test if number is greater than", func(t *testing.T) {
		got := ValidateNumbericalComparison(2, 1, "is greater than")
		want := true
		assertCorrectMessage(t, got, want)
	})

	t.Run("test if number is not greater than", func(t *testing.T) {
		got := ValidateNumbericalComparison(2, 1, "is not greater than")
		want := false
		assertCorrectMessage(t, got, want)
	})

	t.Run("test if number is less than", func(t *testing.T) {
		got := ValidateNumbericalComparison(1, 2, "is less than")
		want := true
		assertCorrectMessage(t, got, want)
	})

	t.Run("test if number is not less than", func(t *testing.T) {
		got := ValidateNumbericalComparison(1, 2, "is not less than")
		want := false
		assertCorrectMessage(t, got, want)
	})
}

func TestIsString(t *testing.T) {
	t.Run("test if value is a string", func(t *testing.T) {
		got := IsString("hello")
		want := true
		assertCorrectMessage(t, got, want)
	})

	t.Run("test if value is not a string", func(t *testing.T) {
		got := IsString(1)
		want := false
		assertCorrectMessage(t, got, want)
	})
}

func TestIsAssertionPassing(t *testing.T) {
	t.Run("test if actual int 5 is equal to expected string 5", func(t *testing.T) {
		got := IsAssertionPassing(5, "5", "is not equal to")
		want := false
		assertCorrectMessage(t, got, want)
	})
	t.Run("test if int 5 is less than int 6", func(t *testing.T) {
		got := IsAssertionPassing(5, 6, "is greater than")
		want := false
		assertCorrectMessage(t, got, want)
	})

	t.Run("test if string hello is not equal to string HELLO", func(t *testing.T) {
		got := IsAssertionPassing("hello", "HELLO", "is not equal to")
		want := true
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}