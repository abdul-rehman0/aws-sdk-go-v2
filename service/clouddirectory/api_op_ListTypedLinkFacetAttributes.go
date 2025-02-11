// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListTypedLinkFacetAttributesRequest
type ListTypedLinkFacetAttributesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to retrieve.
	MaxResults *int64 `min:"1" type:"integer"`

	// The unique name of the typed link facet.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The pagination token.
	NextToken *string `type:"string"`

	// The Amazon Resource Name (ARN) that is associated with the schema. For more
	// information, see arns.
	//
	// SchemaArn is a required field
	SchemaArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`
}

// String returns the string representation
func (s ListTypedLinkFacetAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTypedLinkFacetAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTypedLinkFacetAttributesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.SchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTypedLinkFacetAttributesInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListTypedLinkFacetAttributesResponse
type ListTypedLinkFacetAttributesOutput struct {
	_ struct{} `type:"structure"`

	// An ordered set of attributes associate with the typed link.
	Attributes []TypedLinkAttributeDefinition `type:"list"`

	// The pagination token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListTypedLinkFacetAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTypedLinkFacetAttributesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Attributes != nil {
		v := s.Attributes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Attributes", metadata)
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

const opListTypedLinkFacetAttributes = "ListTypedLinkFacetAttributes"

// ListTypedLinkFacetAttributesRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Returns a paginated list of all attribute definitions for a particular TypedLinkFacet.
// For more information, see Typed Links (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
//
//    // Example sending a request using ListTypedLinkFacetAttributesRequest.
//    req := client.ListTypedLinkFacetAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListTypedLinkFacetAttributes
func (c *Client) ListTypedLinkFacetAttributesRequest(input *ListTypedLinkFacetAttributesInput) ListTypedLinkFacetAttributesRequest {
	op := &aws.Operation{
		Name:       opListTypedLinkFacetAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/typedlink/facet/attributes",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTypedLinkFacetAttributesInput{}
	}

	req := c.newRequest(op, input, &ListTypedLinkFacetAttributesOutput{})
	return ListTypedLinkFacetAttributesRequest{Request: req, Input: input, Copy: c.ListTypedLinkFacetAttributesRequest}
}

// ListTypedLinkFacetAttributesRequest is the request type for the
// ListTypedLinkFacetAttributes API operation.
type ListTypedLinkFacetAttributesRequest struct {
	*aws.Request
	Input *ListTypedLinkFacetAttributesInput
	Copy  func(*ListTypedLinkFacetAttributesInput) ListTypedLinkFacetAttributesRequest
}

// Send marshals and sends the ListTypedLinkFacetAttributes API request.
func (r ListTypedLinkFacetAttributesRequest) Send(ctx context.Context) (*ListTypedLinkFacetAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTypedLinkFacetAttributesResponse{
		ListTypedLinkFacetAttributesOutput: r.Request.Data.(*ListTypedLinkFacetAttributesOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTypedLinkFacetAttributesRequestPaginator returns a paginator for ListTypedLinkFacetAttributes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTypedLinkFacetAttributesRequest(input)
//   p := clouddirectory.NewListTypedLinkFacetAttributesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTypedLinkFacetAttributesPaginator(req ListTypedLinkFacetAttributesRequest) ListTypedLinkFacetAttributesPaginator {
	return ListTypedLinkFacetAttributesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListTypedLinkFacetAttributesInput
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

// ListTypedLinkFacetAttributesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTypedLinkFacetAttributesPaginator struct {
	aws.Pager
}

func (p *ListTypedLinkFacetAttributesPaginator) CurrentPage() *ListTypedLinkFacetAttributesOutput {
	return p.Pager.CurrentPage().(*ListTypedLinkFacetAttributesOutput)
}

// ListTypedLinkFacetAttributesResponse is the response type for the
// ListTypedLinkFacetAttributes API operation.
type ListTypedLinkFacetAttributesResponse struct {
	*ListTypedLinkFacetAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTypedLinkFacetAttributes request.
func (r *ListTypedLinkFacetAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
