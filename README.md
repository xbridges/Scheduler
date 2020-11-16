# scheduler
Scheduler to start on time.
A scheduler that runs at "0:00,  0:30... minutes per hour" required for resident programs.

## usage
  ```
  scheduler, nexttime := NewScheduler(Interval, Offset)
  
  go func(){
      break timeticker:
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
  ```