# 字符串哈希算法

字符串哈希是一种将字符串映射为数值的技术，常用于快速字符串匹配、子串查找等问题。

## 算法原理

### 多项式哈希
字符串哈希通常使用多项式哈希函数：
```
hash(s) = (s[0] * base^(n-1) + s[1] * base^(n-2) + ... + s[n-1] * base^0) mod p
```

其中：
- `base` 是哈希基数（常用131、233等）
- `p` 是大质数模数（常用1e9+7）
- `s[i]` 是字符的ASCII值

### 前缀哈希
为了快速计算任意子串的哈希值，我们预计算前缀哈希：
```
prefix[i] = hash(s[0:i])
```

任意子串 `s[l:r]` 的哈希值可以通过前缀哈希计算：
```
hash(s[l:r]) = (prefix[r] - prefix[l] * base^(r-l)) mod p

举例说明该公式：
假设 base = x
abcde = a * x^4 + b * x^3 + c * x^2 + d * x^1 + e * x^0
ab = a * x^1 + b * x^0
如果要计算 cde 的哈希值，则有：
cde = c * x^2 + d * x^1 + e * x^0 这个可以快速计算

prefix[3] = a * x^2 + b * x^1 + c * x^0
prefix[5] = a * x^4 + b * x^3 + c * x^2 + d * x^1 + e * x^0
prefix[5] - prefix[3]*x^2 = c * x^2 + d * x^1 + e * x^0 = hash(cde)

因此我们可以事先求前缀哈希，然后就计算任意子串的哈希值。
```

## 解决的问题

给定字符串 `s` 和模式串 `p`，求 `s` 中有多少个子串等于 `p`。

## 算法实现

### 方法1: 基础前缀哈希
```go
func CountSubstring(s, p string) int {
    sh := NewStringHash(s)
    pHash := sh.GetStringHash(p)
    
    count := 0
    for i := 0; i <= len(s)-len(p); i++ {
        if sh.GetHash(i, i+len(p)-1) == pHash {
            count++
        }
    }
    return count
}
```

