package steamail_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSteamail(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Steamail Suite")
}
