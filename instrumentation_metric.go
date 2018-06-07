package instru

const (
	keyEvaluationTime = "_evaluation_time"
	keyCounter        = "_counter"
)

type InstrumentationMetric map[string]interface{}

// NewInstrumentationMetric create a new instance of InstrumentationMetric
func NewInstrumentationMetric() InstrumentationMetric {
	return InstrumentationMetric(make(map[string]interface{}))
}

// Return EvaluationMetric. It will create empty metric if not available.
func (m InstrumentationMetric) EvaluationMetric() *EvaluationMetric {
	metric, ok := m[keyEvaluationTime]
	if !ok {
		metric = NewEvaluationMetric()
		m[keyEvaluationTime] = metric
	}

	return metric.(*EvaluationMetric)
}

// Return CounterMetric. It will create empty metric if not available.
func (m InstrumentationMetric) CounterMetric() *CounterMetric {
	metric, ok := m[keyCounter]
	if !ok {
		metric = NewCounterMetric()
		m[keyCounter] = metric
	}
	return metric.(*CounterMetric)
}

// Put value to InstrumentationMetric
func (m InstrumentationMetric) Put(key string, val interface{}) {
	m[key] = val
}

// Get value from InstrumentationMetric
func (m InstrumentationMetric) Get(key string) interface{} {
	return m[key]
}
