// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetConnectivityInfoRequest
type GetConnectivityInfoInput struct {
	_ struct{} `type:"structure"`

	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"ThingName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetConnectivityInfoInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetConnectivityInfoInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetConnectivityInfoInput"}

	if s.ThingName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetConnectivityInfoInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ThingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about a Greengrass core's connectivity.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetConnectivityInfoResponse
type GetConnectivityInfoOutput struct {
	_ struct{} `type:"structure"`

	// Connectivity info list.
	ConnectivityInfo []ConnectivityInfo `type:"list"`

	// A message about the connectivity info request.
	Message *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s GetConnectivityInfoOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetConnectivityInfoOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ConnectivityInfo != nil {
		v := s.ConnectivityInfo

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ConnectivityInfo", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Message != nil {
		v := *s.Message

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "message", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetConnectivityInfo = "GetConnectivityInfo"

// GetConnectivityInfoRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves the connectivity information for a core.
//
//    // Example sending a request using GetConnectivityInfoRequest.
//    req := client.GetConnectivityInfoRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetConnectivityInfo
func (c *Client) GetConnectivityInfoRequest(input *GetConnectivityInfoInput) GetConnectivityInfoRequest {
	op := &aws.Operation{
		Name:       opGetConnectivityInfo,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/things/{ThingName}/connectivityInfo",
	}

	if input == nil {
		input = &GetConnectivityInfoInput{}
	}

	req := c.newRequest(op, input, &GetConnectivityInfoOutput{})
	return GetConnectivityInfoRequest{Request: req, Input: input, Copy: c.GetConnectivityInfoRequest}
}

// GetConnectivityInfoRequest is the request type for the
// GetConnectivityInfo API operation.
type GetConnectivityInfoRequest struct {
	*aws.Request
	Input *GetConnectivityInfoInput
	Copy  func(*GetConnectivityInfoInput) GetConnectivityInfoRequest
}

// Send marshals and sends the GetConnectivityInfo API request.
func (r GetConnectivityInfoRequest) Send(ctx context.Context) (*GetConnectivityInfoResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetConnectivityInfoResponse{
		GetConnectivityInfoOutput: r.Request.Data.(*GetConnectivityInfoOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetConnectivityInfoResponse is the response type for the
// GetConnectivityInfo API operation.
type GetConnectivityInfoResponse struct {
	*GetConnectivityInfoOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetConnectivityInfo request.
func (r *GetConnectivityInfoResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
