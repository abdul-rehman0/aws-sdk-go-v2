// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/StartBackupJobInput
type StartBackupJobInput struct {
	_ struct{} `type:"structure"`

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and
	// the AWS Region where they are created. They consist of lowercase letters,
	// numbers, and hyphens.
	//
	// BackupVaultName is a required field
	BackupVaultName *string `type:"string" required:"true"`

	// The amount of time AWS Backup attempts a backup before canceling the job
	// and returning an error.
	CompleteWindowMinutes *int64 `type:"long"`

	// Specifies the IAM role ARN used to create the target recovery point; for
	// example, arn:aws:iam::123456789012:role/S3Access.
	//
	// IamRoleArn is a required field
	IamRoleArn *string `type:"string" required:"true"`

	// A customer chosen string that can be used to distinguish between calls to
	// StartBackupJob. Idempotency tokens time out after one hour. Therefore, if
	// you call StartBackupJob multiple times with the same idempotency token within
	// one hour, AWS Backup recognizes that you are requesting only one backup job
	// and initiates only one. If you change the idempotency token for each call,
	// AWS Backup recognizes that you are requesting to start multiple backups.
	IdempotencyToken *string `type:"string"`

	// The lifecycle defines when a protected resource is transitioned to cold storage
	// and when it expires. AWS Backup will transition and expire backups automatically
	// according to the lifecycle that you define.
	//
	// Backups transitioned to cold storage must be stored in cold storage for a
	// minimum of 90 days. Therefore, the “expire after days” setting must be
	// 90 days greater than the “transition to cold after days” setting. The
	// “transition to cold after days” setting cannot be changed after a backup
	// has been transitioned to cold.
	Lifecycle *Lifecycle `type:"structure"`

	// To help organize your resources, you can assign your own metadata to the
	// resources that you create. Each tag is a key-value pair.
	RecoveryPointTags map[string]string `type:"map"`

	// An Amazon Resource Name (ARN) that uniquely identifies a resource. The format
	// of the ARN depends on the resource type.
	//
	// ResourceArn is a required field
	ResourceArn *string `type:"string" required:"true"`

	// The amount of time in minutes before beginning a backup.
	StartWindowMinutes *int64 `type:"long"`
}

// String returns the string representation
func (s StartBackupJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartBackupJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartBackupJobInput"}

	if s.BackupVaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupVaultName"))
	}

	if s.IamRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("IamRoleArn"))
	}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartBackupJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.BackupVaultName != nil {
		v := *s.BackupVaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BackupVaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CompleteWindowMinutes != nil {
		v := *s.CompleteWindowMinutes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CompleteWindowMinutes", protocol.Int64Value(v), metadata)
	}
	if s.IamRoleArn != nil {
		v := *s.IamRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IamRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IdempotencyToken != nil {
		v := *s.IdempotencyToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IdempotencyToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Lifecycle != nil {
		v := s.Lifecycle

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Lifecycle", v, metadata)
	}
	if s.RecoveryPointTags != nil {
		v := s.RecoveryPointTags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "RecoveryPointTags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.ResourceArn != nil {
		v := *s.ResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ResourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StartWindowMinutes != nil {
		v := *s.StartWindowMinutes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StartWindowMinutes", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/StartBackupJobOutput
type StartBackupJobOutput struct {
	_ struct{} `type:"structure"`

	// Uniquely identifies a request to AWS Backup to back up a resource.
	BackupJobId *string `type:"string"`

	// The date and time that a backup job is started, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time `type:"timestamp" timestampFormat:"unix"`

	// An ARN that uniquely identifies a recovery point; for example, arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	RecoveryPointArn *string `type:"string"`
}

// String returns the string representation
func (s StartBackupJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartBackupJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BackupJobId != nil {
		v := *s.BackupJobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BackupJobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationDate != nil {
		v := *s.CreationDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.RecoveryPointArn != nil {
		v := *s.RecoveryPointArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RecoveryPointArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opStartBackupJob = "StartBackupJob"

// StartBackupJobRequest returns a request value for making API operation for
// AWS Backup.
//
// Starts a job to create a one-time backup of the specified resource.
//
//    // Example sending a request using StartBackupJobRequest.
//    req := client.StartBackupJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/StartBackupJob
func (c *Client) StartBackupJobRequest(input *StartBackupJobInput) StartBackupJobRequest {
	op := &aws.Operation{
		Name:       opStartBackupJob,
		HTTPMethod: "PUT",
		HTTPPath:   "/backup-jobs",
	}

	if input == nil {
		input = &StartBackupJobInput{}
	}

	req := c.newRequest(op, input, &StartBackupJobOutput{})
	return StartBackupJobRequest{Request: req, Input: input, Copy: c.StartBackupJobRequest}
}

// StartBackupJobRequest is the request type for the
// StartBackupJob API operation.
type StartBackupJobRequest struct {
	*aws.Request
	Input *StartBackupJobInput
	Copy  func(*StartBackupJobInput) StartBackupJobRequest
}

// Send marshals and sends the StartBackupJob API request.
func (r StartBackupJobRequest) Send(ctx context.Context) (*StartBackupJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartBackupJobResponse{
		StartBackupJobOutput: r.Request.Data.(*StartBackupJobOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartBackupJobResponse is the response type for the
// StartBackupJob API operation.
type StartBackupJobResponse struct {
	*StartBackupJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartBackupJob request.
func (r *StartBackupJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
