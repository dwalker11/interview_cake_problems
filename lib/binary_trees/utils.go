package binarytrees

func Count[K comparable, V any](root *BSTNode[K, V]) int {
	if root == nil {
		return 0
	}

	return 1 + Count(root.left) + Count(root.right)
}

func CheckBST[K comparable, V any](root *BSTNode[K, V], low K, high K) bool {
	if root == nil {
		return true
	}

	if root.key < low || root.key > high {
		return false
	}

	if !CheckBST(root.left, low, root.key) {
		return false
	}

	return CheckBST(root.right, root.key, high)
}
