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

func TestClientCreateFeedback(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("sqlmock init failed", err)
	}
	defer db.Close()

	serverAddress := startTestGrpcServer(t,
		repo.NewFeedbackRepo(sqlx.NewDb(db, "")), nil, 2)
	client := newTestGrpcClient(t, serverAddress)

	// valid request
	nf1 := fb.NewFeedback{UserId: 42, ClassroomId: 24, Comment: "hello1"}

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO reaction.feedback").
		ExpectQuery().
		WithArgs(nf1.UserId, nf1.ClassroomId, nf1.Comment).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	reqCreate := &fb.CreateFeedbackV1Request{NewFeedback: &nf1}
	respCreate, err := client.CreateFeedbackV1(context.Background(), reqCreate)
	require.NoError(t, err)
	require.NotNil(t, respCreate)
	require.Equal(t, uint64(1), respCreate.FeedbackId)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// invalid request, must fail on validation
	nf2 := fb.NewFeedback{UserId: 0, ClassroomId: 240, Comment: "hello2"}
	reqCreate = &fb.CreateFeedbackV1Request{NewFeedback: &nf2}
	respCreate, err = client.CreateFeedbackV1(context.Background(), reqCreate)
	require.Error(t, err)
	require.Nil(t, respCreate)
}

func TestClientCreateMultiFeedback(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("sqlmock init failed", err)
	}
	defer db.Close()

	serverAddress := startTestGrpcServer(t,
		repo.NewFeedbackRepo(sqlx.NewDb(db, "")), nil, 2)
	client := newTestGrpcClient(t, serverAddress)

	// valid request
	nf1 := fb.NewFeedback{UserId: 42, ClassroomId: 24, Comment: "hello1"}
	nf2 := fb.NewFeedback{UserId: 420, ClassroomId: 240, Comment: "hello2"}
	nf3 := fb.NewFeedback{UserId: 4200, ClassroomId: 2400, Comment: "hello3"}
	nf4 := fb.NewFeedback{UserId: 42000, ClassroomId: 24000, Comment: "hello4"}

	reqMultiCreate := &fb.CreateMultiFeedbackV1Request{
		NewFeedbacks: []*fb.NewFeedback{
			&nf1,
			&nf2,
			&nf3,
			&nf4,
		},
	}

	// sequence numbers
	assignedNumbers := []*sqlmock.Rows{}
	for i := 1; i <= len(reqMultiCreate.NewFeedbacks); i++ {
		assignedNumbers = append(assignedNumbers, sqlmock.NewRows([]string{"id"}).AddRow(i))
	}

	// assume that feedbacks will be split into 2 chunks of size 2
	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO reaction.feedback")
	mock.ExpectQuery("INSERT INTO reaction.feedback").
		WithArgs(nf1.UserId, nf1.ClassroomId, nf1.Comment).
		WillReturnRows(assignedNumbers[0])
	mock.ExpectQuery("INSERT INTO reaction.feedback").
		WithArgs(nf2.UserId, nf2.ClassroomId, nf2.Comment).
		WillReturnRows(assignedNumbers[1])
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO reaction.feedback")
	mock.ExpectQuery("INSERT INTO reaction.feedback").
		WithArgs(nf3.UserId, nf3.ClassroomId, nf3.Comment).
		WillReturnRows(assignedNumbers[2])
	mock.ExpectQuery("INSERT INTO reaction.feedback").
		WithArgs(nf4.UserId, nf4.ClassroomId, nf4.Comment).
		WillReturnRows(assignedNumbers[3])
	mock.ExpectCommit()

	respMultiCreate, err := client.CreateMultiFeedbackV1(context.Background(), reqMultiCreate)
	require.NoError(t, err)
	require.NotNil(t, respMultiCreate)
	require.Equal(t, []uint64{1, 2, 3, 4}, respMultiCreate.FeedbackIds)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// invalid request
	nf5 := fb.NewFeedback{UserId: 0, ClassroomId: 24, Comment: "hello1"}
	reqMultiCreate = &fb.CreateMultiFeedbackV1Request{
		NewFeedbacks: []*fb.NewFeedback{
			&nf5,
		},
	}

	respMultiCreate, err = client.CreateMultiFeedbackV1(context.Background(), reqMultiCreate)
	require.Error(t, err)
	require.Nil(t, respMultiCreate)
}

