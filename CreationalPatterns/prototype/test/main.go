package main

import "designPattern/CreationalPatterns/prototype"

func main() {
	file1 := &prototype.File{Name: "file1"}
	file2 := &prototype.File{Name: "file2"}
	file3 := &prototype.File{Name: "file3"}

	folder1 := &prototype.Folder{Name: "folder1", Children: []prototype.Prototype{file1}}
	folder2 := &prototype.Folder{Name: "folder2", Children: []prototype.Prototype{folder1, file2, file3}}

	folder1.Print(" ")
	clonedfolder := folder2.Clone()
	clonedfolder.Print(" ")

}
