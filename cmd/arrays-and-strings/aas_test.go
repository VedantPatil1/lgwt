package aas

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		got := Sum(nums)
		want := 15

		if got != want {
			t.Errorf("Got   %d \nWanted %d\n Given %v", got, want, nums)
		}
	})
	t.Run("Collection of 3 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3}

		got := Sum(nums)
		want := 6

		if got != want {
			t.Errorf("Got   %d \nWanted %d\n Given %v", got, want, nums)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	nums := []int{5, 5, 5, 5, 5}
	for b.Loop() {
		Sum(nums)
	}
}

func ExampleSum() {
	nums := []int{2, 2, 2, 2, 2}
	fmt.Println(Sum(nums))
	// Output: 10
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{0, 8})
	want := []int{6, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v Wanted %v", got, want)
	}
}

func TestSumAllTail(t *testing.T) {
	got := SumAllTail([]int{1, 2, 3}, []int{0, 8})
	want := []int{5, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v Wanted %v", got, want)
	}
}
