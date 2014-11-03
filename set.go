package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

type Set struct {
	store map[string]interface{}
}

func NewSet() *Set {
	return &Set{store: map[string]interface{}{}}
}

func (s *Set) Add(object interface{}) bool {
	h := sha1.New()
	io.WriteString(h, fmt.Sprintf("%v", object))
	hash := fmt.Sprintf("% x", h.Sum(nil))
	_, ok := s.store[hash]
	s.store[hash] = object
	return !ok
}

func (s *Set) Len() int {
	return len(s.store)
}

func (s *Set) GetAll() []interface{} {
	var objects []interface{}
	for _, object := range s.store {
		objects = append(objects, object)
	}
	return objects
}

func (s *Set) Strings() []string {
	var strings []string
	for _, object := range s.store {
		strings = append(strings, object.(string))
	}
	return strings
}
