package instru

import (
	"time"
)

var DefaultInstrumentation Instrumentation = NewInstrumentation()
var DefaultExposer Exposer
var DefaultCallback Callback

var errorCh = make(chan error)
var callbackStop chan int

var OnErrorFunc func(error)

func init() {
	go loopError()
}

// Create new Evaluation in default instrumentation
func Evaluate(label string) Evaluation {
	return DefaultInstrumentation.Evaluate(label)
}

// Create new Counter in default instrumentation
func Count(label string) Counter {
	return DefaultInstrumentation.Count(label)
}

// Return metric of default instrumentation
func Metric(label string) InstrumentationMetric {
	return DefaultInstrumentation.Metric(label)
}

// Expose default instrumentation
func Expose(exposer Exposer) {
	go func() {
		errorCh <- exposer.Expose(DefaultInstrumentation)
	}()
	DefaultExposer = exposer
}

// Expose default instrumentation with restful api
func ExposeWithRestful(addr string) {
	Expose(NewRestfulExposer(addr))
}

// Stop expose default instrumentation
func StopExpose() {
	if DefaultExposer != nil {
		DefaultExposer.Stop()
	}
	DefaultExposer = nil
}

// Set callback for default instrumentation
func SetCallback(interval time.Duration, callback Callback) {
	tick := time.Tick(interval)

	DefaultCallback = callback
	callbackStop = make(chan int)

	go func() {
		for {
			select {
			case <-tick:
				err := callback.OnCallback(DefaultInstrumentation)
				fireError(err)
			case <-callbackStop:
				return
			}
		}
	}()
}

// Set web callback for default instrumentation
func SetWebCallback(interval time.Duration, url string) {
	SetCallback(interval, NewWebCallback(url))
}

func UnsetCallback() {
	if callbackStop != nil {
		callbackStop <- 1
	}

	callbackStop = nil
	DefaultCallback = nil
}

func GetEventCount(label, event string) int64 {
	return DefaultInstrumentation.Metric(label).CounterMetric().TotalEvent(event)
}

func Flush() {
	DefaultInstrumentation.Flush()
}

func loopError() {
	for err := range errorCh {
		fireError(err)
	}
}

func fireError(err error) {
	if OnErrorFunc != nil {
		OnErrorFunc(err)
	}
}
