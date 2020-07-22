// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package frauddetector

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetEventPredictionInput struct {
	_ struct{} `type:"structure"`

	// The detector ID.
	//
	// DetectorId is a required field
	DetectorId *string `locationName:"detectorId" type:"string" required:"true"`

	// The detector version ID.
	DetectorVersionId *string `locationName:"detectorVersionId" type:"string"`

	// The entity type (associated with the detector's event type) and specific
	// entity ID representing who performed the event. If an entity id is not available,
	// use "UNKNOWN."
	//
	// Entities is a required field
	Entities []Entity `locationName:"entities" type:"list" required:"true"`

	// The unique ID used to identify the event.
	//
	// EventId is a required field
	EventId *string `locationName:"eventId" type:"string" required:"true"`

	// Timestamp that defines when the event under evaluation occurred.
	//
	// EventTimestamp is a required field
	EventTimestamp *string `locationName:"eventTimestamp" type:"string" required:"true"`

	// The event type associated with the detector specified for the prediction.
	//
	// EventTypeName is a required field
	EventTypeName *string `locationName:"eventTypeName" type:"string" required:"true"`

	// Names of the event type's variables you defined in Amazon Fraud Detector
	// to represent data elements and their corresponding values for the event you
	// are sending for evaluation.
	//
	// EventVariables is a required field
	EventVariables map[string]string `locationName:"eventVariables" min:"1" type:"map" required:"true"`

	// The Amazon SageMaker model endpoint input data blobs.
	ExternalModelEndpointDataBlobs map[string]ModelEndpointDataBlob `locationName:"externalModelEndpointDataBlobs" type:"map" sensitive:"true"`
}

// String returns the string representation
func (s GetEventPredictionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetEventPredictionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetEventPredictionInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}

	if s.Entities == nil {
		invalidParams.Add(aws.NewErrParamRequired("Entities"))
	}

	if s.EventId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventId"))
	}

	if s.EventTimestamp == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventTimestamp"))
	}

	if s.EventTypeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventTypeName"))
	}

	if s.EventVariables == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventVariables"))
	}
	if s.EventVariables != nil && len(s.EventVariables) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventVariables", 1))
	}
	if s.Entities != nil {
		for i, v := range s.Entities {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Entities", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.ExternalModelEndpointDataBlobs != nil {
		for i, v := range s.ExternalModelEndpointDataBlobs {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ExternalModelEndpointDataBlobs", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetEventPredictionOutput struct {
	_ struct{} `type:"structure"`

	// The model scores. Amazon Fraud Detector generates model scores between 0
	// and 1000, where 0 is low fraud risk and 1000 is high fraud risk. Model scores
	// are directly related to the false positive rate (FPR). For example, a score
	// of 600 corresponds to an estimated 10% false positive rate whereas a score
	// of 900 corresponds to an estimated 2% false positive rate.
	ModelScores []ModelScores `locationName:"modelScores" type:"list"`

	// The results.
	RuleResults []RuleResult `locationName:"ruleResults" type:"list"`
}

// String returns the string representation
func (s GetEventPredictionOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetEventPrediction = "GetEventPrediction"

// GetEventPredictionRequest returns a request value for making API operation for
// Amazon Fraud Detector.
//
// Evaluates an event against a detector version. If a version ID is not provided,
// the detector’s (ACTIVE) version is used.
//
//    // Example sending a request using GetEventPredictionRequest.
//    req := client.GetEventPredictionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/frauddetector-2019-11-15/GetEventPrediction
func (c *Client) GetEventPredictionRequest(input *GetEventPredictionInput) GetEventPredictionRequest {
	op := &aws.Operation{
		Name:       opGetEventPrediction,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetEventPredictionInput{}
	}

	req := c.newRequest(op, input, &GetEventPredictionOutput{})

	return GetEventPredictionRequest{Request: req, Input: input, Copy: c.GetEventPredictionRequest}
}

// GetEventPredictionRequest is the request type for the
// GetEventPrediction API operation.
type GetEventPredictionRequest struct {
	*aws.Request
	Input *GetEventPredictionInput
	Copy  func(*GetEventPredictionInput) GetEventPredictionRequest
}

// Send marshals and sends the GetEventPrediction API request.
func (r GetEventPredictionRequest) Send(ctx context.Context) (*GetEventPredictionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetEventPredictionResponse{
		GetEventPredictionOutput: r.Request.Data.(*GetEventPredictionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetEventPredictionResponse is the response type for the
// GetEventPrediction API operation.
type GetEventPredictionResponse struct {
	*GetEventPredictionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetEventPrediction request.
func (r *GetEventPredictionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
