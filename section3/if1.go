package main

import "fmt"

func main() {
	//IF 문 : 반드시 Boolean 검사 -> 1, 0 으로 안됨 ~!
	// 소괄호 사용 X

	var a int = 20
	b := 20

	if a >= 15 {
		fmt.Println("over 15")
	}

	if b >= 25 {
		fmt.Println("over 25")
	}

	//에러 발생1
	/*
		if b >= 25
		{

		}
	*/

	//에러 발생2
	/*
		if b >= 25
			fmt.Println("over 25")
	*/

	// 에러 발생3
	/*
		if c := 1; c {

		}
	*/

	if c := 40; c >= 35 {
		fmt.Println("over 35")
	}

	// c += 1 에러발생

}
