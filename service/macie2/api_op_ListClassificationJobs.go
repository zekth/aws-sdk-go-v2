// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Specifies criteria for filtering, sorting, and paginating the results of
// a request for information about classification jobs.
type ListClassificationJobsInput struct {
	_ struct{} `type:"structure"`

	// Specifies criteria for filtering the results of a request for information
	// about classification jobs.
	FilterCriteria *ListJobsFilterCriteria `locationName:"filterCriteria" type:"structure"`

	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	NextToken *string `locationName:"nextToken" type:"string"`

	// Specifies criteria for sorting the results of a request for information about
	// classification jobs.
	SortCriteria *ListJobsSortCriteria `locationName:"sortCriteria" type:"structure"`
}

// String returns the string representation
func (s ListClassificationJobsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListClassificationJobsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FilterCriteria != nil {
		v := s.FilterCriteria

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "filterCriteria", v, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SortCriteria != nil {
		v := s.SortCriteria

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "sortCriteria", v, metadata)
	}
	return nil
}

// Provides the results of a request for information about one or more classification
// jobs.
type ListClassificationJobsOutput struct {
	_ struct{} `type:"structure"`

	Items []JobSummary `locationName:"items" type:"list"`

	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListClassificationJobsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListClassificationJobsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Items != nil {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "items", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListClassificationJobs = "ListClassificationJobs"

// ListClassificationJobsRequest returns a request value for making API operation for
// Amazon Macie 2.
//
// Retrieves a subset of information about one or more classification jobs.
//
//    // Example sending a request using ListClassificationJobsRequest.
//    req := client.ListClassificationJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie2-2020-01-01/ListClassificationJobs
func (c *Client) ListClassificationJobsRequest(input *ListClassificationJobsInput) ListClassificationJobsRequest {
	op := &aws.Operation{
		Name:       opListClassificationJobs,
		HTTPMethod: "POST",
		HTTPPath:   "/jobs/list",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListClassificationJobsInput{}
	}

	req := c.newRequest(op, input, &ListClassificationJobsOutput{})

	return ListClassificationJobsRequest{Request: req, Input: input, Copy: c.ListClassificationJobsRequest}
}

// ListClassificationJobsRequest is the request type for the
// ListClassificationJobs API operation.
type ListClassificationJobsRequest struct {
	*aws.Request
	Input *ListClassificationJobsInput
	Copy  func(*ListClassificationJobsInput) ListClassificationJobsRequest
}

// Send marshals and sends the ListClassificationJobs API request.
func (r ListClassificationJobsRequest) Send(ctx context.Context) (*ListClassificationJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListClassificationJobsResponse{
		ListClassificationJobsOutput: r.Request.Data.(*ListClassificationJobsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListClassificationJobsRequestPaginator returns a paginator for ListClassificationJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListClassificationJobsRequest(input)
//   p := macie2.NewListClassificationJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListClassificationJobsPaginator(req ListClassificationJobsRequest) ListClassificationJobsPaginator {
	return ListClassificationJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListClassificationJobsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListClassificationJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListClassificationJobsPaginator struct {
	aws.Pager
}

func (p *ListClassificationJobsPaginator) CurrentPage() *ListClassificationJobsOutput {
	return p.Pager.CurrentPage().(*ListClassificationJobsOutput)
}

// ListClassificationJobsResponse is the response type for the
// ListClassificationJobs API operation.
type ListClassificationJobsResponse struct {
	*ListClassificationJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListClassificationJobs request.
func (r *ListClassificationJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
