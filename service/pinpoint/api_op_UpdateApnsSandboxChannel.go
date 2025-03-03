// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateApnsSandboxChannelRequest
type UpdateApnsSandboxChannelInput struct {
	_ struct{} `type:"structure" payload:"APNSSandboxChannelRequest"`

	// Specifies the status and settings of the APNs (Apple Push Notification service)
	// sandbox channel for an application.
	//
	// APNSSandboxChannelRequest is a required field
	APNSSandboxChannelRequest *APNSSandboxChannelRequest `type:"structure" required:"true"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateApnsSandboxChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateApnsSandboxChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateApnsSandboxChannelInput"}

	if s.APNSSandboxChannelRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("APNSSandboxChannelRequest"))
	}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateApnsSandboxChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.APNSSandboxChannelRequest != nil {
		v := s.APNSSandboxChannelRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "APNSSandboxChannelRequest", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateApnsSandboxChannelResponse
type UpdateApnsSandboxChannelOutput struct {
	_ struct{} `type:"structure" payload:"APNSSandboxChannelResponse"`

	// Provides information about the status and settings of the APNs (Apple Push
	// Notification service) sandbox channel for an application.
	//
	// APNSSandboxChannelResponse is a required field
	APNSSandboxChannelResponse *APNSSandboxChannelResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateApnsSandboxChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateApnsSandboxChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.APNSSandboxChannelResponse != nil {
		v := s.APNSSandboxChannelResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "APNSSandboxChannelResponse", v, metadata)
	}
	return nil
}

const opUpdateApnsSandboxChannel = "UpdateApnsSandboxChannel"

// UpdateApnsSandboxChannelRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Updates the APNs sandbox channel settings for an application.
//
//    // Example sending a request using UpdateApnsSandboxChannelRequest.
//    req := client.UpdateApnsSandboxChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateApnsSandboxChannel
func (c *Client) UpdateApnsSandboxChannelRequest(input *UpdateApnsSandboxChannelInput) UpdateApnsSandboxChannelRequest {
	op := &aws.Operation{
		Name:       opUpdateApnsSandboxChannel,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/apps/{application-id}/channels/apns_sandbox",
	}

	if input == nil {
		input = &UpdateApnsSandboxChannelInput{}
	}

	req := c.newRequest(op, input, &UpdateApnsSandboxChannelOutput{})
	return UpdateApnsSandboxChannelRequest{Request: req, Input: input, Copy: c.UpdateApnsSandboxChannelRequest}
}

// UpdateApnsSandboxChannelRequest is the request type for the
// UpdateApnsSandboxChannel API operation.
type UpdateApnsSandboxChannelRequest struct {
	*aws.Request
	Input *UpdateApnsSandboxChannelInput
	Copy  func(*UpdateApnsSandboxChannelInput) UpdateApnsSandboxChannelRequest
}

// Send marshals and sends the UpdateApnsSandboxChannel API request.
func (r UpdateApnsSandboxChannelRequest) Send(ctx context.Context) (*UpdateApnsSandboxChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateApnsSandboxChannelResponse{
		UpdateApnsSandboxChannelOutput: r.Request.Data.(*UpdateApnsSandboxChannelOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateApnsSandboxChannelResponse is the response type for the
// UpdateApnsSandboxChannel API operation.
type UpdateApnsSandboxChannelResponse struct {
	*UpdateApnsSandboxChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApnsSandboxChannel request.
func (r *UpdateApnsSandboxChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
