// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to perform a predictive inbox placement test. Predictive inbox
// placement tests can help you predict how your messages will be handled by
// various email providers around the world. When you perform a predictive inbox
// placement test, you provide a sample message that contains the content that
// you plan to send to your customers. Amazon Pinpoint then sends that message
// to special email addresses spread across several major email providers. After
// about 24 hours, the test is complete, and you can use the GetDeliverabilityTestReport
// operation to view the results of the test.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/CreateDeliverabilityTestReportRequest
type CreateDeliverabilityTestReportInput struct {
	_ struct{} `type:"structure"`

	// The HTML body of the message that you sent when you performed the predictive
	// inbox placement test.
	//
	// Content is a required field
	Content *EmailContent `type:"structure" required:"true"`

	// The email address that the predictive inbox placement test email was sent
	// from.
	//
	// FromEmailAddress is a required field
	FromEmailAddress *string `type:"string" required:"true"`

	// A unique name that helps you to identify the predictive inbox placement test
	// when you retrieve the results.
	ReportName *string `type:"string"`

	// An array of objects that define the tags (keys and values) that you want
	// to associate with the predictive inbox placement test.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateDeliverabilityTestReportInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDeliverabilityTestReportInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDeliverabilityTestReportInput"}

	if s.Content == nil {
		invalidParams.Add(aws.NewErrParamRequired("Content"))
	}

	if s.FromEmailAddress == nil {
		invalidParams.Add(aws.NewErrParamRequired("FromEmailAddress"))
	}
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			invalidParams.AddNested("Content", err.(aws.ErrInvalidParams))
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
func (s CreateDeliverabilityTestReportInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Content != nil {
		v := s.Content

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Content", v, metadata)
	}
	if s.FromEmailAddress != nil {
		v := *s.FromEmailAddress

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FromEmailAddress", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ReportName != nil {
		v := *s.ReportName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ReportName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// Information about the predictive inbox placement test that you created.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/CreateDeliverabilityTestReportResponse
type CreateDeliverabilityTestReportOutput struct {
	_ struct{} `type:"structure"`

	// The status of the predictive inbox placement test. If the status is IN_PROGRESS,
	// then the predictive inbox placement test is currently running. Predictive
	// inbox placement tests are usually complete within 24 hours of creating the
	// test. If the status is COMPLETE, then the test is finished, and you can use
	// the GetDeliverabilityTestReport to view the results of the test.
	//
	// DeliverabilityTestStatus is a required field
	DeliverabilityTestStatus DeliverabilityTestStatus `type:"string" required:"true" enum:"true"`

	// A unique string that identifies the predictive inbox placement test.
	//
	// ReportId is a required field
	ReportId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateDeliverabilityTestReportOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDeliverabilityTestReportOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.DeliverabilityTestStatus) > 0 {
		v := s.DeliverabilityTestStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeliverabilityTestStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ReportId != nil {
		v := *s.ReportId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ReportId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateDeliverabilityTestReport = "CreateDeliverabilityTestReport"

// CreateDeliverabilityTestReportRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Create a new predictive inbox placement test. Predictive inbox placement
// tests can help you predict how your messages will be handled by various email
// providers around the world. When you perform a predictive inbox placement
// test, you provide a sample message that contains the content that you plan
// to send to your customers. Amazon Pinpoint then sends that message to special
// email addresses spread across several major email providers. After about
// 24 hours, the test is complete, and you can use the GetDeliverabilityTestReport
// operation to view the results of the test.
//
//    // Example sending a request using CreateDeliverabilityTestReportRequest.
//    req := client.CreateDeliverabilityTestReportRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/CreateDeliverabilityTestReport
func (c *Client) CreateDeliverabilityTestReportRequest(input *CreateDeliverabilityTestReportInput) CreateDeliverabilityTestReportRequest {
	op := &aws.Operation{
		Name:       opCreateDeliverabilityTestReport,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/email/deliverability-dashboard/test",
	}

	if input == nil {
		input = &CreateDeliverabilityTestReportInput{}
	}

	req := c.newRequest(op, input, &CreateDeliverabilityTestReportOutput{})
	return CreateDeliverabilityTestReportRequest{Request: req, Input: input, Copy: c.CreateDeliverabilityTestReportRequest}
}

// CreateDeliverabilityTestReportRequest is the request type for the
// CreateDeliverabilityTestReport API operation.
type CreateDeliverabilityTestReportRequest struct {
	*aws.Request
	Input *CreateDeliverabilityTestReportInput
	Copy  func(*CreateDeliverabilityTestReportInput) CreateDeliverabilityTestReportRequest
}

// Send marshals and sends the CreateDeliverabilityTestReport API request.
func (r CreateDeliverabilityTestReportRequest) Send(ctx context.Context) (*CreateDeliverabilityTestReportResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDeliverabilityTestReportResponse{
		CreateDeliverabilityTestReportOutput: r.Request.Data.(*CreateDeliverabilityTestReportOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDeliverabilityTestReportResponse is the response type for the
// CreateDeliverabilityTestReport API operation.
type CreateDeliverabilityTestReportResponse struct {
	*CreateDeliverabilityTestReportOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDeliverabilityTestReport request.
func (r *CreateDeliverabilityTestReportResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
