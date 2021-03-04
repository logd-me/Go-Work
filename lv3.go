/*
@Title : lv3
@Description :
@Author : 谭靖渝
@Update : 2021/2/27 9:58
*/
package main

import (
	"fmt"
	"protest/MyContext"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Second)
	ctx,cancel:=MyContext.MyWithCancel()
	go worker1(ctx, ticker)
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(50 * time.Millisecond)
}

func worker1( ctx MyContext.MyEemCtx, t *time.Ticker) {
	// Ticker是一个计时器，每过一段时间就向管道Ticker.C中发送一个Time结构体

	go func() {
		i:=1
		defer fmt.Println("协程退出")
		// Using stop channel explicit exit
		for {
			select {
			case <-ctx.MyDone():
				// ctx.Done()是一个方法，返回一个chan struct{}类型的管道
				// 当主函数中调用了cancel()方法，就会关闭这个管道
				// 按照我们前面讲到，这时会从管道拿到一个零值，用于通知结束
				fmt.Println("接收停止信号")
				return
			case <-t.C:

				fmt.Println("正在工作",i)
				i++
				// do something
			}
		}
	}()
	//return
}