package arrays

func ConcatInt(left []int, right []int) []int {
	result := make([]int, 0)
	copy(result, left)
	copy(result, right)
	return result
}

func PopInt(array *[]int) int {
	length := len(*array)
	end := (*array)[length-1]
	*array = (*array)[0 : length-1]
	return end
}

func PushInt(array *[]int, val int) {
	*array = append(*array, val)
}

func ShiftInt(array *[]int) int {
	length := len(*array)
	start := (*array)[0]
	*array = (*array)[0 : length-1]
	return start
}

func UnshiftInt(array *[]int, val int)  {
	newArray := []int{val}
	*array = append(newArray, *array...)
}

func IndexOfInt(array *[]int, val int) int  {
	for index, arrayVal := range *array {
		if arrayVal == val {
			return index
		}
	}
	return -1
}

func LastIndexOfInt(array *[]int, val int) int  {
	orderIndex := IndexOfInt(array, val)
	if orderIndex == -1 {
		return -1
	}
	return len(*array) - 1 - orderIndex
}

func IncludesInt(array *[]int, val int) bool  {
	exists := false
	for _, arrayVal := range *array {
		if arrayVal == val {
			exists = true
			break
		}
	}
	return exists
}