func TestClientRemoveFeedback(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("sqlmock init failed", err)
	}
	defer db.Close()

	serverAddress := startTestGrpcServer(t,
		repo.NewFeedbackRepo(sqlx.NewDb(db, "")), nil, 2)
	client := newTestGrpcClient(t, serverAddress)

	nf1 := fb.NewFeedback{UserId: 42, ClassroomId: 24, Comment: "hello1"}

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO reaction.feedback").
		ExpectQuery().
		WithArgs(nf1.UserId, nf1.ClassroomId, nf1.Comment).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	reqCreate := &fb.CreateFeedbackV1Request{NewFeedback: &nf1}
	respCreate, err := client.CreateFeedbackV1(context.Background(), reqCreate)
	require.NoError(t, err)
	require.NotNil(t, respCreate)
	require.Equal(t, uint64(1), respCreate.FeedbackId)

	mock.ExpectQuery("SELECT 1 FROM reaction.feedback").
		WithArgs(respCreate.FeedbackId).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectExec("DELETE FROM reaction.feedback").
		WithArgs(respCreate.FeedbackId).WillReturnResult(sqlmock.NewResult(1, 1))

	reqRemove := &fb.RemoveFeedbackV1Request{FeedbackId: respCreate.FeedbackId}
	respRemove, err := client.RemoveFeedbackV1(context.Background(), reqRemove)

	require.NoError(t, err)
	require.NotNil(t, respRemove)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// try to remove the it the second time
	reqRemove = &fb.RemoveFeedbackV1Request{FeedbackId: respCreate.FeedbackId}
	respRemove, err = client.RemoveFeedbackV1(context.Background(), reqRemove)
	require.Error(t, err)
	require.Nil(t, respRemove)

	// invalid reqest
	reqRemove = &fb.RemoveFeedbackV1Request{FeedbackId: 0}
	respRemove, err = client.RemoveFeedbackV1(context.Background(), reqRemove)
	require.Error(t, err)
	require.Nil(t, respRemove)
}

func TestClientDescribeFeedback(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("sqlmock init failed", err)
	}
	defer db.Close()

	serverAddress := startTestGrpcServer(t,
		repo.NewFeedbackRepo(sqlx.NewDb(db, "")), nil, 2)
	client := newTestGrpcClient(t, serverAddress)

	nf := fb.NewFeedback{UserId: 42, ClassroomId: 24, Comment: "hello1"}

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO reaction.feedback").
		ExpectQuery().
		WithArgs(nf.UserId, nf.ClassroomId, nf.Comment).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	reqCreate := &fb.CreateFeedbackV1Request{NewFeedback: &nf}
	respCreate, err1 := client.CreateFeedbackV1(context.Background(), reqCreate)
	require.NoError(t, err1)
	require.NotNil(t, respCreate)
	require.Equal(t, uint64(1), respCreate.FeedbackId)

	returned := sqlmock.NewRows([]string{
		"id",
		"user_id",
		"classroom_id",
		"comment"},
	).AddRow(1, nf.UserId, nf.ClassroomId, nf.Comment)

	mock.ExpectQuery("SELECT id, user_id, classroom_id, comment FROM reaction.feedback").
		WithArgs(respCreate.FeedbackId).WillReturnRows(returned)

	// valid request
	reqDescribe := &fb.DescribeFeedbackV1Request{FeedbackId: respCreate.FeedbackId}
	respDescribe, err := client.DescribeFeedbackV1(context.Background(), reqDescribe)
	require.NoError(t, err)
	require.NotNil(t, respDescribe)
	require.Equal(t, respDescribe.Feedback.FeedbackId, uint64(1))
	require.Equal(t, respDescribe.Feedback.UserId, nf.UserId)
	require.Equal(t, respDescribe.Feedback.ClassroomId, nf.ClassroomId)
	require.Equal(t, respDescribe.Feedback.Comment, nf.Comment)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// missing id
	reqDescribe = &fb.DescribeFeedbackV1Request{FeedbackId: respCreate.FeedbackId + 1}
	respDescribe, err = client.DescribeFeedbackV1(context.Background(), reqDescribe)
	require.Error(t, err)
	require.Nil(t, respDescribe)

	// invalid request
	reqDescribe = &fb.DescribeFeedbackV1Request{FeedbackId: 0}
	respDescribe, err = client.DescribeFeedbackV1(context.Background(), reqDescribe)
	require.Error(t, err)
	require.Nil(t, respDescribe)
}

