// 用来检测要将整数值转为字符串，使用哪个函数会更好的基准
// 测试示例。先使用strconv.FormatInt函数，最后使用strconv.Itoa
package listing28_test

import (
	"fmt"
	"strconv"
	"testing"
)

// BenchmarkSprintf 对fmt.Sprintf函数进行基准测试
func BenchmarkSprintf(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}

// BenchmarkFormat 对strconv.FormatInt函数进行基准测试
func BenchmarkFormat(b *testing.B) {
	number := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

// BenchmarkItoa 对strconv.Itoa函数进行基准测试
func BenchmarkItoa(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}
