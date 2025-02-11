// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/CreateMeshInput
type CreateMeshInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of therequest. Up to 36 letters, numbers, hyphens, and underscores are allowed.
	ClientToken *string `locationName:"clientToken" type:"string" idempotencyToken:"true"`

	// The name to use for the service mesh.
	//
	// MeshName is a required field
	MeshName *string `locationName:"meshName" min:"1" type:"string" required:"true"`

	// The service mesh specification to apply.
	Spec *MeshSpec `locationName:"spec" type:"structure"`

	// Optional metadata that you can apply to the service mesh to assist with categorization
	// and organization. Each tag consists of a key and an optional value, both
	// of which you define. Tag keys can have a maximum character length of 128
	// characters, and tag values can have a maximum length of 256 characters.
	Tags []TagRef `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreateMeshInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMeshInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateMeshInput"}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}
	if s.Spec != nil {
		if err := s.Spec.Validate(); err != nil {
			invalidParams.AddNested("Spec", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMeshInput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.MeshName != nil {
		v := *s.MeshName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "meshName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Spec != nil {
		v := s.Spec

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "spec", v, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/CreateMeshOutput
type CreateMeshOutput struct {
	_ struct{} `type:"structure" payload:"Mesh"`

	// The full description of your service mesh following the create call.
	//
	// Mesh is a required field
	Mesh *MeshData `locationName:"mesh" type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateMeshOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMeshOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Mesh != nil {
		v := s.Mesh

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "mesh", v, metadata)
	}
	return nil
}

const opCreateMesh = "CreateMesh"

// CreateMeshRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Creates a service mesh. A service mesh is a logical boundary for network
// traffic between the services that reside within it.
//
// After you create your service mesh, you can create virtual services, virtual
// nodes, virtual routers, and routes to distribute traffic between the applications
// in your mesh.
//
//    // Example sending a request using CreateMeshRequest.
//    req := client.CreateMeshRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/CreateMesh
func (c *Client) CreateMeshRequest(input *CreateMeshInput) CreateMeshRequest {
	op := &aws.Operation{
		Name:       opCreateMesh,
		HTTPMethod: "PUT",
		HTTPPath:   "/v20190125/meshes",
	}

	if input == nil {
		input = &CreateMeshInput{}
	}

	req := c.newRequest(op, input, &CreateMeshOutput{})
	return CreateMeshRequest{Request: req, Input: input, Copy: c.CreateMeshRequest}
}

// CreateMeshRequest is the request type for the
// CreateMesh API operation.
type CreateMeshRequest struct {
	*aws.Request
	Input *CreateMeshInput
	Copy  func(*CreateMeshInput) CreateMeshRequest
}

// Send marshals and sends the CreateMesh API request.
func (r CreateMeshRequest) Send(ctx context.Context) (*CreateMeshResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMeshResponse{
		CreateMeshOutput: r.Request.Data.(*CreateMeshOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMeshResponse is the response type for the
// CreateMesh API operation.
type CreateMeshResponse struct {
	*CreateMeshOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMesh request.
func (r *CreateMeshResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
