// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package xray

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/xray-2016-04-12/GetTraceSummariesRequest
type GetTraceSummariesInput struct {
	_ struct{} `type:"structure"`

	// The end of the time frame for which to retrieve traces.
	//
	// EndTime is a required field
	EndTime *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`

	// Specify a filter expression to retrieve trace summaries for services or requests
	// that meet certain requirements.
	FilterExpression *string `type:"string"`

	// Specify the pagination token returned by a previous request to retrieve the
	// next page of results.
	NextToken *string `type:"string"`

	// Set to true to get summaries for only a subset of available traces.
	Sampling *bool `type:"boolean"`

	// A paramater to indicate whether to enable sampling on trace summaries. Input
	// parameters are Name and Value.
	SamplingStrategy *SamplingStrategy `type:"structure"`

	// The start of the time frame for which to retrieve traces.
	//
	// StartTime is a required field
	StartTime *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`

	// A parameter to indicate whether to query trace summaries by TraceId or Event
	// time.
	TimeRangeType TimeRangeType `type:"string" enum:"true"`
}

// String returns the string representation
func (s GetTraceSummariesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTraceSummariesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTraceSummariesInput"}

	if s.EndTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndTime"))
	}

	if s.StartTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartTime"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetTraceSummariesInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.EndTime != nil {
		v := *s.EndTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EndTime", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.FilterExpression != nil {
		v := *s.FilterExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FilterExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Sampling != nil {
		v := *s.Sampling

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Sampling", protocol.BoolValue(v), metadata)
	}
	if s.SamplingStrategy != nil {
		v := s.SamplingStrategy

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SamplingStrategy", v, metadata)
	}
	if s.StartTime != nil {
		v := *s.StartTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StartTime", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if len(s.TimeRangeType) > 0 {
		v := s.TimeRangeType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TimeRangeType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/xray-2016-04-12/GetTraceSummariesResult
type GetTraceSummariesOutput struct {
	_ struct{} `type:"structure"`

	// The start time of this page of results.
	ApproximateTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// If the requested time frame contained more than one page of results, you
	// can use this token to retrieve the next page. The first page contains the
	// most most recent results, closest to the end of the time frame.
	NextToken *string `type:"string"`

	// Trace IDs and metadata for traces that were found in the specified time frame.
	TraceSummaries []TraceSummary `type:"list"`

	// The total number of traces processed, including traces that did not match
	// the specified filter expression.
	TracesProcessedCount *int64 `type:"long"`
}

// String returns the string representation
func (s GetTraceSummariesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetTraceSummariesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApproximateTime != nil {
		v := *s.ApproximateTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ApproximateTime", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TraceSummaries != nil {
		v := s.TraceSummaries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "TraceSummaries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.TracesProcessedCount != nil {
		v := *s.TracesProcessedCount

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TracesProcessedCount", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opGetTraceSummaries = "GetTraceSummaries"

// GetTraceSummariesRequest returns a request value for making API operation for
// AWS X-Ray.
//
// Retrieves IDs and metadata for traces available for a specified time frame
// using an optional filter. To get the full traces, pass the trace IDs to BatchGetTraces.
//
// A filter expression can target traced requests that hit specific service
// nodes or edges, have errors, or come from a known user. For example, the
// following filter expression targets traces that pass through api.example.com:
//
// service("api.example.com")
//
// This filter expression finds traces that have an annotation named account
// with the value 12345:
//
// annotation.account = "12345"
//
// For a full list of indexed fields and keywords that you can use in filter
// expressions, see Using Filter Expressions (https://docs.aws.amazon.com/xray/latest/devguide/xray-console-filters.html)
// in the AWS X-Ray Developer Guide.
//
//    // Example sending a request using GetTraceSummariesRequest.
//    req := client.GetTraceSummariesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/xray-2016-04-12/GetTraceSummaries
func (c *Client) GetTraceSummariesRequest(input *GetTraceSummariesInput) GetTraceSummariesRequest {
	op := &aws.Operation{
		Name:       opGetTraceSummaries,
		HTTPMethod: "POST",
		HTTPPath:   "/TraceSummaries",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetTraceSummariesInput{}
	}

	req := c.newRequest(op, input, &GetTraceSummariesOutput{})
	return GetTraceSummariesRequest{Request: req, Input: input, Copy: c.GetTraceSummariesRequest}
}

// GetTraceSummariesRequest is the request type for the
// GetTraceSummaries API operation.
type GetTraceSummariesRequest struct {
	*aws.Request
	Input *GetTraceSummariesInput
	Copy  func(*GetTraceSummariesInput) GetTraceSummariesRequest
}

// Send marshals and sends the GetTraceSummaries API request.
func (r GetTraceSummariesRequest) Send(ctx context.Context) (*GetTraceSummariesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTraceSummariesResponse{
		GetTraceSummariesOutput: r.Request.Data.(*GetTraceSummariesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetTraceSummariesRequestPaginator returns a paginator for GetTraceSummaries.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetTraceSummariesRequest(input)
//   p := xray.NewGetTraceSummariesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetTraceSummariesPaginator(req GetTraceSummariesRequest) GetTraceSummariesPaginator {
	return GetTraceSummariesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetTraceSummariesInput
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

// GetTraceSummariesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetTraceSummariesPaginator struct {
	aws.Pager
}

func (p *GetTraceSummariesPaginator) CurrentPage() *GetTraceSummariesOutput {
	return p.Pager.CurrentPage().(*GetTraceSummariesOutput)
}

// GetTraceSummariesResponse is the response type for the
// GetTraceSummaries API operation.
type GetTraceSummariesResponse struct {
	*GetTraceSummariesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTraceSummaries request.
func (r *GetTraceSummariesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
