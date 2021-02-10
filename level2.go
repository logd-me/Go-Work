/*
@Title : level 2
@Description : 接口断言
@Author : 谭靖渝
@Update : 2021/2/6 22:13
*/
package main

import "fmt"

func main()  {
	//制作接口切片
	i := make([]interface{},3)
	i[0] = "你好吗"
	i[1] = 32
	i[2] = true
	receiver(i)
	
}

// @title : receiver
// @description : 接口断言判断
// @auth : 谭靖渝 22:47 2021/2/6
// @param :  [v]
// @return :
func receiver(v []interface{}) {
	for _, data := range v{
		switch value:=data.(type) {
		case string:
			fmt.Println("我接收到的是",value,"这个是string")
		case int:
			fmt.Println("我接收到的是",value,"这个是int")
		case bool:
			fmt.Println("我接收到的是",value,"这个是bool")
		}
	}
}