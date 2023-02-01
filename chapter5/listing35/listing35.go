// 这个示例程序展示bytes.Buffer也可以用于io.Copy函数
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// main 是应用程序的入口
func main() {
	var b bytes.Buffer

	// 将字符创写入Buffer
	b.Write([]byte("Hello"))

	// 使用Fprintf将字符串拼接到Buffer
	fmt.Fprintf(&b, "World!")

	// 将Buffer的内容写到Stdout
	io.Copy(os.Stdout, &b)
}
