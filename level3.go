/*
@Title : level 3
@Description : 哔哩哔哩
@Author : 谭靖渝
@Update : 2021/2/7 20:49
*/
package main

import (
	"fmt"
	"os"
	"time"
)

type Author struct {
	Name string //名字
	VIP bool //是否是高贵的带会员
	Icon string //头像
	Signature string//签名
	Focus int //关注人数
}

type video struct {
	Author//视频的作者信息
	title string//视频名称
	Likes int//点赞数
	collections int//收藏数
	coin int //硬币数
	BriefIntroduction string//视频简介

}

func main() {
	var  vid video
	var  aut Author
	var option int
	register(&aut)
	fmt.Println("注册完成,你可以做一下的事来熟悉哔哩哔哩哦~")
	fmt.Println("-------------------------------------------")
	//选项循环
	for  {
		fmt.Println("\n1.查看个人信息  2.发布视频  3.查看发布视频信息  4.观看视频  5.更改用户信息  6.成为大会员！")
		fmt.Println("不想看了就随便按一个键结束这次哔哩哔哩旅行哦~")
		fmt.Scanf("%d",&option)
		//各个选项内容
		switch option {
		//查看个人信息
		case 1:
			DisplayInfoA(&aut)
		//发布视频
		case 2:
			fmt.Println("哇，你的视频真的好看，请给他取一个名字吧！")
			var vtitle,introduce string
			fmt.Scanf("%s",&vtitle)
			vid = PostVideo(aut.Name,vtitle)
			fmt.Println("这么好看的视频，怎么少的了简介呢~")
			fmt.Scanf("%s",&introduce)
			vid.BriefIntroduction = introduce
			fmt.Print(".")
			time.Sleep(time.Second*1)
			fmt.Print(".")
			time.Sleep(time.Second*1)
			fmt.Println(".")
			fmt.Println("发布成功！")
		//查看发布视频信息
		case 3:
			DisplayInfoV(&vid)
		//观看视频
		case 4:
			look(&vid)
		//更改用户信息
		case 5:
			Change(&aut)
		//成为大会员
		case 6:
			Vip(&aut)
		default:
			fmt.Println("欢迎下次↓光↑临~~~~~")
			os.Exit(1)
		}
	}
}

// @title : like
// @description : 点赞数加一
//@auth : 谭靖渝 11:16 2021/2/10
// @param :
// @return :
func (v *video)like() {
	v.Likes++
}

// @title : Collection
// @description : 收藏数加一
// @auth : 谭靖渝 11:16 2021/2/10
// @param :
// @return :
func (v *video) Collection(){

	v.collections++
}

// @title : CoinOperated
// @description : 硬币数加一
// @auth : 谭靖渝 11:17 2021/2/10
// @param :
// @return :
func (v *video)CoinOperated() {
	v.coin++
}

// @title : sanlian
// @description : 一键三连
// @auth : 谭靖渝 11:18 2021/2/10
// @param :
// @return :
func (v *video)sanlian() {
	v.Collection()
	v.like()
	v.CoinOperated()
}

// @title : PostVideo
// @description : 发布视频
// @auth : 谭靖渝 11:19 2021/2/10
// @param :  [Name, title]
// @return :
func PostVideo(Name,title string)video  {
	var v video
	v.title = title
	v.Author.Name = Name
	return v
}

// @title : register
// @description : 注册成员
// @auth : 谭靖渝 11:19 2021/2/10
// @param :  [aut]
// @return :
func register(aut *Author){
	fmt.Println("欢迎来到哔哩哔哩,要想成为一个up🐖，就要先注册账号哦")
	fmt.Println("-------------------------------------------")
	fmt.Print("昵称 : ")
	fmt.Scanf("%s",&aut.Name)
	fmt.Print("签名 : ")
	fmt.Scanf("%s",&aut.Signature)
	fmt.Print("头像 : ")
	fmt.Scanf("%s",&aut.Icon)
	fmt.Println("-----------------注册成功-------------------")
	fmt.Println("        （ ゜- ゜）つロ 干杯~bilibili        ")
	fmt.Println("请确认信息 : 昵称 : ",aut.Name,"  签名  : ",aut.Signature,"  头像 : ",aut.Icon)
}

// @title : DisplayInfoA
// @description : 显示成员信息
// @auth : 谭靖渝 11:20 2021/2/10
// @param :  [aut]
// @return :
func DisplayInfoA(aut *Author)  {
	fmt.Println("信息查询中......")
	time.Sleep(time.Second * 2)
	fmt.Println(" 昵称 : ",aut.Name)
	fmt.Println(" 签名 : ",aut.Signature)
	fmt.Println(" 头像 : ",aut.Icon)
	if aut.VIP {
		fmt.Println(" 高贵会员 : 是")
	}else {
		fmt.Println(" 高贵会员 : 否")
	}
}

// @title : DisplayInfoV
// @description : 显示视频信息
// @auth : 谭靖渝 11:20 2021/2/10
// @param :  [vid]
// @return :
func DisplayInfoV(vid *video)  {
	fmt.Println("作者名 : ",vid.Name)
	fmt.Println("视频名 : ",vid.title)
	fmt.Println("视频简介 : ",vid.BriefIntroduction)
	fmt.Println("视频的 : © ",vid.coin," 👍 ",vid.Likes," ⭐︎",vid.collections)
}

// @title : Vip
// @description : 充值大会员
// @auth : 谭靖渝 11:20 2021/2/10
// @param :  [aut]
// @return :
func Vip(aut *Author)  {
	var judge string
	fmt.Println("开通大会员，提前看刺客五六七哦~✂")
	fmt.Println("------------Yes------------No--------------")
	fmt.Scanf("%s",&judge)
	switch judge {
	case "Yes":
		aut.VIP = true
		fmt.Println("恭喜大大成为大会员~")
	case "No":
		aut.VIP = false
		fmt.Println("毕竟大家要恰饭的啦~")
	}
}

// @title : look
// @description : 观看视频
// @auth : 谭靖渝 11:21 2021/2/10
// @param :  [vid]
// @return :
func look(vid *video)  {
	var judge int
	fmt.Println("你看完了这个视频，感觉怎么样呢")
	fmt.Println("1.一键三连  2.点赞  3.投币  4.收藏")
	fmt.Scanf("%d",&judge)
	switch judge {
	case 1:
		vid.sanlian()
	case 2:
		vid.like()
	case 3:
		vid.CoinOperated()
	case 4:
		vid.Collection()
	}
}

// @title : Change
// @description : 修改用户信息
// @auth : 谭靖渝 11:21 2021/2/10
// @param :  [aut]
// @return :
func Change(aut *Author)  {
	fmt.Print("昵称 : ")
	fmt.Scanf("%s",&aut.Name)
	fmt.Print("签名 : ")
	fmt.Scanf("%s",&aut.Signature)
	fmt.Print("头像 : ")
	fmt.Scanf("%s",&aut.Icon)
	fmt.Print(".")
	time.Sleep(time.Second*1)
	fmt.Print(".")
	time.Sleep(time.Second*1)
	fmt.Println(".")
	time.Sleep(time.Second*1)
	fmt.Println("-----------------修改成功-------------------")
}