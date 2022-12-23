// Code generated by smithy-go-codegen DO NOT EDIT.

package connectparticipant

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCompleteAttachmentUpload struct {
}

func (*validateOpCompleteAttachmentUpload) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCompleteAttachmentUpload) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CompleteAttachmentUploadInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCompleteAttachmentUploadInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateParticipantConnection struct {
}

func (*validateOpCreateParticipantConnection) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateParticipantConnection) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateParticipantConnectionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateParticipantConnectionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisconnectParticipant struct {
}

func (*validateOpDisconnectParticipant) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisconnectParticipant) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisconnectParticipantInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisconnectParticipantInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetAttachment struct {
}

func (*validateOpGetAttachment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetAttachment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetAttachmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetAttachmentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetTranscript struct {
}

func (*validateOpGetTranscript) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTranscript) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTranscriptInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTranscriptInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSendEvent struct {
}

func (*validateOpSendEvent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSendEvent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SendEventInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSendEventInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSendMessage struct {
}

func (*validateOpSendMessage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSendMessage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SendMessageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSendMessageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartAttachmentUpload struct {
}

func (*validateOpStartAttachmentUpload) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartAttachmentUpload) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartAttachmentUploadInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartAttachmentUploadInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCompleteAttachmentUploadValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCompleteAttachmentUpload{}, middleware.After)
}

func addOpCreateParticipantConnectionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateParticipantConnection{}, middleware.After)
}

func addOpDisconnectParticipantValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisconnectParticipant{}, middleware.After)
}

func addOpGetAttachmentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetAttachment{}, middleware.After)
}

func addOpGetTranscriptValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTranscript{}, middleware.After)
}

func addOpSendEventValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSendEvent{}, middleware.After)
}

func addOpSendMessageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSendMessage{}, middleware.After)
}

func addOpStartAttachmentUploadValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartAttachmentUpload{}, middleware.After)
}

func validateOpCompleteAttachmentUploadInput(v *CompleteAttachmentUploadInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CompleteAttachmentUploadInput"}
	if v.AttachmentIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AttachmentIds"))
	}
	if v.ClientToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientToken"))
	}
	if v.ConnectionToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectionToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateParticipantConnectionInput(v *CreateParticipantConnectionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateParticipantConnectionInput"}
	if v.ParticipantToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ParticipantToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisconnectParticipantInput(v *DisconnectParticipantInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisconnectParticipantInput"}
	if v.ConnectionToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectionToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetAttachmentInput(v *GetAttachmentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetAttachmentInput"}
	if v.AttachmentId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AttachmentId"))
	}
	if v.ConnectionToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectionToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetTranscriptInput(v *GetTranscriptInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTranscriptInput"}
	if v.ConnectionToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectionToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSendEventInput(v *SendEventInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SendEventInput"}
	if v.ContentType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContentType"))
	}
	if v.ConnectionToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectionToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSendMessageInput(v *SendMessageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SendMessageInput"}
	if v.ContentType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContentType"))
	}
	if v.Content == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Content"))
	}
	if v.ConnectionToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectionToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartAttachmentUploadInput(v *StartAttachmentUploadInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartAttachmentUploadInput"}
	if v.ContentType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContentType"))
	}
	if v.AttachmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AttachmentName"))
	}
	if v.ClientToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientToken"))
	}
	if v.ConnectionToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConnectionToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
