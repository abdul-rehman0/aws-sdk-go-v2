// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/DescribeSchemaRequest
type DescribeSchemaInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the schema to retrieve.
	//
	// SchemaArn is a required field
	SchemaArn *string `locationName:"schemaArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeSchemaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSchemaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSchemaInput"}

	if s.SchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/DescribeSchemaResponse
type DescribeSchemaOutput struct {
	_ struct{} `type:"structure"`

	// The requested schema.
	Schema *DatasetSchema `locationName:"schema" type:"structure"`
}

// String returns the string representation
func (s DescribeSchemaOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSchema = "DescribeSchema"

// DescribeSchemaRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Describes a schema. For more information on schemas, see CreateSchema.
//
//    // Example sending a request using DescribeSchemaRequest.
//    req := client.DescribeSchemaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/DescribeSchema
func (c *Client) DescribeSchemaRequest(input *DescribeSchemaInput) DescribeSchemaRequest {
	op := &aws.Operation{
		Name:       opDescribeSchema,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSchemaInput{}
	}

	req := c.newRequest(op, input, &DescribeSchemaOutput{})
	return DescribeSchemaRequest{Request: req, Input: input, Copy: c.DescribeSchemaRequest}
}

// DescribeSchemaRequest is the request type for the
// DescribeSchema API operation.
type DescribeSchemaRequest struct {
	*aws.Request
	Input *DescribeSchemaInput
	Copy  func(*DescribeSchemaInput) DescribeSchemaRequest
}

// Send marshals and sends the DescribeSchema API request.
func (r DescribeSchemaRequest) Send(ctx context.Context) (*DescribeSchemaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSchemaResponse{
		DescribeSchemaOutput: r.Request.Data.(*DescribeSchemaOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSchemaResponse is the response type for the
// DescribeSchema API operation.
type DescribeSchemaResponse struct {
	*DescribeSchemaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSchema request.
func (r *DescribeSchemaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
