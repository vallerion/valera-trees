package bst

import "trees/utils"

type BST struct {
	root       *NodeTree
	comparator utils.Comparator
}

type NodeTree struct {
	key, value  interface{}
	left, right *NodeTree
}

func ConstructIntTree() *BST {
	return &BST{nil, &utils.ComparatorInt{}}
}

func (tree *BST) Search(key interface{}) (interface{}, bool) {
	node := tree.searchNode(key)
	return node, node != nil
}

func (tree *BST) searchNode(key interface{}) *NodeTree {
	node := tree.root

	for node != nil {
		comparison := tree.comparator.Compare(node.key, key)

		if comparison.IsEqual() {
			return node
		} else if comparison.IsLess() {
			node = node.right
		} else {
			node = node.left
		}
	}

	return nil
}

func (tree *BST) Insert(key, value interface{}) {
	if tree.root == nil {
		tree.root = &NodeTree{key, value, nil, nil}
		return
	}

	node := tree.root

	for node != nil {
		comparison := tree.comparator.Compare(node.key, key)

		if comparison.IsEqual() {
			if node.right == nil {
				node.right = &NodeTree{key, value, nil, nil}
			} else {
				// equals value should be inserted into right subtree
				temp := node.right
				node.right = &NodeTree{key, value, nil, temp}
			}
			return
		} else if comparison.IsLess() {
			if node.right == nil {
				node.right = &NodeTree{key, value, nil, nil}
				return
			} else {
				node = node.right
			}
		} else {
			if node.left == nil {
				node.left = &NodeTree{key, value, nil, nil}
				return
			} else {
				node = node.left
			}
		}
	}
}

// Delete
// Return true if node was successfully deleted
func (tree *BST) Delete(key interface{}) bool {
	node := tree.root
	var prev *NodeTree
	fromLeft := false

	for node != nil {
		comparison := tree.comparator.Compare(node.key, key)

		if comparison.IsEqual() {
			if node.left == nil && node.right == nil { // if node does not have any children make it null
				// if prev is null it means root is found node
				if prev == nil {
					tree.root = nil
				} else if fromLeft == true { // if we came from parent to left child
					prev.left = nil
				} else {
					prev.right = nil
				}
			} else if node.left == nil { // if node has only right (may be)
				if prev == nil {
					tree.root = node.right
				} else if fromLeft == true {
					prev.left = node.right
				} else {
					prev.right = node.right
				}
			} else if node.right == nil {
				if prev == nil {
					tree.root = node.left
				} else if fromLeft == true {
					prev.left = node.left
				} else {
					prev.right = node.left
				}
			} else {
				// if both children exists
				// we should take most left child of right child
				// it will replace deleted node

				parentMostLeft, mostLeft := node, node.right
				for mostLeft.left != nil {
					parentMostLeft = mostLeft
					mostLeft = mostLeft.left
				}

				if parentMostLeft == node {
					parentMostLeft.right = mostLeft.right
				} else {
					parentMostLeft.left = mostLeft.right
				}

				node.key = mostLeft.key
				node.value = mostLeft.value
			}

			return true

		} else if comparison.IsLess() {
			prev = node
			node = node.right
			fromLeft = false
		} else {
			prev = node
			node = node.left
			fromLeft = true
		}
	}

	return false
}

func (tree *BST) Update(key, newValue interface{}) bool {
	node := tree.searchNode(key)

	if node == nil {
		return false
	}

	node.value = newValue
	return true
}

func (tree *BST) CreateFromSortedArray(arr []int) {
	tree.root = arrayToTreeHelper(arr, 0, len(arr)-1)
}

func arrayToTreeHelper(arr []int, start, end int) *NodeTree {
	if start > end {
		return nil
	}

	middle := start + (end-start)/2

	return &NodeTree{
		arr[middle],
		arr[middle],
		arrayToTreeHelper(arr, start, middle-1),
		arrayToTreeHelper(arr, middle+1, end),
	}
}

func (tree *BST) ToArray() []int {
	arr := make([]int, 0)
	treeToArrayHelper(tree.root, &arr)
	return arr
}

func treeToArrayHelper(root *NodeTree, arr *[]int) {
	if root == nil {
		return
	}

	treeToArrayHelper(root.left, arr)
	*arr = append(*arr, root.value.(int))
	treeToArrayHelper(root.right, arr)
}
