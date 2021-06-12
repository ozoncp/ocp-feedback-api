package grpc_service_test

import (
	"context"
	"testing"

	"github.com/ozoncp/ocp-feedback-api/internal/repo"
	fb "github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api"
	"github.com/stretchr/testify/require"
)

func TestClientCreateFeedback(t *testing.T) {
	t.Parallel()

	serverAddress := startTestGrpcServer(t,
		repo.NewInMemoryFeedbackRepo(), nil)
	client := newTestGrpcClient(t, serverAddress)

	// valid request
	newFeedback1 := fb.NewFeedback{UserId: 42, ClassroomId: 24, Comment: "hello1"}
	req1 := &fb.CreateFeedbackV1Request{NewFeedback: &newFeedback1}
	resp1, err1 := client.CreateFeedbackV1(context.Background(), req1)
	require.NoError(t, err1)
	require.NotNil(t, resp1)
	require.Equal(t, uint64(1), resp1.FeedbackId)

	// invalid request
	newFeedback2 := fb.NewFeedback{UserId: 0, ClassroomId: 240, Comment: "hello2"}
	req2 := &fb.CreateFeedbackV1Request{NewFeedback: &newFeedback2}
	resp2, err2 := client.CreateFeedbackV1(context.Background(), req2)
	require.Error(t, err2)
	require.Nil(t, resp2)
}

func TestClientCreateMultiFeedback(t *testing.T) {
	t.Parallel()

	serverAddress := startTestGrpcServer(t,
		repo.NewInMemoryFeedbackRepo(), nil)
	client := newTestGrpcClient(t, serverAddress)

	// valid request
	newFeedback1 := fb.NewFeedback{UserId: 42, ClassroomId: 24, Comment: "hello1"}
	newFeedback2 := fb.NewFeedback{UserId: 420, ClassroomId: 240, Comment: "hello2"}
	newFeedback3 := fb.NewFeedback{UserId: 4200, ClassroomId: 2400, Comment: "hello3"}
	newFeedback4 := fb.NewFeedback{UserId: 42000, ClassroomId: 24000, Comment: "hello4"}

	req1 := &fb.CreateMultiFeedbackV1Request{NewFeedbacks: []*fb.NewFeedback{
		&newFeedback1,
		&newFeedback2,
		&newFeedback3,
		&newFeedback4,
	},
	}

	resp, err := client.CreateMultiFeedbackV1(context.Background(), req1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, []uint64{1, 2, 3, 4}, resp.FeedbackIds)

	// invalid request
	newFeedback5 := fb.NewFeedback{UserId: 0, ClassroomId: 24, Comment: "hello1"}
	req2 := &fb.CreateMultiFeedbackV1Request{NewFeedbacks: []*fb.NewFeedback{
		&newFeedback5,
	},
	}

	resp2, err := client.CreateMultiFeedbackV1(context.Background(), req2)
	require.Error(t, err)
	require.Nil(t, resp2)
}

func TestClientRemoveFeedback(t *testing.T) {
	t.Parallel()

	serverAddress := startTestGrpcServer(t,
		repo.NewInMemoryFeedbackRepo(), nil)
	client := newTestGrpcClient(t, serverAddress)

	newFeedback1 := fb.NewFeedback{UserId: 42, ClassroomId: 24, Comment: "hello1"}
	newFeedback2 := fb.NewFeedback{UserId: 420, ClassroomId: 240, Comment: "hello2"}
	newFeedback3 := fb.NewFeedback{UserId: 4200, ClassroomId: 2400, Comment: "hello3"}
	newFeedback4 := fb.NewFeedback{UserId: 42000, ClassroomId: 24000, Comment: "hello4"}

	reqCreate := &fb.CreateMultiFeedbackV1Request{NewFeedbacks: []*fb.NewFeedback{
		&newFeedback1,
		&newFeedback2,
		&newFeedback3,
		&newFeedback4,
	},
	}

	respCreate, err := client.CreateMultiFeedbackV1(context.Background(), reqCreate)
	require.NoError(t, err)
	require.NotNil(t, respCreate)

	// remove the second newFeedback2 record
	reqRemove1 := &fb.RemoveFeedbackV1Request{FeedbackId: respCreate.FeedbackIds[1]}
	resRemove1, err := client.RemoveFeedbackV1(context.Background(), reqRemove1)

	require.NoError(t, err)
	require.NotNil(t, resRemove1)

	// try to remove the second time
	reqRemove2 := &fb.RemoveFeedbackV1Request{FeedbackId: respCreate.FeedbackIds[1]}
	respRemove2, err := client.RemoveFeedbackV1(context.Background(), reqRemove2)
	require.Error(t, err)
	require.Nil(t, respRemove2)

	// invalid reqest
	reqRemove3 := &fb.RemoveFeedbackV1Request{FeedbackId: 0}
	respRemove3, err := client.RemoveFeedbackV1(context.Background(), reqRemove3)
	require.Error(t, err)
	require.Nil(t, respRemove3)

}

