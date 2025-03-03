// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A complex type that contains information about a request to update a health
// check.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/UpdateHealthCheckRequest
type UpdateHealthCheckInput struct {
	_ struct{} `locationName:"UpdateHealthCheckRequest" type:"structure" xmlURI:"https://route53.amazonaws.com/doc/2013-04-01/"`

	// A complex type that identifies the CloudWatch alarm that you want Amazon
	// Route 53 health checkers to use to determine whether the specified health
	// check is healthy.
	AlarmIdentifier *AlarmIdentifier `type:"structure"`

	// A complex type that contains one ChildHealthCheck element for each health
	// check that you want to associate with a CALCULATED health check.
	ChildHealthChecks []string `locationNameList:"ChildHealthCheck" type:"list"`

	// Stops Route 53 from performing health checks. When you disable a health check,
	// here's what happens:
	//
	//    * Health checks that check the health of endpoints: Route 53 stops submitting
	//    requests to your application, server, or other resource.
	//
	//    * Calculated health checks: Route 53 stops aggregating the status of the
	//    referenced health checks.
	//
	//    * Health checks that monitor CloudWatch alarms: Route 53 stops monitoring
	//    the corresponding CloudWatch metrics.
	//
	// After you disable a health check, Route 53 considers the status of the health
	// check to always be healthy. If you configured DNS failover, Route 53 continues
	// to route traffic to the corresponding resources. If you want to stop routing
	// traffic to a resource, change the value of Inverted (https://docs.aws.amazon.com/Route53/latest/APIReference/API_UpdateHealthCheck.html#Route53-UpdateHealthCheck-request-Inverted).
	//
	// Charges for a health check still apply when the health check is disabled.
	// For more information, see Amazon Route 53 Pricing (http://aws.amazon.com/route53/pricing/).
	Disabled *bool `type:"boolean"`

	// Specify whether you want Amazon Route 53 to send the value of FullyQualifiedDomainName
	// to the endpoint in the client_hello message during TLS negotiation. This
	// allows the endpoint to respond to HTTPS health check requests with the applicable
	// SSL/TLS certificate.
	//
	// Some endpoints require that HTTPS requests include the host name in the client_hello
	// message. If you don't enable SNI, the status of the health check will be
	// SSL alert handshake_failure. A health check can also have that status for
	// other reasons. If SNI is enabled and you're still getting the error, check
	// the SSL/TLS configuration on your endpoint and confirm that your certificate
	// is valid.
	//
	// The SSL/TLS certificate on your endpoint includes a domain name in the Common
	// Name field and possibly several more in the Subject Alternative Names field.
	// One of the domain names in the certificate should match the value that you
	// specify for FullyQualifiedDomainName. If the endpoint responds to the client_hello
	// message with a certificate that does not include the domain name that you
	// specified in FullyQualifiedDomainName, a health checker will retry the handshake.
	// In the second attempt, the health checker will omit FullyQualifiedDomainName
	// from the client_hello message.
	EnableSNI *bool `type:"boolean"`

	// The number of consecutive health checks that an endpoint must pass or fail
	// for Amazon Route 53 to change the current status of the endpoint from unhealthy
	// to healthy or vice versa. For more information, see How Amazon Route 53 Determines
	// Whether an Endpoint Is Healthy (http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-determining-health-of-endpoints.html)
	// in the Amazon Route 53 Developer Guide.
	//
	// If you don't specify a value for FailureThreshold, the default value is three
	// health checks.
	FailureThreshold *int64 `min:"1" type:"integer"`

	// Amazon Route 53 behavior depends on whether you specify a value for IPAddress.
	//
	// If a health check already has a value for IPAddress, you can change the value.
	// However, you can't update an existing health check to add or remove the value
	// of IPAddress.
	//
	// If you specify a value for IPAddress:
	//
	// Route 53 sends health check requests to the specified IPv4 or IPv6 address
	// and passes the value of FullyQualifiedDomainName in the Host header for all
	// health checks except TCP health checks. This is typically the fully qualified
	// DNS name of the endpoint on which you want Route 53 to perform health checks.
	//
	// When Route 53 checks the health of an endpoint, here is how it constructs
	// the Host header:
	//
	//    * If you specify a value of 80 for Port and HTTP or HTTP_STR_MATCH for
	//    Type, Route 53 passes the value of FullyQualifiedDomainName to the endpoint
	//    in the Host header.
	//
	//    * If you specify a value of 443 for Port and HTTPS or HTTPS_STR_MATCH
	//    for Type, Route 53 passes the value of FullyQualifiedDomainName to the
	//    endpoint in the Host header.
	//
	//    * If you specify another value for Port and any value except TCP for Type,
	//    Route 53 passes FullyQualifiedDomainName:Port to the endpoint in the Host
	//    header.
	//
	// If you don't specify a value for FullyQualifiedDomainName, Route 53 substitutes
	// the value of IPAddress in the Host header in each of the above cases.
	//
	// If you don't specify a value for IPAddress:
	//
	// If you don't specify a value for IPAddress, Route 53 sends a DNS request
	// to the domain that you specify in FullyQualifiedDomainName at the interval
	// you specify in RequestInterval. Using an IPv4 address that is returned by
	// DNS, Route 53 then checks the health of the endpoint.
	//
	// If you don't specify a value for IPAddress, Route 53 uses only IPv4 to send
	// health checks to the endpoint. If there's no resource record set with a type
	// of A for the name that you specify for FullyQualifiedDomainName, the health
	// check fails with a "DNS resolution failed" error.
	//
	// If you want to check the health of weighted, latency, or failover resource
	// record sets and you choose to specify the endpoint only by FullyQualifiedDomainName,
	// we recommend that you create a separate health check for each endpoint. For
	// example, create a health check for each HTTP server that is serving content
	// for www.example.com. For the value of FullyQualifiedDomainName, specify the
	// domain name of the server (such as us-east-2-www.example.com), not the name
	// of the resource record sets (www.example.com).
	//
	// In this configuration, if the value of FullyQualifiedDomainName matches the
	// name of the resource record sets and you then associate the health check
	// with those resource record sets, health check results will be unpredictable.
	//
	// In addition, if the value of Type is HTTP, HTTPS, HTTP_STR_MATCH, or HTTPS_STR_MATCH,
	// Route 53 passes the value of FullyQualifiedDomainName in the Host header,
	// as it does when you specify a value for IPAddress. If the value of Type is
	// TCP, Route 53 doesn't pass a Host header.
	FullyQualifiedDomainName *string `type:"string"`

	// The ID for the health check for which you want detailed information. When
	// you created the health check, CreateHealthCheck returned the ID in the response,
	// in the HealthCheckId element.
	//
	// HealthCheckId is a required field
	HealthCheckId *string `location:"uri" locationName:"HealthCheckId" type:"string" required:"true"`

	// A sequential counter that Amazon Route 53 sets to 1 when you create a health
	// check and increments by 1 each time you update settings for the health check.
	//
	// We recommend that you use GetHealthCheck or ListHealthChecks to get the current
	// value of HealthCheckVersion for the health check that you want to update,
	// and that you include that value in your UpdateHealthCheck request. This prevents
	// Route 53 from overwriting an intervening update:
	//
	//    * If the value in the UpdateHealthCheck request matches the value of HealthCheckVersion
	//    in the health check, Route 53 updates the health check with the new settings.
	//
	//    * If the value of HealthCheckVersion in the health check is greater, the
	//    health check was changed after you got the version number. Route 53 does
	//    not update the health check, and it returns a HealthCheckVersionMismatch
	//    error.
	HealthCheckVersion *int64 `min:"1" type:"long"`

	// The number of child health checks that are associated with a CALCULATED health
	// that Amazon Route 53 must consider healthy for the CALCULATED health check
	// to be considered healthy. To specify the child health checks that you want
	// to associate with a CALCULATED health check, use the ChildHealthChecks and
	// ChildHealthCheck elements.
	//
	// Note the following:
	//
	//    * If you specify a number greater than the number of child health checks,
	//    Route 53 always considers this health check to be unhealthy.
	//
	//    * If you specify 0, Route 53 always considers this health check to be
	//    healthy.
	HealthThreshold *int64 `type:"integer"`

	// The IPv4 or IPv6 IP address for the endpoint that you want Amazon Route 53
	// to perform health checks on. If you don't specify a value for IPAddress,
	// Route 53 sends a DNS request to resolve the domain name that you specify
	// in FullyQualifiedDomainName at the interval that you specify in RequestInterval.
	// Using an IP address that is returned by DNS, Route 53 then checks the health
	// of the endpoint.
	//
	// Use one of the following formats for the value of IPAddress:
	//
	//    * IPv4 address: four values between 0 and 255, separated by periods (.),
	//    for example, 192.0.2.44.
	//
	//    * IPv6 address: eight groups of four hexadecimal values, separated by
	//    colons (:), for example, 2001:0db8:85a3:0000:0000:abcd:0001:2345. You
	//    can also shorten IPv6 addresses as described in RFC 5952, for example,
	//    2001:db8:85a3::abcd:1:2345.
	//
	// If the endpoint is an EC2 instance, we recommend that you create an Elastic
	// IP address, associate it with your EC2 instance, and specify the Elastic
	// IP address for IPAddress. This ensures that the IP address of your instance
	// never changes. For more information, see the applicable documentation:
	//
	//    * Linux: Elastic IP Addresses (EIP) (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html)
	//    in the Amazon EC2 User Guide for Linux Instances
	//
	//    * Windows: Elastic IP Addresses (EIP) (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/elastic-ip-addresses-eip.html)
	//    in the Amazon EC2 User Guide for Windows Instances
	//
	// If a health check already has a value for IPAddress, you can change the value.
	// However, you can't update an existing health check to add or remove the value
	// of IPAddress.
	//
	// For more information, see FullyQualifiedDomainName (https://docs.aws.amazon.com/Route53/latest/APIReference/API_UpdateHealthCheck.html#Route53-UpdateHealthCheck-request-FullyQualifiedDomainName).
	//
	// Constraints: Route 53 can't check the health of endpoints for which the IP
	// address is in local, private, non-routable, or multicast ranges. For more
	// information about IP addresses for which you can't create health checks,
	// see the following documents:
	//
	//    * RFC 5735, Special Use IPv4 Addresses (https://tools.ietf.org/html/rfc5735)
	//
	//    * RFC 6598, IANA-Reserved IPv4 Prefix for Shared Address Space (https://tools.ietf.org/html/rfc6598)
	//
	//    * RFC 5156, Special-Use IPv6 Addresses (https://tools.ietf.org/html/rfc5156)
	IPAddress *string `type:"string"`

	// When CloudWatch has insufficient data about the metric to determine the alarm
	// state, the status that you want Amazon Route 53 to assign to the health check:
	//
	//    * Healthy: Route 53 considers the health check to be healthy.
	//
	//    * Unhealthy: Route 53 considers the health check to be unhealthy.
	//
	//    * LastKnownStatus: Route 53 uses the status of the health check from the
	//    last time CloudWatch had sufficient data to determine the alarm state.
	//    For new health checks that have no last known status, the default status
	//    for the health check is healthy.
	InsufficientDataHealthStatus InsufficientDataHealthStatus `type:"string" enum:"true"`

	// Specify whether you want Amazon Route 53 to invert the status of a health
	// check, for example, to consider a health check unhealthy when it otherwise
	// would be considered healthy.
	Inverted *bool `type:"boolean"`

	// The port on the endpoint on which you want Amazon Route 53 to perform health
	// checks.
	Port *int64 `min:"1" type:"integer"`

	// A complex type that contains one Region element for each region that you
	// want Amazon Route 53 health checkers to check the specified endpoint from.
	Regions []HealthCheckRegion `locationNameList:"Region" min:"3" type:"list"`

	// A complex type that contains one ResettableElementName element for each element
	// that you want to reset to the default value. Valid values for ResettableElementName
	// include the following:
	//
	//    * ChildHealthChecks: Amazon Route 53 resets ChildHealthChecks (https://docs.aws.amazon.com/Route53/latest/APIReference/API_HealthCheckConfig.html#Route53-Type-HealthCheckConfig-ChildHealthChecks)
	//    to null.
	//
	//    * FullyQualifiedDomainName: Route 53 resets FullyQualifiedDomainName (https://docs.aws.amazon.com/Route53/latest/APIReference/API_UpdateHealthCheck.html#Route53-UpdateHealthCheck-request-FullyQualifiedDomainName).
	//    to null.
	//
	//    * Regions: Route 53 resets the Regions (https://docs.aws.amazon.com/Route53/latest/APIReference/API_HealthCheckConfig.html#Route53-Type-HealthCheckConfig-Regions)
	//    list to the default set of regions.
	//
	//    * ResourcePath: Route 53 resets ResourcePath (https://docs.aws.amazon.com/Route53/latest/APIReference/API_HealthCheckConfig.html#Route53-Type-HealthCheckConfig-ResourcePath)
	//    to null.
	ResetElements []ResettableElementName `locationNameList:"ResettableElementName" type:"list"`

	// The path that you want Amazon Route 53 to request when performing health
	// checks. The path can be any value for which your endpoint will return an
	// HTTP status code of 2xx or 3xx when the endpoint is healthy, for example
	// the file /docs/route53-health-check.html. You can also include query string
	// parameters, for example, /welcome.html?language=jp&login=y.
	//
	// Specify this value only if you want to change it.
	ResourcePath *string `type:"string"`

	// If the value of Type is HTTP_STR_MATCH or HTTP_STR_MATCH, the string that
	// you want Amazon Route 53 to search for in the response body from the specified
	// resource. If the string appears in the response body, Route 53 considers
	// the resource healthy. (You can't change the value of Type when you update
	// a health check.)
	SearchString *string `type:"string"`
}

