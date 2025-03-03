// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/CreateDetectorRequest
type CreateDetectorInput struct {
	_ struct{} `type:"structure"`

	// The idempotency token for the create request.
	ClientToken *string `locationName:"clientToken" type:"string" idempotencyToken:"true"`

	// A boolean value that specifies whether the detector is to be enabled.
	//
	// Enable is a required field
	Enable *bool `locationName:"enable" type:"boolean" required:"true"`

	// A enum value that specifies how frequently customer got Finding updates published.
	FindingPublishingFrequency FindingPublishingFrequency `locationName:"findingPublishingFrequency" type:"string" enum:"true"`

	// The tags to be added to a new detector resource.
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`
}

// String returns the string representation
func (s CreateDetectorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDetectorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDetectorInput"}

	if s.Enable == nil {
		invalidParams.Add(aws.NewErrParamRequired("Enable"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDetectorInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	var ClientToken string
	if s.ClientToken != nil {
		ClientToken = *s.ClientToken
	} else {
		ClientToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Enable != nil {
		v := *s.Enable

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enable", protocol.BoolValue(v), metadata)
	}
	if len(s.FindingPublishingFrequency) > 0 {
		v := s.FindingPublishingFrequency

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "findingPublishingFrequency", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/CreateDetectorResponse
type CreateDetectorOutput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the created detector.
	DetectorId *string `locationName:"detectorId" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateDetectorOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDetectorOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateDetector = "CreateDetector"

// CreateDetectorRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Creates a single Amazon GuardDuty detector. A detector is an object that
// represents the GuardDuty service. A detector must be created in order for
// GuardDuty to become operational.
//
//    // Example sending a request using CreateDetectorRequest.
//    req := client.CreateDetectorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/CreateDetector
func (c *Client) CreateDetectorRequest(input *CreateDetectorInput) CreateDetectorRequest {
	op := &aws.Operation{
		Name:       opCreateDetector,
		HTTPMethod: "POST",
		HTTPPath:   "/detector",
	}

	if input == nil {
		input = &CreateDetectorInput{}
	}

	req := c.newRequest(op, input, &CreateDetectorOutput{})
	return CreateDetectorRequest{Request: req, Input: input, Copy: c.CreateDetectorRequest}
}

// CreateDetectorRequest is the request type for the
// CreateDetector API operation.
type CreateDetectorRequest struct {
	*aws.Request
	Input *CreateDetectorInput
	Copy  func(*CreateDetectorInput) CreateDetectorRequest
}

// Send marshals and sends the CreateDetector API request.
func (r CreateDetectorRequest) Send(ctx context.Context) (*CreateDetectorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDetectorResponse{
		CreateDetectorOutput: r.Request.Data.(*CreateDetectorOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDetectorResponse is the response type for the
// CreateDetector API operation.
type CreateDetectorResponse struct {
	*CreateDetectorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDetector request.
func (r *CreateDetectorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
