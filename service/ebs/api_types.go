// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ebs

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// A block of data in an Amazon Elastic Block Store snapshot.
type Block struct {
	_ struct{} `type:"structure"`

	// The block index.
	BlockIndex *int64 `type:"integer"`

	// The block token for the block index.
	BlockToken *string `type:"string"`
}

// String returns the string representation
func (s Block) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Block) MarshalFields(e protocol.FieldEncoder) error {
	if s.BlockIndex != nil {
		v := *s.BlockIndex

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BlockIndex", protocol.Int64Value(v), metadata)
	}
	if s.BlockToken != nil {
		v := *s.BlockToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BlockToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// A block of data in an Amazon Elastic Block Store snapshot that is different
// from another snapshot of the same volume/snapshot lineage.
type ChangedBlock struct {
	_ struct{} `type:"structure" sensitive:"true"`

	// The block index.
	BlockIndex *int64 `type:"integer"`

	// The block token for the block index of the FirstSnapshotId specified in the
	// ListChangedBlocks operation. This value is absent if the first snapshot does
	// not have the changed block that is on the second snapshot.
	FirstBlockToken *string `type:"string"`

	// The block token for the block index of the SecondSnapshotId specified in
	// the ListChangedBlocks operation.
	SecondBlockToken *string `type:"string"`
}

// String returns the string representation
func (s ChangedBlock) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ChangedBlock) MarshalFields(e protocol.FieldEncoder) error {
	if s.BlockIndex != nil {
		v := *s.BlockIndex

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BlockIndex", protocol.Int64Value(v), metadata)
	}
	if s.FirstBlockToken != nil {
		v := *s.FirstBlockToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FirstBlockToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SecondBlockToken != nil {
		v := *s.SecondBlockToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SecondBlockToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Describes a tag.
type Tag struct {
	_ struct{} `type:"structure"`

	// The key of the tag.
	Key *string `type:"string"`

	// The value of the tag.
	Value *string `type:"string"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Tag) MarshalFields(e protocol.FieldEncoder) error {
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Key", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Value != nil {
		v := *s.Value

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Value", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}
