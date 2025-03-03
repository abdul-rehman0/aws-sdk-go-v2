// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateOTAUpdateInput struct {
	_ struct{} `type:"structure"`

	// A list of additional OTA update parameters which are name-value pairs.
	AdditionalParameters map[string]string `locationName:"additionalParameters" type:"map"`

	// Configuration for the rollout of OTA updates.
	AwsJobExecutionsRolloutConfig *AwsJobExecutionsRolloutConfig `locationName:"awsJobExecutionsRolloutConfig" type:"structure"`

	// The description of the OTA update.
	Description *string `locationName:"description" type:"string"`

	// The files to be streamed by the OTA update.
	//
	// Files is a required field
	Files []OTAUpdateFile `locationName:"files" min:"1" type:"list" required:"true"`

	// The ID of the OTA update to be created.
	//
	// OtaUpdateId is a required field
	OtaUpdateId *string `location:"uri" locationName:"otaUpdateId" min:"1" type:"string" required:"true"`

	// The IAM role that allows access to the AWS IoT Jobs service.
	//
	// RoleArn is a required field
	RoleArn *string `locationName:"roleArn" min:"20" type:"string" required:"true"`

	// Metadata which can be used to manage updates.
	Tags []Tag `locationName:"tags" type:"list"`

	// Specifies whether the update will continue to run (CONTINUOUS), or will be
	// complete after all the things specified as targets have completed the update
	// (SNAPSHOT). If continuous, the update may also be run on a thing when a change
	// is detected in a target. For example, an update will run on a thing when
	// the thing is added to a target group, even after the update was completed
	// by all things originally in the group. Valid values: CONTINUOUS | SNAPSHOT.
	TargetSelection TargetSelection `locationName:"targetSelection" type:"string" enum:"true"`

	// The targeted devices to receive OTA updates.
	//
	// Targets is a required field
	Targets []string `locationName:"targets" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s CreateOTAUpdateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateOTAUpdateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateOTAUpdateInput"}

	if s.Files == nil {
		invalidParams.Add(aws.NewErrParamRequired("Files"))
	}
	if s.Files != nil && len(s.Files) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Files", 1))
	}

	if s.OtaUpdateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OtaUpdateId"))
	}
	if s.OtaUpdateId != nil && len(*s.OtaUpdateId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OtaUpdateId", 1))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 20))
	}

	if s.Targets == nil {
		invalidParams.Add(aws.NewErrParamRequired("Targets"))
	}
	if s.Targets != nil && len(s.Targets) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Targets", 1))
	}
	if s.AwsJobExecutionsRolloutConfig != nil {
		if err := s.AwsJobExecutionsRolloutConfig.Validate(); err != nil {
			invalidParams.AddNested("AwsJobExecutionsRolloutConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.Files != nil {
		for i, v := range s.Files {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Files", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateOTAUpdateInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AdditionalParameters != nil {
		v := s.AdditionalParameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "additionalParameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.AwsJobExecutionsRolloutConfig != nil {
		v := s.AwsJobExecutionsRolloutConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "awsJobExecutionsRolloutConfig", v, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Files != nil {
		v := s.Files

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "files", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.RoleArn != nil {
		v := *s.RoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "roleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.TargetSelection) > 0 {
		v := s.TargetSelection

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "targetSelection", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Targets != nil {
		v := s.Targets

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "targets", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.OtaUpdateId != nil {
		v := *s.OtaUpdateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "otaUpdateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateOTAUpdateOutput struct {
	_ struct{} `type:"structure"`

	// The AWS IoT job ARN associated with the OTA update.
	AwsIotJobArn *string `locationName:"awsIotJobArn" type:"string"`

	// The AWS IoT job ID associated with the OTA update.
	AwsIotJobId *string `locationName:"awsIotJobId" type:"string"`

	// The OTA update ARN.
	OtaUpdateArn *string `locationName:"otaUpdateArn" type:"string"`

	// The OTA update ID.
	OtaUpdateId *string `locationName:"otaUpdateId" min:"1" type:"string"`

	// The OTA update status.
	OtaUpdateStatus OTAUpdateStatus `locationName:"otaUpdateStatus" type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateOTAUpdateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateOTAUpdateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AwsIotJobArn != nil {
		v := *s.AwsIotJobArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "awsIotJobArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AwsIotJobId != nil {
		v := *s.AwsIotJobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "awsIotJobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OtaUpdateArn != nil {
		v := *s.OtaUpdateArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "otaUpdateArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OtaUpdateId != nil {
		v := *s.OtaUpdateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "otaUpdateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.OtaUpdateStatus) > 0 {
		v := s.OtaUpdateStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "otaUpdateStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opCreateOTAUpdate = "CreateOTAUpdate"

// CreateOTAUpdateRequest returns a request value for making API operation for
// AWS IoT.
//
// Creates an AWS IoT OTAUpdate on a target group of things or groups.
//
//    // Example sending a request using CreateOTAUpdateRequest.
//    req := client.CreateOTAUpdateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateOTAUpdateRequest(input *CreateOTAUpdateInput) CreateOTAUpdateRequest {
	op := &aws.Operation{
		Name:       opCreateOTAUpdate,
		HTTPMethod: "POST",
		HTTPPath:   "/otaUpdates/{otaUpdateId}",
	}

	if input == nil {
		input = &CreateOTAUpdateInput{}
	}

	req := c.newRequest(op, input, &CreateOTAUpdateOutput{})
	return CreateOTAUpdateRequest{Request: req, Input: input, Copy: c.CreateOTAUpdateRequest}
}

// CreateOTAUpdateRequest is the request type for the
// CreateOTAUpdate API operation.
type CreateOTAUpdateRequest struct {
	*aws.Request
	Input *CreateOTAUpdateInput
	Copy  func(*CreateOTAUpdateInput) CreateOTAUpdateRequest
}

// Send marshals and sends the CreateOTAUpdate API request.
func (r CreateOTAUpdateRequest) Send(ctx context.Context) (*CreateOTAUpdateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateOTAUpdateResponse{
		CreateOTAUpdateOutput: r.Request.Data.(*CreateOTAUpdateOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateOTAUpdateResponse is the response type for the
// CreateOTAUpdate API operation.
type CreateOTAUpdateResponse struct {
	*CreateOTAUpdateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateOTAUpdate request.
func (r *CreateOTAUpdateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
