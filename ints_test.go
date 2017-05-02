package pancake_test

import (
	"testing"
	. "github.com/vosmith/pancake"
	"fmt"
)

func TestFlattenInt641D(t *testing.T) {
	arr := []int{1,2,3}
	arr2, err := FlattenInt64(arr)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Failed, error")
	}
	if arr[0] != arr2[0] {
		t.Fatal("Failed, Not equal")
	}
	fmt.Println(arr2)
}

func TestFlattenInt642D(t *testing.T) {
	arr := [][]int64{
		{1,2,3},
		{4,5,6},
		{7,8,9},
		{0,11,12}}
	arr2, err := FlattenInt64(arr)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Failed, error")
	}
	if arr[2][0] != arr2[6] {
		t.Fatal("Failed, Not equal")
	}
	fmt.Println(arr2)
}

func TestFlattenInt643D(t *testing.T) {
	arr := [][][]int64{
		[][]int64{
			[]int64{1,2,3},
			[]int64{4,5,6}},
		[][]int64{
			[]int64{7,8,9},
			[]int64{10,11,12}}}
	arr2, err := FlattenInt64(arr)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Failed, error")
	}
	if arr[0][0][1] != arr2[1] {
		t.Fatal("Failed, Not equal")
	}
	fmt.Println(arr2)
}
