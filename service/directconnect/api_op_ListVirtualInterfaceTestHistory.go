// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListVirtualInterfaceTestHistoryInput struct {
	_ struct{} `type:"structure"`

	// The BGP peers that were placed in the DOWN state during the virtual interface
	// failover test.
	BgpPeers []string `locationName:"bgpPeers" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	//
	// If MaxResults is given a value larger than 100, only 100 results are returned.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The token for the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The status of the virtual interface failover test.
	Status *string `locationName:"status" type:"string"`

	// The ID of the virtual interface failover test.
	TestId *string `locationName:"testId" type:"string"`

	// The ID of the virtual interface that was tested.
	VirtualInterfaceId *string `locationName:"virtualInterfaceId" type:"string"`
}

// String returns the string representation
func (s ListVirtualInterfaceTestHistoryInput) String() string {
	return awsutil.Prettify(s)
}

type ListVirtualInterfaceTestHistoryOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The ID of the tested virtual interface.
	VirtualInterfaceTestHistory []VirtualInterfaceTestHistory `locationName:"virtualInterfaceTestHistory" type:"list"`
}

// String returns the string representation
func (s ListVirtualInterfaceTestHistoryOutput) String() string {
	return awsutil.Prettify(s)
}

const opListVirtualInterfaceTestHistory = "ListVirtualInterfaceTestHistory"

// ListVirtualInterfaceTestHistoryRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Lists the virtual interface failover test history.
//
//    // Example sending a request using ListVirtualInterfaceTestHistoryRequest.
//    req := client.ListVirtualInterfaceTestHistoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/ListVirtualInterfaceTestHistory
func (c *Client) ListVirtualInterfaceTestHistoryRequest(input *ListVirtualInterfaceTestHistoryInput) ListVirtualInterfaceTestHistoryRequest {
	op := &aws.Operation{
		Name:       opListVirtualInterfaceTestHistory,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListVirtualInterfaceTestHistoryInput{}
	}

	req := c.newRequest(op, input, &ListVirtualInterfaceTestHistoryOutput{})

	return ListVirtualInterfaceTestHistoryRequest{Request: req, Input: input, Copy: c.ListVirtualInterfaceTestHistoryRequest}
}

// ListVirtualInterfaceTestHistoryRequest is the request type for the
// ListVirtualInterfaceTestHistory API operation.
type ListVirtualInterfaceTestHistoryRequest struct {
	*aws.Request
	Input *ListVirtualInterfaceTestHistoryInput
	Copy  func(*ListVirtualInterfaceTestHistoryInput) ListVirtualInterfaceTestHistoryRequest
}

// Send marshals and sends the ListVirtualInterfaceTestHistory API request.
func (r ListVirtualInterfaceTestHistoryRequest) Send(ctx context.Context) (*ListVirtualInterfaceTestHistoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListVirtualInterfaceTestHistoryResponse{
		ListVirtualInterfaceTestHistoryOutput: r.Request.Data.(*ListVirtualInterfaceTestHistoryOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListVirtualInterfaceTestHistoryResponse is the response type for the
// ListVirtualInterfaceTestHistory API operation.
type ListVirtualInterfaceTestHistoryResponse struct {
	*ListVirtualInterfaceTestHistoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListVirtualInterfaceTestHistory request.
func (r *ListVirtualInterfaceTestHistoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
