// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/BatchGetDevEndpointsRequest
type BatchGetDevEndpointsInput struct {
	_ struct{} `type:"structure"`

	// The list of DevEndpoint names, which might be the names returned from the
	// ListDevEndpoint operation.
	//
	// DevEndpointNames is a required field
	DevEndpointNames []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchGetDevEndpointsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchGetDevEndpointsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchGetDevEndpointsInput"}

	if s.DevEndpointNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("DevEndpointNames"))
	}
	if s.DevEndpointNames != nil && len(s.DevEndpointNames) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DevEndpointNames", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/BatchGetDevEndpointsResponse
type BatchGetDevEndpointsOutput struct {
	_ struct{} `type:"structure"`

	// A list of DevEndpoint definitions.
	DevEndpoints []DevEndpoint `type:"list"`

	// A list of DevEndpoints not found.
	DevEndpointsNotFound []string `min:"1" type:"list"`
}

// String returns the string representation
func (s BatchGetDevEndpointsOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchGetDevEndpoints = "BatchGetDevEndpoints"

// BatchGetDevEndpointsRequest returns a request value for making API operation for
// AWS Glue.
//
// Returns a list of resource metadata for a given list of development endpoint
// names. After calling the ListDevEndpoints operation, you can call this operation
// to access the data to which you have been granted permissions. This operation
// supports all IAM permissions, including permission conditions that uses tags.
//
//    // Example sending a request using BatchGetDevEndpointsRequest.
//    req := client.BatchGetDevEndpointsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/BatchGetDevEndpoints
func (c *Client) BatchGetDevEndpointsRequest(input *BatchGetDevEndpointsInput) BatchGetDevEndpointsRequest {
	op := &aws.Operation{
		Name:       opBatchGetDevEndpoints,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchGetDevEndpointsInput{}
	}

	req := c.newRequest(op, input, &BatchGetDevEndpointsOutput{})
	return BatchGetDevEndpointsRequest{Request: req, Input: input, Copy: c.BatchGetDevEndpointsRequest}
}

// BatchGetDevEndpointsRequest is the request type for the
// BatchGetDevEndpoints API operation.
type BatchGetDevEndpointsRequest struct {
	*aws.Request
	Input *BatchGetDevEndpointsInput
	Copy  func(*BatchGetDevEndpointsInput) BatchGetDevEndpointsRequest
}

// Send marshals and sends the BatchGetDevEndpoints API request.
func (r BatchGetDevEndpointsRequest) Send(ctx context.Context) (*BatchGetDevEndpointsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchGetDevEndpointsResponse{
		BatchGetDevEndpointsOutput: r.Request.Data.(*BatchGetDevEndpointsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchGetDevEndpointsResponse is the response type for the
// BatchGetDevEndpoints API operation.
type BatchGetDevEndpointsResponse struct {
	*BatchGetDevEndpointsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchGetDevEndpoints request.
func (r *BatchGetDevEndpointsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
