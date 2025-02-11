// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetDomainNamesRequest
type GetDomainNamesInput struct {
	_ struct{} `type:"structure"`

	MaxResults *string `location:"querystring" locationName:"maxResults" type:"string"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetDomainNamesInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDomainNamesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetDomainNamesResponse
type GetDomainNamesOutput struct {
	_ struct{} `type:"structure"`

	Items []DomainName `locationName:"items" type:"list"`

	// The next page of elements from this collection. Not valid for the last element
	// of the collection.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetDomainNamesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDomainNamesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Items != nil {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "items", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetDomainNames = "GetDomainNames"

// GetDomainNamesRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Gets the domain names for an AWS account.
//
//    // Example sending a request using GetDomainNamesRequest.
//    req := client.GetDomainNamesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetDomainNames
func (c *Client) GetDomainNamesRequest(input *GetDomainNamesInput) GetDomainNamesRequest {
	op := &aws.Operation{
		Name:       opGetDomainNames,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/domainnames",
	}

	if input == nil {
		input = &GetDomainNamesInput{}
	}

	req := c.newRequest(op, input, &GetDomainNamesOutput{})
	return GetDomainNamesRequest{Request: req, Input: input, Copy: c.GetDomainNamesRequest}
}

// GetDomainNamesRequest is the request type for the
// GetDomainNames API operation.
type GetDomainNamesRequest struct {
	*aws.Request
	Input *GetDomainNamesInput
	Copy  func(*GetDomainNamesInput) GetDomainNamesRequest
}

// Send marshals and sends the GetDomainNames API request.
func (r GetDomainNamesRequest) Send(ctx context.Context) (*GetDomainNamesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDomainNamesResponse{
		GetDomainNamesOutput: r.Request.Data.(*GetDomainNamesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDomainNamesResponse is the response type for the
// GetDomainNames API operation.
type GetDomainNamesResponse struct {
	*GetDomainNamesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDomainNames request.
func (r *GetDomainNamesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
