package proposal_grpc_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestProposalGrpc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProposalGrpc Suite")
}
