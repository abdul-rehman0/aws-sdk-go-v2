// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/DescribeClusterOperationRequest
type DescribeClusterOperationInput struct {
	_ struct{} `type:"structure"`

	// ClusterOperationArn is a required field
	ClusterOperationArn *string `location:"uri" locationName:"clusterOperationArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeClusterOperationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeClusterOperationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeClusterOperationInput"}

	if s.ClusterOperationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterOperationArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeClusterOperationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ClusterOperationArn != nil {
		v := *s.ClusterOperationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "clusterOperationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about a cluster operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/DescribeClusterOperationResponse
type DescribeClusterOperationOutput struct {
	_ struct{} `type:"structure"`

	// Cluster operation information
	ClusterOperationInfo *ClusterOperationInfo `locationName:"clusterOperationInfo" type:"structure"`
}

// String returns the string representation
func (s DescribeClusterOperationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeClusterOperationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClusterOperationInfo != nil {
		v := s.ClusterOperationInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "clusterOperationInfo", v, metadata)
	}
	return nil
}

const opDescribeClusterOperation = "DescribeClusterOperation"

// DescribeClusterOperationRequest returns a request value for making API operation for
// Managed Streaming for Kafka.
//
// Returns a description of the cluster operation specified by the ARN.
//
//    // Example sending a request using DescribeClusterOperationRequest.
//    req := client.DescribeClusterOperationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/DescribeClusterOperation
func (c *Client) DescribeClusterOperationRequest(input *DescribeClusterOperationInput) DescribeClusterOperationRequest {
	op := &aws.Operation{
		Name:       opDescribeClusterOperation,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/operations/{clusterOperationArn}",
	}

	if input == nil {
		input = &DescribeClusterOperationInput{}
	}

	req := c.newRequest(op, input, &DescribeClusterOperationOutput{})
	return DescribeClusterOperationRequest{Request: req, Input: input, Copy: c.DescribeClusterOperationRequest}
}

// DescribeClusterOperationRequest is the request type for the
// DescribeClusterOperation API operation.
type DescribeClusterOperationRequest struct {
	*aws.Request
	Input *DescribeClusterOperationInput
	Copy  func(*DescribeClusterOperationInput) DescribeClusterOperationRequest
}

// Send marshals and sends the DescribeClusterOperation API request.
func (r DescribeClusterOperationRequest) Send(ctx context.Context) (*DescribeClusterOperationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeClusterOperationResponse{
		DescribeClusterOperationOutput: r.Request.Data.(*DescribeClusterOperationOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeClusterOperationResponse is the response type for the
// DescribeClusterOperation API operation.
type DescribeClusterOperationResponse struct {
	*DescribeClusterOperationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeClusterOperation request.
func (r *DescribeClusterOperationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
