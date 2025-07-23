package interations

import (
	"testing"
	"fmt"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a",5)
	want := "aaaaa"
	if got != want {
		t.Errorf("Got %q Expected %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a",7)
	}
}

func ExampleRepeat() {
	res := Repeat("a", 2)
	fmt.Println(res)
	// Output: aa
}
