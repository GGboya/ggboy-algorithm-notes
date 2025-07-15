package math

import (
	"testing"
)

func TestNumToBaseX(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		x        int
		expected int
	}{
		{
			name:     "十进制转二进制 - 10",
			num:      10,
			x:        2,
			expected: 1010,
		},
		{
			name:     "十进制转二进制 - 15",
			num:      15,
			x:        2,
			expected: 1111,
		},
		{
			name:     "十进制转八进制 - 25",
			num:      25,
			x:        8,
			expected: 31,
		},
		{
			name:     "十进制转八进制 - 100",
			num:      100,
			x:        8,
			expected: 144,
		},
		{
			name:     "十进制转十六进制 - 255",
			num:      255,
			x:        16,
			expected: 255,
		},
		{
			name:     "十进制转十六进制 - 1000",
			num:      1000,
			x:        16,
			expected: 1896,
		},
		{
			name:     "十进制转三进制 - 10",
			num:      10,
			x:        3,
			expected: 101,
		},
		{
			name:     "十进制转五进制 - 25",
			num:      25,
			x:        5,
			expected: 100,
		},
		{
			name:     "零值测试",
			num:      0,
			x:        2,
			expected: 0,
		},
		{
			name:     "十进制转十进制 - 123",
			num:      123,
			x:        10,
			expected: 123,
		},
		{
			name:     "十进制转七进制 - 49",
			num:      49,
			x:        7,
			expected: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NumToBaseX(tt.num, tt.x)
			if result != tt.expected {
				t.Errorf("NumToBaseX(%d, %d) = %d, want %d", tt.num, tt.x, result, tt.expected)
			}
		})
	}
}

// 基准测试
func BenchmarkNumToBaseX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumToBaseX(1000, 2)
	}
}
