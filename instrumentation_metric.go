package instru

var (
	KeyEvaluationTime = "_evaluation_time"
	KeyCounter        = "_counter"
)

type InstrumentationMetric map[string]interface{}

// NewInstrumentationMetric create a new instance of InstrumentationMetric
func NewInstrumentationMetric() InstrumentationMetric {
	return InstrumentationMetric(make(map[string]interface{}))
}

// Return EvaluationMetric. It will create empty metric if not available.
func (m InstrumentationMetric) EvaluationMetric() *EvaluationMetric {
	metric, ok := m[KeyEvaluationTime]
	if !ok {
		metric = NewEvaluationMetric()
		m[KeyEvaluationTime] = metric
	}

	return metric.(*EvaluationMetric)
}

// Return CounterMetric. It will create empty metric if not available.
func (m InstrumentationMetric) CounterMetric() *CounterMetric {
	metric, ok := m[KeyCounter]
	if !ok {
		metric = NewCounterMetric()
		m[KeyCounter] = metric
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
