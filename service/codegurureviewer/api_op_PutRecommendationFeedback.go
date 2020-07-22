// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codegurureviewer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type PutRecommendationFeedbackInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the CodeReview (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_CodeReview.html)
	// object.
	//
	// CodeReviewArn is a required field
	CodeReviewArn *string `min:"1" type:"string" required:"true"`

	// List for storing reactions. Reactions are utf-8 text code for emojis. If
	// you send an empty list it clears all your feedback.
	//
	// Reactions is a required field
	Reactions []Reaction `type:"list" required:"true"`

	// The recommendation ID that can be used to track the provided recommendations
	// and then to collect the feedback.
	//
	// RecommendationId is a required field
	RecommendationId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutRecommendationFeedbackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutRecommendationFeedbackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutRecommendationFeedbackInput"}

	if s.CodeReviewArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CodeReviewArn"))
	}
	if s.CodeReviewArn != nil && len(*s.CodeReviewArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CodeReviewArn", 1))
	}

	if s.Reactions == nil {
		invalidParams.Add(aws.NewErrParamRequired("Reactions"))
	}

	if s.RecommendationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RecommendationId"))
	}
	if s.RecommendationId != nil && len(*s.RecommendationId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RecommendationId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutRecommendationFeedbackInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CodeReviewArn != nil {
		v := *s.CodeReviewArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CodeReviewArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Reactions != nil {
		v := s.Reactions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Reactions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.RecommendationId != nil {
		v := *s.RecommendationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RecommendationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type PutRecommendationFeedbackOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutRecommendationFeedbackOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutRecommendationFeedbackOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutRecommendationFeedback = "PutRecommendationFeedback"

// PutRecommendationFeedbackRequest returns a request value for making API operation for
// Amazon CodeGuru Reviewer.
//
// Stores customer feedback for a CodeGuru Reviewer recommendation. When this
// API is called again with different reactions the previous feedback is overwritten.
//
//    // Example sending a request using PutRecommendationFeedbackRequest.
//    req := client.PutRecommendationFeedbackRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codeguru-reviewer-2019-09-19/PutRecommendationFeedback
func (c *Client) PutRecommendationFeedbackRequest(input *PutRecommendationFeedbackInput) PutRecommendationFeedbackRequest {
	op := &aws.Operation{
		Name:       opPutRecommendationFeedback,
		HTTPMethod: "PUT",
		HTTPPath:   "/feedback",
	}

	if input == nil {
		input = &PutRecommendationFeedbackInput{}
	}

	req := c.newRequest(op, input, &PutRecommendationFeedbackOutput{})

	return PutRecommendationFeedbackRequest{Request: req, Input: input, Copy: c.PutRecommendationFeedbackRequest}
}

// PutRecommendationFeedbackRequest is the request type for the
// PutRecommendationFeedback API operation.
type PutRecommendationFeedbackRequest struct {
	*aws.Request
	Input *PutRecommendationFeedbackInput
	Copy  func(*PutRecommendationFeedbackInput) PutRecommendationFeedbackRequest
}

// Send marshals and sends the PutRecommendationFeedback API request.
func (r PutRecommendationFeedbackRequest) Send(ctx context.Context) (*PutRecommendationFeedbackResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutRecommendationFeedbackResponse{
		PutRecommendationFeedbackOutput: r.Request.Data.(*PutRecommendationFeedbackOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutRecommendationFeedbackResponse is the response type for the
// PutRecommendationFeedback API operation.
type PutRecommendationFeedbackResponse struct {
	*PutRecommendationFeedbackOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutRecommendationFeedback request.
func (r *PutRecommendationFeedbackResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
