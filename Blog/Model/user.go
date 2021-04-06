/*
@Title : user
@Description :
@Author : 谭靖渝
@Update : 2021/3/29 11:23
*/
package Model

type User struct {
	//用户的账号
	ID string `json:"id" form:"id"`
	//用户姓名
	Username string `json:"username" form:"username"`
	//用户的密码
	PassWord string `json:"password" form:"password"`
	//用户的博客
	Speech string `json:"speech" form:"speech"`
	//用户博客的点赞数
	Like int `json:"like" form:"like"`
}

//用户注册,并且返回相应的信息。
func (users *User)RegisterUser(username string,password string)(user *User)  {
	UID:=Create(username,password)
	//查询创建好的用户，进行返回
	return  Query(UID)
}

//用户登录，如果登录成功则返回相应信息,0是账号不对，1是密码不对
func (users *User)LoginUser(UID string,password string)(in interface{})  {
	u:=Query(UID)
	if u.ID==""{
		return 0
	}else if password==u.PassWord{
		return u
	}
	return 1
}
