package vector

import (
	"encoding/json"
	"fmt"
)

func New[T any]() *Vector[T] {
	return From[T](nil)
}

func From[T any](ts *[]T) *Vector[T] {
	return &Vector[T]{vs: ts}
}

type Vector[T any] struct{ vs *[]T }

// Push adds an element to the end of the vector.
func (v *Vector[T]) Push(t T) {
	if v.vs == nil {
		v.vs = &[]T{}
	}
	*v.vs = append(*v.vs, t)
}

// Pop removes the last element from the vector and returns it.
// If the vector is empty, Pop returns false.
func (v *Vector[T]) Pop() (T, bool) {
	var t T
	if v.Len() == 0 {
		return t, false
	}
	t = (*v.vs)[v.Len()-1]
	*v.vs = (*v.vs)[:v.Len()-1]
	return t, true
}

// IsEmpty returns true if the vector is empty.
func (v *Vector[T]) IsEmpty() bool {
	return v.Len() == 0
}

// Len returns the number of elements in the vector.
func (v *Vector[T]) Len() int {
	return len(*v.vs)
}

// Clear removes all elements from the vector.
func (v *Vector[T]) Clear() {
	*v.vs = (*v.vs)[:0]
}

// At returns the element at the given index.
// If the index is out of range, At returns false.
// Equivalent to v[i], but returns false if i is out of range.
func (v *Vector[T]) At(i int) (T, bool) {
	var t T
	if i < 0 || i >= v.Len() {
		return t, false
	}
	t = (*v.vs)[i]
	return t, true
}

// Set sets the element at the given index to the given value.
// If the index is out of range, Set returns false.
// Equivalent to v[i] = t, but returns false if i is out of range.
func (v *Vector[T]) Set(i int, t T) bool {
	if i < 0 || i >= v.Len() {
		return false
	}
	(*v.vs)[i] = t
	return true
}

// Insert inserts the given element at the given index.
func (v *Vector[T]) Insert(i int, t T) bool {
	if i < 0 || i >= v.Len() {
		return false
	}
	*v.vs = append(*v.vs, t)
	copy((*v.vs)[i+1:], (*v.vs)[i:])
	(*v.vs)[i] = t
	return true
}

// Delete removes the element at the given index.
func (v *Vector[T]) Delete(i int) bool {
	if i < 0 || i >= v.Len() {
		return false
	}
	copy((*v.vs)[i:], (*v.vs)[i+1:])
	*v.vs = (*v.vs)[:v.Len()-1]
	return true
}

// Remove removes the first element equal to the given value.
func (v *Vector[T]) Remove(t T) bool {
	var a any = t
	var b any
	for i, item := range *v.vs {
		b = item
		if a == b {
			v.Delete(i)
			return true
		}
	}
	return false
}

// Swap swaps the elements at the given indices.
func (v *Vector[T]) Swap(i, j int) bool {
	if i < 0 || i >= v.Len() || j < 0 || j >= v.Len() {
		return false
	}
	(*v.vs)[i], (*v.vs)[j] = (*v.vs)[j], (*v.vs)[i]
	return true
}

// Slice returns a slice of the vector from i to j.
func (v *Vector[T]) Slice(i, j int) *Vector[T] {
	if i < 0 || i >= v.Len() || j < 0 || j >= v.Len() {
		return nil
	}
	item := (*v.vs)[i:j]
	return From(&item)
}

// Reverse reverses the elements of the vector.
func (v *Vector[T]) Reverse() {
	for i, j := 0, v.Len()-1; i < j; i, j = i+1, j-1 {
		(*v.vs)[i], (*v.vs)[j] = (*v.vs)[j], (*v.vs)[i]
	}
}

// Foreach calls the given function for each element of the vector.
func (v *Vector[T]) Foreach(f func(T)) {
	for _, item := range *v.vs {
		f(item)
	}
}

// ForeachIndex calls the given function for each element of the vector,
func (v *Vector[T]) ForeachIndex(f func(int, T)) {
	for i, item := range *v.vs {
		f(i, item)
	}
}

// Filter returns a new vector containing all elements for which the given function returns true.
func (v *Vector[T]) Filter(f func(T) bool) *Vector[T] {
	var item Vector[T]
	for _, t := range *v.vs {
		if f(t) {
			item.Push(t)
		}
	}
	return &item
}

// Map returns a new vector containing the results of applying the given function to each element of the vector.
func (v *Vector[T]) Map(f func(T) T) *Vector[T] {
	var item Vector[T]
	for _, t := range *v.vs {
		item.Push(f(t))
	}
	return &item
}

// Reduce applies a function to each element of the vector, returning the result.
func (v *Vector[T]) Reduce(f func(T, T) T) T {
	var t T
	for _, item := range *v.vs {
		t = f(t, item)
	}
	return t
}

// ReduceIndex applies a function to each element of the vector, returning the result.
func (v *Vector[T]) ReduceIndex(f func(int, T, T) T) T {
	var t T
	for i, item := range *v.vs {
		t = f(i, t, item)
	}
	return t
}

// Find returns the first element for which the given function returns true.
func (v *Vector[T]) Find(f func(T) bool) (T, bool) {
	var t T
	for _, item := range *v.vs {
		if f(item) {
			t = item
			return t, true
		}
	}
	return t, false
}

// FindIndex returns the index of the first element for which the given function returns true.
func (v *Vector[T]) FindIndex(f func(T) bool) (int, bool) {
	for i, item := range *v.vs {
		if f(item) {
			return i, true
		}
	}
	return 0, false
}

// FindLast returns the last element for which the given function returns true.
func (v *Vector[T]) FindLast(f func(T) bool) (T, bool) {
	var t T
	for i := v.Len() - 1; i >= 0; i-- {
		if f((*v.vs)[i]) {
			t = (*v.vs)[i]
			return t, true
		}
	}
	return t, false
}

// FindLastIndex returns the index of the last element for which the given function returns true.
func (v *Vector[T]) FindLastIndex(f func(T) bool) (int, bool) {
	for i := v.Len() - 1; i >= 0; i-- {
		if f((*v.vs)[i]) {
			return i, true
		}
	}
	return 0, false
}

// Prototype returns the prototype of the vector.
func (v *Vector[T]) Prototype() []T {
	return *v.vs
}

// Collect returns the elements of the vector.
func (v *Vector[T]) Collect() []T {
	return *v.vs
}

// UnmarshalJSON unmarshals the vector from JSON.
func (v *Vector[T]) UnmarshalJSON(bytes []byte) error {
	var item []T
	if err := json.Unmarshal(bytes, &item); err != nil {
		return err
	}
	v.vs = &item
	return nil
}

// MarshalJSON marshals the vector to JSON.
func (v *Vector[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.vs)
}

func (v *Vector[T]) String() string {
	return fmt.Sprintf("%v", *v.vs)
}
