// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeSubnetsRequest
type DescribeSubnetsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * availability-zone - The Availability Zone for the subnet. You can also
	//    use availabilityZone as the filter name.
	//
	//    * availability-zone-id - The ID of the Availability Zone for the subnet.
	//    You can also use availabilityZoneId as the filter name.
	//
	//    * available-ip-address-count - The number of IPv4 addresses in the subnet
	//    that are available.
	//
	//    * cidr-block - The IPv4 CIDR block of the subnet. The CIDR block you specify
	//    must exactly match the subnet's CIDR block for information to be returned
	//    for the subnet. You can also use cidr or cidrBlock as the filter names.
	//
	//    * default-for-az - Indicates whether this is the default subnet for the
	//    Availability Zone. You can also use defaultForAz as the filter name.
	//
	//    * ipv6-cidr-block-association.ipv6-cidr-block - An IPv6 CIDR block associated
	//    with the subnet.
	//
	//    * ipv6-cidr-block-association.association-id - An association ID for an
	//    IPv6 CIDR block associated with the subnet.
	//
	//    * ipv6-cidr-block-association.state - The state of an IPv6 CIDR block
	//    associated with the subnet.
	//
	//    * owner-id - The ID of the AWS account that owns the subnet.
	//
	//    * state - The state of the subnet (pending | available).
	//
	//    * subnet-arn - The Amazon Resource Name (ARN) of the subnet.
	//
	//    * subnet-id - The ID of the subnet.
	//
	//    * tag:<key> - The key/value combination of a tag assigned to the resource.
	//    Use the tag key in the filter name and the tag value as the filter value.
	//    For example, to find all resources that have a tag with the key Owner
	//    and the value TeamA, specify tag:Owner for the filter name and TeamA for
	//    the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. Use this filter
	//    to find all resources assigned a tag with a specific key, regardless of
	//    the tag value.
	//
	//    * vpc-id - The ID of the VPC for the subnet.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`

	// One or more subnet IDs.
	//
	// Default: Describes all your subnets.
	SubnetIds []string `locationName:"SubnetId" locationNameList:"SubnetId" type:"list"`
}

// String returns the string representation
func (s DescribeSubnetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSubnetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSubnetsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeSubnetsResult
type DescribeSubnetsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about one or more subnets.
	Subnets []Subnet `locationName:"subnetSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeSubnetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSubnets = "DescribeSubnets"

// DescribeSubnetsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more of your subnets.
//
// For more information, see Your VPC and Subnets (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using DescribeSubnetsRequest.
//    req := client.DescribeSubnetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeSubnets
func (c *Client) DescribeSubnetsRequest(input *DescribeSubnetsInput) DescribeSubnetsRequest {
	op := &aws.Operation{
		Name:       opDescribeSubnets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeSubnetsInput{}
	}

	req := c.newRequest(op, input, &DescribeSubnetsOutput{})
	return DescribeSubnetsRequest{Request: req, Input: input, Copy: c.DescribeSubnetsRequest}
}

// DescribeSubnetsRequest is the request type for the
// DescribeSubnets API operation.
type DescribeSubnetsRequest struct {
	*aws.Request
	Input *DescribeSubnetsInput
	Copy  func(*DescribeSubnetsInput) DescribeSubnetsRequest
}

// Send marshals and sends the DescribeSubnets API request.
func (r DescribeSubnetsRequest) Send(ctx context.Context) (*DescribeSubnetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSubnetsResponse{
		DescribeSubnetsOutput: r.Request.Data.(*DescribeSubnetsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeSubnetsRequestPaginator returns a paginator for DescribeSubnets.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeSubnetsRequest(input)
//   p := ec2.NewDescribeSubnetsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeSubnetsPaginator(req DescribeSubnetsRequest) DescribeSubnetsPaginator {
	return DescribeSubnetsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeSubnetsInput
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

// DescribeSubnetsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeSubnetsPaginator struct {
	aws.Pager
}

func (p *DescribeSubnetsPaginator) CurrentPage() *DescribeSubnetsOutput {
	return p.Pager.CurrentPage().(*DescribeSubnetsOutput)
}

// DescribeSubnetsResponse is the response type for the
// DescribeSubnets API operation.
type DescribeSubnetsResponse struct {
	*DescribeSubnetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSubnets request.
func (r *DescribeSubnetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
