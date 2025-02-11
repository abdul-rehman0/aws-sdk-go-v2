// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/ModifyReplicationTaskMessage
type ModifyReplicationTaskInput struct {
	_ struct{} `type:"structure"`

	// Indicates when you want a change data capture (CDC) operation to start. Use
	// either CdcStartPosition or CdcStartTime to specify when you want a CDC operation
	// to start. Specifying both values results in an error.
	//
	// The value can be in date, checkpoint, or LSN/SCN format.
	//
	// Date Example: --cdc-start-position “2018-03-08T12:12:12”
	//
	// Checkpoint Example: --cdc-start-position "checkpoint:V1#27#mysql-bin-changelog.157832:1975:-1:2002:677883278264080:mysql-bin-changelog.157832:1876#0#0#*#0#93"
	//
	// LSN Example: --cdc-start-position “mysql-bin-changelog.000024:373”
	CdcStartPosition *string `type:"string"`

	// Indicates the start time for a change data capture (CDC) operation. Use either
	// CdcStartTime or CdcStartPosition to specify when you want a CDC operation
	// to start. Specifying both values results in an error.
	//
	// Timestamp Example: --cdc-start-time “2018-03-08T12:12:12”
	CdcStartTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Indicates when you want a change data capture (CDC) operation to stop. The
	// value can be either server time or commit time.
	//
	// Server time example: --cdc-stop-position “server_time:3018-02-09T12:12:12”
	//
	// Commit time example: --cdc-stop-position “commit_time: 3018-02-09T12:12:12
	// “
	CdcStopPosition *string `type:"string"`

	// The migration type. Valid values: full-load | cdc | full-load-and-cdc
	MigrationType MigrationTypeValue `type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) of the replication task.
	//
	// ReplicationTaskArn is a required field
	ReplicationTaskArn *string `type:"string" required:"true"`

	// The replication task identifier.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 255 alphanumeric characters or hyphens.
	//
	//    * First character must be a letter.
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens.
	ReplicationTaskIdentifier *string `type:"string"`

	// JSON file that contains settings for the task, such as target metadata settings.
	ReplicationTaskSettings *string `type:"string"`

	// When using the AWS CLI or boto3, provide the path of the JSON file that contains
	// the table mappings. Precede the path with file://. When working with the
	// DMS API, provide the JSON as the parameter value, for example: --table-mappings
	// file://mappingfile.json
	TableMappings *string `type:"string"`
}

// String returns the string representation
func (s ModifyReplicationTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyReplicationTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyReplicationTaskInput"}

	if s.ReplicationTaskArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationTaskArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/ModifyReplicationTaskResponse
type ModifyReplicationTaskOutput struct {
	_ struct{} `type:"structure"`

	// The replication task that was modified.
	ReplicationTask *ReplicationTask `type:"structure"`
}

// String returns the string representation
func (s ModifyReplicationTaskOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyReplicationTask = "ModifyReplicationTask"

// ModifyReplicationTaskRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Modifies the specified replication task.
//
// You can't modify the task endpoints. The task must be stopped before you
// can modify it.
//
// For more information about AWS DMS tasks, see Working with Migration Tasks
// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.html) in the
// AWS Database Migration Service User Guide.
//
//    // Example sending a request using ModifyReplicationTaskRequest.
//    req := client.ModifyReplicationTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/ModifyReplicationTask
func (c *Client) ModifyReplicationTaskRequest(input *ModifyReplicationTaskInput) ModifyReplicationTaskRequest {
	op := &aws.Operation{
		Name:       opModifyReplicationTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyReplicationTaskInput{}
	}

	req := c.newRequest(op, input, &ModifyReplicationTaskOutput{})
	return ModifyReplicationTaskRequest{Request: req, Input: input, Copy: c.ModifyReplicationTaskRequest}
}

// ModifyReplicationTaskRequest is the request type for the
// ModifyReplicationTask API operation.
type ModifyReplicationTaskRequest struct {
	*aws.Request
	Input *ModifyReplicationTaskInput
	Copy  func(*ModifyReplicationTaskInput) ModifyReplicationTaskRequest
}

// Send marshals and sends the ModifyReplicationTask API request.
func (r ModifyReplicationTaskRequest) Send(ctx context.Context) (*ModifyReplicationTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyReplicationTaskResponse{
		ModifyReplicationTaskOutput: r.Request.Data.(*ModifyReplicationTaskOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyReplicationTaskResponse is the response type for the
// ModifyReplicationTask API operation.
type ModifyReplicationTaskResponse struct {
	*ModifyReplicationTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyReplicationTask request.
func (r *ModifyReplicationTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
