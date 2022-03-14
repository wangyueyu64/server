package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

//数据库配置
const (
	userName = "root"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "wyy"
)

// DB Db数据库连接池
var DB *sql.DB

// InitDB 注意方法名大写，就是public
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
}
func InsertUser(userId string, password string, email string, department string, authorityLevel int) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("start tx fail")
		return false
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO user (`user_id`, `password`, `email`, `department`, `authority_level` ) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Printf("Prepare fail error:%v\n", err)
		return false
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(userId, password, email, department, authorityLevel)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true
}
