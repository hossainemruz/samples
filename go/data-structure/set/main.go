package main

import "fmt"

type Set struct {
	container map[string]struct{}
}

func NewSet() *Set {
	return &Set{
		container: make(map[string]struct{}),
	}
}

func (s *Set) Add(key string) {
	s.container[key] = struct{}{}
}

func (s *Set) Remove(key string) {
	delete(s.container, key)
}

func (s *Set) Contains(key string) bool {
	_, exist := s.container[key]
	return exist
}

func (s *Set) Size() int {
	return len(s.container)
}

// Test set
func main() {

	// Declare a new set
	s := NewSet()

	// Insert sample data
	s.Add("John")
	s.Add("James")
	s.Add("Dave")

	// Check set size
	fmt.Println("Set Size: ", s.Size())

	// Check sample data existence
	fmt.Println("Dave: ", s.Contains("Dave")) // true
	fmt.Println("Eve: ", s.Contains("Eve"))   // false

	// Remove some element
	s.Remove("Dave")

	// Check data has been removed
	fmt.Println("Dave: ", s.Contains("Dave")) // false

	// Check set size
	fmt.Println("Set Size: ", s.Size())
}
