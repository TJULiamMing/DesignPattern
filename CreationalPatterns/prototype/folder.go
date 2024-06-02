package prototype

import "fmt"

type Folder struct {
	Name     string
	Children []Prototype
}

func (folder *Folder) Print(indentation string) {
	fmt.Println(indentation + folder.Name)
	for _, child := range folder.Children {
		child.Print(indentation + indentation)
	}
}

func (folder *Folder) Clone() Prototype {
	cloned := &Folder{
		Name: folder.Name,
	}

	for _, child := range folder.Children {
		cloned.Children = append(cloned.Children, child.Clone())
	}

	return cloned
}
