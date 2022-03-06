package main

import "fmt"

func main() {
	//배열 순회

	//예제1
	arr1 := [5]int{1, 10, 100, 1000, 10000}

	//len 길이 반복
	for i := 0; i < len(arr1); i++ {
		fmt.Println("ex1 : ", arr1[i])
	}

	for i, val := range arr1 {
		fmt.Println(i, val)
	}

	for _, val := range arr1 {
		fmt.Println(val)
	}

	//인덱스 생략2
	for v := range arr1 {
		fmt.Println("ex4 : ", v)
	}

}
