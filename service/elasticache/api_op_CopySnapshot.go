// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a CopySnapshotMessage operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/CopySnapshotMessage
type CopySnapshotInput struct {
	_ struct{} `type:"structure"`

	// The name of an existing snapshot from which to make a copy.
	//
	// SourceSnapshotName is a required field
	SourceSnapshotName *string `type:"string" required:"true"`

	// The Amazon S3 bucket to which the snapshot is exported. This parameter is
	// used only when exporting a snapshot for external access.
	//
	// When using this parameter to export a snapshot, be sure Amazon ElastiCache
	// has the needed permissions to this S3 bucket. For more information, see Step
	// 2: Grant ElastiCache Access to Your Amazon S3 Bucket (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html#backups-exporting-grant-access)
	// in the Amazon ElastiCache User Guide.
	//
	// For more information, see Exporting a Snapshot (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Snapshots.Exporting.html)
	// in the Amazon ElastiCache User Guide.
	TargetBucket *string `type:"string"`

	// A name for the snapshot copy. ElastiCache does not permit overwriting a snapshot,
	// therefore this name must be unique within its context - ElastiCache or an
	// Amazon S3 bucket if exporting.
	//
	// TargetSnapshotName is a required field
	TargetSnapshotName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CopySnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CopySnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CopySnapshotInput"}

	if s.SourceSnapshotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceSnapshotName"))
	}

	if s.TargetSnapshotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetSnapshotName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/CopySnapshotResult
type CopySnapshotOutput struct {
	_ struct{} `type:"structure"`

	// Represents a copy of an entire Redis cluster as of the time when the snapshot
	// was taken.
	Snapshot *Snapshot `type:"structure"`
}

// String returns the string representation
func (s CopySnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opCopySnapshot = "CopySnapshot"

// CopySnapshotRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Makes a copy of an existing snapshot.
//
// This operation is valid for Redis only.
//
// Users or groups that have permissions to use the CopySnapshot operation can
// create their own Amazon S3 buckets and copy snapshots to it. To control access
// to your snapshots, use an IAM policy to control who has the ability to use
// the CopySnapshot operation. For more information about using IAM to control
// the use of ElastiCache operations, see Exporting Snapshots (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html)
// and Authentication & Access Control (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/IAM.html).
//
// You could receive the following error messages.
//
// Error Messages
//
//    * Error Message: The S3 bucket %s is outside of the region. Solution:
//    Create an Amazon S3 bucket in the same region as your snapshot. For more
//    information, see Step 1: Create an Amazon S3 Bucket (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html#backups-exporting-create-s3-bucket)
//    in the ElastiCache User Guide.
//
//    * Error Message: The S3 bucket %s does not exist. Solution: Create an
//    Amazon S3 bucket in the same region as your snapshot. For more information,
//    see Step 1: Create an Amazon S3 Bucket (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html#backups-exporting-create-s3-bucket)
//    in the ElastiCache User Guide.
//
//    * Error Message: The S3 bucket %s is not owned by the authenticated user.
//    Solution: Create an Amazon S3 bucket in the same region as your snapshot.
//    For more information, see Step 1: Create an Amazon S3 Bucket (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html#backups-exporting-create-s3-bucket)
//    in the ElastiCache User Guide.
//
//    * Error Message: The authenticated user does not have sufficient permissions
//    to perform the desired activity. Solution: Contact your system administrator
//    to get the needed permissions.
//
//    * Error Message: The S3 bucket %s already contains an object with key
//    %s. Solution: Give the TargetSnapshotName a new and unique value. If exporting
//    a snapshot, you could alternatively create a new Amazon S3 bucket and
//    use this same value for TargetSnapshotName.
//
//    * Error Message: ElastiCache has not been granted READ permissions %s
//    on the S3 Bucket. Solution: Add List and Read permissions on the bucket.
//    For more information, see Step 2: Grant ElastiCache Access to Your Amazon
//    S3 Bucket (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html#backups-exporting-grant-access)
//    in the ElastiCache User Guide.
//
//    * Error Message: ElastiCache has not been granted WRITE permissions %s
//    on the S3 Bucket. Solution: Add Upload/Delete permissions on the bucket.
//    For more information, see Step 2: Grant ElastiCache Access to Your Amazon
//    S3 Bucket (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html#backups-exporting-grant-access)
//    in the ElastiCache User Guide.
//
//    * Error Message: ElastiCache has not been granted READ_ACP permissions
//    %s on the S3 Bucket. Solution: Add View Permissions on the bucket. For
//    more information, see Step 2: Grant ElastiCache Access to Your Amazon
//    S3 Bucket (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html#backups-exporting-grant-access)
//    in the ElastiCache User Guide.
//
//    // Example sending a request using CopySnapshotRequest.
//    req := client.CopySnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/CopySnapshot
func (c *Client) CopySnapshotRequest(input *CopySnapshotInput) CopySnapshotRequest {
	op := &aws.Operation{
		Name:       opCopySnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CopySnapshotInput{}
	}

	req := c.newRequest(op, input, &CopySnapshotOutput{})
	return CopySnapshotRequest{Request: req, Input: input, Copy: c.CopySnapshotRequest}
}

// CopySnapshotRequest is the request type for the
// CopySnapshot API operation.
type CopySnapshotRequest struct {
	*aws.Request
	Input *CopySnapshotInput
	Copy  func(*CopySnapshotInput) CopySnapshotRequest
}

// Send marshals and sends the CopySnapshot API request.
func (r CopySnapshotRequest) Send(ctx context.Context) (*CopySnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CopySnapshotResponse{
		CopySnapshotOutput: r.Request.Data.(*CopySnapshotOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CopySnapshotResponse is the response type for the
// CopySnapshot API operation.
type CopySnapshotResponse struct {
	*CopySnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CopySnapshot request.
func (r *CopySnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
