package arrays

import "testing"

func TestSum(t *testing.T) {

	t.Run("When using fixed size ", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("When using undefined size ", func(t *testing.T) {

		nosizenumbers := []int{1, 2, 3}

		got := Sum(nosizenumbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, nosizenumbers)
		}
	})

}
