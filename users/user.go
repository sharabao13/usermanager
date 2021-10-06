package users

import (
	"crypto/md5"
	"fmt"
	"github.com/howeyc/gopass"
	"strconv"
	"strings"
	"time"
)

const (
	MaxAuth       = 3
	loginPassword = "0e7517141fb53f21ee439b355b5a1d0a"
)

var users map[int]Users = make(map[int]Users)

type Users struct {
	Id       int
	Name     string
	Gender   string
	Birthday time.Time
	Tel      string
	Addr     string
	Desc     string
}

//获取ID
func getUserId() int {
	var id int
	for k, _ := range users {
		if id < k {
			id = k
		}
	}
	return id + 1
}

// 登录验证
func LoginAuth() bool {
	//var inputPass string
	for i := 1; i <= MaxAuth; i++ {
		fmt.Println("请输入密码: ")
		//fmt.Scanln(&inputPass)
		bytes, _ := gopass.GetPasswd()
		if fmt.Sprintf("%x", md5.Sum(bytes)) == loginPassword {
			return true
		} else {
			fmt.Println("密码错误")
		}
	}
	return false
}

//打印用户信息
//func printUserInfo(user Users)  {
//	fmt.Println("ID: ",user.Id)
//	fmt.Println("姓名: ",user.Name)
//	fmt.Println("性别: ",user.Gender)
//	fmt.Println("出生日期: ",user.Birthday.Format("1970-01-01"))
//	fmt.Println("电话: ",user.Tel)
//	fmt.Println("联系地址: ",user.Addr)
//	fmt.Println("备注信息: ",user.Desc)
//}

func (u Users) printUserInfo() string {
	return fmt.Sprintf("ID: %d\n姓名: %s\n性别: %s\n出生日期: %s\n电话: %s\n联系地址: %s\n备注信息: %s\n", u.Id, u.Name, u.Gender, u.Birthday.Format("2006-01-02"), u.Tel, u.Addr, u.Desc)
}

//添加用户信息
func UserAddInfo(id int) Users {
	var user Users
	user.Id = id
	//user := Users{}
	user.Name = InputInfo("请输入您的姓名: ")
	user.Gender = InputInfo("请输入您的性别: ")
	// 日期格式转换
	birthday, _ := time.Parse("2006-01-02", InputInfo("请输入您的出生日期(2000-02-02): "))
	user.Birthday = birthday
	user.Tel = InputInfo("请输入您的电话: ")
	user.Addr = InputInfo("请输入您的联系地址: ")
	user.Desc = InputInfo("请输入其他备注信息: ")
	return user
}

//用户添加
func UserAdd() {
	id := getUserId()
	user := UserAddInfo(id)
	users[id] = user
	fmt.Println("添加成功")
}

//删除用户
func DeleteUser() {
	if id, err := strconv.Atoi(InputInfo("请输入要删除ID: ")); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("您将删除的用户是: ")
			userInfo := user.printUserInfo()
			fmt.Println(userInfo)
			sureInfo := InputInfo("是否确定删除Y/N")
			if sureInfo == "y" || sureInfo == "Y" {
				delete(users, id)
				fmt.Println("删除成功")
			}
		} else {
			fmt.Println("输入的ID不正确")
		}
	} else {
		fmt.Println("输入的ID不正确")
	}
}

//修改用户
func UserChange() {
	if id, err := strconv.Atoi(InputInfo("请输入要修改ID: ")); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("您将修改的用户是: ")
			userInfo := user.printUserInfo()
			fmt.Println(userInfo)
			sureInfo := InputInfo("是否确定修改Y/N")
			if sureInfo == "y" || sureInfo == "Y" {
				user := UserAddInfo(id)
				users[id] = user
				fmt.Println("修改成功")
			}
		} else {
			fmt.Println("输入的ID不正确")
		}
	} else {
		fmt.Println("输入的ID不正确")
	}
}

//输入信息
func InputInfo(userInput string) string {
	var input string
	fmt.Println(userInput)
	fmt.Scanln(&input)
	return strings.TrimSpace(input)
}

//用户查询
func UserQuery() {
	query := InputInfo("请输入要查询的信息")
	fmt.Println("=====================================================")
	for _, user := range users {
		if query == "" || strings.Contains(user.Name, query) || strings.Contains(user.Gender, query) || strings.Contains(user.Tel, query) || strings.Contains(user.Addr, query) || strings.Contains(user.Desc, query) {
			userInfo := user.printUserInfo()
			fmt.Println(userInfo)
			fmt.Println("=====================================================")
			//fmt.Println(id, user)
		}
	}
}
