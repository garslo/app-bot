package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAppBot(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "App-Bot Suite")
}
