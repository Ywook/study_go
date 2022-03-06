package main

import "fmt"

func main() {
	//배열
	//배열은 용량, 길이가 항상 같다.
	//배열 vs 슬라이스 차이점 중요
	//길이 고정 vs 길이 가변
	//값 타임 vs 참조 타입
	//복사 전달, vs 참조 값 전달 (메모리 주소 값)
	//전체 비교연산자 사용 가능 vs 비교 연산자 사용불가
	//대부분 슬라이스 사용함.

	//cap() : 배열, 슬라이스 용량
	//len() : 배열 슬라이스 개수

	var arr1 [5]int
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	var arr3 = [5]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4, 5}
	arr5 := [5]int{1, 2, 3}
	arr6 := [...]int{1, 2, 3, 4, 5} //배열크기 자동 맞춤
	arr7 := [5][5]int{
		{1, 2, 3, 4, 5},
		{},
		{},
		{},
		{},
	}
	fmt.Println(arr1)
	fmt.Println(arr2)

	fmt.Printf(" %d %v\n", arr1, len(arr1), arr1)
	fmt.Printf("%-5T %d %v\n", arr2, len(arr1), arr2)
	fmt.Printf("%-5T %d %v\n", arr3, len(arr1), arr3)
	fmt.Printf("%-5T %d %v\n", arr4, len(arr1), arr4)
	fmt.Printf("%-5T %d %v\n", arr5, len(arr1), arr5)
	fmt.Printf("%-5T %d %v\n", arr6, len(arr1), arr6)
	fmt.Printf("%-5T %d %v\n", arr7, len(arr1), arr7)

	arr8 := [5]int{1, 2, 3, 4, 5}
	arr9 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}

	arr10 := [...]string{"kim", "lee", "park"}

	fmt.Printf("%-5T %d %v\n", arr8, len(arr8), arr8)
	fmt.Printf("%-5T %d %v\n", arr9, len(arr9), arr9)
	fmt.Printf("%-5T %d %v\n", arr10, len(arr10), arr10)

}