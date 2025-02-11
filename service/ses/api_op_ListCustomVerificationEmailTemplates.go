// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to list the existing custom verification email templates
// for your account.
//
// For more information about custom verification email templates, see Using
// Custom Verification Email Templates (ses/latest/DeveloperGuide/custom-verification-emails.html)
// in the Amazon SES Developer Guide.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/ListCustomVerificationEmailTemplatesRequest
type ListCustomVerificationEmailTemplatesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of custom verification email templates to return. This
	// value must be at least 1 and less than or equal to 50. If you do not specify
	// a value, or if you specify a value less than 1 or greater than 50, the operation
	// will return up to 50 results.
	MaxResults *int64 `min:"1" type:"integer"`

	// An array the contains the name and creation time stamp for each template
	// in your Amazon SES account.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListCustomVerificationEmailTemplatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListCustomVerificationEmailTemplatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListCustomVerificationEmailTemplatesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A paginated list of custom verification email templates.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/ListCustomVerificationEmailTemplatesResponse
type ListCustomVerificationEmailTemplatesOutput struct {
	_ struct{} `type:"structure"`

	// A list of the custom verification email templates that exist in your account.
	CustomVerificationEmailTemplates []CustomVerificationEmailTemplate `type:"list"`

	// A token indicating that there are additional custom verification email templates
	// available to be listed. Pass this token to a subsequent call to ListTemplates
	// to retrieve the next 50 custom verification email templates.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListCustomVerificationEmailTemplatesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListCustomVerificationEmailTemplates = "ListCustomVerificationEmailTemplates"

// ListCustomVerificationEmailTemplatesRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Lists the existing custom verification email templates for your account in
// the current AWS Region.
//
// For more information about custom verification email templates, see Using
// Custom Verification Email Templates (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/custom-verification-emails.html)
// in the Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using ListCustomVerificationEmailTemplatesRequest.
//    req := client.ListCustomVerificationEmailTemplatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/ListCustomVerificationEmailTemplates
func (c *Client) ListCustomVerificationEmailTemplatesRequest(input *ListCustomVerificationEmailTemplatesInput) ListCustomVerificationEmailTemplatesRequest {
	op := &aws.Operation{
		Name:       opListCustomVerificationEmailTemplates,
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
		input = &ListCustomVerificationEmailTemplatesInput{}
	}

	req := c.newRequest(op, input, &ListCustomVerificationEmailTemplatesOutput{})
	return ListCustomVerificationEmailTemplatesRequest{Request: req, Input: input, Copy: c.ListCustomVerificationEmailTemplatesRequest}
}

// ListCustomVerificationEmailTemplatesRequest is the request type for the
// ListCustomVerificationEmailTemplates API operation.
type ListCustomVerificationEmailTemplatesRequest struct {
	*aws.Request
	Input *ListCustomVerificationEmailTemplatesInput
	Copy  func(*ListCustomVerificationEmailTemplatesInput) ListCustomVerificationEmailTemplatesRequest
}

// Send marshals and sends the ListCustomVerificationEmailTemplates API request.
func (r ListCustomVerificationEmailTemplatesRequest) Send(ctx context.Context) (*ListCustomVerificationEmailTemplatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListCustomVerificationEmailTemplatesResponse{
		ListCustomVerificationEmailTemplatesOutput: r.Request.Data.(*ListCustomVerificationEmailTemplatesOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListCustomVerificationEmailTemplatesRequestPaginator returns a paginator for ListCustomVerificationEmailTemplates.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListCustomVerificationEmailTemplatesRequest(input)
//   p := ses.NewListCustomVerificationEmailTemplatesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListCustomVerificationEmailTemplatesPaginator(req ListCustomVerificationEmailTemplatesRequest) ListCustomVerificationEmailTemplatesPaginator {
	return ListCustomVerificationEmailTemplatesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListCustomVerificationEmailTemplatesInput
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

// ListCustomVerificationEmailTemplatesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListCustomVerificationEmailTemplatesPaginator struct {
	aws.Pager
}

func (p *ListCustomVerificationEmailTemplatesPaginator) CurrentPage() *ListCustomVerificationEmailTemplatesOutput {
	return p.Pager.CurrentPage().(*ListCustomVerificationEmailTemplatesOutput)
}

// ListCustomVerificationEmailTemplatesResponse is the response type for the
// ListCustomVerificationEmailTemplates API operation.
type ListCustomVerificationEmailTemplatesResponse struct {
	*ListCustomVerificationEmailTemplatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListCustomVerificationEmailTemplates request.
func (r *ListCustomVerificationEmailTemplatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
