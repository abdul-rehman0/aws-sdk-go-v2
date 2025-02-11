// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DeleteLabelsRequest
type DeleteLabelsInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Do not set this field when using administrative
	// API actions, as in accessing the API using AWS credentials.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string"`

	// Flag to request removal of all labels from the specified resource.
	DeleteAll *bool `location:"querystring" locationName:"deleteAll" type:"boolean"`

	// List of labels to delete from the resource.
	Labels []string `location:"querystring" locationName:"labels" type:"list"`

	// The ID of the resource.
	//
	// ResourceId is a required field
	ResourceId *string `location:"uri" locationName:"ResourceId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteLabelsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLabelsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLabelsInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if s.ResourceId != nil && len(*s.ResourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteLabelsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.AuthenticationToken != nil {
		v := *s.AuthenticationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Authentication", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceId != nil {
		v := *s.ResourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ResourceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DeleteAll != nil {
		v := *s.DeleteAll

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "deleteAll", protocol.BoolValue(v), metadata)
	}
	if s.Labels != nil {
		v := s.Labels

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.QueryTarget, "labels", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DeleteLabelsResponse
type DeleteLabelsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteLabelsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteLabelsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteLabels = "DeleteLabels"

// DeleteLabelsRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Deletes the specified list of labels from a resource.
//
//    // Example sending a request using DeleteLabelsRequest.
//    req := client.DeleteLabelsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DeleteLabels
func (c *Client) DeleteLabelsRequest(input *DeleteLabelsInput) DeleteLabelsRequest {
	op := &aws.Operation{
		Name:       opDeleteLabels,
		HTTPMethod: "DELETE",
		HTTPPath:   "/api/v1/resources/{ResourceId}/labels",
	}

	if input == nil {
		input = &DeleteLabelsInput{}
	}

	req := c.newRequest(op, input, &DeleteLabelsOutput{})
	return DeleteLabelsRequest{Request: req, Input: input, Copy: c.DeleteLabelsRequest}
}

// DeleteLabelsRequest is the request type for the
// DeleteLabels API operation.
type DeleteLabelsRequest struct {
	*aws.Request
	Input *DeleteLabelsInput
	Copy  func(*DeleteLabelsInput) DeleteLabelsRequest
}

// Send marshals and sends the DeleteLabels API request.
func (r DeleteLabelsRequest) Send(ctx context.Context) (*DeleteLabelsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteLabelsResponse{
		DeleteLabelsOutput: r.Request.Data.(*DeleteLabelsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteLabelsResponse is the response type for the
// DeleteLabels API operation.
type DeleteLabelsResponse struct {
	*DeleteLabelsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteLabels request.
func (r *DeleteLabelsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
