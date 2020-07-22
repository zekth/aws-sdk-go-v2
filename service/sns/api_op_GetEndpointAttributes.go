// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input for GetEndpointAttributes action.
type GetEndpointAttributesInput struct {
	_ struct{} `type:"structure"`

	// EndpointArn for GetEndpointAttributes input.
	//
	// EndpointArn is a required field
	EndpointArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetEndpointAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetEndpointAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetEndpointAttributesInput"}

	if s.EndpointArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response from GetEndpointAttributes of the EndpointArn.
type GetEndpointAttributesOutput struct {
	_ struct{} `type:"structure"`

	// Attributes include the following:
	//
	//    * CustomUserData – arbitrary user data to associate with the endpoint.
	//    Amazon SNS does not use this data. The data must be in UTF-8 format and
	//    less than 2KB.
	//
	//    * Enabled – flag that enables/disables delivery to the endpoint. Amazon
	//    SNS will set this to false when a notification service indicates to Amazon
	//    SNS that the endpoint is invalid. Users can set it back to true, typically
	//    after updating Token.
	//
	//    * Token – device token, also referred to as a registration id, for an
	//    app and mobile device. This is returned from the notification service
	//    when an app and mobile device are registered with the notification service.
	//    The device token for the iOS platform is returned in lowercase.
	Attributes map[string]string `type:"map"`
}

// String returns the string representation
func (s GetEndpointAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetEndpointAttributes = "GetEndpointAttributes"

// GetEndpointAttributesRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Retrieves the endpoint attributes for a device on one of the supported push
// notification services, such as GCM (Firebase Cloud Messaging) and APNS. For
// more information, see Using Amazon SNS Mobile Push Notifications (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html).
//
//    // Example sending a request using GetEndpointAttributesRequest.
//    req := client.GetEndpointAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/GetEndpointAttributes
func (c *Client) GetEndpointAttributesRequest(input *GetEndpointAttributesInput) GetEndpointAttributesRequest {
	op := &aws.Operation{
		Name:       opGetEndpointAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetEndpointAttributesInput{}
	}

	req := c.newRequest(op, input, &GetEndpointAttributesOutput{})

	return GetEndpointAttributesRequest{Request: req, Input: input, Copy: c.GetEndpointAttributesRequest}
}

// GetEndpointAttributesRequest is the request type for the
// GetEndpointAttributes API operation.
type GetEndpointAttributesRequest struct {
	*aws.Request
	Input *GetEndpointAttributesInput
	Copy  func(*GetEndpointAttributesInput) GetEndpointAttributesRequest
}

// Send marshals and sends the GetEndpointAttributes API request.
func (r GetEndpointAttributesRequest) Send(ctx context.Context) (*GetEndpointAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetEndpointAttributesResponse{
		GetEndpointAttributesOutput: r.Request.Data.(*GetEndpointAttributesOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetEndpointAttributesResponse is the response type for the
// GetEndpointAttributes API operation.
type GetEndpointAttributesResponse struct {
	*GetEndpointAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetEndpointAttributes request.
func (r *GetEndpointAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
