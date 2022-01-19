package pancake

import (
	"testing"
	"fmt"
)

func TestInt1D(t *testing.T) {
	arr := []int{1,2,3}
	arr2, err := Ints(arr)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Failed, error")
	}
	if arr[0] != arr2[0] {
		t.Fatal("Failed, Not equal")
	}
	fmt.Println(arr2)
}

func TestInt2D(t *testing.T) {
	arr := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
		{0,11,12}}
	arr2, err := Ints(arr)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Failed, error")
	}
	if arr[2][0] != arr2[6] {
		t.Fatal("Failed, Not equal")
	}
	fmt.Println(arr2)
}

func TestInt3D(t *testing.T) {
	arr := [][][]int{
		{
			{1,2,3},
			{4,5,6}},
		{
			{7,8,9},
			{10,11,12}}}
	arr2, err := Ints(arr)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Failed, error")
	}
	if arr[0][0][1] != arr2[1] {
		t.Fatal("Failed, Not equal")
	}
	fmt.Println(arr2)
}
