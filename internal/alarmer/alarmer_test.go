package alarmer_test

import (
	"sync/atomic"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"

	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-feedback-api/internal/alarmer"
)

var _ = Describe("Alarmer", func() {

	Describe("Init call", func() {

		var (
			controller *gomock.Controller
		)

		BeforeEach(func() {
			controller = gomock.NewController(GinkgoT())
		})

		AfterEach(func() {
			controller.Finish()
		})

		Context("Alarmer looping", func() {
			It("should be closed correctly", func() {
				al := alarmer.New(10 * time.Millisecond)
				al.Init()
				timer := time.NewTimer(100 * time.Millisecond)
				go func() {
					defer al.Close()
					<-timer.C
				}()
				Eventually(al.Alarm()).Should(BeClosed())
			})

			It("should notify correct number of times (approximately) ", func() {
				al := alarmer.New(40 * time.Millisecond)
				var ticks int32
				al.Init()
				timer := time.NewTimer(400 * time.Millisecond)
				go func() {
					defer al.Close()
					<-timer.C
					// timers are not very accurate
					Ω(atomic.LoadInt32(&ticks)).Should(BeNumerically(">=", 9))
				}()
				for range al.Alarm() {
					atomic.AddInt32(&ticks, 1)
				}
			})
		})

	})

})
