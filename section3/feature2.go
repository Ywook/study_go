package main

import "fmt"

func main() {
	//문장 끝 세미콜론(;) 주의
	// Go 컴파일러가 자동으로 문장 끝에 세미콜론을 삽입한다
	// 두 문장을 한 문장에 표현할 경우 명시적으로 세미콜론 사용(권장 x)
	//반복문 및 제어문(조건문) 사용할 경우 주의

	//예제 1

	for i := 0; i <= 10; i++ {
		fmt.Println("ex1 : ", i)
		fmt.Println(i)
	}

	//예제2
	for j := 10; j >= 0; j-- {
		fmt.Println("ex2 : ", j)
	}
}
