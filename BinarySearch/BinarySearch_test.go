package binarysearch

import (
	"testing"
)

func TestFindFirstGreaterOrEqual(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Empty array",
			nums:     []int{},
			target:   5,
			expected: -1, // 空数组应该返回-1
		},
		{
			name:     "Target smaller than all elements",
			nums:     []int{5, 7, 9, 11},
			target:   2,
			expected: 0, // First element >= 2 is at index 0 (value 5)
		},
		{
			name:     "Target equal to first element",
			nums:     []int{5, 7, 9, 11},
			target:   5,
			expected: 0, // First element >= 5 is at index 0 (value 5)
		},
		{
			name:     "Target between elements",
			nums:     []int{5, 7, 9, 11},
			target:   8,
			expected: 2, // First element >= 8 is at index 2 (value 9)
		},
		{
			name:     "Target equal to middle element",
			nums:     []int{5, 7, 9, 11},
			target:   7,
			expected: 1, // First element >= 7 is at index 1 (value 7)
		},
		{
			name:     "Target equal to last element",
			nums:     []int{5, 7, 9, 11},
			target:   11,
			expected: 3, // First element >= 11 is at index 3 (value 11)
		},
		{
			name:     "Target larger than all elements",
			nums:     []int{5, 7, 9, 11},
			target:   15,
			expected: -1, // 没有大于等于15的元素，应返回-1
		},
		{
			name:     "Array with duplicates, target in array",
			nums:     []int{1, 3, 3, 5, 7, 7, 7, 9},
			target:   7,
			expected: 4, // First occurrence of element >= 7 (value 7)
		},
		{
			name:     "Array with duplicates, target between elements",
			nums:     []int{1, 3, 3, 5, 7, 7, 7, 9},
			target:   6,
			expected: 4, // First element >= 6 is at index 4 (value 7)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findFirstGreaterOrEqual(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("findFirstGreaterOrEqual(%v, %d) = %d, want %d",
					tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}

func TestFindFirstGreater(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Empty array",
			nums:     []int{},
			target:   5,
			expected: -1, // 空数组应该返回-1
		},
		{
			name:     "Target smaller than all elements",
			nums:     []int{5, 7, 9, 11},
			target:   2,
			expected: 0, // First element > 2 is at index 0 (value 5)
		},
		{
			name:     "Target equal to first element",
			nums:     []int{5, 7, 9, 11},
			target:   5,
			expected: 1, // First element > 5 is at index 1 (value 7)
		},
		{
			name:     "Target between elements",
			nums:     []int{5, 7, 9, 11},
			target:   8,
			expected: 2, // First element > 8 is at index 2 (value 9)
		},
		{
			name:     "Target equal to middle element",
			nums:     []int{5, 7, 9, 11},
			target:   7,
			expected: 2, // First element > 7 is at index 2 (value 9)
		},
		{
			name:     "Target equal to last element",
			nums:     []int{5, 7, 9, 11},
			target:   11,
			expected: -1, // 没有大于11的元素，应返回-1
		},
		{
			name:     "Target larger than all elements",
			nums:     []int{5, 7, 9, 11},
			target:   15,
			expected: -1, // 没有大于15的元素，应返回-1
		},
		{
			name:     "Array with duplicates, target in array",
			nums:     []int{1, 3, 3, 5, 7, 7, 7, 9},
			target:   7,
			expected: 7, // First element > 7 is at index 7 (value 9)
		},
		{
			name:     "Array with duplicates, target between elements",
			nums:     []int{1, 3, 3, 5, 7, 7, 7, 9},
			target:   6,
			expected: 4, // First element > 6 is at index 4 (value 7)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findFirstGreater(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("findFirstGreater(%v, %d) = %d, want %d",
					tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}

func TestFindLastLessOrEqual(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Empty array",
			nums:     []int{},
			target:   5,
			expected: -1, // 空数组应该返回-1
		},
		{
			name:     "Target smaller than all elements",
			nums:     []int{5, 7, 9, 11},
			target:   2,
			expected: -1, // 没有小于等于2的元素，应返回-1
		},
		{
			name:     "Target equal to first element",
			nums:     []int{5, 7, 9, 11},
			target:   5,
			expected: 0, // Last element <= 5 is at index 0 (value 5)
		},
		{
			name:     "Target between elements",
			nums:     []int{5, 7, 9, 11},
			target:   8,
			expected: 1, // Last element <= 8 is at index 1 (value 7)
		},
		{
			name:     "Target equal to middle element",
			nums:     []int{5, 7, 9, 11},
			target:   7,
			expected: 1, // Last element <= 7 is at index 1 (value 7)
		},
		{
			name:     "Target equal to last element",
			nums:     []int{5, 7, 9, 11},
			target:   11,
			expected: 3, // Last element <= 11 is at index 3 (value 11)
		},
		{
			name:     "Target larger than all elements",
			nums:     []int{5, 7, 9, 11},
			target:   15,
			expected: 3, // Last element <= 15 is at index 3 (value 11)
		},
		{
			name:     "Array with duplicates, target in array",
			nums:     []int{1, 3, 3, 5, 7, 7, 7, 9},
			target:   7,
			expected: 6, // Last occurrence of element <= 7 is at index 6 (value 7)
		},
		{
			name:     "Array with duplicates, target between elements",
			nums:     []int{1, 3, 3, 5, 7, 7, 7, 9},
			target:   6,
			expected: 3, // Last element <= 6 is at index 3 (value 5)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findLastLessOrEqual(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("findLastLessOrEqual(%v, %d) = %d, want %d",
					tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}

