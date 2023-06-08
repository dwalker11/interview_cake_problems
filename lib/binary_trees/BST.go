package binarytrees

import "fmt"

type BST[K comparable, V any] struct {
	root  *BSTNode[K, V]
	count int
}

func (b *BST[K, V]) Clear() {
	clearHelp(b.root)
	b.root = nil
	b.count = 0
}

func (b *BST[K, V]) Find(key K) (V, bool) {
	return findhelp(b.root, key)
}

func (b *BST[K, V]) Insert(key K, value V) {
	b.root = insertHelp(b.root, key, value)
	b.count++
}

func (b *BST[K, V]) Remove(key K) (V, bool) {
	temp, ok := findhelp[K, V](b.root, key)

	if ok {
		b.root = removehelp(b.root, key)
		b.count--
	}

	return temp, ok
}

func (b *BST[K, V]) RemoveAny() (V, bool) {
	if b.root == nil {
		var zero V
		return zero, false
	}

	temp := b.root.val
	b.root = removehelp[K, V](b.root, b.root.key)
	b.count--

	return temp, true
}

func (b *BST[K, V]) Print() {
	if b.root == nil {
		fmt.Println("The BST is empty")
	}

	printhelp(b.root, 0)
}

func getmin[K comparable, V any](root *BSTNode[K, V]) *BSTNode[K, V] {
	if root.left == nil {
		return root
	}

	return getmin[K, V](root.left)
}

func deletemin[K comparable, V any](root *BSTNode[K, V]) *BSTNode[K, V] {
	if root.left == nil {
		return root.right
	}

	root.left = deletemin[K, V](root.left)

	return root
}

func findhelp[K comparable, V any](root *BSTNode[K, V], key K) (V, bool) {
	if root == nil {
		var zero V
		return zero, false
	}

	if key < root.key {
		return findhelp[K, V](root.left, key)
	}

	if key > root.key {
		return findhelp[K, V](root.right, key)
	}

	return root.val, true
}

func insertHelp[K comparable, V any](root *BSTNode[K, V], key K, value V) *BSTNode[K, V] {
	if root == nil {
		return &BSTNode[K, V]{key: key, val: value}
	}

	if k < root.key {
		root.left = insertHelp[K, V](root.left, key, value)
	}

	if k > root.key {
		root.right = insertHelp[K, V](root.right, key, value)
	}

	return root
}

func removehelp[K comparable, V any](root *BSTNode[K, V], key K) *BSTNode[K, V] {
	if root == nil {
		return nil
	}

	switch {
	case key < root.key:
		root.left = removehelp[K, V](root.left, key)
	case key > root.key:
		root.right = removehelp[K, V](root.right, key)
	default:
		if root.left == nil {
			root = root.right
		} else if root.right == nil {
			root = root.left
		} else {
			temp := getmin[K, V](root.right)
			root.val = temp.val
			root.key = temp.key
			root.right = deletemin[K, V](root.right)
		}
	}

	return root
}

func clearHelp[K comparable, V any](root *BSTNode[K, V]) {
	//
}

func printhelp[K comparable, V any](root *BSTNode[K, V], level int) {
	if root == nil {
		return
	}

	printhelp[K, V](root.left, level+1)

	for i := 0; i < level; i++ {
		fmt.Print(" ")
	}

	fmt.Println(root.key)

	printhelp[K, V](root.right, level+1)
}
