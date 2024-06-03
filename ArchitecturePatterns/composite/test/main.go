package main

import (
	"designPattern/ArchitecturePatterns/composite"
	"fmt"
)

// 组合模式适用于生成树形结果的问题，例如：文件系统
func main() {
	file1 := &composite.File{Name: "file1"}
	file2 := &composite.File{Name: "file2"}
	file3 := &composite.File{Name: "file3"}

	folder1 := &composite.Folder{Name: "folder1", Children: []composite.Composite{file1}}
	folder1.Add(file2)

	folder2 := &composite.Folder{Name: "folder2", Children: []composite.Composite{file3}}
	folder2.Add(folder1)

	fmt.Println("下面是folder1")
	folder1.Print(" ")
	fmt.Println("下面是folder2")
	folder2.Print(" ")

}
