//자료형 : 맵(1)

package main

import "fmt"

func main() {
	//map
	//맵 : 해시테이블, 딕셔너리(파이썬), Key-value로 자료 저장
	//레퍼런트 타입
	//비교 연산자 사용 불가능(참조 타입이므로)
	//특징 : 참조타입(Key)로 사용 불가능, 값(Value)으로 모든 타입 사용 가능
	//make 함수 및 축약(리터럴)로 초기화 가능
	//순서 없음.

	//예제1
	var map1 map[string]int = make(map[string]int)
	var map2 = make(map[string]int)

	map3 := make(map[string]int)

	fmt.Println("ex1 : ", map1)
	fmt.Println("ex1 : ", map2)
	fmt.Println("ex1 : ", map3)
	fmt.Println()

	map4 := map[string]int{} //Json 형태
	map4["apple"] = 26
	map4["orange"] = 33
	map4["banana"] = 40

	map5 := map[string]int{
		"apple":  15,
		"banana": 40,
		"oragne": 23,
	}

	map6 := make(map[string]int, 10)
	map6["apple"] = 26
	map6["orange"] = 33
	map6["banana"] = 40

	fmt.Println("ex1 : ", map4, cap(map4))
	fmt.Println("ex1 : ", map5, cap(map5))
	fmt.Println("ex1 : ", map6, cap(map6))
	fmt.Println("ex1 : ", map6["orange"])

	fmt.Println()

}
