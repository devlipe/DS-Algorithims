package structures

// A Set is a data structure that stores a collection of items just once.
type Set[T comparable] map[T]struct{}

// NewSet returns a new Set.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{}
}

// Create a new set with elements that was passed as arguments.
func NewSetWithElements[T comparable](elements ...T) *Set[T] {
	set := NewSet[T]()
	for _, element := range elements {
		set.Add(element)
	}
	return set
}

// Add an element to the set. Returns true if the element was adde.
// False if the element was already in the set.
func (s *Set[T]) Add(element T) bool {
	_, exists := (*s)[element]
	(*s)[element] = struct{}{}
	return !exists
}

// Return the number of elements in the set
func (s *Set[T]) Len() int {
	return len(*s)
}

// Tests if two sets are equal.
func (s *Set[T]) Equals(other *Set[T]) bool {
	if other == nil || s.Len() != other.Len() {
		return false
	}

	for value := range *s {
		if !other.Contains(value) {
			return false
		}
	}
	return true
}

// Returns whether the set contains the value or not.
func (s Set[T]) Contains(value T) bool {
	_, has := s[value]
	return has
}

// Merge will add the elements of other to the receiver set
func (s *Set[T]) Merge(other *Set[T]) {
	if other == nil {
		return
	}

	for value := range *other {
		s.Add(value)
	}
}

// Remove deletes a value from the set and return the status of the operation
func (s *Set[T]) Remove(value T) bool {
	if _, exists := (*s)[value]; exists {
		delete(*s, value)
		return true
	}
	return false
}

// Execute the function for each element in the set
func (s *Set[T]) Each(f func(T, func())) {
	stopped := false

	stop := func() { stopped = true }

	for v := range *s {

		f(v, stop)
		if stopped {
			return
		}
	}
}