// String returns the string representation
func (s UpdateHealthCheckInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateHealthCheckInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateHealthCheckInput"}
	if s.FailureThreshold != nil && *s.FailureThreshold < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("FailureThreshold", 1))
	}

	if s.HealthCheckId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HealthCheckId"))
	}
	if s.HealthCheckVersion != nil && *s.HealthCheckVersion < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("HealthCheckVersion", 1))
	}
	if s.Port != nil && *s.Port < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Port", 1))
	}
	if s.Regions != nil && len(s.Regions) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("Regions", 3))
	}
	if s.AlarmIdentifier != nil {
		if err := s.AlarmIdentifier.Validate(); err != nil {
			invalidParams.AddNested("AlarmIdentifier", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateHealthCheckInput) MarshalFields(e protocol.FieldEncoder) error {

	e.SetFields(protocol.BodyTarget, "UpdateHealthCheckRequest", protocol.FieldMarshalerFunc(func(e protocol.FieldEncoder) error {
		if s.AlarmIdentifier != nil {
			v := s.AlarmIdentifier

			metadata := protocol.Metadata{}
			e.SetFields(protocol.BodyTarget, "AlarmIdentifier", v, metadata)
		}
		if s.ChildHealthChecks != nil {
			v := s.ChildHealthChecks

			metadata := protocol.Metadata{ListLocationName: "ChildHealthCheck"}
			ls0 := e.List(protocol.BodyTarget, "ChildHealthChecks", metadata)
			ls0.Start()
			for _, v1 := range v {
				ls0.ListAddValue(protocol.StringValue(v1))
			}
			ls0.End()

		}
		if s.Disabled != nil {
			v := *s.Disabled

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "Disabled", protocol.BoolValue(v), metadata)
		}
		if s.EnableSNI != nil {
			v := *s.EnableSNI

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "EnableSNI", protocol.BoolValue(v), metadata)
		}
		if s.FailureThreshold != nil {
			v := *s.FailureThreshold

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "FailureThreshold", protocol.Int64Value(v), metadata)
		}
		if s.FullyQualifiedDomainName != nil {
			v := *s.FullyQualifiedDomainName

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "FullyQualifiedDomainName", protocol.StringValue(v), metadata)
		}
		if s.HealthCheckVersion != nil {
			v := *s.HealthCheckVersion

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "HealthCheckVersion", protocol.Int64Value(v), metadata)
		}
		if s.HealthThreshold != nil {
			v := *s.HealthThreshold

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "HealthThreshold", protocol.Int64Value(v), metadata)
		}
		if s.IPAddress != nil {
			v := *s.IPAddress

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "IPAddress", protocol.StringValue(v), metadata)
		}
		if len(s.InsufficientDataHealthStatus) > 0 {
			v := s.InsufficientDataHealthStatus

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "InsufficientDataHealthStatus", v, metadata)
		}
		if s.Inverted != nil {
			v := *s.Inverted

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "Inverted", protocol.BoolValue(v), metadata)
		}
		if s.Port != nil {
			v := *s.Port

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "Port", protocol.Int64Value(v), metadata)
		}
		if s.Regions != nil {
			v := s.Regions

			metadata := protocol.Metadata{ListLocationName: "Region"}
			ls0 := e.List(protocol.BodyTarget, "Regions", metadata)
			ls0.Start()
			for _, v1 := range v {
				ls0.ListAddValue(protocol.StringValue(v1))
			}
			ls0.End()

		}
		if s.ResetElements != nil {
			v := s.ResetElements

			metadata := protocol.Metadata{ListLocationName: "ResettableElementName"}
			ls0 := e.List(protocol.BodyTarget, "ResetElements", metadata)
			ls0.Start()
			for _, v1 := range v {
				ls0.ListAddValue(protocol.StringValue(v1))
			}
			ls0.End()

		}
		if s.ResourcePath != nil {
			v := *s.ResourcePath

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "ResourcePath", protocol.StringValue(v), metadata)
		}
		if s.SearchString != nil {
			v := *s.SearchString

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "SearchString", protocol.StringValue(v), metadata)
		}
		return nil
	}), protocol.Metadata{XMLNamespaceURI: "https://route53.amazonaws.com/doc/2013-04-01/"})
	if s.HealthCheckId != nil {
		v := *s.HealthCheckId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "HealthCheckId", protocol.StringValue(v), metadata)
	}
	return nil
}

