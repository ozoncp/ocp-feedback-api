package proposal_grpc_test

import (
	"context"
	"net"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-feedback-api/internal/repo"
	"github.com/ozoncp/ocp-feedback-api/internal/server/proposal_grpc"
	pr "github.com/ozoncp/ocp-feedback-api/pkg/ocp-proposal-api"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateProposal(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("sqlmock init failed", err)
	}
	defer db.Close()

	serverAddress := startTestGrpcServer(t,
		nil, repo.NewProposalRepo(sqlx.NewDb(db, "")), 2)
	client := newTestGrpcClient(t, serverAddress)

	// valid request
	np1 := pr.Proposal{UserId: 10, LessonId: 20, DocumentId: 30}

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO reaction.proposal").
		ExpectQuery().
		WithArgs(np1.UserId, np1.LessonId, np1.DocumentId).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	reqCreate := &pr.CreateProposalV1Request{Proposal: &np1}
	respCreate, err := client.CreateProposalV1(context.Background(), reqCreate)
	require.NoError(t, err)
	require.NotNil(t, respCreate)
	require.Equal(t, uint64(1), respCreate.ProposalId)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// invalid request, must fail on validation
	np2 := pr.Proposal{UserId: 0, LessonId: 20, DocumentId: 30}
	reqCreate = &pr.CreateProposalV1Request{Proposal: &np2}
	respCreate, err = client.CreateProposalV1(context.Background(), reqCreate)
	require.Error(t, err)
	require.Nil(t, respCreate)
}

func TestClientCreateMultiProposal(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("sqlmock init failed", err)
	}
	defer db.Close()

	serverAddress := startTestGrpcServer(t,
		nil, repo.NewProposalRepo(sqlx.NewDb(db, "")), 2)
	client := newTestGrpcClient(t, serverAddress)

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

	// assume that proposals will be split into 2 chunks of size 2
	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO reaction.proposal")
	mock.ExpectQuery("INSERT INTO reaction.proposal").
		WithArgs(pr1.UserId, pr1.LessonId, pr1.DocumentId).
		WillReturnRows(assignedNumbers[0])
	mock.ExpectQuery("INSERT INTO reaction.proposal").
		WithArgs(pr2.UserId, pr2.LessonId, pr2.DocumentId).
		WillReturnRows(assignedNumbers[1])
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO reaction.proposal")
	mock.ExpectQuery("INSERT INTO reaction.proposal").
		WithArgs(pr3.UserId, pr3.LessonId, pr3.DocumentId).
		WillReturnRows(assignedNumbers[2])
	mock.ExpectQuery("INSERT INTO reaction.proposal").
		WithArgs(pr4.UserId, pr4.LessonId, pr4.DocumentId).
		WillReturnRows(assignedNumbers[3])
	mock.ExpectCommit()

	respMultiCreate, err := client.CreateMultiProposalV1(context.Background(), reqMultiCreate)
	require.NoError(t, err)
	require.NotNil(t, respMultiCreate)
	require.Equal(t, []uint64{1, 2, 3, 4}, respMultiCreate.Proposals)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// invalid request
	np5 := pr.Proposal{UserId: 0, LessonId: 24, DocumentId: 100}
	reqMultiCreate = &pr.CreateMultiProposalV1Request{
		Proposals: []*pr.Proposal{
			&np5,
		},
	}

	respMultiCreate, err = client.CreateMultiProposalV1(context.Background(), reqMultiCreate)
	require.Error(t, err)
	require.Nil(t, respMultiCreate)
}

func TestClientRemoveProposal(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("sqlmock init failed", err)
	}
	defer db.Close()

	serverAddress := startTestGrpcServer(t,
		nil, repo.NewProposalRepo(sqlx.NewDb(db, "")), 2)
	client := newTestGrpcClient(t, serverAddress)

	np1 := pr.Proposal{UserId: 10, LessonId: 20, DocumentId: 30}

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO reaction.proposal").
		ExpectQuery().
		WithArgs(np1.UserId, np1.LessonId, np1.DocumentId).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	reqCreate := &pr.CreateProposalV1Request{Proposal: &np1}
	respCreate, err := client.CreateProposalV1(context.Background(), reqCreate)
	require.NoError(t, err)
	require.NotNil(t, respCreate)
	require.Equal(t, uint64(1), respCreate.ProposalId)

	mock.ExpectQuery("DELETE FROM reaction.proposal").
		WithArgs(respCreate.ProposalId).WillReturnRows(
		sqlmock.NewRows([]string{"id"}).AddRow(1))

	reqRemove := &pr.RemoveProposalV1Request{ProposalId: respCreate.ProposalId}
	respRemove, err := client.RemoveProposalV1(context.Background(), reqRemove)
	require.NoError(t, err)
	require.NotNil(t, respRemove)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// try to remove the it the second time
	reqRemove = &pr.RemoveProposalV1Request{ProposalId: respCreate.ProposalId}
	respRemove, err = client.RemoveProposalV1(context.Background(), reqRemove)
	require.Error(t, err)
	require.Nil(t, respRemove)

	// invalid reqest
	reqRemove = &pr.RemoveProposalV1Request{ProposalId: 0}
	respRemove, err = client.RemoveProposalV1(context.Background(), reqRemove)
	require.Error(t, err)
	require.Nil(t, respRemove)
}

