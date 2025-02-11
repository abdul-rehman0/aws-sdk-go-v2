// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to specify the Amazon SNS topic to which Amazon SES
// will publish bounce, complaint, or delivery notifications for emails sent
// with that identity as the Source. For information about Amazon SES notifications,
// see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/notifications-via-sns.html).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/SetIdentityNotificationTopicRequest
type SetIdentityNotificationTopicInput struct {
	_ struct{} `type:"structure"`

	// The identity (email address or domain) that you want to set the Amazon SNS
	// topic for.
	//
	// You can only specify a verified identity for this parameter.
	//
	// You can specify an identity by using its name or by using its Amazon Resource
	// Name (ARN). The following examples are all valid identities: sender@example.com,
	// example.com, arn:aws:ses:us-east-1:123456789012:identity/example.com.
	//
	// Identity is a required field
	Identity *string `type:"string" required:"true"`

	// The type of notifications that will be published to the specified Amazon
	// SNS topic.
	//
	// NotificationType is a required field
	NotificationType NotificationType `type:"string" required:"true" enum:"true"`

	// The Amazon Resource Name (ARN) of the Amazon SNS topic. If the parameter
	// is omitted from the request or a null value is passed, SnsTopic is cleared
	// and publishing is disabled.
	SnsTopic *string `type:"string"`
}

// String returns the string representation
func (s SetIdentityNotificationTopicInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetIdentityNotificationTopicInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetIdentityNotificationTopicInput"}

	if s.Identity == nil {
		invalidParams.Add(aws.NewErrParamRequired("Identity"))
	}
	if len(s.NotificationType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("NotificationType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An empty element returned on a successful request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/SetIdentityNotificationTopicResponse
type SetIdentityNotificationTopicOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetIdentityNotificationTopicOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetIdentityNotificationTopic = "SetIdentityNotificationTopic"

// SetIdentityNotificationTopicRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Sets an Amazon Simple Notification Service (Amazon SNS) topic to use when
// delivering notifications. When you use this operation, you specify a verified
// identity, such as an email address or domain. When you send an email that
// uses the chosen identity in the Source field, Amazon SES sends notifications
// to the topic you specified. You can send bounce, complaint, or delivery notifications
// (or any combination of the three) to the Amazon SNS topic that you specify.
//
// You can execute this operation no more than once per second.
//
// For more information about feedback notification, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/notifications.html).
//
//    // Example sending a request using SetIdentityNotificationTopicRequest.
//    req := client.SetIdentityNotificationTopicRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/SetIdentityNotificationTopic
func (c *Client) SetIdentityNotificationTopicRequest(input *SetIdentityNotificationTopicInput) SetIdentityNotificationTopicRequest {
	op := &aws.Operation{
		Name:       opSetIdentityNotificationTopic,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetIdentityNotificationTopicInput{}
	}

	req := c.newRequest(op, input, &SetIdentityNotificationTopicOutput{})
	return SetIdentityNotificationTopicRequest{Request: req, Input: input, Copy: c.SetIdentityNotificationTopicRequest}
}

// SetIdentityNotificationTopicRequest is the request type for the
// SetIdentityNotificationTopic API operation.
type SetIdentityNotificationTopicRequest struct {
	*aws.Request
	Input *SetIdentityNotificationTopicInput
	Copy  func(*SetIdentityNotificationTopicInput) SetIdentityNotificationTopicRequest
}

// Send marshals and sends the SetIdentityNotificationTopic API request.
func (r SetIdentityNotificationTopicRequest) Send(ctx context.Context) (*SetIdentityNotificationTopicResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetIdentityNotificationTopicResponse{
		SetIdentityNotificationTopicOutput: r.Request.Data.(*SetIdentityNotificationTopicOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetIdentityNotificationTopicResponse is the response type for the
// SetIdentityNotificationTopic API operation.
type SetIdentityNotificationTopicResponse struct {
	*SetIdentityNotificationTopicOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetIdentityNotificationTopic request.
func (r *SetIdentityNotificationTopicResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
