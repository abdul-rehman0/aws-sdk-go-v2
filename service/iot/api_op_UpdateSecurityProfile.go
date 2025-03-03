// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateSecurityProfileInput struct {
	_ struct{} `type:"structure"`

	// A list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the profile's behaviors but it is also retained for
	// any metric specified here.
	AdditionalMetricsToRetain []string `locationName:"additionalMetricsToRetain" type:"list"`

	// Where the alerts are sent. (Alerts are always sent to the console.)
	AlertTargets map[string]AlertTarget `locationName:"alertTargets" type:"map"`

	// Specifies the behaviors that, when violated by a device (thing), cause an
	// alert.
	Behaviors []Behavior `locationName:"behaviors" type:"list"`

	// If true, delete all additionalMetricsToRetain defined for this security profile.
	// If any additionalMetricsToRetain are defined in the current invocation an
	// exception occurs.
	DeleteAdditionalMetricsToRetain *bool `locationName:"deleteAdditionalMetricsToRetain" type:"boolean"`

	// If true, delete all alertTargets defined for this security profile. If any
	// alertTargets are defined in the current invocation an exception occurs.
	DeleteAlertTargets *bool `locationName:"deleteAlertTargets" type:"boolean"`

	// If true, delete all behaviors defined for this security profile. If any behaviors
	// are defined in the current invocation an exception occurs.
	DeleteBehaviors *bool `locationName:"deleteBehaviors" type:"boolean"`

	// The expected version of the security profile. A new version is generated
	// whenever the security profile is updated. If you specify a value that is
	// different than the actual version, a VersionConflictException is thrown.
	ExpectedVersion *int64 `location:"querystring" locationName:"expectedVersion" type:"long"`

	// A description of the security profile.
	SecurityProfileDescription *string `locationName:"securityProfileDescription" type:"string"`

	// The name of the security profile you want to update.
	//
	// SecurityProfileName is a required field
	SecurityProfileName *string `location:"uri" locationName:"securityProfileName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateSecurityProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSecurityProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSecurityProfileInput"}

	if s.SecurityProfileName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecurityProfileName"))
	}
	if s.SecurityProfileName != nil && len(*s.SecurityProfileName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecurityProfileName", 1))
	}
	if s.AlertTargets != nil {
		for i, v := range s.AlertTargets {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "AlertTargets", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Behaviors != nil {
		for i, v := range s.Behaviors {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Behaviors", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateSecurityProfileInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AdditionalMetricsToRetain != nil {
		v := s.AdditionalMetricsToRetain

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "additionalMetricsToRetain", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.AlertTargets != nil {
		v := s.AlertTargets

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "alertTargets", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.Behaviors != nil {
		v := s.Behaviors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "behaviors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.DeleteAdditionalMetricsToRetain != nil {
		v := *s.DeleteAdditionalMetricsToRetain

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deleteAdditionalMetricsToRetain", protocol.BoolValue(v), metadata)
	}
	if s.DeleteAlertTargets != nil {
		v := *s.DeleteAlertTargets

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deleteAlertTargets", protocol.BoolValue(v), metadata)
	}
	if s.DeleteBehaviors != nil {
		v := *s.DeleteBehaviors

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deleteBehaviors", protocol.BoolValue(v), metadata)
	}
	if s.SecurityProfileDescription != nil {
		v := *s.SecurityProfileDescription

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "securityProfileDescription", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SecurityProfileName != nil {
		v := *s.SecurityProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "securityProfileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ExpectedVersion != nil {
		v := *s.ExpectedVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "expectedVersion", protocol.Int64Value(v), metadata)
	}
	return nil
}

type UpdateSecurityProfileOutput struct {
	_ struct{} `type:"structure"`

	// A list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the security profile's behaviors but it is also retained
	// for any metric specified here.
	AdditionalMetricsToRetain []string `locationName:"additionalMetricsToRetain" type:"list"`

	// Where the alerts are sent. (Alerts are always sent to the console.)
	AlertTargets map[string]AlertTarget `locationName:"alertTargets" type:"map"`

	// Specifies the behaviors that, when violated by a device (thing), cause an
	// alert.
	Behaviors []Behavior `locationName:"behaviors" type:"list"`

	// The time the security profile was created.
	CreationDate *time.Time `locationName:"creationDate" type:"timestamp" timestampFormat:"unix"`

	// The time the security profile was last modified.
	LastModifiedDate *time.Time `locationName:"lastModifiedDate" type:"timestamp" timestampFormat:"unix"`

	// The ARN of the security profile that was updated.
	SecurityProfileArn *string `locationName:"securityProfileArn" type:"string"`

	// The description of the security profile.
	SecurityProfileDescription *string `locationName:"securityProfileDescription" type:"string"`

	// The name of the security profile that was updated.
	SecurityProfileName *string `locationName:"securityProfileName" min:"1" type:"string"`

	// The updated version of the security profile.
	Version *int64 `locationName:"version" type:"long"`
}

// String returns the string representation
func (s UpdateSecurityProfileOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateSecurityProfileOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AdditionalMetricsToRetain != nil {
		v := s.AdditionalMetricsToRetain

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "additionalMetricsToRetain", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.AlertTargets != nil {
		v := s.AlertTargets

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "alertTargets", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.Behaviors != nil {
		v := s.Behaviors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "behaviors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.CreationDate != nil {
		v := *s.CreationDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "creationDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.LastModifiedDate != nil {
		v := *s.LastModifiedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastModifiedDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.SecurityProfileArn != nil {
		v := *s.SecurityProfileArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "securityProfileArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SecurityProfileDescription != nil {
		v := *s.SecurityProfileDescription

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "securityProfileDescription", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SecurityProfileName != nil {
		v := *s.SecurityProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "securityProfileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opUpdateSecurityProfile = "UpdateSecurityProfile"

// UpdateSecurityProfileRequest returns a request value for making API operation for
// AWS IoT.
//
// Updates a Device Defender security profile.
//
//    // Example sending a request using UpdateSecurityProfileRequest.
//    req := client.UpdateSecurityProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateSecurityProfileRequest(input *UpdateSecurityProfileInput) UpdateSecurityProfileRequest {
	op := &aws.Operation{
		Name:       opUpdateSecurityProfile,
		HTTPMethod: "PATCH",
		HTTPPath:   "/security-profiles/{securityProfileName}",
	}

	if input == nil {
		input = &UpdateSecurityProfileInput{}
	}

	req := c.newRequest(op, input, &UpdateSecurityProfileOutput{})
	return UpdateSecurityProfileRequest{Request: req, Input: input, Copy: c.UpdateSecurityProfileRequest}
}

// UpdateSecurityProfileRequest is the request type for the
// UpdateSecurityProfile API operation.
type UpdateSecurityProfileRequest struct {
	*aws.Request
	Input *UpdateSecurityProfileInput
	Copy  func(*UpdateSecurityProfileInput) UpdateSecurityProfileRequest
}

// Send marshals and sends the UpdateSecurityProfile API request.
func (r UpdateSecurityProfileRequest) Send(ctx context.Context) (*UpdateSecurityProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSecurityProfileResponse{
		UpdateSecurityProfileOutput: r.Request.Data.(*UpdateSecurityProfileOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSecurityProfileResponse is the response type for the
// UpdateSecurityProfile API operation.
type UpdateSecurityProfileResponse struct {
	*UpdateSecurityProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSecurityProfile request.
func (r *UpdateSecurityProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