func TestClientDescribeProposal(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("sqlmock init failed", err)
	}
	defer db.Close()

	serverAddress := startTestGrpcServer(t,
		nil, repo.NewProposalRepo(sqlx.NewDb(db, "")), 2)
	client := newTestGrpcClient(t, serverAddress)

	np := pr.Proposal{UserId: 10, LessonId: 20, DocumentId: 30}

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO reaction.proposal").
		ExpectQuery().
		WithArgs(np.UserId, np.LessonId, np.DocumentId).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	reqCreate := &pr.CreateProposalV1Request{Proposal: &np}
	respCreate, err1 := client.CreateProposalV1(context.Background(), reqCreate)
	require.NoError(t, err1)
	require.NotNil(t, respCreate)
	require.Equal(t, uint64(1), respCreate.ProposalId)

	returned := sqlmock.NewRows([]string{
		"id",
		"user_id",
		"lesson_id",
		"document_id"},
	).AddRow(1, np.UserId, np.LessonId, np.DocumentId)

	mock.ExpectQuery("SELECT id, user_id, lesson_id, document_id FROM reaction.proposal").
		WithArgs(respCreate.ProposalId).WillReturnRows(returned)

	// valid request
	reqDescribe := &pr.DescribeProposalV1Request{ProposalId: respCreate.ProposalId}
	respDescribe, err := client.DescribeProposalV1(context.Background(), reqDescribe)
	require.NoError(t, err)
	require.NotNil(t, respDescribe)
	require.Equal(t, respDescribe.Proposal.ProposalId, uint64(1))
	require.Equal(t, respDescribe.Proposal.UserId, np.UserId)
	require.Equal(t, respDescribe.Proposal.LessonId, np.LessonId)
	require.Equal(t, respDescribe.Proposal.DocumentId, np.DocumentId)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// missing id
	reqDescribe = &pr.DescribeProposalV1Request{ProposalId: respCreate.ProposalId + 1}
	respDescribe, err = client.DescribeProposalV1(context.Background(), reqDescribe)
	require.Error(t, err)
	require.Nil(t, respDescribe)

	// invalid request
	reqDescribe = &pr.DescribeProposalV1Request{ProposalId: 0}
	respDescribe, err = client.DescribeProposalV1(context.Background(), reqDescribe)
	require.Error(t, err)
	require.Nil(t, respDescribe)
}

