package sudoku

import (
	"set"
	"testing"
)

func TestSetCreation(t *testing.T) {
	var s = set.NewIntSet()
	if s == nil {
		t.Errorf("Failed to create integer set!")
	}
}

func TestAddContains(t *testing.T) {
	var s = set.NewIntSet()
	if s != nil {
		if s.Size() != 0 {
			t.Errorf("Set size is: %d, prior to adding any members.", s.Size())
		}

		var value = 10
		s.Add(value)

		if !(s.Contains(value)) {
			t.Errorf("Set should contain: %d, however it was not found as a member.", value)
		}

		if s.Size() != 1 {
			t.Errorf("Set should have only 1 member, however: %d members were found.", s.Size())
		}
	}
}

func TestRemove(t *testing.T) {
	var s = set.NewIntSet()
	if s != nil {
		for val := -4; val <= 1000; val++ {
			s.Add(val)
		}
	}

	var sizeBefore = s.Size()
	// Attempt to remove a non member.
	s.Remove(100000)
	var sizeAfter = s.Size()
	if sizeBefore != sizeAfter {
		t.Errorf("Set size changed after attempting to remove a NON-member.")
	}

	// Remove a member.
	s.Remove(0)
	var sizeAfterMemberRemoval = s.Size()

	if sizeBefore == sizeAfterMemberRemoval {
		t.Errorf("Set size did not change after member removal.")
	}
}

func TestGetLastValue(t *testing.T) {
	var s = set.NewIntSet()
	if s != nil {
		// Given an empty set
		v, e := s.GetLastValue()
		// Then no last value exists, so and error should be returned.
		if e == nil {
			t.Errorf("GetLastValue return %d, when error should have been reported.", v)
		}

		// Given a set with 1 value
		var value = 32
		s.Add(value)
		// Then that value should be returned.
		v, e = s.GetLastValue()
		if e != nil {
			t.Errorf("GetLastValue reported an error!")
		}

		if v != value {
			t.Errorf("GetLastValue did not return expected value(%d), returned(%d)!", value, v)
		}

		// Given more than 1 item in the set.
		s.Add(500)
		// Then git last value should return an error.
		v, e = s.GetLastValue()
		if e == nil {
			t.Errorf("GetLastValue should have reported an error since more than 1 items is in the set!")
		}

	}
}

func TestSetEquality(t *testing.T) {
	s := set.NewIntSet()
	if s != nil {
		s2 := set.NewIntSet()
		if s2 != nil {
			if !s.Equals(s2) {
				t.Errorf("Empty sets should be equal!")
			}

			s.Add(1)
			s2.Add(1)
			if !s.Equals(s2) {
				t.Errorf("Sets with single value should be equal.")
			}

			s.Add(2)
			if s.Equals(s2) {
				t.Errorf("Different sets should not be equal.")
			}
		}
	}
}

func TestClear(t *testing.T) {
	s := set.NewIntSet()
	if s != nil {
		if s.Size() != 0 {
			t.Errorf("Set size is not zero.")
		}
		s.Add(100)
		if s.Size() != 1 {
			t.Errorf("Set size is not one.")
		}
		s.Clear()
		if s.Size() != 0 {
			t.Errorf("Cleared Set does not have size zero.")
		}
	}
}
