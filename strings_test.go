package pancake_test

import (
	"testing"
	. "github.com/vosmith/pancake"
	"fmt"
)

func TestFlattenString1D(t *testing.T) {
	arr := []string{"a","b","c"}
	arr2, err := FlattenString(arr)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Failed, error")
	}
	if arr[0] != arr2[0] {
		t.Fatal("Failed, Not equal")
	}
	fmt.Println(arr2)
}

func TestFlattenString2D(t *testing.T) {
	arr := [][]string{
		{"a","b","c"},
		{"d","e","f"},
		{"g","h","i"},
		{"j","k","j"}}
	arr2, err := FlattenString(arr)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Failed, error")
	}
	if arr[2][0] != arr2[6] {
		t.Fatal("Failed, Not equal")
	}
	fmt.Println(arr2)
}

func TestFlattenString3D(t *testing.T) {
	arr := [][][]string{
		[][]string{
			[]string{"a","b","c"},
			[]string{"d","e","f"}},
		[][]string{
			[]string{"g","h","i"},
			[]string{"j","k","l"}}}
	arr2, err := FlattenString(arr)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Failed, error")
	}
	if arr[0][0][1] != arr2[1] {
		t.Fatal("Failed, Not equal")
	}
	fmt.Println(arr2)
}
