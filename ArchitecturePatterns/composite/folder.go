package composite

import "fmt"

type Folder struct {
	Name     string
	Children []Composite
}

func (folder *Folder) Print(indentation string) {
	fmt.Println(indentation + folder.Name)
	for _, child := range folder.Children {
		child.Print(indentation + indentation)
	}
}

func (folder *Folder) Add(c Composite) {
	folder.Children = append(folder.Children, c)
}
