package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Should return 4 with Add(2,2)", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected %d but got %d", expected, sum)
		}
	})

	t.Run("Should return 6 with Add(4,2)", func(t *testing.T) {
		sum := Add(4, 2)
		expected := 6

		if sum != expected {
			t.Errorf("expected %d but got %d", expected, sum)
		}
	})
}

type UseCase struct {
	X int
	Y int
}

// Teste usando exemplo
func ExampleAdd() {
	use_cases := []UseCase{
		{X: 1, Y: 4},
		{X: 3, Y: 5},
		{X: 8, Y: 7},
	}

	for _, usecase := range use_cases {
		fmt.Println(Add(usecase.X, usecase.Y))
	}
	// Output:
	// 5
	// 8
	// 15
}
