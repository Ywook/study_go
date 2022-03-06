package main

//고루틴 동기화 기초(1)
import (
	"fmt"
	"runtime"
	"sync"
)

//구조체 선언(공유 데이터)

type test struct {
	num   int
	mutex sync.Mutex
}

func (c *test) increment() {
	//공유 데이터 수정 전 뮤텍스로 보호
	c.mutex.Lock()
	c.num += 1
	//공유 데이터 수정 후 보호 해제
	c.mutex.Unlock()
}

func (c *test) result() {
	fmt.Println(c.num)
}

func main() {
	//고루틴 동기화 예제
	//실행 흐름 제어 및 변수 동기화 가능
	//공유 데이터 보호가 가장 중요
	//뮤텍스(Mutex)

	//동기화 사용하지 않은 경우 예제
	//시스템 전체 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := test{num: 0}
	done := make(chan bool)

	for i := 1; i <= 10000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched() // CPU 양보
		}()
	}

	for i := 1; i <= 10000; i++ {
		<-done
	}
	c.result()
}
