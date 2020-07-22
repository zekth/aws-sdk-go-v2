// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DetectSyntaxInput struct {
	_ struct{} `type:"structure"`

	// The language code of the input documents. You can specify any of the following
	// languages supported by Amazon Comprehend: German ("de"), English ("en"),
	// Spanish ("es"), French ("fr"), Italian ("it"), or Portuguese ("pt").
	//
	// LanguageCode is a required field
	LanguageCode SyntaxLanguageCode `type:"string" required:"true" enum:"true"`

	// A UTF-8 string. Each string must contain fewer that 5,000 bytes of UTF encoded
	// characters.
	//
	// Text is a required field
	Text *string `min:"1" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s DetectSyntaxInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetectSyntaxInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetectSyntaxInput"}
	if len(s.LanguageCode) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("LanguageCode"))
	}

	if s.Text == nil {
		invalidParams.Add(aws.NewErrParamRequired("Text"))
	}
	if s.Text != nil && len(*s.Text) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Text", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DetectSyntaxOutput struct {
	_ struct{} `type:"structure" sensitive:"true"`

	// A collection of syntax tokens describing the text. For each token, the response
	// provides the text, the token type, where the text begins and ends, and the
	// level of confidence that Amazon Comprehend has that the token is correct.
	// For a list of token types, see how-syntax.
	SyntaxTokens []SyntaxToken `type:"list"`
}

// String returns the string representation
func (s DetectSyntaxOutput) String() string {
	return awsutil.Prettify(s)
}

const opDetectSyntax = "DetectSyntax"

// DetectSyntaxRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Inspects text for syntax and the part of speech of words in the document.
// For more information, how-syntax.
//
//    // Example sending a request using DetectSyntaxRequest.
//    req := client.DetectSyntaxRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/DetectSyntax
func (c *Client) DetectSyntaxRequest(input *DetectSyntaxInput) DetectSyntaxRequest {
	op := &aws.Operation{
		Name:       opDetectSyntax,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetectSyntaxInput{}
	}

	req := c.newRequest(op, input, &DetectSyntaxOutput{})

	return DetectSyntaxRequest{Request: req, Input: input, Copy: c.DetectSyntaxRequest}
}

// DetectSyntaxRequest is the request type for the
// DetectSyntax API operation.
type DetectSyntaxRequest struct {
	*aws.Request
	Input *DetectSyntaxInput
	Copy  func(*DetectSyntaxInput) DetectSyntaxRequest
}

// Send marshals and sends the DetectSyntax API request.
func (r DetectSyntaxRequest) Send(ctx context.Context) (*DetectSyntaxResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetectSyntaxResponse{
		DetectSyntaxOutput: r.Request.Data.(*DetectSyntaxOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetectSyntaxResponse is the response type for the
// DetectSyntax API operation.
type DetectSyntaxResponse struct {
	*DetectSyntaxOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetectSyntax request.
func (r *DetectSyntaxResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
