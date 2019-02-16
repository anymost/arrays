package arrays

func Concat(left []interface{}, right []interface{}) []interface{} {
	result := make([]interface{}, 0)
	copy(result, left)
	copy(result, right)
	return result
}

func Pop(array *[]interface{}) interface{} {
	length := len(*array)
	end := (*array)[length-1]
	*array = (*array)[0 : length-1]
	return end
}

func Push(array *[]interface{}, val interface{}) {
	*array = append(*array, val)
}

func Shift(array *[]interface{}) interface{} {
	length := len(*array)
	start := (*array)[0]
	*array = (*array)[0 : length-1]
	return start
}

func Unshift(array *[]interface{}, val interface{})  {
	newArray := []interface{}{val}
	*array = append(newArray, *array)
}

func IndexOf(array *[]interface{}, val interface{}) int  {
	for index, arrayVal := range *array {
		if arrayVal == val {
			return index
		}
	}
	return -1
}

func LastIndexOf(array *[]interface{}, val interface{}) int  {
	orderIndex := IndexOf(array, val)
	if orderIndex == -1 {
		return -1
	}
	return len(*array) - 1 -IndexOf(array, val)
}

func Includes(array *[]interface{}, val interface{}) bool  {
	exists := false
	for _, arrayVal := range *array {
		if arrayVal == val {
			exists = true
			break
		}
	}
	return exists
}
