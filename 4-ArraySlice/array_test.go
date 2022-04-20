package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		expected := 15
		if got != expected {
			t.Errorf("expected: %v, got: %v", expected, got)
		}
	})

}

func TestSumAll(t *testing.T) {
	t.Run("should sum more than one slice", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected: %v, got %v", expected, got)
		}
	})

	t.Run("should sum one slice", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3})
		expected := []int{6}

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected: %v, got %v", expected, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("make the sums of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
