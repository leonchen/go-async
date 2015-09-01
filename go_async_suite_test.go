package async_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoAsync(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoAsync Suite")
}
