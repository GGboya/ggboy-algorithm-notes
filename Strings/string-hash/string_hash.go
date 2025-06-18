package main

import "fmt"

// StringHash 字符串哈希结构
type StringHash struct {
	s    string  // 原字符串
	hash []int64 // 前缀哈希值
	pow  []int64 // base的幂次
	base int64   // 哈希基数
	mod  int64   // 模数
}

// NewStringHash 创建字符串哈希对象
func NewStringHash(s string) *StringHash {
	const base = 131    // 常用的哈希基数
	const mod = 1e9 + 7 // 大质数作为模数

	n := len(s)
	hash := make([]int64, n+1)
	pow := make([]int64, n+1)

	pow[0] = 1
	for i := 1; i <= n; i++ {
		// 计算前缀哈希值
		hash[i] = (hash[i-1]*base + int64(s[i-1])) % mod
		// 计算base的幂次
		pow[i] = (pow[i-1] * base) % mod
	}

	return &StringHash{
		s:    s,
		hash: hash,
		pow:  pow,
		base: base,
		mod:  mod,
	}
}

// GetHash 获取子串s[l:r+1]的哈希值 (左闭右闭区间)
func (sh *StringHash) GetHash(l, r int) int64 {
	// hash[r+1] - hash[l] * pow[r-l+1]
	result := (sh.hash[r+1] - sh.hash[l]*sh.pow[r-l+1]%sh.mod + sh.mod) % sh.mod
	return result
}

// GetStringHash 获取字符串的哈希值
func (sh *StringHash) GetStringHash(s string) int64 {
	hash := int64(0)
	for i := 0; i < len(s); i++ {
		hash = (hash*sh.base + int64(s[i])) % sh.mod
	}
	return hash
}

// CountSubstring 统计字符串s中有多少个子串等于p
func CountSubstring(s, p string) int {
	if len(p) == 0 || len(p) > len(s) {
		return 0
	}

	// 创建字符串哈希对象
	sh := NewStringHash(s)

	// 计算模式串p的哈希值
	pHash := sh.GetStringHash(p)

	count := 0
	pLen := len(p)

	// 遍历所有长度为len(p)的子串
	for i := 0; i <= len(s)-pLen; i++ {
		// 获取子串s[i:i+pLen]的哈希值
		subHash := sh.GetHash(i, i+pLen-1)
		if subHash == pHash {
			count++
		}
	}

	return count
}

// CountSubstringOptimized 优化版本：使用滚动哈希
func CountSubstringOptimized(s, p string) int {
	if len(p) == 0 || len(p) > len(s) {
		return 0
	}

	const base = 131
	const mod = 1e9 + 7

	pLen := len(p)
	sLen := len(s)

	// 计算模式串p的哈希值
	pHash := int64(0)
	for i := 0; i < pLen; i++ {
		pHash = (pHash*base + int64(p[i])) % mod
	}

	// 计算base^pLen
	basePow := int64(1)
	for i := 0; i < pLen; i++ {
		basePow = (basePow * base) % mod
	}

	// 滚动哈希
	sHash := int64(0)
	count := 0

	for i := 0; i < sLen; i++ {
		// 添加新字符
		sHash = (sHash*base + int64(s[i])) % mod

		// 如果窗口大小超过pLen，移除最左边的字符
		if i >= pLen {
			sHash = (sHash - int64(s[i-pLen])*basePow%mod + mod) % mod
		}

		// 如果窗口大小等于pLen，检查是否匹配
		if i >= pLen-1 && sHash == pHash {
			count++
		}
	}

	return count
}

