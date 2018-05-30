package instru

import (
	"time"
)

var DefaultInstrumentation Instrumentation
var DefaultExposer Exposer
var DefaultCallback Callback

var ErrorCh = make(chan error)
var OnErrorFunc func(error)

var CallbackTick <-chan time.Time
var CallbackStop chan int

func init() {
	DefaultInstrumentation = NewInstrumentation()
	go loopError()
}

func Evaluate(name string) Evaluation {
	return DefaultInstrumentation.Evaluate(name)
}

func Count(name string) Counter {
	return DefaultInstrumentation.Count(name)
}

func Expose(exposer Exposer) {
	DefaultExposer = exposer
	go func() {
		ErrorCh <- DefaultExposer.Expose(DefaultInstrumentation)
	}()
}

func ExposeWithRestful(addr string) {
	Expose(NewRestfulExposer(addr))
}

func StopExpose() {
	if DefaultExposer != nil {
		DefaultExposer.Stop()
	}
	DefaultExposer = nil
}

func SetCallback(interval time.Duration, callback Callback) {
	CallbackTick = time.Tick(interval)
	DefaultCallback = callback
	CallbackStop = make(chan int)

	go loopCallback()
}

func SetWebCallback(interval time.Duration, url string) {
	SetCallback(interval, NewWebCallback(url))
}

func UnsetCallback() {
	if CallbackStop != nil {
		CallbackStop <- 1
	}

	CallbackStop = nil
	DefaultCallback = nil
	CallbackTick = nil

}

func loopCallback() {
	for {
		select {
		case <-CallbackTick:
			err := DefaultCallback.OnCallback(DefaultInstrumentation)
			fireError(err)
		case <-CallbackStop:
			return
		}
	}
}

func loopError() {
	for err := range ErrorCh {
		fireError(err)
	}
}

func fireError(err error) {
	if OnErrorFunc != nil {
		OnErrorFunc(err)
	}
}
