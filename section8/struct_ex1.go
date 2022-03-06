//구조체 심화(1)

package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func NewAccount(number string, balance float64, interest float64) *Account { //포인터 반환 아닌 경우 값 복사
	return &Account{number, balance, interest} //구조체 인스턴스를 생성한 뒤 포인터 리턴
}

func main() {

	kim := Account{number: "245-901", balance: 10000000, interest: 0.015}

	var lee *Account = new(Account)
	lee.number = "245-902" //getter, setter
	lee.balance = 13000000
	lee.interest = 0.025

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)

	park := NewAccount("245-903", 17000000, 0.04)
	fmt.Println("ex2 : ", park)
}
