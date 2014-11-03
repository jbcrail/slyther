package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestEmptySet(t *testing.T) {
	s := NewSet()
	if s.Len() != 0 {
		t.Errorf("NewSet().Len() = %v, want %v", s.Len(), 0)
	}
}

func TestAddUniqueItems(t *testing.T) {
	s := NewSet()
	s.Add("foo")
	s.Add("bar")
	s.Add("baz")
	if s.Len() != 3 {
		t.Errorf("Len() = %v, want %v", s.Len(), 3)
	}
}

func TestAddDuplicateItems(t *testing.T) {
	s := NewSet()
	s.Add("foo")
	s.Add("foo")
	s.Add("bar")
	s.Add("bar")
	s.Add("bar")
	s.Add("baz")
	if s.Len() != 3 {
		t.Errorf("Len() = %v, want %v", s.Len(), 3)
	}
}

func TestStrings(t *testing.T) {
	s := NewSet()
	s.Add("foo")
	s.Add("bar")
	s.Add("baz")
	actual := s.Strings()
	expected := []string{"foo", "bar", "baz"}
	sort.Strings(actual)
	sort.Strings(expected)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Strings() = %v, want %v", actual, expected)
	}
}
