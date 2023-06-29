package main

import "fmt"

type Employee struct {
	name   string
	salary int
}

type Node[T any] struct {
	value T
}

type TreeNode[T any] struct {
	Node[T]
	children []*TreeNode[T]
}

type BinaryTreeNode[T any] struct {
	Node[T]
	left  *BinaryTreeNode[T]
	right *BinaryTreeNode[T]
}

func (z *TreeNode[T]) preorderDFS() {
	if z == nil {
		return
	}

	fmt.Println(z.value)
	for _, child := range z.children {
		child.preorderDFS()
	}
}

func (z *TreeNode[T]) postorderDFS() {
	if z == nil {
		return
	}
	for _, child := range z.children {
		child.postorderDFS()
	}
	fmt.Println(z.value)
}

func (z *TreeNode[T]) iterativeDFS() {
	// this is preorder
	// what does dfs do, it visits the deepest nodes first then makes its way up

	stack := []*TreeNode[T]{z}
	// while curr node is not nil or theres still stuff in the stack
	for len(stack) > 0 {
		// load up the stack
		// we got to the lowest level, now we act
		currNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(currNode.value)
		for i := len(currNode.children) - 1; i >= 0; i-- {
			stack = append(stack, currNode.children[i])
		}
	}
}

func (z *TreeNode[T]) iterativePostorderDFS() {
	stack := []*TreeNode[T]{z}
	var result []T
	// while curr node is not nil or theres still stuff in the stack
	for len(stack) > 0 {
		// load up the stack
		// we got to the lowest level, now we act
		currNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, currNode.value)
		//for i := len(currNode.children) - 1; i >= 0; i-- {
		//	stack = append(stack, currNode.children[i])
		//}
		for _, child := range currNode.children {
			stack = append(stack, child)
		}
	}
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Println(result[i])
	}
}

func (z *TreeNode[T]) bfs() {
	if z == nil {
		return
	}

	q := []*TreeNode[T]{z}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		fmt.Println(node.value)
		for _, child := range node.children {
			q = append(q, child)
		}
	}
}

func (t *BinaryTreeNode[T]) dfs() {
	if t == nil {
		return
	}
	t.left.dfs()
	fmt.Println(t.value)
	t.right.dfs()
}

func (t *BinaryTreeNode[T]) bfs() {
	if t == nil {
		return
	}
	queue := []*BinaryTreeNode[T]{t}

	for len(queue) > 0 {
		treeNode := queue[0]
		queue = queue[1:]
		fmt.Println(treeNode.value)

		if treeNode.left != nil {
			queue = append(queue, treeNode.left)
		}
		if treeNode.right != nil {
			queue = append(queue, treeNode.right)
		}
	}
}

func (t *BinaryTreeNode[T]) iterativeDFS() {
	if t == nil {
		return
	}

	var stack []*BinaryTreeNode[T]

	currentTreeNode := t

	for currentTreeNode != nil || len(stack) > 0 {
		for currentTreeNode != nil {
			stack = append(stack, currentTreeNode)
			currentTreeNode = currentTreeNode.left
		}

		currentTreeNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(currentTreeNode.value)
		currentTreeNode = currentTreeNode.right
	}

}

/*
Can I make a tree object with an implemention of DFS, BFS, and Iterative BFS that takes a generic struct?
*/
