package main

type ItemType = int

type BinTree struct {
	Root *Node
}

type Node struct {
	Value       ItemType
	Left, Right *Node
}

func (tree *BinTree) Insert(value ItemType) {
	if tree.Root == nil {
		tree.Root = &Node{Value: value}
	} else {
		tree.Root.insert(value)
	}
}

func (node *Node) insert(value ItemType) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &Node{Value: value}
			return
		}
		node.Left.insert(value)
	} else {
		if node.Right == nil {
			node.Right = &Node{Value: value}
			return
		}
		node.Right.insert(value)
	}
}

func (node *Node) GetSortedArray() []ItemType {
	var results []ItemType = make([]ItemType, 0, 3)
	if node.Left != nil {
		results = append(results, node.Left.GetSortedArray()...)
	}
	results = append(results, node.Value)
	if node.Right != nil {
		results = append(results, node.Right.GetSortedArray()...)
	}
	return results
}

func (tree *BinTree) Reset() {
	tree.Root = nil
}
