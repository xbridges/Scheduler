# scheduler
  Scheduler to start on time.
  A scheduler that runs at "0:00,  0:30... minutes per hour" required for resident programs.

## usage
  ```
  scheduler, nexttime := NewScheduler(10, 0)
  
  wg := sync.WaitGroup{}
  wg.Add(1)
  go func(){
      break timeticker:
      for {
      i := 0
          select {
          case <-scheduler.C:
              scheduler.Stop() // stop before hevy something.
              func() {
                  fmt.Printf("tick=%v\n", time.Now())
              }
              if i < 100 {
                  n = scheduler.Reset() // restart scheduler
              } else {
                  scheduler.Close()
              }
              fmt.Printf("next-> %v\n", n)
          case <-scheduler.Done:
              fmt.Println("break loop.")
              wg.Done()
              break timeticker
          }
      }
  }()
  wg.Wait()
  
  ```