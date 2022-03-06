package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//고루틴 동기화 고급
	//원자성(Atomic) 사용 ->  기능적으로 분할 불가능한 완전 보증된 일련의 조작, 모두 성공하거나 모두 실패
	//모든 조작이 완료 될 떄까지 다른 프로세스 개입 불가
	//sync/atomic에서 원자적 연산자 제공
	//https://golang.org/pkg/sync/atomic 에서 확인 가능
	//주로 공용 변수에 관한 계산 사용

	//원자성 사용 안할 경우 예제
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := new(sync.WaitGroup)

	var cnt int64 = 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(n int) {
			cnt += 1
			fmt.Println("WaitGroup : ", n)
			wg.Done()
		}(i)
	}

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(n int) {
			cnt -= 1
			fmt.Println("WaitGroup : ", n)
			wg.Done()
		}(i)
	}
	//Add == Done 횟수가 같아야함!
	wg.Wait()
	fmt.Println("WaitGroup End", cnt)

}
