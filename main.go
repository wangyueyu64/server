package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main() {

	InitDB()
	if err := LogInit(); err != nil {
		fmt.Printf("init log err:%v", err)
	}

	//实例化echo对象。
	e := echo.New()

	//注册一个Get请求, 路由地址为: /hello  并且绑定一个控制器函数, 这里使用的是闭包函数。

	e.POST("/user/sendEmail", SendEmail)
	//e.POST("/user/register", Register)
	e.POST("/user/login", Login)
	e.POST("/user/update", UpdateUser)
	e.GET("/user/getLevel", GetLevel)

	e.POST("/computer/add", Add)
	e.DELETE("/computer/del", Del)
	e.GET("/computer/find", Find)
	e.POST("/computer/update", Update)
	e.POST("/computer/updatelevel", UpdateLevel)

	//e.GET("/test", test)

	//启动http server, 并监听8080端口，冒号（:）前面为空的意思就是绑定网卡所有Ip地址，本机支持的所有ip地址都可以访问。
	e.Start(":8081")

	DB.Close()
}

func test(c echo.Context) error {

	return c.String(http.StatusOK, "登陆成功")
}
