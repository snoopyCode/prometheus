package prometheus

import (
	"reflect"
	"sync"
)

// https://blog.golang.org/go-maps-in-action#TOC_6.
// http://stackoverflow.com/questions/1823286/singleton-in-go

// Gauge is a map for gauges
type Gauge struct {
	mu     sync.Mutex
	values map[string]string
}

// NewGauge creates new Gauge
func NewGauge() *Gauge {
	return &Gauge{
		values: make(map[string]string),
	}
}

// Get returns gauge value
func (s *Gauge) Get(key string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.values[key]
}

// Set sets gauge value
func (s *Gauge) Set(key string, value string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[key] = value
	return s.values[key]
}

// Keys returns all keys
func (s *Gauge) Keys() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	keys := reflect.ValueOf(s.values).MapKeys()
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}
	return strkeys
}
