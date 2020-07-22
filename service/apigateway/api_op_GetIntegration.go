// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Represents a request to get the integration configuration.
type GetIntegrationInput struct {
	_ struct{} `type:"structure"`

	// [Required] Specifies a get integration request's HTTP method.
	//
	// HttpMethod is a required field
	HttpMethod *string `location:"uri" locationName:"http_method" type:"string" required:"true"`

	// [Required] Specifies a get integration request's resource identifier
	//
	// ResourceId is a required field
	ResourceId *string `location:"uri" locationName:"resource_id" type:"string" required:"true"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetIntegrationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetIntegrationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetIntegrationInput"}

	if s.HttpMethod == nil {
		invalidParams.Add(aws.NewErrParamRequired("HttpMethod"))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetIntegrationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.HttpMethod != nil {
		v := *s.HttpMethod

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "http_method", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceId != nil {
		v := *s.ResourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "resource_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents an HTTP, HTTP_PROXY, AWS, AWS_PROXY, or Mock integration.
//
// In the API Gateway console, the built-in Lambda integration is an AWS integration.
//
// Creating an API (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
type GetIntegrationOutput struct {
	_ struct{} `type:"structure"`

	// A list of request parameters whose values API Gateway caches. To be valid
	// values for cacheKeyParameters, these parameters must also be specified for
	// Method requestParameters.
	CacheKeyParameters []string `locationName:"cacheKeyParameters" type:"list"`

	// Specifies a group of related cached parameters. By default, API Gateway uses
	// the resource ID as the cacheNamespace. You can specify the same cacheNamespace
	// across resources to return the same cached data for requests to different
	// resources.
	CacheNamespace *string `locationName:"cacheNamespace" type:"string"`

	// The (id (https://docs.aws.amazon.com/apigateway/api-reference/resource/vpc-link/#id))
	// of the VpcLink used for the integration when connectionType=VPC_LINK and
	// undefined, otherwise.
	ConnectionId *string `locationName:"connectionId" type:"string"`

	// The type of the network connection to the integration endpoint. The valid
	// value is INTERNET for connections through the public routable internet or
	// VPC_LINK for private connections between API Gateway and a network load balancer
	// in a VPC. The default value is INTERNET.
	ConnectionType ConnectionType `locationName:"connectionType" type:"string" enum:"true"`

	// Specifies how to handle request payload content type conversions. Supported
	// values are CONVERT_TO_BINARY and CONVERT_TO_TEXT, with the following behaviors:
	//
	//    * CONVERT_TO_BINARY: Converts a request payload from a Base64-encoded
	//    string to the corresponding binary blob.
	//
	//    * CONVERT_TO_TEXT: Converts a request payload from a binary blob to a
	//    Base64-encoded string.
	//
	// If this property is not defined, the request payload will be passed through
	// from the method request to integration request without modification, provided
	// that the passthroughBehavior is configured to support payload pass-through.
	ContentHandling ContentHandlingStrategy `locationName:"contentHandling" type:"string" enum:"true"`

	// Specifies the credentials required for the integration, if any. For AWS integrations,
	// three options are available. To specify an IAM Role for API Gateway to assume,
	// use the role's Amazon Resource Name (ARN). To require that the caller's identity
	// be passed through from the request, specify the string arn:aws:iam::\*:user/\*.
	// To use resource-based permissions on supported AWS services, specify null.
	Credentials *string `locationName:"credentials" type:"string"`

	// Specifies the integration's HTTP method type.
	HttpMethod *string `locationName:"httpMethod" type:"string"`

	// Specifies the integration's responses.
	//
	// Example: Get integration responses of a method
	//
	// Request
	//   GET /restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200
	//   HTTP/1.1 Content-Type: application/json Host: apigateway.us-east-1.amazonaws.com
	//   X-Amz-Date: 20160607T191449Z Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160607/us-east-1/apigateway/aws4_request,
	//   SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
	// Response
	//
	// The successful response returns 200 OK status and a payload as follows:
	//  { "_links": { "curies": { "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-integration-response-{rel}.html",
	//  "name": "integrationresponse", "templated": true }, "self": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200",
	//  "title": "200" }, "integrationresponse:delete": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200"
	//  }, "integrationresponse:update": { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200"
	//  } }, "responseParameters": { "method.response.header.Content-Type": "'application/xml'"
	//  }, "responseTemplates": { "application/json": "$util.urlDecode(\"%3CkinesisStreams%3E#foreach($stream
	//  in $input.path('$.StreamNames'))%3Cstream%3E%3Cname%3E$stream%3C/name%3E%3C/stream%3E#end%3C/kinesisStreams%3E\")\n"
	//  }, "statusCode": "200" }
	//
	// Creating an API (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
	IntegrationResponses map[string]IntegrationResponse `locationName:"integrationResponses" type:"map"`

	// Specifies how the method request body of an unmapped content type will be
	// passed through the integration request to the back end without transformation.
	// A content type is unmapped if no mapping template is defined in the integration
	// or the content type does not match any of the mapped content types, as specified
	// in requestTemplates. The valid value is one of the following:
	//
	//    * WHEN_NO_MATCH: passes the method request body through the integration
	//    request to the back end without transformation when the method request
	//    content type does not match any content type associated with the mapping
	//    templates defined in the integration request.
	//
	//    * WHEN_NO_TEMPLATES: passes the method request body through the integration
	//    request to the back end without transformation when no mapping template
	//    is defined in the integration request. If a template is defined when this
	//    option is selected, the method request of an unmapped content-type will
	//    be rejected with an HTTP 415 Unsupported Media Type response.
	//
	//    * NEVER: rejects the method request with an HTTP 415 Unsupported Media
	//    Type response when either the method request content type does not match
	//    any content type associated with the mapping templates defined in the
	//    integration request or no mapping template is defined in the integration
	//    request.
	PassthroughBehavior *string `locationName:"passthroughBehavior" type:"string"`

	// A key-value map specifying request parameters that are passed from the method
	// request to the back end. The key is an integration request parameter name
	// and the associated value is a method request parameter value or static value
	// that must be enclosed within single quotes and pre-encoded as required by
	// the back end. The method request parameter value must match the pattern of
	// method.request.{location}.{name}, where location is querystring, path, or
	// header and name must be a valid and unique method request parameter name.
	RequestParameters map[string]string `locationName:"requestParameters" type:"map"`

	// Represents a map of Velocity templates that are applied on the request payload
	// based on the value of the Content-Type header sent by the client. The content
	// type value is the key in this map, and the template (as a String) is the
	// value.
	RequestTemplates map[string]string `locationName:"requestTemplates" type:"map"`

	// Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000
	// milliseconds or 29 seconds.
	TimeoutInMillis *int64 `locationName:"timeoutInMillis" type:"integer"`

	// Specifies the TLS configuration for an integration.
	TlsConfig *TlsConfig `locationName:"tlsConfig" type:"structure"`

	// Specifies an API method integration type. The valid value is one of the following:
	//
	//    * AWS: for integrating the API method request with an AWS service action,
	//    including the Lambda function-invoking action. With the Lambda function-invoking
	//    action, this is referred to as the Lambda custom integration. With any
	//    other AWS service action, this is known as AWS integration.
	//
	//    * AWS_PROXY: for integrating the API method request with the Lambda function-invoking
	//    action with the client request passed through as-is. This integration
	//    is also referred to as the Lambda proxy integration.
	//
	//    * HTTP: for integrating the API method request with an HTTP endpoint,
	//    including a private HTTP endpoint within a VPC. This integration is also
	//    referred to as the HTTP custom integration.
	//
	//    * HTTP_PROXY: for integrating the API method request with an HTTP endpoint,
	//    including a private HTTP endpoint within a VPC, with the client request
	//    passed through as-is. This is also referred to as the HTTP proxy integration.
	//
	//    * MOCK: for integrating the API method request with API Gateway as a "loop-back"
	//    endpoint without invoking any backend.
	//
	// For the HTTP and HTTP proxy integrations, each integration can specify a
	// protocol (http/https), port and path. Standard 80 and 443 ports are supported
	// as well as custom ports above 1024. An HTTP or HTTP proxy integration with
	// a connectionType of VPC_LINK is referred to as a private integration and
	// uses a VpcLink to connect API Gateway to a network load balancer of a VPC.
	Type IntegrationType `locationName:"type" type:"string" enum:"true"`

	// Specifies Uniform Resource Identifier (URI) of the integration endpoint.
	//
	//    * For HTTP or HTTP_PROXY integrations, the URI must be a fully formed,
	//    encoded HTTP(S) URL according to the RFC-3986 specification (https://en.wikipedia.org/wiki/Uniform_Resource_Identifier),
	//    for either standard integration, where connectionType is not VPC_LINK,
	//    or private integration, where connectionType is VPC_LINK. For a private
	//    HTTP integration, the URI is not used for routing.
	//
	//    * For AWS or AWS_PROXY integrations, the URI is of the form arn:aws:apigateway:{region}:{subdomain.service|service}:path|action/{service_api}.
	//    Here, {Region} is the API Gateway region (e.g., us-east-1); {service}
	//    is the name of the integrated AWS service (e.g., s3); and {subdomain}
	//    is a designated subdomain supported by certain AWS service for fast host-name
	//    lookup. action can be used for an AWS service action-based API, using
	//    an Action={name}&{p1}={v1}&p2={v2}... query string. The ensuing {service_api}
	//    refers to a supported action {name} plus any required input parameters.
	//    Alternatively, path can be used for an AWS service path-based API. The
	//    ensuing service_api refers to the path to an AWS service resource, including
	//    the region of the integrated AWS service, if applicable. For example,
	//    for integration with the S3 API of GetObject (https://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectGET.html),
	//    the uri can be either arn:aws:apigateway:us-west-2:s3:action/GetObject&Bucket={bucket}&Key={key}
	//    or arn:aws:apigateway:us-west-2:s3:path/{bucket}/{key}
	Uri *string `locationName:"uri" type:"string"`
}

// String returns the string representation
func (s GetIntegrationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetIntegrationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CacheKeyParameters != nil {
		v := s.CacheKeyParameters

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "cacheKeyParameters", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.CacheNamespace != nil {
		v := *s.CacheNamespace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "cacheNamespace", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConnectionId != nil {
		v := *s.ConnectionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "connectionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ConnectionType) > 0 {
		v := s.ConnectionType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "connectionType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.ContentHandling) > 0 {
		v := s.ContentHandling

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "contentHandling", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Credentials != nil {
		v := *s.Credentials

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "credentials", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.HttpMethod != nil {
		v := *s.HttpMethod

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "httpMethod", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IntegrationResponses != nil {
		v := s.IntegrationResponses

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "integrationResponses", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.PassthroughBehavior != nil {
		v := *s.PassthroughBehavior

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "passthroughBehavior", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestParameters != nil {
		v := s.RequestParameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "requestParameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.RequestTemplates != nil {
		v := s.RequestTemplates

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "requestTemplates", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.TimeoutInMillis != nil {
		v := *s.TimeoutInMillis

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "timeoutInMillis", protocol.Int64Value(v), metadata)
	}
	if s.TlsConfig != nil {
		v := s.TlsConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "tlsConfig", v, metadata)
	}
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Uri != nil {
		v := *s.Uri

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "uri", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetIntegration = "GetIntegration"

// GetIntegrationRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Get the integration settings.
//
//    // Example sending a request using GetIntegrationRequest.
//    req := client.GetIntegrationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetIntegrationRequest(input *GetIntegrationInput) GetIntegrationRequest {
	op := &aws.Operation{
		Name:       opGetIntegration,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis/{restapi_id}/resources/{resource_id}/methods/{http_method}/integration",
	}

	if input == nil {
		input = &GetIntegrationInput{}
	}

	req := c.newRequest(op, input, &GetIntegrationOutput{})

	return GetIntegrationRequest{Request: req, Input: input, Copy: c.GetIntegrationRequest}
}

// GetIntegrationRequest is the request type for the
// GetIntegration API operation.
type GetIntegrationRequest struct {
	*aws.Request
	Input *GetIntegrationInput
	Copy  func(*GetIntegrationInput) GetIntegrationRequest
}

// Send marshals and sends the GetIntegration API request.
func (r GetIntegrationRequest) Send(ctx context.Context) (*GetIntegrationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetIntegrationResponse{
		GetIntegrationOutput: r.Request.Data.(*GetIntegrationOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetIntegrationResponse is the response type for the
// GetIntegration API operation.
type GetIntegrationResponse struct {
	*GetIntegrationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetIntegration request.
func (r *GetIntegrationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
