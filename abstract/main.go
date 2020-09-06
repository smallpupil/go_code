package main

import "fmt"

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

//1.存款

func (account *Account) Deposite(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密码错误")
		return
	}
	if money <= 0 {
		fmt.Println("金额错误")
		return
	}
	account.Balance += money
	fmt.Println("存款成功~~")
}

//取款

func (account *Account) WithDraw(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密码错误")
		return
	}
	if money <= 0 || money > account.Balance {
		fmt.Println("金额错误")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功~~")
}

//余额
func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密码错误")
		return
	}
	fmt.Println("账号余额:", account.Balance)
}

func main() {
	account := Account{"gs1111", "666666", 1000.0}
	account.Query("666666")
}
