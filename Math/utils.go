package math

// 十进制转 x 进制
func NumToBaseX(num, x int) int {
	res := 0
	v := []int{}
	for num > 0 {
		v = append(v, num%x)
		num /= x
	}
	for i := len(v) - 1; i >= 0; i-- {
		res = res*10 + v[i]
	}
	return res
}

// 逆元
/*
费马小定理：
如果 p 是一个质数，且整数 a 不是 p 的倍数，那么 a^(p-1) == 1 (mod p)

在实数中，b*b^(-1) = 1。在模运算中，我们希望 b*x == 1 (mod p),x 就是我们要求的逆元
a^(p-1) == 1 (mod p) 推导：
a * a^(p-2) == 1 (mod p)
此时 a^(p-2) 就是 a 的逆元
因此通过快速幂，求出 a 的 p-2 次方，再取模，就是答案
*/
func Inv(a, mod int) int {
	if a == 2 {
		return (mod + 1) / 2
	}
	return Pow(a, mod-2, mod)
}

// 快速幂
func Pow(a, b, mod int) int {
	res := 1
	for b > 0 {
		if b&1 > 0 {
			res = res * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return res
}

// 数论分块
/*
向下取整的数论分块
[n/i]，求 i = 1, 2, 3, ..., n 的和，当 n 比较大的情形，怎么做？
暴力做法是一个一个的枚举，但是通过找规律发现，在连续区间，是存在大量重复的值的。举个例子
n = 7
i = 1, q = 7
i = 2, q = 3
i = 3, q = 2
i = 4, 5, 6, 7, q = 1
看，当 i = 4,5,6,7 时，q 是一样的，能不能快速计算这些重复的值呢

当知道固定的 q，能不能快速求到区间呢？
q = n/i 的下取整
因此 q <= n/i < q + 1
同时取倒数， 1/q >= i/n > 1/(q+1) 有  n/q >= i >= n/(q+1)
所以知道一个固定的 q, 其 i 的区间为 [n/(q+1)+1, n/q] 闭区间
*/

// 该函数用于求 i=1,2,3...,n 的 n/i 下取整数的和
func H(n int) int {
	var l, r int
	l = 1
	ans := 0
	for l <= n {
		q := n / l
		r = n / q
		ans += (r - l + 1) * q
		l = r + 1
	}
	return ans
}
