package instru

type CounterMetric struct {
	Total  int64            `json:"total"`
	Events map[string]int64 `json:"events"`
}

func NewCounterMetric() *CounterMetric {
	return &CounterMetric{
		Events: make(map[string]int64),
	}
}

func (m *CounterMetric) TotalEvent(name string) int64 {
	return m.Events[name]
}
