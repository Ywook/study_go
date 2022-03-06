//데이터 타입 : 문자열 연산(1)

package main

import "fmt"

func main() {
	//문자열 연산
	//추출, 비교, 조합(결합)

	//예제1
	var str1 string = "Golang"
	var str2 string = "World"

	//slicing을 처리해서 출력하면 문자, 배열 그대로 출력하면 아스키코드값
	fmt.Println("ex1 : ", str1[0:2], str1[0])
	fmt.Println("ex1 : ", str2[3:], str2[0])
	fmt.Println("ex1 : ", str2[:4])
	fmt.Println("ex1 : ", str1[1:3])
}
