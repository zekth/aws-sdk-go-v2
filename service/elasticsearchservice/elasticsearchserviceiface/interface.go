// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package elasticsearchserviceiface provides an interface to enable mocking the Amazon Elasticsearch Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package elasticsearchserviceiface

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
)

// ClientAPI provides an interface to enable mocking the
// elasticsearchservice.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Elasticsearch Service.
//    func myFunc(svc elasticsearchserviceiface.ClientAPI) bool {
//        // Make svc.AcceptInboundCrossClusterSearchConnection request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := elasticsearchservice.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        elasticsearchserviceiface.ClientPI
//    }
//    func (m *mockClientClient) AcceptInboundCrossClusterSearchConnection(input *elasticsearchservice.AcceptInboundCrossClusterSearchConnectionInput) (*elasticsearchservice.AcceptInboundCrossClusterSearchConnectionOutput, error) {
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
	AcceptInboundCrossClusterSearchConnectionRequest(*elasticsearchservice.AcceptInboundCrossClusterSearchConnectionInput) elasticsearchservice.AcceptInboundCrossClusterSearchConnectionRequest

	AddTagsRequest(*elasticsearchservice.AddTagsInput) elasticsearchservice.AddTagsRequest

	AssociatePackageRequest(*elasticsearchservice.AssociatePackageInput) elasticsearchservice.AssociatePackageRequest

	CancelElasticsearchServiceSoftwareUpdateRequest(*elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateInput) elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateRequest

	CreateElasticsearchDomainRequest(*elasticsearchservice.CreateElasticsearchDomainInput) elasticsearchservice.CreateElasticsearchDomainRequest

	CreateOutboundCrossClusterSearchConnectionRequest(*elasticsearchservice.CreateOutboundCrossClusterSearchConnectionInput) elasticsearchservice.CreateOutboundCrossClusterSearchConnectionRequest

	CreatePackageRequest(*elasticsearchservice.CreatePackageInput) elasticsearchservice.CreatePackageRequest

	DeleteElasticsearchDomainRequest(*elasticsearchservice.DeleteElasticsearchDomainInput) elasticsearchservice.DeleteElasticsearchDomainRequest

	DeleteElasticsearchServiceRoleRequest(*elasticsearchservice.DeleteElasticsearchServiceRoleInput) elasticsearchservice.DeleteElasticsearchServiceRoleRequest

	DeleteInboundCrossClusterSearchConnectionRequest(*elasticsearchservice.DeleteInboundCrossClusterSearchConnectionInput) elasticsearchservice.DeleteInboundCrossClusterSearchConnectionRequest

	DeleteOutboundCrossClusterSearchConnectionRequest(*elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionInput) elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionRequest

	DeletePackageRequest(*elasticsearchservice.DeletePackageInput) elasticsearchservice.DeletePackageRequest

	DescribeElasticsearchDomainRequest(*elasticsearchservice.DescribeElasticsearchDomainInput) elasticsearchservice.DescribeElasticsearchDomainRequest

	DescribeElasticsearchDomainConfigRequest(*elasticsearchservice.DescribeElasticsearchDomainConfigInput) elasticsearchservice.DescribeElasticsearchDomainConfigRequest

	DescribeElasticsearchDomainsRequest(*elasticsearchservice.DescribeElasticsearchDomainsInput) elasticsearchservice.DescribeElasticsearchDomainsRequest

	DescribeElasticsearchInstanceTypeLimitsRequest(*elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsInput) elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsRequest

	DescribeInboundCrossClusterSearchConnectionsRequest(*elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsInput) elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsRequest

	DescribeOutboundCrossClusterSearchConnectionsRequest(*elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsInput) elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsRequest

	DescribePackagesRequest(*elasticsearchservice.DescribePackagesInput) elasticsearchservice.DescribePackagesRequest

	DescribeReservedElasticsearchInstanceOfferingsRequest(*elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsInput) elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsRequest

	DescribeReservedElasticsearchInstancesRequest(*elasticsearchservice.DescribeReservedElasticsearchInstancesInput) elasticsearchservice.DescribeReservedElasticsearchInstancesRequest

	DissociatePackageRequest(*elasticsearchservice.DissociatePackageInput) elasticsearchservice.DissociatePackageRequest

	GetCompatibleElasticsearchVersionsRequest(*elasticsearchservice.GetCompatibleElasticsearchVersionsInput) elasticsearchservice.GetCompatibleElasticsearchVersionsRequest

	GetUpgradeHistoryRequest(*elasticsearchservice.GetUpgradeHistoryInput) elasticsearchservice.GetUpgradeHistoryRequest

	GetUpgradeStatusRequest(*elasticsearchservice.GetUpgradeStatusInput) elasticsearchservice.GetUpgradeStatusRequest

	ListDomainNamesRequest(*elasticsearchservice.ListDomainNamesInput) elasticsearchservice.ListDomainNamesRequest

	ListDomainsForPackageRequest(*elasticsearchservice.ListDomainsForPackageInput) elasticsearchservice.ListDomainsForPackageRequest

	ListElasticsearchInstanceTypesRequest(*elasticsearchservice.ListElasticsearchInstanceTypesInput) elasticsearchservice.ListElasticsearchInstanceTypesRequest

	ListElasticsearchVersionsRequest(*elasticsearchservice.ListElasticsearchVersionsInput) elasticsearchservice.ListElasticsearchVersionsRequest

	ListPackagesForDomainRequest(*elasticsearchservice.ListPackagesForDomainInput) elasticsearchservice.ListPackagesForDomainRequest

	ListTagsRequest(*elasticsearchservice.ListTagsInput) elasticsearchservice.ListTagsRequest

	PurchaseReservedElasticsearchInstanceOfferingRequest(*elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingInput) elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingRequest

	RejectInboundCrossClusterSearchConnectionRequest(*elasticsearchservice.RejectInboundCrossClusterSearchConnectionInput) elasticsearchservice.RejectInboundCrossClusterSearchConnectionRequest

	RemoveTagsRequest(*elasticsearchservice.RemoveTagsInput) elasticsearchservice.RemoveTagsRequest

	StartElasticsearchServiceSoftwareUpdateRequest(*elasticsearchservice.StartElasticsearchServiceSoftwareUpdateInput) elasticsearchservice.StartElasticsearchServiceSoftwareUpdateRequest

	UpdateElasticsearchDomainConfigRequest(*elasticsearchservice.UpdateElasticsearchDomainConfigInput) elasticsearchservice.UpdateElasticsearchDomainConfigRequest

	UpgradeElasticsearchDomainRequest(*elasticsearchservice.UpgradeElasticsearchDomainInput) elasticsearchservice.UpgradeElasticsearchDomainRequest
}

var _ ClientAPI = (*elasticsearchservice.Client)(nil)
