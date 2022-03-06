//go 초기화 함수 (1)

package main

import (
	"fmt"
)

//main 함수보다 먼저 딱 1번 실행되는 함수
func init() {
	fmt.Println("Init Method Start!")
}

func init() {
	fmt.Println("Init Method Start2!")
}

func init() {
	fmt.Println("Init Method Start3!")
}

func init() {
	fmt.Println("Init Method Start4!")
}

func init() {
	fmt.Println("Init Method Start5!")
}

func main() {
	//init : 패키지 로드시에 가장 먼저 호출
	//가장 먼저 초기화 되는 작업 작성시 유용하다.
	//main 패키지가 아닌 곳에서 init메소드를 사용할 경우 import 시 init 메서드가 호출된다.
	//init 메서드가 여러개 있을 경우 순차적으로 위에서부터 실행된다.

}
