package main

import "fmt"

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历 (根 -> 左 -> 右)
func preorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	// 递归实现
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val) // 访问根节点
		preorder(node.Left)               // 遍历左子树
		preorder(node.Right)              // 遍历右子树
	}

	preorder(root)
	return result
}

// 前序遍历 - 迭代实现
func preorderTraversalIterative(root *TreeNode) []int {
	type stackItem struct {
		node  *TreeNode
		state int // 0: 未访问, 1: 已访问
	}
	result := []int{}
	stack := []stackItem{{node: root, state: 0}}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.node == nil {
			continue
		}

		if current.state == 0 {
			// 按照 右->左->根 的顺序入栈(出栈时就是根->左->右)
			stack = append(stack, stackItem{node: current.node.Right, state: 0})
			stack = append(stack, stackItem{node: current.node.Left, state: 0})
			stack = append(stack, stackItem{node: current.node, state: 1})
		} else {
			result = append(result, current.node.Val)
		}
	}
	return result
}

// 中序遍历 (左 -> 根 -> 右)
func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	// 递归实现
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)                // 遍历左子树
		result = append(result, node.Val) // 访问根节点
		inorder(node.Right)               // 遍历右子树
	}

	inorder(root)
	return result
}

// 中序遍历 - 迭代实现
func inorderTraversalIterative(root *TreeNode) []int {
	type stackItem struct {
		node  *TreeNode
		state int // 0: 未访问, 1: 已访问
	}
	result := []int{}
	stack := []stackItem{{node: root, state: 0}}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.node == nil {
			continue
		}

		if current.state == 0 {
			// 左中右是我们期望的，所以入栈的时候是右中左
			stack = append(stack, stackItem{node: current.node.Right, state: 0})
			stack = append(stack, stackItem{node: current.node, state: 1})
			stack = append(stack, stackItem{node: current.node.Left, state: 0})
		} else {
			result = append(result, current.node.Val)
		}
	}
	return result
}

// 后序遍历 (左 -> 右 -> 根)
func postorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	// 递归实现
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)              // 遍历左子树
		postorder(node.Right)             // 遍历右子树
		result = append(result, node.Val) // 访问根节点
	}

	postorder(root)
	return result
}

// 后序遍历 - 迭代实现
func postorderTraversalIterative(root *TreeNode) []int {
	type stackItem struct {
		node  *TreeNode
		state int // 0: 未访问, 1: 已访问
	}
	result := []int{}
	stack := []stackItem{{node: root, state: 0}}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.node == nil {
			continue
		}

		if current.state == 0 {
			// 左右中是我们期望的，所以入栈的时候是中右左
			stack = append(stack, stackItem{node: current.node, state: 1})
			stack = append(stack, stackItem{node: current.node.Right, state: 0})
			stack = append(stack, stackItem{node: current.node.Left, state: 0})
		} else {
			result = append(result, current.node.Val)
		}
	}
	return result
}

// 层序遍历 (广度优先遍历)
func levelOrderTraversal(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		var currentLevel []int

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			currentLevel = append(currentLevel, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, currentLevel)
	}

	return result
}

// 创建示例二叉树 - 用于演示
func createSampleTree() *TreeNode {
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	return root
}

// 演示函数 - 可以单独调用来查看结果
func demonstrateTraversals() {
	// 创建示例二叉树
	root := createSampleTree()

	fmt.Println("二叉树结构:")
	fmt.Println("       1")
	fmt.Println("      / \\")
	fmt.Println("     2   3")
	fmt.Println("    / \\")
	fmt.Println("   4   5")
	fmt.Println()

	// 前序遍历
	fmt.Println("前序遍历 (根->左->右):")
	fmt.Println("递归实现:", preorderTraversal(root))
	fmt.Println("迭代实现:", preorderTraversalIterative(root))
	fmt.Println()

	// 中序遍历
	fmt.Println("中序遍历 (左->根->右):")
	fmt.Println("递归实现:", inorderTraversal(root))
	fmt.Println("迭代实现:", inorderTraversalIterative(root))
	fmt.Println()

	// 后序遍历
	fmt.Println("后序遍历 (左->右->根):")
	fmt.Println("递归实现:", postorderTraversal(root))
	fmt.Println("迭代实现:", postorderTraversalIterative(root))
	fmt.Println()

	// 层序遍历
	fmt.Println("层序遍历 (广度优先):")
	fmt.Println("结果:", levelOrderTraversal(root))
}
