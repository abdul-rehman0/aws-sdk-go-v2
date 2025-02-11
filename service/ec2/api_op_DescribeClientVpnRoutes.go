// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeClientVpnRoutesRequest
type DescribeClientVpnRoutesInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Client VPN endpoint.
	//
	// ClientVpnEndpointId is a required field
	ClientVpnEndpointId *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters. Filter names and values are case-sensitive.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return for the request in a single page.
	// The remaining results can be seen by sending another request with the nextToken
	// value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token to retrieve the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeClientVpnRoutesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeClientVpnRoutesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeClientVpnRoutesInput"}

	if s.ClientVpnEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientVpnEndpointId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeClientVpnRoutesResult
type DescribeClientVpnRoutesOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the Client VPN endpoint routes.
	Routes []VpnRoute `locationName:"routes" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeClientVpnRoutesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeClientVpnRoutes = "DescribeClientVpnRoutes"

// DescribeClientVpnRoutesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the routes for the specified Client VPN endpoint.
//
//    // Example sending a request using DescribeClientVpnRoutesRequest.
//    req := client.DescribeClientVpnRoutesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeClientVpnRoutes
func (c *Client) DescribeClientVpnRoutesRequest(input *DescribeClientVpnRoutesInput) DescribeClientVpnRoutesRequest {
	op := &aws.Operation{
		Name:       opDescribeClientVpnRoutes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeClientVpnRoutesInput{}
	}

	req := c.newRequest(op, input, &DescribeClientVpnRoutesOutput{})
	return DescribeClientVpnRoutesRequest{Request: req, Input: input, Copy: c.DescribeClientVpnRoutesRequest}
}

// DescribeClientVpnRoutesRequest is the request type for the
// DescribeClientVpnRoutes API operation.
type DescribeClientVpnRoutesRequest struct {
	*aws.Request
	Input *DescribeClientVpnRoutesInput
	Copy  func(*DescribeClientVpnRoutesInput) DescribeClientVpnRoutesRequest
}

// Send marshals and sends the DescribeClientVpnRoutes API request.
func (r DescribeClientVpnRoutesRequest) Send(ctx context.Context) (*DescribeClientVpnRoutesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeClientVpnRoutesResponse{
		DescribeClientVpnRoutesOutput: r.Request.Data.(*DescribeClientVpnRoutesOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeClientVpnRoutesRequestPaginator returns a paginator for DescribeClientVpnRoutes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeClientVpnRoutesRequest(input)
//   p := ec2.NewDescribeClientVpnRoutesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeClientVpnRoutesPaginator(req DescribeClientVpnRoutesRequest) DescribeClientVpnRoutesPaginator {
	return DescribeClientVpnRoutesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeClientVpnRoutesInput
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

// DescribeClientVpnRoutesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeClientVpnRoutesPaginator struct {
	aws.Pager
}

func (p *DescribeClientVpnRoutesPaginator) CurrentPage() *DescribeClientVpnRoutesOutput {
	return p.Pager.CurrentPage().(*DescribeClientVpnRoutesOutput)
}

// DescribeClientVpnRoutesResponse is the response type for the
// DescribeClientVpnRoutes API operation.
type DescribeClientVpnRoutesResponse struct {
	*DescribeClientVpnRoutesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeClientVpnRoutes request.
func (r *DescribeClientVpnRoutesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
