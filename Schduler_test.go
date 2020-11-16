package scheduler

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

const (
	Interval uint = 30
	Offset   uint = 0
)

func TestCreateAndGo(t *testing.T) {

	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		s, n := NewScheduler(Interval, Offset)
		fmt.Printf("next-> %v\n", n)
		wg.Add(1)
		go func() {
			i := 0
		timeticker:
			for {
				select {
				case <-s.C:
					s.Stop() // stop before hevy something.
					f()
					if i >= 100 {
						fmt.Println("call close.")
						close(s.Done)
						fmt.Println("closed")
					}
					i++
					n = s.Reset() // restart scheduler
					fmt.Printf("next-> %v\n", n)
				case <-s.Done:
					fmt.Println("break loop.")
					wg.Done()
					break timeticker
				}
			}
		}()
	}
	wg.Wait()
}

func f() error {
	fmt.Printf("tick=%v\n", time.Now())
	return nil
}
