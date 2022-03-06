package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	//파일 쓰기
	//csv 파일 쓰기 예제
	//패키지 저장소를 통해서 Excel 등 다양한 파일 형식 쓰기, 읽기 가능
	//패키지 Github 주호 : https://github.com/tealeg/xlsx
	//bufio : 파일이 용량이 클 경우 버퍼 사용 권장

	//파일 생성
	file, _ := os.Create("test_write.csv")

	defer file.Close()

	wr := csv.NewWriter(file)
	//wr := csv.NewWriter(bufio.NewWriter(file)) //버퍼 쓰는걸 권장

	//csv 내용 쓰기
	wr.Write([]string{"Kim", "4.8"})
	wr.Write([]string{"Lee", "5.8"})
	wr.Write([]string{"Park", "4.8"})
	wr.Write([]string{"Cho", "5.8"})
	wr.Write([]string{"Hong", "0.8"})

	wr.Flush() //버퍼 -> 파일로 쓰기

	fi, _ := file.Stat()

	fmt.Printf("CSV 쓰기 작업 후 파일 크기(%d bytes)\n", fi.Size())
	fmt.Println("CSV 파일명 : ", fi.Name())
	fmt.Println("운영체제 파일 권한 : ", fi.Mode())
}
