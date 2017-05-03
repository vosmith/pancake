# Pancake
Pancake is a very simple package written in Go to flatten multidimensional
slices.  Currently it support `string`, `int`, and `float64` slices.

## Installation
```
go get github.com/vosmith08/pancake
```

## Usage
```
package main

import (
  "github.com/vosmith08/pancake"
)

func main() {
  // Create a 2-dimensional slice of strings
  arr := [][]string{
		{"a","b","c"},
		{"d","e","f"},
		{"g","h","i"},
		{"j","k","j"}}

  // Flatten it
	arr2, err := pancake.Strings(arr)
  if err != nil {
    panic(err)
  }

  // See the results
  fmt.Println(arr2)
  // >>> [a b c d e f g h i j k j]
}
```

## Future plans
- [ ] Support all primitive data types (bool, byte, chan, func)
- [ ] Support interface?
- [ ] Column-major Support
- [ ] Benchmarks
