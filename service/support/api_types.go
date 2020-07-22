// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package support

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// An attachment to a case communication. The attachment consists of the file
// name and the content of the file.
type Attachment struct {
	_ struct{} `type:"structure"`

	// The content of the attachment file.
	//
	// Data is automatically base64 encoded/decoded by the SDK.
	Data []byte `locationName:"data" type:"blob"`

	// The name of the attachment file.
	FileName *string `locationName:"fileName" type:"string"`
}

// String returns the string representation
func (s Attachment) String() string {
	return awsutil.Prettify(s)
}

// The file name and ID of an attachment to a case communication. You can use
// the ID to retrieve the attachment with the DescribeAttachment operation.
type AttachmentDetails struct {
	_ struct{} `type:"structure"`

	// The ID of the attachment.
	AttachmentId *string `locationName:"attachmentId" type:"string"`

	// The file name of the attachment.
	FileName *string `locationName:"fileName" type:"string"`
}

// String returns the string representation
func (s AttachmentDetails) String() string {
	return awsutil.Prettify(s)
}

// A JSON-formatted object that contains the metadata for a support case. It
// is contained in the response from a DescribeCases request. CaseDetails contains
// the following fields:
//
//    * caseId. The AWS Support case ID requested or returned in the call. The
//    case ID is an alphanumeric string formatted as shown in this example:
//    case-12345678910-2013-c4c1d2bf33c5cf47.
//
//    * categoryCode. The category of problem for the AWS Support case. Corresponds
//    to the CategoryCode values returned by a call to DescribeServices.
//
//    * displayId. The identifier for the case on pages in the AWS Support Center.
//
//    * language. The ISO 639-1 code for the language in which AWS provides
//    support. AWS Support currently supports English ("en") and Japanese ("ja").
//    Language parameters must be passed explicitly for operations that take
//    them.
//
//    * nextToken. A resumption point for pagination.
//
//    * recentCommunications. One or more Communication objects. Fields of these
//    objects are attachments, body, caseId, submittedBy, and timeCreated.
//
//    * serviceCode. The identifier for the AWS service that corresponds to
//    the service code defined in the call to DescribeServices.
//
//    * severityCode. The severity code assigned to the case. Contains one of
//    the values returned by the call to DescribeSeverityLevels. The possible
//    values are: low, normal, high, urgent, and critical.
//
//    * status. The status of the case in the AWS Support Center. Valid values:
//    opened pending-customer-action reopened resolved unassigned work-in-progress
//
//    * subject. The subject line of the case.
//
//    * submittedBy. The email address of the account that submitted the case.
//
//    * timeCreated. The time the case was created, in ISO-8601 format.
type CaseDetails struct {
	_ struct{} `type:"structure"`

	// The AWS Support case ID requested or returned in the call. The case ID is
	// an alphanumeric string formatted as shown in this example: case-12345678910-2013-c4c1d2bf33c5cf47
	CaseId *string `locationName:"caseId" type:"string"`

	// The category of problem for the AWS Support case.
	CategoryCode *string `locationName:"categoryCode" type:"string"`

	// The email addresses that receive copies of communication about the case.
	CcEmailAddresses []string `locationName:"ccEmailAddresses" type:"list"`

	// The ID displayed for the case in the AWS Support Center. This is a numeric
	// string.
	DisplayId *string `locationName:"displayId" type:"string"`

	// The ISO 639-1 code for the language in which AWS provides support. AWS Support
	// currently supports English ("en") and Japanese ("ja"). Language parameters
	// must be passed explicitly for operations that take them.
	Language *string `locationName:"language" type:"string"`

	// The five most recent communications between you and AWS Support Center, including
	// the IDs of any attachments to the communications. Also includes a nextToken
	// that you can use to retrieve earlier communications.
	RecentCommunications *RecentCaseCommunications `locationName:"recentCommunications" type:"structure"`

	// The code for the AWS service. You can get a list of codes and the corresponding
	// service names by calling DescribeServices.
	ServiceCode *string `locationName:"serviceCode" type:"string"`

	// The code for the severity level returned by the call to DescribeSeverityLevels.
	SeverityCode *string `locationName:"severityCode" type:"string"`

	// The status of the case.
	//
	// Valid values:
	//
	//    * opened
	//
	//    * pending-customer-action
	//
	//    * reopened
	//
	//    * resolved
	//
	//    * unassigned
	//
	//    * work-in-progress
	Status *string `locationName:"status" type:"string"`

	// The subject line for the case in the AWS Support Center.
	Subject *string `locationName:"subject" type:"string"`

	// The email address of the account that submitted the case.
	SubmittedBy *string `locationName:"submittedBy" type:"string"`

	// The time that the case was created in the AWS Support Center.
	TimeCreated *string `locationName:"timeCreated" type:"string"`
}

