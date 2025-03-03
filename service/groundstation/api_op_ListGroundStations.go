// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package groundstation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/ListGroundStationsRequest
type ListGroundStationsInput struct {
	_ struct{} `type:"structure"`

	// Maximum number of ground stations returned.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	// Next token that can be supplied in the next call to get the next page of
	// ground stations.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListGroundStationsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListGroundStationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/ListGroundStationsResponse
type ListGroundStationsOutput struct {
	_ struct{} `type:"structure"`

	// List of ground stations.
	GroundStationList []GroundStationData `locationName:"groundStationList" type:"list"`

	// Next token that can be supplied in the next call to get the next page of
	// ground stations.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListGroundStationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListGroundStationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.GroundStationList != nil {
		v := s.GroundStationList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "groundStationList", metadata)
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

const opListGroundStations = "ListGroundStations"

// ListGroundStationsRequest returns a request value for making API operation for
// AWS Ground Station.
//
// Returns a list of ground stations.
//
//    // Example sending a request using ListGroundStationsRequest.
//    req := client.ListGroundStationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/ListGroundStations
func (c *Client) ListGroundStationsRequest(input *ListGroundStationsInput) ListGroundStationsRequest {
	op := &aws.Operation{
		Name:       opListGroundStations,
		HTTPMethod: "GET",
		HTTPPath:   "/groundstation",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListGroundStationsInput{}
	}

	req := c.newRequest(op, input, &ListGroundStationsOutput{})
	return ListGroundStationsRequest{Request: req, Input: input, Copy: c.ListGroundStationsRequest}
}

// ListGroundStationsRequest is the request type for the
// ListGroundStations API operation.
type ListGroundStationsRequest struct {
	*aws.Request
	Input *ListGroundStationsInput
	Copy  func(*ListGroundStationsInput) ListGroundStationsRequest
}

// Send marshals and sends the ListGroundStations API request.
func (r ListGroundStationsRequest) Send(ctx context.Context) (*ListGroundStationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGroundStationsResponse{
		ListGroundStationsOutput: r.Request.Data.(*ListGroundStationsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListGroundStationsRequestPaginator returns a paginator for ListGroundStations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListGroundStationsRequest(input)
//   p := groundstation.NewListGroundStationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListGroundStationsPaginator(req ListGroundStationsRequest) ListGroundStationsPaginator {
	return ListGroundStationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListGroundStationsInput
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

// ListGroundStationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListGroundStationsPaginator struct {
	aws.Pager
}

func (p *ListGroundStationsPaginator) CurrentPage() *ListGroundStationsOutput {
	return p.Pager.CurrentPage().(*ListGroundStationsOutput)
}

// ListGroundStationsResponse is the response type for the
// ListGroundStations API operation.
type ListGroundStationsResponse struct {
	*ListGroundStationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGroundStations request.
func (r *ListGroundStationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
