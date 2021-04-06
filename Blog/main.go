/*
@Title : main
@Description : 运行
@Author : 谭靖渝
@Update : 2021/3/29 11:16
*/
package main

import (
	"Blog/Model"
	"Blog/Serve"
)

func main()  {
	Model.Initial()
	r:=Serve.NewRouter()
	r.Run("127.0.0.1:8080")

}