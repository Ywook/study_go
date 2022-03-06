package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	//함수 고급
	//재귀 함수(Recursuion)
	//프로그램이 보기 쉽고, 코드 간결, 오류 수정이 용이 : 장점
	//코드 이해하기 어렵고, 기억공간을 많이 사용, 무한루프 사용 가능성

	//예제1
	x := fact(7)
	fmt.Println("ex1 : ", x)
}
