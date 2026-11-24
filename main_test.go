package main

import "testing"

func TestMiddle(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5, 6}, 4}, // second middle on even length
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
	}
	for _, c := range cases {
		if got := Middle(fromSlice(c.in)).Val; got != c.want {
			t.Errorf("Middle(%v) = %d, want %d", c.in, got, c.want)
		}
	}
	if Middle(nil) != nil {
		t.Error("Middle(nil) should be nil")
	}
}
