// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetCampaignsRequest
type GetCampaignsInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	PageSize *string `location:"querystring" locationName:"page-size" type:"string"`

	Token *string `location:"querystring" locationName:"token" type:"string"`
}

// String returns the string representation
func (s GetCampaignsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCampaignsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCampaignsInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCampaignsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetCampaignsResponse
type GetCampaignsOutput struct {
	_ struct{} `type:"structure" payload:"CampaignsResponse"`

	// Provides information about the configuration and other settings for all the
	// campaigns that are associated with an application.
	//
	// CampaignsResponse is a required field
	CampaignsResponse *CampaignsResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetCampaignsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCampaignsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CampaignsResponse != nil {
		v := s.CampaignsResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "CampaignsResponse", v, metadata)
	}
	return nil
}

const opGetCampaigns = "GetCampaigns"

// GetCampaignsRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the status, configuration, and other settings
// for all the campaigns that are associated with an application.
//
//    // Example sending a request using GetCampaignsRequest.
//    req := client.GetCampaignsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetCampaigns
func (c *Client) GetCampaignsRequest(input *GetCampaignsInput) GetCampaignsRequest {
	op := &aws.Operation{
		Name:       opGetCampaigns,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/campaigns",
	}

	if input == nil {
		input = &GetCampaignsInput{}
	}

	req := c.newRequest(op, input, &GetCampaignsOutput{})
	return GetCampaignsRequest{Request: req, Input: input, Copy: c.GetCampaignsRequest}
}

// GetCampaignsRequest is the request type for the
// GetCampaigns API operation.
type GetCampaignsRequest struct {
	*aws.Request
	Input *GetCampaignsInput
	Copy  func(*GetCampaignsInput) GetCampaignsRequest
}

// Send marshals and sends the GetCampaigns API request.
func (r GetCampaignsRequest) Send(ctx context.Context) (*GetCampaignsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCampaignsResponse{
		GetCampaignsOutput: r.Request.Data.(*GetCampaignsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCampaignsResponse is the response type for the
// GetCampaigns API operation.
type GetCampaignsResponse struct {
	*GetCampaignsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCampaigns request.
func (r *GetCampaignsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
