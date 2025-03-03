// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicequotas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/ListServiceQuotasRequest
type ListServiceQuotasInput struct {
	_ struct{} `type:"structure"`

	// (Optional) Limits the number of results that you want to include in the response.
	// If you don't include this parameter, the response defaults to a value that's
	// specific to the operation. If additional items exist beyond the specified
	// maximum, the NextToken element is present and has a value (isn't null). Include
	// that value as the NextToken request parameter in the call to the operation
	// to get the next part of the results. You should check NextToken after every
	// operation to ensure that you receive all of the results.
	MaxResults *int64 `min:"1" type:"integer"`

	// (Optional) Use this parameter in a request if you receive a NextToken response
	// in a previous request that indicates that there's more output available.
	// In a subsequent call, set it to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string `type:"string"`

	// The identifier for a service. When performing an operation, use the ServiceCode
	// to specify a particular service.
	//
	// ServiceCode is a required field
	ServiceCode *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListServiceQuotasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListServiceQuotasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListServiceQuotasInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ServiceCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceCode"))
	}
	if s.ServiceCode != nil && len(*s.ServiceCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServiceCode", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/ListServiceQuotasResponse
type ListServiceQuotasOutput struct {
	_ struct{} `type:"structure"`

	// If present in the response, this value indicates there's more output available
	// that what's included in the current response. This can occur even when the
	// response includes no values at all, such as when you ask for a filtered view
	// of a very long list. Use this value in the NextToken request parameter in
	// a subsequent call to the operation to continue processing and get the next
	// part of the output. You should repeat this until the NextToken response element
	// comes back empty (as null).
	NextToken *string `type:"string"`

	// The response information for a quota lists all attribute information for
	// the quota.
	Quotas []ServiceQuota `type:"list"`
}

// String returns the string representation
func (s ListServiceQuotasOutput) String() string {
	return awsutil.Prettify(s)
}

const opListServiceQuotas = "ListServiceQuotas"

// ListServiceQuotasRequest returns a request value for making API operation for
// Service Quotas.
//
// Lists all service quotas for the specified AWS service. This request returns
// a list of the service quotas for the specified service. you'll see the default
// values are the values that AWS provides for the quotas.
//
// Always check the NextToken response parameter when calling any of the List*
// operations. These operations can return an unexpected list of results, even
// when there are more results available. When this happens, the NextToken response
// parameter contains a value to pass the next call to the same API to request
// the next part of the list.
//
//    // Example sending a request using ListServiceQuotasRequest.
//    req := client.ListServiceQuotasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/ListServiceQuotas
func (c *Client) ListServiceQuotasRequest(input *ListServiceQuotasInput) ListServiceQuotasRequest {
	op := &aws.Operation{
		Name:       opListServiceQuotas,
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
		input = &ListServiceQuotasInput{}
	}

	req := c.newRequest(op, input, &ListServiceQuotasOutput{})
	return ListServiceQuotasRequest{Request: req, Input: input, Copy: c.ListServiceQuotasRequest}
}

// ListServiceQuotasRequest is the request type for the
// ListServiceQuotas API operation.
type ListServiceQuotasRequest struct {
	*aws.Request
	Input *ListServiceQuotasInput
	Copy  func(*ListServiceQuotasInput) ListServiceQuotasRequest
}

// Send marshals and sends the ListServiceQuotas API request.
func (r ListServiceQuotasRequest) Send(ctx context.Context) (*ListServiceQuotasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListServiceQuotasResponse{
		ListServiceQuotasOutput: r.Request.Data.(*ListServiceQuotasOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListServiceQuotasRequestPaginator returns a paginator for ListServiceQuotas.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListServiceQuotasRequest(input)
//   p := servicequotas.NewListServiceQuotasRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListServiceQuotasPaginator(req ListServiceQuotasRequest) ListServiceQuotasPaginator {
	return ListServiceQuotasPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListServiceQuotasInput
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

// ListServiceQuotasPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListServiceQuotasPaginator struct {
	aws.Pager
}

func (p *ListServiceQuotasPaginator) CurrentPage() *ListServiceQuotasOutput {
	return p.Pager.CurrentPage().(*ListServiceQuotasOutput)
}

// ListServiceQuotasResponse is the response type for the
// ListServiceQuotas API operation.
type ListServiceQuotasResponse struct {
	*ListServiceQuotasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListServiceQuotas request.
func (r *ListServiceQuotasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
