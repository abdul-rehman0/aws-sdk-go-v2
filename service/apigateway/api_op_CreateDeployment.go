// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Requests API Gateway to create a Deployment resource.
type CreateDeploymentInput struct {
	_ struct{} `type:"structure"`

	// Enables a cache cluster for the Stage resource specified in the input.
	CacheClusterEnabled *bool `locationName:"cacheClusterEnabled" type:"boolean"`

	// Specifies the cache cluster size for the Stage resource specified in the
	// input, if a cache cluster is enabled.
	CacheClusterSize CacheClusterSize `locationName:"cacheClusterSize" type:"string" enum:"true"`

	// The input configuration for the canary deployment when the deployment is
	// a canary release deployment.
	CanarySettings *DeploymentCanarySettings `locationName:"canarySettings" type:"structure"`

	// The description for the Deployment resource to create.
	Description *string `locationName:"description" type:"string"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`

	// The description of the Stage resource for the Deployment resource to create.
	StageDescription *string `locationName:"stageDescription" type:"string"`

	// The name of the Stage resource for the Deployment resource to create.
	StageName *string `locationName:"stageName" type:"string"`

	// Specifies whether active tracing with X-ray is enabled for the Stage.
	TracingEnabled *bool `locationName:"tracingEnabled" type:"boolean"`

	// A map that defines the stage variables for the Stage resource that is associated
	// with the new deployment. Variable names can have alphanumeric and underscore
	// characters, and the values must match [A-Za-z0-9-._~:/?#&=,]+.
	Variables map[string]string `locationName:"variables" type:"map"`
}

// String returns the string representation
func (s CreateDeploymentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDeploymentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDeploymentInput"}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDeploymentInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.CacheClusterEnabled != nil {
		v := *s.CacheClusterEnabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "cacheClusterEnabled", protocol.BoolValue(v), metadata)
	}
	if len(s.CacheClusterSize) > 0 {
		v := s.CacheClusterSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "cacheClusterSize", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.CanarySettings != nil {
		v := s.CanarySettings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "canarySettings", v, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StageDescription != nil {
		v := *s.StageDescription

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stageDescription", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StageName != nil {
		v := *s.StageName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stageName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TracingEnabled != nil {
		v := *s.TracingEnabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "tracingEnabled", protocol.BoolValue(v), metadata)
	}
	if s.Variables != nil {
		v := s.Variables

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "variables", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An immutable representation of a RestApi resource that can be called by users
// using Stages. A deployment must be associated with a Stage for it to be callable
// over the Internet.
//
// To create a deployment, call POST on the Deployments resource of a RestApi.
// To view, update, or delete a deployment, call GET, PATCH, or DELETE on the
// specified deployment resource (/restapis/{restapi_id}/deployments/{deployment_id}).
//
// RestApi, Deployments, Stage, AWS CLI (https://docs.aws.amazon.com/cli/latest/reference/apigateway/get-deployment.html),
// AWS SDKs (https://aws.amazon.com/tools/)
type CreateDeploymentOutput struct {
	_ struct{} `type:"structure"`

	// A summary of the RestApi at the date and time that the deployment resource
	// was created.
	ApiSummary map[string]map[string]MethodSnapshot `locationName:"apiSummary" type:"map"`

	// The date and time that the deployment resource was created.
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp" timestampFormat:"unix"`

	// The description for the deployment resource.
	Description *string `locationName:"description" type:"string"`

	// The identifier for the deployment resource.
	Id *string `locationName:"id" type:"string"`
}

// String returns the string representation
func (s CreateDeploymentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDeploymentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApiSummary != nil {
		v := s.ApiSummary

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "apiSummary", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms1 := ms0.Map(k1)
			ms1.Start()
			for k2, v2 := range v1 {
				ms1.MapSetFields(k2, v2)
			}
			ms1.End()
		}
		ms0.End()

	}
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateDeployment = "CreateDeployment"

// CreateDeploymentRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Creates a Deployment resource, which makes a specified RestApi callable over
// the internet.
//
//    // Example sending a request using CreateDeploymentRequest.
//    req := client.CreateDeploymentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateDeploymentRequest(input *CreateDeploymentInput) CreateDeploymentRequest {
	op := &aws.Operation{
		Name:       opCreateDeployment,
		HTTPMethod: "POST",
		HTTPPath:   "/restapis/{restapi_id}/deployments",
	}

	if input == nil {
		input = &CreateDeploymentInput{}
	}

	req := c.newRequest(op, input, &CreateDeploymentOutput{})
	return CreateDeploymentRequest{Request: req, Input: input, Copy: c.CreateDeploymentRequest}
}

// CreateDeploymentRequest is the request type for the
// CreateDeployment API operation.
type CreateDeploymentRequest struct {
	*aws.Request
	Input *CreateDeploymentInput
	Copy  func(*CreateDeploymentInput) CreateDeploymentRequest
}

// Send marshals and sends the CreateDeployment API request.
func (r CreateDeploymentRequest) Send(ctx context.Context) (*CreateDeploymentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDeploymentResponse{
		CreateDeploymentOutput: r.Request.Data.(*CreateDeploymentOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDeploymentResponse is the response type for the
// CreateDeployment API operation.
type CreateDeploymentResponse struct {
	*CreateDeploymentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDeployment request.
func (r *CreateDeploymentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
