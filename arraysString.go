package arrays

import (
	"fmt"
)

// ConcatString for string like javascript Array.prototype.concat
func ConcatString(left []string, right []string) []string {
	result := make([]string, 0)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

// PopString for string like javascript Array.prototype.pop
func PopString(array []string) string {
	length := len(array)
	end := (array)[length-1]
	array = (array)[0 : length-1]
	return end
}

// PushString for string like javascript Array.prototype.push
func PushString(array []string, val string) {
	array = append(array, val)
}

// ShiftString for string like javascript Array.prototype.shift
func ShiftString(array []string) string {
	length := len(array)
	start := (array)[0]
	array = (array)[0 : length-1]
	return start
}

// UnshiftString for string like javascript Array.prototype.unshift
func UnshiftString(array []string, val string)  {
	newArray := []string{val}
	array = append(newArray, array...)
}

// IndexOfString for string like javascript Array.prototype.indexOf
func IndexOfString(array []string, val string) int  {
	for index, arrayVal := range array {
		if arrayVal == val {
			return index
		}
	}
	return -1
}

// LastIndexOfString for string like javascript Array.prototype.lastIndexOf
func LastIndexOfString(array []string, val string) int  {
	orderIndex := IndexOfString(array, val)
	if orderIndex == -1 {
		return -1
	}
	return len(array) - 1 - orderIndex
}

// IncludesString for string like javascript Array.prototype.includes
func IncludesString(array []string, val string) bool  {
	exists := false
	for _, arrayVal := range array {
		if arrayVal == val {
			exists = true
			break
		}
	}
	return exists
}

// ReverseString for string like javascript Array.prototype.reverse
func ReverseString(array []string) {
	length := len(array)
	for index := range array {
		if index > length / 2 {
			(array)[index], (array)[length - 1 - index] =(array)[length - 1 - index], (array)[index]
		}
	}
}

// SpliceString for string like javascript Array.prototype.splice
func SpliceString(array []string, start int, length int, appends []string)  {
	newArray := make([]string, 0)
	left := (array)[0: start]
	right := (array)[(start+length): len(array)]
	newArray = append(newArray, left...)
	newArray = append(newArray, appends...)
	newArray = append(newArray, right...)
	array = newArray
}

// JoinString for string like javascript Array.prototype.join
func JoinString(array []string, slim string) string {
	str := array[0]
	for index, val := range array {
		if index != 0 {
			str = fmt.Sprintf("%s%s%s", str, slim, val)
		}
	}
	return str
}
