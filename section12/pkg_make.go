//사용자 패키지 작성 및 문서화 예제
package main

import (
	"fmt"
	oper "gostudy/section12/arithmetic" //alias 사용
)

func main() {
	//사용자 패키지 작성 & Go 문서화
	//기준은 GOPATH/src
	//패키지 폴더명(디텍토리명)
	//패키지 파일의 pacakge 이름으로 사용한다.
	//package main을 제외하고 package 문서에 등록
	//지금까지 우리는 패키지를 사용해 왔다.
	//import 할 때 goroot의 pkg를 보고
	//다음에는 gopath의 pkg를 찾음!
	//기존적으로 GOROOT 의 패키지(pkg) 검색 -> GOPATH의 패키지(src/pkg) 검색
	//go install 명령어 실행의 경우 -> GOPATH/pkg에 등록(문서화)
	//godoc -http=:6060(임의의 포트) -> pkg 이동 -> 본인 패키지 메소드 및 주석 확인(패키지, 타입, 메소드) 주석

	//패키지 사용 예제(사칙연산)

	nums := oper.Numbers{100, 10}

	fmt.Println("Package Used(1) : ", nums.Plus())
	fmt.Println("Package Used(2) : ", nums.Minus())
	fmt.Println("Package Used(3) : ", nums.Multi())
	fmt.Println("Package Used(4) : ", nums.Divide())
	fmt.Println("Package Used(5) : ", nums.SquareMinus())
	fmt.Println("Package Used(6) : ", nums.SquarePlus())

	nums.SquareMinus()
}
