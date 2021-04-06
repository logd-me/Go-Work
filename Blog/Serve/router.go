/*
@Title : router
@Description : 设置路由器，路由组
@Author : 谭靖渝
@Update : 2021/3/30 14:27
*/
package Serve

import (
	"Blog/Api"
	"github.com/gin-gonic/gin"
	"net/http"
	"Blog/middleware"
)

func NewRouter() *gin.Engine{
	//设置默认引擎
	r:=gin.Default()
	//加载主页
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})
	//设置路由组
	v1Group := r.Group("v1")
	{
		//验证是否连接成功
		v1Group.POST("/pong",Api.Ping)
		//发表言论
		v1Group.POST("/speech/:id",middleware.JudgeLogin(),Api.Speech)
		//点赞
		v1Group.PUT("/speech/:id",middleware.JudgeLogin(),Api.Likes)
		//登录
		v1Group.POST("/UserLogin",Api.Login)
		//注册
		v1Group.POST("/UserRegister",Api.Register)
		//删除用户
		v1Group.DELETE("/User/:id",Api.DeleteNum)
		//修改密码
		v1Group.PUT("/ChangePassword/:id",middleware.JudgeLogin(),Api.ChangePassword)
		//获取用户信息
		v1Group.GET("/User/:id",middleware.JudgeLogin(),Api.GetUser)
	}
	return r
}
