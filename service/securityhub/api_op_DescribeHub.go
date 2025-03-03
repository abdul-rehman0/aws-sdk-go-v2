// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DescribeHubRequest
type DescribeHubInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the Hub resource to retrieve.
	HubArn *string `type:"string"`
}

// String returns the string representation
func (s DescribeHubInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeHubInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.HubArn != nil {
		v := *s.HubArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "HubArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DescribeHubResponse
type DescribeHubOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the Hub resource retrieved.
	HubArn *string `type:"string"`

	// The date and time when Security Hub was enabled in the account.
	SubscribedAt *string `type:"string"`
}

// String returns the string representation
func (s DescribeHubOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeHubOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.HubArn != nil {
		v := *s.HubArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "HubArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SubscribedAt != nil {
		v := *s.SubscribedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SubscribedAt", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeHub = "DescribeHub"

// DescribeHubRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Returns details about the Hub resource in your account, including the HubArn
// and the time when you enabled Security Hub.
//
//    // Example sending a request using DescribeHubRequest.
//    req := client.DescribeHubRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DescribeHub
func (c *Client) DescribeHubRequest(input *DescribeHubInput) DescribeHubRequest {
	op := &aws.Operation{
		Name:       opDescribeHub,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts",
	}

	if input == nil {
		input = &DescribeHubInput{}
	}

	req := c.newRequest(op, input, &DescribeHubOutput{})
	return DescribeHubRequest{Request: req, Input: input, Copy: c.DescribeHubRequest}
}

// DescribeHubRequest is the request type for the
// DescribeHub API operation.
type DescribeHubRequest struct {
	*aws.Request
	Input *DescribeHubInput
	Copy  func(*DescribeHubInput) DescribeHubRequest
}

// Send marshals and sends the DescribeHub API request.
func (r DescribeHubRequest) Send(ctx context.Context) (*DescribeHubResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeHubResponse{
		DescribeHubOutput: r.Request.Data.(*DescribeHubOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeHubResponse is the response type for the
// DescribeHub API operation.
type DescribeHubResponse struct {
	*DescribeHubOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeHub request.
func (r *DescribeHubResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
