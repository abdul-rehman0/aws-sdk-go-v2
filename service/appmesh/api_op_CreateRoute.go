// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/CreateRouteInput
type CreateRouteInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of therequest. Up to 36 letters, numbers, hyphens, and underscores are allowed.
	ClientToken *string `locationName:"clientToken" type:"string" idempotencyToken:"true"`

	// The name of the service mesh to create the route in.
	//
	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	// The name to use for the route.
	//
	// RouteName is a required field
	RouteName *string `locationName:"routeName" min:"1" type:"string" required:"true"`

	// The route specification to apply.
	//
	// Spec is a required field
	Spec *RouteSpec `locationName:"spec" type:"structure" required:"true"`

	// Optional metadata that you can apply to the route to assist with categorization
	// and organization. Each tag consists of a key and an optional value, both
	// of which you define. Tag keys can have a maximum character length of 128
	// characters, and tag values can have a maximum length of 256 characters.
	Tags []TagRef `locationName:"tags" type:"list"`

	// The name of the virtual router in which to create the route.
	//
	// VirtualRouterName is a required field
	VirtualRouterName *string `location:"uri" locationName:"virtualRouterName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateRouteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRouteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRouteInput"}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}

	if s.RouteName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteName"))
	}
	if s.RouteName != nil && len(*s.RouteName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RouteName", 1))
	}

	if s.Spec == nil {
		invalidParams.Add(aws.NewErrParamRequired("Spec"))
	}

	if s.VirtualRouterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualRouterName"))
	}
	if s.VirtualRouterName != nil && len(*s.VirtualRouterName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VirtualRouterName", 1))
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
func (s CreateRouteInput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.RouteName != nil {
		v := *s.RouteName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "routeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.MeshName != nil {
		v := *s.MeshName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meshName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VirtualRouterName != nil {
		v := *s.VirtualRouterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "virtualRouterName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/CreateRouteOutput
type CreateRouteOutput struct {
	_ struct{} `type:"structure" payload:"Route"`

	// The full description of your mesh following the create call.
	//
	// Route is a required field
	Route *RouteData `locationName:"route" type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateRouteOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateRouteOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Route != nil {
		v := s.Route

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "route", v, metadata)
	}
	return nil
}

const opCreateRoute = "CreateRoute"

// CreateRouteRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Creates a route that is associated with a virtual router.
//
// You can use the prefix parameter in your route specification for path-based
// routing of requests. For example, if your virtual service name is my-service.local
// and you want the route to match requests to my-service.local/metrics, your
// prefix should be /metrics.
//
// If your route matches a request, you can distribute traffic to one or more
// target virtual nodes with relative weighting.
//
//    // Example sending a request using CreateRouteRequest.
//    req := client.CreateRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/CreateRoute
func (c *Client) CreateRouteRequest(input *CreateRouteInput) CreateRouteRequest {
	op := &aws.Operation{
		Name:       opCreateRoute,
		HTTPMethod: "PUT",
		HTTPPath:   "/v20190125/meshes/{meshName}/virtualRouter/{virtualRouterName}/routes",
	}

	if input == nil {
		input = &CreateRouteInput{}
	}

	req := c.newRequest(op, input, &CreateRouteOutput{})
	return CreateRouteRequest{Request: req, Input: input, Copy: c.CreateRouteRequest}
}

// CreateRouteRequest is the request type for the
// CreateRoute API operation.
type CreateRouteRequest struct {
	*aws.Request
	Input *CreateRouteInput
	Copy  func(*CreateRouteInput) CreateRouteRequest
}

// Send marshals and sends the CreateRoute API request.
func (r CreateRouteRequest) Send(ctx context.Context) (*CreateRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRouteResponse{
		CreateRouteOutput: r.Request.Data.(*CreateRouteOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRouteResponse is the response type for the
// CreateRoute API operation.
type CreateRouteResponse struct {
	*CreateRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRoute request.
func (r *CreateRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
