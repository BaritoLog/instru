package instru

import (
	"testing"

	. "github.com/BaritoLog/go-boilerplate/testkit"
)

func TestCounter_Event(t *testing.T) {

	metric := NewCounterMetric()
	counter := NewCounter(metric)

	counter.Event("success")
	counter.Event("success")
	counter.Event("success")
	counter.Event("error")
	counter.Event("error")

	FatalIf(t, metric.Total != 5, "wrong metrix.total")
	ec, _ := metric.Load("success")
	FatalIf(t, ec != 3, "wrong metrix.events.success")
	ec, _ = metric.Load("error")
	FatalIf(t, ec != 2, "wrong metrix.events.success")

}
