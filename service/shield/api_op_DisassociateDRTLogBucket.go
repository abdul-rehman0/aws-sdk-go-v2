// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/DisassociateDRTLogBucketRequest
type DisassociateDRTLogBucketInput struct {
	_ struct{} `type:"structure"`

	// The Amazon S3 bucket that contains your AWS WAF logs.
	//
	// LogBucket is a required field
	LogBucket *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateDRTLogBucketInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateDRTLogBucketInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateDRTLogBucketInput"}

	if s.LogBucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogBucket"))
	}
	if s.LogBucket != nil && len(*s.LogBucket) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("LogBucket", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/DisassociateDRTLogBucketResponse
type DisassociateDRTLogBucketOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateDRTLogBucketOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateDRTLogBucket = "DisassociateDRTLogBucket"

// DisassociateDRTLogBucketRequest returns a request value for making API operation for
// AWS Shield.
//
// Removes the DDoS Response team's (DRT) access to the specified Amazon S3
// bucket containing your AWS WAF logs.
//
// To make a DisassociateDRTLogBucket request, you must be subscribed to the
// Business Support plan (https://aws.amazon.com/premiumsupport/business-support/)
// or the Enterprise Support plan (https://aws.amazon.com/premiumsupport/enterprise-support/).
// However, if you are not subscribed to one of these support plans, but had
// been previously and had granted the DRT access to your account, you can submit
// a DisassociateDRTLogBucket request to remove this access.
//
//    // Example sending a request using DisassociateDRTLogBucketRequest.
//    req := client.DisassociateDRTLogBucketRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/DisassociateDRTLogBucket
func (c *Client) DisassociateDRTLogBucketRequest(input *DisassociateDRTLogBucketInput) DisassociateDRTLogBucketRequest {
	op := &aws.Operation{
		Name:       opDisassociateDRTLogBucket,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateDRTLogBucketInput{}
	}

	req := c.newRequest(op, input, &DisassociateDRTLogBucketOutput{})
	return DisassociateDRTLogBucketRequest{Request: req, Input: input, Copy: c.DisassociateDRTLogBucketRequest}
}

// DisassociateDRTLogBucketRequest is the request type for the
// DisassociateDRTLogBucket API operation.
type DisassociateDRTLogBucketRequest struct {
	*aws.Request
	Input *DisassociateDRTLogBucketInput
	Copy  func(*DisassociateDRTLogBucketInput) DisassociateDRTLogBucketRequest
}

// Send marshals and sends the DisassociateDRTLogBucket API request.
func (r DisassociateDRTLogBucketRequest) Send(ctx context.Context) (*DisassociateDRTLogBucketResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateDRTLogBucketResponse{
		DisassociateDRTLogBucketOutput: r.Request.Data.(*DisassociateDRTLogBucketOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateDRTLogBucketResponse is the response type for the
// DisassociateDRTLogBucket API operation.
type DisassociateDRTLogBucketResponse struct {
	*DisassociateDRTLogBucketOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateDRTLogBucket request.
func (r *DisassociateDRTLogBucketResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
