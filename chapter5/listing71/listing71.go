// 这个示例程序展示公开的结构类型中未公开的字段
// 无法直接访问
package main

import (
	"fmt"
	"goinaction/chapter5/listing71/entities"
)

// main 是应用程序的入口
func main() {
	// 创建entities包中的User类型的值
	u := entities.User{
		Name:  "Bill",
		email: "bill@email.com",
	}

	// ./listing71.go:15:3: 结构字面量中结构entities.User的字段'email'未知

	fmt.Printf("User: %v\n", u)
}
