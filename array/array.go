package array

// DeleteElement removes an element from a slice at a given index
// [T any] is a type constraint, it means that the function can be used with any type
// arr []T is a slice of any type
// index int is the location of the element to remove
// []T is the return type, a slice of any type
func DeleteElement[T any](arr []T, index int) []T {
	// return a new slice with the element at index removed
	// append the first part of the slice to the second part of the slice
	// arr[:index] is the first part of the slice up to the index, :index is exclusive
	// arr[index+1:] is the second part of the slice from the index to the end, index+1: is inclusive
	// ... is the spread operator, it expands the slice into individual elements
	return append(arr[:index], arr[index+1:]...)
}

// Copy creates a new slice with the same elements as the input slice
// [T any] is a type constraint, it means that the function can be used with any type
// arr []T is a slice of any type
// []T is the return type, a slice of any type
func Copy[T any](arr []T) []T {
	// return a new slice with the same elements as the input slice
	// []T(nil) is used to create a new slice, nil is an empty slice
	// arr is the slice to copy
	// ... is the spread operator, it expands the slice into individual elements
	return append([]T(nil), arr...)
}

// Append adds elements to the end of a slice
// [T any] is a type constraint, it means that the function can be used with any type
// arr []T is a slice of any type
// elements ...T is a variadic parameter, it allows the function to accept any number of elements
// []T is the return type, a slice of any type
func Append[T any](arr []T, elements ...T) []T {
	// append the elements to the slice
	// arr is the slice to append to
	// elements... is the spread operator, it expands the elements into individual elements
	return append(arr, elements...)
}

// Prepend adds elements to the start of a slice
// [T any] is a type constraint, it means that the function can be used with any type
// arr []T is a slice of any type
// elements ...T is a variadic parameter, it allows the function to accept any number of elements
// []T is the return type, a slice of any type
func Prepend[T any](arr []T, elements ...T) []T {
	// append the elements to the start of the slice
	// elements is the slice to append
	// arr... is the spread operator, it expands the slice into individual elements
	return append(elements, arr...)
}

// Insert adds elements to a slice at a given index
// [T any] is a type constraint, it means that the function can be used with any type
// arr []T is a slice of any type
// index int is the location to insert the elements
// elements ...T is a variadic parameter, it allows the function to accept any number of elements
// []T is the return type, a slice of any type
func Insert[T any](arr []T, index int, elements ...T) []T {
	// insert the elements at the index
	// append the first part of the slice to the elements, then append the second part of the slice
	// arr[:index] is the first part of the slice up to the index, :index is exclusive
	// arr[index:] is the second part of the slice from the index to the end, index: is inclusive
	// ... is the spread operator, it expands the slice into individual elements
	return append(arr[:index], append(elements, arr[index:]...)...)
}

// Remove removes elements from a slice
// [T comparable] is a type constraint, it means that the function can be used with any comparable type
// arr []T is a slice of any comparable type
// elements ...T is a variadic parameter, it allows the function to accept any number of elements
// []T is the return type, a slice of any comparable type
func Remove[T comparable](arr []T, elements ...T) []T {
	for _, element := range elements { // iterate over the elements to remove
		for i, v := range arr { // iterate over the slice keeping the index in i and the value in v
			if v == element { // if the value is equal to the element
				arr = append(arr[:i], arr[i+1:]...) // remove the element at the index by appending the first part to the second par, excluding the element
			}
		}
	}
	return arr
}

// Contains checks if a slice contains a given element
// [T comparable] is a type constraint, it means that the function can be used with any comparable type
// arr []T is a slice of any comparable type
// element T is the element to check for in the slice
// bool is the return type, a boolean value
func Contains[T comparable](arr []T, element T) bool {
	for _, v := range arr { // iterate over the slice keeping the value in v and discarding the index with _
		if v == element { // if the value is equal to the element
			return true // return true
		}
	}
	return false // return false if the element is not found
}

// GetElementIndex returns the index of an element in a slice
// [T comparable] is a type constraint, it means that the function can be used with any comparable type
// arr []T is a slice of any comparable type
// element T is the element to find in the slice
// int is the return type, an integer value
func GetElementIndex[T comparable](arr []T, element T) int {
	for i, v := range arr { // iterate over the slice keeping the index in i and the value in v
		if v == element { // if the value is equal to the element
			return i // return the index
		}
	}
	return -1 // return -1 if the element is not found
}

// Distinct returns a slice with duplicate elements removed
func Distinct[T comparable](arr []T) []T {
	distinct := make([]T, 0, len(arr))
	seen := make(map[T]struct{})
	for _, v := range arr {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			distinct = append(distinct, v)
		}
	}
	return distinct
}
