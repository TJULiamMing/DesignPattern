package template

import "fmt"

type LoL struct {
	Name string
}

func (l *LoL) Init() error {
	fmt.Printf("LoL init in time\n")
	return nil
}

func (l *LoL) Play() {
	fmt.Printf("play LOL\n")
}

func (l *LoL) Stop(time int) {
	fmt.Printf("LoL stop in %d\n", time)
}
