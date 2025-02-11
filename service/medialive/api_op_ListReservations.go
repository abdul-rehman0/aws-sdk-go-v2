// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/ListReservationsRequest
type ListReservationsInput struct {
	_ struct{} `type:"structure"`

	ChannelClass *string `location:"querystring" locationName:"channelClass" type:"string"`

	Codec *string `location:"querystring" locationName:"codec" type:"string"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	MaximumBitrate *string `location:"querystring" locationName:"maximumBitrate" type:"string"`

	MaximumFramerate *string `location:"querystring" locationName:"maximumFramerate" type:"string"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	Resolution *string `location:"querystring" locationName:"resolution" type:"string"`

	ResourceType *string `location:"querystring" locationName:"resourceType" type:"string"`

	SpecialFeature *string `location:"querystring" locationName:"specialFeature" type:"string"`

	VideoQuality *string `location:"querystring" locationName:"videoQuality" type:"string"`
}

// String returns the string representation
func (s ListReservationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListReservationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListReservationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListReservationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ChannelClass != nil {
		v := *s.ChannelClass

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "channelClass", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Codec != nil {
		v := *s.Codec

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "codec", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.MaximumBitrate != nil {
		v := *s.MaximumBitrate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maximumBitrate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaximumFramerate != nil {
		v := *s.MaximumFramerate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maximumFramerate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Resolution != nil {
		v := *s.Resolution

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "resolution", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceType != nil {
		v := *s.ResourceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "resourceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SpecialFeature != nil {
		v := *s.SpecialFeature

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "specialFeature", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VideoQuality != nil {
		v := *s.VideoQuality

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "videoQuality", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/ListReservationsResponse
type ListReservationsOutput struct {
	_ struct{} `type:"structure"`

	NextToken *string `locationName:"nextToken" type:"string"`

	Reservations []Reservation `locationName:"reservations" type:"list"`
}

// String returns the string representation
func (s ListReservationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListReservationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Reservations != nil {
		v := s.Reservations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "reservations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListReservations = "ListReservations"

// ListReservationsRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// List purchased reservations.
//
//    // Example sending a request using ListReservationsRequest.
//    req := client.ListReservationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/ListReservations
func (c *Client) ListReservationsRequest(input *ListReservationsInput) ListReservationsRequest {
	op := &aws.Operation{
		Name:       opListReservations,
		HTTPMethod: "GET",
		HTTPPath:   "/prod/reservations",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListReservationsInput{}
	}

	req := c.newRequest(op, input, &ListReservationsOutput{})
	return ListReservationsRequest{Request: req, Input: input, Copy: c.ListReservationsRequest}
}

// ListReservationsRequest is the request type for the
// ListReservations API operation.
type ListReservationsRequest struct {
	*aws.Request
	Input *ListReservationsInput
	Copy  func(*ListReservationsInput) ListReservationsRequest
}

// Send marshals and sends the ListReservations API request.
func (r ListReservationsRequest) Send(ctx context.Context) (*ListReservationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListReservationsResponse{
		ListReservationsOutput: r.Request.Data.(*ListReservationsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListReservationsRequestPaginator returns a paginator for ListReservations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListReservationsRequest(input)
//   p := medialive.NewListReservationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListReservationsPaginator(req ListReservationsRequest) ListReservationsPaginator {
	return ListReservationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListReservationsInput
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

// ListReservationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListReservationsPaginator struct {
	aws.Pager
}

func (p *ListReservationsPaginator) CurrentPage() *ListReservationsOutput {
	return p.Pager.CurrentPage().(*ListReservationsOutput)
}

// ListReservationsResponse is the response type for the
// ListReservations API operation.
type ListReservationsResponse struct {
	*ListReservationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListReservations request.
func (r *ListReservationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
