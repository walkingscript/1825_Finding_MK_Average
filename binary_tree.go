package main

type ItemType = int

type BinTree struct {
	Root       *Node
	ItemsCount int
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
	tree.ItemsCount++
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

func (tree *BinTree) Min() ItemType {
	return tree.Root.min()
}

func (tree *BinTree) Max() ItemType {
	return tree.Root.max()
}

func (node *Node) min() ItemType {
	currentNode := node
	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}
	return currentNode.Value
}

func (node *Node) max() ItemType {
	currentNode := node
	for currentNode.Right != nil {
		currentNode = currentNode.Right
	}
	return currentNode.Value
}

func (tree *BinTree) PopLeft() (value ItemType) {
	defer func() { tree.ItemsCount-- }()
	return tree.Root.popLeft()
}

func (tree *BinTree) PopRight() (value ItemType) {
	defer func() { tree.ItemsCount-- }()
	return tree.Root.popRight()
}

func (node *Node) popLeft() (value ItemType) {
	if node == nil {
		panic("popLeft: can't pop value; not enougth items")
	}
	currentNode := node
	if currentNode.Left != nil { // can be removed because in terms m >= 3
		for currentNode.Left.Left != nil {
			currentNode = currentNode.Left
		}
	}
	value = currentNode.Left.Value
	currentNode.Left = nil
	return value
}

func (node *Node) popRight() (value ItemType) {
	if node == nil {
		panic("popRight: can't pop value; not enougth items")
	}
	currentNode := node
	if currentNode.Right != nil { // can be removed because in terms m >= 3
		for currentNode.Right.Right != nil {
			currentNode = currentNode.Right
		}
	}
	value = currentNode.Right.Value
	currentNode.Right = nil
	return value
}

func (tree *BinTree) Sum() (sum ItemType) {
	for _, n := range tree.Root.GetSortedArray() {
		sum += n
	}
	return
}

func (tree *BinTree) Reset() {
	tree.Root = nil
}
