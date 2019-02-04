package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got '%d' want '%d', give '%v'", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	t.Run("two slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9}, []int{1, 1})
		want := []int{3, 9, 2}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2, 4}, []int{0, 9, -1})
	want := []int{6, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}
