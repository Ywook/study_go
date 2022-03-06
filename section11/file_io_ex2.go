package main

import (
	"bufio"
	"fmt"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//파일 읽기, 버퍼 사용 쓰기 ->  bufio 패키지 활용
	//iotuil, bufio 등은 io.Reader, io.Writer 인터페이스를 구현 함 -> 즉, Writer, Read 메소드를 사용 가능
	/*
		type Reader interface{
			Read(p []byte) (n int, err error)
		}

		type Writer interface {
			Wirte(p []byte) (n int, err error)
		}
	*/

	//즉, bufio의 NewReader, NewWriter를 통해서 객체 반환 후 메소드 사용 가능

	//bufio(Buffered io) 패키지 보는게 좋다!
	//https://golang.org/pkg/bufio
	//https://golang.org/pkg/os/#pkg-constants

	/*
		상태
		a -----> b
		b -----> ab
		c -----> abc
		d -----> abcd
		e -----> e		----->abcd
		f -----> ef		----->abcd
		g -----> efg 	----->abcd
		h -----> efgh	----->abcdefgh
	*/

	file, err := os.OpenFile("test_write2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))

	//bufio 파일 쓰기 예제
	wt := bufio.NewWriter(file) //Writer 반환(버퍼 사용)
	wt.WriteString("Hello Golang\nFile Write Test!\n")
	wt.WriteString("Hello Golang\nFile Write Test!\n")
	//에러체크
	errCheck(err)
	fmt.Printf("사용한 Buffer Size : (%d byte)\n", wt.Buffered())
	fmt.Printf("남은 Buffer Size : (%d byte)\n", wt.Available())
	fmt.Printf("전체 Buffer Size : (%d byte)\n", wt.Size())

	wt.Flush() // 버퍼 비우고 반영 ( 버퍼에 내용을 디스크에 기록 )
	fmt.Println("쓰기 작업 완료\n")
	fmt.Println("==========================================")

	rt := bufio.NewReader(file) // Reader 반환
	fi, err := file.Stat()
	errCheck(err)

	b := make([]byte, fi.Size())

	fmt.Println("파일 정보 출력 : ", fi)
	fmt.Println("파일 이름 : ", fi.Name())
	fmt.Println("파일 크기 : ", fi.Size())
	fmt.Println("파일 수정 시간 : ", fi.ModTime())
	fmt.Println("==========================================")

	file.Seek(0, os.SEEK_SET)
	data, _ := rt.Read(b) //읽기(ReadLine, ReadByte, ReadBytes 등)

	fmt.Printf("전체 Buffer Size : (%d byte)\n", rt.Size())
	fmt.Printf("읽기 작업 완료 : (%d byte)\n", data)

	defer file.Close()
}
