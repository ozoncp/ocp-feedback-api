package saver_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
)

var _ = Describe("Saver", func() {

	Describe("Init call", func() {

		var (
			controller *gomock.Controller
			//mockFlusher *mocks.MockFlusher
		)

		BeforeEach(func() {
			controller = gomock.NewController(GinkgoT())
		})

		AfterEach(func() {
			controller.Finish()
		})

		Context("Save is called", func() {
			When("Cap is not reached", func() {
				It("should flush everything", func() {

				})
			})

			When("Cap is reached", func() {
				When("Policy is DropAll", func() {
					It("should drop everything", func() {})
				})

				When("Policy is DropOne", func() {
					It("should drop oldest", func() {})
				})
			})
		})

		When("Close is called", func() {
			It("should flush everything", func() {})
		})

		When("Flush has failed", func() {
			It("should re-flush remaining", func() {})
		})
	})
})

type alarmerStub struct {
	alarms chan struct{}
}

func (a *alarmerStub) Alarm() <-chan struct{} {
	return a.alarms
}
