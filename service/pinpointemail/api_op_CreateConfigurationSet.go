// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to create a configuration set.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/CreateConfigurationSetRequest
type CreateConfigurationSetInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `type:"string" required:"true"`

	// An object that defines the dedicated IP pool that is used to send emails
	// that you send using the configuration set.
	DeliveryOptions *DeliveryOptions `type:"structure"`

	// An object that defines whether or not Amazon Pinpoint collects reputation
	// metrics for the emails that you send that use the configuration set.
	ReputationOptions *ReputationOptions `type:"structure"`

	// An object that defines whether or not Amazon Pinpoint can send email that
	// you send using the configuration set.
	SendingOptions *SendingOptions `type:"structure"`

	// An array of objects that define the tags (keys and values) that you want
	// to associate with the configuration set.
	Tags []Tag `type:"list"`

	// An object that defines the open and click tracking options for emails that
	// you send using the configuration set.
	TrackingOptions *TrackingOptions `type:"structure"`
}

// String returns the string representation
func (s CreateConfigurationSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateConfigurationSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateConfigurationSetInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.TrackingOptions != nil {
		if err := s.TrackingOptions.Validate(); err != nil {
			invalidParams.AddNested("TrackingOptions", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateConfigurationSetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ConfigurationSetName != nil {
		v := *s.ConfigurationSetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ConfigurationSetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DeliveryOptions != nil {
		v := s.DeliveryOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DeliveryOptions", v, metadata)
	}
	if s.ReputationOptions != nil {
		v := s.ReputationOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ReputationOptions", v, metadata)
	}
	if s.SendingOptions != nil {
		v := s.SendingOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SendingOptions", v, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.TrackingOptions != nil {
		v := s.TrackingOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TrackingOptions", v, metadata)
	}
	return nil
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/CreateConfigurationSetResponse
type CreateConfigurationSetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateConfigurationSetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateConfigurationSetOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCreateConfigurationSet = "CreateConfigurationSet"

// CreateConfigurationSetRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Create a configuration set. Configuration sets are groups of rules that you
// can apply to the emails you send using Amazon Pinpoint. You apply a configuration
// set to an email by including a reference to the configuration set in the
// headers of the email. When you apply a configuration set to an email, all
// of the rules in that configuration set are applied to the email.
//
//    // Example sending a request using CreateConfigurationSetRequest.
//    req := client.CreateConfigurationSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/CreateConfigurationSet
func (c *Client) CreateConfigurationSetRequest(input *CreateConfigurationSetInput) CreateConfigurationSetRequest {
	op := &aws.Operation{
		Name:       opCreateConfigurationSet,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/email/configuration-sets",
	}

	if input == nil {
		input = &CreateConfigurationSetInput{}
	}

	req := c.newRequest(op, input, &CreateConfigurationSetOutput{})
	return CreateConfigurationSetRequest{Request: req, Input: input, Copy: c.CreateConfigurationSetRequest}
}

// CreateConfigurationSetRequest is the request type for the
// CreateConfigurationSet API operation.
type CreateConfigurationSetRequest struct {
	*aws.Request
	Input *CreateConfigurationSetInput
	Copy  func(*CreateConfigurationSetInput) CreateConfigurationSetRequest
}

// Send marshals and sends the CreateConfigurationSet API request.
func (r CreateConfigurationSetRequest) Send(ctx context.Context) (*CreateConfigurationSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateConfigurationSetResponse{
		CreateConfigurationSetOutput: r.Request.Data.(*CreateConfigurationSetOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateConfigurationSetResponse is the response type for the
// CreateConfigurationSet API operation.
type CreateConfigurationSetResponse struct {
	*CreateConfigurationSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateConfigurationSet request.
func (r *CreateConfigurationSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
