// 这个示例程序展示如何使用atomic包来提供对数值类型的安全访问
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// counter 是所有goroutine都要增加其值的变量
	counter int64

	// wg用来等待程序结束
	wg sync.WaitGroup
)

// main 是所有Go程序的入口
func main() {
	// 计数加2，表示要等待两个goroutine
	wg.Add(2)

	// 创建两个goroutine
	go incCounter(1)
	go incCounter(2)

	// 等待goroutine结束
	wg.Wait()

	// 显示最终的值
	fmt.Println("Final Counter:", counter)
}

// incCounter 增加包里counter变量的值
func incCounter(id int) {
	// 在函数退出时调用Done来通知main函数工作已经完成
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 安全地对counter加1
		atomic.AddInt64(&counter, 1)

		// 当前goroutine从线程退出，并放回到队列
		runtime.Gosched()
	}
}
