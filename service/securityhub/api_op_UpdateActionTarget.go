// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/UpdateActionTargetRequest
type UpdateActionTargetInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the custom action target to update.
	//
	// ActionTargetArn is a required field
	ActionTargetArn *string `location:"uri" locationName:"ActionTargetArn" type:"string" required:"true"`

	// The updated description for the custom action target.
	Description *string `type:"string"`

	// The updated name of the custom action target.
	Name *string `type:"string"`
}

// String returns the string representation
func (s UpdateActionTargetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateActionTargetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateActionTargetInput"}

	if s.ActionTargetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActionTargetArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateActionTargetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ActionTargetArn != nil {
		v := *s.ActionTargetArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ActionTargetArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/UpdateActionTargetResponse
type UpdateActionTargetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateActionTargetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateActionTargetOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateActionTarget = "UpdateActionTarget"

// UpdateActionTargetRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Updates the name and description of a custom action target in Security Hub.
//
//    // Example sending a request using UpdateActionTargetRequest.
//    req := client.UpdateActionTargetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/UpdateActionTarget
func (c *Client) UpdateActionTargetRequest(input *UpdateActionTargetInput) UpdateActionTargetRequest {
	op := &aws.Operation{
		Name:       opUpdateActionTarget,
		HTTPMethod: "PATCH",
		HTTPPath:   "/actionTargets/{ActionTargetArn+}",
	}

	if input == nil {
		input = &UpdateActionTargetInput{}
	}

	req := c.newRequest(op, input, &UpdateActionTargetOutput{})
	return UpdateActionTargetRequest{Request: req, Input: input, Copy: c.UpdateActionTargetRequest}
}

// UpdateActionTargetRequest is the request type for the
// UpdateActionTarget API operation.
type UpdateActionTargetRequest struct {
	*aws.Request
	Input *UpdateActionTargetInput
	Copy  func(*UpdateActionTargetInput) UpdateActionTargetRequest
}

// Send marshals and sends the UpdateActionTarget API request.
func (r UpdateActionTargetRequest) Send(ctx context.Context) (*UpdateActionTargetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateActionTargetResponse{
		UpdateActionTargetOutput: r.Request.Data.(*UpdateActionTargetOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateActionTargetResponse is the response type for the
// UpdateActionTarget API operation.
type UpdateActionTargetResponse struct {
	*UpdateActionTargetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateActionTarget request.
func (r *UpdateActionTargetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
