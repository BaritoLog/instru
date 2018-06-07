/*
Simple go instrumentation library for flexible push-pull strategy.

1. Instrumenting your code

1.1 Evaluation time

  func myFunction()  {
    eval := instru.Evaluate("myFunction")
    defer eval.Done()

    // some process
  }


Sample of Evaluation Time Metric (in nano seconds)
  {
    "myFunction": {
      "_evaluation_time": {
        "count": 12,
        "avg": 5000,
        "max": 10000,
        "min": 1000,
        "recent": 1000
      }
    }
  }


1.2 Counter

  func myFunction()  {
  	if rand.Int31()%2 == 0 {
  		instru.Counter("myFunction").Event("odd")
  	} else {
  		instru.Counter("myFunction").Event("even")
  	}
  }


Sample of Counter metric
  {
    "myFunction": {
      "_counter": {
        "total": 21,
        "events": {
          "odd": 9,
          "even": 12
        }
      }
    }
  }


1.3 Custom Instrumenting

  func myFunction()  {
    // get fileinfo of my.conf
    _, err := os.Stat("my.conf")

    instru.Metric("myFunc").Put("is_config_exist", os.IsNotExist(err))
  }


Some of metric
  {
    "myFunction": {
      "is_config_exist": true
    }
  }


2. Pull the Instrumentation Metric

2.1 With RESTful API

Expose with RESTful API
  func main()  {
    instru.ExposeWithRestful(":8998")
  }



2.2 With Custom Exposer


  type CustomExposer struct{
  }

  // Expose is required
  func (e *CustomExposer)Expose(instr intru.Instrumenation) (err error)  {
    // TODO: expose implementation
    return
  }

  // Stop is required
  func (e *CustomExposer) Stop()  {
    // TODO: stop implementation
  }

  func main()  {
    exposer := &CustomExposer{}
    instru.Expose(exposer)
  }


3. Push the Instrumentation Metric

3.1 Web Callback

  func main()  {
    interval, _ := time.ParseDuration("5m")
    SetWebCallback(interval, "http://my-server:3042")
  }


3.2 Custom Callback

  type MyCallback struct{
  }

  func (c *MyCallback) OnCallback(instr Instrumentation) (err error) {
    // TODO: OnCallback implementation
    return
  }

  func main()  {
    interval, _ := time.ParseDuration("5m")
    instru.SetCallback(interval, &MyCallback{})
  }
*/
package instru
