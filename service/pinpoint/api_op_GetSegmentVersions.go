// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetSegmentVersionsRequest
type GetSegmentVersionsInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	PageSize *string `location:"querystring" locationName:"page-size" type:"string"`

	// SegmentId is a required field
	SegmentId *string `location:"uri" locationName:"segment-id" type:"string" required:"true"`

	Token *string `location:"querystring" locationName:"token" type:"string"`
}

// String returns the string representation
func (s GetSegmentVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSegmentVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSegmentVersionsInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.SegmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SegmentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSegmentVersionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SegmentId != nil {
		v := *s.SegmentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "segment-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PageSize != nil {
		v := *s.PageSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "page-size", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Token != nil {
		v := *s.Token

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetSegmentVersionsResponse
type GetSegmentVersionsOutput struct {
	_ struct{} `type:"structure" payload:"SegmentsResponse"`

	// Provides information about all the segments that are associated with an application.
	//
	// SegmentsResponse is a required field
	SegmentsResponse *SegmentsResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetSegmentVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSegmentVersionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.SegmentsResponse != nil {
		v := s.SegmentsResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "SegmentsResponse", v, metadata)
	}
	return nil
}

const opGetSegmentVersions = "GetSegmentVersions"

// GetSegmentVersionsRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the configuration, dimension, and other settings
// for all versions of a specific segment that's associated with an application.
//
//    // Example sending a request using GetSegmentVersionsRequest.
//    req := client.GetSegmentVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetSegmentVersions
func (c *Client) GetSegmentVersionsRequest(input *GetSegmentVersionsInput) GetSegmentVersionsRequest {
	op := &aws.Operation{
		Name:       opGetSegmentVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/segments/{segment-id}/versions",
	}

	if input == nil {
		input = &GetSegmentVersionsInput{}
	}

	req := c.newRequest(op, input, &GetSegmentVersionsOutput{})
	return GetSegmentVersionsRequest{Request: req, Input: input, Copy: c.GetSegmentVersionsRequest}
}

// GetSegmentVersionsRequest is the request type for the
// GetSegmentVersions API operation.
type GetSegmentVersionsRequest struct {
	*aws.Request
	Input *GetSegmentVersionsInput
	Copy  func(*GetSegmentVersionsInput) GetSegmentVersionsRequest
}

// Send marshals and sends the GetSegmentVersions API request.
func (r GetSegmentVersionsRequest) Send(ctx context.Context) (*GetSegmentVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSegmentVersionsResponse{
		GetSegmentVersionsOutput: r.Request.Data.(*GetSegmentVersionsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSegmentVersionsResponse is the response type for the
// GetSegmentVersions API operation.
type GetSegmentVersionsResponse struct {
	*GetSegmentVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSegmentVersions request.
func (r *GetSegmentVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
