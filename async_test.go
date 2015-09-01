package async_test

import (
	"github.com/leonchen/go-async"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Async", func() {
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

})
