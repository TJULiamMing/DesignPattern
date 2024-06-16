package main

import (
	"designPattern/BehavioralPattern/memento"
	"fmt"
)

func main() {
	editor := &memento.TextEditor{}
	caretaker := &memento.Caretaker{}

	editor.SetState("State #1")
	caretaker.AddMemento(editor.Save())

	editor.SetState("State #2")
	caretaker.AddMemento(editor.Save())

	editor.SetState("State #3")
	caretaker.AddMemento(editor.Save())

	caretaker.Undo(editor)
	fmt.Println(editor.GetState())

	caretaker.Redo(editor)
	fmt.Println(editor.GetState())
}
