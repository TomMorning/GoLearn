package array

import "testing"

func TestSum(t *testing.T) {
	t.Run("Sum of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("want %d got %d given %v", want, got, numbers)
		}
	})

	t.Run("Sum of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("want %d got %d given %v", want, got, numbers)
		}
	})
}
