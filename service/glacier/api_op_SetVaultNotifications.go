// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Provides options to configure notifications that will be sent when specific
// events happen to a vault.
type SetVaultNotificationsInput struct {
	_ struct{} `type:"structure" payload:"VaultNotificationConfig"`

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`

	// Provides options for specifying notification configuration.
	VaultNotificationConfig *VaultNotificationConfig `locationName:"vaultNotificationConfig" type:"structure"`
}

// String returns the string representation
func (s SetVaultNotificationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetVaultNotificationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetVaultNotificationsInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.VaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SetVaultNotificationsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VaultName != nil {
		v := *s.VaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "vaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VaultNotificationConfig != nil {
		v := s.VaultNotificationConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "vaultNotificationConfig", v, metadata)
	}
	return nil
}

type SetVaultNotificationsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetVaultNotificationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SetVaultNotificationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opSetVaultNotifications = "SetVaultNotifications"

// SetVaultNotificationsRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation configures notifications that will be sent when specific events
// happen to a vault. By default, you don't get any notifications.
//
// To configure vault notifications, send a PUT request to the notification-configuration
// subresource of the vault. The request should include a JSON document that
// provides an Amazon SNS topic and specific events for which you want Amazon
// S3 Glacier to send notifications to the topic.
//
// Amazon SNS topics must grant permission to the vault to be allowed to publish
// notifications to the topic. You can configure a vault to publish a notification
// for the following vault events:
//
//    * ArchiveRetrievalCompleted This event occurs when a job that was initiated
//    for an archive retrieval is completed (InitiateJob). The status of the
//    completed job can be "Succeeded" or "Failed". The notification sent to
//    the SNS topic is the same output as returned from DescribeJob.
//
//    * InventoryRetrievalCompleted This event occurs when a job that was initiated
//    for an inventory retrieval is completed (InitiateJob). The status of the
//    completed job can be "Succeeded" or "Failed". The notification sent to
//    the SNS topic is the same output as returned from DescribeJob.
//
// An AWS account has full permission to perform all operations (actions). However,
// AWS Identity and Access Management (IAM) users don't have any permissions
// by default. You must grant them explicit permission to perform specific actions.
// For more information, see Access Control Using AWS Identity and Access Management
// (IAM) (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html).
//
// For conceptual information and underlying REST API, see Configuring Vault
// Notifications in Amazon S3 Glacier (https://docs.aws.amazon.com/amazonglacier/latest/dev/configuring-notifications.html)
// and Set Vault Notification Configuration (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-notifications-put.html)
// in the Amazon Glacier Developer Guide.
//
//    // Example sending a request using SetVaultNotificationsRequest.
//    req := client.SetVaultNotificationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) SetVaultNotificationsRequest(input *SetVaultNotificationsInput) SetVaultNotificationsRequest {
	op := &aws.Operation{
		Name:       opSetVaultNotifications,
		HTTPMethod: "PUT",
		HTTPPath:   "/{accountId}/vaults/{vaultName}/notification-configuration",
	}

	if input == nil {
		input = &SetVaultNotificationsInput{}
	}

	req := c.newRequest(op, input, &SetVaultNotificationsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SetVaultNotificationsRequest{Request: req, Input: input, Copy: c.SetVaultNotificationsRequest}
}

// SetVaultNotificationsRequest is the request type for the
// SetVaultNotifications API operation.
type SetVaultNotificationsRequest struct {
	*aws.Request
	Input *SetVaultNotificationsInput
	Copy  func(*SetVaultNotificationsInput) SetVaultNotificationsRequest
}

// Send marshals and sends the SetVaultNotifications API request.
func (r SetVaultNotificationsRequest) Send(ctx context.Context) (*SetVaultNotificationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetVaultNotificationsResponse{
		SetVaultNotificationsOutput: r.Request.Data.(*SetVaultNotificationsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetVaultNotificationsResponse is the response type for the
// SetVaultNotifications API operation.
type SetVaultNotificationsResponse struct {
	*SetVaultNotificationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetVaultNotifications request.
func (r *SetVaultNotificationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