func TestFindLastLess(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Empty array",
			nums:     []int{},
			target:   5,
			expected: -1, // 空数组应该返回-1
		},
		{
			name:     "Target smaller than all elements",
			nums:     []int{5, 7, 9, 11},
			target:   2,
			expected: -1, // 没有小于2的元素，应返回-1
		},
		{
			name:     "Target equal to first element",
			nums:     []int{5, 7, 9, 11},
			target:   5,
			expected: -1, // 没有小于5的元素，应返回-1
		},
		{
			name:     "Target between elements",
			nums:     []int{5, 7, 9, 11},
			target:   8,
			expected: 1, // Last element < 8 is at index 1 (value 7)
		},
		{
			name:     "Target equal to middle element",
			nums:     []int{5, 7, 9, 11},
			target:   7,
			expected: 0, // Last element < 7 is at index 0 (value 5)
		},
		{
			name:     "Target equal to last element",
			nums:     []int{5, 7, 9, 11},
			target:   11,
			expected: 2, // Last element < 11 is at index 2 (value 9)
		},
		{
			name:     "Target larger than all elements",
			nums:     []int{5, 7, 9, 11},
			target:   15,
			expected: 3, // Last element < 15 is at index 3 (value 11)
		},
		{
			name:     "Array with duplicates, target in array",
			nums:     []int{1, 3, 3, 5, 7, 7, 7, 9},
			target:   7,
			expected: 3, // Last element < 7 is at index 3 (value 5)
		},
		{
			name:     "Array with duplicates, target between elements",
			nums:     []int{1, 3, 3, 5, 7, 7, 7, 9},
			target:   6,
			expected: 3, // Last element < 6 is at index 3 (value 5)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findLastLess(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("findLastLess(%v, %d) = %d, want %d",
					tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}

// TestEdgeCases tests edge cases for all functions with a single element array
func TestEdgeCases(t *testing.T) {
	// Test with a single element array
	nums := []int{42}

	// findFirstGreaterOrEqual tests
	t.Run("findFirstGreaterOrEqual: Target equal to element", func(t *testing.T) {
		if got := findFirstGreaterOrEqual(nums, 42); got != 0 {
			t.Errorf("findFirstGreaterOrEqual([42], 42) = %d, want 0", got)
		}
	})

	t.Run("findFirstGreaterOrEqual: Target less than element", func(t *testing.T) {
		if got := findFirstGreaterOrEqual(nums, 30); got != 0 {
			t.Errorf("findFirstGreaterOrEqual([42], 30) = %d, want 0", got)
		}
	})

	t.Run("findFirstGreaterOrEqual: Target greater than element", func(t *testing.T) {
		if got := findFirstGreaterOrEqual(nums, 50); got != -1 { // 更新期望值为-1
			t.Errorf("findFirstGreaterOrEqual([42], 50) = %d, want -1", got)
		}
	})

	// findFirstGreater tests
	t.Run("findFirstGreater: Target equal to element", func(t *testing.T) {
		if got := findFirstGreater(nums, 42); got != -1 { // 更新期望值为-1
			t.Errorf("findFirstGreater([42], 42) = %d, want -1", got)
		}
	})

	t.Run("findFirstGreater: Target less than element", func(t *testing.T) {
		if got := findFirstGreater(nums, 30); got != 0 {
			t.Errorf("findFirstGreater([42], 30) = %d, want 0", got)
		}
	})

	t.Run("findFirstGreater: Target greater than element", func(t *testing.T) {
		if got := findFirstGreater(nums, 50); got != -1 { // 更新期望值为-1
			t.Errorf("findFirstGreater([42], 50) = %d, want -1", got)
		}
	})

	// findLastLessOrEqual tests
	t.Run("findLastLessOrEqual: Target equal to element", func(t *testing.T) {
		if got := findLastLessOrEqual(nums, 42); got != 0 { // 更新期望值为0
			t.Errorf("findLastLessOrEqual([42], 42) = %d, want 0", got)
		}
	})

	t.Run("findLastLessOrEqual: Target less than element", func(t *testing.T) {
		if got := findLastLessOrEqual(nums, 30); got != -1 {
			t.Errorf("findLastLessOrEqual([42], 30) = %d, want -1", got)
		}
	})

	t.Run("findLastLessOrEqual: Target greater than element", func(t *testing.T) {
		if got := findLastLessOrEqual(nums, 50); got != 0 { // 更新期望值为0
			t.Errorf("findLastLessOrEqual([42], 50) = %d, want 0", got)
		}
	})

	// findLastLess tests
	t.Run("findLastLess: Target equal to element", func(t *testing.T) {
		if got := findLastLess(nums, 42); got != -1 {
			t.Errorf("findLastLess([42], 42) = %d, want -1", got)
		}
	})

	t.Run("findLastLess: Target less than element", func(t *testing.T) {
		if got := findLastLess(nums, 30); got != -1 {
			t.Errorf("findLastLess([42], 30) = %d, want -1", got)
		}
	})

	t.Run("findLastLess: Target greater than element", func(t *testing.T) {
		if got := findLastLess(nums, 50); got != 0 { // 更新期望值为0
			t.Errorf("findLastLess([42], 50) = %d, want 0", got)
		}
	})
}
