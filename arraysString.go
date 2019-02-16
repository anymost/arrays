package arrays

func ConcatString(left []string, right []string) []string {
	result := make([]string, 0)
	copy(result, left)
	copy(result, right)
	return result
}

func PopString(array *[]string) string {
	length := len(*array)
	end := (*array)[length-1]
	*array = (*array)[0 : length-1]
	return end
}

func PushString(array *[]string, val string) {
	*array = append(*array, val)
}

func ShiftString(array *[]string) string {
	length := len(*array)
	start := (*array)[0]
	*array = (*array)[0 : length-1]
	return start
}

func UnshiftString(array *[]string, val string)  {
	newArray := []string{val}
	*array = append(newArray, *array...)
}

func IndexOfString(array *[]string, val string) int  {
	for index, arrayVal := range *array {
		if arrayVal == val {
			return index
		}
	}
	return -1
}

func LastIndexOfString(array *[]string, val string) int  {
	orderIndex := IndexOfString(array, val)
	if orderIndex == -1 {
		return -1
	}
	return len(*array) - 1 - orderIndex
}

func IncludesString(array *[]string, val string) bool  {
	exists := false
	for _, arrayVal := range *array {
		if arrayVal == val {
			exists = true
			break
		}
	}
	return exists
}
