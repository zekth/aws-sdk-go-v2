// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type PutVoiceConnectorEmergencyCallingConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The emergency calling configuration details.
	//
	// EmergencyCallingConfiguration is a required field
	EmergencyCallingConfiguration *EmergencyCallingConfiguration `type:"structure" required:"true"`

	// The Amazon Chime Voice Connector ID.
	//
	// VoiceConnectorId is a required field
	VoiceConnectorId *string `location:"uri" locationName:"voiceConnectorId" type:"string" required:"true"`
}

// String returns the string representation
func (s PutVoiceConnectorEmergencyCallingConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutVoiceConnectorEmergencyCallingConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutVoiceConnectorEmergencyCallingConfigurationInput"}

	if s.EmergencyCallingConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("EmergencyCallingConfiguration"))
	}

	if s.VoiceConnectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorId"))
	}
	if s.EmergencyCallingConfiguration != nil {
		if err := s.EmergencyCallingConfiguration.Validate(); err != nil {
			invalidParams.AddNested("EmergencyCallingConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutVoiceConnectorEmergencyCallingConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.EmergencyCallingConfiguration != nil {
		v := s.EmergencyCallingConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "EmergencyCallingConfiguration", v, metadata)
	}
	if s.VoiceConnectorId != nil {
		v := *s.VoiceConnectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "voiceConnectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type PutVoiceConnectorEmergencyCallingConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The emergency calling configuration details.
	EmergencyCallingConfiguration *EmergencyCallingConfiguration `type:"structure"`
}

// String returns the string representation
func (s PutVoiceConnectorEmergencyCallingConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutVoiceConnectorEmergencyCallingConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.EmergencyCallingConfiguration != nil {
		v := s.EmergencyCallingConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "EmergencyCallingConfiguration", v, metadata)
	}
	return nil
}

const opPutVoiceConnectorEmergencyCallingConfiguration = "PutVoiceConnectorEmergencyCallingConfiguration"

// PutVoiceConnectorEmergencyCallingConfigurationRequest returns a request value for making API operation for
// Amazon Chime.
//
// Puts emergency calling configuration details to the specified Amazon Chime
// Voice Connector, such as emergency phone numbers and calling countries. Origination
// and termination settings must be enabled for the Amazon Chime Voice Connector
// before emergency calling can be configured.
//
//    // Example sending a request using PutVoiceConnectorEmergencyCallingConfigurationRequest.
//    req := client.PutVoiceConnectorEmergencyCallingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/PutVoiceConnectorEmergencyCallingConfiguration
func (c *Client) PutVoiceConnectorEmergencyCallingConfigurationRequest(input *PutVoiceConnectorEmergencyCallingConfigurationInput) PutVoiceConnectorEmergencyCallingConfigurationRequest {
	op := &aws.Operation{
		Name:       opPutVoiceConnectorEmergencyCallingConfiguration,
		HTTPMethod: "PUT",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}/emergency-calling-configuration",
	}

	if input == nil {
		input = &PutVoiceConnectorEmergencyCallingConfigurationInput{}
	}

	req := c.newRequest(op, input, &PutVoiceConnectorEmergencyCallingConfigurationOutput{})

	return PutVoiceConnectorEmergencyCallingConfigurationRequest{Request: req, Input: input, Copy: c.PutVoiceConnectorEmergencyCallingConfigurationRequest}
}

// PutVoiceConnectorEmergencyCallingConfigurationRequest is the request type for the
// PutVoiceConnectorEmergencyCallingConfiguration API operation.
type PutVoiceConnectorEmergencyCallingConfigurationRequest struct {
	*aws.Request
	Input *PutVoiceConnectorEmergencyCallingConfigurationInput
	Copy  func(*PutVoiceConnectorEmergencyCallingConfigurationInput) PutVoiceConnectorEmergencyCallingConfigurationRequest
}

// Send marshals and sends the PutVoiceConnectorEmergencyCallingConfiguration API request.
func (r PutVoiceConnectorEmergencyCallingConfigurationRequest) Send(ctx context.Context) (*PutVoiceConnectorEmergencyCallingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutVoiceConnectorEmergencyCallingConfigurationResponse{
		PutVoiceConnectorEmergencyCallingConfigurationOutput: r.Request.Data.(*PutVoiceConnectorEmergencyCallingConfigurationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutVoiceConnectorEmergencyCallingConfigurationResponse is the response type for the
// PutVoiceConnectorEmergencyCallingConfiguration API operation.
type PutVoiceConnectorEmergencyCallingConfigurationResponse struct {
	*PutVoiceConnectorEmergencyCallingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutVoiceConnectorEmergencyCallingConfiguration request.
func (r *PutVoiceConnectorEmergencyCallingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
