/*
@Title : init
@Description : 连接数据库
@Author : 谭靖渝
@Update : 2021/3/29 11:05
*/
package Model

import "fmt"

func Initial() {
	//连接数据库
	err := InitMySQL()
	if err != nil {
		fmt.Println(err)
	}
}
