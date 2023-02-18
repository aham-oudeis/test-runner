package value_interpolator

import (
	"testing"
)

func TestInterpolateValue(t *testing.T) {
	t.Run("test interpolating a reference from a nested object", func(t *testing.T) {
		response := map[string]interface{}{};
		response["body"] = map[string]interface{}{};
		response["body"].(map[string]interface{})["country"] = "India"
		got, _ := InterpolateValue("body.country", response)
		want := "India"
		assertCorrectMessage(t, got.(string), want)
	})

	t.Run("test interpolating a reference from a nested object", func(t *testing.T) {
		response := map[string]interface{}{};
		response["body"] = map[string]interface{}{};
		response["body"].(map[string]interface{})["country"] = []string{"India", "USA"}
		got, _ := InterpolateValue("body.country[0]", response)
		want := "India"
		assertCorrectMessage(t, got.(string), want)
	})
}

func TestSplitValues(t *testing.T) {
	t.Run("test splitting string by regex", func(t *testing.T) {
		got := SplitStringByRegex("foo.bar[0]", "\\.|\\[|\\]")
		want := []string{"foo", "bar", "0"}
		assertCorrectMessage(t, got[0], want[0])
		assertCorrectMessage(t, got[1], want[1])
		assertCorrectMessage(t, got[2], want[2])
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}