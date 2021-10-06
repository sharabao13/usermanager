package main

import (
	"fmt"
	upkg "github.com/sharabao13/usermanager/users"
	"os"
)

func main() {
	if !upkg.LoginAuth() {
		fmt.Printf("密码%d次错误，程序退出\n", upkg.MaxAuth)
		return
	}
	menu := `
*********************************
  1. 新建用户
  2. 修改用户
  3. 删除用户
  4. 查询用户
  5. 退出
*********************************
`
	//users := make(map[int]map[string]string)
	callBacks := map[string]func(){
		"1": upkg.UserAdd,
		"2": upkg.UserChange,
		"3": upkg.DeleteUser,
		"4": upkg.UserQuery,
		"5": func() {
			os.Exit(0)
		},
	}
	fmt.Println("欢迎进入xxx用户管理系统")
	//END:
	for {
		fmt.Println(menu)
		oper := upkg.InputInfo("请输入指令")
		callback, ok := callBacks[oper]
		if ok {
			callback()
		} else {
			fmt.Println("指令错误")
		}
	}
}
