package main

import (
	"fmt"
	"gostudy/section4/lib2"
)

func main() {
	//패키지 접근제어
	//변수, 상수, 함수, 메서드, 구조체 등 식별자
	//대문자 : 패키지 외부에서 접근 가능
	//소문자 : 패키지 외부에서 접근 불가

	fmt.Println("100보다 큰 수 ? : ", lib2.CheckNum1(103))
	fmt.Println("1000보다 큰 수 ? : ", lib2.CheckNum2(1004))
}