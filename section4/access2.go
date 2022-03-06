package main

import (
	"fmt"
	checkUp "gostudy/section4/lib"
	_ "gostudy/section4/lib2" //빈 식별자 사용을 통해 임시 import
)

func main() {
	//패키지 접근제어
	//별칭 사용
	//빈 식별자 사용

	fmt.Println("10보다 큰 수 ? : ", checkUp.CheckNum(11))
}
