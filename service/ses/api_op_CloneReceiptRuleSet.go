// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to create a receipt rule set by cloning an existing
// one. You use receipt rule sets to receive email with Amazon SES. For more
// information, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/CloneReceiptRuleSetRequest
type CloneReceiptRuleSetInput struct {
	_ struct{} `type:"structure"`

	// The name of the rule set to clone.
	//
	// OriginalRuleSetName is a required field
	OriginalRuleSetName *string `type:"string" required:"true"`

	// The name of the rule set to create. The name must:
	//
	//    * This value can only contain ASCII letters (a-z, A-Z), numbers (0-9),
	//    underscores (_), or dashes (-).
	//
	//    * Start and end with a letter or number.
	//
	//    * Contain less than 64 characters.
	//
	// RuleSetName is a required field
	RuleSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CloneReceiptRuleSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CloneReceiptRuleSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CloneReceiptRuleSetInput"}

	if s.OriginalRuleSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("OriginalRuleSetName"))
	}

	if s.RuleSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An empty element returned on a successful request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/CloneReceiptRuleSetResponse
type CloneReceiptRuleSetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CloneReceiptRuleSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opCloneReceiptRuleSet = "CloneReceiptRuleSet"

// CloneReceiptRuleSetRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Creates a receipt rule set by cloning an existing one. All receipt rules
// and configurations are copied to the new receipt rule set and are completely
// independent of the source rule set.
//
// For information about setting up rule sets, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-receipt-rule-set.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using CloneReceiptRuleSetRequest.
//    req := client.CloneReceiptRuleSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/CloneReceiptRuleSet
func (c *Client) CloneReceiptRuleSetRequest(input *CloneReceiptRuleSetInput) CloneReceiptRuleSetRequest {
	op := &aws.Operation{
		Name:       opCloneReceiptRuleSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CloneReceiptRuleSetInput{}
	}

	req := c.newRequest(op, input, &CloneReceiptRuleSetOutput{})
	return CloneReceiptRuleSetRequest{Request: req, Input: input, Copy: c.CloneReceiptRuleSetRequest}
}

// CloneReceiptRuleSetRequest is the request type for the
// CloneReceiptRuleSet API operation.
type CloneReceiptRuleSetRequest struct {
	*aws.Request
	Input *CloneReceiptRuleSetInput
	Copy  func(*CloneReceiptRuleSetInput) CloneReceiptRuleSetRequest
}

// Send marshals and sends the CloneReceiptRuleSet API request.
func (r CloneReceiptRuleSetRequest) Send(ctx context.Context) (*CloneReceiptRuleSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CloneReceiptRuleSetResponse{
		CloneReceiptRuleSetOutput: r.Request.Data.(*CloneReceiptRuleSetOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CloneReceiptRuleSetResponse is the response type for the
// CloneReceiptRuleSet API operation.
type CloneReceiptRuleSetResponse struct {
	*CloneReceiptRuleSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CloneReceiptRuleSet request.
func (r *CloneReceiptRuleSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
