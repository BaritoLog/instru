package instru

import (
	"testing"

	. "github.com/BaritoLog/go-boilerplate/testkit"
	"github.com/BaritoLog/go-boilerplate/timekit"
)

func TestEvaluate(t *testing.T) {
	eval := Evaluate("eval01")
	timekit.Sleep("100ms")
	eval.Done()

	eval = Evaluate("eval01")
	timekit.Sleep("300ms")
	eval.Done()

	metric := Instance.GetEvaluationMetric("eval01")
	FatalIf(t, metric.Count != 2, "wrong metric count")
}

func TestCounter(t *testing.T) {
	Count("count01").Event("success")
	Count("count01").Event("fail")
	Count("count01").Event("error")
	Count("count01").Event("fail")

	metric := Instance.GetCounterMetric("count01")
	FatalIf(t, metric.Total != 4, "wrong metric total")
	FatalIf(t, len(metric.Events) != 3, "wrong metric event")
}
