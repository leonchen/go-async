# go-async
async helpers for golang

## Install

````
go get github.com/leonchen/go-async/...
````

## Usage
````
import (
  "github.com/leonchen/go-async/async"
  "fmt"
)

var data = []interface{}{"a", "b", "c"}
var each = func(d interface{}) {
  fmt.Println(d.(string))
}
// run in parallel
async.Each(data, each)

// run in parallel with specified limit
async.EachLimit(data, 2, each)

// run in parallel with limit as the number of cpus
async.EachCPU(data, each)

// run in parallel with maximum number of processes
async.EachProc(data, each)

// process each element and return an array of the results
var eachRes = func(d interface{}) interface{} {
  return d+"o"
}
// results is []interface{}
results := async.Map(data, eachRes)

// run n times in parallel
var eachTimes = func(n int) interface{} {
  return n
}
// res is []interface{}
res := async.Times(5, eachTimes)
````