// String returns the string representation
func (s CaseDetails) String() string {
	return awsutil.Prettify(s)
}

// A JSON-formatted name/value pair that represents the category name and category
// code of the problem, selected from the DescribeServices response for each
// AWS service.
type Category struct {
	_ struct{} `type:"structure"`

	// The category code for the support case.
	Code *string `locationName:"code" type:"string"`

	// The category name for the support case.
	Name *string `locationName:"name" type:"string"`
}

// String returns the string representation
func (s Category) String() string {
	return awsutil.Prettify(s)
}

// A communication associated with an AWS Support case. The communication consists
// of the case ID, the message body, attachment information, the submitter of
// the communication, and the date and time of the communication.
type Communication struct {
	_ struct{} `type:"structure"`

	// Information about the attachments to the case communication.
	AttachmentSet []AttachmentDetails `locationName:"attachmentSet" type:"list"`

	// The text of the communication between the customer and AWS Support.
	Body *string `locationName:"body" min:"1" type:"string"`

	// The AWS Support case ID requested or returned in the call. The case ID is
	// an alphanumeric string formatted as shown in this example: case-12345678910-2013-c4c1d2bf33c5cf47
	CaseId *string `locationName:"caseId" type:"string"`

	// The identity of the account that submitted, or responded to, the support
	// case. Customer entries include the role or IAM user as well as the email
	// address. For example, "AdminRole (Role) <someone@example.com>. Entries from
	// the AWS Support team display "Amazon Web Services," and do not show an email
	// address.
	SubmittedBy *string `locationName:"submittedBy" type:"string"`

	// The time the communication was created.
	TimeCreated *string `locationName:"timeCreated" type:"string"`
}

// String returns the string representation
func (s Communication) String() string {
	return awsutil.Prettify(s)
}

// The five most recent communications associated with the case.
type RecentCaseCommunications struct {
	_ struct{} `type:"structure"`

	// The five most recent communications associated with the case.
	Communications []Communication `locationName:"communications" type:"list"`

	// A resumption point for pagination.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s RecentCaseCommunications) String() string {
	return awsutil.Prettify(s)
}

// Information about an AWS service returned by the DescribeServices operation.
type Service struct {
	_ struct{} `type:"structure"`

	// A list of categories that describe the type of support issue a case describes.
	// Categories consist of a category name and a category code. Category names
	// and codes are passed to AWS Support when you call CreateCase.
	Categories []Category `locationName:"categories" type:"list"`

	// The code for an AWS service returned by the DescribeServices response. The
	// name element contains the corresponding friendly name.
	Code *string `locationName:"code" type:"string"`

	// The friendly name for an AWS service. The code element contains the corresponding
	// code.
	Name *string `locationName:"name" type:"string"`
}

// String returns the string representation
func (s Service) String() string {
	return awsutil.Prettify(s)
}

// A code and name pair that represents the severity level of a support case.
// The available values depend on the support plan for the account. For more
// information, see Choosing a severity (https://docs.aws.amazon.com/awssupport/latest/user/case-management.html#choosing-severity)
// in the AWS Support User Guide.
type SeverityLevel struct {
	_ struct{} `type:"structure"`

	// The code for case severity level.
	//
	// Valid values: low | normal | high | urgent | critical
	Code *string `locationName:"code" type:"string"`

	// The name of the severity level that corresponds to the severity level code.
	//
	// The values returned by the API differ from the values that are displayed
	// in the AWS Support Center. For example, for the code "low", the API name
	// is "Low", but the name in the Support Center is "General guidance". These
	// are the Support Center code/name mappings:
	//
	//    * low: General guidance
	//
	//    * normal: System impaired
	//
	//    * high: Production system impaired
	//
	//    * urgent: Production system down
	//
	//    * critical: Business-critical system down
	//
	// For more information, see Choosing a severity (https://docs.aws.amazon.com/awssupport/latest/user/case-management.html#choosing-severity)
	// in the AWS Support User Guide.
	Name *string `locationName:"name" type:"string"`
}

// String returns the string representation
func (s SeverityLevel) String() string {
	return awsutil.Prettify(s)
}

// The container for summary information that relates to the category of the
// Trusted Advisor check.
type TrustedAdvisorCategorySpecificSummary struct {
	_ struct{} `type:"structure"`

	// The summary information about cost savings for a Trusted Advisor check that
	// is in the Cost Optimizing category.
	CostOptimizing *TrustedAdvisorCostOptimizingSummary `locationName:"costOptimizing" type:"structure"`
}

