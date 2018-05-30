package instru

type Callback interface {
	OnCallback(inst Instrumentation)
}

type dummyCallback struct {
	Instr Instrumentation
}

func (c *dummyCallback) OnCallback(instr Instrumentation) {
	c.Instr = instr
}
