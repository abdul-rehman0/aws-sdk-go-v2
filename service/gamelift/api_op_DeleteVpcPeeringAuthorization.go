// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DeleteVpcPeeringAuthorizationInput
type DeleteVpcPeeringAuthorizationInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for the AWS account that you use to manage your Amazon
	// GameLift fleet. You can find your Account ID in the AWS Management Console
	// under account settings.
	//
	// GameLiftAwsAccountId is a required field
	GameLiftAwsAccountId *string `min:"1" type:"string" required:"true"`

	// Unique identifier for a VPC with resources to be accessed by your Amazon
	// GameLift fleet. The VPC must be in the same region where your fleet is deployed.
	// Look up a VPC ID using the VPC Dashboard (https://console.aws.amazon.com/vpc/)
	// in the AWS Management Console. Learn more about VPC peering in VPC Peering
	// with Amazon GameLift Fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html).
	//
	// PeerVpcId is a required field
	PeerVpcId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVpcPeeringAuthorizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVpcPeeringAuthorizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVpcPeeringAuthorizationInput"}

	if s.GameLiftAwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GameLiftAwsAccountId"))
	}
	if s.GameLiftAwsAccountId != nil && len(*s.GameLiftAwsAccountId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameLiftAwsAccountId", 1))
	}

	if s.PeerVpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PeerVpcId"))
	}
	if s.PeerVpcId != nil && len(*s.PeerVpcId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PeerVpcId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DeleteVpcPeeringAuthorizationOutput
type DeleteVpcPeeringAuthorizationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteVpcPeeringAuthorizationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteVpcPeeringAuthorization = "DeleteVpcPeeringAuthorization"

// DeleteVpcPeeringAuthorizationRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Cancels a pending VPC peering authorization for the specified VPC. If you
// need to delete an existing VPC peering connection, call DeleteVpcPeeringConnection.
//
//    * CreateVpcPeeringAuthorization
//
//    * DescribeVpcPeeringAuthorizations
//
//    * DeleteVpcPeeringAuthorization
//
//    * CreateVpcPeeringConnection
//
//    * DescribeVpcPeeringConnections
//
//    * DeleteVpcPeeringConnection
//
//    // Example sending a request using DeleteVpcPeeringAuthorizationRequest.
//    req := client.DeleteVpcPeeringAuthorizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DeleteVpcPeeringAuthorization
func (c *Client) DeleteVpcPeeringAuthorizationRequest(input *DeleteVpcPeeringAuthorizationInput) DeleteVpcPeeringAuthorizationRequest {
	op := &aws.Operation{
		Name:       opDeleteVpcPeeringAuthorization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteVpcPeeringAuthorizationInput{}
	}

	req := c.newRequest(op, input, &DeleteVpcPeeringAuthorizationOutput{})
	return DeleteVpcPeeringAuthorizationRequest{Request: req, Input: input, Copy: c.DeleteVpcPeeringAuthorizationRequest}
}

// DeleteVpcPeeringAuthorizationRequest is the request type for the
// DeleteVpcPeeringAuthorization API operation.
type DeleteVpcPeeringAuthorizationRequest struct {
	*aws.Request
	Input *DeleteVpcPeeringAuthorizationInput
	Copy  func(*DeleteVpcPeeringAuthorizationInput) DeleteVpcPeeringAuthorizationRequest
}

// Send marshals and sends the DeleteVpcPeeringAuthorization API request.
func (r DeleteVpcPeeringAuthorizationRequest) Send(ctx context.Context) (*DeleteVpcPeeringAuthorizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVpcPeeringAuthorizationResponse{
		DeleteVpcPeeringAuthorizationOutput: r.Request.Data.(*DeleteVpcPeeringAuthorizationOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVpcPeeringAuthorizationResponse is the response type for the
// DeleteVpcPeeringAuthorization API operation.
type DeleteVpcPeeringAuthorizationResponse struct {
	*DeleteVpcPeeringAuthorizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVpcPeeringAuthorization request.
func (r *DeleteVpcPeeringAuthorizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
