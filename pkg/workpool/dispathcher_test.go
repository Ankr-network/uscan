package workpool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Student struct {
	wg *sync.WaitGroup
	id int
}

func (s *Student) Execute() {
	time.Sleep(time.Second * 1)
	fmt.Printf("%d doing...\n", s.id)
	s.wg.Done()
}

func TestStartDispathcher(t *testing.T) {
	start := time.Now()
	var d1 Dispathcher = NewDispathcher(2)
	var d2 Dispathcher = NewDispathcher(2)
	var wg sync.WaitGroup
	for i := 1; i < 9; i++ {
		s := &Student{
			id: i,
			wg: &wg,
		}
		wg.Add(1)
		if i%2 == 0 {
			d1.AddJob(s)
		} else {
			d2.AddJob(s)
		}
	}

	wg.Wait()
	d1.Stop()
	d2.Stop()

	sub := time.Since(start).Seconds()
	t.Log(sub)
}
