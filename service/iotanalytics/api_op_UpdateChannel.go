// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/UpdateChannelRequest
type UpdateChannelInput struct {
	_ struct{} `type:"structure"`

	// The name of the channel to be updated.
	//
	// ChannelName is a required field
	ChannelName *string `location:"uri" locationName:"channelName" min:"1" type:"string" required:"true"`

	// Where channel data is stored.
	ChannelStorage *ChannelStorage `locationName:"channelStorage" type:"structure"`

	// How long, in days, message data is kept for the channel.
	RetentionPeriod *RetentionPeriod `locationName:"retentionPeriod" type:"structure"`
}

// String returns the string representation
func (s UpdateChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateChannelInput"}

	if s.ChannelName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChannelName"))
	}
	if s.ChannelName != nil && len(*s.ChannelName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChannelName", 1))
	}
	if s.ChannelStorage != nil {
		if err := s.ChannelStorage.Validate(); err != nil {
			invalidParams.AddNested("ChannelStorage", err.(aws.ErrInvalidParams))
		}
	}
	if s.RetentionPeriod != nil {
		if err := s.RetentionPeriod.Validate(); err != nil {
			invalidParams.AddNested("RetentionPeriod", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateChannelInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ChannelStorage != nil {
		v := s.ChannelStorage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "channelStorage", v, metadata)
	}
	if s.RetentionPeriod != nil {
		v := s.RetentionPeriod

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "retentionPeriod", v, metadata)
	}
	if s.ChannelName != nil {
		v := *s.ChannelName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "channelName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/UpdateChannelOutput
type UpdateChannelOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateChannel = "UpdateChannel"

// UpdateChannelRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Updates the settings of a channel.
//
//    // Example sending a request using UpdateChannelRequest.
//    req := client.UpdateChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/UpdateChannel
func (c *Client) UpdateChannelRequest(input *UpdateChannelInput) UpdateChannelRequest {
	op := &aws.Operation{
		Name:       opUpdateChannel,
		HTTPMethod: "PUT",
		HTTPPath:   "/channels/{channelName}",
	}

	if input == nil {
		input = &UpdateChannelInput{}
	}

	req := c.newRequest(op, input, &UpdateChannelOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateChannelRequest{Request: req, Input: input, Copy: c.UpdateChannelRequest}
}

// UpdateChannelRequest is the request type for the
// UpdateChannel API operation.
type UpdateChannelRequest struct {
	*aws.Request
	Input *UpdateChannelInput
	Copy  func(*UpdateChannelInput) UpdateChannelRequest
}

// Send marshals and sends the UpdateChannel API request.
func (r UpdateChannelRequest) Send(ctx context.Context) (*UpdateChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateChannelResponse{
		UpdateChannelOutput: r.Request.Data.(*UpdateChannelOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateChannelResponse is the response type for the
// UpdateChannel API operation.
type UpdateChannelResponse struct {
	*UpdateChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateChannel request.
func (r *UpdateChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