func TestClientListFeedbacks(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("sqlmock init failed", err)
	}
	defer db.Close()

	serverAddress := startTestGrpcServer(t,
		repo.NewFeedbackRepo(sqlx.NewDb(db, "")), nil, 1)
	client := newTestGrpcClient(t, serverAddress)

	nf1 := fb.NewFeedback{UserId: 42, ClassroomId: 24, Comment: "hello1"}
	nf2 := fb.NewFeedback{UserId: 420, ClassroomId: 240, Comment: "hello2"}

	reqCreate := &fb.CreateMultiFeedbackV1Request{NewFeedbacks: []*fb.NewFeedback{
		&nf1,
		&nf2,
	},
	}

	// assume that feedbacks won't be split into chunks
	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO reaction.feedback")
	mock.ExpectQuery("INSERT INTO reaction.feedback").
		WithArgs(nf1.UserId, nf1.ClassroomId, nf1.Comment).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectQuery("INSERT INTO reaction.feedback").
		WithArgs(nf2.UserId, nf2.ClassroomId, nf2.Comment).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))
	mock.ExpectCommit()

	respMultiCreate, err := client.CreateMultiFeedbackV1(context.Background(), reqCreate)
	require.NoError(t, err)
	require.NotNil(t, respMultiCreate)

	returned := sqlmock.NewRows([]string{
		"id",
		"user_id",
		"classroom_id",
		"comment"},
	).AddRow(1, nf1.UserId, nf1.ClassroomId, nf1.Comment).AddRow(
		2, nf2.UserId, nf2.ClassroomId, nf2.Comment,
	)

	mock.ExpectQuery("SELECT id, user_id, classroom_id, comment FROM reaction.feedback").
		WithArgs(2, 0).WillReturnRows(returned)

	// valid request
	reqList := &fb.ListFeedbacksV1Request{Limit: 2, Offset: 0}
	respList, err := client.ListFeedbacksV1(context.Background(), reqList)
	require.NoError(t, err)
	require.NotNil(t, respList)
	require.Equal(t, len(respList.Feedbacks), 2)

	require.Equal(t, respList.Feedbacks[0].FeedbackId, uint64(1))
	require.Equal(t, respList.Feedbacks[0].UserId, nf1.UserId)
	require.Equal(t, respList.Feedbacks[0].ClassroomId, nf1.ClassroomId)
	require.Equal(t, respList.Feedbacks[0].Comment, nf1.Comment)

	require.Equal(t, respList.Feedbacks[1].FeedbackId, uint64(2))
	require.Equal(t, respList.Feedbacks[1].UserId, nf2.UserId)
	require.Equal(t, respList.Feedbacks[1].ClassroomId, nf2.ClassroomId)
	require.Equal(t, respList.Feedbacks[1].Comment, nf2.Comment)

	mock.ExpectQuery("SELECT id, user_id, classroom_id, comment FROM reaction.feedback").
		WithArgs(1, 4).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"user_id",
		"classroom_id",
		"comment"}))

	// wrong offset
	reqList = &fb.ListFeedbacksV1Request{Limit: 1, Offset: 4}
	respList, err = client.ListFeedbacksV1(context.Background(), reqList)
	require.NoError(t, err)
	require.NotNil(t, respList)
	require.Equal(t, len(respList.Feedbacks), 0)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// invalid request
	reqList = &fb.ListFeedbacksV1Request{Limit: 0, Offset: 1}
	respList, err = client.ListFeedbacksV1(context.Background(), reqList)
	require.Error(t, err)
	require.Nil(t, respList)
}
