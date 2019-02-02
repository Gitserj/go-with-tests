package strings

import (
	"strings"
	"testing"
)

func TestCount(t *testing.T) {

	assertCorrectTest := func(t *testing.T, test, expected int) {
		t.Helper()
		if test != expected {
			t.Errorf("expected '%d' but received '%d'", expected, test)
		}
	}

	t.Run("counting 's' in 'test'", func(t *testing.T) {
		test := strings.Count("test", "s")
		expected := 1

		assertCorrectTest(t, test, expected)
	})

	t.Run("counting 't' in 'test'", func(t *testing.T) {
		test := strings.Count("test", "t")
		expected := 2

		assertCorrectTest(t, test, expected)
	})

	t.Run("counting 'e' in 'test'", func(t *testing.T) {
		test := strings.Count("test", "e")
		expected := 1

		assertCorrectTest(t, test, expected)
	})

	t.Run("counting 'E' in 'test'", func(t *testing.T) {
		test := strings.Count("test", "e")
		expected := 1

		assertCorrectTest(t, test, expected)
	})

	t.Run("counting 'd' in 'test'", func(t *testing.T) {
		test := strings.Count("test", "d")
		expected := 0

		assertCorrectTest(t, test, expected)
	})

	t.Run("counting '' in 'test'", func(t *testing.T) {
		test := strings.Count("test", "")
		expected := 5

		assertCorrectTest(t, test, expected)
	})

	t.Run("counting 'e' in ''", func(t *testing.T) {
		test := strings.Count("", "e")
		expected := 0

		assertCorrectTest(t, test, expected)
	})

	t.Run("counting 'es' in 'test'", func(t *testing.T) {
		test := strings.Count("test", "es")
		expected := 1

		assertCorrectTest(t, test, expected)
	})

	t.Run("counting 'test' in 'test'", func(t *testing.T) {
		test := strings.Count("test", "test")
		expected := 1

		assertCorrectTest(t, test, expected)
	})

}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Count("test", "es")
	}
}
