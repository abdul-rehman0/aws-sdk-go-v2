// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/DescribeScheduledActionsType
type DescribeScheduledActionsInput struct {
	_ struct{} `type:"structure"`

	// The name of the Auto Scaling group.
	AutoScalingGroupName *string `min:"1" type:"string"`

	// The latest scheduled start time to return. If scheduled action names are
	// provided, this parameter is ignored.
	EndTime *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The maximum number of items to return with this call. The default value is
	// 50 and the maximum value is 100.
	MaxRecords *int64 `type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`

	// The names of one or more scheduled actions. You can specify up to 50 actions.
	// If you omit this parameter, all scheduled actions are described. If you specify
	// an unknown scheduled action, it is ignored with no error.
	ScheduledActionNames []string `type:"list"`

	// The earliest scheduled start time to return. If scheduled action names are
	// provided, this parameter is ignored.
	StartTime *time.Time `type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s DescribeScheduledActionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeScheduledActionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeScheduledActionsInput"}
	if s.AutoScalingGroupName != nil && len(*s.AutoScalingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoScalingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/ScheduledActionsType
type DescribeScheduledActionsOutput struct {
	_ struct{} `type:"structure"`

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this
	// string for the NextToken value when requesting the next set of items. This
	// value is null when there are no more items to return.
	NextToken *string `type:"string"`

	// The scheduled actions.
	ScheduledUpdateGroupActions []ScheduledUpdateGroupAction `type:"list"`
}

// String returns the string representation
func (s DescribeScheduledActionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeScheduledActions = "DescribeScheduledActions"

// DescribeScheduledActionsRequest returns a request value for making API operation for
// Auto Scaling.
//
// Describes the actions scheduled for your Auto Scaling group that haven't
// run or that have not reached their end time. To describe the actions that
// have already run, use DescribeScalingActivities.
//
//    // Example sending a request using DescribeScheduledActionsRequest.
//    req := client.DescribeScheduledActionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/DescribeScheduledActions
func (c *Client) DescribeScheduledActionsRequest(input *DescribeScheduledActionsInput) DescribeScheduledActionsRequest {
	op := &aws.Operation{
		Name:       opDescribeScheduledActions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeScheduledActionsInput{}
	}

	req := c.newRequest(op, input, &DescribeScheduledActionsOutput{})
	return DescribeScheduledActionsRequest{Request: req, Input: input, Copy: c.DescribeScheduledActionsRequest}
}

// DescribeScheduledActionsRequest is the request type for the
// DescribeScheduledActions API operation.
type DescribeScheduledActionsRequest struct {
	*aws.Request
	Input *DescribeScheduledActionsInput
	Copy  func(*DescribeScheduledActionsInput) DescribeScheduledActionsRequest
}

// Send marshals and sends the DescribeScheduledActions API request.
func (r DescribeScheduledActionsRequest) Send(ctx context.Context) (*DescribeScheduledActionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeScheduledActionsResponse{
		DescribeScheduledActionsOutput: r.Request.Data.(*DescribeScheduledActionsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeScheduledActionsRequestPaginator returns a paginator for DescribeScheduledActions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeScheduledActionsRequest(input)
//   p := autoscaling.NewDescribeScheduledActionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeScheduledActionsPaginator(req DescribeScheduledActionsRequest) DescribeScheduledActionsPaginator {
	return DescribeScheduledActionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeScheduledActionsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeScheduledActionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeScheduledActionsPaginator struct {
	aws.Pager
}

func (p *DescribeScheduledActionsPaginator) CurrentPage() *DescribeScheduledActionsOutput {
	return p.Pager.CurrentPage().(*DescribeScheduledActionsOutput)
}

// DescribeScheduledActionsResponse is the response type for the
// DescribeScheduledActions API operation.
type DescribeScheduledActionsResponse struct {
	*DescribeScheduledActionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeScheduledActions request.
func (r *DescribeScheduledActionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
