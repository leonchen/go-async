package async_test

import (
	"github.com/leonchen/go-async"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Async", func() {
	// Each
	It("should run func for each element", func() {
		data := []interface{}{1, 3, 5}

		count := 0
		async.Each(data, func(d interface{}) {
			count += d.(int)
		})

		Expect(count).To(Equal(9))
	})

	// EachLimit
	It("should run func for each element with limit", func() {
		data := []interface{}{1, 3, 5, 7, 9}

		ch := make(chan int, len(data))
		async.EachLimit(data, 2, func(d interface{}) {
			if d.(int) == 5 {
				Expect(len(ch)).To(Equal(2))
			}
			if d.(int) == 9 {
				Expect(len(ch)).To(Equal(4))
			}
			ch <- 1
		})
	})

	// Map
	It("should return each result for map", func() {
		data := []interface{}{1, 3, 5}
		ret := async.Map(data, func(d interface{}) interface{} {
			return d.(int) + 1
		})

		Expect(len(ret)).To(Equal(3))
		Expect(ret[0].(int)).To(Equal(2))
		Expect(ret[1].(int)).To(Equal(4))
		Expect(ret[2].(int)).To(Equal(6))
	})

	// Times
	It("should return each result for times", func() {
		ret := async.Times(5, func(n int) interface{} {
			return n
		})

		Expect(len(ret)).To(Equal(5))
		Expect(ret[0].(int)).To(Equal(0))
		Expect(ret[4].(int)).To(Equal(4))
	})

})
