package strings

import (
	"strings"
	"testing"
)

func TestContains(t *testing.T) {

	assertCorrectTest := func(t *testing.T, test, expected bool) {
		t.Helper()
		if test != expected {
			t.Errorf("expected '%t' but received '%t'", expected, test)
		}
	}

	t.Run("search 'es' in 'test'", func(t *testing.T) {
		test := strings.Contains("test", "es")
		expected := true

		assertCorrectTest(t, test, expected)
	})

	t.Run("expected false", func(t *testing.T) {
		test := strings.Contains("test", " ")
		expected := false

		assertCorrectTest(t, test, expected)
	})

}
