// 这个示例程序使用接口展示多态行为
package main

import "fmt"

// notifier 是一个定义了通知类行为的接口
type notifier interface {
	notify()
}

// user 在程序里定义一个用户类型
type user struct {
	name  string
	email string
}

// notify 使用指针接收者实现了notifier接口
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// admin 定义了程序里的管理员
type admin struct {
	name  string
	email string
}

// notify 使用指针接收者实现了notifier接口
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name,
		a.email)
}

// main 是应用程序的入口
func main() {
	// 创建一个user值并传给sendNotification
	bill := user{"Bill", "bill@email.com"}
	sendNotification(&bill)

	// 创建一个admin值并传给sendNotification
	lisa := admin{"Lisa", "lisa@email.com"}
	sendNotification(&lisa)
}

// sendNotification 接受一个实现了notifier接口值并发送通知
func sendNotification(n notifier) {
	n.notify()
}
