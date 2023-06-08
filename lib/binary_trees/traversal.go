package binarytrees

import "fmt"

func Preorder[K, V comparable](root *BSTNode[K, V]) {
	if root == nil {
		return
	}

	visit(root)
	Preorder(root.left)
	Preorder(root.right)
}

func Inorder[K, V comparable](root *BSTNode[K, V]) {
	if root == nil {
		return
	}

	Inorder(root.left)
	visit(root)
	Inorder(root.right)
}

func Postorder[K, V comparable](root *BSTNode[K, V]) {
	if root == nil {
		return
	}

	Postorder(root.left)
	Postorder(root.right)
	visit(root)
}

func visit[K, V comparable](node *BSTNode[K, V]) {
	fmt.Println(node.val)
}
