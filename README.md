# slice_fn

This go module provides simple methods for manipulations of slices via callback functions. Examples can be found under `slice_fn_test.go`.

## Examples

### Map
The Map() method creates a new slice/array populated with the results of
calling a provided function on every element in the calling array.
```golang
    in := []int{1, 2, 3, 4, 5}

	actual := slice_fn.Map(in, func(val int) int { return val + 1 })
    
    ->> [2, 3, 4, 5, 6]
```

### Filter
The Filter() method creates a copy of a portion of a given array,
filtered down to just the elements from the given array that pass the test implemented by the provided function.
```golang
    in := []int{1, 2, 2, 4, 2}

	actual := slice_fn.Filter(in, func(val int) bool { return val == 2 })    

    ->> [2, 2, 2]
```

### Reduce
The Reduce() method executes a user-supplied "reducer" callback function on each element of the array,
in order, passing in the return value from the calculation on the preceding element.
The final result of running the reducer across all elements of the array is a single value.
```golang
    in := []int{1, 2, 2, 4, 2}

	actual := slice_fn.Reduce(in, 0, func(previous int, current int) int { return previous + current })

    ->> 11
```

### Find
The Find() method returns the first element in the provided array that satisfies the provided testing function.
If no values satisfy the testing function, object/type with initialized values is returned. (Meaning f.e. int is used then 0 is returned)
```golang
    input := []int{1, 2, 3, 4, 5}

	actual := slice_fn.Find(input, func(val int) bool { return val == 4 })
    
    ->> 4
```

### FindIndex
The FindIndex() method returns the index of the first element in an array that satisfies the provided testing function.
If no elements satisfy the testing function, -1 is returned.
```golang
    input := []int{1, 2, 3, 4, 5}

	actual := slice_fn.FindIndex(input, func(val int) bool { return val == 3 })

    ->> 2
```

### FindLast
The FindLast() method returns the value of the last element in an array that satisfies the provided testing function.
If no values satisfy the testing function, object/type with initialized values is returned. (Meaning f.e. int is used then 0 is returned)
```golang
    input := []int{1, 5, 3, 4, 5}

	actual := slice_fn.FindLast(input, func(val int) bool { return val == 5 })

    ->> 5
```

### FindLastIndex
The FindLastIndex() method returns the value of the last element in an array that satisfies the provided testing function.
If no values satisfy the testing function, -1 is returned.
```golang
    input := []int{1, 5, 3, 4, 5}

	actual := slice_fn.FindLast(input, func(val int) bool { return val == 5 })

    ->> 4
```

### ForEach
The ForEach() method executes a provided function once for each array element.
```golang
    in := []int{1, 2, 2, 4, 2}
    sum := 0

	slice_fn.ForEach(input, func(val int) { sum += val }) 

    ->> 11
```