// CountSubstringWithDoubleHash 双哈希版本，减少碰撞概率
func CountSubstringWithDoubleHash(s, p string) int {
	if len(p) == 0 || len(p) > len(s) {
		return 0
	}

	const base1, mod1 = 131, 1e9 + 7
	const base2, mod2 = 137, 1e9 + 9

	pLen := len(p)
	sLen := len(s)

	// 计算模式串p的双哈希值
	pHash1, pHash2 := int64(0), int64(0)
	for i := 0; i < pLen; i++ {
		pHash1 = (pHash1*base1 + int64(p[i])) % mod1
		pHash2 = (pHash2*base2 + int64(p[i])) % mod2
	}

	// 计算base的幂次
	basePow1, basePow2 := int64(1), int64(1)
	for i := 0; i < pLen; i++ {
		basePow1 = (basePow1 * base1) % mod1
		basePow2 = (basePow2 * base2) % mod2
	}

	// 滚动双哈希
	sHash1, sHash2 := int64(0), int64(0)
	count := 0

	for i := 0; i < sLen; i++ {
		// 添加新字符
		sHash1 = (sHash1*base1 + int64(s[i])) % mod1
		sHash2 = (sHash2*base2 + int64(s[i])) % mod2

		// 移除最左边的字符
		if i >= pLen {
			sHash1 = (sHash1 - int64(s[i-pLen])*basePow1%mod1 + mod1) % mod1
			sHash2 = (sHash2 - int64(s[i-pLen])*basePow2%mod2 + mod2) % mod2
		}

		// 检查双哈希是否都匹配
		if i >= pLen-1 && sHash1 == pHash1 && sHash2 == pHash2 {
			count++
		}
	}

	return count
}

// FindAllSubstringPositions 找到所有匹配子串的位置
func FindAllSubstringPositions(s, p string) []int {
	if len(p) == 0 || len(p) > len(s) {
		return []int{}
	}

	sh := NewStringHash(s)
	pHash := sh.GetStringHash(p)

	var positions []int
	pLen := len(p)

	for i := 0; i <= len(s)-pLen; i++ {
		subHash := sh.GetHash(i, i+pLen-1)
		if subHash == pHash {
			positions = append(positions, i)
		}
	}

	return positions
}

// 演示函数
func demonstrateStringHash() {
	fmt.Println("=== 字符串哈希算法演示 ===")
	fmt.Println()

	// 测试用例
	testCases := []struct {
		s string
		p string
	}{
		{"abcabcabc", "abc"},
		{"aaaa", "aa"},
		{"hello world hello", "hello"},
		{"abababab", "abab"},
		{"", "a"},
		{"a", ""},
		{"same", "same"},
		{"different", "xyz"},
	}

	for i, tc := range testCases {
		fmt.Printf("测试用例 %d:\n", i+1)
		fmt.Printf("字符串 s: \"%s\"\n", tc.s)
		fmt.Printf("模式串 p: \"%s\"\n", tc.p)

		// 测试不同方法
		count1 := CountSubstring(tc.s, tc.p)
		count2 := CountSubstringOptimized(tc.s, tc.p)
		count3 := CountSubstringWithDoubleHash(tc.s, tc.p)

		fmt.Printf("基础方法结果: %d\n", count1)
		fmt.Printf("滚动哈希结果: %d\n", count2)
		fmt.Printf("双哈希结果: %d\n", count3)

		// 找到所有位置
		positions := FindAllSubstringPositions(tc.s, tc.p)
		fmt.Printf("匹配位置: %v\n", positions)

		// 验证一致性
		if count1 == count2 && count2 == count3 {
			fmt.Printf("✅ 所有方法结果一致\n")
		} else {
			fmt.Printf("❌ 方法结果不一致\n")
		}

		fmt.Println()
	}

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	longS := ""
	for i := 0; i < 1000; i++ {
		longS += "abcdefgh"
	}
	pattern := "abcd"

	fmt.Printf("长字符串长度: %d\n", len(longS))
	fmt.Printf("模式串: \"%s\"\n", pattern)

	result := CountSubstringOptimized(longS, pattern)
	fmt.Printf("匹配数量: %d\n", result)
}

func main() {
	demonstrateStringHash()
}
