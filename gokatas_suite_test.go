package katas_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGokatas(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gokatas Suite")
}
