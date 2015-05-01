package sample_test

import (
	. "github.com/OneRedOak/sample"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sample", func() {
    Describe("Adding and retreving Person object from MongoDB", func() {
        Context("inspecting their name", func() {
            It("should result 'Ale'", func() {
                Expect(GetResult()).To(Equal("Alen"))
            })
        })
    })
})
