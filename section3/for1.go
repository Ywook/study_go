package main

import "fmt"

func main() {
	//반복문 - For
	//Go에서 유일하게 제공하는 반복분

	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", i)
	}
	/*
		//에러 발생 1

		for i := 0; i < 5; i++
		{

		}

		//에러 발생 2
		for i:= 0; i < 5; i++
			fmt.Println("ex 1 " , i)
	*/
	/*  무한루프 패턴
	for {
		fmt.Println("ex2 : Helllo, Golang!")
		fmt.Println("ex2 : Infinite loop!")
	}
	*/

	loc := []string{"Seoul", "Busan", "Incheon"}

	for index, name := range loc {
		fmt.Println("ex3 : ", index, name)
	}
}
