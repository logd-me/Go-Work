/*
@Title : DataBase
@Description : 数据库的链接，增删改查
@Author : 谭靖渝
@Update : 2021/3/29 11:08
*/
package Model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)
//
var DB *sql.DB
var m sync.Mutex

func InitMySQL()(err error) {
	//链接数据库
	dsn := "root:root@(127.0.0.1:3306)/Blog?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	//测试数据库是否链接成功
	err = DB.Ping()
	if err != nil {
		return err
	}
	return nil
}


func Create(username string,password string)(UID string)  {
	//创建账户，用于注册
	UID=Random(6)
	sqlStr := "insert into users(id,username,password) values (?,?,?)"
	m.Lock()
	_, err := DB.Exec(sqlStr,UID,username,password)
	m.Unlock()
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	return UID
}

func Query(UID string)(user *User)  {
	//查询账号的相关信息
	sqlStr :="select id,username,password,speech,likes from users where id = ?"
	//将相关的信息存储到响应的结构体里面
	var u User
	err := DB.QueryRow(sqlStr,UID ).Scan(&u.ID, &u.Username, &u.PassWord,&u.Speech,&u.Like)
	if err != nil {
		fmt.Println(err)
	}
	return &u
}

func Update(UID string,judge int,speech string,password string)(user *User)  {
	u := Query(UID)
	if judge==1 {
		//点赞数加一
		m.Lock()
		u.Like++
		sqlStr := "update users set likes=? where id = ?"
		m.Unlock()
		_, err := DB.Exec(sqlStr, u.Like, UID)
		if err != nil {
			fmt.Printf("update failed, err:%v\n", err)
			return
		}
	}else if judge==2 {
		//更新说说，将原本的说说点赞数变成零
		m.Lock()
		u.Like=0
		sqlStr := "update users set speech=?,likes=? where id = ?"
		_, err := DB.Exec(sqlStr, speech,u.Like, UID)
		m.Unlock()
		if err != nil {
			fmt.Printf("update failed, err:%v\n", err)
			return
		}
	}else if judge==3{
		m.Lock()
		sqlStr :="update users set password=? where id = ?"
		_, err := DB.Exec(sqlStr,password, UID)
		m.Unlock()
		if err != nil {
			fmt.Printf("update failed, err:%v\n", err)
			return
		}
	}
	//更行过后重新查询一遍，并且返回结果
	u1 := Query(UID)
	return u1
}

func Delete(UID string)(judge bool)  {
	//删除账号，
	m.Lock()
	sqlStr := "delete from users where id = ?"
	_, err := DB.Exec(sqlStr, UID)
	m.Unlock()
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return false
	}
	return true
}

