// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/RunPipelineActivityRequest
type RunPipelineActivityInput struct {
	_ struct{} `type:"structure"`

	// The sample message payloads on which the pipeline activity is run.
	//
	// Payloads is a required field
	Payloads [][]byte `locationName:"payloads" min:"1" type:"list" required:"true"`

	// The pipeline activity that is run. This must not be a 'channel' activity
	// or a 'datastore' activity because these activities are used in a pipeline
	// only to load the original message and to store the (possibly) transformed
	// message. If a 'lambda' activity is specified, only short-running Lambda functions
	// (those with a timeout of less than 30 seconds or less) can be used.
	//
	// PipelineActivity is a required field
	PipelineActivity *PipelineActivity `locationName:"pipelineActivity" type:"structure" required:"true"`
}

// String returns the string representation
func (s RunPipelineActivityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RunPipelineActivityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RunPipelineActivityInput"}

	if s.Payloads == nil {
		invalidParams.Add(aws.NewErrParamRequired("Payloads"))
	}
	if s.Payloads != nil && len(s.Payloads) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Payloads", 1))
	}

	if s.PipelineActivity == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineActivity"))
	}
	if s.PipelineActivity != nil {
		if err := s.PipelineActivity.Validate(); err != nil {
			invalidParams.AddNested("PipelineActivity", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RunPipelineActivityInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Payloads != nil {
		v := s.Payloads

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "payloads", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.BytesValue(v1)})
		}
		ls0.End()

	}
	if s.PipelineActivity != nil {
		v := s.PipelineActivity

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "pipelineActivity", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/RunPipelineActivityResponse
type RunPipelineActivityOutput struct {
	_ struct{} `type:"structure"`

	// In case the pipeline activity fails, the log message that is generated.
	LogResult *string `locationName:"logResult" type:"string"`

	// The enriched or transformed sample message payloads as base64-encoded strings.
	// (The results of running the pipeline activity on each input sample message
	// payload, encoded in base64.)
	Payloads [][]byte `locationName:"payloads" min:"1" type:"list"`
}

// String returns the string representation
func (s RunPipelineActivityOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RunPipelineActivityOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.LogResult != nil {
		v := *s.LogResult

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "logResult", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Payloads != nil {
		v := s.Payloads

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "payloads", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.BytesValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opRunPipelineActivity = "RunPipelineActivity"

// RunPipelineActivityRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Simulates the results of running a pipeline activity on a message payload.
//
//    // Example sending a request using RunPipelineActivityRequest.
//    req := client.RunPipelineActivityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/RunPipelineActivity
func (c *Client) RunPipelineActivityRequest(input *RunPipelineActivityInput) RunPipelineActivityRequest {
	op := &aws.Operation{
		Name:       opRunPipelineActivity,
		HTTPMethod: "POST",
		HTTPPath:   "/pipelineactivities/run",
	}

	if input == nil {
		input = &RunPipelineActivityInput{}
	}

	req := c.newRequest(op, input, &RunPipelineActivityOutput{})
	return RunPipelineActivityRequest{Request: req, Input: input, Copy: c.RunPipelineActivityRequest}
}

// RunPipelineActivityRequest is the request type for the
// RunPipelineActivity API operation.
type RunPipelineActivityRequest struct {
	*aws.Request
	Input *RunPipelineActivityInput
	Copy  func(*RunPipelineActivityInput) RunPipelineActivityRequest
}

// Send marshals and sends the RunPipelineActivity API request.
func (r RunPipelineActivityRequest) Send(ctx context.Context) (*RunPipelineActivityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RunPipelineActivityResponse{
		RunPipelineActivityOutput: r.Request.Data.(*RunPipelineActivityOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RunPipelineActivityResponse is the response type for the
// RunPipelineActivity API operation.
type RunPipelineActivityResponse struct {
	*RunPipelineActivityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RunPipelineActivity request.
func (r *RunPipelineActivityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
