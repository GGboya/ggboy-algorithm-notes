package main

import (
	"reflect"
	"testing"
)

// 辅助函数：比较两个int切片是否相等（处理nil和空切片的情况）
func equalIntSlices(a, b []int) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	return reflect.DeepEqual(a, b)
}

// 辅助函数：比较两个二维int切片是否相等
func equalInt2DSlices(a, b [][]int) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	return reflect.DeepEqual(a, b)
}

// 测试用例结构
type testCase struct {
	name     string
	tree     *TreeNode
	expected []int
}

// 层序遍历测试用例结构
type levelOrderTestCase struct {
	name     string
	tree     *TreeNode
	expected [][]int
}

// 创建测试用的二叉树
func createTestTrees() map[string]*TreeNode {
	trees := make(map[string]*TreeNode)

	// 空树
	trees["empty"] = nil

	// 单节点树
	trees["single"] = &TreeNode{Val: 1}

	// 标准测试树
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	trees["standard"] = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}

	// 左偏树
	//   1
	//  /
	// 2
	///
	//3
	trees["left_skewed"] = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}

	// 右偏树
	// 1
	//  \
	//   2
	//    \
	//     3
	trees["right_skewed"] = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 3},
		},
	}

	// 完全二叉树
	//       1
	//      / \
	//     2   3
	//    / \ / \
	//   4  5 6  7
	trees["complete"] = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}

	return trees
}

// 前序遍历测试
func TestPreorderTraversal(t *testing.T) {
	trees := createTestTrees()

	testCases := []testCase{
		{"empty tree", trees["empty"], []int{}},
		{"single node", trees["single"], []int{1}},
		{"standard tree", trees["standard"], []int{1, 2, 4, 5, 3}},
		{"left skewed", trees["left_skewed"], []int{1, 2, 3}},
		{"right skewed", trees["right_skewed"], []int{1, 2, 3}},
		{"complete tree", trees["complete"], []int{1, 2, 4, 5, 3, 6, 7}},
	}

	for _, tc := range testCases {
		t.Run("recursive_"+tc.name, func(t *testing.T) {
			result := preorderTraversal(tc.tree)
			if !equalIntSlices(result, tc.expected) {
				t.Errorf("preorderTraversal(%s) = %v, want %v", tc.name, result, tc.expected)
			}
		})

		t.Run("iterative_"+tc.name, func(t *testing.T) {
			result := preorderTraversalIterative(tc.tree)
			if !equalIntSlices(result, tc.expected) {
				t.Errorf("preorderTraversalIterative(%s) = %v, want %v", tc.name, result, tc.expected)
			}
		})
	}
}

// 中序遍历测试
func TestInorderTraversal(t *testing.T) {
	trees := createTestTrees()

	testCases := []testCase{
		{"empty tree", trees["empty"], []int{}},
		{"single node", trees["single"], []int{1}},
		{"standard tree", trees["standard"], []int{4, 2, 5, 1, 3}},
		{"left skewed", trees["left_skewed"], []int{3, 2, 1}},
		{"right skewed", trees["right_skewed"], []int{1, 2, 3}},
		{"complete tree", trees["complete"], []int{4, 2, 5, 1, 6, 3, 7}},
	}

	for _, tc := range testCases {
		t.Run("recursive_"+tc.name, func(t *testing.T) {
			result := inorderTraversal(tc.tree)
			if !equalIntSlices(result, tc.expected) {
				t.Errorf("inorderTraversal(%s) = %v, want %v", tc.name, result, tc.expected)
			}
		})

		t.Run("iterative_"+tc.name, func(t *testing.T) {
			result := inorderTraversalIterative(tc.tree)
			if !equalIntSlices(result, tc.expected) {
				t.Errorf("inorderTraversalIterative(%s) = %v, want %v", tc.name, result, tc.expected)
			}
		})
	}
}

