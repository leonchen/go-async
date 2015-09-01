package async

import (
	"runtime"
	"sync"
)

func Each(data []interface{}, cb func(interface{})) {
	l := len(data)
	if l < 1 {
		return
	}

	ch := make(chan int, l)
	defer close(ch)
	c := make(chan int, 1)
	defer close(c)

	for i := 0; i < l; i++ {
		ch <- 1
	}

	for _, d := range data {
		d := d
		go func() {
			defer func() {
				<-ch
				if len(ch) == 0 {
					c <- 1
				}
			}()
			cb(d)
		}()
	}
	<-c
}

func EachLimit(data []interface{}, limit int, cb func(interface{})) {
	if len(data) < 1 {
		return
	}
	eachLimit(data, limit, cb)
}

func eachLimit(data []interface{}, limit int, cb func(interface{})) {
	ch := make(chan int, limit)
	defer close(ch)

	var wg sync.WaitGroup
	l := len(data)
	wg.Add(l)

	for _, d := range data {
		d := d
		ch <- 1
		go func() {
			cb(d)
			<-ch
			wg.Done()
		}()
	}
	wg.Wait()
}

func EachCPU(data []interface{}, cb func(interface{})) {
	if len(data) < 1 {
		return
	}
	cores := runtime.NumCPU()
	eachLimit(data, cores, cb)
}

func EachProc(data []interface{}, cb func(interface{})) {
	if len(data) < 1 {
		return
	}
	cores := runtime.NumCPU()
	procs := runtime.GOMAXPROCS(cores)
	eachLimit(data, procs, cb)
}

func Map(data []interface{}, cb func(interface{}) interface{}) []interface{} {
	l := len(data)
	ret := make([]interface{}, l, l)
	if l < 1 {
		return ret
	}

	ch := make(chan int, l)
	defer close(ch)
	for i := 0; i < l; i++ {
		ch <- 1
	}

	c := make(chan int, 1)
	defer close(c)

	for idx, d := range data {
		idx := idx
		d := d
		go func() {
			defer func() {
				<-ch
				if len(ch) == 0 {
					c <- 1
				}
			}()
			ret[idx] = cb(d)
		}()
	}

	<-c
	return ret
}
