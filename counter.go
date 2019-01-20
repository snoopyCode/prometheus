package prometheus

import (
	"reflect"
	"sync"
)

// https://blog.golang.org/go-maps-in-action#TOC_6.
// http://stackoverflow.com/questions/1823286/singleton-in-go

type counter struct {
	mu     sync.Mutex
	values map[string]int64
}

// Get returns the value of the counter
func (s *counter) Get(key string) int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.values[key]
}

// Incr increases the counter
func (s *counter) Incr(key string) int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[key]++
	return s.values[key]
}

// Keys returns all keys
func (s *counter) Keys() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	keys := reflect.ValueOf(s.values).MapKeys()
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}
	return strkeys
}

// All returns the values of the counter as map
func (s *counter) All() map[string]int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.values
}
