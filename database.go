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
		logger.Errorf("opon database fail")
		return
	}
	fmt.Println("connnect success")
}
func InsertUser(userId string, password string, email string, department string, authorityLevel int) error {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		logger.Errorf("start tx fail")
		return nil
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO user (`user_id`, `password`, `email`, `department`, `authority_level` ) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		logger.Errorf("Prepare fail error:%v\n", err)
		return err
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(userId, password, email, department, authorityLevel)
	if err != nil {
		logger.Errorf("Exec fail")
		return err
	}

	//将事务提交
	tx.Commit()
	logger.Println(res.LastInsertId())
	return nil
}

func Verify(userId string, password string) (bool, error) {

	//查询表
	rows, err := DB.Query(" select password from user where user_id = ?", userId)
	if err != nil {
		logger.Errorf("Query err :%v", err)
		return false, err
	}

	//用完关闭
	defer rows.Close()

	//数据库中记录的密码

	var p string
	for rows.Next() {
		if err = rows.Scan(&p); err != nil {
			logger.Errorf("scan err :%v", err)
			return false, err
		}
	}

	if p != password {
		return false, nil
	}

	return true, nil
}
func AddComputer(computer_id string, mac string, model string, os string) error {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		logger.Errorf("start tx fail")
		return nil
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO computer (`computer_id`, `mac`, `model`, `os`) VALUES (?, ?, ?, ?)")
	if err != nil {
		logger.Errorf("Prepare fail error:%v\n", err)
		return err
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(computer_id, mac, model, os)
	if err != nil {
		logger.Errorf("Exec fail")
		return err
	}

	//将事务提交
	tx.Commit()
	logger.Println(res.LastInsertId())
	return nil
}
func DelComuputer(id string) error {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		logger.Errorf("start tx fail")
		return err
	}
	//准备sql语句
	stmt, err := tx.Prepare("DELETE FROM computer WHERE computer_id = ?")
	if err != nil {
		logger.Errorf("Prepare fail error:%v\n", err)
		return err
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(id)
	if err != nil {
		logger.Errorf("Exec fail")
		return err
	}

	//将事务提交
	tx.Commit()
	logger.Println(res.LastInsertId())
	return nil
}

type Computer struct {
	ComputerId string `json:"computer_id"`
	Mac        string `json:"mac"`
	Model      string `json:"model"`
	Os         string `json:"os"`
	User       string `json:"user"`
	Level      int    `json:"level"`
}

func FindComputer(id string, model string, os string, user string) ([]Computer, error) {

	var rows *sql.Rows
	var err error
	//不筛选
	if id == "" && model == "" && os == "" && user == "" {
		rows, err = DB.Query(" select * from computer ")
		if err != nil {
			logger.Errorf("Query err :%v", err)
			return nil, err
		}
	}
	//只用id
	if id != "" && model == "" && os == "" && user == "" {
		rows, err = DB.Query(" select * from computer where computer_id = ?", id)
		if err != nil {
			logger.Errorf("Query err :%v", err)
			return nil, err
		}
	}
	//只用model
	if id == "" && model != "" && os == "" && user == "" {
		rows, err = DB.Query(" select * from computer where model = ?", model)
		if err != nil {
			logger.Errorf("Query err :%v", err)
			return nil, err
		}
	}
	//只用os
	if id == "" && model == "" && os != "" && user == "" {
		rows, err = DB.Query(" select * from computer where os = ?", os)
		if err != nil {
			logger.Errorf("Query err :%v", err)
			return nil, err
		}
	}
	//只用user
	if id == "" && model == "" && os == "" && user != "" {
		rows, err = DB.Query(" select * from computer where  user = ?", user)
		if err != nil {
			logger.Errorf("Query err :%v", err)
			return nil, err
		}
	}
	//model os
	if id == "" && model != "" && os != "" && user == "" {
		rows, err = DB.Query(" select * from computer where model = ? and os = ?", model, os)
		if err != nil {
			logger.Errorf("Query err :%v", err)
			return nil, err
		}
	}
	//model user
	if id == "" && model != "" && os == "" && user != "" {
		rows, err = DB.Query(" select * from computer where model = ? and user = ?", model, user)
		if err != nil {
			logger.Errorf("Query err :%v", err)
			return nil, err
		}
	}
	//os user
	if id == "" && model == "" && os != "" && user != "" {
		rows, err = DB.Query(" select * from computer where os = ? and user = ?", os, user)
		if err != nil {
			logger.Errorf("Query err :%v", err)
			return nil, err
		}
	}
	//model os user
	if id == "" && model != "" && os != "" && user != "" {
		rows, err = DB.Query(" select * from computer where model = ? and os = ? and user = ?", model, os, user)
		if err != nil {
			logger.Errorf("Query err :%v", err)
			return nil, err
		}
	}

	//用完关闭
	defer rows.Close()

	var computer Computer
	computerSlice := make([]Computer, 0)

	for rows.Next() {
		if err = rows.Scan(&computer.ComputerId, &computer.Mac, &computer.Model, &computer.Os, &computer.User, &computer.Level); err != nil {
			logger.Errorf("scan err :%v", err)
			return nil, err
		}
		computerSlice = append(computerSlice, computer)
	}

	return computerSlice, nil
}

func UpdateComputer(id string, model string, os string, user string) error {

	var res sql.Result
	var err error

	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		logger.Errorf("start tx fail")
		return err
	}

	//model
	if model != "" && os == "" && user == "" {
		stmt, err := tx.Prepare("UPDATE computer SET model = ? WHERE computer_id = ?")
		if err != nil {
			logger.Errorf("Prepare fail error:%v\n", err)
			return err
		}
		//将参数传递到sql语句中并且执行
		res, err = stmt.Exec(model, id)
		if err != nil {
			logger.Errorf("Exec fail")
			return err
		}
	}

	//只用os
	if model == "" && os != "" && user == "" {
		stmt, err := tx.Prepare("UPDATE computer SET os = ? WHERE computer_id = ?")
		if err != nil {
			logger.Errorf("Prepare fail error:%v\n", err)
			return err
		}
		//将参数传递到sql语句中并且执行
		res, err = stmt.Exec(os, id)
		if err != nil {
			logger.Errorf("Exec fail")
			return err
		}
	}

	//只用user
	if model == "" && os == "" && user != "" {
		stmt, err := tx.Prepare("UPDATE computer SET user = ? WHERE computer_id = ?")
		if err != nil {
			logger.Errorf("Prepare fail error:%v\n", err)
			return err
		}
		//将参数传递到sql语句中并且执行
		res, err = stmt.Exec(user, id)
		if err != nil {
			logger.Errorf("Exec fail")
			return err
		}
	}

	//model os
	if model != "" && os != "" && user == "" {
		stmt, err := tx.Prepare("UPDATE computer SET model = ?, os = ? WHERE computer_id = ?")
		if err != nil {
			logger.Errorf("Prepare fail error:%v\n", err)
			return err
		}
		//将参数传递到sql语句中并且执行
		res, err = stmt.Exec(model, os, id)
		if err != nil {
			logger.Errorf("Exec fail")
			return err
		}
	}

	//model user
	if model != "" && os == "" && user != "" {
		stmt, err := tx.Prepare("UPDATE computer SET model = ?, user = ? WHERE computer_id = ?")
		if err != nil {
			logger.Errorf("Prepare fail error:%v\n", err)
			return err
		}
		//将参数传递到sql语句中并且执行
		res, err = stmt.Exec(model, user, id)
		if err != nil {
			logger.Errorf("Exec fail")
			return err
		}
	}

	//os user
	if model == "" && os != "" && user != "" {
		stmt, err := tx.Prepare("UPDATE computer SET os = ?, user = ? WHERE computer_id = ?")
		if err != nil {
			logger.Errorf("Prepare fail error:%v\n", err)
			return err
		}
		//将参数传递到sql语句中并且执行
		res, err = stmt.Exec(os, user, id)
		if err != nil {
			logger.Errorf("Exec fail")
			return err
		}
	}

	//model os user
	if model != "" && os != "" && user != "" {
		stmt, err := tx.Prepare("UPDATE computer SET model = ?, os = ?, user = ? WHERE computer_id = ?")
		if err != nil {
			logger.Errorf("Prepare fail error:%v\n", err)
			return err
		}
		//将参数传递到sql语句中并且执行
		res, err = stmt.Exec(model, os, user, id)
		if err != nil {
			logger.Errorf("Exec fail")
			return err
		}
	}

	//将事务提交
	tx.Commit()
	logger.Println(res.LastInsertId())
	return nil

}
func UpdateComputerLevel(id string, level string) error {

	var res sql.Result
	var err error

	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		logger.Errorf("start tx fail")
		return err
	}

	stmt, err := tx.Prepare("UPDATE computer SET security_level = ? WHERE computer_id = ?")
	if err != nil {
		logger.Errorf("Prepare fail error:%v\n", err)
		return err
	}
	//将参数传递到sql语句中并且执行
	res, err = stmt.Exec(level, id)
	if err != nil {
		logger.Errorf("Exec fail")
		return err
	}

	//将事务提交
	tx.Commit()
	logger.Println(res.LastInsertId())
	return nil

}
func UpdateUsermsg(id string, department string, level string, leader string) error {

	var res sql.Result
	var err error

	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		logger.Errorf("start tx fail")
		return err
	}

	//department
	if department != "" && level == "" && leader == "" {
		stmt, err := tx.Prepare("UPDATE user SET department = ? WHERE user_id = ?")
		if err != nil {
			logger.Errorf("Prepare fail error:%v\n", err)
			return err
		}
		//将参数传递到sql语句中并且执行
		res, err = stmt.Exec(department, id)
		if err != nil {
			logger.Errorf("Exec fail")
			return err
		}
	}

	//level
	if department == "" && level != "" && leader == "" {
		stmt, err := tx.Prepare("UPDATE user SET authority_level= ? WHERE user_id = ?")
		if err != nil {
			logger.Errorf("Prepare fail error:%v\n", err)
			return err
		}
		//将参数传递到sql语句中并且执行
		res, err = stmt.Exec(level, id)
		if err != nil {
			logger.Errorf("Exec fail")
			return err
		}
	}

	//leader
	if department == "" && level == "" && leader != "" {
		stmt, err := tx.Prepare("UPDATE user SET leader= ? WHERE user_id = ?")
		if err != nil {
			logger.Errorf("Prepare fail error:%v\n", err)
			return err
		}
		//将参数传递到sql语句中并且执行
		res, err = stmt.Exec(leader, id)
		if err != nil {
			logger.Errorf("Exec fail")
			return err
		}
	}

	//department level
	if department != "" && level != "" && leader == "" {
		stmt, err := tx.Prepare("UPDATE user SET department= ?, authority_level = ? WHERE user_id = ?")
		if err != nil {
			logger.Errorf("Prepare fail error:%v\n", err)
			return err
		}
		//将参数传递到sql语句中并且执行
		res, err = stmt.Exec(department, level, id)
		if err != nil {
			logger.Errorf("Exec fail")
			return err
		}
	}

	//department leader
	if department != "" && level == "" && leader != "" {
		stmt, err := tx.Prepare("UPDATE user SET department= ?, leader = ? WHERE user_id = ?")
		if err != nil {
			logger.Errorf("Prepare fail error:%v\n", err)
			return err
		}
		//将参数传递到sql语句中并且执行
		res, err = stmt.Exec(department, leader, id)
		if err != nil {
			logger.Errorf("Exec fail")
			return err
		}
	}

	//level leader
	if department == "" && level != "" && leader != "" {
		stmt, err := tx.Prepare("UPDATE user SET authority_level= ?, leader = ? WHERE user_id = ?")
		if err != nil {
			logger.Errorf("Prepare fail error:%v\n", err)
			return err
		}
		//将参数传递到sql语句中并且执行
		res, err = stmt.Exec(level, leader, id)
		if err != nil {
			logger.Errorf("Exec fail")
			return err
		}
	}

	//
	if department != "" && level != "" && leader != "" {
		stmt, err := tx.Prepare("UPDATE user SET department = ?, authority_level = ?, leader = ? WHERE user_id = ?")
		if err != nil {
			logger.Errorf("Prepare fail error:%v\n", err)
			return err
		}
		//将参数传递到sql语句中并且执行
		res, err = stmt.Exec(department, level, leader, id)
		if err != nil {
			logger.Errorf("Exec fail")
			return err
		}
	}

	//将事务提交
	tx.Commit()
	logger.Println(res.LastInsertId())

	return nil
}
func GetUserLevel(userid string) (string, error) {
	var rows *sql.Rows
	var err error
	//不筛选

	//只用id
	rows, err = DB.Query(" select authority_level from user where user_id = ?", userid)
	if err != nil {
		logger.Errorf("Query err :%v", err)
		return "", err
	}

	//用完关闭
	defer rows.Close()

	var level string
	for rows.Next() {
		if err = rows.Scan(&level); err != nil {
			logger.Errorf("scan err :%v", err)
			return "", err
		}
	}
	return level, nil
}
