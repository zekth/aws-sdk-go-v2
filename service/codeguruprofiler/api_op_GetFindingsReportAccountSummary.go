// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codeguruprofiler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The structure representing the GetFindingsReportAccountSummaryRequest.
type GetFindingsReportAccountSummaryInput struct {
	_ struct{} `type:"structure"`

	// A Boolean value indicating whether to only return reports from daily profiles.
	// If set to True, only analysis data from daily profiles is returned. If set
	// to False, analysis data is returned from smaller time windows (for example,
	// one hour).
	DailyReportsOnly *bool `location:"querystring" locationName:"dailyReportsOnly" type:"boolean"`

	// The maximum number of results returned by GetFindingsReportAccountSummary
	// in paginated output. When this parameter is used, GetFindingsReportAccountSummary
	// only returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another GetFindingsReportAccountSummary request with the returned nextToken
	// value.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The nextToken value returned from a previous paginated GetFindingsReportAccountSummary
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `location:"querystring" locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s GetFindingsReportAccountSummaryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFindingsReportAccountSummaryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetFindingsReportAccountSummaryInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFindingsReportAccountSummaryInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DailyReportsOnly != nil {
		v := *s.DailyReportsOnly

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "dailyReportsOnly", protocol.BoolValue(v), metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The structure representing the GetFindingsReportAccountSummaryResponse.
type GetFindingsReportAccountSummaryOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken value to include in a future GetFindingsReportAccountSummary
	// request. When the results of a GetFindingsReportAccountSummary request exceed
	// maxResults, this value can be used to retrieve the next page of results.
	// This value is null when there are no more results to return.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// The return list of FindingsReportSummary (https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_FindingsReportSummary.html)
	// objects taht contain summaries of analysis results for all profiling groups
	// in your AWS account.
	//
	// ReportSummaries is a required field
	ReportSummaries []FindingsReportSummary `locationName:"reportSummaries" type:"list" required:"true"`
}

// String returns the string representation
func (s GetFindingsReportAccountSummaryOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFindingsReportAccountSummaryOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ReportSummaries != nil {
		v := s.ReportSummaries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "reportSummaries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetFindingsReportAccountSummary = "GetFindingsReportAccountSummary"

// GetFindingsReportAccountSummaryRequest returns a request value for making API operation for
// Amazon CodeGuru Profiler.
//
// Returns a list of FindingsReportSummary (https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_FindingsReportSummary.html)
// objects that contain analysis results for all profiling groups in your AWS
// account.
//
//    // Example sending a request using GetFindingsReportAccountSummaryRequest.
//    req := client.GetFindingsReportAccountSummaryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codeguruprofiler-2019-07-18/GetFindingsReportAccountSummary
func (c *Client) GetFindingsReportAccountSummaryRequest(input *GetFindingsReportAccountSummaryInput) GetFindingsReportAccountSummaryRequest {
	op := &aws.Operation{
		Name:       opGetFindingsReportAccountSummary,
		HTTPMethod: "GET",
		HTTPPath:   "/internal/findingsReports",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetFindingsReportAccountSummaryInput{}
	}

	req := c.newRequest(op, input, &GetFindingsReportAccountSummaryOutput{})

	return GetFindingsReportAccountSummaryRequest{Request: req, Input: input, Copy: c.GetFindingsReportAccountSummaryRequest}
}

// GetFindingsReportAccountSummaryRequest is the request type for the
// GetFindingsReportAccountSummary API operation.
type GetFindingsReportAccountSummaryRequest struct {
	*aws.Request
	Input *GetFindingsReportAccountSummaryInput
	Copy  func(*GetFindingsReportAccountSummaryInput) GetFindingsReportAccountSummaryRequest
}

// Send marshals and sends the GetFindingsReportAccountSummary API request.
func (r GetFindingsReportAccountSummaryRequest) Send(ctx context.Context) (*GetFindingsReportAccountSummaryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFindingsReportAccountSummaryResponse{
		GetFindingsReportAccountSummaryOutput: r.Request.Data.(*GetFindingsReportAccountSummaryOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetFindingsReportAccountSummaryRequestPaginator returns a paginator for GetFindingsReportAccountSummary.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetFindingsReportAccountSummaryRequest(input)
//   p := codeguruprofiler.NewGetFindingsReportAccountSummaryRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetFindingsReportAccountSummaryPaginator(req GetFindingsReportAccountSummaryRequest) GetFindingsReportAccountSummaryPaginator {
	return GetFindingsReportAccountSummaryPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetFindingsReportAccountSummaryInput
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

// GetFindingsReportAccountSummaryPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetFindingsReportAccountSummaryPaginator struct {
	aws.Pager
}

func (p *GetFindingsReportAccountSummaryPaginator) CurrentPage() *GetFindingsReportAccountSummaryOutput {
	return p.Pager.CurrentPage().(*GetFindingsReportAccountSummaryOutput)
}

// GetFindingsReportAccountSummaryResponse is the response type for the
// GetFindingsReportAccountSummary API operation.
type GetFindingsReportAccountSummaryResponse struct {
	*GetFindingsReportAccountSummaryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFindingsReportAccountSummary request.
func (r *GetFindingsReportAccountSummaryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
