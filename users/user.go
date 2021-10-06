package users

import (
	"crypto/md5"
	"fmt"
	"github.com/howeyc/gopass"
	"strconv"
	"strings"
)

const (
	MaxAuth       = 3
	loginPassword = "0e7517141fb53f21ee439b355b5a1d0a"
)

var users map[int]map[string]string = make(map[int]map[string]string)

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

//添加用户信息
func UserAddInfo() map[string]string {
	user := map[string]string{}
	user["name"] = InputInfo("请输入您的姓名: ")
	user["age"] = InputInfo("请输入您的年龄: ")
	user["tel"] = InputInfo("请输入您的电话: ")
	user["addr"] = InputInfo("请输入您的住址: ")
	return user
}

//用户添加
func UserAdd() {
	id := getUserId()
	user := UserAddInfo()
	users[id] = user
	fmt.Println("添加成功")
}

//删除用户
func DeleteUser() {
	if id, err := strconv.Atoi(InputInfo("请输入要删除ID: ")); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("您将删除的用户是: ")
			fmt.Println(user)
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
			fmt.Println(user)
			sureInfo := InputInfo("是否确定修改Y/N")
			if sureInfo == "y" || sureInfo == "Y" {
				user := UserAddInfo()
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
	titel := fmt.Sprintf("%3s|%10s|%3s|%10s|%20s", "Id", "Name", "AGE", "Tel", "Addr")
	fmt.Println(titel)
	fmt.Println(strings.Repeat("-", len(titel)))
	for id, user := range users {
		if query == "" || strings.Contains(user["name"], query) || strings.Contains(user["age"], query) || strings.Contains(user["tel"], query) || strings.Contains(user["addr"], query) {
			fmt.Printf("%3d|%10s|%3s|%10s|%20s", id, user["name"], user["age"], user["tel"], user["addr"])
			//fmt.Println(id, user)
		}
	}
}
