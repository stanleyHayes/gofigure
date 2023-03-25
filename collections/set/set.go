package set

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type Set[T constraints.Integer | constraints.Float | string] struct {
	size     uint64
	elements []T
}

// New Create a new set of type T
func New[T constraints.Integer | constraints.Float | string](elements ...T) *Set[T] {
	set := &Set[T]{
		size:     uint64(0),
		elements: make([]T, 0),
	}
	for _, element := range elements {
		set.Add(element)
	}
	return set
}

// Size Return the number of elements in the set
func (s *Set[T]) Size() uint64 {
	return s.size
}

// Length Return the number of elements in the set
// Usage
// size := set.Size()
// fmt.Println(size)
// Returns uint64
func (s *Set[T]) Length() uint64 {
	return s.size
}

// Cardinality Return the number of elements in the set
// Usage
// size := set.Cardinality()
// fmt.Println(size)
// Returns uint64
func (s *Set[T]) Cardinality() uint64 {
	return s.size
}

// In Tests if an item is a member of the set
func (s *Set[T]) In(element T) bool {
	for _, item := range s.elements {
		if item == element {
			return true
		}
	}
	return false
}

// NotIn Tests element for non-membership in set
func (s *Set[T]) NotIn(element T) bool {
	for _, item := range s.elements {
		if item == element {
			return false
		}
	}
	return true
}

// IsDisjoint Return true if the set has no elements in common with other
func (s *Set[T]) IsDisjoint(other *Set[T]) bool {
	for _, element := range s.elements {
		if other.In(element) {
			return false
		}
	}
	return true
}

// IsSubset Test whether every element in the set is in other
func (s *Set[T]) IsSubset(other *Set[T]) bool {
	for _, element := range other.elements {
		if other.NotIn(element) {
			return false
		}
	}
	return true
}

// IsSuperset Test whether every element in other is in the set
func (s *Set[T]) IsSuperset(other *Set[T]) bool {
	for _, element := range other.elements {
		if s.NotIn(element) {
			return false
		}
	}
	return true
}

// Union Return a new set with elements from set and all others
func (s *Set[T]) Union(others ...Set[T]) *Set[T] {
	size := uint64(0)
	set := New[T]()
	for _, other := range others {
		size += other.size
		set.elements = append(set.elements, other.elements...)
	}
	set.size = size
	return set
}

// Intersection Return a new set with elements common to the set and all others
func (s *Set[T]) Intersection(others ...Set[T]) *Set[T] {
	set := New[T]()
	for _, element := range s.elements {
		found := true
		for _, other := range others {
			if other.NotIn(element) {
				found = false
				break
			}
		}
		if found {
			set.Add(element)
		}
	}
	return set
}

// Difference Return a new set with elements that are not in others
func (s *Set[T]) Difference(others ...Set[T]) *Set[T] {
	set := New[T]()
	for _, element := range s.elements {
		found := true
		for _, other := range others {
			if other.In(element) {
				found = false
				break
			}
		}
		if found {
			set.Add(element)
		}
	}
	return set
}

// Copy Return a shallow copy of set
func (s *Set[T]) Copy() *Set[T] {
	set := New[T]()
	for _, element := range s.elements {
		set.Add(element)
	}
	return set
}

// SymmetricDifference Returns a new set with elements in either the set or others but not both
func (s *Set[T]) SymmetricDifference(other Set[T]) *Set[T] {
	set := New[T]()
	for _, element := range s.elements {
		if other.NotIn(element) {
			set.Add(element)
		}
	}
	for _, element := range other.elements {
		if s.NotIn(element) {
			set.Add(element)
		}
	}
	return set
}

// Add Adds elements to the set if elements not already in set
func (s *Set[T]) Add(elements ...T) {
	for _, element := range elements {
		if s.NotIn(element) {
			s.size += 1
			s.elements = append(s.elements, element)
		}
	}
}

// Remove Removes an elements present in set
func (s *Set[T]) Remove(elements ...T) {
	for index, element := range elements {
		if s.In(element) {
			s.size -= 1
			slices.Delete(s.elements, index, index+1)
		}
	}
}

// String Prints the elements of a set

// Clear Removes all the elements in the set
func (s *Set[T]) Clear() {
	slices.Delete(s.elements, 0, int(s.Size()))
	s.size = uint64(0)
}

// From Create a set from another iterable - list or dictionary

// Index Returns the index of the element
func (s *Set[T]) Index(item T) (int64, bool) {
	for i, element := range s.elements {
		if element == item {
			return int64(i), true
		}
	}
	return int64(-1), false
}

// IsEmpty Returns true if the size of the set is 0 otherwise false
func (s *Set[T]) IsEmpty() bool {
	return s.size == 0
}
