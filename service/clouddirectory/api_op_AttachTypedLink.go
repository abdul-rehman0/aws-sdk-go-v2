// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/AttachTypedLinkRequest
type AttachTypedLinkInput struct {
	_ struct{} `type:"structure"`

	// A set of attributes that are associated with the typed link.
	//
	// Attributes is a required field
	Attributes []AttributeNameAndValue `type:"list" required:"true"`

	// The Amazon Resource Name (ARN) of the directory where you want to attach
	// the typed link.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// Identifies the source object that the typed link will attach to.
	//
	// SourceObjectReference is a required field
	SourceObjectReference *ObjectReference `type:"structure" required:"true"`

	// Identifies the target object that the typed link will attach to.
	//
	// TargetObjectReference is a required field
	TargetObjectReference *ObjectReference `type:"structure" required:"true"`

	// Identifies the typed link facet that is associated with the typed link.
	//
	// TypedLinkFacet is a required field
	TypedLinkFacet *TypedLinkSchemaAndFacetName `type:"structure" required:"true"`
}

// String returns the string representation
func (s AttachTypedLinkInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachTypedLinkInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AttachTypedLinkInput"}

	if s.Attributes == nil {
		invalidParams.Add(aws.NewErrParamRequired("Attributes"))
	}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if s.SourceObjectReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceObjectReference"))
	}

	if s.TargetObjectReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetObjectReference"))
	}

	if s.TypedLinkFacet == nil {
		invalidParams.Add(aws.NewErrParamRequired("TypedLinkFacet"))
	}
	if s.Attributes != nil {
		for i, v := range s.Attributes {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Attributes", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.TypedLinkFacet != nil {
		if err := s.TypedLinkFacet.Validate(); err != nil {
			invalidParams.AddNested("TypedLinkFacet", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AttachTypedLinkInput) MarshalFields(e protocol.FieldEncoder) error {

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
	if s.SourceObjectReference != nil {
		v := s.SourceObjectReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SourceObjectReference", v, metadata)
	}
	if s.TargetObjectReference != nil {
		v := s.TargetObjectReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TargetObjectReference", v, metadata)
	}
	if s.TypedLinkFacet != nil {
		v := s.TypedLinkFacet

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TypedLinkFacet", v, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/AttachTypedLinkResponse
type AttachTypedLinkOutput struct {
	_ struct{} `type:"structure"`

	// Returns a typed link specifier as output.
	TypedLinkSpecifier *TypedLinkSpecifier `type:"structure"`
}

// String returns the string representation
func (s AttachTypedLinkOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AttachTypedLinkOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.TypedLinkSpecifier != nil {
		v := s.TypedLinkSpecifier

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TypedLinkSpecifier", v, metadata)
	}
	return nil
}

const opAttachTypedLink = "AttachTypedLink"

// AttachTypedLinkRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Attaches a typed link to a specified source and target object. For more information,
// see Typed Links (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
//
//    // Example sending a request using AttachTypedLinkRequest.
//    req := client.AttachTypedLinkRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/AttachTypedLink
func (c *Client) AttachTypedLinkRequest(input *AttachTypedLinkInput) AttachTypedLinkRequest {
	op := &aws.Operation{
		Name:       opAttachTypedLink,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/typedlink/attach",
	}

	if input == nil {
		input = &AttachTypedLinkInput{}
	}

	req := c.newRequest(op, input, &AttachTypedLinkOutput{})
	return AttachTypedLinkRequest{Request: req, Input: input, Copy: c.AttachTypedLinkRequest}
}

// AttachTypedLinkRequest is the request type for the
// AttachTypedLink API operation.
type AttachTypedLinkRequest struct {
	*aws.Request
	Input *AttachTypedLinkInput
	Copy  func(*AttachTypedLinkInput) AttachTypedLinkRequest
}

// Send marshals and sends the AttachTypedLink API request.
func (r AttachTypedLinkRequest) Send(ctx context.Context) (*AttachTypedLinkResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachTypedLinkResponse{
		AttachTypedLinkOutput: r.Request.Data.(*AttachTypedLinkOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachTypedLinkResponse is the response type for the
// AttachTypedLink API operation.
type AttachTypedLinkResponse struct {
	*AttachTypedLinkOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachTypedLink request.
func (r *AttachTypedLinkResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
