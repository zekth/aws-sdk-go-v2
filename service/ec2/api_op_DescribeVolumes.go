// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeVolumesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The filters.
	//
	//    * attachment.attach-time - The time stamp when the attachment initiated.
	//
	//    * attachment.delete-on-termination - Whether the volume is deleted on
	//    instance termination.
	//
	//    * attachment.device - The device name specified in the block device mapping
	//    (for example, /dev/sda1).
	//
	//    * attachment.instance-id - The ID of the instance the volume is attached
	//    to.
	//
	//    * attachment.status - The attachment state (attaching | attached | detaching).
	//
	//    * availability-zone - The Availability Zone in which the volume was created.
	//
	//    * create-time - The time stamp when the volume was created.
	//
	//    * encrypted - Indicates whether the volume is encrypted (true | false)
	//
	//    * multi-attach-enabled - Indicates whether the volume is enabled for Multi-Attach
	//    (true | false)
	//
	//    * fast-restored - Indicates whether the volume was created from a snapshot
	//    that is enabled for fast snapshot restore (true | false).
	//
	//    * size - The size of the volume, in GiB.
	//
	//    * snapshot-id - The snapshot from which the volume was created.
	//
	//    * status - The status of the volume (creating | available | in-use | deleting
	//    | deleted | error).
	//
	//    * tag:<key> - The key/value combination of a tag assigned to the resource.
	//    Use the tag key in the filter name and the tag value as the filter value.
	//    For example, to find all resources that have a tag with the key Owner
	//    and the value TeamA, specify tag:Owner for the filter name and TeamA for
	//    the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. Use this filter
	//    to find all resources assigned a tag with a specific key, regardless of
	//    the tag value.
	//
	//    * volume-id - The volume ID.
	//
	//    * volume-type - The Amazon EBS volume type. This can be gp2 for General
	//    Purpose SSD, io1 for Provisioned IOPS SSD, st1 for Throughput Optimized
	//    HDD, sc1 for Cold HDD, or standard for Magnetic volumes.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of volume results returned by DescribeVolumes in paginated
	// output. When this parameter is used, DescribeVolumes only returns MaxResults
	// results in a single page along with a NextToken response element. The remaining
	// results of the initial request can be seen by sending another DescribeVolumes
	// request with the returned NextToken value. This value can be between 5 and
	// 500; if MaxResults is given a value larger than 500, only 500 results are
	// returned. If this parameter is not used, then DescribeVolumes returns all
	// results. You cannot specify this parameter and the volume IDs parameter in
	// the same request.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The NextToken value returned from a previous paginated DescribeVolumes request
	// where MaxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// NextToken value. This value is null when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The volume IDs.
	VolumeIds []string `locationName:"VolumeId" locationNameList:"VolumeId" type:"list"`
}

// String returns the string representation
func (s DescribeVolumesInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeVolumesOutput struct {
	_ struct{} `type:"structure"`

	// The NextToken value to include in a future DescribeVolumes request. When
	// the results of a DescribeVolumes request exceed MaxResults, this value can
	// be used to retrieve the next page of results. This value is null when there
	// are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the volumes.
	Volumes []Volume `locationName:"volumeSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeVolumesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeVolumes = "DescribeVolumes"

// DescribeVolumesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified EBS volumes or all of your EBS volumes.
//
// If you are describing a long list of volumes, we recommend that you paginate
// the output to make the list more manageable. The MaxResults parameter sets
// the maximum number of results returned in a single page. If the list of results
// exceeds your MaxResults value, then that number of results is returned along
// with a NextToken value that can be passed to a subsequent DescribeVolumes
// request to retrieve the remaining results.
//
// For more information about EBS volumes, see Amazon EBS Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumes.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DescribeVolumesRequest.
//    req := client.DescribeVolumesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeVolumes
func (c *Client) DescribeVolumesRequest(input *DescribeVolumesInput) DescribeVolumesRequest {
	op := &aws.Operation{
		Name:       opDescribeVolumes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeVolumesInput{}
	}

	req := c.newRequest(op, input, &DescribeVolumesOutput{})

	return DescribeVolumesRequest{Request: req, Input: input, Copy: c.DescribeVolumesRequest}
}

// DescribeVolumesRequest is the request type for the
// DescribeVolumes API operation.
type DescribeVolumesRequest struct {
	*aws.Request
	Input *DescribeVolumesInput
	Copy  func(*DescribeVolumesInput) DescribeVolumesRequest
}

// Send marshals and sends the DescribeVolumes API request.
func (r DescribeVolumesRequest) Send(ctx context.Context) (*DescribeVolumesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVolumesResponse{
		DescribeVolumesOutput: r.Request.Data.(*DescribeVolumesOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeVolumesRequestPaginator returns a paginator for DescribeVolumes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeVolumesRequest(input)
//   p := ec2.NewDescribeVolumesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeVolumesPaginator(req DescribeVolumesRequest) DescribeVolumesPaginator {
	return DescribeVolumesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeVolumesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeVolumesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeVolumesPaginator struct {
	aws.Pager
}

func (p *DescribeVolumesPaginator) CurrentPage() *DescribeVolumesOutput {
	return p.Pager.CurrentPage().(*DescribeVolumesOutput)
}

// DescribeVolumesResponse is the response type for the
// DescribeVolumes API operation.
type DescribeVolumesResponse struct {
	*DescribeVolumesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVolumes request.
func (r *DescribeVolumesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
