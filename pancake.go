package pancake

import (
  "reflect"
  "regexp"
  "errors"
)

func FlattenString(a interface{}) ([]string, error){
  return flattenDepthString(reflect.ValueOf(a), getDepth(a))
}

func FlattenInt64(a interface{}) ([]int64, error){
  return flattenDepthInt64(reflect.ValueOf(a), getDepth(a))
}

func FlattenFloat64(a interface{}) ([]float64, error){
  return flattenDepthFloat64(reflect.ValueOf(a), getDepth(a))
}

func flattenDepthString(a reflect.Value, depth int) ([]string, error) {
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

func flattenDepthInt64(a reflect.Value, depth int) ([]int64, error) {
  if depth < 1 {
    return []int64{}, errors.New("input must be an slice of ints")
  } else if depth == 1 {
    intArr := make([]int64, a.Len())

    for i := 0; i < a.Len(); i++ {
      intArr[i] = a.Index(i).Int()
    }

    return intArr, nil
  } else {
    intArr := []int64{}

    for i := 0; i < a.Len(); i++ {
      res, err := flattenDepthInt64(a.Index(i), depth -1)
      if err != nil{
        return []int64{}, err
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
