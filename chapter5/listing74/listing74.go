// 这个示例程序展示公开的结构类型中如何访问未公开的内嵌类型的例子
package main

import (
	"fmt"

	"goinaction/chapter5/listing74/entities"
)

// main 是应用程序的入口
func main() {
	// 创建entitles包中的Admin类型的值
	a := entities.Admin{
		Rights: 10,
	}

	// 设置未公开的内部类型的公开字段的值
	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)
}
