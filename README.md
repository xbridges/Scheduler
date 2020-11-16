# scheduler
Scheduler to start on time.

## usage
  scheduler, nexttime := NewScheduler(Interval, Offset)
  
  break timeticker:
  go func(){
      for {
          select {
          case <-s.C:
              s.Stop() // stop before hevy something.
              func() {
                  fmt.Printf("tick=%v\n", time.Now())
              }
              n = s.Reset() // restart scheduler
              fmt.Printf("next-> %v\n", n)
          case <-s.Done:
              fmt.Println("break loop.")
              wg.Done()
              break timeticker
          }
      }
  }()
