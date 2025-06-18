package main

import (
	"strings"
	"testing"
)

// 测试基础字符串哈希功能
func TestStringHash(t *testing.T) {
	s := "abcdefgh"
	sh := NewStringHash(s)

	// 测试获取子串哈希值
	hash1 := sh.GetHash(0, 2) // "abc"
	hash2 := sh.GetHash(3, 5) // "def"
	hash3 := sh.GetHash(0, 2) // "abc" again

	// 相同子串应该有相同哈希值
	if hash1 != hash3 {
		t.Errorf("相同子串的哈希值应该相等: %d != %d", hash1, hash3)
	}

	// 不同子串应该有不同哈希值（大概率）
	if hash1 == hash2 {
		t.Logf("警告: 不同子串产生了相同哈希值（可能的碰撞）")
	}
}

// 测试计数子串功能
func TestCountSubstring(t *testing.T) {
	testCases := []struct {
		s        string
		p        string
		expected int
	}{
		{"abcabcabc", "abc", 3},
		{"aaaa", "aa", 3},
		{"hello world hello", "hello", 2},
		{"abababab", "abab", 3},
		{"", "a", 0},
		{"a", "", 0},
		{"same", "same", 1},
		{"different", "xyz", 0},
		{"abcdef", "abcdef", 1},
		{"aaaaaa", "aaa", 4},
	}

	for _, tc := range testCases {
		t.Run(tc.s+"_"+tc.p, func(t *testing.T) {
			// 测试所有三种方法
			result1 := CountSubstring(tc.s, tc.p)
			result2 := CountSubstringOptimized(tc.s, tc.p)
			result3 := CountSubstringWithDoubleHash(tc.s, tc.p)

			// 验证结果正确性
			if result1 != tc.expected {
				t.Errorf("CountSubstring(%q, %q) = %d, want %d", tc.s, tc.p, result1, tc.expected)
			}
			if result2 != tc.expected {
				t.Errorf("CountSubstringOptimized(%q, %q) = %d, want %d", tc.s, tc.p, result2, tc.expected)
			}
			if result3 != tc.expected {
				t.Errorf("CountSubstringWithDoubleHash(%q, %q) = %d, want %d", tc.s, tc.p, result3, tc.expected)
			}

			// 验证方法一致性
			if result1 != result2 || result2 != result3 {
				t.Errorf("方法结果不一致: %d, %d, %d", result1, result2, result3)
			}
		})
	}
}

// 测试查找所有位置功能
func TestFindAllSubstringPositions(t *testing.T) {
	testCases := []struct {
		s        string
		p        string
		expected []int
	}{
		{"abcabcabc", "abc", []int{0, 3, 6}},
		{"aaaa", "aa", []int{0, 1, 2}},
		{"hello world hello", "hello", []int{0, 12}},
		{"abababab", "abab", []int{0, 2, 4}},
		{"same", "same", []int{0}},
		{"different", "xyz", []int{}},
	}

	for _, tc := range testCases {
		t.Run(tc.s+"_"+tc.p, func(t *testing.T) {
			result := FindAllSubstringPositions(tc.s, tc.p)

			if len(result) != len(tc.expected) {
				t.Errorf("位置数量不匹配: got %v, want %v", result, tc.expected)
				return
			}

			for i, pos := range result {
				if pos != tc.expected[i] {
					t.Errorf("位置不匹配: got %v, want %v", result, tc.expected)
					break
				}
			}
		})
	}
}

// 与标准库方法对比测试
func TestCompareWithStandardLibrary(t *testing.T) {
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
		{"abcdefghijklmnopqrstuvwxyz", "def"},
		{"aaaaaaaaaa", "aaa"},
	}

	for _, tc := range testCases {
		t.Run(tc.s+"_"+tc.p, func(t *testing.T) {
			// 使用字符串哈希方法
			hashResult := CountSubstring(tc.s, tc.p)

			// 使用标准库方法计算期望结果
			expected := 0
			if tc.p != "" {
				start := 0
				for {
					index := strings.Index(tc.s[start:], tc.p)
					if index == -1 {
						break
					}
					expected++
					start += index + 1
				}
			}

			if hashResult != expected {
				t.Errorf("与标准库结果不一致: hash=%d, stdlib=%d", hashResult, expected)
			}
		})
	}
}

// 边界情况测试
func TestEdgeCases(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		p        string
		expected int
	}{
		{"空字符串和空模式", "", "", 0},
		{"空字符串和非空模式", "", "a", 0},
		{"非空字符串和空模式", "abc", "", 0},
		{"模式长度大于字符串", "ab", "abc", 0},
		{"单字符匹配", "a", "a", 1},
		{"单字符不匹配", "a", "b", 0},
		{"重复字符", "aaaa", "a", 4},
		{"完全匹配", "hello", "hello", 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := CountSubstring(tc.s, tc.p)
			if result != tc.expected {
				t.Errorf("%s: got %d, want %d", tc.name, result, tc.expected)
			}
		})
	}
}

// 性能基准测试
func BenchmarkCountSubstring(b *testing.B) {
	// 创建长字符串
	longString := strings.Repeat("abcdefgh", 1000)
	pattern := "abcd"

	b.Run("Basic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CountSubstring(longString, pattern)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CountSubstringOptimized(longString, pattern)
		}
	})

	b.Run("DoubleHash", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CountSubstringWithDoubleHash(longString, pattern)
		}
	})
}

// 测试哈希碰撞
func TestHashCollision(t *testing.T) {
	// 这个测试主要是为了观察是否有哈希碰撞
	// 在实际应用中，我们通常使用双哈希来减少碰撞概率

	s := "abcdefghijklmnopqrstuvwxyz"
	sh := NewStringHash(s)

	hashMap := make(map[int64]string)
	collisions := 0

	// 测试所有长度为3的子串
	for i := 0; i <= len(s)-3; i++ {
		substr := s[i : i+3]
		hash := sh.GetHash(i, i+2)

		if existing, exists := hashMap[hash]; exists {
			if existing != substr {
				collisions++
				t.Logf("哈希碰撞: %q 和 %q 有相同哈希值 %d", existing, substr, hash)
			}
		} else {
			hashMap[hash] = substr
		}
	}

	if collisions > 0 {
		t.Logf("发现 %d 个哈希碰撞", collisions)
	}
}
