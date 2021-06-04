package saver_test

import (
	"sync"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozoncp/ocp-feedback-api/internal/mocks"
	"github.com/ozoncp/ocp-feedback-api/internal/models"
	"github.com/ozoncp/ocp-feedback-api/internal/saver"
	//. "github.com/onsi/gomega"
)

var _ = Describe("Saver", func() {

	Describe("Init call", func() {

		var (
			controller  *gomock.Controller
			mockFlusher *mocks.MockFlusher
		)

		BeforeEach(func() {
			controller = gomock.NewController(GinkgoT())
			mockFlusher = mocks.NewMockFlusher(controller)
		})

		AfterEach(func() {
			controller.Finish()
		})

		Context("Save is called", func() {
			When("Cap is not reached", func() {
				It("should flush everything", func() {
					entities := []models.Entity{
						&entityStub{id: 1},
						&entityStub{id: 2},
						&entityStub{id: 3},
						&entityStub{id: 4},
					}
					var wg sync.WaitGroup
					wg.Add(1)
					defer wg.Wait()

					al := &alarmerStub{}
					saver, _ := saver.New(len(entities), saver.DropAll, al, mockFlusher)
					mockFlusher.EXPECT().Flush(gomock.Eq(entities)).Do(func(entities []models.Entity) { wg.Done() })
					saver.Init()

					for i := 0; i < len(entities); i++ {
						saver.Save(entities[i])
					}
					saver.Close()
				})
			})

			// When("Cap is reached", func() {
			// 	When("Policy is DropAll", func() {
			// 		It("should drop everything", func() {})
			// 	})

			// 	When("Policy is DropOne", func() {
			// 		It("should drop oldest", func() {})
			// 	})
			// })
		})

		// When("Close is called", func() {
		// 	It("should flush everything", func() {})
		// })

		// When("Flush has failed", func() {
		// 	It("should re-flush remaining", func() {})
		// })
	})
})

type alarmerStub struct {
	alarms chan struct{}
}

func (a *alarmerStub) Alarm() <-chan struct{} {
	return a.alarms
}

// func (a *alarmerStub) alarm() {
// 	a.alarms <- struct{}{}
// }

type entityStub struct {
	id uint64
}

func (d *entityStub) ObjectId() uint64 {
	return d.id
}
