package main

import "testing"

// TestAdd calls Add with a series of ints, checking for a valid return
func TestAdd(t *testing.T) {
	cases := []struct {
		one int
		two int
	}{
		{
			one: 10,
			two: 20,
		},
		{
			one: 1,
			two: 2,
		},
		{
			one: 1000000,
			two: 2000000,
		},
	}

	for i, c := range cases {
		actual := Add(c.one, c.two)
		expected := c.one + c.two

		if actual != expected {
			t.Fatalf("Test case %d: expected %d, got %d\n", i, expected, actual)
		}
	}
}
