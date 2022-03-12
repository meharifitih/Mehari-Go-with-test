package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	got := sumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v  want %v", got, want)
	}
}

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	// t.Run("collection of any size", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3}
	// 	got := sum(numbers)
	// 	want := 6

	// 	if got != want {
	// 		t.Errorf("got %d wnat %d given %v", got, want, numbers)
	// 	}
	// })

}
