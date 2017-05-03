package pancake_test

import (
	"testing"
	. "github.com/vosmith/pancake"
	"fmt"
)

func TestFloat641D(t *testing.T) {
	arr := []float64{1.0,2.0,3.0}
	arr2, err := Float64s(arr)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Failed, error")
	}
	if arr[0] != arr2[0] {
		t.Fatal("Failed, Not equal")
	}
	fmt.Println(arr2)
}

func TestFloat642D(t *testing.T) {
	arr := [][]float64{
		{1.0,2.0,3.0},
		{4.0,5.0,6.0},
		{7.0,8.0,9.0},
		{0.0,11.0,12.0}}
	arr2, err := Float64s(arr)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Failed, error")
	}
	if arr[2][0] != arr2[6] {
		t.Fatal("Failed, Not equal")
	}
	fmt.Println(arr2)
}

func TestFloat643D(t *testing.T) {
	arr := [][][]float64{
		[][]float64{
			[]float64{1.0,2.0,3.0},
			[]float64{4.0,5.0,6.0}},
		[][]float64{
			[]float64{7.0,8.0,9.0},
			[]float64{10.0,11.0,12.0}}}
	arr2, err := Float64s(arr)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Failed, error")
	}
	if arr[0][0][1] != arr2[1] {
		t.Fatal("Failed, Not equal")
	}
	fmt.Println(arr2)
}
