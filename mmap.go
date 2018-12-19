// Package mmap implements a concurrent-safe map.
//
// Simply the basic operations on maps but guarded with RWMutex.

package mmap

import (
	"sync"
)

type mmap struct {
	m map[interface{}]interface{}
	sync.RWMutex
}

// New returns a pointer to newly created instance of mmap.
func New() *mmap {
	m := new(mmap)
	m.m = make(map[interface{}]interface{})
	return m
}

// Set assigns value to specified key.
func (m *mmap) Set(key interface{}, value interface{}) {
	m.Lock()
	m.m[key] = value
	m.Unlock()
}

// Get returns value assigned to key.
func (m *mmap) Get(key interface{}) interface{} {
	m.Lock()
	value, _ := m.m[key]
	m.Unlock()
	return value
}

// Has returns true if key exists in map. False otherwise.
func (m *mmap) Has(key interface{}) bool {
	m.Lock()
	_, ok := m.m[key]
	m.Unlock()
	return ok
}

// Remove key with its value from map.
func (m *mmap) Remove(key interface{}) {
	m.Lock()
	delete(m.m, key)
	m.Unlock()
}

// Count returns total number of keys in map.
func (m *mmap) Count() int {
	m.Lock()
	count := len(m.m)
	m.Unlock()
	return count
}

// Items returns a slice of values.
func (m *mmap) Items() []interface{} {
	values := make([]interface{}, 0)

	m.Lock()
	for _, value := range m.m {
		values = append(values, value)
	}
	m.Unlock()

	return values
}

// Iterate takes a function as an argument and executes it in loop over map.
func (m *mmap) Iterate(f func(key interface{}, value interface{})) {
	m.Lock()
	for key, value := range m.m {
		f(key, value)
	}
	m.Unlock()
}
