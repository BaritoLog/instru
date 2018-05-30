package instru

import (
	"time"
)

var Instance = NewInstrumentation()
var ExposerInstance Exposer
var ErrorCh = make(chan error)
var OnErrorFunc func(error)

var CallbackInstance Callback
var CallbackTick <-chan time.Time
var CallbackStop chan int

func init() {
	go loopError()
}

func Evaluate(name string) Evaluation {
	return Instance.Evaluate(name)
}

func Count(name string) Counter {
	return Instance.Count(name)
}

func Expose(exposer Exposer) {
	ExposerInstance = exposer
	go func() {
		ErrorCh <- ExposerInstance.Expose(Instance)
	}()
}

func ExposeWithRestful(addr string) {
	Expose(NewRestfulExposer(addr))
}

func StopExpose() {
	if ExposerInstance != nil {
		ExposerInstance.Stop()
	}
	ExposerInstance = nil
}

func SetCallback(interval time.Duration, callback Callback) {
	CallbackTick = time.Tick(interval)
	CallbackInstance = callback
	CallbackStop = make(chan int)

	go loopCallback()
}

func UnsetCallback() {
	if CallbackStop != nil {
		CallbackStop <- 1
	}

	CallbackStop = nil
	CallbackInstance = nil
	CallbackTick = nil

}

func loopCallback() {
	for {
		select {
		case <-CallbackTick:
			err := CallbackInstance.OnCallback(Instance)
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
