package grpc_service_test

import (
	"context"
	"testing"

	"github.com/ozoncp/ocp-feedback-api/internal/repo"
	fb "github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api"
	"github.com/stretchr/testify/require"
)

func TestClientCreateProposal(t *testing.T) {
	t.Parallel()

	serverAddress := startTestGrpcServer(t,
		nil, repo.NewInMemoryProposalRepo())
	client := newTestGrpcClient(t, serverAddress)

	// valid request
	newProposal1 := fb.NewProposal{UserId: 42, LessonId: 24, DocumentId: 7}
	req1 := &fb.CreateProposalV1Request{NewProposal: &newProposal1}
	res1, err1 := client.CreateProposalV1(context.Background(), req1)
	require.NoError(t, err1)
	require.NotNil(t, res1)
	require.Equal(t, uint64(1), res1.ProposalId)

	// invalid request
	newProposal2 := fb.NewProposal{UserId: 0, LessonId: 24, DocumentId: 7}
	req2 := &fb.CreateProposalV1Request{NewProposal: &newProposal2}
	resp2, err2 := client.CreateProposalV1(context.Background(), req2)
	require.Error(t, err2)
	require.Nil(t, resp2)
}

func TestClientCreateMultiProposal(t *testing.T) {
	t.Parallel()

	serverAddress := startTestGrpcServer(t,
		nil, repo.NewInMemoryProposalRepo())
	client := newTestGrpcClient(t, serverAddress)

	newProposal1 := fb.NewProposal{UserId: 42, LessonId: 24, DocumentId: 7}
	newProposal2 := fb.NewProposal{UserId: 420, LessonId: 240, DocumentId: 70}
	newProposal3 := fb.NewProposal{UserId: 4200, LessonId: 2400, DocumentId: 700}
	newProposal4 := fb.NewProposal{UserId: 42000, LessonId: 24000, DocumentId: 7000}

	req1 := &fb.CreateMultiProposalV1Request{NewProposals: []*fb.NewProposal{
		&newProposal1,
		&newProposal2,
		&newProposal3,
		&newProposal4,
	},
	}

	// valid request
	resp, err := client.CreateMultiProposalV1(context.Background(), req1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, []uint64{1, 2, 3, 4}, resp.ProposalIds)

	// invalid request
	newProposal5 := fb.NewProposal{UserId: 0, LessonId: 24, DocumentId: 10}
	req2 := &fb.CreateMultiProposalV1Request{NewProposals: []*fb.NewProposal{
		&newProposal5,
	},
	}

	resp2, err := client.CreateMultiProposalV1(context.Background(), req2)
	require.Error(t, err)
	require.Nil(t, resp2)
}

func TestClientRemoveProposal(t *testing.T) {
	t.Parallel()

	serverAddress := startTestGrpcServer(t,
		nil, repo.NewInMemoryProposalRepo())
	client := newTestGrpcClient(t, serverAddress)

	newProposal1 := fb.NewProposal{UserId: 42, LessonId: 24, DocumentId: 7}
	newProposal2 := fb.NewProposal{UserId: 420, LessonId: 240, DocumentId: 70}
	newProposal3 := fb.NewProposal{UserId: 4200, LessonId: 2400, DocumentId: 700}
	newProposal4 := fb.NewProposal{UserId: 42000, LessonId: 24000, DocumentId: 7000}

	req1 := &fb.CreateMultiProposalV1Request{NewProposals: []*fb.NewProposal{
		&newProposal1,
		&newProposal2,
		&newProposal3,
		&newProposal4,
	},
	}

	respCreate, err := client.CreateMultiProposalV1(context.Background(), req1)
	require.NoError(t, err)
	require.NotNil(t, respCreate)

	// remove the second newProposal2 record
	reqRemove1 := &fb.RemoveProposalV1Request{ProposalId: respCreate.ProposalIds[1]}
	respRemove1, err := client.RemoveProposalV1(context.Background(), reqRemove1)

	require.NoError(t, err)
	require.NotNil(t, respRemove1)

	// try to remove the second time
	reqRemove2 := &fb.RemoveProposalV1Request{ProposalId: respCreate.ProposalIds[1]}
	resRemove2, err := client.RemoveProposalV1(context.Background(), reqRemove2)
	require.Error(t, err)
	require.Nil(t, resRemove2)

	// invalid reqest
	reqRemove3 := &fb.RemoveProposalV1Request{ProposalId: 0}
	resRemove3, err := client.RemoveProposalV1(context.Background(), reqRemove3)
	require.Error(t, err)
	require.Nil(t, resRemove3)

}

