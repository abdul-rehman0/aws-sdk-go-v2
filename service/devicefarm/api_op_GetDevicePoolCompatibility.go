// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to the get device pool compatibility operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/GetDevicePoolCompatibilityRequest
type GetDevicePoolCompatibilityInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the app that is associated with the specified device pool.
	AppArn *string `locationName:"appArn" min:"32" type:"string"`

	// An object containing information about the settings for a run.
	Configuration *ScheduleRunConfiguration `locationName:"configuration" type:"structure"`

	// The device pool's ARN.
	//
	// DevicePoolArn is a required field
	DevicePoolArn *string `locationName:"devicePoolArn" min:"32" type:"string" required:"true"`

	// Information about the uploaded test to be run against the device pool.
	Test *ScheduleRunTest `locationName:"test" type:"structure"`

	// The test type for the specified device pool.
	//
	// Allowed values include the following:
	//
	//    * BUILTIN_FUZZ: The built-in fuzz type.
	//
	//    * BUILTIN_EXPLORER: For Android, an app explorer that will traverse an
	//    Android app, interacting with it and capturing screenshots at the same
	//    time.
	//
	//    * APPIUM_JAVA_JUNIT: The Appium Java JUnit type.
	//
	//    * APPIUM_JAVA_TESTNG: The Appium Java TestNG type.
	//
	//    * APPIUM_PYTHON: The Appium Python type.
	//
	//    * APPIUM_NODE: The Appium Node.js type.
	//
	//    * APPIUM_RUBY: The Appium Ruby type.
	//
	//    * APPIUM_WEB_JAVA_JUNIT: The Appium Java JUnit type for web apps.
	//
	//    * APPIUM_WEB_JAVA_TESTNG: The Appium Java TestNG type for web apps.
	//
	//    * APPIUM_WEB_PYTHON: The Appium Python type for web apps.
	//
	//    * APPIUM_WEB_NODE: The Appium Node.js type for web apps.
	//
	//    * APPIUM_WEB_RUBY: The Appium Ruby type for web apps.
	//
	//    * CALABASH: The Calabash type.
	//
	//    * INSTRUMENTATION: The Instrumentation type.
	//
	//    * UIAUTOMATION: The uiautomation type.
	//
	//    * UIAUTOMATOR: The uiautomator type.
	//
	//    * XCTEST: The Xcode test type.
	//
	//    * XCTEST_UI: The Xcode UI test type.
	TestType TestType `locationName:"testType" type:"string" enum:"true"`
}

// String returns the string representation
func (s GetDevicePoolCompatibilityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDevicePoolCompatibilityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDevicePoolCompatibilityInput"}
	if s.AppArn != nil && len(*s.AppArn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("AppArn", 32))
	}

	if s.DevicePoolArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DevicePoolArn"))
	}
	if s.DevicePoolArn != nil && len(*s.DevicePoolArn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("DevicePoolArn", 32))
	}
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			invalidParams.AddNested("Configuration", err.(aws.ErrInvalidParams))
		}
	}
	if s.Test != nil {
		if err := s.Test.Validate(); err != nil {
			invalidParams.AddNested("Test", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the result of describe device pool compatibility request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/GetDevicePoolCompatibilityResult
type GetDevicePoolCompatibilityOutput struct {
	_ struct{} `type:"structure"`

	// Information about compatible devices.
	CompatibleDevices []DevicePoolCompatibilityResult `locationName:"compatibleDevices" type:"list"`

	// Information about incompatible devices.
	IncompatibleDevices []DevicePoolCompatibilityResult `locationName:"incompatibleDevices" type:"list"`
}

// String returns the string representation
func (s GetDevicePoolCompatibilityOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDevicePoolCompatibility = "GetDevicePoolCompatibility"

// GetDevicePoolCompatibilityRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Gets information about compatibility with a device pool.
//
//    // Example sending a request using GetDevicePoolCompatibilityRequest.
//    req := client.GetDevicePoolCompatibilityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/GetDevicePoolCompatibility
func (c *Client) GetDevicePoolCompatibilityRequest(input *GetDevicePoolCompatibilityInput) GetDevicePoolCompatibilityRequest {
	op := &aws.Operation{
		Name:       opGetDevicePoolCompatibility,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDevicePoolCompatibilityInput{}
	}

	req := c.newRequest(op, input, &GetDevicePoolCompatibilityOutput{})
	return GetDevicePoolCompatibilityRequest{Request: req, Input: input, Copy: c.GetDevicePoolCompatibilityRequest}
}

// GetDevicePoolCompatibilityRequest is the request type for the
// GetDevicePoolCompatibility API operation.
type GetDevicePoolCompatibilityRequest struct {
	*aws.Request
	Input *GetDevicePoolCompatibilityInput
	Copy  func(*GetDevicePoolCompatibilityInput) GetDevicePoolCompatibilityRequest
}

// Send marshals and sends the GetDevicePoolCompatibility API request.
func (r GetDevicePoolCompatibilityRequest) Send(ctx context.Context) (*GetDevicePoolCompatibilityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDevicePoolCompatibilityResponse{
		GetDevicePoolCompatibilityOutput: r.Request.Data.(*GetDevicePoolCompatibilityOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDevicePoolCompatibilityResponse is the response type for the
// GetDevicePoolCompatibility API operation.
type GetDevicePoolCompatibilityResponse struct {
	*GetDevicePoolCompatibilityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDevicePoolCompatibility request.
func (r *GetDevicePoolCompatibilityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
