package augustine_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAugustine(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Augustine's Suite")
}
