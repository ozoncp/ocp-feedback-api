package feedback_grpc_test

import (
	"context"
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Shopify/sarama/mocks"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	ocp_mocks "github.com/ozoncp/ocp-feedback-api/internal/mocks"
	"github.com/ozoncp/ocp-feedback-api/internal/producer"
	"github.com/ozoncp/ocp-feedback-api/internal/repo"
	"github.com/ozoncp/ocp-feedback-api/internal/server/feedback_grpc"
	fb "github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api"
	"github.com/rs/zerolog"
)

var _ = Describe("GrpcService", func() {
	var (
		controller   *gomock.Controller
		mockProm     *ocp_mocks.MockPromMetrics
		ctx          context.Context
		sqlmk        sqlmock.Sqlmock
		cancel       context.CancelFunc
		asynProdMock *mocks.AsyncProducer
		db           *sql.DB
		grpcService  *feedback_grpc.FeedbackService
		p            producer.Producer
	)

	BeforeEach(func() {
		zerolog.SetGlobalLevel(zerolog.Disabled)
		controller = gomock.NewController(GinkgoT())
		mockProm = ocp_mocks.NewMockPromMetrics(controller)
		ctx = context.Background()
		ctx, cancel = context.WithCancel(ctx)
		asynProdMock = mocks.NewAsyncProducer(controller.T, nil)
		prod, err := producer.New("feedbacks", asynProdMock)
		p = prod
		if err != nil {
			controller.T.Errorf("unable to create sarama mock: %s", err)
		}
		prod.Init(ctx)
		db, sqlmk, _ = sqlmock.New()

	})

	AfterEach(func() {
		db.Close()
		controller.Finish()
	})

	When("CreateFeedback is called", func() {
		It("should process a message", func() {
			repo := repo.NewFeedbackRepo(sqlx.NewDb(db, ""))
			grpcService = feedback_grpc.New(repo, p, mockProm, 2)

			asynProdMock.ExpectInputAndSucceed()
			mockProm.EXPECT().IncCreate()

			nf1 := fb.Feedback{UserId: 10, ClassroomId: 24, Comment: "hello1"}

			sqlmk.ExpectBegin()
			sqlmk.ExpectPrepare("INSERT INTO reaction.feedback").
				ExpectQuery().
				WithArgs(nf1.UserId, nf1.ClassroomId, nf1.Comment).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			sqlmk.ExpectCommit()

			reqCreate := &fb.CreateFeedbackV1Request{Feedback: &nf1}
			respCreate, err := grpcService.CreateFeedbackV1(ctx, reqCreate)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(respCreate).ShouldNot(BeNil())
			Ω(respCreate.FeedbackId).Should(Equal(uint64(1)))

			if err := sqlmk.ExpectationsWereMet(); err != nil {
				controller.T.Errorf("there were unfulfilled expectations: %s", err)
			}
			// invalid request, must fail on validation
			nf2 := fb.Feedback{UserId: 0, ClassroomId: 240, Comment: "hello2"}
			reqCreate = &fb.CreateFeedbackV1Request{Feedback: &nf2}
			respCreate, err = grpcService.CreateFeedbackV1(ctx, reqCreate)
			Ω(respCreate).Should(BeNil())
			Ω(err).Should(HaveOccurred())
			cancel()
			p.Close()
		})
	})

	When("CreateMultiFeedback is called", func() {
		It("should process a message", func() {
			repo := repo.NewFeedbackRepo(sqlx.NewDb(db, ""))
			grpcService = feedback_grpc.New(repo, p, mockProm, 2)

			asynProdMock.ExpectInputAndSucceed()
			asynProdMock.ExpectInputAndSucceed()
			asynProdMock.ExpectInputAndSucceed()
			asynProdMock.ExpectInputAndSucceed()
			mockProm.EXPECT().IncCreate().Times(4)

			// valid request
			nf1 := fb.Feedback{UserId: 42, ClassroomId: 24, Comment: "hello1"}
			nf2 := fb.Feedback{UserId: 420, ClassroomId: 240, Comment: "hello2"}
			nf3 := fb.Feedback{UserId: 4200, ClassroomId: 2400, Comment: "hello3"}
			nf4 := fb.Feedback{UserId: 42000, ClassroomId: 24000, Comment: "hello4"}

			reqMultiCreate := &fb.CreateMultiFeedbackV1Request{
				Feedbacks: []*fb.Feedback{
					&nf1,
					&nf2,
					&nf3,
					&nf4,
				},
			}

			// sequence numbers
			assignedNumbers := []*sqlmock.Rows{}
			for i := 1; i <= len(reqMultiCreate.Feedbacks); i++ {
				assignedNumbers = append(assignedNumbers, sqlmock.NewRows([]string{"id"}).AddRow(i))
			}

			// assume that feedbacks will be split into 2 chunks of size 2
			sqlmk.ExpectBegin()
			sqlmk.ExpectPrepare("INSERT INTO reaction.feedback")
			sqlmk.ExpectQuery("INSERT INTO reaction.feedback").
				WithArgs(nf1.UserId, nf1.ClassroomId, nf1.Comment).
				WillReturnRows(assignedNumbers[0])
			sqlmk.ExpectQuery("INSERT INTO reaction.feedback").
				WithArgs(nf2.UserId, nf2.ClassroomId, nf2.Comment).
				WillReturnRows(assignedNumbers[1])
			sqlmk.ExpectCommit()

			sqlmk.ExpectBegin()
			sqlmk.ExpectPrepare("INSERT INTO reaction.feedback")
			sqlmk.ExpectQuery("INSERT INTO reaction.feedback").
				WithArgs(nf3.UserId, nf3.ClassroomId, nf3.Comment).
				WillReturnRows(assignedNumbers[2])
			sqlmk.ExpectQuery("INSERT INTO reaction.feedback").
				WithArgs(nf4.UserId, nf4.ClassroomId, nf4.Comment).
				WillReturnRows(assignedNumbers[3])
			sqlmk.ExpectCommit()

			respMultiCreate, err := grpcService.CreateMultiFeedbackV1(ctx, reqMultiCreate)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(respMultiCreate).ShouldNot(BeNil())
			Ω(respMultiCreate.FeedbackId).Should(Equal([]uint64{1, 2, 3, 4}))

			if err := sqlmk.ExpectationsWereMet(); err != nil {
				controller.T.Errorf("there were unfulfilled expectations: %s", err)
			}

			// invalid request
			nf5 := fb.Feedback{UserId: 0, ClassroomId: 24, Comment: "hello1"}
			reqMultiCreate = &fb.CreateMultiFeedbackV1Request{
				Feedbacks: []*fb.Feedback{
					&nf5,
				},
			}

			respMultiCreate, err = grpcService.CreateMultiFeedbackV1(ctx, reqMultiCreate)
			Ω(respMultiCreate).Should(BeNil())
			Ω(err).Should(HaveOccurred())
			cancel()
			p.Close()
		})

		When("RemoveFeedback is called", func() {
			It("should process a message", func() {
				repo := repo.NewFeedbackRepo(sqlx.NewDb(db, ""))
				grpcService = feedback_grpc.New(repo, p, mockProm, 1)

				asynProdMock.ExpectInputAndSucceed()
				asynProdMock.ExpectInputAndSucceed()
				mockProm.EXPECT().IncCreate()
				mockProm.EXPECT().IncRemove()

				nf1 := fb.Feedback{UserId: 42, ClassroomId: 24, Comment: "hello1"}

				sqlmk.ExpectBegin()
				sqlmk.ExpectPrepare("INSERT INTO reaction.feedback").
					ExpectQuery().
					WithArgs(nf1.UserId, nf1.ClassroomId, nf1.Comment).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				sqlmk.ExpectCommit()

				reqCreate := &fb.CreateFeedbackV1Request{Feedback: &nf1}

				respCreate, err := grpcService.CreateFeedbackV1(ctx, reqCreate)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(respCreate).ShouldNot(BeNil())
				Ω(respCreate.FeedbackId).Should(Equal(uint64(1)))

				sqlmk.ExpectQuery("DELETE FROM reaction.feedback").
					WithArgs(respCreate.FeedbackId).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				reqRemove := &fb.RemoveFeedbackV1Request{FeedbackId: respCreate.FeedbackId}

				respRemove, err := grpcService.RemoveFeedbackV1(ctx, reqRemove)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(respRemove).ShouldNot(BeNil())

				if err := sqlmk.ExpectationsWereMet(); err != nil {
					controller.T.Errorf("there were unfulfilled expectations: %s", err)
				}

				// try to remove the it the second time
				reqRemove = &fb.RemoveFeedbackV1Request{FeedbackId: respCreate.FeedbackId}
				respRemove, err = grpcService.RemoveFeedbackV1(ctx, reqRemove)
				Ω(respRemove).Should(BeNil())
				Ω(err).Should(HaveOccurred())

				// invalid reqest
				reqRemove = &fb.RemoveFeedbackV1Request{FeedbackId: 0}
				respRemove, err = grpcService.RemoveFeedbackV1(ctx, reqRemove)
				Ω(respRemove).Should(BeNil())
				Ω(err).Should(HaveOccurred())
				cancel()
				p.Close()
			})
		})

		When("DescribeFeedback is called", func() {
			It("should process a message", func() {
				repo := repo.NewFeedbackRepo(sqlx.NewDb(db, ""))
				grpcService = feedback_grpc.New(repo, p, mockProm, 1)

				asynProdMock.ExpectInputAndSucceed()
				mockProm.EXPECT().IncCreate()

				nf := fb.Feedback{UserId: 42, ClassroomId: 24, Comment: "hello1"}

				sqlmk.ExpectBegin()
				sqlmk.ExpectPrepare("INSERT INTO reaction.feedback").
					ExpectQuery().
					WithArgs(nf.UserId, nf.ClassroomId, nf.Comment).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				sqlmk.ExpectCommit()

				reqCreate := &fb.CreateFeedbackV1Request{Feedback: &nf}
				respCreate, err := grpcService.CreateFeedbackV1(ctx, reqCreate)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(respCreate).ShouldNot(BeNil())
				Ω(respCreate.FeedbackId).Should(Equal(uint64(1)))

				returned := sqlmock.NewRows([]string{
					"id",
					"user_id",
					"classroom_id",
					"comment"},
				).AddRow(1, nf.UserId, nf.ClassroomId, nf.Comment)

				sqlmk.ExpectQuery("SELECT id, user_id, classroom_id, comment FROM reaction.feedback").
					WithArgs(respCreate.FeedbackId).WillReturnRows(returned)

				// valid request
				reqDescribe := &fb.DescribeFeedbackV1Request{FeedbackId: respCreate.FeedbackId}
				respDescribe, err := grpcService.DescribeFeedbackV1(ctx, reqDescribe)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(respDescribe).ShouldNot(BeNil())
				Ω(respDescribe.Feedback.FeedbackId).Should(Equal(uint64(1)))
				Ω(respDescribe.Feedback.UserId).Should(Equal(nf.UserId))
				Ω(respDescribe.Feedback.ClassroomId).Should(Equal(nf.ClassroomId))
				Ω(respDescribe.Feedback.Comment).Should(Equal(nf.Comment))

				if err := sqlmk.ExpectationsWereMet(); err != nil {
					controller.T.Errorf("there were unfulfilled expectations: %s", err)
				}

				// missing id
				reqDescribe = &fb.DescribeFeedbackV1Request{FeedbackId: respCreate.FeedbackId + 1}
				respDescribe, err = grpcService.DescribeFeedbackV1(ctx, reqDescribe)
				Ω(respDescribe).Should(BeNil())
				Ω(err).Should(HaveOccurred())

				// invalid request
				reqDescribe = &fb.DescribeFeedbackV1Request{FeedbackId: 0}
				respDescribe, err = grpcService.DescribeFeedbackV1(ctx, reqDescribe)
				Ω(respDescribe).Should(BeNil())
				Ω(err).Should(HaveOccurred())
				cancel()
				p.Close()
			})
		})

		When("ListFeedback is called", func() {
			It("should process a message", func() {
				repo := repo.NewFeedbackRepo(sqlx.NewDb(db, ""))
				grpcService = feedback_grpc.New(repo, p, mockProm, 2)

				asynProdMock.ExpectInputAndSucceed()
				asynProdMock.ExpectInputAndSucceed()
				mockProm.EXPECT().IncCreate()
				mockProm.EXPECT().IncCreate()

				nf1 := fb.Feedback{UserId: 42, ClassroomId: 24, Comment: "hello1"}
				nf2 := fb.Feedback{UserId: 420, ClassroomId: 240, Comment: "hello2"}

				reqCreate := &fb.CreateMultiFeedbackV1Request{Feedbacks: []*fb.Feedback{
					&nf1,
					&nf2,
				},
				}

				// assume that feedbacks won't be split into chunks
				sqlmk.ExpectBegin()
				sqlmk.ExpectPrepare("INSERT INTO reaction.feedback")
				sqlmk.ExpectQuery("INSERT INTO reaction.feedback").
					WithArgs(nf1.UserId, nf1.ClassroomId, nf1.Comment).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				sqlmk.ExpectQuery("INSERT INTO reaction.feedback").
					WithArgs(nf2.UserId, nf2.ClassroomId, nf2.Comment).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))
				sqlmk.ExpectCommit()

				respMultiCreate, err := grpcService.CreateMultiFeedbackV1(ctx, reqCreate)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(respMultiCreate).ShouldNot(BeNil())

				returned := sqlmock.NewRows([]string{
					"id",
					"user_id",
					"classroom_id",
					"comment"},
				).AddRow(1, nf1.UserId, nf1.ClassroomId, nf1.Comment).AddRow(
					2, nf2.UserId, nf2.ClassroomId, nf2.Comment,
				)

				sqlmk.ExpectQuery("SELECT id, user_id, classroom_id, comment FROM reaction.feedback").
					WithArgs(2, 0).WillReturnRows(returned)

				// valid request
				reqList := &fb.ListFeedbacksV1Request{Limit: 2, Offset: 0}
				respList, err := grpcService.ListFeedbacksV1(ctx, reqList)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(respList).ShouldNot(BeNil())
				Ω(len(respList.Feedbacks)).Should(Equal(2))

				Ω(respList.Feedbacks[0].FeedbackId).Should(Equal(uint64(1)))
				Ω(respList.Feedbacks[0].UserId).Should(Equal(nf1.UserId))
				Ω(respList.Feedbacks[0].ClassroomId).Should(Equal(nf1.ClassroomId))
				Ω(respList.Feedbacks[0].Comment).Should(Equal(nf1.Comment))

				Ω(respList.Feedbacks[1].FeedbackId).Should(Equal(uint64(2)))
				Ω(respList.Feedbacks[1].UserId).Should(Equal(nf2.UserId))
				Ω(respList.Feedbacks[1].ClassroomId).Should(Equal(nf2.ClassroomId))
				Ω(respList.Feedbacks[1].Comment).Should(Equal(nf2.Comment))

				sqlmk.ExpectQuery("SELECT id, user_id, classroom_id, comment FROM reaction.feedback").
					WithArgs(1, 4).WillReturnRows(sqlmock.NewRows([]string{
					"id",
					"user_id",
					"classroom_id",
					"comment"}))

				// wrong offset
				reqList = &fb.ListFeedbacksV1Request{Limit: 1, Offset: 4}
				respList, err = grpcService.ListFeedbacksV1(ctx, reqList)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(respList).ShouldNot(BeNil())
				Ω(len(respList.Feedbacks)).Should(Equal(0))

				if err := sqlmk.ExpectationsWereMet(); err != nil {
					controller.T.Errorf("there were unfulfilled expectations: %s", err)
				}

				// invalid request
				reqList = &fb.ListFeedbacksV1Request{Limit: 0, Offset: 1}
				respList, err = grpcService.ListFeedbacksV1(ctx, reqList)
				Ω(respList).Should(BeNil())
				Ω(err).Should(HaveOccurred())
				cancel()
				p.Close()
			})
		})

		When("UpdateFeedback is called", func() {
			It("should process a message", func() {
				repo := repo.NewFeedbackRepo(sqlx.NewDb(db, ""))
				grpcService = feedback_grpc.New(repo, p, mockProm, 1)

				asynProdMock.ExpectInputAndSucceed()
				asynProdMock.ExpectInputAndSucceed()
				mockProm.EXPECT().IncCreate()
				mockProm.EXPECT().IncUpdate()

				nf1 := fb.Feedback{UserId: 42, ClassroomId: 24, Comment: "hello1"}

				sqlmk.ExpectBegin()
				sqlmk.ExpectPrepare("INSERT INTO reaction.feedback").
					ExpectQuery().
					WithArgs(nf1.UserId, nf1.ClassroomId, nf1.Comment).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				sqlmk.ExpectCommit()

				reqCreate := &fb.CreateFeedbackV1Request{Feedback: &nf1}
				respCreate, err := grpcService.CreateFeedbackV1(ctx, reqCreate)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(respCreate).ShouldNot(BeNil())
				Ω(respCreate.FeedbackId).Should(Equal(uint64(1)))

				nf2 := fb.Feedback{FeedbackId: 1, UserId: 10, ClassroomId: 20, Comment: "hi"}
				sqlmk.ExpectQuery("UPDATE reaction.feedback").
					WithArgs(nf2.UserId, nf2.ClassroomId, nf2.Comment, nf2.FeedbackId).WillReturnRows(
					sqlmock.NewRows([]string{"id"}).AddRow(1),
				)

				reqUpdate := &fb.UpdateFeedbackV1Request{Feedback: &nf2}
				respUpdate, err := grpcService.UpdateFeedbackV1(ctx, reqUpdate)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(respUpdate).ShouldNot(BeNil())

				returned := sqlmock.NewRows([]string{
					"id",
					"user_id",
					"classroom_id",
					"comment"},
				).AddRow(1, nf2.UserId, nf2.ClassroomId, nf2.Comment)

				sqlmk.ExpectQuery("SELECT id, user_id, classroom_id, comment FROM reaction.feedback").
					WithArgs(respCreate.FeedbackId).WillReturnRows(returned)

				// valid request
				reqDescribe := &fb.DescribeFeedbackV1Request{FeedbackId: respCreate.FeedbackId}
				respDescribe, err := grpcService.DescribeFeedbackV1(ctx, reqDescribe)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(respDescribe).ShouldNot(BeNil())
				Ω(respDescribe.Feedback.FeedbackId).Should(Equal(uint64(1)))
				Ω(respDescribe.Feedback.UserId).Should(Equal(nf2.UserId))
				Ω(respDescribe.Feedback.ClassroomId).Should(Equal(nf2.ClassroomId))
				Ω(respDescribe.Feedback.Comment).Should(Equal(nf2.Comment))

				if err := sqlmk.ExpectationsWereMet(); err != nil {
					controller.T.Errorf("there were unfulfilled expectations: %s", err)
				}

				// invalid reqest
				reqUpdate = &fb.UpdateFeedbackV1Request{}
				respUpdate, err = grpcService.UpdateFeedbackV1(ctx, reqUpdate)
				Ω(respUpdate).Should(BeNil())
				Ω(err).Should(HaveOccurred())
				cancel()
				p.Close()

			})
		})

	})

})
