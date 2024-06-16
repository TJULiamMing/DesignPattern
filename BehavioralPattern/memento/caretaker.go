package memento

type Caretaker struct {
	mementos     []Memento
	currentIndex int
}

func (c *Caretaker) AddMemento(m Memento) {
	c.mementos = append(c.mementos, m)
	c.currentIndex = len(c.mementos) - 1
}

func (c *Caretaker) Undo(t *TextEditor) {
	if c.currentIndex > 0 {
		c.currentIndex--
		m := c.mementos[c.currentIndex]
		t.Restore(m)
	}
}

func (c *Caretaker) Redo(t *TextEditor) {
	if c.currentIndex < len(c.mementos)-1 {
		c.currentIndex++
		m := c.mementos[c.currentIndex]
		t.Restore(m)
	}
}
