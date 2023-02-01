// 这个示例程序展示Go语言里如何使用接口
package main

import "fmt"

// notifier 是一个定义了通知类行为的接口
type notifier interface {
	notify()
}

// user 在程序里定义了一个用户类型
type user struct {
	name  string
	email string
}

// notify 是使用指针接收者实现的方法
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// main 是应用程序的入口
func main() {
	// 创建一个user类型的值，并发送通知
	u := user{"Bill", "bill@example.com"}

	sendNotifiaction(u)

	// ./listing36.go:29:18: 不能将u（类型是user）作为
	// sendNotification的参数类型notifier：
	// user类型并没有实现notifier（notify方法使用指针接收者声明）
}

// sendNotifiaction 接收一个实现了notifier接口的值并发送通知
func sendNotifiaction(n notifier) {
	n.notify()
}
