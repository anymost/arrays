package arrays

func ConcatFloat64(left []float64, right []float64) []float64 {
	result := make([]float64, 0)
	copy(result, left)
	copy(result, right)
	return result
}

func PopFloat64(array *[]float64) float64 {
	length := len(*array)
	end := (*array)[length-1]
	*array = (*array)[0 : length-1]
	return end
}

func PushFloat64(array *[]float64, val float64) {
	*array = append(*array, val)
}

func ShiftFloat64(array *[]float64) float64 {
	length := len(*array)
	start := (*array)[0]
	*array = (*array)[0 : length-1]
	return start
}

func UnshiftFloat64(array *[]float64, val float64)  {
	newArray := []float64{val}
	*array = append(newArray, *array...)
}

func IndexOfFloat64(array *[]float64, val float64) int  {
	for index, arrayVal := range *array {
		if arrayVal == val {
			return index
		}
	}
	return -1
}

func LastIndexOfFloat64(array *[]float64, val float64) int  {
	orderIndex := IndexOfFloat64(array, val)
	if orderIndex == -1 {
		return -1
	}
	return len(*array) - 1 - orderIndex
}

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
