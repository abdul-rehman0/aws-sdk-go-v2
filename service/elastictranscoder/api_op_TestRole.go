// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elastictranscoder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The TestRoleRequest structure.
type TestRoleInput struct {
	_ struct{} `deprecated:"true" type:"structure"`

	// The Amazon S3 bucket that contains media files to be transcoded. The action
	// attempts to read from this bucket.
	//
	// InputBucket is a required field
	InputBucket *string `type:"string" required:"true"`

	// The Amazon S3 bucket that Elastic Transcoder writes transcoded media files
	// to. The action attempts to read from this bucket.
	//
	// OutputBucket is a required field
	OutputBucket *string `type:"string" required:"true"`

	// The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder
	// to test.
	//
	// Role is a required field
	Role *string `type:"string" required:"true"`

	// The ARNs of one or more Amazon Simple Notification Service (Amazon SNS) topics
	// that you want the action to send a test notification to.
	//
	// Topics is a required field
	Topics []string `type:"list" required:"true"`
}

// String returns the string representation
func (s TestRoleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TestRoleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TestRoleInput"}

	if s.InputBucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputBucket"))
	}

	if s.OutputBucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("OutputBucket"))
	}

	if s.Role == nil {
		invalidParams.Add(aws.NewErrParamRequired("Role"))
	}

	if s.Topics == nil {
		invalidParams.Add(aws.NewErrParamRequired("Topics"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TestRoleInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.InputBucket != nil {
		v := *s.InputBucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "InputBucket", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OutputBucket != nil {
		v := *s.OutputBucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "OutputBucket", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Role != nil {
		v := *s.Role

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Role", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Topics != nil {
		v := s.Topics

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Topics", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// The TestRoleResponse structure.
type TestRoleOutput struct {
	_ struct{} `deprecated:"true" type:"structure"`

	// If the Success element contains false, this value is an array of one or more
	// error messages that were generated during the test process.
	Messages []string `type:"list"`

	// If the operation is successful, this value is true; otherwise, the value
	// is false.
	Success *string `type:"string"`
}

// String returns the string representation
func (s TestRoleOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TestRoleOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Messages != nil {
		v := s.Messages

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Messages", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Success != nil {
		v := *s.Success

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Success", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opTestRole = "TestRole"

// TestRoleRequest returns a request value for making API operation for
// Amazon Elastic Transcoder.
//
// The TestRole operation tests the IAM role used to create the pipeline.
//
// The TestRole action lets you determine whether the IAM role you are using
// has sufficient permissions to let Elastic Transcoder perform tasks associated
// with the transcoding process. The action attempts to assume the specified
// IAM role, checks read access to the input and output buckets, and tries to
// send a test notification to Amazon SNS topics that you specify.
//
//    // Example sending a request using TestRoleRequest.
//    req := client.TestRoleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) TestRoleRequest(input *TestRoleInput) TestRoleRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, TestRole, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opTestRole,
		HTTPMethod: "POST",
		HTTPPath:   "/2012-09-25/roleTests",
	}

	if input == nil {
		input = &TestRoleInput{}
	}

	req := c.newRequest(op, input, &TestRoleOutput{})
	return TestRoleRequest{Request: req, Input: input, Copy: c.TestRoleRequest}
}

// TestRoleRequest is the request type for the
// TestRole API operation.
type TestRoleRequest struct {
	*aws.Request
	Input *TestRoleInput
	Copy  func(*TestRoleInput) TestRoleRequest
}

// Send marshals and sends the TestRole API request.
func (r TestRoleRequest) Send(ctx context.Context) (*TestRoleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TestRoleResponse{
		TestRoleOutput: r.Request.Data.(*TestRoleOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TestRoleResponse is the response type for the
// TestRole API operation.
type TestRoleResponse struct {
	*TestRoleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TestRole request.
func (r *TestRoleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
