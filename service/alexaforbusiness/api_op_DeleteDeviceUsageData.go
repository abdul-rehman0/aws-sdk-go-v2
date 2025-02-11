// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DeleteDeviceUsageDataRequest
type DeleteDeviceUsageDataInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the device.
	//
	// DeviceArn is a required field
	DeviceArn *string `type:"string" required:"true"`

	// The type of usage data to delete.
	//
	// DeviceUsageType is a required field
	DeviceUsageType DeviceUsageType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s DeleteDeviceUsageDataInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDeviceUsageDataInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDeviceUsageDataInput"}

	if s.DeviceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceArn"))
	}
	if len(s.DeviceUsageType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("DeviceUsageType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DeleteDeviceUsageDataResponse
type DeleteDeviceUsageDataOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDeviceUsageDataOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDeviceUsageData = "DeleteDeviceUsageData"

// DeleteDeviceUsageDataRequest returns a request value for making API operation for
// Alexa For Business.
//
// When this action is called for a specified shared device, it allows authorized
// users to delete the device's entire previous history of voice input data
// and associated response data. This action can be called once every 24 hours
// for a specific shared device.
//
// When this action is called for a specified shared device, it allows authorized
// users to delete the device's entire previous history of voice input data.
// This action can be called once every 24 hours for a specific shared device.
//
//    // Example sending a request using DeleteDeviceUsageDataRequest.
//    req := client.DeleteDeviceUsageDataRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DeleteDeviceUsageData
func (c *Client) DeleteDeviceUsageDataRequest(input *DeleteDeviceUsageDataInput) DeleteDeviceUsageDataRequest {
	op := &aws.Operation{
		Name:       opDeleteDeviceUsageData,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDeviceUsageDataInput{}
	}

	req := c.newRequest(op, input, &DeleteDeviceUsageDataOutput{})
	return DeleteDeviceUsageDataRequest{Request: req, Input: input, Copy: c.DeleteDeviceUsageDataRequest}
}

// DeleteDeviceUsageDataRequest is the request type for the
// DeleteDeviceUsageData API operation.
type DeleteDeviceUsageDataRequest struct {
	*aws.Request
	Input *DeleteDeviceUsageDataInput
	Copy  func(*DeleteDeviceUsageDataInput) DeleteDeviceUsageDataRequest
}

// Send marshals and sends the DeleteDeviceUsageData API request.
func (r DeleteDeviceUsageDataRequest) Send(ctx context.Context) (*DeleteDeviceUsageDataResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDeviceUsageDataResponse{
		DeleteDeviceUsageDataOutput: r.Request.Data.(*DeleteDeviceUsageDataOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDeviceUsageDataResponse is the response type for the
// DeleteDeviceUsageData API operation.
type DeleteDeviceUsageDataResponse struct {
	*DeleteDeviceUsageDataOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDeviceUsageData request.
func (r *DeleteDeviceUsageDataResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
