package main

import (
	"fmt"
	"github.com/go-gomail/gomail"
	"github.com/labstack/echo"
	"math/rand"
	"net/http"
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
		fmt.Printf("Verify err:%v", err)
		return c.String(http.StatusInternalServerError, "服务器错误！")
	}
	if ok == false {
		return c.String(http.StatusForbidden, "密码错误！")
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
		panic(err)
	}
	return nil
}

//// SendMsg 向手机发送验证码
//func SendMsg(tel string, code string) string {
//	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "<accesskeyId>", "<accessSecret>")
//	request := dysmsapi.CreateSendSmsRequest()
//	request.Scheme = "https"
//	request.PhoneNumbers = tel             //手机号变量值
//	request.SignName = "凌睿工作室"             //签名
//	request.TemplateCode = "SMS_19586XXXX" //模板编码
//	request.TemplateParam = "{\"code\":\"" + code + "\"}"
//	response, err := client.SendSms(request)
//	fmt.Println(response.Code)
//	if response.Code == "isv.BUSINESS_LIMIT_CONTROL" {
//		return "frequency_limit"
//	}
//	if err != nil {
//		fmt.Print(err.Error())
//		return "failed"
//	}
//	return "success"
//}
//
//// Code 随机验证码
//func Code() string {
//	rand.Seed(time.Now().UnixNano())
//	code := rand.Intn(899999) + 100000
//	res := strconv.Itoa(code) //转字符串返回
//	return res
//}
//
//// 接收手机号并发送验证码
//func getValidationHandler(c *gin.Context) {
//	var user User
//	c.ShouldBind(&user)
//	code := Code()
//	fmt.Println(code)
//
//	sendRes := SendMsg(user.Tel, code)
//	if sendRes == "failed" {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"msg": "error",
//		})
//	} else {
//		if !SetRedis(user.UserId, code) {
//			c.JSON(http.StatusInternalServerError, gin.H{
//				"msg": "error",
//			})
//		}
//		c.JSON(http.StatusOK, gin.H{
//			"msg": sendRes,
//		})
//	}
//}
//
//// Validation 在注册时检查验证码
//func Validation(validation string, userId string) int {
//	var flag int
//	getcode := GetRedis(userId)
//
//	if validation == getcode {
//		flag = 1
//	} else {
//		flag = 0
//	}
//	return flag
//}
//func registerHandler(c *gin.Context) {
//	var user User
//	err := c.BindJSON(&user)
//	if err != nil {
//		fmt.Println(err)
//		c.JSON(http.StatusBadRequest, gin.H{
//			"msg": "error",
//		})
//		return
//	}
//
//	if Validation(user.Validation, user.UserId) == 0 {
//		c.JSON(http.StatusOK, gin.H{
//			"msg": "wrong_code",
//		})
//		return
//	}
//}
//func SetRedis(userId string, code string) bool {
//	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
//	if err != nil {
//		fmt.Println("connect redis error :", err)
//		return false
//	}
//	defer conn.Close()
//	_, err = conn.Do("SET", userId, code)
//	if err != nil {
//		fmt.Println("redis set error:", err)
//		return false
//	}
//	_, err = conn.Do("expire", userId, 300)
//	if err != nil {
//		fmt.Println("set expire error: ", err)
//		return false
//	}
//	return true
//}
//
//func GetRedis(userId string) string {
//	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
//	if err != nil {
//		fmt.Println("connect redis error :", err)
//	}
//	defer conn.Close()
//	code, err := redis.String(conn.Do("GET", userId))
//	if err != nil {
//		fmt.Println("redis get error:", err)
//	}
//	return code
//}
