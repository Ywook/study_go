package main

import "fmt"

func main() {
	//슬라이스 추가 및 병합
	//예제 1
	/*
		용량을 초과하는 것 만큼 추가 될 경우
		기존 용량 * 2 의 신규 배열을 만들고 데이터를 할당하게 된다.
		그래서 용량은 예측을 하고 넣어주는 게 좋고 append는 신중하게 사용해야함.
	*/

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{8, 9, 10, 11, 12}
	s3 := []int{13, 14, 15, 16, 17}

	s1 = append(s1, 6, 7)
	s2 = append(s1, s2...)      //슬라이스를 삽입할 경우 사용
	s3 = append(s2, s3[0:3]...) // 추출 후 병합

	fmt.Println("ex1  : ", s1, cap(s1))
	fmt.Println("ex1  : ", s2, cap(s2))
	fmt.Println("ex1  : ", s3, cap(s3))

	s4 := make([]int, 0, 5)
	for i := 0; i < 15; i++ {
		s4 = append(s4, i)
		fmt.Printf("ex2 -> len : %d, cap : %d, %v\n", len(s4), cap(s4), s4)
	}
}
