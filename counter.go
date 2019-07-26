package prometheus

import (
	"fmt"
	"reflect"
	"sync"
)

// https://blog.golang.org/go-maps-in-action#TOC_6.
// http://stackoverflow.com/questions/1823286/singleton-in-go

// Counter with a map of counters
type Counter struct {
	mu     sync.Mutex
	values map[string]int64
}

// NewCounter creates new Counter
func NewCounter() *Counter {
	return &Counter{
		values: make(map[string]int64),
	}
}

// Get returns the value of the counter
func (s *Counter) Get(key string) int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.values[key]
}

// Incr increases the counter
func (s *Counter) Incr(key string) int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[key]++
	return s.values[key]
}

// Keys returns all keys
func (s *Counter) Keys() []string {
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
func (s *Counter) All() map[string]int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.values
}

// BuildCounterMetric returns a counter metric as string
func BuildCounterMetric(metric string, metricsDesc string, label string, labelValue string, value int64) string {
	return fmt.Sprintf("\n# HELP %s %s\n# TYPE %s counter\n%s{%s=\"%s\"} %d", metric, metricsDesc, metric, metric, label, labelValue, value)
}

// BuildCounterTwoLabelMetric returns a counter metric as string with 2 labels
func BuildCounterTwoLabelMetric(metric string, metricsDesc string, label1 string, label1Value string, label2 string, label2Value string, value int64) string {
	return fmt.Sprintf("\n# HELP %s %s\n# TYPE %s counter\n%s{%s=\"%s\",%s=\"%s\"} %d", metric, metricsDesc, metric, metric, label1, label1Value, label2, label2Value, value)
}