func TestClientListProposal(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("sqlmock init failed", err)
	}
	defer db.Close()

	serverAddress := startTestGrpcServer(t,
		nil, repo.NewProposalRepo(sqlx.NewDb(db, "")), 1)
	client := newTestGrpcClient(t, serverAddress)

	np1 := pr.Proposal{UserId: 42, LessonId: 24, DocumentId: 50}
	np2 := pr.Proposal{UserId: 420, LessonId: 240, DocumentId: 500}

	reqCreate := &pr.CreateMultiProposalV1Request{Proposals: []*pr.Proposal{
		&np1,
		&np2,
	},
	}

	// assume that feedbacks won't be split into chunks
	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO reaction.proposal")
	mock.ExpectQuery("INSERT INTO reaction.proposal").
		WithArgs(np1.UserId, np1.LessonId, np1.DocumentId).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectQuery("INSERT INTO reaction.proposal").
		WithArgs(np2.UserId, np2.LessonId, np2.DocumentId).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))
	mock.ExpectCommit()

	respMultiCreate, err := client.CreateMultiProposalV1(context.Background(), reqCreate)
	require.NoError(t, err)
	require.NotNil(t, respMultiCreate)

	returned := sqlmock.NewRows([]string{
		"id",
		"user_id",
		"lesson_id",
		"document_id"},
	).AddRow(1, np1.UserId, np1.LessonId, np1.DocumentId).AddRow(
		2, np2.UserId, np2.LessonId, np2.DocumentId,
	)

	mock.ExpectQuery("SELECT id, user_id, lesson_id, document_id FROM reaction.proposal").
		WithArgs(2, 0).WillReturnRows(returned)

	// valid request
	reqList := &pr.ListProposalsV1Request{Limit: 2, Offset: 0}
	respList, err := client.ListProposalsV1(context.Background(), reqList)
	require.NoError(t, err)
	require.NotNil(t, respList)
	require.Equal(t, len(respList.Proposals), 2)

	require.Equal(t, respList.Proposals[0].ProposalId, uint64(1))
	require.Equal(t, respList.Proposals[0].UserId, np1.UserId)
	require.Equal(t, respList.Proposals[0].LessonId, np1.LessonId)
	require.Equal(t, respList.Proposals[0].DocumentId, np1.DocumentId)

	require.Equal(t, respList.Proposals[1].ProposalId, uint64(2))
	require.Equal(t, respList.Proposals[1].UserId, np2.UserId)
	require.Equal(t, respList.Proposals[1].LessonId, np2.LessonId)
	require.Equal(t, respList.Proposals[1].DocumentId, np2.DocumentId)

	mock.ExpectQuery("SELECT id, user_id, lesson_id, document_id FROM reaction.proposal").
		WithArgs(1, 4).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"user_id",
		"lesson_id",
		"document_id"}))

	// wrong offset
	reqList = &pr.ListProposalsV1Request{Limit: 1, Offset: 4}
	respList, err = client.ListProposalsV1(context.Background(), reqList)
	require.NoError(t, err)
	require.NotNil(t, respList)
	require.Equal(t, len(respList.Proposals), 0)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// invalid request
	reqList = &pr.ListProposalsV1Request{Limit: 0, Offset: 1}
	respList, err = client.ListProposalsV1(context.Background(), reqList)
	require.Error(t, err)
	require.Nil(t, respList)
}

func TestClientUpdateProposal(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("sqlmock init failed", err)
	}
	defer db.Close()

	serverAddress := startTestGrpcServer(t,
		nil, repo.NewProposalRepo(sqlx.NewDb(db, "")), 2)
	client := newTestGrpcClient(t, serverAddress)

	np1 := pr.Proposal{UserId: 42, LessonId: 24, DocumentId: 100}

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO reaction.proposal").
		ExpectQuery().
		WithArgs(np1.UserId, np1.LessonId, np1.DocumentId).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	reqCreate := &pr.CreateProposalV1Request{Proposal: &np1}
	respCreate, err := client.CreateProposalV1(context.Background(), reqCreate)
	require.NoError(t, err)
	require.NotNil(t, respCreate)
	require.Equal(t, uint64(1), respCreate.ProposalId)

	np2 := pr.Proposal{ProposalId: 1, UserId: 10, LessonId: 20, DocumentId: 30}
	mock.ExpectQuery("UPDATE reaction.proposal").
		WithArgs(np2.UserId, np2.LessonId, np2.DocumentId, np2.ProposalId).WillReturnRows(
		sqlmock.NewRows([]string{"id"}).AddRow(1),
	)

	reqUpdate := &pr.UpdateProposalV1Request{Proposal: &np2}
	respUpdate, err := client.UpdateProposalV1(context.Background(), reqUpdate)
	require.NoError(t, err)
	require.NotNil(t, respUpdate)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// invalid reqest
	reqUpdate = &pr.UpdateProposalV1Request{}
	respUpdate, err = client.UpdateProposalV1(context.Background(), reqUpdate)
	require.Error(t, err)
	require.Nil(t, respUpdate)
}

func startTestGrpcServer(t *testing.T,
	feedbackRepo repo.Repo,
	proposalRepo repo.Repo,
	chunks int,
) string {
	zerolog.SetGlobalLevel(zerolog.Disabled)

	service := proposal_grpc.New(proposalRepo, chunks)
	grpcServer := grpc.NewServer()
	pr.RegisterOcpProposalApiServer(grpcServer, service)
	listener, err := net.Listen("tcp", ":0") // random available port
	require.NoError(t, err)
	go func() {
		err := grpcServer.Serve(listener)
		require.NoError(t, err)
	}()
	return listener.Addr().String()
}

func newTestGrpcClient(t *testing.T, serverAddress string) pr.OcpProposalApiClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	return pr.NewOcpProposalApiClient(conn)
}
