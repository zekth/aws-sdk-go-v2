// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package secretsmanageriface provides an interface to enable mocking the AWS Secrets Manager service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package secretsmanageriface

import (
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// ClientAPI provides an interface to enable mocking the
// secretsmanager.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Secrets Manager.
//    func myFunc(svc secretsmanageriface.ClientAPI) bool {
//        // Make svc.CancelRotateSecret request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := secretsmanager.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        secretsmanageriface.ClientPI
//    }
//    func (m *mockClientClient) CancelRotateSecret(input *secretsmanager.CancelRotateSecretInput) (*secretsmanager.CancelRotateSecretOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	CancelRotateSecretRequest(*secretsmanager.CancelRotateSecretInput) secretsmanager.CancelRotateSecretRequest

	CreateSecretRequest(*secretsmanager.CreateSecretInput) secretsmanager.CreateSecretRequest

	DeleteResourcePolicyRequest(*secretsmanager.DeleteResourcePolicyInput) secretsmanager.DeleteResourcePolicyRequest

	DeleteSecretRequest(*secretsmanager.DeleteSecretInput) secretsmanager.DeleteSecretRequest

	DescribeSecretRequest(*secretsmanager.DescribeSecretInput) secretsmanager.DescribeSecretRequest

	GetRandomPasswordRequest(*secretsmanager.GetRandomPasswordInput) secretsmanager.GetRandomPasswordRequest

	GetResourcePolicyRequest(*secretsmanager.GetResourcePolicyInput) secretsmanager.GetResourcePolicyRequest

	GetSecretValueRequest(*secretsmanager.GetSecretValueInput) secretsmanager.GetSecretValueRequest

	ListSecretVersionIdsRequest(*secretsmanager.ListSecretVersionIdsInput) secretsmanager.ListSecretVersionIdsRequest

	ListSecretsRequest(*secretsmanager.ListSecretsInput) secretsmanager.ListSecretsRequest

	PutResourcePolicyRequest(*secretsmanager.PutResourcePolicyInput) secretsmanager.PutResourcePolicyRequest

	PutSecretValueRequest(*secretsmanager.PutSecretValueInput) secretsmanager.PutSecretValueRequest

	RestoreSecretRequest(*secretsmanager.RestoreSecretInput) secretsmanager.RestoreSecretRequest

	RotateSecretRequest(*secretsmanager.RotateSecretInput) secretsmanager.RotateSecretRequest

	TagResourceRequest(*secretsmanager.TagResourceInput) secretsmanager.TagResourceRequest

	UntagResourceRequest(*secretsmanager.UntagResourceInput) secretsmanager.UntagResourceRequest

	UpdateSecretRequest(*secretsmanager.UpdateSecretInput) secretsmanager.UpdateSecretRequest

	UpdateSecretVersionStageRequest(*secretsmanager.UpdateSecretVersionStageInput) secretsmanager.UpdateSecretVersionStageRequest

	ValidateResourcePolicyRequest(*secretsmanager.ValidateResourcePolicyInput) secretsmanager.ValidateResourcePolicyRequest
}

var _ ClientAPI = (*secretsmanager.Client)(nil)
