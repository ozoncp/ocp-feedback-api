package proposal_grpc_test

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
	"github.com/ozoncp/ocp-feedback-api/internal/server/proposal_grpc"
	pr "github.com/ozoncp/ocp-feedback-api/pkg/ocp-proposal-api"
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
		grpcService  *proposal_grpc.ProposalService
		p            producer.Producer
	)

	BeforeEach(func() {
		zerolog.SetGlobalLevel(zerolog.Disabled)
		controller = gomock.NewController(GinkgoT())
		mockProm = ocp_mocks.NewMockPromMetrics(controller)
		ctx = context.Background()
		ctx, cancel = context.WithCancel(ctx)
		asynProdMock = mocks.NewAsyncProducer(controller.T, nil)
		prod, err := producer.New("proposals", asynProdMock)
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

	When("CreateProposal is called", func() {
		It("should process a message", func() {
			repo := repo.NewProposalRepo(sqlx.NewDb(db, ""))
			grpcService = proposal_grpc.New(repo, p, mockProm, 2)

			asynProdMock.ExpectInputAndSucceed()
			mockProm.EXPECT().IncCreate()

			np1 := pr.Proposal{UserId: 10, LessonId: 20, DocumentId: 30}

			sqlmk.ExpectBegin()
			sqlmk.ExpectPrepare("INSERT INTO reaction.proposal").
				ExpectQuery().
				WithArgs(np1.UserId, np1.LessonId, np1.DocumentId).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			sqlmk.ExpectCommit()

			reqCreate := &pr.CreateProposalV1Request{Proposal: &np1}
			respCreate, err := grpcService.CreateProposalV1(ctx, reqCreate)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(respCreate).ShouldNot(BeNil())
			Ω(respCreate.ProposalId).Should(Equal(uint64(1)))

			if err := sqlmk.ExpectationsWereMet(); err != nil {
				controller.T.Errorf("there were unfulfilled expectations: %s", err)
			}
			// // invalid request, must fail on validation
			np2 := pr.Proposal{UserId: 0, LessonId: 240, DocumentId: 544}
			reqCreate = &pr.CreateProposalV1Request{Proposal: &np2}
			respCreate, err = grpcService.CreateProposalV1(ctx, reqCreate)
			Ω(respCreate).Should(BeNil())
			Ω(err).Should(HaveOccurred())
			cancel()
			p.Close()
		})
	})

	When("CreateMultiProposal is called", func() {
		It("should process a message", func() {
			repo := repo.NewProposalRepo(sqlx.NewDb(db, ""))
			grpcService = proposal_grpc.New(repo, p, mockProm, 2)

			asynProdMock.ExpectInputAndSucceed()
			asynProdMock.ExpectInputAndSucceed()
			asynProdMock.ExpectInputAndSucceed()
			asynProdMock.ExpectInputAndSucceed()
			mockProm.EXPECT().IncCreate().Times(4)

			// valid request
			pr1 := pr.Proposal{UserId: 42, LessonId: 24, DocumentId: 50}
			pr2 := pr.Proposal{UserId: 420, LessonId: 240, DocumentId: 500}
			pr3 := pr.Proposal{UserId: 4200, LessonId: 2400, DocumentId: 5000}
			pr4 := pr.Proposal{UserId: 42000, LessonId: 24000, DocumentId: 50000}

			reqMultiCreate := &pr.CreateMultiProposalV1Request{
				Proposals: []*pr.Proposal{
					&pr1,
					&pr2,
					&pr3,
					&pr4,
				},
			}

			// sequence numbers
			assignedNumbers := []*sqlmock.Rows{}
			for i := 1; i <= len(reqMultiCreate.Proposals); i++ {
				assignedNumbers = append(assignedNumbers, sqlmock.NewRows([]string{"id"}).AddRow(i))
			}

			// assume that feedbacks will be split into 2 chunks of size 2
			sqlmk.ExpectBegin()
			sqlmk.ExpectPrepare("INSERT INTO reaction.proposal")
			sqlmk.ExpectQuery("INSERT INTO reaction.proposal").
				WithArgs(pr1.UserId, pr1.LessonId, pr1.DocumentId).
				WillReturnRows(assignedNumbers[0])
			sqlmk.ExpectQuery("INSERT INTO reaction.proposal").
				WithArgs(pr2.UserId, pr2.LessonId, pr2.DocumentId).
				WillReturnRows(assignedNumbers[1])
			sqlmk.ExpectCommit()

			sqlmk.ExpectBegin()
			sqlmk.ExpectPrepare("INSERT INTO reaction.proposal")
			sqlmk.ExpectQuery("INSERT INTO reaction.proposal").
				WithArgs(pr3.UserId, pr3.LessonId, pr3.DocumentId).
				WillReturnRows(assignedNumbers[2])
			sqlmk.ExpectQuery("INSERT INTO reaction.proposal").
				WithArgs(pr4.UserId, pr4.LessonId, pr4.DocumentId).
				WillReturnRows(assignedNumbers[3])
			sqlmk.ExpectCommit()

			respMultiCreate, err := grpcService.CreateMultiProposalV1(ctx, reqMultiCreate)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(respMultiCreate).ShouldNot(BeNil())
			Ω(respMultiCreate.Proposals).Should(Equal([]uint64{1, 2, 3, 4}))

			if err := sqlmk.ExpectationsWereMet(); err != nil {
				controller.T.Errorf("there were unfulfilled expectations: %s", err)
			}

			// invalid request
			np5 := pr.Proposal{UserId: 0, LessonId: 24, DocumentId: 100}
			reqMultiCreate = &pr.CreateMultiProposalV1Request{
				Proposals: []*pr.Proposal{
					&np5,
				},
			}

			respMultiCreate, err = grpcService.CreateMultiProposalV1(ctx, reqMultiCreate)
			Ω(respMultiCreate).Should(BeNil())
			Ω(err).Should(HaveOccurred())
			cancel()
			p.Close()
		})
	})

	When("RemoveProposal is called", func() {
		It("should process a message", func() {
			repo := repo.NewProposalRepo(sqlx.NewDb(db, ""))
			grpcService = proposal_grpc.New(repo, p, mockProm, 2)

			asynProdMock.ExpectInputAndSucceed()
			asynProdMock.ExpectInputAndSucceed()
			mockProm.EXPECT().IncCreate()
			mockProm.EXPECT().IncRemove()

			np1 := pr.Proposal{UserId: 10, LessonId: 20, DocumentId: 30}

			sqlmk.ExpectBegin()
			sqlmk.ExpectPrepare("INSERT INTO reaction.proposal").
				ExpectQuery().
				WithArgs(np1.UserId, np1.LessonId, np1.DocumentId).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			sqlmk.ExpectCommit()

			reqCreate := &pr.CreateProposalV1Request{Proposal: &np1}
			respCreate, err := grpcService.CreateProposalV1(ctx, reqCreate)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(respCreate).ShouldNot(BeNil())
			Ω(respCreate.ProposalId).Should(Equal(uint64(1)))

			sqlmk.ExpectQuery("DELETE FROM reaction.proposal").
				WithArgs(respCreate.ProposalId).WillReturnRows(
				sqlmock.NewRows([]string{"id"}).AddRow(1))

			reqRemove := &pr.RemoveProposalV1Request{ProposalId: respCreate.ProposalId}
			respRemove, err := grpcService.RemoveProposalV1(ctx, reqRemove)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(respRemove).ShouldNot(BeNil())

			if err := sqlmk.ExpectationsWereMet(); err != nil {
				controller.T.Errorf("there were unfulfilled expectations: %s", err)
			}

			// try to remove the it the second time
			reqRemove = &pr.RemoveProposalV1Request{ProposalId: respCreate.ProposalId}
			respRemove, err = grpcService.RemoveProposalV1(ctx, reqRemove)
			Ω(respRemove).Should(BeNil())
			Ω(err).Should(HaveOccurred())

			// invalid reqest
			reqRemove = &pr.RemoveProposalV1Request{ProposalId: 0}
			respRemove, err = grpcService.RemoveProposalV1(ctx, reqRemove)
			Ω(respRemove).Should(BeNil())
			Ω(err).Should(HaveOccurred())
			cancel()
			p.Close()
		})
	})

	When("DescribeProposal is called", func() {
		It("should process a message", func() {
			repo := repo.NewProposalRepo(sqlx.NewDb(db, ""))
			grpcService = proposal_grpc.New(repo, p, mockProm, 2)

			asynProdMock.ExpectInputAndSucceed()
			mockProm.EXPECT().IncCreate()

			np := pr.Proposal{UserId: 10, LessonId: 20, DocumentId: 30}

			sqlmk.ExpectBegin()
			sqlmk.ExpectPrepare("INSERT INTO reaction.proposal").
				ExpectQuery().
				WithArgs(np.UserId, np.LessonId, np.DocumentId).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			sqlmk.ExpectCommit()

			reqCreate := &pr.CreateProposalV1Request{Proposal: &np}
			respCreate, err := grpcService.CreateProposalV1(ctx, reqCreate)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(respCreate).ShouldNot(BeNil())
			Ω(respCreate.ProposalId).Should(Equal(uint64(1)))

			returned := sqlmock.NewRows([]string{
				"id",
				"user_id",
				"lesson_id",
				"document_id"},
			).AddRow(1, np.UserId, np.LessonId, np.DocumentId)

			sqlmk.ExpectQuery("SELECT id, user_id, lesson_id, document_id FROM reaction.proposal").
				WithArgs(respCreate.ProposalId).WillReturnRows(returned)

			// valid request
			reqDescribe := &pr.DescribeProposalV1Request{ProposalId: respCreate.ProposalId}
			respDescribe, err := grpcService.DescribeProposalV1(ctx, reqDescribe)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(respDescribe).ShouldNot(BeNil())
			Ω(respDescribe.Proposal.ProposalId).Should(Equal(uint64(1)))
			Ω(respDescribe.Proposal.UserId).Should(Equal(np.UserId))
			Ω(respDescribe.Proposal.LessonId).Should(Equal(np.LessonId))
			Ω(respDescribe.Proposal.DocumentId).Should(Equal(np.DocumentId))

			if err := sqlmk.ExpectationsWereMet(); err != nil {
				controller.T.Errorf("there were unfulfilled expectations: %s", err)
			}

			// missing id
			reqDescribe = &pr.DescribeProposalV1Request{ProposalId: respCreate.ProposalId + 1}
			respDescribe, err = grpcService.DescribeProposalV1(ctx, reqDescribe)
			Ω(respDescribe).Should(BeNil())
			Ω(err).Should(HaveOccurred())

			// invalid request
			reqDescribe = &pr.DescribeProposalV1Request{ProposalId: 0}
			respDescribe, err = grpcService.DescribeProposalV1(ctx, reqDescribe)
			Ω(respDescribe).Should(BeNil())
			Ω(err).Should(HaveOccurred())
			cancel()
			p.Close()
		})
	})

	When("ListProposal is called", func() {
		It("should process a message", func() {
			repo := repo.NewProposalRepo(sqlx.NewDb(db, ""))
			grpcService = proposal_grpc.New(repo, p, mockProm, 1)

			asynProdMock.ExpectInputAndSucceed()
			asynProdMock.ExpectInputAndSucceed()
			mockProm.EXPECT().IncCreate()
			mockProm.EXPECT().IncCreate()

			np1 := pr.Proposal{UserId: 42, LessonId: 24, DocumentId: 50}
			np2 := pr.Proposal{UserId: 420, LessonId: 240, DocumentId: 500}

			reqCreate := &pr.CreateMultiProposalV1Request{Proposals: []*pr.Proposal{
				&np1,
				&np2,
			},
			}

			// assume that feedbacks won't be split into chunks
			sqlmk.ExpectBegin()
			sqlmk.ExpectPrepare("INSERT INTO reaction.proposal")
			sqlmk.ExpectQuery("INSERT INTO reaction.proposal").
				WithArgs(np1.UserId, np1.LessonId, np1.DocumentId).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			sqlmk.ExpectQuery("INSERT INTO reaction.proposal").
				WithArgs(np2.UserId, np2.LessonId, np2.DocumentId).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))
			sqlmk.ExpectCommit()

			respMultiCreate, err := grpcService.CreateMultiProposalV1(ctx, reqCreate)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(respMultiCreate).ShouldNot(BeNil())

			returned := sqlmock.NewRows([]string{
				"id",
				"user_id",
				"lesson_id",
				"document_id"},
			).AddRow(1, np1.UserId, np1.LessonId, np1.DocumentId).AddRow(
				2, np2.UserId, np2.LessonId, np2.DocumentId,
			)

			sqlmk.ExpectQuery("SELECT id, user_id, lesson_id, document_id FROM reaction.proposal").
				WithArgs(2, 0).WillReturnRows(returned)

			// valid request
			reqList := &pr.ListProposalsV1Request{Limit: 2, Offset: 0}
			respList, err := grpcService.ListProposalsV1(ctx, reqList)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(respList).ShouldNot(BeNil())
			Ω(len(respList.Proposals)).Should(Equal(2))

			Ω(respList.Proposals[0].ProposalId).Should(Equal(uint64(1)))
			Ω(respList.Proposals[0].UserId).Should(Equal(np1.UserId))
			Ω(respList.Proposals[0].LessonId).Should(Equal(np1.LessonId))
			Ω(respList.Proposals[0].DocumentId).Should(Equal(np1.DocumentId))

			Ω(respList.Proposals[1].ProposalId).Should(Equal(uint64(2)))
			Ω(respList.Proposals[1].UserId).Should(Equal(np2.UserId))
			Ω(respList.Proposals[1].LessonId).Should(Equal(np2.LessonId))
			Ω(respList.Proposals[1].DocumentId).Should(Equal(np2.DocumentId))

			sqlmk.ExpectQuery("SELECT id, user_id, lesson_id, document_id FROM reaction.proposal").
				WithArgs(1, 4).WillReturnRows(sqlmock.NewRows([]string{
				"id",
				"user_id",
				"lesson_id",
				"document_id"}))

			// wrong offset
			reqList = &pr.ListProposalsV1Request{Limit: 1, Offset: 4}
			respList, err = grpcService.ListProposalsV1(ctx, reqList)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(respList).ShouldNot(BeNil())
			Ω(len(respList.Proposals)).Should(Equal(0))

			if err := sqlmk.ExpectationsWereMet(); err != nil {
				controller.T.Errorf("there were unfulfilled expectations: %s", err)
			}

			// invalid request
			reqList = &pr.ListProposalsV1Request{Limit: 0, Offset: 1}
			respList, err = grpcService.ListProposalsV1(ctx, reqList)
			Ω(respList).Should(BeNil())
			Ω(err).Should(HaveOccurred())
			cancel()
			p.Close()
		})
	})

	When("UpdateProposal is called", func() {
		It("should process a message", func() {
			repo := repo.NewProposalRepo(sqlx.NewDb(db, ""))
			grpcService = proposal_grpc.New(repo, p, mockProm, 1)

			asynProdMock.ExpectInputAndSucceed()
			asynProdMock.ExpectInputAndSucceed()
			mockProm.EXPECT().IncCreate()
			mockProm.EXPECT().IncUpdate()

			np1 := pr.Proposal{UserId: 42, LessonId: 24, DocumentId: 100}

			sqlmk.ExpectBegin()
			sqlmk.ExpectPrepare("INSERT INTO reaction.proposal").
				ExpectQuery().
				WithArgs(np1.UserId, np1.LessonId, np1.DocumentId).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			sqlmk.ExpectCommit()

			reqCreate := &pr.CreateProposalV1Request{Proposal: &np1}
			respCreate, err := grpcService.CreateProposalV1(ctx, reqCreate)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(respCreate).ShouldNot(BeNil())
			Ω(respCreate.ProposalId).Should(Equal(uint64(1)))

			np2 := pr.Proposal{ProposalId: 1, UserId: 10, LessonId: 20, DocumentId: 30}
			sqlmk.ExpectQuery("UPDATE reaction.proposal").
				WithArgs(np2.UserId, np2.LessonId, np2.DocumentId, np2.ProposalId).WillReturnRows(
				sqlmock.NewRows([]string{"id"}).AddRow(1),
			)

			reqUpdate := &pr.UpdateProposalV1Request{Proposal: &np2}
			respUpdate, err := grpcService.UpdateProposalV1(ctx, reqUpdate)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(respUpdate).ShouldNot(BeNil())

			returned := sqlmock.NewRows([]string{
				"id",
				"user_id",
				"lesson_id",
				"classroom_id"},
			).AddRow(1, np2.UserId, np2.LessonId, np2.DocumentId)

			sqlmk.ExpectQuery("SELECT id, user_id, lesson_id, document_id FROM reaction.proposal").
				WithArgs(respCreate.ProposalId).WillReturnRows(returned)

			// valid request
			reqDescribe := &pr.DescribeProposalV1Request{ProposalId: respCreate.ProposalId}
			respDescribe, err := grpcService.DescribeProposalV1(ctx, reqDescribe)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(respDescribe).ShouldNot(BeNil())
			Ω(respDescribe.Proposal.ProposalId).Should(Equal(uint64(1)))
			Ω(respDescribe.Proposal.UserId).Should(Equal(np2.UserId))
			Ω(respDescribe.Proposal.LessonId).Should(Equal(np2.LessonId))
			Ω(respDescribe.Proposal.DocumentId).Should(Equal(np2.DocumentId))

			if err := sqlmk.ExpectationsWereMet(); err != nil {
				controller.T.Errorf("there were unfulfilled expectations: %s", err)
			}

			// invalid reqest
			reqUpdate = &pr.UpdateProposalV1Request{}
			respUpdate, err = grpcService.UpdateProposalV1(ctx, reqUpdate)
			Ω(respUpdate).Should(BeNil())
			Ω(err).Should(HaveOccurred())
			cancel()
			p.Close()

		})
	})

})
