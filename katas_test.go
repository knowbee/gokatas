package katas_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/knowbee/gokatas"
)

var _ = Describe("Katas", func() {
	It("return hello", func() {
		Expect(Hello()).To(Equal("hello world"))
	})
})
