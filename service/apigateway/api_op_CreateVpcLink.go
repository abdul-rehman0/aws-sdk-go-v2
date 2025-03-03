// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Creates a VPC link, under the caller's account in a selected region, in an
// asynchronous operation that typically takes 2-4 minutes to complete and become
// operational. The caller must have permissions to create and update VPC Endpoint
// services.
type CreateVpcLinkInput struct {
	_ struct{} `type:"structure"`

	// The description of the VPC link.
	Description *string `locationName:"description" type:"string"`

	// [Required] The name used to label and identify the VPC link.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/].
	// The tag key can be up to 128 characters and must not start with aws:. The
	// tag value can be up to 256 characters.
	Tags map[string]string `locationName:"tags" type:"map"`

	// [Required] The ARNs of network load balancers of the VPC targeted by the
	// VPC link. The network load balancers must be owned by the same AWS account
	// of the API owner.
	//
	// TargetArns is a required field
	TargetArns []string `locationName:"targetArns" type:"list" required:"true"`
}

// String returns the string representation
func (s CreateVpcLinkInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVpcLinkInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVpcLinkInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.TargetArns == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetArns"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateVpcLinkInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.TargetArns != nil {
		v := s.TargetArns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "targetArns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// A API Gateway VPC link for a RestApi to access resources in an Amazon Virtual
// Private Cloud (VPC).
//
// To enable access to a resource in an Amazon Virtual Private Cloud through
// Amazon API Gateway, you, as an API developer, create a VpcLink resource targeted
// for one or more network load balancers of the VPC and then integrate an API
// method with a private integration that uses the VpcLink. The private integration
// has an integration type of HTTP or HTTP_PROXY and has a connection type of
// VPC_LINK. The integration uses the connectionId property to identify the
// VpcLink used.
type CreateVpcLinkOutput struct {
	_ struct{} `type:"structure"`

	// The description of the VPC link.
	Description *string `locationName:"description" type:"string"`

	// The identifier of the VpcLink. It is used in an Integration to reference
	// this VpcLink.
	Id *string `locationName:"id" type:"string"`

	// The name used to label and identify the VPC link.
	Name *string `locationName:"name" type:"string"`

	// The status of the VPC link. The valid values are AVAILABLE, PENDING, DELETING,
	// or FAILED. Deploying an API will wait if the status is PENDING and will fail
	// if the status is DELETING.
	Status VpcLinkStatus `locationName:"status" type:"string" enum:"true"`

	// A description about the VPC link status.
	StatusMessage *string `locationName:"statusMessage" type:"string"`

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string `locationName:"tags" type:"map"`

	// The ARNs of network load balancers of the VPC targeted by the VPC link. The
	// network load balancers must be owned by the same AWS account of the API owner.
	TargetArns []string `locationName:"targetArns" type:"list"`
}

// String returns the string representation
func (s CreateVpcLinkOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateVpcLinkOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StatusMessage != nil {
		v := *s.StatusMessage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "statusMessage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.TargetArns != nil {
		v := s.TargetArns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "targetArns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opCreateVpcLink = "CreateVpcLink"

// CreateVpcLinkRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Creates a VPC link, under the caller's account in a selected region, in an
// asynchronous operation that typically takes 2-4 minutes to complete and become
// operational. The caller must have permissions to create and update VPC Endpoint
// services.
//
//    // Example sending a request using CreateVpcLinkRequest.
//    req := client.CreateVpcLinkRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateVpcLinkRequest(input *CreateVpcLinkInput) CreateVpcLinkRequest {
	op := &aws.Operation{
		Name:       opCreateVpcLink,
		HTTPMethod: "POST",
		HTTPPath:   "/vpclinks",
	}

	if input == nil {
		input = &CreateVpcLinkInput{}
	}

	req := c.newRequest(op, input, &CreateVpcLinkOutput{})
	return CreateVpcLinkRequest{Request: req, Input: input, Copy: c.CreateVpcLinkRequest}
}

// CreateVpcLinkRequest is the request type for the
// CreateVpcLink API operation.
type CreateVpcLinkRequest struct {
	*aws.Request
	Input *CreateVpcLinkInput
	Copy  func(*CreateVpcLinkInput) CreateVpcLinkRequest
}

// Send marshals and sends the CreateVpcLink API request.
func (r CreateVpcLinkRequest) Send(ctx context.Context) (*CreateVpcLinkResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVpcLinkResponse{
		CreateVpcLinkOutput: r.Request.Data.(*CreateVpcLinkOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVpcLinkResponse is the response type for the
// CreateVpcLink API operation.
type CreateVpcLinkResponse struct {
	*CreateVpcLinkOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVpcLink request.
func (r *CreateVpcLinkResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
