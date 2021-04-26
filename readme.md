#### Golang Arrays
##### an array utils like JavaScript Array, includes contains, concat, pop, push, shift, unshift, indexOf, lastIndexOf, includes,reverse, splice, join 

```
useage: go get -u github.com/anymost/arrays
```
```
import (
    "github.com/anymost/arrays"
)

func main() {
    array = []int{1, 2, 3, 4}
    arrays.push(&array, 5)
    fmt.Println(array) // result is [1, 2, 3, 4, 5]
}
```