func TestClientDescribeProposal(t *testing.T) {
	t.Parallel()

	serverAddress := startTestGrpcServer(t,
		nil, repo.NewInMemoryProposalRepo())
	client := newTestGrpcClient(t, serverAddress)

	newProposal1 := fb.NewProposal{UserId: 42, LessonId: 24, DocumentId: 7}
	req1 := &fb.CreateProposalV1Request{NewProposal: &newProposal1}
	resp1, err1 := client.CreateProposalV1(context.Background(), req1)
	require.NoError(t, err1)
	require.NotNil(t, resp1)

	// valid request
	reqDescribe1 := &fb.DescribeProposalV1Request{ProposalId: resp1.ProposalId}
	respDescribe1, err := client.DescribeProposalV1(context.Background(), reqDescribe1)
	require.NoError(t, err)
	require.NotNil(t, respDescribe1)
	require.Equal(t, respDescribe1.Proposal.ProposalId, uint64(1))
	require.Equal(t, respDescribe1.Proposal.UserId, newProposal1.UserId)
	require.Equal(t, respDescribe1.Proposal.LessonId, newProposal1.LessonId)
	require.Equal(t, respDescribe1.Proposal.DocumentId, newProposal1.DocumentId)

	// missing id
	reqDescribe2 := &fb.DescribeProposalV1Request{ProposalId: resp1.ProposalId + 1}
	respDescribe2, err := client.DescribeProposalV1(context.Background(), reqDescribe2)
	require.Error(t, err)
	require.Nil(t, respDescribe2)

	// invalid request
	reqDescribe3 := &fb.DescribeProposalV1Request{ProposalId: 0}
	respDescribe3, err := client.DescribeProposalV1(context.Background(), reqDescribe3)
	require.Error(t, err)
	require.Nil(t, respDescribe3)
}

func TestClientListProposals(t *testing.T) {
	t.Parallel()

	serverAddress := startTestGrpcServer(t,
		nil, repo.NewInMemoryProposalRepo())
	client := newTestGrpcClient(t, serverAddress)

	newProposal1 := fb.NewProposal{UserId: 42, LessonId: 24, DocumentId: 7}
	newProposal2 := fb.NewProposal{UserId: 420, LessonId: 240, DocumentId: 70}
	newProposal3 := fb.NewProposal{UserId: 4200, LessonId: 2400, DocumentId: 700}
	newProposal4 := fb.NewProposal{UserId: 42000, LessonId: 24000, DocumentId: 7000}

	req1 := &fb.CreateMultiProposalV1Request{NewProposals: []*fb.NewProposal{
		&newProposal1,
		&newProposal2,
		&newProposal3,
		&newProposal4,
	},
	}

	resp, err := client.CreateMultiProposalV1(context.Background(), req1)
	require.NoError(t, err)
	require.NotNil(t, resp)

	// valid request
	reqList1 := &fb.ListProposalsV1Request{Limit: 2, Offset: 1}
	respList1, err := client.ListProposalsV1(context.Background(), reqList1)
	require.NoError(t, err)
	require.NotNil(t, respList1)

	require.Equal(t, respList1.Proposals[0].ProposalId, uint64(2))
	require.Equal(t, respList1.Proposals[0].UserId, newProposal2.UserId)
	require.Equal(t, respList1.Proposals[0].LessonId, newProposal2.LessonId)
	require.Equal(t, respList1.Proposals[0].DocumentId, newProposal2.DocumentId)

	require.Equal(t, respList1.Proposals[1].ProposalId, uint64(3))
	require.Equal(t, respList1.Proposals[1].UserId, newProposal3.UserId)
	require.Equal(t, respList1.Proposals[1].LessonId, newProposal3.LessonId)
	require.Equal(t, respList1.Proposals[1].DocumentId, newProposal3.DocumentId)

	// wrong offset
	reqList2 := &fb.ListProposalsV1Request{Limit: 1, Offset: 4}
	respList2, err := client.ListProposalsV1(context.Background(), reqList2)
	require.NoError(t, err)
	require.NotNil(t, respList2)
	require.Equal(t, len(respList2.Proposals), 0)

	// invalid request
	reqList3 := &fb.ListProposalsV1Request{Limit: 0, Offset: 1}
	respList3, err := client.ListProposalsV1(context.Background(), reqList3)
	require.Error(t, err)
	require.Nil(t, respList3)
}
