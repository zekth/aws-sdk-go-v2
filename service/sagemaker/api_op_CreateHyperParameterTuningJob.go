// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateHyperParameterTuningJobInput struct {
	_ struct{} `type:"structure"`

	// The HyperParameterTuningJobConfig object that describes the tuning job, including
	// the search strategy, the objective metric used to evaluate training jobs,
	// ranges of parameters to search, and resource limits for the tuning job. For
	// more information, see How Hyperparameter Tuning Works (https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-how-it-works.html).
	//
	// HyperParameterTuningJobConfig is a required field
	HyperParameterTuningJobConfig *HyperParameterTuningJobConfig `type:"structure" required:"true"`

	// The name of the tuning job. This name is the prefix for the names of all
	// training jobs that this tuning job launches. The name must be unique within
	// the same AWS account and AWS Region. The name must have { } to { } characters.
	// Valid characters are a-z, A-Z, 0-9, and : + = @ _ % - (hyphen). The name
	// is not case sensitive.
	//
	// HyperParameterTuningJobName is a required field
	HyperParameterTuningJobName *string `min:"1" type:"string" required:"true"`

	// An array of key-value pairs. You can use tags to categorize your AWS resources
	// in different ways, for example, by purpose, owner, or environment. For more
	// information, see AWS Tagging Strategies (https://aws.amazon.com/answers/account-management/aws-tagging-strategies/).
	//
	// Tags that you specify for the tuning job are also added to all training jobs
	// that the tuning job launches.
	Tags []Tag `type:"list"`

	// The HyperParameterTrainingJobDefinition object that describes the training
	// jobs that this tuning job launches, including static hyperparameters, input
	// data configuration, output data configuration, resource configuration, and
	// stopping condition.
	TrainingJobDefinition *HyperParameterTrainingJobDefinition `type:"structure"`

	// A list of the HyperParameterTrainingJobDefinition objects launched for this
	// tuning job.
	TrainingJobDefinitions []HyperParameterTrainingJobDefinition `min:"1" type:"list"`

	// Specifies the configuration for starting the hyperparameter tuning job using
	// one or more previous tuning jobs as a starting point. The results of previous
	// tuning jobs are used to inform which combinations of hyperparameters to search
	// over in the new tuning job.
	//
	// All training jobs launched by the new hyperparameter tuning job are evaluated
	// by using the objective metric. If you specify IDENTICAL_DATA_AND_ALGORITHM
	// as the WarmStartType value for the warm start configuration, the training
	// job that performs the best in the new tuning job is compared to the best
	// training jobs from the parent tuning jobs. From these, the training job that
	// performs the best as measured by the objective metric is returned as the
	// overall best training job.
	//
	// All training jobs launched by parent hyperparameter tuning jobs and the new
	// hyperparameter tuning jobs count against the limit of training jobs for the
	// tuning job.
	WarmStartConfig *HyperParameterTuningJobWarmStartConfig `type:"structure"`
}

// String returns the string representation
func (s CreateHyperParameterTuningJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateHyperParameterTuningJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateHyperParameterTuningJobInput"}

	if s.HyperParameterTuningJobConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("HyperParameterTuningJobConfig"))
	}

	if s.HyperParameterTuningJobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("HyperParameterTuningJobName"))
	}
	if s.HyperParameterTuningJobName != nil && len(*s.HyperParameterTuningJobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HyperParameterTuningJobName", 1))
	}
	if s.TrainingJobDefinitions != nil && len(s.TrainingJobDefinitions) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TrainingJobDefinitions", 1))
	}
	if s.HyperParameterTuningJobConfig != nil {
		if err := s.HyperParameterTuningJobConfig.Validate(); err != nil {
			invalidParams.AddNested("HyperParameterTuningJobConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.TrainingJobDefinition != nil {
		if err := s.TrainingJobDefinition.Validate(); err != nil {
			invalidParams.AddNested("TrainingJobDefinition", err.(aws.ErrInvalidParams))
		}
	}
	if s.TrainingJobDefinitions != nil {
		for i, v := range s.TrainingJobDefinitions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TrainingJobDefinitions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.WarmStartConfig != nil {
		if err := s.WarmStartConfig.Validate(); err != nil {
			invalidParams.AddNested("WarmStartConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateHyperParameterTuningJobOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the tuning job. Amazon SageMaker assigns
	// an ARN to a hyperparameter tuning job when you create it.
	//
	// HyperParameterTuningJobArn is a required field
	HyperParameterTuningJobArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateHyperParameterTuningJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateHyperParameterTuningJob = "CreateHyperParameterTuningJob"

// CreateHyperParameterTuningJobRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Starts a hyperparameter tuning job. A hyperparameter tuning job finds the
// best version of a model by running many training jobs on your dataset using
// the algorithm you choose and values for hyperparameters within ranges that
// you specify. It then chooses the hyperparameter values that result in a model
// that performs the best, as measured by an objective metric that you choose.
//
//    // Example sending a request using CreateHyperParameterTuningJobRequest.
//    req := client.CreateHyperParameterTuningJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/CreateHyperParameterTuningJob
func (c *Client) CreateHyperParameterTuningJobRequest(input *CreateHyperParameterTuningJobInput) CreateHyperParameterTuningJobRequest {
	op := &aws.Operation{
		Name:       opCreateHyperParameterTuningJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateHyperParameterTuningJobInput{}
	}

	req := c.newRequest(op, input, &CreateHyperParameterTuningJobOutput{})

	return CreateHyperParameterTuningJobRequest{Request: req, Input: input, Copy: c.CreateHyperParameterTuningJobRequest}
}

// CreateHyperParameterTuningJobRequest is the request type for the
// CreateHyperParameterTuningJob API operation.
type CreateHyperParameterTuningJobRequest struct {
	*aws.Request
	Input *CreateHyperParameterTuningJobInput
	Copy  func(*CreateHyperParameterTuningJobInput) CreateHyperParameterTuningJobRequest
}

// Send marshals and sends the CreateHyperParameterTuningJob API request.
func (r CreateHyperParameterTuningJobRequest) Send(ctx context.Context) (*CreateHyperParameterTuningJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateHyperParameterTuningJobResponse{
		CreateHyperParameterTuningJobOutput: r.Request.Data.(*CreateHyperParameterTuningJobOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateHyperParameterTuningJobResponse is the response type for the
// CreateHyperParameterTuningJob API operation.
type CreateHyperParameterTuningJobResponse struct {
	*CreateHyperParameterTuningJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateHyperParameterTuningJob request.
func (r *CreateHyperParameterTuningJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
