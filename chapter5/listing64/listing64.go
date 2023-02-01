// 这个示例展示无法从另一个包里访问未公开的标识符
package main

import (
	"fmt"

	"goinaction/chapter5/listing64/counters"
)

// main是应用程序的入口
func main() {
	// 创建一个未公开的类型的变量
	// 并将其初始化为10
	counter := counters.alertCounter(10)

	// ./listing64.go:15: 不能引用未公开的名字
	//                counters.alertCounter
	// ./listing64.go:15: 未定义: counters.alertCounter

	fmt.Printf("Counter: %d\n", counter)
}
