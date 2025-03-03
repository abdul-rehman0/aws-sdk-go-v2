// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/GetPublicKeyRequest
type GetPublicKeyInput struct {
	_ struct{} `type:"structure"`

	// Request the ID for the public key.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPublicKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPublicKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPublicKeyInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPublicKeyInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/GetPublicKeyResult
type GetPublicKeyOutput struct {
	_ struct{} `type:"structure" payload:"PublicKey"`

	// The current version of the public key. For example: E2QWRUHAPOMQZL.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// Return the public key.
	PublicKey *PublicKey `type:"structure"`
}

// String returns the string representation
func (s GetPublicKeyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPublicKeyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.PublicKey != nil {
		v := s.PublicKey

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "PublicKey", v, metadata)
	}
	return nil
}

const opGetPublicKey = "GetPublicKey2019_03_26"

// GetPublicKeyRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Get the public key information.
//
//    // Example sending a request using GetPublicKeyRequest.
//    req := client.GetPublicKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/GetPublicKey
func (c *Client) GetPublicKeyRequest(input *GetPublicKeyInput) GetPublicKeyRequest {
	op := &aws.Operation{
		Name:       opGetPublicKey,
		HTTPMethod: "GET",
		HTTPPath:   "/2019-03-26/public-key/{Id}",
	}

	if input == nil {
		input = &GetPublicKeyInput{}
	}

	req := c.newRequest(op, input, &GetPublicKeyOutput{})
	return GetPublicKeyRequest{Request: req, Input: input, Copy: c.GetPublicKeyRequest}
}

// GetPublicKeyRequest is the request type for the
// GetPublicKey API operation.
type GetPublicKeyRequest struct {
	*aws.Request
	Input *GetPublicKeyInput
	Copy  func(*GetPublicKeyInput) GetPublicKeyRequest
}

// Send marshals and sends the GetPublicKey API request.
func (r GetPublicKeyRequest) Send(ctx context.Context) (*GetPublicKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPublicKeyResponse{
		GetPublicKeyOutput: r.Request.Data.(*GetPublicKeyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPublicKeyResponse is the response type for the
// GetPublicKey API operation.
type GetPublicKeyResponse struct {
	*GetPublicKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPublicKey request.
func (r *GetPublicKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
