// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure for update webhook request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/UpdateWebhookRequest
type UpdateWebhookInput struct {
	_ struct{} `type:"structure"`

	// Name for a branch, part of an Amplify App.
	BranchName *string `locationName:"branchName" min:"1" type:"string"`

	// Description for a webhook.
	Description *string `locationName:"description" type:"string"`

	// Unique Id for a webhook.
	//
	// WebhookId is a required field
	WebhookId *string `location:"uri" locationName:"webhookId" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateWebhookInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateWebhookInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateWebhookInput"}
	if s.BranchName != nil && len(*s.BranchName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BranchName", 1))
	}

	if s.WebhookId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WebhookId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateWebhookInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.BranchName != nil {
		v := *s.BranchName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "branchName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.WebhookId != nil {
		v := *s.WebhookId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "webhookId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure for the update webhook request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/UpdateWebhookResult
type UpdateWebhookOutput struct {
	_ struct{} `type:"structure"`

	// Webhook structure.
	//
	// Webhook is a required field
	Webhook *Webhook `locationName:"webhook" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateWebhookOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateWebhookOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Webhook != nil {
		v := s.Webhook

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "webhook", v, metadata)
	}
	return nil
}

const opUpdateWebhook = "UpdateWebhook"

// UpdateWebhookRequest returns a request value for making API operation for
// AWS Amplify.
//
// Update a webhook.
//
//    // Example sending a request using UpdateWebhookRequest.
//    req := client.UpdateWebhookRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/UpdateWebhook
func (c *Client) UpdateWebhookRequest(input *UpdateWebhookInput) UpdateWebhookRequest {
	op := &aws.Operation{
		Name:       opUpdateWebhook,
		HTTPMethod: "POST",
		HTTPPath:   "/webhooks/{webhookId}",
	}

	if input == nil {
		input = &UpdateWebhookInput{}
	}

	req := c.newRequest(op, input, &UpdateWebhookOutput{})
	return UpdateWebhookRequest{Request: req, Input: input, Copy: c.UpdateWebhookRequest}
}

// UpdateWebhookRequest is the request type for the
// UpdateWebhook API operation.
type UpdateWebhookRequest struct {
	*aws.Request
	Input *UpdateWebhookInput
	Copy  func(*UpdateWebhookInput) UpdateWebhookRequest
}

// Send marshals and sends the UpdateWebhook API request.
func (r UpdateWebhookRequest) Send(ctx context.Context) (*UpdateWebhookResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateWebhookResponse{
		UpdateWebhookOutput: r.Request.Data.(*UpdateWebhookOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateWebhookResponse is the response type for the
// UpdateWebhook API operation.
type UpdateWebhookResponse struct {
	*UpdateWebhookOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateWebhook request.
func (r *UpdateWebhookResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
