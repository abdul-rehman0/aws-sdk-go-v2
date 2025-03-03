// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to set the position of a receipt rule in a receipt rule
// set. You use receipt rule sets to receive email with Amazon SES. For more
// information, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/SetReceiptRulePositionRequest
type SetReceiptRulePositionInput struct {
	_ struct{} `type:"structure"`

	// The name of the receipt rule after which to place the specified receipt rule.
	After *string `type:"string"`

	// The name of the receipt rule to reposition.
	//
	// RuleName is a required field
	RuleName *string `type:"string" required:"true"`

	// The name of the receipt rule set that contains the receipt rule to reposition.
	//
	// RuleSetName is a required field
	RuleSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SetReceiptRulePositionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetReceiptRulePositionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetReceiptRulePositionInput"}

	if s.RuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleName"))
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
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/SetReceiptRulePositionResponse
type SetReceiptRulePositionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetReceiptRulePositionOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetReceiptRulePosition = "SetReceiptRulePosition"

// SetReceiptRulePositionRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Sets the position of the specified receipt rule in the receipt rule set.
//
// For information about managing receipt rules, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-managing-receipt-rules.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using SetReceiptRulePositionRequest.
//    req := client.SetReceiptRulePositionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/SetReceiptRulePosition
func (c *Client) SetReceiptRulePositionRequest(input *SetReceiptRulePositionInput) SetReceiptRulePositionRequest {
	op := &aws.Operation{
		Name:       opSetReceiptRulePosition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetReceiptRulePositionInput{}
	}

	req := c.newRequest(op, input, &SetReceiptRulePositionOutput{})
	return SetReceiptRulePositionRequest{Request: req, Input: input, Copy: c.SetReceiptRulePositionRequest}
}

// SetReceiptRulePositionRequest is the request type for the
// SetReceiptRulePosition API operation.
type SetReceiptRulePositionRequest struct {
	*aws.Request
	Input *SetReceiptRulePositionInput
	Copy  func(*SetReceiptRulePositionInput) SetReceiptRulePositionRequest
}

// Send marshals and sends the SetReceiptRulePosition API request.
func (r SetReceiptRulePositionRequest) Send(ctx context.Context) (*SetReceiptRulePositionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetReceiptRulePositionResponse{
		SetReceiptRulePositionOutput: r.Request.Data.(*SetReceiptRulePositionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetReceiptRulePositionResponse is the response type for the
// SetReceiptRulePosition API operation.
type SetReceiptRulePositionResponse struct {
	*SetReceiptRulePositionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetReceiptRulePosition request.
func (r *SetReceiptRulePositionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
