// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a map resource in your AWS account, which provides map tiles of
// different styles sourced from global location data providers.
func (c *Client) CreateMap(ctx context.Context, params *CreateMapInput, optFns ...func(*Options)) (*CreateMapOutput, error) {
	if params == nil {
		params = &CreateMapInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMap", params, optFns, c.addOperationCreateMapMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMapOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMapInput struct {

	// Specifies the map style selected from an available data provider.
	//
	// This member is required.
	Configuration *types.MapConfiguration

	// The name for the map resource. Requirements:
	//
	// * Must contain only alphanumeric
	// characters (A–Z, a–z, 0–9), hyphens (-), periods (.), and underscores (_).
	//
	// *
	// Must be a unique map resource name.
	//
	// * No spaces allowed. For example,
	// ExampleMap.
	//
	// This member is required.
	MapName *string

	// Specifies the pricing plan for your map resource. For additional details and
	// restrictions on each pricing plan option, see Amazon Location Service pricing
	// (https://aws.amazon.com/location/pricing/).
	//
	// This member is required.
	PricingPlan types.PricingPlan

	// An optional description for the map resource.
	Description *string

	// Applies one or more tags to the map resource. A tag is a key-value pair helps
	// manage, identify, search, and filter your resources by labelling them. Format:
	// "key" : "value" Restrictions:
	//
	// * Maximum 50 tags per resource
	//
	// * Each resource
	// tag must be unique with a maximum of one value.
	//
	// * Maximum key length: 128
	// Unicode characters in UTF-8
	//
	// * Maximum value length: 256 Unicode characters in
	// UTF-8
	//
	// * Can use alphanumeric characters (A–Z, a–z, 0–9), and the following
	// characters: + - = . _ : / @.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateMapOutput struct {

	// The timestamp for when the map resource was created in ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// This member is required.
	CreateTime *time.Time

	// The Amazon Resource Name (ARN) for the map resource. Used to specify a resource
	// across all AWS.
	//
	// * Format example: arn:aws:geo:region:account-id:maps/ExampleMap
	//
	// This member is required.
	MapArn *string

	// The name of the map resource.
	//
	// This member is required.
	MapName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMapMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMap{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMap{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateMapValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMap(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateMap(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "CreateMap",
	}
}
