// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eventbridge

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/PutPermissionRequest
type PutPermissionInput struct {
	_ struct{} `type:"structure"`

	// The action that you're enabling the other account to perform. Currently,
	// this must be events:PutEvents.
	//
	// Action is a required field
	Action *string `min:"1" type:"string" required:"true"`

	// This parameter enables you to limit the permission to accounts that fulfill
	// a certain condition, such as being a member of a certain AWS organization.
	// For more information about AWS Organizations, see What Is AWS Organizations?
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_introduction.html)
	// in the AWS Organizations User Guide.
	//
	// If you specify Condition with an AWS organization ID and specify "*" as the
	// value for Principal, you grant permission to all the accounts in the named
	// organization.
	//
	// The Condition is a JSON string that must contain Type, Key, and Value fields.
	Condition *Condition `type:"structure"`

	// The event bus associated with the rule. If you omit this, the default event
	// bus is used.
	EventBusName *string `min:"1" type:"string"`

	// The 12-digit AWS account ID that you are permitting to put events to your
	// default event bus. Specify "*" to permit any account to put events to your
	// default event bus.
	//
	// If you specify "*" without specifying Condition, avoid creating rules that
	// might match undesirable events. To create more secure rules, make sure that
	// the event pattern for each rule contains an account field with a specific
	// account ID to receive events from. Rules with an account field don't match
	// any events sent from other accounts.
	//
	// Principal is a required field
	Principal *string `min:"1" type:"string" required:"true"`

	// An identifier string for the external account that you're granting permissions
	// to. If you later want to revoke the permission for this external account,
	// specify this StatementId when you run RemovePermission.
	//
	// StatementId is a required field
	StatementId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutPermissionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutPermissionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutPermissionInput"}

	if s.Action == nil {
		invalidParams.Add(aws.NewErrParamRequired("Action"))
	}
	if s.Action != nil && len(*s.Action) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Action", 1))
	}
	if s.EventBusName != nil && len(*s.EventBusName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventBusName", 1))
	}

	if s.Principal == nil {
		invalidParams.Add(aws.NewErrParamRequired("Principal"))
	}
	if s.Principal != nil && len(*s.Principal) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Principal", 1))
	}

	if s.StatementId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StatementId"))
	}
	if s.StatementId != nil && len(*s.StatementId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StatementId", 1))
	}
	if s.Condition != nil {
		if err := s.Condition.Validate(); err != nil {
			invalidParams.AddNested("Condition", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/PutPermissionOutput
type PutPermissionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutPermissionOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutPermission = "PutPermission"

// PutPermissionRequest returns a request value for making API operation for
// Amazon EventBridge.
//
// Running PutPermission permits the specified AWS account or AWS organization
// to put events to the specified event bus. Rules in your account are triggered
// by these events arriving to an event bus in your account.
//
// For another account to send events to your account, that external account
// must have a rule with your account's event bus as a target.
//
// To enable multiple AWS accounts to put events to an event bus, run PutPermission
// once for each of these accounts. Or, if all the accounts are members of the
// same AWS organization, you can run PutPermission once specifying Principal
// as "*" and specifying the AWS organization ID in Condition, to grant permissions
// to all accounts in that organization.
//
// If you grant permissions using an organization, then accounts in that organization
// must specify a RoleArn with proper permissions when they use PutTarget to
// add your account's event bus as a target. For more information, see Sending
// and Receiving Events Between AWS Accounts (https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html)
// in the Amazon EventBridge User Guide.
//
// The permission policy on an event bus can't exceed 10 KB in size.
//
//    // Example sending a request using PutPermissionRequest.
//    req := client.PutPermissionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/PutPermission
func (c *Client) PutPermissionRequest(input *PutPermissionInput) PutPermissionRequest {
	op := &aws.Operation{
		Name:       opPutPermission,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutPermissionInput{}
	}

	req := c.newRequest(op, input, &PutPermissionOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutPermissionRequest{Request: req, Input: input, Copy: c.PutPermissionRequest}
}

// PutPermissionRequest is the request type for the
// PutPermission API operation.
type PutPermissionRequest struct {
	*aws.Request
	Input *PutPermissionInput
	Copy  func(*PutPermissionInput) PutPermissionRequest
}

// Send marshals and sends the PutPermission API request.
func (r PutPermissionRequest) Send(ctx context.Context) (*PutPermissionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutPermissionResponse{
		PutPermissionOutput: r.Request.Data.(*PutPermissionOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutPermissionResponse is the response type for the
// PutPermission API operation.
type PutPermissionResponse struct {
	*PutPermissionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutPermission request.
func (r *PutPermissionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
