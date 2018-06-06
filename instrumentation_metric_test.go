package instru

import (
	"testing"

	. "github.com/BaritoLog/go-boilerplate/testkit"
)

func TestInstrumentationMetric(t *testing.T) {
	metric := NewInstrumentationMetric()
	metric.Put("some-key", "some-value")

	FatalIf(t, metric.Get("some-key") != "some-value", "wrong value for some-key")
	FatalIf(t, metric.Get("wrong-wrong") != nil, "wrong value for ")
}

func TestInstrumentationMetric_EvaluationMetric(t *testing.T) {
	metric := NewInstrumentationMetric()
	evalMetric := metric.EvaluationMetric()

	FatalIf(t, evalMetric == nil, "evalMetric can't be nil")
	FatalIf(t, metric.EvaluationMetric() != evalMetric, "metric.EvaluationMetric() must return same metric")
}

func TestInstrumentationMetric_CounterMetric(t *testing.T) {
	metric := NewInstrumentationMetric()
	counterMetric := metric.CounterMetric()

	FatalIf(t, counterMetric == nil, "counterMetric can't be nil")
	FatalIf(t, metric.CounterMetric() != counterMetric, "metric.CounterMetric() must return same metric")
}
