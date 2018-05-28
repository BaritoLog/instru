package instru

var Instance = NewInstrumentation()
var ExposerInstance Exposer
var ErrorCh = make(chan error)
var OnErrorFunc func(error)

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

func loopError() {
	for err := range ErrorCh {
		if OnErrorFunc != nil {
			OnErrorFunc(err)
		}
	}
}
