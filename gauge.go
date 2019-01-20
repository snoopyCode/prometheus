package prometheus

import (
	"fmt"
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

// BuildGaugeMetric returns a gauge metric as string
func BuildGaugeMetric(metric string, metricsDesc string, label string, labelValue string, value string) string {
	return fmt.Sprintf("\n# HELP %s %s\n# TYPE %s gauge\n%s{%s=\"%s\"} %s", metric, metricsDesc, metric, metric, label, labelValue, value)
}
