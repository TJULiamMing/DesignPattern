package memento

type Originator interface {
	Save() Memento
	Restore(m Memento)
}

type TextEditor struct {
	state string
}

func (t *TextEditor) Save() Memento {
	return &textMemento{state: t.state}
}

func (t *TextEditor) Restore(m Memento) {
	t.state = m.GetState()
}

func (t *TextEditor) SetState(state string) {
	t.state = state
}

func (t *TextEditor) GetState() string {
	return t.state
}