// String returns the string representation
func (s TrustedAdvisorCategorySpecificSummary) String() string {
	return awsutil.Prettify(s)
}

// The description and metadata for a Trusted Advisor check.
type TrustedAdvisorCheckDescription struct {
	_ struct{} `type:"structure"`

	// The category of the Trusted Advisor check.
	//
	// Category is a required field
	Category *string `locationName:"category" type:"string" required:"true"`

	// The description of the Trusted Advisor check, which includes the alert criteria
	// and recommended operations (contains HTML markup).
	//
	// Description is a required field
	Description *string `locationName:"description" type:"string" required:"true"`

	// The unique identifier for the Trusted Advisor check.
	//
	// Id is a required field
	Id *string `locationName:"id" type:"string" required:"true"`

	// The column headings for the data returned by the Trusted Advisor check. The
	// order of the headings corresponds to the order of the data in the Metadata
	// element of the TrustedAdvisorResourceDetail for the check. Metadata contains
	// all the data that is shown in the Excel download, even in those cases where
	// the UI shows just summary data.
	//
	// Metadata is a required field
	Metadata []string `locationName:"metadata" type:"list" required:"true"`

	// The display name for the Trusted Advisor check.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`
}

// String returns the string representation
func (s TrustedAdvisorCheckDescription) String() string {
	return awsutil.Prettify(s)
}

// The refresh status of a Trusted Advisor check.
type TrustedAdvisorCheckRefreshStatus struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the Trusted Advisor check.
	//
	// CheckId is a required field
	CheckId *string `locationName:"checkId" type:"string" required:"true"`

	// The amount of time, in milliseconds, until the Trusted Advisor check is eligible
	// for refresh.
	//
	// MillisUntilNextRefreshable is a required field
	MillisUntilNextRefreshable *int64 `locationName:"millisUntilNextRefreshable" type:"long" required:"true"`

	// The status of the Trusted Advisor check for which a refresh has been requested:
	//
	//    * none: The check is not refreshed or the non-success status exceeds the
	//    timeout
	//
	//    * enqueued: The check refresh requests has entered the refresh queue
	//
	//    * processing: The check refresh request is picked up by the rule processing
	//    engine
	//
	//    * success: The check is successfully refreshed
	//
	//    * abandoned: The check refresh has failed
	//
	// Status is a required field
	Status *string `locationName:"status" type:"string" required:"true"`
}

// String returns the string representation
func (s TrustedAdvisorCheckRefreshStatus) String() string {
	return awsutil.Prettify(s)
}

// The results of a Trusted Advisor check returned by DescribeTrustedAdvisorCheckResult.
type TrustedAdvisorCheckResult struct {
	_ struct{} `type:"structure"`

	// Summary information that relates to the category of the check. Cost Optimizing
	// is the only category that is currently supported.
	//
	// CategorySpecificSummary is a required field
	CategorySpecificSummary *TrustedAdvisorCategorySpecificSummary `locationName:"categorySpecificSummary" type:"structure" required:"true"`

	// The unique identifier for the Trusted Advisor check.
	//
	// CheckId is a required field
	CheckId *string `locationName:"checkId" type:"string" required:"true"`

	// The details about each resource listed in the check result.
	//
	// FlaggedResources is a required field
	FlaggedResources []TrustedAdvisorResourceDetail `locationName:"flaggedResources" type:"list" required:"true"`

	// Details about AWS resources that were analyzed in a call to Trusted Advisor
	// DescribeTrustedAdvisorCheckSummaries.
	//
	// ResourcesSummary is a required field
	ResourcesSummary *TrustedAdvisorResourcesSummary `locationName:"resourcesSummary" type:"structure" required:"true"`

	// The alert status of the check: "ok" (green), "warning" (yellow), "error"
	// (red), or "not_available".
	//
	// Status is a required field
	Status *string `locationName:"status" type:"string" required:"true"`

	// The time of the last refresh of the check.
	//
	// Timestamp is a required field
	Timestamp *string `locationName:"timestamp" type:"string" required:"true"`
}

// String returns the string representation
func (s TrustedAdvisorCheckResult) String() string {
	return awsutil.Prettify(s)
}

