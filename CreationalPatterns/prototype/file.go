package prototype

import "fmt"

type File struct {
	Name string
}

func (f File) Print(inentation string) {
	fmt.Println(inentation + f.Name)
}

func (f File) Clone() Prototype {
	return &File{Name: f.Name}
}
