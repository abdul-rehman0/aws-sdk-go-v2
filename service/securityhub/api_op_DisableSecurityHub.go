// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DisableSecurityHubRequest
type DisableSecurityHubInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisableSecurityHubInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisableSecurityHubInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DisableSecurityHubResponse
type DisableSecurityHubOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisableSecurityHubOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisableSecurityHubOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDisableSecurityHub = "DisableSecurityHub"

// DisableSecurityHubRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Disables Security Hub in your account only in the current Region. To disable
// Security Hub in all Regions, you must submit one request per Region where
// you have enabled Security Hub. When you disable Security Hub for a master
// account, it doesn't disable Security Hub for any associated member accounts.
//
// When you disable Security Hub, your existing findings and insights and any
// Security Hub configuration settings are deleted after 90 days and can't be
// recovered. Any standards that were enabled are disabled, and your master
// and member account associations are removed. If you want to save your existing
// findings, you must export them before you disable Security Hub.
//
//    // Example sending a request using DisableSecurityHubRequest.
//    req := client.DisableSecurityHubRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DisableSecurityHub
func (c *Client) DisableSecurityHubRequest(input *DisableSecurityHubInput) DisableSecurityHubRequest {
	op := &aws.Operation{
		Name:       opDisableSecurityHub,
		HTTPMethod: "DELETE",
		HTTPPath:   "/accounts",
	}

	if input == nil {
		input = &DisableSecurityHubInput{}
	}

	req := c.newRequest(op, input, &DisableSecurityHubOutput{})
	return DisableSecurityHubRequest{Request: req, Input: input, Copy: c.DisableSecurityHubRequest}
}

// DisableSecurityHubRequest is the request type for the
// DisableSecurityHub API operation.
type DisableSecurityHubRequest struct {
	*aws.Request
	Input *DisableSecurityHubInput
	Copy  func(*DisableSecurityHubInput) DisableSecurityHubRequest
}

// Send marshals and sends the DisableSecurityHub API request.
func (r DisableSecurityHubRequest) Send(ctx context.Context) (*DisableSecurityHubResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisableSecurityHubResponse{
		DisableSecurityHubOutput: r.Request.Data.(*DisableSecurityHubOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisableSecurityHubResponse is the response type for the
// DisableSecurityHub API operation.
type DisableSecurityHubResponse struct {
	*DisableSecurityHubOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisableSecurityHub request.
func (r *DisableSecurityHubResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
