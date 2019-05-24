package arrays

import (
	"fmt"
)

// ConcatInt for int like javascript Array.prototype.concat
func ConcatInt(left []int, right []int) []int {
	result := make([]int, 0)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

// PopInt for int like javascript Array.prototype.pop
func PopInt(array []int) int {
	length := len(array)
	end := (array)[length-1]
	array = (array)[0 : length-1]
	return end
}

// PushInt for int like javascript Array.prototype.push
func PushInt(array []int, val int) {
	array = append(array, val)
}

// ShiftInt for int like javascript Array.prototype.shift
func ShiftInt(array []int) int {
	length := len(array)
	start := (array)[0]
	array = (array)[0 : length-1]
	return start
}

// UnshiftInt for int like javascript Array.prototype.unshift
func UnshiftInt(array []int, val int)  {
	newArray := []int{val}
	array = append(newArray, array...)
}

// IndexOfInt for int like javascript Array.prototype.indexOf
func IndexOfInt(array []int, val int) int  {
	for index, arrayVal := range array {
		if arrayVal == val {
			return index
		}
	}
	return -1
}

// LastIndexOfInt for int like javascript Array.prototype.lastIndexOf
func LastIndexOfInt(array []int, val int) int  {
	orderIndex := IndexOfInt(array, val)
	if orderIndex == -1 {
		return -1
	}
	return len(array) - 1 - orderIndex
}

// IncludesInt for int like javascript Array.prototype.includes
func IncludesInt(array []int, val int) bool  {
	exists := false
	for _, arrayVal := range array {
		if arrayVal == val {
			exists = true
			break
		}
	}
	return exists
}

// ReverseInt for int like javascript Array.prototype.reverse
func ReverseInt(array []int) {
	length := len(array)
	for index := range array {
		if index > length / 2 {
			(array)[index], (array)[length - 1 - index] =(array)[length - 1 - index], (array)[index]
		}
	}
}

// SpliceInt for int like javascript Array.prototype.splice
func SpliceInt(array []int, start int, length int, appends []int)  {
	newArray := make([]int, 0)
	left := (array)[0: start]
	right := (array)[(start+length): len(array)]
	newArray = append(newArray, left...)
	newArray = append(newArray, appends...)
	newArray = append(newArray, right...)
	array = newArray
}

// JoinInt for int like javascript Array.prototype.join
func JoinInt(array []int, slim string) string {
	str := fmt.Sprintf("%d", array[0])
	for index, val := range array {
		if index != 0 {
			str = fmt.Sprintf("%s%s%d", str, slim, val)
		}
	}
	return str
}

