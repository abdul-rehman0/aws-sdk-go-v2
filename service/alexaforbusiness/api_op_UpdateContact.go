// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/UpdateContactRequest
type UpdateContactInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the contact to update.
	//
	// ContactArn is a required field
	ContactArn *string `type:"string" required:"true"`

	// The updated display name of the contact.
	DisplayName *string `min:"1" type:"string"`

	// The updated first name of the contact.
	FirstName *string `min:"1" type:"string"`

	// The updated last name of the contact.
	LastName *string `min:"1" type:"string"`

	// The updated phone number of the contact. The phone number type defaults to
	// WORK. You can either specify PhoneNumber or PhoneNumbers. We recommend that
	// you use PhoneNumbers, which lets you specify the phone number type and multiple
	// numbers.
	PhoneNumber *string `type:"string"`

	// The list of phone numbers for the contact.
	PhoneNumbers []PhoneNumber `type:"list"`

	// The list of SIP addresses for the contact.
	SipAddresses []SipAddress `type:"list"`
}

// String returns the string representation
func (s UpdateContactInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateContactInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateContactInput"}

	if s.ContactArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContactArn"))
	}
	if s.DisplayName != nil && len(*s.DisplayName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DisplayName", 1))
	}
	if s.FirstName != nil && len(*s.FirstName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FirstName", 1))
	}
	if s.LastName != nil && len(*s.LastName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LastName", 1))
	}
	if s.PhoneNumbers != nil {
		for i, v := range s.PhoneNumbers {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "PhoneNumbers", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SipAddresses != nil {
		for i, v := range s.SipAddresses {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SipAddresses", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/UpdateContactResponse
type UpdateContactOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateContactOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateContact = "UpdateContact"

// UpdateContactRequest returns a request value for making API operation for
// Alexa For Business.
//
// Updates the contact details by the contact ARN.
//
//    // Example sending a request using UpdateContactRequest.
//    req := client.UpdateContactRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/UpdateContact
func (c *Client) UpdateContactRequest(input *UpdateContactInput) UpdateContactRequest {
	op := &aws.Operation{
		Name:       opUpdateContact,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateContactInput{}
	}

	req := c.newRequest(op, input, &UpdateContactOutput{})
	return UpdateContactRequest{Request: req, Input: input, Copy: c.UpdateContactRequest}
}

// UpdateContactRequest is the request type for the
// UpdateContact API operation.
type UpdateContactRequest struct {
	*aws.Request
	Input *UpdateContactInput
	Copy  func(*UpdateContactInput) UpdateContactRequest
}

// Send marshals and sends the UpdateContact API request.
func (r UpdateContactRequest) Send(ctx context.Context) (*UpdateContactResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateContactResponse{
		UpdateContactOutput: r.Request.Data.(*UpdateContactOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateContactResponse is the response type for the
// UpdateContact API operation.
type UpdateContactResponse struct {
	*UpdateContactOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateContact request.
func (r *UpdateContactResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
