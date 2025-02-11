// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteEndpointRequest
type DeleteEndpointInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// EndpointId is a required field
	EndpointId *string `location:"uri" locationName:"endpoint-id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteEndpointInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.EndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteEndpointInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EndpointId != nil {
		v := *s.EndpointId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "endpoint-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteEndpointResponse
type DeleteEndpointOutput struct {
	_ struct{} `type:"structure" payload:"EndpointResponse"`

	// Provides information about the channel type and other settings for an endpoint.
	//
	// EndpointResponse is a required field
	EndpointResponse *EndpointResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteEndpointOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.EndpointResponse != nil {
		v := s.EndpointResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "EndpointResponse", v, metadata)
	}
	return nil
}

const opDeleteEndpoint = "DeleteEndpoint"

// DeleteEndpointRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Deletes an endpoint from an application.
//
//    // Example sending a request using DeleteEndpointRequest.
//    req := client.DeleteEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteEndpoint
func (c *Client) DeleteEndpointRequest(input *DeleteEndpointInput) DeleteEndpointRequest {
	op := &aws.Operation{
		Name:       opDeleteEndpoint,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/apps/{application-id}/endpoints/{endpoint-id}",
	}

	if input == nil {
		input = &DeleteEndpointInput{}
	}

	req := c.newRequest(op, input, &DeleteEndpointOutput{})
	return DeleteEndpointRequest{Request: req, Input: input, Copy: c.DeleteEndpointRequest}
}

// DeleteEndpointRequest is the request type for the
// DeleteEndpoint API operation.
type DeleteEndpointRequest struct {
	*aws.Request
	Input *DeleteEndpointInput
	Copy  func(*DeleteEndpointInput) DeleteEndpointRequest
}

// Send marshals and sends the DeleteEndpoint API request.
func (r DeleteEndpointRequest) Send(ctx context.Context) (*DeleteEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteEndpointResponse{
		DeleteEndpointOutput: r.Request.Data.(*DeleteEndpointOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteEndpointResponse is the response type for the
// DeleteEndpoint API operation.
type DeleteEndpointResponse struct {
	*DeleteEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteEndpoint request.
func (r *DeleteEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
