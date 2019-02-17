package arrays

import (
	"fmt"
)

// ConcatFloat64 for float64 like javascript Array.prototype.concat
func ConcatFloat64(left []float64, right []float64) []float64 {
	result := make([]float64, 0)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

// PopFloat64 for float64 like javascript Array.prototype.pop
func PopFloat64(array *[]float64) float64 {
	length := len(*array)
	end := (*array)[length-1]
	*array = (*array)[0 : length-1]
	return end
}

// PushFloat64 for float64 like javascript Array.prototype.push
func PushFloat64(array *[]float64, val float64) {
	*array = append(*array, val)
}

// ShiftFloat64 for float64 like javascript Array.prototype.shift
func ShiftFloat64(array *[]float64) float64 {
	length := len(*array)
	start := (*array)[0]
	*array = (*array)[0 : length-1]
	return start
}

// UnshiftFloat64 for float64 like javascript Array.prototype.unshift
func UnshiftFloat64(array *[]float64, val float64)  {
	newArray := []float64{val}
	*array = append(newArray, *array...)
}

// IndexOfFloat64 for float64 like javascript Array.prototype.indexOf
func IndexOfFloat64(array *[]float64, val float64) int  {
	for index, arrayVal := range *array {
		if arrayVal == val {
			return index
		}
	}
	return -1
}

// LastIndexOfFloat64 for float64 like javascript Array.prototype.lastIndexOf
func LastIndexOfFloat64(array *[]float64, val float64) int  {
	orderIndex := IndexOfFloat64(array, val)
	if orderIndex == -1 {
		return -1
	}
	return len(*array) - 1 - orderIndex
}

// IncludesFloat64 for float64 like javascript Array.prototype.includes
func IncludesFloat64(array *[]float64, val float64) bool  {
	exists := false
	for _, arrayVal := range *array {
		if arrayVal == val {
			exists = true
			break
		}
	}
	return exists
}

// ReverseFloat64 for float64 like javascript Array.prototype.reverse
func ReverseFloat64(array *[]float64) {
	length := len(*array)
	for index := range *array {
		if index > length / 2 {
			(*array)[index], (*array)[length - 1 - index] =(*array)[length - 1 - index], (*array)[index]
		}
	}
}

// SpliceFloat64 for float64 like javascript Array.prototype.splice
func SpliceFloat64(array *[]float64, start int, length int, appends []float64)  {
	newArray := make([]float64, 0)
	left := (*array)[0: start]
	right := (*array)[(start+length): len(*array)]
	newArray = append(newArray, left...)
	newArray = append(newArray, appends...)
	newArray = append(newArray, right...)
	*array = newArray
}

// JoinFloat64 for float64 like javascript Array.prototype.join
func JoinFloat64(array []float64, slim string) string {
	str := fmt.Sprintf("%f", array[0])
	for index, val := range array {
		if index != 0 {
			str = fmt.Sprintf("%s%s%f", str, slim, val)
		}
	}
	return str
}