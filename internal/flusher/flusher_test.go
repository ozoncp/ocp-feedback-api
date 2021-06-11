package flusher_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-feedback-api/internal/flusher"
	"github.com/ozoncp/ocp-feedback-api/internal/mocks"
	"github.com/ozoncp/ocp-feedback-api/internal/models"
)

// top-level describe container
var _ = Describe("Flusher", func() {

	Describe("Constructor call", func() {

		When("arguments are invalid", func() {
			It("should return an error", func() {
				By("receiving a chunk size which is < 0")
				got, err := flusher.New(-1, &repoStub{})
				Ω(err).Should(HaveOccurred())
				Ω(got).Should(BeNil())

				By("receiving a nil repo")
				got, err = flusher.New(1, nil)
				Ω(err).Should(HaveOccurred())
				Ω(got).Should(BeNil())
			})
		})

		When("arguments are valid", func() {
			It("should return valid object", func() {
				got, err := flusher.New(1, &repoStub{})
				Ω(err).ShouldNot(HaveOccurred())
				Ω(got).ShouldNot(BeNil())
			})
		})
	})

	Describe("Flush call", func() {

		var (
			err        error
			controller *gomock.Controller
			mockRepo   *mocks.MockBatchAdder
			entities   []models.Entity
			remain     []models.Entity
			f          flusher.Flusher
			ctx        context.Context
		)

		BeforeEach(func() {
			// prevent vars from being mutated by It blocks
			err = nil
			controller = gomock.NewController(GinkgoT())
			mockRepo = mocks.NewMockBatchAdder(controller)
			entities = nil
			remain = nil
			f = nil
			ctx = context.Background()
		})

		AfterEach(func() {
			controller.Finish()
		})

		When("slice is nil", func() {
			It("should return error", func() {
				f, _ = flusher.New(1, mockRepo)
				remain, err = f.Flush(ctx, entities)
				Ω(err).Should(HaveOccurred())
				Ω(remain).Should(BeNil())
			})
		})

		When("flush operation successful", func() {
			It("should flush all entities if thier number is divisible by chunk size ", func() {
				entities = []models.Entity{
					&entityStub{id: 1},
					&entityStub{id: 2},
					&entityStub{id: 3},
					&entityStub{id: 4},
				}
				f, _ = flusher.New(len(entities)/2, mockRepo)
				gomock.InOrder(
					mockRepo.EXPECT().AddEntities(ctx, entities[:2]),
					mockRepo.EXPECT().AddEntities(ctx, entities[2:]),
				)
				remain, err = f.Flush(ctx, entities)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(remain).Should(BeNil())
			})

			It("should flush all entities if thier number is not divisible by chunk size ", func() {
				entities = []models.Entity{
					&entityStub{id: 1},
					&entityStub{id: 2},
					&entityStub{id: 3},
					&entityStub{id: 4},
					&entityStub{id: 5},
				}
				f, _ = flusher.New(2, mockRepo)
				gomock.InOrder(
					mockRepo.EXPECT().AddEntities(ctx, entities[:2]),
					mockRepo.EXPECT().AddEntities(ctx, entities[2:4]),
					mockRepo.EXPECT().AddEntities(ctx, entities[4:]),
				)
				remain, err = f.Flush(ctx, entities)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(remain).Should(BeNil())
			})
			It("should flush all entities at once if chunk size is zero", func() {
				entities = []models.Entity{
					&entityStub{id: 1},
					&entityStub{id: 2},
					&entityStub{id: 3},
					&entityStub{id: 4},
				}
				f, _ = flusher.New(0, mockRepo)

				mockRepo.EXPECT().AddEntities(ctx, entities)

				remain, err = f.Flush(ctx, entities)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(remain).Should(BeNil())
			})
			It("should flush all entities if chunk size is greater than number of entities", func() {
				entities = []models.Entity{
					&entityStub{id: 1},
					&entityStub{id: 2},
					&entityStub{id: 3},
					&entityStub{id: 4},
				}
				f, _ = flusher.New(len(entities)+1, mockRepo)

				mockRepo.EXPECT().AddEntities(ctx, entities)

				remain, err = f.Flush(ctx, entities)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(remain).Should(BeNil())
			})
		})

		When("flush operation fails", func() {
			It("should return remaining entities if AddEntities fails in the middle of operation", func() {
				entities = []models.Entity{
					&entityStub{id: 1},
					&entityStub{id: 2},
					&entityStub{id: 3},
					&entityStub{id: 4},
				}
				f, _ = flusher.New(2, mockRepo)

				gomock.InOrder(
					mockRepo.EXPECT().AddEntities(ctx, entities[:2]),
					mockRepo.EXPECT().AddEntities(ctx, entities[2:]).Return(nil, errors.New("add entities fails")),
				)

				remain, err = f.Flush(ctx, entities)
				Ω(err).Should(HaveOccurred())
				Ω(remain).Should(BeEquivalentTo(entities[2:]))
			})

			It("should return all entities if AddEntities fails in the beginning of operation", func() {
				entities = []models.Entity{
					&entityStub{id: 1},
					&entityStub{id: 2},
					&entityStub{id: 3},
					&entityStub{id: 4},
				}
				f, _ = flusher.New(2, mockRepo)

				mockRepo.EXPECT().AddEntities(ctx, entities[:2]).Return(nil, errors.New("add entities fails"))

				remain, err = f.Flush(ctx, entities)
				Ω(err).Should(HaveOccurred())
				Ω(remain).Should(BeEquivalentTo(entities))
			})
		})
	})
})

type repoStub struct{}

func (d *repoStub) AddEntities(ctx context.Context, entities ...models.Entity) ([]uint64, error) {
	return nil, nil
}

type entityStub struct {
	id uint64
}

func (d *entityStub) ObjectId() uint64 {
	return d.id
}
