package main

import (
	"fmt"
	"github.com/go-gomail/gomail"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var codeMap = make(map[string]string)

//func Register(c echo.Context) error {
//	account := c.QueryParam("account")
//	password := c.QueryParam("password")
//	email := c.QueryParam("email")
//
//
//	return c.String(http.StatusOK, "register success")
//}

var temp int = 0

func Login(c echo.Context) error {
	account := c.QueryParam("account")
	password := c.QueryParam("password")
	code := c.QueryParam("code")

	_, ok := codeMap[account] // 如果key1存在则ok == true，否则ok为false
	if ok == false {
		return c.String(http.StatusForbidden, "请先发送验证码")
	}

	if codeMap[account] == "" {
		return c.String(http.StatusForbidden, "验证码失效，请重新发送")
	}

	if code != codeMap[account] {
		return c.String(http.StatusForbidden, "验证码错误！")
	}

	ok, err := Verify(account, password)
	if err != nil {
		logger.Errorf("Verify err:%v", err)
		return c.String(http.StatusInternalServerError, "服务器错误！")
	}
	if ok == false {
		if temp < 2 {
			temp++
			return c.String(http.StatusForbidden, "密码错误！")
		}
		if temp == 2 {
			temp = 0
			// 获取用户等级
			l, err := GetUserLevel(account)
			if err != nil {
				logger.Errorf("GetUserLevel err :%v", err)
				return err
			}

			l2, err2 := strconv.Atoi(l)
			if err2 != nil {
				logger.Errorf("Atoi err :%v", err)
				return err
			}

			l2--
			err = UpdateUsermsg(account, "", strconv.Itoa(l2), "")
			if err != nil {
				logger.Errorf("UpdateUsermsg err : %v", err)
				return err
			}
			return c.String(http.StatusForbidden, "连续三次密码错误！用户安全等级已调整！")
		}

	}

	return c.String(http.StatusOK, "登陆成功")
}

func SendEmail(c echo.Context) error {

	account := c.QueryParam("account")
	email := c.QueryParam("email")

	rand.Seed(time.Now().Unix())
	num := fmt.Sprintf("%06v", rand.Int31n(1000000))

	codeMap[account] = num
	fmt.Println(num)

	m := gomail.NewMessage()
	m.SetAddressHeader("From", "2578103136@qq.com", "网络资产管理") // 发件人
	m.SetHeader("To",                                         // 收件人
		m.FormatAddress(email, ""),
	)
	m.SetHeader("Subject", "网络资产管理系统") // 主题
	m.SetBody("text/html", ""+num)     // 正文

	d := gomail.NewPlainDialer("smtp.qq.com", 465, "2578103136@qq.com", "wgcojyvwwfyuebag") // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		logger.Errorf("send verification code err: %v", err)
		return c.String(http.StatusInternalServerError, "服务器错误！")
	}
	return nil
}

func Add(c echo.Context) error {

	macAddr := c.QueryParam("mac")
	model := c.QueryParam("model")
	os := c.QueryParam("os")

	Uuid := fmt.Sprintf("%v", uuid.NewString())

	if err := AddComputer(Uuid, macAddr, model, os); err != nil {
		logger.Errorf("AddComputer err :%v", err)
		return c.String(http.StatusInternalServerError, "服务器错误！")
	}

	return nil
}

func Del(c echo.Context) error {

	computerId := c.QueryParam("id")

	if err := DelComuputer(computerId); err != nil {
		logger.Errorf("DelComputer err :%v", err)
		return c.String(http.StatusInternalServerError, "服务器错误！")
	}

	return nil
}

func Find(c echo.Context) error {
	computerId := c.QueryParam("id")
	model := c.QueryParam("model")
	os := c.QueryParam("os")
	user := c.QueryParam("user")

	result, err := FindComputer(computerId, model, os, user)
	if err != nil {
		logger.Errorf("FindComputer err :%v", err)
		return c.String(http.StatusInternalServerError, "服务器错误！")
	}

	return c.JSON(http.StatusOK, result)
}

func Update(c echo.Context) error {
	computerId := c.QueryParam("id")
	model := c.QueryParam("model")
	os := c.QueryParam("os")
	user := c.QueryParam("user")

	err := UpdateComputer(computerId, model, os, user)
	if err != nil {
		logger.Errorf("UpdateComputer err :%v", err)
		return c.String(http.StatusInternalServerError, "服务器错误！")
	}

	return nil

}

func UpdateUser(c echo.Context) error {
	userId := c.QueryParam("id")
	department := c.QueryParam("department")
	level := c.QueryParam("level")
	leader := c.QueryParam("leader")

	err := UpdateUsermsg(userId, department, level, leader)
	if err != nil {
		logger.Errorf("UpdateUser err :%v", err)
		return c.String(http.StatusInternalServerError, "服务器错误！")
	}

	return nil
}
func UpdateLevel(c echo.Context) error {
	computerId := c.QueryParam("id")
	level := c.QueryParam("level")

	err := UpdateComputerLevel(computerId, level)
	if err != nil {
		logger.Errorf("UpdateComputer err :%v", err)
		return c.String(http.StatusInternalServerError, "服务器错误！")
	}

	return nil

}

func GetLevel(c echo.Context) error {

	userId := c.QueryParam("userid")

	result, err := GetUserLevel(userId)
	if err != nil {
		logger.Errorf("getUserLevel err :%v", err)
		return c.String(http.StatusInternalServerError, "服务器错误！")
	}

	return c.String(http.StatusOK, result)
}
