// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acmpca

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/RestoreCertificateAuthorityRequest
type RestoreCertificateAuthorityInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that was returned when you called the CreateCertificateAuthority
	// action. This must be of the form:
	//
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	//
	// CertificateAuthorityArn is a required field
	CertificateAuthorityArn *string `min:"5" type:"string" required:"true"`
}

// String returns the string representation
func (s RestoreCertificateAuthorityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreCertificateAuthorityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestoreCertificateAuthorityInput"}

	if s.CertificateAuthorityArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if s.CertificateAuthorityArn != nil && len(*s.CertificateAuthorityArn) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateAuthorityArn", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/RestoreCertificateAuthorityOutput
type RestoreCertificateAuthorityOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RestoreCertificateAuthorityOutput) String() string {
	return awsutil.Prettify(s)
}

const opRestoreCertificateAuthority = "RestoreCertificateAuthority"

// RestoreCertificateAuthorityRequest returns a request value for making API operation for
// AWS Certificate Manager Private Certificate Authority.
//
// Restores a certificate authority (CA) that is in the DELETED state. You can
// restore a CA during the period that you defined in the PermanentDeletionTimeInDays
// parameter of the DeleteCertificateAuthority action. Currently, you can specify
// 7 to 30 days. If you did not specify a PermanentDeletionTimeInDays value,
// by default you can restore the CA at any time in a 30 day period. You can
// check the time remaining in the restoration period of a private CA in the
// DELETED state by calling the DescribeCertificateAuthority or ListCertificateAuthorities
// actions. The status of a restored CA is set to its pre-deletion status when
// the RestoreCertificateAuthority action returns. To change its status to ACTIVE,
// call the UpdateCertificateAuthority action. If the private CA was in the
// PENDING_CERTIFICATE state at deletion, you must use the ImportCertificateAuthorityCertificate
// action to import a certificate authority into the private CA before it can
// be activated. You cannot restore a CA after the restoration period has ended.
//
//    // Example sending a request using RestoreCertificateAuthorityRequest.
//    req := client.RestoreCertificateAuthorityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/RestoreCertificateAuthority
func (c *Client) RestoreCertificateAuthorityRequest(input *RestoreCertificateAuthorityInput) RestoreCertificateAuthorityRequest {
	op := &aws.Operation{
		Name:       opRestoreCertificateAuthority,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestoreCertificateAuthorityInput{}
	}

	req := c.newRequest(op, input, &RestoreCertificateAuthorityOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RestoreCertificateAuthorityRequest{Request: req, Input: input, Copy: c.RestoreCertificateAuthorityRequest}
}

// RestoreCertificateAuthorityRequest is the request type for the
// RestoreCertificateAuthority API operation.
type RestoreCertificateAuthorityRequest struct {
	*aws.Request
	Input *RestoreCertificateAuthorityInput
	Copy  func(*RestoreCertificateAuthorityInput) RestoreCertificateAuthorityRequest
}

// Send marshals and sends the RestoreCertificateAuthority API request.
func (r RestoreCertificateAuthorityRequest) Send(ctx context.Context) (*RestoreCertificateAuthorityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreCertificateAuthorityResponse{
		RestoreCertificateAuthorityOutput: r.Request.Data.(*RestoreCertificateAuthorityOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreCertificateAuthorityResponse is the response type for the
// RestoreCertificateAuthority API operation.
type RestoreCertificateAuthorityResponse struct {
	*RestoreCertificateAuthorityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreCertificateAuthority request.
func (r *RestoreCertificateAuthorityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
