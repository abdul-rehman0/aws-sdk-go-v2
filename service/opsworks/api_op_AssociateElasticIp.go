// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/AssociateElasticIpRequest
type AssociateElasticIpInput struct {
	_ struct{} `type:"structure"`

	// The Elastic IP address.
	//
	// ElasticIp is a required field
	ElasticIp *string `type:"string" required:"true"`

	// The instance ID.
	InstanceId *string `type:"string"`
}

// String returns the string representation
func (s AssociateElasticIpInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateElasticIpInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateElasticIpInput"}

	if s.ElasticIp == nil {
		invalidParams.Add(aws.NewErrParamRequired("ElasticIp"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/AssociateElasticIpOutput
type AssociateElasticIpOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateElasticIpOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateElasticIp = "AssociateElasticIp"

// AssociateElasticIpRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Associates one of the stack's registered Elastic IP addresses with a specified
// instance. The address must first be registered with the stack by calling
// RegisterElasticIp. For more information, see Resource Management (https://docs.aws.amazon.com/opsworks/latest/userguide/resources.html).
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using AssociateElasticIpRequest.
//    req := client.AssociateElasticIpRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/AssociateElasticIp
func (c *Client) AssociateElasticIpRequest(input *AssociateElasticIpInput) AssociateElasticIpRequest {
	op := &aws.Operation{
		Name:       opAssociateElasticIp,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateElasticIpInput{}
	}

	req := c.newRequest(op, input, &AssociateElasticIpOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return AssociateElasticIpRequest{Request: req, Input: input, Copy: c.AssociateElasticIpRequest}
}

// AssociateElasticIpRequest is the request type for the
// AssociateElasticIp API operation.
type AssociateElasticIpRequest struct {
	*aws.Request
	Input *AssociateElasticIpInput
	Copy  func(*AssociateElasticIpInput) AssociateElasticIpRequest
}

// Send marshals and sends the AssociateElasticIp API request.
func (r AssociateElasticIpRequest) Send(ctx context.Context) (*AssociateElasticIpResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateElasticIpResponse{
		AssociateElasticIpOutput: r.Request.Data.(*AssociateElasticIpOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateElasticIpResponse is the response type for the
// AssociateElasticIp API operation.
type AssociateElasticIpResponse struct {
	*AssociateElasticIpOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateElasticIp request.
func (r *AssociateElasticIpResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
