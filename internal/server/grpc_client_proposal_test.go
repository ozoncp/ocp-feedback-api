package grpc_service_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-feedback-api/internal/repo"
	fb "github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api"
	"github.com/stretchr/testify/require"
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
	np1 := fb.Proposal{UserId: 10, LessonId: 20, DocumentId: 30}

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO reaction.proposal").
		ExpectQuery().
		WithArgs(np1.UserId, np1.LessonId, np1.DocumentId).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	reqCreate := &fb.CreateProposalV1Request{Proposal: &np1}
	respCreate, err := client.CreateProposalV1(context.Background(), reqCreate)
	require.NoError(t, err)
	require.NotNil(t, respCreate)
	require.Equal(t, uint64(1), respCreate.Proposal)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// invalid request, must fail on validation
	np2 := fb.Proposal{UserId: 0, LessonId: 20, DocumentId: 30}
	reqCreate = &fb.CreateProposalV1Request{Proposal: &np2}
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
	pr1 := fb.Proposal{UserId: 42, LessonId: 24, DocumentId: 50}
	pr2 := fb.Proposal{UserId: 420, LessonId: 240, DocumentId: 500}
	pr3 := fb.Proposal{UserId: 4200, LessonId: 2400, DocumentId: 5000}
	pr4 := fb.Proposal{UserId: 42000, LessonId: 24000, DocumentId: 50000}

	reqMultiCreate := &fb.CreateMultiProposalV1Request{
		Proposals: []*fb.Proposal{
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
	np5 := fb.Proposal{UserId: 0, LessonId: 24, DocumentId: 100}
	reqMultiCreate = &fb.CreateMultiProposalV1Request{
		Proposals: []*fb.Proposal{
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

	np1 := fb.Proposal{UserId: 10, LessonId: 20, DocumentId: 30}

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO reaction.proposal").
		ExpectQuery().
		WithArgs(np1.UserId, np1.LessonId, np1.DocumentId).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	reqCreate := &fb.CreateProposalV1Request{Proposal: &np1}
	respCreate, err := client.CreateProposalV1(context.Background(), reqCreate)
	require.NoError(t, err)
	require.NotNil(t, respCreate)
	require.Equal(t, uint64(1), respCreate.Proposal)

	mock.ExpectQuery("SELECT 1 FROM reaction.proposal").
		WithArgs(respCreate.Proposal).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectExec("DELETE FROM reaction.proposal").
		WithArgs(respCreate.Proposal).WillReturnResult(sqlmock.NewResult(1, 1))

	reqRemove := &fb.RemoveProposalV1Request{Proposal: respCreate.Proposal}
	respRemove, err := client.RemoveProposalV1(context.Background(), reqRemove)
	require.NoError(t, err)
	require.NotNil(t, respRemove)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// try to remove the it the second time
	reqRemove = &fb.RemoveProposalV1Request{Proposal: respCreate.Proposal}
	respRemove, err = client.RemoveProposalV1(context.Background(), reqRemove)
	require.Error(t, err)
	require.Nil(t, respRemove)

	// invalid reqest
	reqRemove = &fb.RemoveProposalV1Request{Proposal: 0}
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

	np := fb.Proposal{UserId: 10, LessonId: 20, DocumentId: 30}

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO reaction.proposal").
		ExpectQuery().
		WithArgs(np.UserId, np.LessonId, np.DocumentId).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	reqCreate := &fb.CreateProposalV1Request{Proposal: &np}
	respCreate, err1 := client.CreateProposalV1(context.Background(), reqCreate)
	require.NoError(t, err1)
	require.NotNil(t, respCreate)
	require.Equal(t, uint64(1), respCreate.Proposal)

	returned := sqlmock.NewRows([]string{
		"id",
		"user_id",
		"lesson_id",
		"document_id"},
	).AddRow(1, np.UserId, np.LessonId, np.DocumentId)

	mock.ExpectQuery("SELECT id, user_id, lesson_id, document_id FROM reaction.proposal").
		WithArgs(respCreate.Proposal).WillReturnRows(returned)

	// valid request
	reqDescribe := &fb.DescribeProposalV1Request{Proposal: respCreate.Proposal}
	respDescribe, err := client.DescribeProposalV1(context.Background(), reqDescribe)
	require.NoError(t, err)
	require.NotNil(t, respDescribe)
	require.Equal(t, respDescribe.Proposal.Id, uint64(1))
	require.Equal(t, respDescribe.Proposal.UserId, np.UserId)
	require.Equal(t, respDescribe.Proposal.LessonId, np.LessonId)
	require.Equal(t, respDescribe.Proposal.DocumentId, np.DocumentId)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// missing id
	reqDescribe = &fb.DescribeProposalV1Request{Proposal: respCreate.Proposal + 1}
	respDescribe, err = client.DescribeProposalV1(context.Background(), reqDescribe)
	require.Error(t, err)
	require.Nil(t, respDescribe)

	// invalid request
	reqDescribe = &fb.DescribeProposalV1Request{Proposal: 0}
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

	np1 := fb.Proposal{UserId: 42, LessonId: 24, DocumentId: 50}
	np2 := fb.Proposal{UserId: 420, LessonId: 240, DocumentId: 500}

	reqCreate := &fb.CreateMultiProposalV1Request{Proposals: []*fb.Proposal{
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
	reqList := &fb.ListProposalsV1Request{Limit: 2, Offset: 0}
	respList, err := client.ListProposalsV1(context.Background(), reqList)
	require.NoError(t, err)
	require.NotNil(t, respList)
	require.Equal(t, len(respList.Proposals), 2)

	require.Equal(t, respList.Proposals[0].Id, uint64(1))
	require.Equal(t, respList.Proposals[0].UserId, np1.UserId)
	require.Equal(t, respList.Proposals[0].LessonId, np1.LessonId)
	require.Equal(t, respList.Proposals[0].DocumentId, np1.DocumentId)

	require.Equal(t, respList.Proposals[1].Id, uint64(2))
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
	reqList = &fb.ListProposalsV1Request{Limit: 1, Offset: 4}
	respList, err = client.ListProposalsV1(context.Background(), reqList)
	require.NoError(t, err)
	require.NotNil(t, respList)
	require.Equal(t, len(respList.Proposals), 0)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// invalid request
	reqList = &fb.ListProposalsV1Request{Limit: 0, Offset: 1}
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

	np1 := fb.Proposal{UserId: 42, LessonId: 24, DocumentId: 100}

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO reaction.proposal").
		ExpectQuery().
		WithArgs(np1.UserId, np1.LessonId, np1.DocumentId).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	reqCreate := &fb.CreateProposalV1Request{Proposal: &np1}
	respCreate, err := client.CreateProposalV1(context.Background(), reqCreate)
	require.NoError(t, err)
	require.NotNil(t, respCreate)
	require.Equal(t, uint64(1), respCreate.Proposal)

	mock.ExpectQuery("SELECT 1 FROM reaction.proposal").
		WithArgs(respCreate.Proposal).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	np2 := fb.Proposal{Id: 1, UserId: 10, LessonId: 20, DocumentId: 30}
	mock.ExpectExec("UPDATE reaction.proposal").
		WithArgs(np2.UserId, np2.LessonId, np2.DocumentId, np2.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	reqUpdate := &fb.UpdateProposalV1Request{Proposal: &np2}
	respUpdate, err := client.UpdateProposalV1(context.Background(), reqUpdate)
	require.NoError(t, err)
	require.NotNil(t, respUpdate)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// invalid reqest
	reqUpdate = &fb.UpdateProposalV1Request{}
	respUpdate, err = client.UpdateProposalV1(context.Background(), reqUpdate)
	require.Error(t, err)
	require.Nil(t, respUpdate)
}
