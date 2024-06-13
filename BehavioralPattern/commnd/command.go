package commnd

type Command interface {
	Execute()
}

type LightOn struct {
	L Light
}

func (c *LightOn) Execute() {
	c.L.On()
}

type LightOff struct {
	L Light
}

func (c *LightOff) Execute() {
	c.L.Off()
}
