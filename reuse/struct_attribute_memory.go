package reuse

type Node struct{
	Key int
}

func NewNode(key int) Node {
	return Node{Key: key}
}

type Tree struct {
	Root Node
}

func NewTree(root Node) Tree {
	return Tree{Root: root}
}