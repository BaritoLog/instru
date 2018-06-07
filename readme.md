# Instru

Simple go [instrumentation](https://en.wikipedia.org/wiki/Instrumentation_%28computer_programming%29) library for flexible [push-pull strategy](https://en.wikipedia.org/wiki/Push%E2%80%93pull_strategy).

Find us at [godoc](https://godoc.org/github.com/BaritoLog/instru)

## 1. Instrumenting your code

#### 1.1 Evaluation time

```go
func myFunction()  {
  eval := instru.Evaluate("myFunction")
  defer eval.Done()
  
  // some process
}
```

Sample of Evaluation Time Metric (in nano seconds)
```js
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
```

#### 1.2 Counter

```go
func myFunction()  {
	if rand.Int31()%2 == 0 {
		instru.Counter("myFunction").Event("odd")
	} else {
		instru.Counter("myFunction").Event("even")
	}
}
```

Sample of Counter metric 
```js
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
```

### 1.3 Custom Instrumenting

```go
func myFunction()  {
  // get fileinfo of my.conf
  _, err := os.Stat("my.conf")
  
  instru.Metric("myFunc").Put("is_config_exist", os.IsNotExist(err))
}
```

Some of metric
```js
{
  "myFunction": {
    "is_config_exist": true
  }
}
```



## 2. Pull the Instrumentation Metric 

#### 2.1 With RESTful API

Expose with RESTful API
```go
func main()  {
  instru.ExposeWithRestful(":8998")
}
```

Retrieve the instrumentation metric
```sh
curl http://localhost:8998
```

Sample of Instrumentation metric
```js
{
  "myFunc": {
    "_evaluation_time": {
      "count": 12,
      "avg": 5000,
      "max": 10000,
      "min": 1000,
      "recent": 1000
    },
    "_counter": {
      "total": 21,
      "events": {
        "odd": 9,
        "even": 12
      }
    },
    "is_config_exist": true
  }
}
```

#### 2.2 With Custom Exposer

```go
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
```
```go
func main()  {
  exposer := &CustomExposer{}
  instru.Expose(exposer)
}
```

## 3. Push the Instrumentation Metric 

#### 3.1 Web Callback
```go
func main()  {
  interval, _ := time.ParseDuration("5m")
  SetWebCallback(interval, "http://my-server:3042")
}
```

#### 3.2 Custom Callback

```go
type MyCallback struct{
}

func (c *MyCallback) OnCallback(instr Instrumentation) (err error) {
  // TODO: OnCallback implementation
  return
}
```
```go
func main()  {
  interval, _ := time.ParseDuration("5m")
  instru.SetCallback(interval, &MyCallback{})  
}
```
