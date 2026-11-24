package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// LevelOrder returns the tree's values level by level, left to right.
// Snapshotting the queue length at each round is what separates levels.
func LevelOrder(root *TreeNode) [][]int {
	out := [][]int{}
	if root == nil {
		return out
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue) // nodes belonging to the current level
		level := make([]int, 0, size)
		for i := 0; i < size; i++ {
			n := queue[0]
			queue = queue[1:]
			level = append(level, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		out = append(out, level)
	}
	return out
}

// SpiralOrder is level order with every second level reversed (the "spiral"
// or zig-zag form the sheet pairs with it).
func SpiralOrder(root *TreeNode) [][]int {
	out := LevelOrder(root)
	for i := 1; i < len(out); i += 2 {
		row := out[i]
		for l, r := 0, len(row)-1; l < r; l, r = l+1, r-1 {
			row[l], row[r] = row[r], row[l]
		}
	}
	return out
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(3, tn(9, nil, nil), tn(20, tn(15, nil, nil), tn(7, nil, nil)))
	fmt.Println(LevelOrder(root))
	fmt.Println(SpiralOrder(root))
}
