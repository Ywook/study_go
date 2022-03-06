//구조체 심화(3)

package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) CalculateD(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a *Account) CalculateF(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

//일반 함수로 전달할 때는 포인터를 받을 때 제대로 &를 써서해야함
//하지만 리시버를 사용하면 굳이 그렇게 안해도됨 ~!
func main() {
	kim := Account{"245-901", 10000000, 0.015}
	lee := Account{"245-901", 20000000, 0.025}

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)

	kim.CalculateD(10000000)
	lee.CalculateF(10000000)

	fmt.Println("ex2 : ", int(kim.balance))
	fmt.Println("ex2 : ", int(lee.balance))
}
