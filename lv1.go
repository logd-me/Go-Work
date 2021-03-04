/*
@Title : lv1
@Description : 并发协程同步求素数
@Author : 谭靖渝
@Update : 2021/2/22 11:00
*/

//筛选算法核心算法，素数的n倍数不会是素数
package main

import (
	"fmt"
)

// @title : create
// @description : 产生1~50000的自然数数据流
// @auth : 谭靖渝 11:51 2021/2/22
// @param :  [ch]
// @return :
func create(ch chan int) {

	for i := 2; i < 50000; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
	//关闭产生完毕的ch，不然会因为filter内访问没有关闭的通道导致无限访问，造成死锁
	close(ch)
}


// @title : filter
// @description : 筛子算法
// @auth : 谭靖渝 11:03 2021/2/22
// @param :  [in, out, prime]
// @return :
func filter(in, out chan int, prime int) {
	for {


		//进行对prime的筛选
		//防止最后死循环在filter之中
		i, ok := <-in
		if !ok {
			break
		}
		//判断in送来的数为素数就放入out管道
		if i%prime != 0 {
			out <- i
		}
	}
	//防止访问导致协程死锁
	close(out)

}

func main() {
	ch := make(chan int)
	//产生自然数对
	go create(ch)
	//循环进行筛选
	 for {
		prime := <-ch
		 if prime==0 {
			 break
		 }
		fmt.Print(prime, " \n")
		ch1 := make(chan int)
		//进行筛选,每一趟对最小素数进行筛选，是这个素数的倍数全部被排除
		go filter(ch, ch1, prime)
		//将筛选过的通道赋值给下一次要进行筛选的通道
		ch = ch1

	}
}
