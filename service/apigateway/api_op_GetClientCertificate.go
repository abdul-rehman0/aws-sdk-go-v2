// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to get information about the current ClientCertificate resource.
type GetClientCertificateInput struct {
	_ struct{} `type:"structure"`

	// [Required] The identifier of the ClientCertificate resource to be described.
	//
	// ClientCertificateId is a required field
	ClientCertificateId *string `location:"uri" locationName:"clientcertificate_id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetClientCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetClientCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetClientCertificateInput"}

	if s.ClientCertificateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientCertificateId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetClientCertificateInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ClientCertificateId != nil {
		v := *s.ClientCertificateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "clientcertificate_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents a client certificate used to configure client-side SSL authentication
// while sending requests to the integration endpoint.
//
// Client certificates are used to authenticate an API by the backend server.
// To authenticate an API client (or user), use IAM roles and policies, a custom
// Authorizer or an Amazon Cognito user pool.
//
// Use Client-Side Certificate (https://docs.aws.amazon.com/apigateway/latest/developerguide/getting-started-client-side-ssl-authentication.html)
type GetClientCertificateOutput struct {
	_ struct{} `type:"structure"`

	// The identifier of the client certificate.
	ClientCertificateId *string `locationName:"clientCertificateId" type:"string"`

	// The timestamp when the client certificate was created.
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp" timestampFormat:"unix"`

	// The description of the client certificate.
	Description *string `locationName:"description" type:"string"`

	// The timestamp when the client certificate will expire.
	ExpirationDate *time.Time `locationName:"expirationDate" type:"timestamp" timestampFormat:"unix"`

	// The PEM-encoded public key of the client certificate, which can be used to
	// configure certificate authentication in the integration endpoint .
	PemEncodedCertificate *string `locationName:"pemEncodedCertificate" type:"string"`

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s GetClientCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetClientCertificateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClientCertificateId != nil {
		v := *s.ClientCertificateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientCertificateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.ExpirationDate != nil {
		v := *s.ExpirationDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "expirationDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.PemEncodedCertificate != nil {
		v := *s.PemEncodedCertificate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "pemEncodedCertificate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	return nil
}

const opGetClientCertificate = "GetClientCertificate"

// GetClientCertificateRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Gets information about the current ClientCertificate resource.
//
//    // Example sending a request using GetClientCertificateRequest.
//    req := client.GetClientCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetClientCertificateRequest(input *GetClientCertificateInput) GetClientCertificateRequest {
	op := &aws.Operation{
		Name:       opGetClientCertificate,
		HTTPMethod: "GET",
		HTTPPath:   "/clientcertificates/{clientcertificate_id}",
	}

	if input == nil {
		input = &GetClientCertificateInput{}
	}

	req := c.newRequest(op, input, &GetClientCertificateOutput{})
	return GetClientCertificateRequest{Request: req, Input: input, Copy: c.GetClientCertificateRequest}
}

// GetClientCertificateRequest is the request type for the
// GetClientCertificate API operation.
type GetClientCertificateRequest struct {
	*aws.Request
	Input *GetClientCertificateInput
	Copy  func(*GetClientCertificateInput) GetClientCertificateRequest
}

// Send marshals and sends the GetClientCertificate API request.
func (r GetClientCertificateRequest) Send(ctx context.Context) (*GetClientCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetClientCertificateResponse{
		GetClientCertificateOutput: r.Request.Data.(*GetClientCertificateOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetClientCertificateResponse is the response type for the
// GetClientCertificate API operation.
type GetClientCertificateResponse struct {
	*GetClientCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetClientCertificate request.
func (r *GetClientCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
