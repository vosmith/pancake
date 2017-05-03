// pancake contains a few helper functions for flattening multidimensional
// slices of certain types.  It currently supports string, int, and float64.
package pancake

import (
  "reflect"
  "regexp"
  "errors"
)

// Strings takes a multidimensional slice of strings and flattens it into
// a 1-dimensional slice of strings in row major order.
func Strings(a interface{}) ([]string, error){
  // Yep, I get it.  I think I want Generics too...
  return flattenDepthString(reflect.ValueOf(a), getDepth(a))
}

// Ints takes a multidimensional slice of ints and flattens it into
// a 1-dimensional slice of ints in row major order.
func Ints(a interface{}) ([]int, error){
  return flattenDepthInt(reflect.ValueOf(a), getDepth(a))
}

// Float64s takes a multidimensional slice of float64 and flattens it into
// a 1-dimensional slice of float64 in row major order.
func Float64s(a interface{}) ([]float64, error){
  return flattenDepthFloat64(reflect.ValueOf(a), getDepth(a))
}

func flattenDepthString(a reflect.Value, depth int) ([]string, error) {
  // Ah, ok.  Pattern matching would be nice to have in Go.
  if depth < 1 {
    return []string{}, errors.New("input must be an slice of strings")
  } else if depth == 1 {
    stringArr := make([]string, a.Len())

    for i := 0; i < a.Len(); i++ {
      stringArr[i] = a.Index(i).String()
    }

    return stringArr, nil
  } else {
    stringArr := []string{}

    for i := 0; i < a.Len(); i++ {
      res, err := flattenDepthString(a.Index(i), depth -1)
      if err != nil{
        return []string{}, err
      }
      stringArr = append(stringArr, res...)
    }
    return stringArr, nil
  }
}

func flattenDepthInt(a reflect.Value, depth int) ([]int, error) {
  if depth < 1 {
    return []int{}, errors.New("input must be an slice of ints")
  } else if depth == 1 {
    intArr := make([]int, a.Len())

    for i := 0; i < a.Len(); i++ {
      intArr[i] = int(a.Index(i).Int())
    }

    return intArr, nil
  } else {
    intArr := []int{}

    for i := 0; i < a.Len(); i++ {
      res, err := flattenDepthInt(a.Index(i), depth -1)
      if err != nil{
        return []int{}, err
      }
      intArr = append(intArr, res...)
    }
    return intArr, nil
  }
}

func flattenDepthFloat64(a reflect.Value, depth int) ([]float64, error) {
  if depth < 1 {
    return []float64{}, errors.New("input must be an slice of float64s")
  } else if depth == 1 {
    floatArr := make([]float64, a.Len())

    for i := 0; i < a.Len(); i++ {
      floatArr[i] = a.Index(i).Float()
    }

    return floatArr, nil
  } else {
    floatArr := []float64{}

    for i := 0; i < a.Len(); i++ {
      res, err := flattenDepthFloat64(a.Index(i), depth -1)
      if err != nil{
        return []float64{}, err
      }
      floatArr = append(floatArr, res...)
    }
    return floatArr, nil
  }
}

func getDepth(a interface{}) int {
  typ := reflect.TypeOf(a)
  bracketExp,_ := regexp.Compile("\\[\\]")
  return len(bracketExp.FindAllStringIndex(typ.String(),-1))
}
