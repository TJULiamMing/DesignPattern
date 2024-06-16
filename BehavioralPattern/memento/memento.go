package memento

type Memento interface {
	GetState() string
}

type textMemento struct {
	state string
}

func (t *textMemento) GetState() string {
	return t.state
}
