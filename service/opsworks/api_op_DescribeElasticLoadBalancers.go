// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeElasticLoadBalancersRequest
type DescribeElasticLoadBalancersInput struct {
	_ struct{} `type:"structure"`

	// A list of layer IDs. The action describes the Elastic Load Balancing instances
	// for the specified layers.
	LayerIds []string `type:"list"`

	// A stack ID. The action describes the stack's Elastic Load Balancing instances.
	StackId *string `type:"string"`
}

// String returns the string representation
func (s DescribeElasticLoadBalancersInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the response to a DescribeElasticLoadBalancers request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeElasticLoadBalancersResult
type DescribeElasticLoadBalancersOutput struct {
	_ struct{} `type:"structure"`

	// A list of ElasticLoadBalancer objects that describe the specified Elastic
	// Load Balancing instances.
	ElasticLoadBalancers []ElasticLoadBalancer `type:"list"`
}

// String returns the string representation
func (s DescribeElasticLoadBalancersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeElasticLoadBalancers = "DescribeElasticLoadBalancers"

// DescribeElasticLoadBalancersRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Describes a stack's Elastic Load Balancing instances.
//
// This call accepts only one resource-identifying parameter.
//
// Required Permissions: To use this action, an IAM user must have a Show, Deploy,
// or Manage permissions level for the stack, or an attached policy that explicitly
// grants permissions. For more information about user permissions, see Managing
// User Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using DescribeElasticLoadBalancersRequest.
//    req := client.DescribeElasticLoadBalancersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeElasticLoadBalancers
func (c *Client) DescribeElasticLoadBalancersRequest(input *DescribeElasticLoadBalancersInput) DescribeElasticLoadBalancersRequest {
	op := &aws.Operation{
		Name:       opDescribeElasticLoadBalancers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeElasticLoadBalancersInput{}
	}

	req := c.newRequest(op, input, &DescribeElasticLoadBalancersOutput{})
	return DescribeElasticLoadBalancersRequest{Request: req, Input: input, Copy: c.DescribeElasticLoadBalancersRequest}
}

// DescribeElasticLoadBalancersRequest is the request type for the
// DescribeElasticLoadBalancers API operation.
type DescribeElasticLoadBalancersRequest struct {
	*aws.Request
	Input *DescribeElasticLoadBalancersInput
	Copy  func(*DescribeElasticLoadBalancersInput) DescribeElasticLoadBalancersRequest
}

// Send marshals and sends the DescribeElasticLoadBalancers API request.
func (r DescribeElasticLoadBalancersRequest) Send(ctx context.Context) (*DescribeElasticLoadBalancersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeElasticLoadBalancersResponse{
		DescribeElasticLoadBalancersOutput: r.Request.Data.(*DescribeElasticLoadBalancersOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeElasticLoadBalancersResponse is the response type for the
// DescribeElasticLoadBalancers API operation.
type DescribeElasticLoadBalancersResponse struct {
	*DescribeElasticLoadBalancersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeElasticLoadBalancers request.
func (r *DescribeElasticLoadBalancersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
