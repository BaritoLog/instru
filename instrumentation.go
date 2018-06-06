package instru

import (
	"time"
)

type Instrumentation interface {
	Evaluate(label string) Evaluation
	Count(label string) Counter
	Metric(label string) InstrumentationMetric
	Flush()
}

type instrumentation struct {
	Metrics map[string]InstrumentationMetric `json:"metrics"`
}

func NewInstrumentation() Instrumentation {
	return &instrumentation{
		Metrics: make(map[string]InstrumentationMetric),
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
func (i *instrumentation) Metric(label string) InstrumentationMetric {
	instrMetric, ok := i.Metrics[label]
	if !ok {
		instrMetric = NewInstrumentationMetric()
		i.Metrics[label] = instrMetric
	}

	return instrMetric
}

func (i *instrumentation) Flush() {
	i.Metrics = make(map[string]InstrumentationMetric)
}
