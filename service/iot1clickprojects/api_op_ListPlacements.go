// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot1clickprojects

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/ListPlacementsRequest
type ListPlacementsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return per request. If not set, a default
	// value of 100 is used.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" min:"1" type:"string"`

	// The project containing the placements to be listed.
	//
	// ProjectName is a required field
	ProjectName *string `location:"uri" locationName:"projectName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListPlacementsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPlacementsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPlacementsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.ProjectName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectName"))
	}
	if s.ProjectName != nil && len(*s.ProjectName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPlacementsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ProjectName != nil {
		v := *s.ProjectName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "projectName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/ListPlacementsResponse
type ListPlacementsOutput struct {
	_ struct{} `type:"structure"`

	// The token used to retrieve the next set of results - will be effectively
	// empty if there are no further results.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// An object listing the requested placements.
	//
	// Placements is a required field
	Placements []PlacementSummary `locationName:"placements" type:"list" required:"true"`
}

// String returns the string representation
func (s ListPlacementsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPlacementsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Placements != nil {
		v := s.Placements

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "placements", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListPlacements = "ListPlacements"

// ListPlacementsRequest returns a request value for making API operation for
// AWS IoT 1-Click Projects Service.
//
// Lists the placement(s) of a project.
//
//    // Example sending a request using ListPlacementsRequest.
//    req := client.ListPlacementsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/ListPlacements
func (c *Client) ListPlacementsRequest(input *ListPlacementsInput) ListPlacementsRequest {
	op := &aws.Operation{
		Name:       opListPlacements,
		HTTPMethod: "GET",
		HTTPPath:   "/projects/{projectName}/placements",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListPlacementsInput{}
	}

	req := c.newRequest(op, input, &ListPlacementsOutput{})
	return ListPlacementsRequest{Request: req, Input: input, Copy: c.ListPlacementsRequest}
}

// ListPlacementsRequest is the request type for the
// ListPlacements API operation.
type ListPlacementsRequest struct {
	*aws.Request
	Input *ListPlacementsInput
	Copy  func(*ListPlacementsInput) ListPlacementsRequest
}

// Send marshals and sends the ListPlacements API request.
func (r ListPlacementsRequest) Send(ctx context.Context) (*ListPlacementsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPlacementsResponse{
		ListPlacementsOutput: r.Request.Data.(*ListPlacementsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPlacementsRequestPaginator returns a paginator for ListPlacements.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPlacementsRequest(input)
//   p := iot1clickprojects.NewListPlacementsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPlacementsPaginator(req ListPlacementsRequest) ListPlacementsPaginator {
	return ListPlacementsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListPlacementsInput
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

// ListPlacementsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPlacementsPaginator struct {
	aws.Pager
}

func (p *ListPlacementsPaginator) CurrentPage() *ListPlacementsOutput {
	return p.Pager.CurrentPage().(*ListPlacementsOutput)
}

// ListPlacementsResponse is the response type for the
// ListPlacements API operation.
type ListPlacementsResponse struct {
	*ListPlacementsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPlacements request.
func (r *ListPlacementsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
