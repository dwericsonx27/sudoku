package set

import (
	"errors"
)

// IntSet is the data structure to hold a set of integers.
type IntSet struct {
	members map[int]bool
}

// NewIntSet creates a empty set of integers.
func NewIntSet() *IntSet {
	return &IntSet{make(map[int]bool)}
}

// Add inserts another member to the set.
func (set *IntSet) Add(i int) bool {
	_, found := set.members[i]
	set.members[i] = true
	return !found //False if it existed already
}

// Contains looks to see if an integer is a member or not.
func (set *IntSet) Contains(i int) bool {
	_, found := set.members[i]
	return found //true if it existed already
}

// Remove eliminates an specified integer.
func (set *IntSet) Remove(i int) {
	delete(set.members, i)
}

// Size returns the number of unique members in the set.
func (set *IntSet) Size() int {
	return len(set.members)
}

// GetLastValue returns the value of the last remaining element of the set.
func (set *IntSet) GetLastValue() (int, error) {
	if len(set.members) == 1 {
		for key, value := range set.members {
			if value {
				return key, nil
			}
		}
	} else if len(set.members) > 1 {
		return 0, errors.New("more than 1 value in the set")
	} else {
		return 0, errors.New("no values in the set")
	}
	return 0, errors.New("set inconsistent state")
}

// GetAllMembers puts all unique elements into an array.
func (set *IntSet) GetAllMembers() []int {
	size := set.Size()
	if size == 0 {
		return nil
	}
	index := 0
	rtnval := make([]int, size)
	for k := range set.members {
		rtnval[index] = k
		index++
	}
	return rtnval
}

// Equals tests for equality between 2 sets.
func (set *IntSet) Equals(subject *IntSet) bool {
	if set.Size() == subject.Size() {
		for key, value := range set.members {
			if value {
				if !subject.Contains(key) {
					return false
				}
			}
		}
		return true
	}

	return false
}

// Clear removes all members of the set.
func (set *IntSet) Clear() {
	for k := range set.members {
		delete(set.members, k)
	}
}
