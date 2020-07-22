// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package snowball

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateClusterInput struct {
	_ struct{} `type:"structure"`

	// The ID for the address that you want the cluster shipped to.
	//
	// AddressId is a required field
	AddressId *string `min:"40" type:"string" required:"true"`

	// An optional description of this specific cluster, for example Environmental
	// Data Cluster-01.
	Description *string `min:"1" type:"string"`

	// The forwarding address ID for a cluster. This field is not supported in most
	// regions.
	ForwardingAddressId *string `min:"40" type:"string"`

	// The type of job for this cluster. Currently, the only job type supported
	// for clusters is LOCAL_USE.
	//
	// JobType is a required field
	JobType JobType `type:"string" required:"true" enum:"true"`

	// The KmsKeyARN value that you want to associate with this cluster. KmsKeyARN
	// values are created by using the CreateKey (https://docs.aws.amazon.com/kms/latest/APIReference/API_CreateKey.html)
	// API action in AWS Key Management Service (AWS KMS).
	KmsKeyARN *string `type:"string"`

	// The Amazon Simple Notification Service (Amazon SNS) notification settings
	// for this cluster.
	Notification *Notification `type:"structure"`

	// The resources associated with the cluster job. These resources include Amazon
	// S3 buckets and optional AWS Lambda functions written in the Python language.
	//
	// Resources is a required field
	Resources *JobResource `type:"structure" required:"true"`

	// The RoleARN that you want to associate with this cluster. RoleArn values
	// are created by using the CreateRole (https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html)
	// API action in AWS Identity and Access Management (IAM).
	//
	// RoleARN is a required field
	RoleARN *string `type:"string" required:"true"`

	// The shipping speed for each node in this cluster. This speed doesn't dictate
	// how soon you'll get each Snowball Edge device, rather it represents how quickly
	// each device moves to its destination while in transit. Regional shipping
	// speeds are as follows:
	//
	//    * In Australia, you have access to express shipping. Typically, Snowballs
	//    shipped express are delivered in about a day.
	//
	//    * In the European Union (EU), you have access to express shipping. Typically,
	//    Snowballs shipped express are delivered in about a day. In addition, most
	//    countries in the EU have access to standard shipping, which typically
	//    takes less than a week, one way.
	//
	//    * In India, Snowballs are delivered in one to seven days.
	//
	//    * In the United States of America (US), you have access to one-day shipping
	//    and two-day shipping.
	//
	//    * In Australia, you have access to express shipping. Typically, devices
	//    shipped express are delivered in about a day.
	//
	//    * In the European Union (EU), you have access to express shipping. Typically,
	//    Snowball Edges shipped express are delivered in about a day. In addition,
	//    most countries in the EU have access to standard shipping, which typically
	//    takes less than a week, one way.
	//
	//    * In India, Snowball Edges are delivered in one to seven days.
	//
	//    * In the US, you have access to one-day shipping and two-day shipping.
	//
	// ShippingOption is a required field
	ShippingOption ShippingOption `type:"string" required:"true" enum:"true"`

	// The type of AWS Snowball device to use for this cluster.
	//
	// For cluster jobs, AWS Snowball currently supports only the EDGE device type.
	SnowballType SnowballType `type:"string" enum:"true"`

	// The tax documents required in your AWS Region.
	TaxDocuments *TaxDocuments `type:"structure"`
}

// String returns the string representation
func (s CreateClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateClusterInput"}

	if s.AddressId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AddressId"))
	}
	if s.AddressId != nil && len(*s.AddressId) < 40 {
		invalidParams.Add(aws.NewErrParamMinLen("AddressId", 40))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if s.ForwardingAddressId != nil && len(*s.ForwardingAddressId) < 40 {
		invalidParams.Add(aws.NewErrParamMinLen("ForwardingAddressId", 40))
	}
	if len(s.JobType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("JobType"))
	}

	if s.Resources == nil {
		invalidParams.Add(aws.NewErrParamRequired("Resources"))
	}

	if s.RoleARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleARN"))
	}
	if len(s.ShippingOption) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ShippingOption"))
	}
	if s.Resources != nil {
		if err := s.Resources.Validate(); err != nil {
			invalidParams.AddNested("Resources", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateClusterOutput struct {
	_ struct{} `type:"structure"`

	// The automatically generated ID for a cluster.
	ClusterId *string `min:"39" type:"string"`
}

// String returns the string representation
func (s CreateClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateCluster = "CreateCluster"

// CreateClusterRequest returns a request value for making API operation for
// Amazon Import/Export Snowball.
//
// Creates an empty cluster. Each cluster supports five nodes. You use the CreateJob
// action separately to create the jobs for each of these nodes. The cluster
// does not ship until these five node jobs have been created.
//
//    // Example sending a request using CreateClusterRequest.
//    req := client.CreateClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/CreateCluster
func (c *Client) CreateClusterRequest(input *CreateClusterInput) CreateClusterRequest {
	op := &aws.Operation{
		Name:       opCreateCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateClusterInput{}
	}

	req := c.newRequest(op, input, &CreateClusterOutput{})

	return CreateClusterRequest{Request: req, Input: input, Copy: c.CreateClusterRequest}
}

// CreateClusterRequest is the request type for the
// CreateCluster API operation.
type CreateClusterRequest struct {
	*aws.Request
	Input *CreateClusterInput
	Copy  func(*CreateClusterInput) CreateClusterRequest
}

// Send marshals and sends the CreateCluster API request.
func (r CreateClusterRequest) Send(ctx context.Context) (*CreateClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateClusterResponse{
		CreateClusterOutput: r.Request.Data.(*CreateClusterOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateClusterResponse is the response type for the
// CreateCluster API operation.
type CreateClusterResponse struct {
	*CreateClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCluster request.
func (r *CreateClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
