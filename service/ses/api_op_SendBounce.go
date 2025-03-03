// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to send a bounce message to the sender of an email you
// received through Amazon SES.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/SendBounceRequest
type SendBounceInput struct {
	_ struct{} `type:"structure"`

	// The address to use in the "From" header of the bounce message. This must
	// be an identity that you have verified with Amazon SES.
	//
	// BounceSender is a required field
	BounceSender *string `type:"string" required:"true"`

	// This parameter is used only for sending authorization. It is the ARN of the
	// identity that is associated with the sending authorization policy that permits
	// you to use the address in the "From" header of the bounce. For more information
	// about sending authorization, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
	BounceSenderArn *string `type:"string"`

	// A list of recipients of the bounced message, including the information required
	// to create the Delivery Status Notifications (DSNs) for the recipients. You
	// must specify at least one BouncedRecipientInfo in the list.
	//
	// BouncedRecipientInfoList is a required field
	BouncedRecipientInfoList []BouncedRecipientInfo `type:"list" required:"true"`

	// Human-readable text for the bounce message to explain the failure. If not
	// specified, the text will be auto-generated based on the bounced recipient
	// information.
	Explanation *string `type:"string"`

	// Message-related DSN fields. If not specified, Amazon SES will choose the
	// values.
	MessageDsn *MessageDsn `type:"structure"`

	// The message ID of the message to be bounced.
	//
	// OriginalMessageId is a required field
	OriginalMessageId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SendBounceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendBounceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SendBounceInput"}

	if s.BounceSender == nil {
		invalidParams.Add(aws.NewErrParamRequired("BounceSender"))
	}

	if s.BouncedRecipientInfoList == nil {
		invalidParams.Add(aws.NewErrParamRequired("BouncedRecipientInfoList"))
	}

	if s.OriginalMessageId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OriginalMessageId"))
	}
	if s.BouncedRecipientInfoList != nil {
		for i, v := range s.BouncedRecipientInfoList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "BouncedRecipientInfoList", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.MessageDsn != nil {
		if err := s.MessageDsn.Validate(); err != nil {
			invalidParams.AddNested("MessageDsn", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents a unique message ID.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/SendBounceResponse
type SendBounceOutput struct {
	_ struct{} `type:"structure"`

	// The message ID of the bounce message.
	MessageId *string `type:"string"`
}

// String returns the string representation
func (s SendBounceOutput) String() string {
	return awsutil.Prettify(s)
}

const opSendBounce = "SendBounce"

// SendBounceRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Generates and sends a bounce message to the sender of an email you received
// through Amazon SES. You can only use this API on an email up to 24 hours
// after you receive it.
//
// You cannot use this API to send generic bounces for mail that was not received
// by Amazon SES.
//
// For information about receiving email through Amazon SES, see the Amazon
// SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using SendBounceRequest.
//    req := client.SendBounceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/SendBounce
func (c *Client) SendBounceRequest(input *SendBounceInput) SendBounceRequest {
	op := &aws.Operation{
		Name:       opSendBounce,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SendBounceInput{}
	}

	req := c.newRequest(op, input, &SendBounceOutput{})
	return SendBounceRequest{Request: req, Input: input, Copy: c.SendBounceRequest}
}

// SendBounceRequest is the request type for the
// SendBounce API operation.
type SendBounceRequest struct {
	*aws.Request
	Input *SendBounceInput
	Copy  func(*SendBounceInput) SendBounceRequest
}

// Send marshals and sends the SendBounce API request.
func (r SendBounceRequest) Send(ctx context.Context) (*SendBounceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SendBounceResponse{
		SendBounceOutput: r.Request.Data.(*SendBounceOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SendBounceResponse is the response type for the
// SendBounce API operation.
type SendBounceResponse struct {
	*SendBounceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SendBounce request.
func (r *SendBounceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
