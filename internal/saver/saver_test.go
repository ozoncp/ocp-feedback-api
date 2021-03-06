package saver_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-feedback-api/internal/mocks"
	"github.com/ozoncp/ocp-feedback-api/internal/models"
	"github.com/ozoncp/ocp-feedback-api/internal/saver"
)

var _ = Describe("Saver", func() {

	Describe("Constructor call", func() {
		var (
			controller *gomock.Controller
			flusher    *flusherStub
			alarmer    *alarmerStub
		)

		BeforeEach(func() {
			controller = gomock.NewController(GinkgoT())
			flusher = &flusherStub{}
			alarmer = &alarmerStub{make(chan struct{})}
		})

		AfterEach(func() {
			controller.Finish()
		})

		When("arguments are invalid", func() {
			It("should return an error", func() {
				By("receiving invalid capacity")
				got, err := saver.New(0, saver.DropAll, alarmer, flusher)
				Ω(err).Should(HaveOccurred())
				Ω(got).Should(BeNil())

				By("receiving a nil alarmer")
				got, err = saver.New(1, saver.DropAll, nil, flusher)
				Ω(err).Should(HaveOccurred())
				Ω(got).Should(BeNil())

				By("receiving a nil flusher")
				got, err = saver.New(1, saver.DropAll, alarmer, nil)
				Ω(err).Should(HaveOccurred())
				Ω(got).Should(BeNil())
			})
		})

		When("arguments are valid", func() {
			It("should return valid object", func() {
				got, err := saver.New(1, saver.DropAll, alarmer, flusher)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(got).ShouldNot(BeNil())
			})
		})
	})

	Describe("Init call", func() {

		var (
			controller  *gomock.Controller
			mockFlusher *mocks.MockFlusher
			alarmer     *alarmerStub
		)

		BeforeEach(func() {
			controller = gomock.NewController(GinkgoT())
			mockFlusher = mocks.NewMockFlusher(controller)
			alarmer = &alarmerStub{make(chan struct{})}
		})

		AfterEach(func() {
			controller.Finish()
		})

		Context("Flushing", func() {
			When("Cap is not reached", func() {
				When("Save is not called before closing ", func() {
					It("should flush everything on alarm and close shouldn't flush anything", func() {
						entities := []models.Entity{
							&entityStub{id: 1},
							&entityStub{id: 2},
							&entityStub{id: 3},
							&entityStub{id: 4},
						}

						saver, _ := saver.New(len(entities)+1, saver.DropAll, alarmer, mockFlusher)
						ctx, cancel := context.WithCancel(context.Background())
						defer saver.WaitClosed()
						defer cancel()

						gomock.InOrder(
							mockFlusher.EXPECT().Flush(ctx, gomock.Eq(entities)),
							mockFlusher.EXPECT().Flush(ctx, gomock.Eq(entities[:0])),
						)

						saver.Init(ctx)

						for i := 0; i < len(entities); i++ {
							saver.Save(entities[i])
						}
						alarmer.alarm()
					})
				})
				When("Save is called before closing ", func() {
					It("should flush something on alarm and the rest after close has been called", func() {
						entities := []models.Entity{
							&entityStub{id: 1},
							&entityStub{id: 2},
							&entityStub{id: 3},
							&entityStub{id: 4},
						}
						saver, _ := saver.New(len(entities)+1, saver.DropAll, alarmer, mockFlusher)

						ctx, cancel := context.WithCancel(context.Background())
						defer saver.WaitClosed()
						defer cancel()

						gomock.InOrder(
							mockFlusher.EXPECT().Flush(ctx, gomock.Eq(entities[:2])),
							mockFlusher.EXPECT().Flush(ctx, gomock.Eq(entities[2:])),
						)

						saver.Init(ctx)

						for i := 0; i < 2; i++ {
							saver.Save(entities[i])
						}
						alarmer.alarm()

						for i := 2; i < len(entities); i++ {
							saver.Save(entities[i])
						}
					})
				})

			})

			When("Cap is reached", func() {
				When("Policy is DropAll", func() {
					It("should drop everything", func() {
						entities := []models.Entity{
							&entityStub{id: 1},
							&entityStub{id: 2},
							&entityStub{id: 3},
							&entityStub{id: 4},
						}

						newEntities := []models.Entity{
							&entityStub{id: 10},
							&entityStub{id: 20},
						}

						saver, _ := saver.New(len(entities), saver.DropAll, alarmer, mockFlusher)
						ctx, cancel := context.WithCancel(context.Background())
						defer saver.WaitClosed()
						defer cancel()

						gomock.InOrder(
							mockFlusher.EXPECT().Flush(ctx, gomock.Eq(newEntities)),
							mockFlusher.EXPECT().Flush(ctx, gomock.Eq(newEntities[:0])),
						)

						saver.Init(ctx)

						for i := 0; i < len(entities); i++ {
							saver.Save(entities[i])
						}

						for i := 0; i < len(newEntities); i++ {
							saver.Save(newEntities[i])
						}

						alarmer.alarm()
					})
				})
			})

			When("Policy is DropOne", func() {
				It("should drop the oldest one", func() {
					entities := []models.Entity{
						&entityStub{id: 1},
						&entityStub{id: 2},
						&entityStub{id: 3},
						&entityStub{id: 4},
					}

					shifted := []models.Entity{
						&entityStub{id: 3},
						&entityStub{id: 4},
						&entityStub{id: 5},
						&entityStub{id: 6},
					}

					saver, _ := saver.New(len(entities), saver.DropOne, alarmer, mockFlusher)
					ctx, cancel := context.WithCancel(context.Background())
					defer saver.WaitClosed()
					defer cancel()

					gomock.InOrder(
						mockFlusher.EXPECT().Flush(ctx, gomock.Eq(shifted)),
						mockFlusher.EXPECT().Flush(ctx, gomock.Eq(shifted[:0])),
					)

					saver.Init(ctx)

					for i := 0; i < len(entities); i++ {
						saver.Save(entities[i])
					}

					saver.Save(&entityStub{id: 5})
					saver.Save(&entityStub{id: 6})

					alarmer.alarm()
				})
			})
		})
		When("Close is called", func() {
			It("should flush everything at once", func() {
				entities := []models.Entity{
					&entityStub{id: 1},
					&entityStub{id: 2},
					&entityStub{id: 3},
					&entityStub{id: 4},
				}

				saver, _ := saver.New(len(entities)+1, saver.DropAll, alarmer, mockFlusher)
				ctx, cancel := context.WithCancel(context.Background())
				defer saver.WaitClosed()
				defer cancel()

				mockFlusher.EXPECT().Flush(ctx, gomock.Eq(entities))

				saver.Init(ctx)

				for i := 0; i < len(entities); i++ {
					saver.Save(entities[i])
				}
				cancel()
			})
		})

		When("Flush has failed", func() {
			It("should re-flush remaining", func() {
				entities := []models.Entity{
					&entityStub{id: 1},
					&entityStub{id: 2},
					&entityStub{id: 3},
					&entityStub{id: 4},
				}

				saver, _ := saver.New(len(entities), saver.DropAll, alarmer, mockFlusher)
				ctx, cancel := context.WithCancel(context.Background())
				defer saver.WaitClosed()
				defer cancel()

				gomock.InOrder(
					mockFlusher.EXPECT().Flush(ctx, gomock.Eq(entities)).
						Return([]models.Entity{&entityStub{id: 3}, &entityStub{id: 4}},
							errors.New("flushing failed")),
					mockFlusher.EXPECT().Flush(ctx, gomock.Eq(entities[2:])),
				)

				saver.Init(ctx)

				for i := 0; i < len(entities); i++ {
					saver.Save(entities[i])
				}
				alarmer.alarm()
			})
		})

		When("Flush has failed on close", func() {
			It("should handle an error", func() {
				entities := []models.Entity{
					&entityStub{id: 1},
					&entityStub{id: 2},
					&entityStub{id: 3},
					&entityStub{id: 4},
				}

				saver, _ := saver.New(len(entities), saver.DropAll, alarmer, mockFlusher)
				ctx, cancel := context.WithCancel(context.Background())
				defer saver.WaitClosed()
				defer cancel()

				mockFlusher.EXPECT().Flush(ctx, gomock.Eq(entities)).
					Return(entities, errors.New("flushing failed"))

				saver.Init(ctx)

				for i := 0; i < len(entities); i++ {
					saver.Save(entities[i])
				}
			})
		})
	})

})

type alarmerStub struct {
	alarms chan struct{}
}

func (a *alarmerStub) Alarm() <-chan struct{} {
	return a.alarms
}

func (a *alarmerStub) alarm() {
	a.alarms <- struct{}{}
}

type flusherStub struct {
}

func (f *flusherStub) Flush(ctx context.Context, entities []models.Entity) ([]models.Entity, error) {
	return nil, nil
}

type entityStub struct {
	id uint64
}

func (d *entityStub) ObjectId() uint64 {
	return d.id
}
