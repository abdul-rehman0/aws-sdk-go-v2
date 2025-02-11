// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to get information about the current ApiKey resource.
type GetApiKeyInput struct {
	_ struct{} `type:"structure"`

	// [Required] The identifier of the ApiKey resource.
	//
	// ApiKey is a required field
	ApiKey *string `location:"uri" locationName:"api_Key" type:"string" required:"true"`

	// A boolean flag to specify whether (true) or not (false) the result contains
	// the key value.
	IncludeValue *bool `location:"querystring" locationName:"includeValue" type:"boolean"`
}

// String returns the string representation
func (s GetApiKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetApiKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetApiKeyInput"}

	if s.ApiKey == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiKey"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetApiKeyInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ApiKey != nil {
		v := *s.ApiKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "api_Key", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IncludeValue != nil {
		v := *s.IncludeValue

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "includeValue", protocol.BoolValue(v), metadata)
	}
	return nil
}

// A resource that can be distributed to callers for executing Method resources
// that require an API key. API keys can be mapped to any Stage on any RestApi,
// which indicates that the callers with the API key can make requests to that
// stage.
//
// Use API Keys (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-api-keys.html)
type GetApiKeyOutput struct {
	_ struct{} `type:"structure"`

	// The timestamp when the API Key was created.
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp" timestampFormat:"unix"`

	// An AWS Marketplace customer identifier , when integrating with the AWS SaaS
	// Marketplace.
	CustomerId *string `locationName:"customerId" type:"string"`

	// The description of the API Key.
	Description *string `locationName:"description" type:"string"`

	// Specifies whether the API Key can be used by callers.
	Enabled *bool `locationName:"enabled" type:"boolean"`

	// The identifier of the API Key.
	Id *string `locationName:"id" type:"string"`

	// The timestamp when the API Key was last updated.
	LastUpdatedDate *time.Time `locationName:"lastUpdatedDate" type:"timestamp" timestampFormat:"unix"`

	// The name of the API Key.
	Name *string `locationName:"name" type:"string"`

	// A list of Stage resources that are associated with the ApiKey resource.
	StageKeys []string `locationName:"stageKeys" type:"list"`

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string `locationName:"tags" type:"map"`

	// The value of the API Key.
	Value *string `locationName:"value" type:"string"`
}

// String returns the string representation
func (s GetApiKeyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetApiKeyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.CustomerId != nil {
		v := *s.CustomerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "customerId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Enabled != nil {
		v := *s.Enabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enabled", protocol.BoolValue(v), metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastUpdatedDate != nil {
		v := *s.LastUpdatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StageKeys != nil {
		v := s.StageKeys

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "stageKeys", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Value != nil {
		v := *s.Value

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "value", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetApiKey = "GetApiKey"

// GetApiKeyRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Gets information about the current ApiKey resource.
//
//    // Example sending a request using GetApiKeyRequest.
//    req := client.GetApiKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetApiKeyRequest(input *GetApiKeyInput) GetApiKeyRequest {
	op := &aws.Operation{
		Name:       opGetApiKey,
		HTTPMethod: "GET",
		HTTPPath:   "/apikeys/{api_Key}",
	}

	if input == nil {
		input = &GetApiKeyInput{}
	}

	req := c.newRequest(op, input, &GetApiKeyOutput{})
	return GetApiKeyRequest{Request: req, Input: input, Copy: c.GetApiKeyRequest}
}

// GetApiKeyRequest is the request type for the
// GetApiKey API operation.
type GetApiKeyRequest struct {
	*aws.Request
	Input *GetApiKeyInput
	Copy  func(*GetApiKeyInput) GetApiKeyRequest
}

// Send marshals and sends the GetApiKey API request.
func (r GetApiKeyRequest) Send(ctx context.Context) (*GetApiKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetApiKeyResponse{
		GetApiKeyOutput: r.Request.Data.(*GetApiKeyOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetApiKeyResponse is the response type for the
// GetApiKey API operation.
type GetApiKeyResponse struct {
	*GetApiKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetApiKey request.
func (r *GetApiKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
