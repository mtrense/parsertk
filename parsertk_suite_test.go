package parsertk_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestParsertk(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parsertk Suite")
}
