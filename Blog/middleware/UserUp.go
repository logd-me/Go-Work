/*
@Title : UserUp
@Description : 中间件，判断用户是否登录
@Author : 谭靖渝
@Update : 2021/3/30 14:38
*/
package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JudgeLogin()gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("User"); err == nil {
			if cookie == "Yes" {
				c.Next()
				return
			}
		}
		// 返回错误,并且跳转到登录页面
		c.Redirect(http.StatusMovedPermanently,"127.0.0.1/v1/UserLogin")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}
