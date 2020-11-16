package scheduler

import (
	"time"
)

// Scheduler struct
type Scheduler struct {
	*time.Ticker
	Done     chan struct{}
	Interval uint // value unit is second
	Offset   uint // value unit is second
}

/*
 function name: getNextSchedule
 getNextSchedule provide next timer schedule interval from now and execution time.
 This function inner call only.
*/
func getNextSchedule(interval uint, offset uint) (time.Duration, time.Time) {

	is := time.Duration(interval) * time.Second // interval value convert to second
	os := time.Duration(offset) * time.Second   // offset value convert to second
	n := time.Now()                             // current time

	// ticker values may fluctuate few (+/-)milli seconds.
	ns := n.Round(is) // ex) interval 60: tick=59.99 or 60.20 -> 60
	ns = ns.Add(is)   //     60 + interval(60) = 120(60s)

	adj := ns.Sub(n) + os // (next - now) + offses => 60(59.9 or 60.1 ....) + offset
	ns = ns.Add(os)       // next fire time.
	// fmt.Printf("now: %v, next: %v, ajusted: %v\n", n, ns, adj)

	// return adjusted interval and next time of fire.
	return adj, ns
}

/*
 function name: NewScheduler
 This is the schedule struct factory.
 The schedule struct use as like time.ticker by embed time.ticker.
 Feature: Set interval time(second) the shcedule struct fired at each hour just 00 second.
  # Interval: 3,600 offset:   0 -> fired 00:00:00, 01:00:00, 02:00:00, ..., 23:00:00 (each an hour)
  # Interval: 3,600 offset: 300 -> fired 00:05:00, 01:05:00, 02:05:00, ..., 23:05:00 (each an hour 5 minute)
  # Interval:   600 offset:   0 -> fired 00:10:00, 00:20:00, 00:30:00, ..., 23:50:00 (each 5 minute)
  # Interval:    10 offset:   0 -> fired 00:00:00, 00:00:10, 00:00:20, ..., 23:50:50 (each 10 second)
*/
func NewScheduler(interval uint, offset uint) (*Scheduler, time.Time) {

	// Fire at just interval.
	// ex)each hour 00 minute, each 10 minutes 05 seconds,...
	n, ns := getNextSchedule(interval, offset)
	t := &Scheduler{Interval: interval, Offset: offset, Ticker: time.NewTicker(n), Done: make(chan struct{})}
	return t, ns
}

/*
 function name: Reset
 Reset function need to set next time after fire.
!!! Must be called with Stop() !!!
*/
func (schduler *Scheduler) Reset() time.Time {
	n, ns := getNextSchedule(schduler.Interval, schduler.Offset)
	schduler.Ticker = time.NewTicker(n)
	return ns
}

/*
 function name: Close
 Close have to be called by teminate sequence.
!!! Do not to close Done channel directly from outside. !!!
*/
func (schduler *Scheduler) Close() {
	close(schduler.Done)
	schduler.Stop()
}
