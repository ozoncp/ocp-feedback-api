package feedback_grpc_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFeedbackGrpc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FeedbackGrpc Suite")
}
