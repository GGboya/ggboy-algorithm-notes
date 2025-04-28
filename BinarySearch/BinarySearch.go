package binarysearch

// 以下的二分，建立在数组升序的基础上

// 寻找数组中第一个大于等于 target 的下标，如果不存在返回-1
func findFirstGreaterOrEqual(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	l, r := 0, n-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	// 检查找到的元素是否满足条件
	if nums[l] >= target {
		return l
	}
	return -1 // 找不到大于等于target的元素
}

// 寻找数组中第一个大于 target 的下标，如果不存在返回-1
func findFirstGreater(nums []int, target int) int {
	idx := findFirstGreaterOrEqual(nums, target+1)
	if idx == -1 {
		return -1
	}
	return idx
}

// 寻找数组中最后一个小于等于 target 的下标，如果不存在返回-1
func findLastLessOrEqual(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	l, r := 0, n-1
	for l < r {
		mid := (l + r + 1) >> 1
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}

	// 检查找到的元素是否满足条件
	if nums[l] <= target {
		return l
	}
	return -1 // 找不到小于等于target的元素
}

// 寻找数组中最后一个小于 target 的下标，如果不存在返回-1
func findLastLess(nums []int, target int) int {
	idx := findLastLessOrEqual(nums, target-1)
	if idx == -1 {
		return -1
	}
	return idx
}
