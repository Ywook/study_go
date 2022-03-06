//구조체 기본(1)
package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	//다양한 선언 방법
	//&struct, struct : &sturct 포인터를 받아오고 역참조를 또 하기 때문에 속도는 조금 느리다.
	//인터페이스 메소드를 선언만 해둔 후 -> 오버라이딩 해서 메소드에 포인터 리시버를 사용할 경우 반드시 &struct로 넘겨야한다.

	//선언 방법1
	var kim *Account = new(Account) //new를 쓸때는 별표를 붙여야함!
	kim.number = "245-901"
	kim.balance = 10000000
	kim.interest = 0.015

	//선언 방법2
	hong := &Account{number: "245-902", balance: 15000000, interest: 0.04}

	//선언 방법3
	lee := new(Account) //new 에서는 안에 값 설정 불가능
	lee.number = "245-903"
	lee.balance = 130000000
	lee.interest = 0.025

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", hong)
	fmt.Println("ex1 : ", lee)

	fmt.Printf("ex2 : %#v\n", kim)
	fmt.Printf("ex2 : %#v\n", hong)
	fmt.Printf("ex2 : %#v\n", lee)

	fmt.Println()
	//예제2
	fmt.Println("ex3 : ", kim.Calculate())
	fmt.Println("ex3 : ", hong.Calculate())
	fmt.Println("ex3 : ", lee.Calculate())
}
