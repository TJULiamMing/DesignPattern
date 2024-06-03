package composite

import "fmt"

type File struct {
	Name string
}

func (f File) Print(inentation string) {
	fmt.Println(inentation + f.Name)
}
