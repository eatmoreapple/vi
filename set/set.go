package set

import (
	"fmt"
)

// Set defines an collection of unique elements.
type Set[T comparable] map[T]struct{}

// New creates a new set.
func New[T comparable]() Set[T] {
	return make(Set[T])
}

// From creates a new set from a slice.
func From[T comparable](ts []T) Set[T] {
	s := New[T]()
	for _, t := range ts {
		s.Add(t)
	}
	return s
}

// Add adds an element to the set.
func (s Set[T]) Add(t T) {
	s[t] = struct{}{}
}

// Remove removes an element from the set.
func (s Set[T]) Remove(t T) {
	delete(s, t)
}

// Contains returns true if the set contains the given element.
func (s Set[T]) Contains(t T) bool {
	_, ok := s[t]
	return ok
}

// Len returns the number of elements in the set.
func (s Set[T]) Len() int {
	return len(s)
}

// Clear removes all elements from the set.
func (s Set[T]) Clear() {
	for t := range s {
		delete(s, t)
	}
}

// ToSlice returns the set as a slice.
func (s Set[T]) ToSlice() []T {
	ts := make([]T, 0, s.Len())
	for t := range s {
		ts = append(ts, t)
	}
	return ts
}

// Union returns the union of two sets.
func (s Set[T]) Union(s2 Set[T]) Set[T] {
	s3 := New[T]()
	for t := range s {
		s3.Add(t)
	}
	for t := range s2 {
		s3.Add(t)
	}
	return s3
}

// Intersection returns the intersection of two sets.
func (s Set[T]) Intersection(s2 Set[T]) Set[T] {
	s3 := New[T]()
	for t := range s {
		if s2.Contains(t) {
			s3.Add(t)
		}
	}
	return s3
}

// Difference returns the difference of two sets.
func (s Set[T]) Difference(s2 Set[T]) Set[T] {
	s3 := New[T]()
	for t := range s {
		if !s2.Contains(t) {
			s3.Add(t)
		}
	}
	return s3
}

// SymmetricDifference returns the symmetric difference of two sets.
func (s Set[T]) SymmetricDifference(s2 Set[T]) Set[T] {
	s3 := New[T]()
	for t := range s {
		if !s2.Contains(t) {
			s3.Add(t)
		}
	}
	for t := range s2 {
		if !s.Contains(t) {
			s3.Add(t)
		}
	}
	return s3
}

// IsSubset returns true if s is a subset of s2.
func (s Set[T]) IsSubset(s2 Set[T]) bool {
	for t := range s {
		if !s2.Contains(t) {
			return false
		}
	}
	return true
}

// IsSuperset returns true if s is a superset of s2.
func (s Set[T]) IsSuperset(s2 Set[T]) bool {
	return s2.IsSubset(s)
}

// IsDisjoint returns true if s and s2 are disjoint.
func (s Set[T]) IsDisjoint(s2 Set[T]) bool {
	for t := range s {
		if s2.Contains(t) {
			return false
		}
	}
	return true
}

// Equal returns true if s and s2 are equal.
func (s Set[T]) Equal(s2 Set[T]) bool {
	return s.IsSubset(s2) && s2.IsSubset(s)
}

// Clone returns a copy of the set.
func (s Set[T]) Clone() Set[T] {
	s2 := New[T]()
	for t := range s {
		s2.Add(t)
	}
	return s2
}

// String returns a string representation of the set.
func (s Set[T]) String() string {
	return fmt.Sprintf("%v", s.ToSlice())
}
