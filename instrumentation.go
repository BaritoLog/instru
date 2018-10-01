package instru

import (
	"encoding/json"
	"time"
)

type Instrumentation interface {
	Evaluate(label string) Evaluation
	Count(label string) Counter
	Metric(label string) *InstrumentationMetric
	Metrics() map[string]*InstrumentationMetric
	ToJson() ([]byte, error)
	Flush()
}

type instrumentation struct {
	metrics map[string]*InstrumentationMetric
}

func NewInstrumentation() Instrumentation {
	return &instrumentation{
		metrics: make(map[string]*InstrumentationMetric),
	}
}

func (i *instrumentation) Evaluate(label string) Evaluation {
	return NewEvaluation(
		time.Now(),
		i.Metric(label).EvaluationMetric(),
	)
}

func (i *instrumentation) Count(label string) Counter {
	return NewCounter(
		i.Metric(label).CounterMetric(),
	)
}

// Return InstrumentationMetric. It will create empty metric if not exist
func (i *instrumentation) Metric(label string) *InstrumentationMetric {
	instrMetric, ok := i.metrics[label]
	if !ok {
		instrMetric = NewInstrumentationMetric()
		i.metrics[label] = instrMetric
	}

	return instrMetric
}

func (i *instrumentation) Metrics() map[string]*InstrumentationMetric {
	return i.metrics
}

func (i *instrumentation) ToJson() (data []byte, err error) {

	metrics := make(map[string]interface{})
	for key, value := range i.Metrics() {
		metricResult := make(map[string]interface{})
		value.Data.Range(func(key interface{}, value interface{}) bool {
			metricResult[key.(string)] = value
			return true
		})

		metrics[key] = metricResult
	}

	result := make(map[string]interface{})
	result["metrics"] = metrics

	data, err = json.Marshal(result)
	return

}

func (i *instrumentation) Flush() {
	i.metrics = make(map[string]*InstrumentationMetric)
}
