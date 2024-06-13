package commnd

import "fmt"

type Light struct {
}

func (l *Light) On() {
	fmt.Println("light on")
}

func (l *Light) Off() {
	fmt.Println("light off")
}
