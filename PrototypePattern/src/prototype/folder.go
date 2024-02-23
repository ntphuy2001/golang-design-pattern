package prototype

import "fmt"

type Folder struct {
	Name     string
	Children []INode
}

func (f *Folder) Print(s string) {
	fmt.Println(s + f.Name)
	for _, node := range f.Children {
		node.Print(s + s)
	}
}

func (f *Folder) Clone() INode {
	cloneFolder := &Folder{
		Name:     f.Name + "_Clone",
		Children: nil,
	}

	for _, node := range f.Children {
		cloneFolder.Children = append(cloneFolder.Children, node.Clone())
	}

	return cloneFolder
}
