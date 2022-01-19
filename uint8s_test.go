package pancake

import (
	"testing"
	"fmt"
)

func TestUint81D(t *testing.T) {
	arr := []uint8{1,2,3}
	arr2, err := Uint8s(arr)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Failed, error")
	}
	if arr[0] != arr2[0] {
		t.Fatal("Failed, Not equal")
	}
	fmt.Println(arr2)
}

func TestUint82D(t *testing.T) {
	arr := [][]uint8{
		{1,2,3},
		{4,5,6},
		{7,8,9},
		{0,11,12}}
	arr2, err := Uint8s(arr)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Failed, error")
	}
	if arr[2][0] != arr2[6] {
		t.Fatal("Failed, Not equal")
	}
	fmt.Println(arr2)
}

func TestUint83D(t *testing.T) {
	arr := [][][]uint8{
		{
			{1,2,3},
			{4,5,6}},
		{
			{7,8,9},
			{10,11,12}}}
	arr2, err := Uint8s(arr)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Failed, error")
	}
	if arr[0][0][1] != arr2[1] {
		t.Fatal("Failed, Not equal")
	}
	fmt.Println(arr2)
}
