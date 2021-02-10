/*
@Title : level 1
@Description : 人类的四种接口
@Author : 谭靖渝
@Update : 2021/2/6 21:19
*/
package main

import "fmt"

//人类
type person struct {
	name string
	age int
	gender string
}
//鸽子
type dove interface {
	gugugu()//鸽
}
//复读机
type repeater interface {
	repeat(str string)
}
//柠檬精
type lenmon interface {
	suan()
}
//真香怪
type fragrant interface {
	aimaya()
}
//四种接口的总和
type TheSum interface {
	dove
	repeater
	lenmon
	fragrant
	realization(str string)
}


func main() {
	per := person{"m", 18,"mike" }
	per.name = "人类"
	per.age = 18
	per.gender = "未知"
	//声明接口变量
	var g TheSum
	g = &per
	//实现
	g.realization("复读机")

	
}

// @title : gugugu
// @description : dove接口的方法
// @auth : 谭靖渝 22:05 2021/2/6
// @param :
// @return :
func (p *person)gugugu()  {
	fmt.Println("咕咕咕，对不起我又鸽了")
}

// @title : repeat
// @description : repeater的方法
// @auth : 谭靖渝 22:06 2021/2/6
// @param :  [str]
// @return :
func(p *person)repeat(str string)() {
	fmt.Println(str)
}

// @title : suan
// @description : lemon的接口方法
// @auth : 谭靖渝 22:07 2021/2/6
// @param :
// @return :
func (p *person)suan() {
	fmt.Println("不过如此")
}

// @title : aimaya
// @description : fragrant的接口方法
// @auth : 谭靖渝 22:07 2021/2/6
// @param :
// @return :
func (p *person)aimaya()  {
	fmt.Println("哎妈呀，真香")
}

// @title : realization
// @description : TheSum 的接口方法
// @auth : 谭靖渝 22:11 2021/2/6
// @param :  [str]
// @return :
func (p *person)realization(str string){
	p.gugugu()
	p.repeat(str)
	p.aimaya()
	p.suan()
}