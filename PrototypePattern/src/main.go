package main

import (
	p "PrototypePattern/src/prototype"
	"fmt"
)

func main() {
	file1 := &p.File{Name: "File1"}
	file2 := &p.File{Name: "File2"}
	file3 := &p.File{Name: "File3"}

	folder1 := &p.Folder{
		Children: []p.INode{file1, file2},
		Name:     "Folder1",
	}

	folder2 := &p.Folder{
		Children: []p.INode{folder1, file2, file3},
		Name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
}
