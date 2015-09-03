package async

import (
	"runtime"
	"sync"
)

func Each(data []interface{}, f func(interface{})) {
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
			f(d)
		}()
	}
	<-c
}

func EachLimit(data []interface{}, limit int, f func(interface{})) {
	if len(data) < 1 {
		return
	}
	eachLimit(data, limit, f)
}

func eachLimit(data []interface{}, limit int, f func(interface{})) {
	ch := make(chan int, limit)
	defer close(ch)

	var wg sync.WaitGroup
	l := len(data)
	wg.Add(l)

	for _, d := range data {
		d := d
		ch <- 1
		go func() {
			f(d)
			<-ch
			wg.Done()
		}()
	}
	wg.Wait()
}

func EachCPU(data []interface{}, f func(interface{})) {
	if len(data) < 1 {
		return
	}
	cores := runtime.NumCPU()
	eachLimit(data, cores, f)
}

func EachProc(data []interface{}, f func(interface{})) {
	if len(data) < 1 {
		return
	}
	cores := runtime.NumCPU()
	procs := runtime.GOMAXPROCS(cores)
	eachLimit(data, procs, f)
}

func Map(data []interface{}, f func(interface{}) interface{}) []interface{} {
	l := len(data)
	ret := make([]interface{}, l, l)
	if l < 1 {
		return ret
	}

	var wg sync.WaitGroup
	wg.Add(l)

	for idx, d := range data {
		idx := idx
		d := d
		go func() {
			defer wg.Done()
			ret[idx] = f(d)
		}()
	}

	wg.Wait()
	return ret
}

func Times(n int, f func(int) interface{}) []interface{} {
	ret := make([]interface{}, n, n)
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		i := i
		go func() {
			defer wg.Done()
			ret[i] = f(i)
		}()
	}

	wg.Wait()
	return ret
}
