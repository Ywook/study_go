//사용자 정의 타입(4)
package main

import "fmt"

func abcd(a *int) {
	*a = 77
}

type shoppingBasket struct {
	cnt   int
	price int
}

func (b shoppingBasket) purchase() int {
	return b.cnt * b.price
}

//원본이 수정 (참조 형식)
func (b *shoppingBasket) rePurchaseP(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

//원본 수정X (값 전달 형식)
func (b shoppingBasket) rePurchaseD(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

func main() {
	//구조체는 역참조, 참조 값 자동으로 진행됨
	//리시버(구조체 메소드) 전달(값, 참조) 형식
	//함수는 기본적으로 값 호출 -> 변수의 값이 복사 후 내부 전달(원본 수정 X) -> 맵, 슬라이스 등은 참조 전달
	//리시버(구조체)도 마찬가지로 포인터를 활용해서 메소드 내에서 원본 수정이 가능함

	//예제1
	bs1 := shoppingBasket{3, 5000}
	fmt.Println("ex1(totPirce) : ", bs1.purchase())

	//참조 전달 (원본 값 수정)
	bs1.rePurchaseP(7, 5000)
	fmt.Println("ex1(totPirce) : ", bs1.purchase())

	//값 전달 (원본 수정 X)
	bs1.rePurchaseD(10, 0)
	fmt.Println("ex1(totPirce) : ", bs1.purchase())
}
