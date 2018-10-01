package instru

import (
	"sync"
)

const (
	keyEvaluationTime = "_evaluation_time"
	keyCounter        = "_counter"
)

type InstrumentationMetric struct {
	Data sync.Map
}

// NewInstrumentationMetric create a new instance of InstrumentationMetric
func NewInstrumentationMetric() *InstrumentationMetric {
	return &InstrumentationMetric{}
}

// Return EvaluationMetric. It will create empty metric if not available.
func (m *InstrumentationMetric) EvaluationMetric() *EvaluationMetric {
	metric, ok := m.Data.Load(keyEvaluationTime)
	if !ok {
		metric = NewEvaluationMetric()
		m.Put(keyEvaluationTime, metric)
	}

	return metric.(*EvaluationMetric)
}

// Return CounterMetric. It will create empty metric if not available.
func (m *InstrumentationMetric) CounterMetric() *CounterMetric {
	metric, ok := m.Data.Load(keyCounter)
	if !ok {
		metric = NewCounterMetric()
		m.Put(keyCounter, metric)
	}
	return metric.(*CounterMetric)
}

// Put value to InstrumentationMetric
func (m *InstrumentationMetric) Put(key string, val interface{}) {
	m.Data.Store(key, val)
}

// Get value from InstrumentationMetric
func (m *InstrumentationMetric) Get(key string) interface{} {
	val, _ := m.Data.Load(key)
	return val
}
