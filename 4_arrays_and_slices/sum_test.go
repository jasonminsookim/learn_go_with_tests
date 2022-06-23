package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{5, 10, 15}

	got := SumAll(slice1, slice2)
	want := []int{10, 30}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Sum an array of size 10.
		slice_1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		Sum(slice_1)
	}
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice1 := []int{12, 20, 100, 1, 5, 10}
		slice2 := []int{15, 3, 1}
		SumAll(slice1, slice2)
	}

}
