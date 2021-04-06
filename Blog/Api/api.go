/*
@Title : api
@Description : 各种接口
@Author : 谭靖渝
@Update : 2021/3/29 20:15
*/
package Api

import (
	"Blog/Model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//状态验证
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

//用户登录,并且设置cookie
func Login(c *gin.Context) {
	var user Model.User
	//获取从服务器发来的UID,password,并且返回相应的信息
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
	} else {
		if user.LoginUser(user.ID, user.PassWord) == 0 {
			c.JSON(404, gin.H{
				"error": "账号错误",
			})
		} else if user.LoginUser(user.ID, user.PassWord) == 1 {
			c.JSON(404, gin.H{
				"error": "密码错误",
			})
		} else {
			//登录成功后设置cookie并且返回相关信息
			c.SetCookie("User", "Yes", 3600, "/", "127.0.0.1",4,false, false)
			c.JSON(http.StatusOK, *(user.LoginUser(user.ID, user.PassWord).(*Model.User)))
		}
	}
}

//用户注册
func Register(c *gin.Context) {
	//获取前端发来的username 和 password,注册成功则返回相关信息
	var user Model.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, *(user.RegisterUser(user.Username, user.PassWord)))
	}
}
//修改密码
func ChangePassword(c *gin.Context)  {
	//获取相应的id
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	var user Model.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,*(Model.Update(id,3,"",user.PassWord)))
	}
}

//为某账号的言论点赞
func Likes(c *gin.Context) {
	//获取相应的的ID
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	var user Model.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,*(Model.Update(id,3,user.Speech,"")))
	}
}

//发表言论
func Speech(c *gin.Context)  {
	//获取相应的的ID
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	//绑定获取言论，并且更新
	var user Model.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,*(Model.Update(id,2,user.Speech,"")))
	}
}

//删除账号
func DeleteNum(c *gin.Context)  {
	//获取相应的id
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	//删除对应的账号
	if Model.Delete(id) {
		c.JSON(http.StatusOK,gin.H{
			"massage":"删除成功",
		})
	}else {
		c.JSON(404,gin.H{
			"error":"删除失败",
		})
	}
}

//获取账号相关信息，不用判断是否为空，前面有中间件
func GetUser(c *gin.Context)  {
	//获取相应的id
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	fmt.Println(id)
	//查询相关的id信息并且返回
	c.JSON(http.StatusOK,*(Model.Query(id)))
}
