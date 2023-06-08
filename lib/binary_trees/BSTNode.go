package binarytrees

type BSTNode[K comparable, V any] struct {
	key   K
	val   V
	left  *BSTNode[K, V]
	right *BSTNode[K, V]
}

func (b BSTNode[K, V]) IsLeaf() bool {
	return b.left == nil && b.right == nil
}

type VarBinNode interface {
	IsLeaf() bool
}

// Leaf Node //
type Operand string

type LeafNode struct {
	operand Operand
}

func (b LeafNode) IsLeaf() bool {
	return true
}

// Internal Node //
type Operator string

type IntlNode struct {
	operator Operator
	left     VarBinNode
	right    VarBinNode
}

func (b IntlNode) IsLeaf() bool {
	return false
}
