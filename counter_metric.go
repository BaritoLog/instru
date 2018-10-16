package instru

import (
	"sync"
)

type CounterMetric struct {
	sync.RWMutex
	Total  int64            `json:"total"`
	Events map[string]int64 `json:"events"`
}

func NewCounterMetric() *CounterMetric {
	return &CounterMetric{
		Events: make(map[string]int64),
	}
}

func (m *CounterMetric) TotalEvent(name string) int64 {
	return m.Events[name]
}

func (m *CounterMetric) Load(key string) (value int64, ok bool) {
	m.RLock()
	result, ok := m.Events[key]
	m.RUnlock()
	return result, ok
}

func (m *CounterMetric) Delete(key string) {
	m.Lock()
	delete(m.Events, key)
	m.Unlock()
}

func (m *CounterMetric) Store(key string, value int64) {
	m.Lock()
	m.Events[key] = value
	m.Unlock()
}