// A summary of a Trusted Advisor check result, including the alert status,
// last refresh, and number of resources examined.
type TrustedAdvisorCheckSummary struct {
	_ struct{} `type:"structure"`

	// Summary information that relates to the category of the check. Cost Optimizing
	// is the only category that is currently supported.
	//
	// CategorySpecificSummary is a required field
	CategorySpecificSummary *TrustedAdvisorCategorySpecificSummary `locationName:"categorySpecificSummary" type:"structure" required:"true"`

	// The unique identifier for the Trusted Advisor check.
	//
	// CheckId is a required field
	CheckId *string `locationName:"checkId" type:"string" required:"true"`

	// Specifies whether the Trusted Advisor check has flagged resources.
	HasFlaggedResources *bool `locationName:"hasFlaggedResources" type:"boolean"`

	// Details about AWS resources that were analyzed in a call to Trusted Advisor
	// DescribeTrustedAdvisorCheckSummaries.
	//
	// ResourcesSummary is a required field
	ResourcesSummary *TrustedAdvisorResourcesSummary `locationName:"resourcesSummary" type:"structure" required:"true"`

	// The alert status of the check: "ok" (green), "warning" (yellow), "error"
	// (red), or "not_available".
	//
	// Status is a required field
	Status *string `locationName:"status" type:"string" required:"true"`

	// The time of the last refresh of the check.
	//
	// Timestamp is a required field
	Timestamp *string `locationName:"timestamp" type:"string" required:"true"`
}

// String returns the string representation
func (s TrustedAdvisorCheckSummary) String() string {
	return awsutil.Prettify(s)
}

// The estimated cost savings that might be realized if the recommended operations
// are taken.
type TrustedAdvisorCostOptimizingSummary struct {
	_ struct{} `type:"structure"`

	// The estimated monthly savings that might be realized if the recommended operations
	// are taken.
	//
	// EstimatedMonthlySavings is a required field
	EstimatedMonthlySavings *float64 `locationName:"estimatedMonthlySavings" type:"double" required:"true"`

	// The estimated percentage of savings that might be realized if the recommended
	// operations are taken.
	//
	// EstimatedPercentMonthlySavings is a required field
	EstimatedPercentMonthlySavings *float64 `locationName:"estimatedPercentMonthlySavings" type:"double" required:"true"`
}

// String returns the string representation
func (s TrustedAdvisorCostOptimizingSummary) String() string {
	return awsutil.Prettify(s)
}

// Contains information about a resource identified by a Trusted Advisor check.
type TrustedAdvisorResourceDetail struct {
	_ struct{} `type:"structure"`

	// Specifies whether the AWS resource was ignored by Trusted Advisor because
	// it was marked as suppressed by the user.
	IsSuppressed *bool `locationName:"isSuppressed" type:"boolean"`

	// Additional information about the identified resource. The exact metadata
	// and its order can be obtained by inspecting the TrustedAdvisorCheckDescription
	// object returned by the call to DescribeTrustedAdvisorChecks. Metadata contains
	// all the data that is shown in the Excel download, even in those cases where
	// the UI shows just summary data.
	//
	// Metadata is a required field
	Metadata []string `locationName:"metadata" type:"list" required:"true"`

	// The AWS region in which the identified resource is located.
	Region *string `locationName:"region" type:"string"`

	// The unique identifier for the identified resource.
	//
	// ResourceId is a required field
	ResourceId *string `locationName:"resourceId" type:"string" required:"true"`

	// The status code for the resource identified in the Trusted Advisor check.
	//
	// Status is a required field
	Status *string `locationName:"status" type:"string" required:"true"`
}

// String returns the string representation
func (s TrustedAdvisorResourceDetail) String() string {
	return awsutil.Prettify(s)
}

// Details about AWS resources that were analyzed in a call to Trusted Advisor
// DescribeTrustedAdvisorCheckSummaries.
type TrustedAdvisorResourcesSummary struct {
	_ struct{} `type:"structure"`

	// The number of AWS resources that were flagged (listed) by the Trusted Advisor
	// check.
	//
	// ResourcesFlagged is a required field
	ResourcesFlagged *int64 `locationName:"resourcesFlagged" type:"long" required:"true"`

	// The number of AWS resources ignored by Trusted Advisor because information
	// was unavailable.
	//
	// ResourcesIgnored is a required field
	ResourcesIgnored *int64 `locationName:"resourcesIgnored" type:"long" required:"true"`

	// The number of AWS resources that were analyzed by the Trusted Advisor check.
	//
	// ResourcesProcessed is a required field
	ResourcesProcessed *int64 `locationName:"resourcesProcessed" type:"long" required:"true"`

	// The number of AWS resources ignored by Trusted Advisor because they were
	// marked as suppressed by the user.
	//
	// ResourcesSuppressed is a required field
	ResourcesSuppressed *int64 `locationName:"resourcesSuppressed" type:"long" required:"true"`
}

// String returns the string representation
func (s TrustedAdvisorResourcesSummary) String() string {
	return awsutil.Prettify(s)
}
