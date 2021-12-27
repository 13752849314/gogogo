package PV

import (
	"fmt"
	"sync"
	"time"
)

var mutex = sync.Mutex{}
var signal = sync.WaitGroup{}
var tickets = 100

func soleTickets(n int) {
	for {
		mutex.Lock()
		if tickets == 0 {
			fmt.Printf("%d号窗口票卖完了!\n", n)
			mutex.Unlock()
			break
		}
		fmt.Printf("%d窗口正在买票,剩余票数:%d\n", n, tickets)
		tickets--
		time.Sleep(100)
		mutex.Unlock()
	}
	signal.Done()
}

func Test() {
	signal.Add(4)
	for i := 0; i < 4; i++ {
		go soleTickets(i)
	}
	signal.Wait()
	fmt.Println("买票结束!")
}
