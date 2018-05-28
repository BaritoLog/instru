package instru

var Instance = NewInstrumentation()

func Evaluate(name string) Evaluation {
	return Instance.Evaluate(name)
}

func Count(name string) Counter {
	return Instance.Count(name)
}
