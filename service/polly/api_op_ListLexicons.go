// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package polly

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/ListLexiconsInput
type ListLexiconsInput struct {
	_ struct{} `type:"structure"`

	// An opaque pagination token returned from previous ListLexicons operation.
	// If present, indicates where to continue the list of lexicons.
	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s ListLexiconsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListLexiconsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/ListLexiconsOutput
type ListLexiconsOutput struct {
	_ struct{} `type:"structure"`

	// A list of lexicon names and attributes.
	Lexicons []LexiconDescription `type:"list"`

	// The pagination token to use in the next request to continue the listing of
	// lexicons. NextToken is returned only if the response is truncated.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListLexiconsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListLexiconsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Lexicons != nil {
		v := s.Lexicons

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Lexicons", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListLexicons = "ListLexicons"

// ListLexiconsRequest returns a request value for making API operation for
// Amazon Polly.
//
// Returns a list of pronunciation lexicons stored in an AWS Region. For more
// information, see Managing Lexicons (http://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html).
//
//    // Example sending a request using ListLexiconsRequest.
//    req := client.ListLexiconsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/ListLexicons
func (c *Client) ListLexiconsRequest(input *ListLexiconsInput) ListLexiconsRequest {
	op := &aws.Operation{
		Name:       opListLexicons,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/lexicons",
	}

	if input == nil {
		input = &ListLexiconsInput{}
	}

	req := c.newRequest(op, input, &ListLexiconsOutput{})
	return ListLexiconsRequest{Request: req, Input: input, Copy: c.ListLexiconsRequest}
}

// ListLexiconsRequest is the request type for the
// ListLexicons API operation.
type ListLexiconsRequest struct {
	*aws.Request
	Input *ListLexiconsInput
	Copy  func(*ListLexiconsInput) ListLexiconsRequest
}

// Send marshals and sends the ListLexicons API request.
func (r ListLexiconsRequest) Send(ctx context.Context) (*ListLexiconsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListLexiconsResponse{
		ListLexiconsOutput: r.Request.Data.(*ListLexiconsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListLexiconsResponse is the response type for the
// ListLexicons API operation.
type ListLexiconsResponse struct {
	*ListLexiconsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListLexicons request.
func (r *ListLexiconsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
