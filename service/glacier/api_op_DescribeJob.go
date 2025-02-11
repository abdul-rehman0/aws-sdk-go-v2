// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Provides options for retrieving a job description.
type DescribeJobInput struct {
	_ struct{} `type:"structure"`

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The ID of the job to describe.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" type:"string" required:"true"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeJobInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}

	if s.VaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeJobInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VaultName != nil {
		v := *s.VaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "vaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Contains the description of an Amazon S3 Glacier job.
type DescribeJobOutput struct {
	_ struct{} `type:"structure"`

	// The job type. This value is either ArchiveRetrieval, InventoryRetrieval,
	// or Select.
	Action ActionCode `type:"string" enum:"true"`

	// The archive ID requested for a select job or archive retrieval. Otherwise,
	// this field is null.
	ArchiveId *string `type:"string"`

	// The SHA256 tree hash of the entire archive for an archive retrieval. For
	// inventory retrieval or select jobs, this field is null.
	ArchiveSHA256TreeHash *string `type:"string"`

	// For an archive retrieval job, this value is the size in bytes of the archive
	// being requested for download. For an inventory retrieval or select job, this
	// value is null.
	ArchiveSizeInBytes *int64 `type:"long"`

	// The job status. When a job is completed, you get the job's output using Get
	// Job Output (GET output).
	Completed *bool `type:"boolean"`

	// The UTC time that the job request completed. While the job is in progress,
	// the value is null.
	CompletionDate *string `type:"string"`

	// The UTC date when the job was created. This value is a string representation
	// of ISO 8601 date format, for example "2012-03-20T17:03:43.221Z".
	CreationDate *string `type:"string"`

	// Parameters used for range inventory retrieval.
	InventoryRetrievalParameters *InventoryRetrievalJobDescription `type:"structure"`

	// For an inventory retrieval job, this value is the size in bytes of the inventory
	// requested for download. For an archive retrieval or select job, this value
	// is null.
	InventorySizeInBytes *int64 `type:"long"`

	// The job description provided when initiating the job.
	JobDescription *string `type:"string"`

	// An opaque string that identifies an Amazon S3 Glacier job.
	JobId *string `type:"string"`

	// Contains the job output location.
	JobOutputPath *string `type:"string"`

	// Contains the location where the data from the select job is stored.
	OutputLocation *OutputLocation `type:"structure"`

	// The retrieved byte range for archive retrieval jobs in the form StartByteValue-EndByteValue.
	// If no range was specified in the archive retrieval, then the whole archive
	// is retrieved. In this case, StartByteValue equals 0 and EndByteValue equals
	// the size of the archive minus 1. For inventory retrieval or select jobs,
	// this field is null.
	RetrievalByteRange *string `type:"string"`

	// For an archive retrieval job, this value is the checksum of the archive.
	// Otherwise, this value is null.
	//
	// The SHA256 tree hash value for the requested range of an archive. If the
	// InitiateJob request for an archive specified a tree-hash aligned range, then
	// this field returns a value.
	//
	// If the whole archive is retrieved, this value is the same as the ArchiveSHA256TreeHash
	// value.
	//
	// This field is null for the following:
	//
	//    * Archive retrieval jobs that specify a range that is not tree-hash aligned
	//
	//    * Archival jobs that specify a range that is equal to the whole archive,
	//    when the job status is InProgress
	//
	//    * Inventory jobs
	//
	//    * Select jobs
	SHA256TreeHash *string `type:"string"`

	// An Amazon SNS topic that receives notification.
	SNSTopic *string `type:"string"`

	// Contains the parameters used for a select.
	SelectParameters *SelectParameters `type:"structure"`

	// The status code can be InProgress, Succeeded, or Failed, and indicates the
	// status of the job.
	StatusCode StatusCode `type:"string" enum:"true"`

	// A friendly message that describes the job status.
	StatusMessage *string `type:"string"`

	// The tier to use for a select or an archive retrieval. Valid values are Expedited,
	// Standard, or Bulk. Standard is the default.
	Tier *string `type:"string"`

	// The Amazon Resource Name (ARN) of the vault from which an archive retrieval
	// was requested.
	VaultARN *string `type:"string"`
}

// String returns the string representation
func (s DescribeJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Action) > 0 {
		v := s.Action

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Action", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ArchiveId != nil {
		v := *s.ArchiveId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ArchiveId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ArchiveSHA256TreeHash != nil {
		v := *s.ArchiveSHA256TreeHash

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ArchiveSHA256TreeHash", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ArchiveSizeInBytes != nil {
		v := *s.ArchiveSizeInBytes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ArchiveSizeInBytes", protocol.Int64Value(v), metadata)
	}
	if s.Completed != nil {
		v := *s.Completed

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Completed", protocol.BoolValue(v), metadata)
	}
	if s.CompletionDate != nil {
		v := *s.CompletionDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CompletionDate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationDate != nil {
		v := *s.CreationDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationDate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.InventoryRetrievalParameters != nil {
		v := s.InventoryRetrievalParameters

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "InventoryRetrievalParameters", v, metadata)
	}
	if s.InventorySizeInBytes != nil {
		v := *s.InventorySizeInBytes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "InventorySizeInBytes", protocol.Int64Value(v), metadata)
	}
	if s.JobDescription != nil {
		v := *s.JobDescription

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "JobDescription", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "JobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobOutputPath != nil {
		v := *s.JobOutputPath

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "JobOutputPath", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OutputLocation != nil {
		v := s.OutputLocation

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "OutputLocation", v, metadata)
	}
	if s.RetrievalByteRange != nil {
		v := *s.RetrievalByteRange

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RetrievalByteRange", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SHA256TreeHash != nil {
		v := *s.SHA256TreeHash

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SHA256TreeHash", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SNSTopic != nil {
		v := *s.SNSTopic

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SNSTopic", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SelectParameters != nil {
		v := s.SelectParameters

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SelectParameters", v, metadata)
	}
	if len(s.StatusCode) > 0 {
		v := s.StatusCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StatusCode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StatusMessage != nil {
		v := *s.StatusMessage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StatusMessage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tier != nil {
		v := *s.Tier

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Tier", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VaultARN != nil {
		v := *s.VaultARN

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VaultARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeJob = "DescribeJob"

// DescribeJobRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation returns information about a job you previously initiated,
// including the job initiation date, the user who initiated the job, the job
// status code/message and the Amazon SNS topic to notify after Amazon S3 Glacier
// (Glacier) completes the job. For more information about initiating a job,
// see InitiateJob.
//
// This operation enables you to check the status of your job. However, it is
// strongly recommended that you set up an Amazon SNS topic and specify it in
// your initiate job request so that Glacier can notify the topic after it completes
// the job.
//
// A job ID will not expire for at least 24 hours after Glacier completes the
// job.
//
// An AWS account has full permission to perform all operations (actions). However,
// AWS Identity and Access Management (IAM) users don't have any permissions
// by default. You must grant them explicit permission to perform specific actions.
// For more information, see Access Control Using AWS Identity and Access Management
// (IAM) (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html).
//
// For more information about using this operation, see the documentation for
// the underlying REST API Describe Job (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-describe-job-get.html)
// in the Amazon Glacier Developer Guide.
//
//    // Example sending a request using DescribeJobRequest.
//    req := client.DescribeJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeJobRequest(input *DescribeJobInput) DescribeJobRequest {
	op := &aws.Operation{
		Name:       opDescribeJob,
		HTTPMethod: "GET",
		HTTPPath:   "/{accountId}/vaults/{vaultName}/jobs/{jobId}",
	}

	if input == nil {
		input = &DescribeJobInput{}
	}

	req := c.newRequest(op, input, &DescribeJobOutput{})
	return DescribeJobRequest{Request: req, Input: input, Copy: c.DescribeJobRequest}
}

// DescribeJobRequest is the request type for the
// DescribeJob API operation.
type DescribeJobRequest struct {
	*aws.Request
	Input *DescribeJobInput
	Copy  func(*DescribeJobInput) DescribeJobRequest
}

// Send marshals and sends the DescribeJob API request.
func (r DescribeJobRequest) Send(ctx context.Context) (*DescribeJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeJobResponse{
		DescribeJobOutput: r.Request.Data.(*DescribeJobOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeJobResponse is the response type for the
// DescribeJob API operation.
type DescribeJobResponse struct {
	*DescribeJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeJob request.
func (r *DescribeJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