// A complex type that contains the response to the UpdateHealthCheck request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/UpdateHealthCheckResponse
type UpdateHealthCheckOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains the response to an UpdateHealthCheck request.
	//
	// HealthCheck is a required field
	HealthCheck *HealthCheck `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateHealthCheckOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateHealthCheckOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.HealthCheck != nil {
		v := s.HealthCheck

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "HealthCheck", v, metadata)
	}
	return nil
}

const opUpdateHealthCheck = "UpdateHealthCheck"

// UpdateHealthCheckRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Updates an existing health check. Note that some values can't be updated.
//
// For more information about updating health checks, see Creating, Updating,
// and Deleting Health Checks (http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-creating-deleting.html)
// in the Amazon Route 53 Developer Guide.
//
//    // Example sending a request using UpdateHealthCheckRequest.
//    req := client.UpdateHealthCheckRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/UpdateHealthCheck
func (c *Client) UpdateHealthCheckRequest(input *UpdateHealthCheckInput) UpdateHealthCheckRequest {
	op := &aws.Operation{
		Name:       opUpdateHealthCheck,
		HTTPMethod: "POST",
		HTTPPath:   "/2013-04-01/healthcheck/{HealthCheckId}",
	}

	if input == nil {
		input = &UpdateHealthCheckInput{}
	}

	req := c.newRequest(op, input, &UpdateHealthCheckOutput{})
	return UpdateHealthCheckRequest{Request: req, Input: input, Copy: c.UpdateHealthCheckRequest}
}

// UpdateHealthCheckRequest is the request type for the
// UpdateHealthCheck API operation.
type UpdateHealthCheckRequest struct {
	*aws.Request
	Input *UpdateHealthCheckInput
	Copy  func(*UpdateHealthCheckInput) UpdateHealthCheckRequest
}

// Send marshals and sends the UpdateHealthCheck API request.
func (r UpdateHealthCheckRequest) Send(ctx context.Context) (*UpdateHealthCheckResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateHealthCheckResponse{
		UpdateHealthCheckOutput: r.Request.Data.(*UpdateHealthCheckOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateHealthCheckResponse is the response type for the
// UpdateHealthCheck API operation.
type UpdateHealthCheckResponse struct {
	*UpdateHealthCheckOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateHealthCheck request.
func (r *UpdateHealthCheckResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
