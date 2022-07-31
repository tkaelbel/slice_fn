package slice_fn

/*
TODO:
FlatMap
Some
Sort
*/

/*
The Map() method creates a new slice/array populated with the results of
calling a provided function on every element in the calling array.
*/
func Map[T, U any](s []T, f func(T) U) []U {
	r := make([]U, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

/*
The Filter() method creates a copy of a portion of a given array,
filtered down to just the elements from the given array that pass the test implemented by the provided function.
*/
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

/*
The Reduce() method executes a user-supplied "reducer" callback function on each element of the array,
in order, passing in the return value from the calculation on the preceding element.
The final result of running the reducer across all elements of the array is a single value.
*/
func Reduce[T, U any](s []T, init U, f func(U, T) U) U {
	r := init
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

/*
TODO: Maybe return err instead of empty object?
The Find() method returns the first element in the provided array that satisfies the provided testing function.
If no values satisfy the testing function, object/type with initialized values is returned. (Meaning f.e. int is used then 0 is returned)
*/
func Find[T any](s []T, f func(T) bool) T {
	var result T
	for _, v := range s {
		if f(v) {
			return v
		}
	}
	return result
}

/*
The FindIndex() method returns the index of the first element in an array that satisfies the provided testing function.
If no elements satisfy the testing function, -1 is returned.
*/
func FindIndex[T any](s []T, f func(T) bool) int {
	for idx, v := range s {
		if f(v) {
			return idx
		}
	}
	return -1
}

/*
The FindLast() method returns the value of the last element in an array that satisfies the provided testing function.
If no values satisfy the testing function, object/type with initialized values is returned. (Meaning f.e. int is used then 0 is returned)
*/
func FindLast[T any](s []T, f func(T) bool) T {
	var result T
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return s[i]
		}
	}

	return result
}

/*
The FindLastIndex() method returns the value of the last element in an array that satisfies the provided testing function.
If no values satisfy the testing function, -1 is returned.
*/
func FindLastIndex[T any](s []T, f func(T) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return i
		}
	}
	return -1
}

/* The ForEach() method executes a provided function once for each array element. */
func ForEach[T any](s []T, f func(T)) {
	for _, v := range s {
		f(v)
	}
}
