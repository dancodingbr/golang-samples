package bubblesort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithDifferentSequenceOfIntegers(t *testing.T) {
	got := []int{389, 1261, 4, 29, 34, 3322, 10324, 546, 365724, 17}
	BubbleSort(got)
	wanted := []int{4, 17, 29, 34, 389, 546, 1261, 3322, 10324, 365724}

	if reflect.DeepEqual(got, wanted) == false {
		t.Errorf("got: %v; wanted: %v\n", got, wanted)
	} else {
		fmt.Printf("got: %v; wanted: %v\n", got, wanted)
	}
}

func TestWithAtLeastOneNegativeNumber(t *testing.T) {
	got := []int{389, 1261, 4, -21, 34, 3322, 10324, 546, 365724, 17}
	BubbleSort(got)
	wanted := []int{-21, 4, 17, 34, 389, 546, 1261, 3322, 10324, 365724}

	if reflect.DeepEqual(got, wanted) == false {
		t.Errorf("got: %v; wanted: %v\n", got, wanted)
	} else {
		fmt.Printf("got: %v; wanted: %v\n", got, wanted)
	}
}

func TestSwapWithTwoElements(t *testing.T) {
	got := []int{3322, 546}
	Swap(got, 0)
	wanted := []int{546, 3322}

	if reflect.DeepEqual(got, wanted) == false {
		t.Errorf("got: %v; wanted: %v\n", got, wanted)
	} else {
		fmt.Printf("got: %v; wanted: %v\n", got, wanted)
	}
}

func TestSwapMoreThanTwoElements(t *testing.T) {
	got := []int{3322, 546, 10324, 365724, 17}
	Swap(got, 3)
	wanted := []int{3322, 546, 10324, 17, 365724}

	if reflect.DeepEqual(got, wanted) == false {
		t.Errorf("got: %v; wanted: %v\n", got, wanted)
	} else {
		fmt.Printf("got: %v; wanted: %v\n", got, wanted)
	}
}

func TestSwapWhenIndexMinorThanZero(t *testing.T) {
	got := []int{3322, 546}
	Swap(got, -1)
	wanted := []int{3322, 546}

	if reflect.DeepEqual(got, wanted) == false {
		t.Errorf("got: %v; wanted: %v\n", got, wanted)
	} else {
		fmt.Printf("got: %v; wanted: %v\n", got, wanted)
	}
}

func TestSwapWhenIndexMajorOrEqualsThanSliceLength(t *testing.T) {
	got := []int{3322, 546}
	Swap(got, len(got))
	wanted := []int{3322, 546}

	if reflect.DeepEqual(got, wanted) == false {
		t.Errorf("got: %v; wanted: %v\n", got, wanted)
	} else {
		fmt.Printf("got: %v; wanted: %v\n", got, wanted)
	}
}