func TestClientDescribeFeedback(t *testing.T) {
	t.Parallel()

	serverAddress := startTestGrpcServer(t,
		repo.NewInMemoryFeedbackRepo(), nil)
	client := newTestGrpcClient(t, serverAddress)

	newFeedback1 := fb.NewFeedback{UserId: 42, ClassroomId: 24, Comment: "hello1"}
	req1 := &fb.CreateFeedbackV1Request{NewFeedback: &newFeedback1}
	resp1, err1 := client.CreateFeedbackV1(context.Background(), req1)
	require.NoError(t, err1)
	require.NotNil(t, resp1)

	// valid request
	reqDescribe1 := &fb.DescribeFeedbackV1Request{FeedbackId: resp1.FeedbackId}
	respDescribe1, err := client.DescribeFeedbackV1(context.Background(), reqDescribe1)
	require.NoError(t, err)
	require.NotNil(t, respDescribe1)
	require.Equal(t, respDescribe1.Feedback.FeedbackId, uint64(1))
	require.Equal(t, respDescribe1.Feedback.UserId, newFeedback1.UserId)
	require.Equal(t, respDescribe1.Feedback.ClassroomId, newFeedback1.ClassroomId)
	require.Equal(t, respDescribe1.Feedback.Comment, newFeedback1.Comment)

	// missing id
	reqDescribe2 := &fb.DescribeFeedbackV1Request{FeedbackId: resp1.FeedbackId + 1}
	respDescribe2, err := client.DescribeFeedbackV1(context.Background(), reqDescribe2)
	require.Error(t, err)
	require.Nil(t, respDescribe2)

	// invalid request
	reqDescribe3 := &fb.DescribeFeedbackV1Request{FeedbackId: 0}
	respDescribe3, err := client.DescribeFeedbackV1(context.Background(), reqDescribe3)
	require.Error(t, err)
	require.Nil(t, respDescribe3)
}

func TestClientListFeedbacks(t *testing.T) {
	t.Parallel()

	serverAddress := startTestGrpcServer(t,
		repo.NewInMemoryFeedbackRepo(), nil)
	client := newTestGrpcClient(t, serverAddress)

	newFeedback1 := fb.NewFeedback{UserId: 42, ClassroomId: 24, Comment: "hello1"}
	newFeedback2 := fb.NewFeedback{UserId: 420, ClassroomId: 240, Comment: "hello2"}
	newFeedback3 := fb.NewFeedback{UserId: 4200, ClassroomId: 2400, Comment: "hello3"}
	newFeedback4 := fb.NewFeedback{UserId: 42000, ClassroomId: 24000, Comment: "hello4"}

	req1 := &fb.CreateMultiFeedbackV1Request{NewFeedbacks: []*fb.NewFeedback{
		&newFeedback1,
		&newFeedback2,
		&newFeedback3,
		&newFeedback4,
	},
	}
	resp, err := client.CreateMultiFeedbackV1(context.Background(), req1)
	require.NoError(t, err)
	require.NotNil(t, resp)

	// valid request
	reqList1 := &fb.ListFeedbacksV1Request{Limit: 2, Offset: 1}
	respList1, err := client.ListFeedbacksV1(context.Background(), reqList1)
	require.NoError(t, err)
	require.NotNil(t, respList1)
	require.Equal(t, len(respList1.Feedbacks), 2)

	require.Equal(t, respList1.Feedbacks[0].FeedbackId, uint64(2))
	require.Equal(t, respList1.Feedbacks[0].UserId, newFeedback2.UserId)
	require.Equal(t, respList1.Feedbacks[0].ClassroomId, newFeedback2.ClassroomId)
	require.Equal(t, respList1.Feedbacks[0].Comment, newFeedback2.Comment)

	require.Equal(t, respList1.Feedbacks[1].FeedbackId, uint64(3))
	require.Equal(t, respList1.Feedbacks[1].UserId, newFeedback3.UserId)
	require.Equal(t, respList1.Feedbacks[1].ClassroomId, newFeedback3.ClassroomId)
	require.Equal(t, respList1.Feedbacks[1].Comment, newFeedback3.Comment)

	// wrong offset
	reqList2 := &fb.ListFeedbacksV1Request{Limit: 1, Offset: 4}
	respList2, err := client.ListFeedbacksV1(context.Background(), reqList2)
	require.NoError(t, err)
	require.NotNil(t, respList2)
	require.Equal(t, len(respList2.Feedbacks), 0)

	// invalid request
	reqList3 := &fb.ListFeedbacksV1Request{Limit: 0, Offset: 1}
	respList3, err := client.ListFeedbacksV1(context.Background(), reqList3)
	require.Error(t, err)
	require.Nil(t, respList3)
}
