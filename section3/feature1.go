// Go 특징 (1)
package main

import "fmt"

func main() {
	//Go : 모호하거나, 애매한 문법 금지
	//수치(후위) 연산자 허용 (i++), 전치(전위) 연산자 비허용 (++i) -> x
	//증감연산 반환 값 없음 sum := i++
	//포인터(Pointer 허용, 연산 비허용)

	sum, i := 0, 0

	for i <= 100 {
		// sum += i++ 예외 발생
		sum += i
		i++
		//++i 예외 발생
	}

	fmt.Println("ex1 : ", sum)
}
