// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package macie2iface provides an interface to enable mocking the Amazon Macie 2 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package macie2iface

import (
	"github.com/aws/aws-sdk-go-v2/service/macie2"
)

// ClientAPI provides an interface to enable mocking the
// macie2.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Macie 2.
//    func myFunc(svc macie2iface.ClientAPI) bool {
//        // Make svc.AcceptInvitation request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := macie2.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        macie2iface.ClientPI
//    }
//    func (m *mockClientClient) AcceptInvitation(input *macie2.AcceptInvitationInput) (*macie2.AcceptInvitationOutput, error) {
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
	AcceptInvitationRequest(*macie2.AcceptInvitationInput) macie2.AcceptInvitationRequest

	BatchGetCustomDataIdentifiersRequest(*macie2.BatchGetCustomDataIdentifiersInput) macie2.BatchGetCustomDataIdentifiersRequest

	CreateClassificationJobRequest(*macie2.CreateClassificationJobInput) macie2.CreateClassificationJobRequest

	CreateCustomDataIdentifierRequest(*macie2.CreateCustomDataIdentifierInput) macie2.CreateCustomDataIdentifierRequest

	CreateFindingsFilterRequest(*macie2.CreateFindingsFilterInput) macie2.CreateFindingsFilterRequest

	CreateInvitationsRequest(*macie2.CreateInvitationsInput) macie2.CreateInvitationsRequest

	CreateMemberRequest(*macie2.CreateMemberInput) macie2.CreateMemberRequest

	CreateSampleFindingsRequest(*macie2.CreateSampleFindingsInput) macie2.CreateSampleFindingsRequest

	DeclineInvitationsRequest(*macie2.DeclineInvitationsInput) macie2.DeclineInvitationsRequest

	DeleteCustomDataIdentifierRequest(*macie2.DeleteCustomDataIdentifierInput) macie2.DeleteCustomDataIdentifierRequest

	DeleteFindingsFilterRequest(*macie2.DeleteFindingsFilterInput) macie2.DeleteFindingsFilterRequest

	DeleteInvitationsRequest(*macie2.DeleteInvitationsInput) macie2.DeleteInvitationsRequest

	DeleteMemberRequest(*macie2.DeleteMemberInput) macie2.DeleteMemberRequest

	DescribeBucketsRequest(*macie2.DescribeBucketsInput) macie2.DescribeBucketsRequest

	DescribeClassificationJobRequest(*macie2.DescribeClassificationJobInput) macie2.DescribeClassificationJobRequest

	DescribeOrganizationConfigurationRequest(*macie2.DescribeOrganizationConfigurationInput) macie2.DescribeOrganizationConfigurationRequest

	DisableMacieRequest(*macie2.DisableMacieInput) macie2.DisableMacieRequest

	DisableOrganizationAdminAccountRequest(*macie2.DisableOrganizationAdminAccountInput) macie2.DisableOrganizationAdminAccountRequest

	DisassociateFromMasterAccountRequest(*macie2.DisassociateFromMasterAccountInput) macie2.DisassociateFromMasterAccountRequest

	DisassociateMemberRequest(*macie2.DisassociateMemberInput) macie2.DisassociateMemberRequest

	EnableMacieRequest(*macie2.EnableMacieInput) macie2.EnableMacieRequest

	EnableOrganizationAdminAccountRequest(*macie2.EnableOrganizationAdminAccountInput) macie2.EnableOrganizationAdminAccountRequest

	GetBucketStatisticsRequest(*macie2.GetBucketStatisticsInput) macie2.GetBucketStatisticsRequest

	GetClassificationExportConfigurationRequest(*macie2.GetClassificationExportConfigurationInput) macie2.GetClassificationExportConfigurationRequest

	GetCustomDataIdentifierRequest(*macie2.GetCustomDataIdentifierInput) macie2.GetCustomDataIdentifierRequest

	GetFindingStatisticsRequest(*macie2.GetFindingStatisticsInput) macie2.GetFindingStatisticsRequest

	GetFindingsRequest(*macie2.GetFindingsInput) macie2.GetFindingsRequest

	GetFindingsFilterRequest(*macie2.GetFindingsFilterInput) macie2.GetFindingsFilterRequest

	GetInvitationsCountRequest(*macie2.GetInvitationsCountInput) macie2.GetInvitationsCountRequest

	GetMacieSessionRequest(*macie2.GetMacieSessionInput) macie2.GetMacieSessionRequest

	GetMasterAccountRequest(*macie2.GetMasterAccountInput) macie2.GetMasterAccountRequest

	GetMemberRequest(*macie2.GetMemberInput) macie2.GetMemberRequest

	GetUsageStatisticsRequest(*macie2.GetUsageStatisticsInput) macie2.GetUsageStatisticsRequest

	GetUsageTotalsRequest(*macie2.GetUsageTotalsInput) macie2.GetUsageTotalsRequest

	ListClassificationJobsRequest(*macie2.ListClassificationJobsInput) macie2.ListClassificationJobsRequest

	ListCustomDataIdentifiersRequest(*macie2.ListCustomDataIdentifiersInput) macie2.ListCustomDataIdentifiersRequest

	ListFindingsRequest(*macie2.ListFindingsInput) macie2.ListFindingsRequest

	ListFindingsFiltersRequest(*macie2.ListFindingsFiltersInput) macie2.ListFindingsFiltersRequest

	ListInvitationsRequest(*macie2.ListInvitationsInput) macie2.ListInvitationsRequest

	ListMembersRequest(*macie2.ListMembersInput) macie2.ListMembersRequest

	ListOrganizationAdminAccountsRequest(*macie2.ListOrganizationAdminAccountsInput) macie2.ListOrganizationAdminAccountsRequest

	ListTagsForResourceRequest(*macie2.ListTagsForResourceInput) macie2.ListTagsForResourceRequest

	PutClassificationExportConfigurationRequest(*macie2.PutClassificationExportConfigurationInput) macie2.PutClassificationExportConfigurationRequest

	TagResourceRequest(*macie2.TagResourceInput) macie2.TagResourceRequest

	TestCustomDataIdentifierRequest(*macie2.TestCustomDataIdentifierInput) macie2.TestCustomDataIdentifierRequest

	UntagResourceRequest(*macie2.UntagResourceInput) macie2.UntagResourceRequest

	UpdateClassificationJobRequest(*macie2.UpdateClassificationJobInput) macie2.UpdateClassificationJobRequest

	UpdateFindingsFilterRequest(*macie2.UpdateFindingsFilterInput) macie2.UpdateFindingsFilterRequest

	UpdateMacieSessionRequest(*macie2.UpdateMacieSessionInput) macie2.UpdateMacieSessionRequest

	UpdateMemberSessionRequest(*macie2.UpdateMemberSessionInput) macie2.UpdateMemberSessionRequest

	UpdateOrganizationConfigurationRequest(*macie2.UpdateOrganizationConfigurationInput) macie2.UpdateOrganizationConfigurationRequest
}

var _ ClientAPI = (*macie2.Client)(nil)
