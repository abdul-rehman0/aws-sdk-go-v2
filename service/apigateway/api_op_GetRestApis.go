// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The GET request to list existing RestApis defined for your collection.
type GetRestApisInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of returned results per page. The default value is 25
	// and the maximum value is 500.
	Limit *int64 `location:"querystring" locationName:"limit" type:"integer"`

	// The current pagination position in the paged result set.
	Position *string `location:"querystring" locationName:"position" type:"string"`
}

// String returns the string representation
func (s GetRestApisInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetRestApisInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.Position != nil {
		v := *s.Position

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "position", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Contains references to your APIs and links that guide you in how to interact
// with your collection. A collection offers a paginated view of your APIs.
//
// Create an API (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
type GetRestApisOutput struct {
	_ struct{} `type:"structure"`

	// The current page of elements from this collection.
	Items []RestApi `locationName:"item" type:"list"`

	Position *string `locationName:"position" type:"string"`
}

// String returns the string representation
func (s GetRestApisOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetRestApisOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Items != nil {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "item", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Position != nil {
		v := *s.Position

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "position", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetRestApis = "GetRestApis"

// GetRestApisRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Lists the RestApis resources for your collection.
//
//    // Example sending a request using GetRestApisRequest.
//    req := client.GetRestApisRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetRestApisRequest(input *GetRestApisInput) GetRestApisRequest {
	op := &aws.Operation{
		Name:       opGetRestApis,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"position"},
			OutputTokens:    []string{"position"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetRestApisInput{}
	}

	req := c.newRequest(op, input, &GetRestApisOutput{})
	return GetRestApisRequest{Request: req, Input: input, Copy: c.GetRestApisRequest}
}

// GetRestApisRequest is the request type for the
// GetRestApis API operation.
type GetRestApisRequest struct {
	*aws.Request
	Input *GetRestApisInput
	Copy  func(*GetRestApisInput) GetRestApisRequest
}

// Send marshals and sends the GetRestApis API request.
func (r GetRestApisRequest) Send(ctx context.Context) (*GetRestApisResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRestApisResponse{
		GetRestApisOutput: r.Request.Data.(*GetRestApisOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetRestApisRequestPaginator returns a paginator for GetRestApis.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetRestApisRequest(input)
//   p := apigateway.NewGetRestApisRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetRestApisPaginator(req GetRestApisRequest) GetRestApisPaginator {
	return GetRestApisPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetRestApisInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// GetRestApisPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetRestApisPaginator struct {
	aws.Pager
}

func (p *GetRestApisPaginator) CurrentPage() *GetRestApisOutput {
	return p.Pager.CurrentPage().(*GetRestApisOutput)
}

// GetRestApisResponse is the response type for the
// GetRestApis API operation.
type GetRestApisResponse struct {
	*GetRestApisOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRestApis request.
func (r *GetRestApisResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