// 后序遍历测试
func TestPostorderTraversal(t *testing.T) {
	trees := createTestTrees()

	testCases := []testCase{
		{"empty tree", trees["empty"], []int{}},
		{"single node", trees["single"], []int{1}},
		{"standard tree", trees["standard"], []int{4, 5, 2, 3, 1}},
		{"left skewed", trees["left_skewed"], []int{3, 2, 1}},
		{"right skewed", trees["right_skewed"], []int{3, 2, 1}},
		{"complete tree", trees["complete"], []int{4, 5, 2, 6, 7, 3, 1}},
	}

	for _, tc := range testCases {
		t.Run("recursive_"+tc.name, func(t *testing.T) {
			result := postorderTraversal(tc.tree)
			if !equalIntSlices(result, tc.expected) {
				t.Errorf("postorderTraversal(%s) = %v, want %v", tc.name, result, tc.expected)
			}
		})

		t.Run("iterative_"+tc.name, func(t *testing.T) {
			result := postorderTraversalIterative(tc.tree)
			if !equalIntSlices(result, tc.expected) {
				t.Errorf("postorderTraversalIterative(%s) = %v, want %v", tc.name, result, tc.expected)
			}
		})
	}
}

// 层序遍历测试
func TestLevelOrderTraversal(t *testing.T) {
	trees := createTestTrees()

	testCases := []levelOrderTestCase{
		{"empty tree", trees["empty"], [][]int{}},
		{"single node", trees["single"], [][]int{{1}}},
		{"standard tree", trees["standard"], [][]int{{1}, {2, 3}, {4, 5}}},
		{"left skewed", trees["left_skewed"], [][]int{{1}, {2}, {3}}},
		{"right skewed", trees["right_skewed"], [][]int{{1}, {2}, {3}}},
		{"complete tree", trees["complete"], [][]int{{1}, {2, 3}, {4, 5, 6, 7}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := levelOrderTraversal(tc.tree)
			if !equalInt2DSlices(result, tc.expected) {
				t.Errorf("levelOrderTraversal(%s) = %v, want %v", tc.name, result, tc.expected)
			}
		})
	}
}

// 基准测试 - 测试性能
func BenchmarkPreorderTraversal(b *testing.B) {
	tree := createTestTrees()["complete"]

	b.Run("recursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			preorderTraversal(tree)
		}
	})

	b.Run("iterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			preorderTraversalIterative(tree)
		}
	})
}

func BenchmarkInorderTraversal(b *testing.B) {
	tree := createTestTrees()["complete"]

	b.Run("recursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			inorderTraversal(tree)
		}
	})

	b.Run("iterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			inorderTraversalIterative(tree)
		}
	})
}

func BenchmarkPostorderTraversal(b *testing.B) {
	tree := createTestTrees()["complete"]

	b.Run("recursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			postorderTraversal(tree)
		}
	})

	b.Run("iterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			postorderTraversalIterative(tree)
		}
	})
}

func BenchmarkLevelOrderTraversal(b *testing.B) {
	tree := createTestTrees()["complete"]

	for i := 0; i < b.N; i++ {
		levelOrderTraversal(tree)
	}
}

// 测试递归和迭代实现的一致性
func TestConsistencyBetweenRecursiveAndIterative(t *testing.T) {
	trees := createTestTrees()

	for name, tree := range trees {
		t.Run(name, func(t *testing.T) {
			// 前序遍历一致性
			preorderRec := preorderTraversal(tree)
			preorderIter := preorderTraversalIterative(tree)
			if !equalIntSlices(preorderRec, preorderIter) {
				t.Errorf("Preorder inconsistency for %s: recursive=%v, iterative=%v",
					name, preorderRec, preorderIter)
			}

			// 中序遍历一致性
			inorderRec := inorderTraversal(tree)
			inorderIter := inorderTraversalIterative(tree)
			if !equalIntSlices(inorderRec, inorderIter) {
				t.Errorf("Inorder inconsistency for %s: recursive=%v, iterative=%v",
					name, inorderRec, inorderIter)
			}

			// 后序遍历一致性
			postorderRec := postorderTraversal(tree)
			postorderIter := postorderTraversalIterative(tree)
			if !equalIntSlices(postorderRec, postorderIter) {
				t.Errorf("Postorder inconsistency for %s: recursive=%v, iterative=%v",
					name, postorderRec, postorderIter)
			}
		})
	}
}
