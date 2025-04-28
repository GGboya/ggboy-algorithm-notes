package kmp

type KMP struct {
}

// 获取pattern的next数组
func (k *KMP) getNext(pattern string) []int {
	n := len(pattern)
	next := make([]int, n)
	for i, j := 1, 0; i < n; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = next[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		next[i] = j
	}
	return next
}

// 在s中查找pattern, 返回pattern在s中的起始位置, 不存在返回-1
func (k *KMP) search(s string, pattern string) int {
	n := len(s)
	next := k.getNext(pattern)
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && s[i] != pattern[j] {
			j = next[j-1]
		}
		if s[i] == pattern[j] {
			j++
		}
		if j == len(pattern) {
			return i - j + 1
		}
	}
	return -1
}

func NewKMP() *KMP {
	return &KMP{}
}

// 例题[459] 重复的子字符串
// 给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。
func repeatedSubstringPattern(s string) bool {
	k := NewKMP()
	doubleS := s + s
	doubleS = doubleS[1 : len(doubleS)-1]
	/*
		充分性：
			如果s可以由一个子串重复多次构成，那么s+s去掉首尾字符后，s一定在s+s中出现过
		必要性：
			如果s+s去掉首尾字符后，s没有出现过，那么s一定不能由一个子串重复多次构成

		充分性很好理解，不多解释。
		去除首尾字符之后，新的字符 T = (s+s)[1:len(s)-1]
		假设 s = T[i:i+n], 那么 s 在 (s+s)中的位置就为 (s+s)[i+1:i+n+1]
		容易看出，s 一定同时处于两个 s 中，s = s[i+1:] + s[:i+1]
		s 可以表示为它自己的一个循环移位 = s 由某个子串重复构成
		接下来证明上述结论：
		## 严格的数学证明
		假设字符串S的长度为n，且S能表示为自身的一个非平凡循环移位（即移位量k满足0<k<n），那么：

		S = S[k:n] + S[0:k]

		我们需要证明S由某个子串重复构成。

		### 引理：如果S = S[k:n] + S[0:k]，则S由长度为gcd(n,k)的子串重复构成

		证明：
		1. 定义d = gcd(n,k)，即n和k的最大公约数
		2. 我们可以将n和k表示为：n = d·m，k = d·p，其中m和p是互质的整数

		3. 由于S = S[k:n] + S[0:k]，我们可以迭代应用这个等式：
		   - S = S[k:n] + S[0:k]
		   - = S[2k:n] + S[k:2k] + S[0:k]（应用等式到S[k:n]）
		   - = S[3k:n] + S[2k:3k] + S[k:2k] + S[0:k]
		   - ...以此类推

		4. 通过不断应用这个等式，我们最终会得到：
		   S = S[n-k:n] + S[n-2k:n-k] + ... + S[k:2k] + S[0:k]

		5. 注意到这把S分成了n/k个长度为k的块（如果n不能被k整除，最后一个块可能更短）

		6. 关键点：当我们应用等式m次（其中m是n/d与k/d的最小公倍数除以k/d）时，我们会回到起点。这是因为(m·k) mod n = 0。

		7. 这意味着S被分割成了若干个完全相同的子串，每个子串长度为d = gcd(n,k)

		8. 因此，S由一个长度为gcd(n,k)的子串重复n/gcd(n,k)次构成
	*/
	return k.search(doubleS, s) != -1
}
