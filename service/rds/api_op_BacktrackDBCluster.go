// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/BacktrackDBClusterMessage
type BacktrackDBClusterInput struct {
	_ struct{} `type:"structure"`

	// The timestamp of the time to backtrack the DB cluster to, specified in ISO
	// 8601 format. For more information about ISO 8601, see the ISO8601 Wikipedia
	// page. (http://en.wikipedia.org/wiki/ISO_8601)
	//
	// If the specified time is not a consistent time for the DB cluster, Aurora
	// automatically chooses the nearest possible consistent time for the DB cluster.
	//
	// Constraints:
	//
	//    * Must contain a valid ISO 8601 timestamp.
	//
	//    * Can't contain a timestamp set in the future.
	//
	// Example: 2017-07-08T18:00Z
	//
	// BacktrackTo is a required field
	BacktrackTo *time.Time `type:"timestamp" timestampFormat:"iso8601" required:"true"`

	// The DB cluster identifier of the DB cluster to be backtracked. This parameter
	// is stored as a lowercase string.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 alphanumeric characters or hyphens.
	//
	//    * First character must be a letter.
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example: my-cluster1
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// A value that indicates whether to force the DB cluster to backtrack when
	// binary logging is enabled. Otherwise, an error occurs when binary logging
	// is enabled.
	Force *bool `type:"boolean"`

	// A value that indicates whether to backtrack the DB cluster to the earliest
	// possible backtrack time when BacktrackTo is set to a timestamp earlier than
	// the earliest backtrack time. When this parameter is disabled and BacktrackTo
	// is set to a timestamp earlier than the earliest backtrack time, an error
	// occurs.
	UseEarliestTimeOnPointInTimeUnavailable *bool `type:"boolean"`
}

// String returns the string representation
func (s BacktrackDBClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BacktrackDBClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BacktrackDBClusterInput"}

	if s.BacktrackTo == nil {
		invalidParams.Add(aws.NewErrParamRequired("BacktrackTo"))
	}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// This data type is used as a response element in the DescribeDBClusterBacktracks
// action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DBClusterBacktrack
type BacktrackDBClusterOutput struct {
	_ struct{} `type:"structure"`

	// Contains the backtrack identifier.
	BacktrackIdentifier *string `type:"string"`

	// The timestamp of the time at which the backtrack was requested.
	BacktrackRequestCreationTime *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The timestamp of the time to which the DB cluster was backtracked.
	BacktrackTo *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The timestamp of the time from which the DB cluster was backtracked.
	BacktrackedFrom *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// Contains a user-supplied DB cluster identifier. This identifier is the unique
	// key that identifies a DB cluster.
	DBClusterIdentifier *string `type:"string"`

	// The status of the backtrack. This property returns one of the following values:
	//
	//    * applying - The backtrack is currently being applied to or rolled back
	//    from the DB cluster.
	//
	//    * completed - The backtrack has successfully been applied to or rolled
	//    back from the DB cluster.
	//
	//    * failed - An error occurred while the backtrack was applied to or rolled
	//    back from the DB cluster.
	//
	//    * pending - The backtrack is currently pending application to or rollback
	//    from the DB cluster.
	Status *string `type:"string"`
}

// String returns the string representation
func (s BacktrackDBClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opBacktrackDBCluster = "BacktrackDBCluster"

// BacktrackDBClusterRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Backtracks a DB cluster to a specific time, without creating a new DB cluster.
//
// For more information on backtracking, see Backtracking an Aurora DB Cluster
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Managing.Backtrack.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using BacktrackDBClusterRequest.
//    req := client.BacktrackDBClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/BacktrackDBCluster
func (c *Client) BacktrackDBClusterRequest(input *BacktrackDBClusterInput) BacktrackDBClusterRequest {
	op := &aws.Operation{
		Name:       opBacktrackDBCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BacktrackDBClusterInput{}
	}

	req := c.newRequest(op, input, &BacktrackDBClusterOutput{})
	return BacktrackDBClusterRequest{Request: req, Input: input, Copy: c.BacktrackDBClusterRequest}
}

// BacktrackDBClusterRequest is the request type for the
// BacktrackDBCluster API operation.
type BacktrackDBClusterRequest struct {
	*aws.Request
	Input *BacktrackDBClusterInput
	Copy  func(*BacktrackDBClusterInput) BacktrackDBClusterRequest
}

// Send marshals and sends the BacktrackDBCluster API request.
func (r BacktrackDBClusterRequest) Send(ctx context.Context) (*BacktrackDBClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BacktrackDBClusterResponse{
		BacktrackDBClusterOutput: r.Request.Data.(*BacktrackDBClusterOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BacktrackDBClusterResponse is the response type for the
// BacktrackDBCluster API operation.
type BacktrackDBClusterResponse struct {
	*BacktrackDBClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BacktrackDBCluster request.
func (r *BacktrackDBClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
