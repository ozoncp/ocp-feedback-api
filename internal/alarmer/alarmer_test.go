package alarmer_test

import (
	"fmt"
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
				ticks := 0
				al.Init()
				timer := time.NewTimer(400 * time.Millisecond)
				go func() {
					defer al.Close()
					<-timer.C
					fmt.Println("ticks", ticks)
					// timers are not very accurate
					Ω(ticks).Should(BeNumerically(">=", 9))
				}()

				for range al.Alarm() {
					ticks++
				}
			})
		})

	})

})
