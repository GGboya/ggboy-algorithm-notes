package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	in  *bufio.Reader
	out *bufio.Writer
)

const (
	INF int = 1e9
)

/*
巨型多米诺
你要选择一些多米诺，满足下面的情形
1、最左边的多米洛编号是 1
2、最右边的多米诺编号是 n
3、推倒多米诺 1 的时候，n 也会倒下去
问是否存在这样的排列，如果有，最少的多米诺个数是多少个。
*/
func solve() {
	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	// 找到 l， 每次二分找到 <= 2l 的 最大元素
	start := arr[0]
	target := arr[n-1]

	if start*2 >= target {
		fmt.Fprintln(out, 2)
		return
	}

	sort.Ints(arr)
	ans := 2
	flag := 0

	for {
		if start*2 >= target {
			flag = 1
			break
		}
		l, r := 0, n-1
		for l < r {
			mid := (l + r + 1) >> 1
			if arr[mid] <= 2*start {
				l = mid
			} else {
				r = mid - 1
			}
		}
		if arr[l] <= 2*start && arr[l] != start {
			// l 这个多米诺，可以倒下去
			start = arr[l]
			ans += 1
		} else {
			break
		}
	}
	if flag == 0 {
		fmt.Fprintln(out, -1)
	} else {
		fmt.Fprintln(out, ans)
	}
}

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		solve()
	}
}
