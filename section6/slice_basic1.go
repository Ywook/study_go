package main

import "fmt"

func main() {
	//길이 x (가변) -> 동적으로 크기가 늘어난다. 레퍼런스(참조 값 타입)
	//슬라이느 (길이 & 용량) 크기가 동적으로 할당 가능

	//2가지 선언 방법ㅂ 1. 배열처럼 선언, 2. make 함수 : make(자료형, 길이, 용량(생략시 길이))

	var slice1 []int //할당 크기를 적용하지 않으면 slice 로 생성된다.
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	slice3[4] = 10

	fmt.Printf("%-5T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("%-5T %d %d %v\n", slice1, len(slice2), cap(slice2), slice2)
	fmt.Printf("%-5T %d %d %v\n", slice1, len(slice3), cap(slice3), slice3)
	fmt.Printf("%-5T %d %d %v\n", slice1, len(slice4), cap(slice4), slice4)

	// 예제2
	// 용량 초과가 날 경우 메모리를 재할당함 그래서 크기가 큰 경우 성능 저하가 발생할 수 있음
	var slice5 []int = make([]int, 5, 10)
	var slice6 = make([]int, 5)
	slice7 := make([]int, 5, 100)
	slice8 := make([]int, 5)

	slice6[2] = 7
	fmt.Println()
	fmt.Printf("%-5T %d %d %v\n", slice5, len(slice1), cap(slice5), slice5)
	fmt.Printf("%-5T %d %d %v\n", slice6, len(slice2), cap(slice6), slice6)
	fmt.Printf("%-5T %d %d %v\n", slice7, len(slice3), cap(slice7), slice7)
	fmt.Printf("%-5T %d %d %v\n", slice8, len(slice4), cap(slice8), slice8)

	// 예제3
	var slice9 []int //nil 슬라이스(길이와 용량 0)

	if slice9 == nil {
		fmt.Println("This is Nil !")
	}

}
