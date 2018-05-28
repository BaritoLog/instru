package instru

type CounterMetric struct {
	Total  int32
	Events map[string]int32
}

func NewCounterMetric() *CounterMetric {
	return &CounterMetric{
		Events: make(map[string]int32),
	}
}
