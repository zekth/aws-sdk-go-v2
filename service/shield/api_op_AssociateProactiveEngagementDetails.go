// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateProactiveEngagementDetailsInput struct {
	_ struct{} `type:"structure"`

	// A list of email addresses and phone numbers that the DDoS Response Team (DRT)
	// can use to contact you for escalations to the DRT and to initiate proactive
	// customer support.
	//
	// To enable proactive engagement, the contact list must include at least one
	// phone number.
	//
	// The contacts that you provide here replace any contacts that were already
	// defined. If you already have contacts defined and want to use them, retrieve
	// the list using DescribeEmergencyContactSettings and then provide it here.
	//
	// EmergencyContactList is a required field
	EmergencyContactList []EmergencyContact `type:"list" required:"true"`
}

// String returns the string representation
func (s AssociateProactiveEngagementDetailsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateProactiveEngagementDetailsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateProactiveEngagementDetailsInput"}

	if s.EmergencyContactList == nil {
		invalidParams.Add(aws.NewErrParamRequired("EmergencyContactList"))
	}
	if s.EmergencyContactList != nil {
		for i, v := range s.EmergencyContactList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "EmergencyContactList", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateProactiveEngagementDetailsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateProactiveEngagementDetailsOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateProactiveEngagementDetails = "AssociateProactiveEngagementDetails"

// AssociateProactiveEngagementDetailsRequest returns a request value for making API operation for
// AWS Shield.
//
// Initializes proactive engagement and sets the list of contacts for the DDoS
// Response Team (DRT) to use. You must provide at least one phone number in
// the emergency contact list.
//
// After you have initialized proactive engagement using this call, to disable
// or enable proactive engagement, use the calls DisableProactiveEngagement
// and EnableProactiveEngagement.
//
// This call defines the list of email addresses and phone numbers that the
// DDoS Response Team (DRT) can use to contact you for escalations to the DRT
// and to initiate proactive customer support.
//
// The contacts that you provide in the request replace any contacts that were
// already defined. If you already have contacts defined and want to use them,
// retrieve the list using DescribeEmergencyContactSettings and then provide
// it to this call.
//
//    // Example sending a request using AssociateProactiveEngagementDetailsRequest.
//    req := client.AssociateProactiveEngagementDetailsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/AssociateProactiveEngagementDetails
func (c *Client) AssociateProactiveEngagementDetailsRequest(input *AssociateProactiveEngagementDetailsInput) AssociateProactiveEngagementDetailsRequest {
	op := &aws.Operation{
		Name:       opAssociateProactiveEngagementDetails,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateProactiveEngagementDetailsInput{}
	}

	req := c.newRequest(op, input, &AssociateProactiveEngagementDetailsOutput{})

	return AssociateProactiveEngagementDetailsRequest{Request: req, Input: input, Copy: c.AssociateProactiveEngagementDetailsRequest}
}

// AssociateProactiveEngagementDetailsRequest is the request type for the
// AssociateProactiveEngagementDetails API operation.
type AssociateProactiveEngagementDetailsRequest struct {
	*aws.Request
	Input *AssociateProactiveEngagementDetailsInput
	Copy  func(*AssociateProactiveEngagementDetailsInput) AssociateProactiveEngagementDetailsRequest
}

// Send marshals and sends the AssociateProactiveEngagementDetails API request.
func (r AssociateProactiveEngagementDetailsRequest) Send(ctx context.Context) (*AssociateProactiveEngagementDetailsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateProactiveEngagementDetailsResponse{
		AssociateProactiveEngagementDetailsOutput: r.Request.Data.(*AssociateProactiveEngagementDetailsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateProactiveEngagementDetailsResponse is the response type for the
// AssociateProactiveEngagementDetails API operation.
type AssociateProactiveEngagementDetailsResponse struct {
	*AssociateProactiveEngagementDetailsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateProactiveEngagementDetails request.
func (r *AssociateProactiveEngagementDetailsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
