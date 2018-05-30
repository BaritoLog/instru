package instru

import (
	"reflect"
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

func TestExposeWithRestful(t *testing.T) {
	ExposeWithRestful(":65501")
	FatalIf(t, ExposerInstance == nil, "ExposerInstance can't be nil")
	FatalIf(t, reflect.TypeOf(ExposerInstance).String() != "*instru.restfulExposer", "wrong type ExposerInstance")

	timekit.Sleep("1ms")

	StopExpose()
	FatalIf(t, ExposerInstance != nil, "ExposerInstance must be nil")
}

func TestExposeWithRestful_Error(t *testing.T) {
	var err error
	OnErrorFunc = func(err0 error) {
		err = err0
	}

	ExposeWithRestful(":65502")
	ExposeWithRestful(":65502") // same address
	defer StopExpose()

	timekit.Sleep("1ms")

	FatalIfWrongError(t, err, "listen tcp :65502: bind: address already in use")
}

func TestSetCallback(t *testing.T) {
	callback := &dummyCallback{}

	SetCallback(timekit.Duration("2ms"), callback)
	FatalIf(t, CallbackInstance != callback, "CallbackInstanct can't be nil")

	timekit.Sleep("5ms")
	FatalIf(t, callback.Instr != Instance, "callback.instrument is wrong")

	UnsetCallback()
	FatalIf(t, CallbackStop != nil, "CallbackStop must be nil")
	FatalIf(t, CallbackInstance != nil, "CallbackInstance must be nil")
	FatalIf(t, CallbackTick != nil, "CallbackTick must be nil")
}
