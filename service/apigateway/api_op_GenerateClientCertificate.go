// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to generate a ClientCertificate resource.
type GenerateClientCertificateInput struct {
	_ struct{} `type:"structure"`

	// The description of the ClientCertificate.
	Description *string `locationName:"description" type:"string"`

	// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/].
	// The tag key can be up to 128 characters and must not start with aws:. The
	// tag value can be up to 256 characters.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s GenerateClientCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GenerateClientCertificateInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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

// Represents a client certificate used to configure client-side SSL authentication
// while sending requests to the integration endpoint.
//
// Client certificates are used to authenticate an API by the backend server.
// To authenticate an API client (or user), use IAM roles and policies, a custom
// Authorizer or an Amazon Cognito user pool.
//
// Use Client-Side Certificate (https://docs.aws.amazon.com/apigateway/latest/developerguide/getting-started-client-side-ssl-authentication.html)
type GenerateClientCertificateOutput struct {
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
func (s GenerateClientCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GenerateClientCertificateOutput) MarshalFields(e protocol.FieldEncoder) error {
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

const opGenerateClientCertificate = "GenerateClientCertificate"

// GenerateClientCertificateRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Generates a ClientCertificate resource.
//
//    // Example sending a request using GenerateClientCertificateRequest.
//    req := client.GenerateClientCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GenerateClientCertificateRequest(input *GenerateClientCertificateInput) GenerateClientCertificateRequest {
	op := &aws.Operation{
		Name:       opGenerateClientCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/clientcertificates",
	}

	if input == nil {
		input = &GenerateClientCertificateInput{}
	}

	req := c.newRequest(op, input, &GenerateClientCertificateOutput{})
	return GenerateClientCertificateRequest{Request: req, Input: input, Copy: c.GenerateClientCertificateRequest}
}

// GenerateClientCertificateRequest is the request type for the
// GenerateClientCertificate API operation.
type GenerateClientCertificateRequest struct {
	*aws.Request
	Input *GenerateClientCertificateInput
	Copy  func(*GenerateClientCertificateInput) GenerateClientCertificateRequest
}

// Send marshals and sends the GenerateClientCertificate API request.
func (r GenerateClientCertificateRequest) Send(ctx context.Context) (*GenerateClientCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GenerateClientCertificateResponse{
		GenerateClientCertificateOutput: r.Request.Data.(*GenerateClientCertificateOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GenerateClientCertificateResponse is the response type for the
// GenerateClientCertificate API operation.
type GenerateClientCertificateResponse struct {
	*GenerateClientCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GenerateClientCertificate request.
func (r *GenerateClientCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
