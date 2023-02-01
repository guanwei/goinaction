// 这个示例展示无法从另一个包里访问未公开的标识符的值
package main

import (
	"fmt"

	"goinaction/chapter5/listing68/counters"
)

// main 是应用程序的入口
func main() {
	// 使用counters包公开的New函数来创建一个未公开的类型的变量
	counter := counters.New(10)

	fmt.Printf("Counter: %d\n", counter)
}
