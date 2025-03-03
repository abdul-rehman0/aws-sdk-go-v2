// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eventbridge

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// This structure specifies the VPC subnets and security groups for the task
// and whether a public IP address is to be used. This structure is relevant
// only for ECS tasks that use the awsvpc network mode.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/AwsVpcConfiguration
type AwsVpcConfiguration struct {
	_ struct{} `type:"structure"`

	// Specifies whether the task's elastic network interface receives a public
	// IP address. You can specify ENABLED only when LaunchType in EcsParameters
	// is set to FARGATE.
	AssignPublicIp AssignPublicIp `type:"string" enum:"true"`

	// Specifies the security groups associated with the task. These security groups
	// must all be in the same VPC. You can specify as many as five security groups.
	// If you don't specify a security group, the default security group for the
	// VPC is used.
	SecurityGroups []string `type:"list"`

	// Specifies the subnets associated with the task. These subnets must all be
	// in the same VPC. You can specify as many as 16 subnets.
	//
	// Subnets is a required field
	Subnets []string `type:"list" required:"true"`
}

// String returns the string representation
func (s AwsVpcConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AwsVpcConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AwsVpcConfiguration"}

	if s.Subnets == nil {
		invalidParams.Add(aws.NewErrParamRequired("Subnets"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The array properties for the submitted job, such as the size of the array.
// The array size can be between 2 and 10,000. If you specify array properties
// for a job, it becomes an array job. This parameter is used only if the target
// is an AWS Batch job.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/BatchArrayProperties
type BatchArrayProperties struct {
	_ struct{} `type:"structure"`

	// The size of the array, if this is an array batch job. Valid values are integers
	// between 2 and 10,000.
	Size *int64 `type:"integer"`
}

// String returns the string representation
func (s BatchArrayProperties) String() string {
	return awsutil.Prettify(s)
}

// The custom parameters to be used when the target is an AWS Batch job.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/BatchParameters
type BatchParameters struct {
	_ struct{} `type:"structure"`

	// The array properties for the submitted job, such as the size of the array.
	// The array size can be between 2 and 10,000. If you specify array properties
	// for a job, it becomes an array job. This parameter is used only if the target
	// is an AWS Batch job.
	ArrayProperties *BatchArrayProperties `type:"structure"`

	// The ARN or name of the job definition to use if the event target is an AWS
	// Batch job. This job definition must already exist.
	//
	// JobDefinition is a required field
	JobDefinition *string `type:"string" required:"true"`

	// The name to use for this execution of the job, if the target is an AWS Batch
	// job.
	//
	// JobName is a required field
	JobName *string `type:"string" required:"true"`

	// The retry strategy to use for failed jobs if the target is an AWS Batch job.
	// The retry strategy is the number of times to retry the failed job execution.
	// Valid values are 1–10. When you specify a retry strategy here, it overrides
	// the retry strategy defined in the job definition.
	RetryStrategy *BatchRetryStrategy `type:"structure"`
}

// String returns the string representation
func (s BatchParameters) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchParameters) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchParameters"}

	if s.JobDefinition == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobDefinition"))
	}

	if s.JobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The retry strategy to use for failed jobs if the target is an AWS Batch job.
// If you specify a retry strategy here, it overrides the retry strategy defined
// in the job definition.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/BatchRetryStrategy
type BatchRetryStrategy struct {
	_ struct{} `type:"structure"`

	// The number of times to attempt to retry, if the job fails. Valid values are
	// 1–10.
	Attempts *int64 `type:"integer"`
}

// String returns the string representation
func (s BatchRetryStrategy) String() string {
	return awsutil.Prettify(s)
}

// A JSON string that you can use to limit the event bus permissions that you're
// granting to only accounts that fulfill the condition. Currently, the only
// supported condition is membership in a certain AWS organization. The string
// must contain Type, Key, and Value fields. The Value field specifies the ID
// of the AWS organization. The following is an example value for Condition:
//
// '{"Type" : "StringEquals", "Key": "aws:PrincipalOrgID", "Value": "o-1234567890"}'
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/Condition
type Condition struct {
	_ struct{} `type:"structure"`

	// The key for the condition. Currently, the only supported key is aws:PrincipalOrgID.
	//
	// Key is a required field
	Key *string `type:"string" required:"true"`

	// The type of condition. Currently, the only supported value is StringEquals.
	//
	// Type is a required field
	Type *string `type:"string" required:"true"`

	// The value for the key. Currently, this must be the ID of the organization.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Condition) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Condition) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Condition"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}

	if s.Type == nil {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The custom parameters to be used when the target is an Amazon ECS task.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/EcsParameters
type EcsParameters struct {
	_ struct{} `type:"structure"`

	// Specifies an ECS task group for the task. The maximum length is 255 characters.
	Group *string `type:"string"`

	// Specifies the launch type on which your task is running. The launch type
	// that you specify here must match one of the launch type (compatibilities)
	// of the target task. The FARGATE value is supported only in the Regions where
	// AWS Fargate with Amazon ECS is supported. For more information, see AWS Fargate
	// on Amazon ECS (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/AWS-Fargate.html)
	// in the Amazon Elastic Container Service Developer Guide.
	LaunchType LaunchType `type:"string" enum:"true"`

	// Use this structure if the ECS task uses the awsvpc network mode. This structure
	// specifies the VPC subnets and security groups associated with the task and
	// whether a public IP address is to be used. This structure is required if
	// LaunchType is FARGATE because the awsvpc mode is required for Fargate tasks.
	//
	// If you specify NetworkConfiguration when the target ECS task doesn't use
	// the awsvpc network mode, the task fails.
	NetworkConfiguration *NetworkConfiguration `type:"structure"`

	// Specifies the platform version for the task. Specify only the numeric portion
	// of the platform version, such as 1.1.0.
	//
	// This structure is used only if LaunchType is FARGATE. For more information
	// about valid platform versions, see AWS Fargate Platform Versions (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	PlatformVersion *string `type:"string"`

	// The number of tasks to create based on TaskDefinition. The default is 1.
	TaskCount *int64 `min:"1" type:"integer"`

	// The ARN of the task definition to use if the event target is an Amazon ECS
	// task.
	//
	// TaskDefinitionArn is a required field
	TaskDefinitionArn *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s EcsParameters) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EcsParameters) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EcsParameters"}
	if s.TaskCount != nil && *s.TaskCount < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("TaskCount", 1))
	}

	if s.TaskDefinitionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskDefinitionArn"))
	}
	if s.TaskDefinitionArn != nil && len(*s.TaskDefinitionArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TaskDefinitionArn", 1))
	}
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			invalidParams.AddNested("NetworkConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An event bus receives events from a source and routes them to rules associated
// with that event bus. Your account's default event bus receives rules from
// AWS services. A custom event bus can receive rules from AWS services as well
// as your custom applications and services. A partner event bus receives events
// from an event source created by an SaaS partner. These events come from the
// partners services or applications.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/EventBus
type EventBus struct {
	_ struct{} `type:"structure"`

	// The ARN of the event bus.
	Arn *string `type:"string"`

	// The name of the event bus.
	Name *string `type:"string"`

	// The permissions policy of the event bus, describing which other AWS accounts
	// can write events to this event bus.
	Policy *string `type:"string"`
}

// String returns the string representation
func (s EventBus) String() string {
	return awsutil.Prettify(s)
}

// A partner event source is created by an SaaS partner. If a customer creates
// a partner event bus that matches this event source, that AWS account can
// receive events from the partner's applications or services.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/EventSource
type EventSource struct {
	_ struct{} `type:"structure"`

	// The ARN of the event source.
	Arn *string `type:"string"`

	// The name of the partner that created the event source.
	CreatedBy *string `type:"string"`

	// The date and time when the event source was created.
	CreationTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The date and time when the event source will expire if the AWS account doesn't
	// create a matching event bus for it.
	ExpirationTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The name of the event source.
	Name *string `type:"string"`

	// The state of the event source. If it's ACTIVE, you have already created a
	// matching event bus for this event source, and that event bus is active. If
	// it's PENDING, either you haven't yet created a matching event bus, or that
	// event bus is deactivated. If it's DELETED, you have created a matching event
	// bus, but the event source has since been deleted.
	State EventSourceState `type:"string" enum:"true"`
}

// String returns the string representation
func (s EventSource) String() string {
	return awsutil.Prettify(s)
}

// Contains the parameters needed for you to provide custom input to a target
// based on one or more pieces of data extracted from the event.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/InputTransformer
type InputTransformer struct {
	_ struct{} `type:"structure"`

	// Map of JSON paths to be extracted from the event. You can then insert these
	// in the template in InputTemplate to produce the output to be sent to the
	// target.
	//
	// InputPathsMap is an array key-value pairs, where each value is a valid JSON
	// path. You can have as many as 10 key-value pairs. You must use JSON dot notation,
	// not bracket notation.
	//
	// The keys can't start with "AWS".
	InputPathsMap map[string]string `type:"map"`

	// Input template where you specify placeholders that will be filled with the
	// values of the keys from InputPathsMap to customize the data sent to the target.
	// Enclose each InputPathsMaps value in brackets: <value>. The InputTemplate
	// must be valid JSON.
	//
	// If InputTemplate is a JSON object (surrounded by curly braces), the following
	// restrictions apply:
	//
	//    * The placeholder can't be used as an object key
	//
	//    * Object values can't include quote marks
	//
	// The following example shows the syntax for using InputPathsMap and InputTemplate.
	//
	// "InputTransformer":
	//
	// {
	//
	// "InputPathsMap": {"instance": "$.detail.instance","status": "$.detail.status"},
	//
	// "InputTemplate": "<instance> is in state <status>"
	//
	// }
	//
	// To have the InputTemplate include quote marks within a JSON string, escape
	// each quote marks with a slash, as in the following example:
	//
	// "InputTransformer":
	//
	// {
	//
	// "InputPathsMap": {"instance": "$.detail.instance","status": "$.detail.status"},
	//
	// "InputTemplate": "<instance> is in state \"<status>\""
	//
	// }
	//
	// InputTemplate is a required field
	InputTemplate *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s InputTransformer) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InputTransformer) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "InputTransformer"}

	if s.InputTemplate == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputTemplate"))
	}
	if s.InputTemplate != nil && len(*s.InputTemplate) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InputTemplate", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// This object enables you to specify a JSON path to extract from the event
// and use as the partition key for the Amazon Kinesis data stream so that you
// can control the shard that the event goes to. If you don't include this parameter,
// the default is to use the eventId as the partition key.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/KinesisParameters
type KinesisParameters struct {
	_ struct{} `type:"structure"`

	// The JSON path to be extracted from the event and used as the partition key.
	// For more information, see Amazon Kinesis Streams Key Concepts (https://docs.aws.amazon.com/streams/latest/dev/key-concepts.html#partition-key)
	// in the Amazon Kinesis Streams Developer Guide.
	//
	// PartitionKeyPath is a required field
	PartitionKeyPath *string `type:"string" required:"true"`
}

// String returns the string representation
func (s KinesisParameters) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *KinesisParameters) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "KinesisParameters"}

	if s.PartitionKeyPath == nil {
		invalidParams.Add(aws.NewErrParamRequired("PartitionKeyPath"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// This structure specifies the network configuration for an ECS task.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/NetworkConfiguration
type NetworkConfiguration struct {
	_ struct{} `type:"structure"`

	// Use this structure to specify the VPC subnets and security groups for the
	// task and whether a public IP address is to be used. This structure is relevant
	// only for ECS tasks that use the awsvpc network mode.
	AwsvpcConfiguration *AwsVpcConfiguration `locationName:"awsvpcConfiguration" type:"structure"`
}

// String returns the string representation
func (s NetworkConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *NetworkConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "NetworkConfiguration"}
	if s.AwsvpcConfiguration != nil {
		if err := s.AwsvpcConfiguration.Validate(); err != nil {
			invalidParams.AddNested("AwsvpcConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A partner event source is created by an SaaS partner. If a customer creates
// a partner event bus that matches this event source, that AWS account can
// receive events from the partner's applications or services.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/PartnerEventSource
type PartnerEventSource struct {
	_ struct{} `type:"structure"`

	// The ARN of the partner event source.
	Arn *string `type:"string"`

	// The name of the partner event source.
	Name *string `type:"string"`
}

// String returns the string representation
func (s PartnerEventSource) String() string {
	return awsutil.Prettify(s)
}

// The AWS account that a partner event source has been offered to.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/PartnerEventSourceAccount
type PartnerEventSourceAccount struct {
	_ struct{} `type:"structure"`

	// The AWS account ID that the partner event source was offered to.
	Account *string `min:"12" type:"string"`

	// The date and time when the event source was created.
	CreationTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The date and time when the event source will expire if the AWS account doesn't
	// create a matching event bus for it.
	ExpirationTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The state of the event source. If it's ACTIVE, you have already created a
	// matching event bus for this event source, and that event bus is active. If
	// it's PENDING, either you haven't yet created a matching event bus, or that
	// event bus is deactivated. If it's DELETED, you have created a matching event
	// bus, but the event source has since been deleted.
	State EventSourceState `type:"string" enum:"true"`
}

// String returns the string representation
func (s PartnerEventSourceAccount) String() string {
	return awsutil.Prettify(s)
}

// Represents an event to be submitted.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/PutEventsRequestEntry
type PutEventsRequestEntry struct {
	_ struct{} `type:"structure"`

	// A valid JSON string. There is no other schema imposed. The JSON string can
	// contain fields and nested subobjects.
	Detail *string `type:"string"`

	// Free-form string used to decide which fields to expect in the event detail.
	DetailType *string `type:"string"`

	// The event bus that will receive the event. Only the rules that are associated
	// with this event bus can match the event.
	EventBusName *string `min:"1" type:"string"`

	// AWS resources, identified by Amazon Resource Name (ARN), that the event primarily
	// concerns. Any number, including zero, can be present.
	Resources []string `type:"list"`

	// The source of the event. This field is required.
	Source *string `type:"string"`

	// The timestamp of the event, per RFC3339 (https://www.rfc-editor.org/rfc/rfc3339.txt).
	// If no timestamp is provided, the timestamp of the PutEvents call is used.
	Time *time.Time `type:"timestamp" timestampFormat:"unix"`
}

// String returns the string representation
func (s PutEventsRequestEntry) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutEventsRequestEntry) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutEventsRequestEntry"}
	if s.EventBusName != nil && len(*s.EventBusName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventBusName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents an event that failed to be submitted.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/PutEventsResultEntry
type PutEventsResultEntry struct {
	_ struct{} `type:"structure"`

	// The error code that indicates why the event submission failed.
	ErrorCode *string `type:"string"`

	// The error message that explains why the event submission failed.
	ErrorMessage *string `type:"string"`

	// The ID of the event.
	EventId *string `type:"string"`
}

// String returns the string representation
func (s PutEventsResultEntry) String() string {
	return awsutil.Prettify(s)
}

// The details about an event generated by an SaaS partner.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/PutPartnerEventsRequestEntry
type PutPartnerEventsRequestEntry struct {
	_ struct{} `type:"structure"`

	// A valid JSON string. There is no other schema imposed. The JSON string can
	// contain fields and nested subobjects.
	Detail *string `type:"string"`

	// A free-form string used to decide which fields to expect in the event detail.
	DetailType *string `type:"string"`

	// AWS resources, identified by Amazon Resource Name (ARN), that the event primarily
	// concerns. Any number, including zero, can be present.
	Resources []string `type:"list"`

	// The event source that is generating the evntry.
	Source *string `type:"string"`

	// The date and time of the event.
	Time *time.Time `type:"timestamp" timestampFormat:"unix"`
}

// String returns the string representation
func (s PutPartnerEventsRequestEntry) String() string {
	return awsutil.Prettify(s)
}

// Represents an event that a partner tried to generate but failed.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/PutPartnerEventsResultEntry
type PutPartnerEventsResultEntry struct {
	_ struct{} `type:"structure"`

	// The error code that indicates why the event submission failed.
	ErrorCode *string `type:"string"`

	// The error message that explains why the event submission failed.
	ErrorMessage *string `type:"string"`

	// The ID of the event.
	EventId *string `type:"string"`
}

// String returns the string representation
func (s PutPartnerEventsResultEntry) String() string {
	return awsutil.Prettify(s)
}

// Represents a target that failed to be added to a rule.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/PutTargetsResultEntry
type PutTargetsResultEntry struct {
	_ struct{} `type:"structure"`

	// The error code that indicates why the target addition failed. If the value
	// is ConcurrentModificationException, too many requests were made at the same
	// time.
	ErrorCode *string `type:"string"`

	// The error message that explains why the target addition failed.
	ErrorMessage *string `type:"string"`

	// The ID of the target.
	TargetId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s PutTargetsResultEntry) String() string {
	return awsutil.Prettify(s)
}

// Represents a target that failed to be removed from a rule.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/RemoveTargetsResultEntry
type RemoveTargetsResultEntry struct {
	_ struct{} `type:"structure"`

	// The error code that indicates why the target removal failed. If the value
	// is ConcurrentModificationException, too many requests were made at the same
	// time.
	ErrorCode *string `type:"string"`

	// The error message that explains why the target removal failed.
	ErrorMessage *string `type:"string"`

	// The ID of the target.
	TargetId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s RemoveTargetsResultEntry) String() string {
	return awsutil.Prettify(s)
}

// Contains information about a rule in Amazon EventBridge.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/Rule
type Rule struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the rule.
	Arn *string `min:"1" type:"string"`

	// The description of the rule.
	Description *string `type:"string"`

	// The event bus associated with the rule.
	EventBusName *string `min:"1" type:"string"`

	// The event pattern of the rule. For more information, see Event Patterns (https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html)
	// in the Amazon EventBridge User Guide.
	EventPattern *string `type:"string"`

	// If an AWS service created the rule on behalf of your account, this field
	// displays the principal name of the service that created the rule.
	ManagedBy *string `min:"1" type:"string"`

	// The name of the rule.
	Name *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the role that is used for target invocation.
	RoleArn *string `min:"1" type:"string"`

	// The scheduling expression: for example, "cron(0 20 * * ? *)" or "rate(5 minutes)".
	ScheduleExpression *string `type:"string"`

	// The state of the rule.
	State RuleState `type:"string" enum:"true"`
}

// String returns the string representation
func (s Rule) String() string {
	return awsutil.Prettify(s)
}

// This parameter contains the criteria (either InstanceIds or a tag) used to
// specify which EC2 instances are to be sent the command.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/RunCommandParameters
type RunCommandParameters struct {
	_ struct{} `type:"structure"`

	// Currently, we support including only one RunCommandTarget block, which specifies
	// either an array of InstanceIds or a tag.
	//
	// RunCommandTargets is a required field
	RunCommandTargets []RunCommandTarget `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s RunCommandParameters) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RunCommandParameters) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RunCommandParameters"}

	if s.RunCommandTargets == nil {
		invalidParams.Add(aws.NewErrParamRequired("RunCommandTargets"))
	}
	if s.RunCommandTargets != nil && len(s.RunCommandTargets) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RunCommandTargets", 1))
	}
	if s.RunCommandTargets != nil {
		for i, v := range s.RunCommandTargets {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "RunCommandTargets", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about the EC2 instances that are to be sent the command, specified
// as key-value pairs. Each RunCommandTarget block can include only one key,
// but this key can specify multiple values.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/RunCommandTarget
type RunCommandTarget struct {
	_ struct{} `type:"structure"`

	// Can be either tag: tag-key or InstanceIds.
	//
	// Key is a required field
	Key *string `min:"1" type:"string" required:"true"`

	// If Key is tag: tag-key, Values is a list of tag values. If Key is InstanceIds,
	// Values is a list of Amazon EC2 instance IDs.
	//
	// Values is a required field
	Values []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s RunCommandTarget) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RunCommandTarget) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RunCommandTarget"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.Values == nil {
		invalidParams.Add(aws.NewErrParamRequired("Values"))
	}
	if s.Values != nil && len(s.Values) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Values", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// This structure includes the custom parameter to be used when the target is
// an SQS FIFO queue.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/SqsParameters
type SqsParameters struct {
	_ struct{} `type:"structure"`

	// The FIFO message group ID to use as the target.
	MessageGroupId *string `type:"string"`
}

// String returns the string representation
func (s SqsParameters) String() string {
	return awsutil.Prettify(s)
}

// A key-value pair associated with an AWS resource. In EventBridge, rules support
// tagging.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/Tag
type Tag struct {
	_ struct{} `type:"structure"`

	// A string that you can use to assign a value. The combination of tag keys
	// and values can help you organize and categorize your resources.
	//
	// Key is a required field
	Key *string `min:"1" type:"string" required:"true"`

	// The value for the specified tag key.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Targets are the resources to be invoked when a rule is triggered. For a complete
// list of services and resources that can be set as a target, see PutTargets.
//
// If you're setting the event bus of another account as the target and that
// account granted permission to your account through an organization instead
// of directly by the account ID, you must specify a RoleArn with proper permissions
// in the Target structure. For more information, see Sending and Receiving
// Events Between AWS Accounts (https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html)
// in the Amazon EventBridge User Guide.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/Target
type Target struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the target.
	//
	// Arn is a required field
	Arn *string `min:"1" type:"string" required:"true"`

	// If the event target is an AWS Batch job, this contains the job definition,
	// job name, and other parameters. For more information, see Jobs (https://docs.aws.amazon.com/batch/latest/userguide/jobs.html)
	// in the AWS Batch User Guide.
	BatchParameters *BatchParameters `type:"structure"`

	// Contains the Amazon ECS task definition and task count to be used if the
	// event target is an Amazon ECS task. For more information about Amazon ECS
	// tasks, see Task Definitions (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html)
	// in the Amazon EC2 Container Service Developer Guide.
	EcsParameters *EcsParameters `type:"structure"`

	// The ID of the target.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// Valid JSON text passed to the target. In this case, nothing from the event
	// itself is passed to the target. For more information, see The JavaScript
	// Object Notation (JSON) Data Interchange Format (http://www.rfc-editor.org/rfc/rfc7159.txt).
	Input *string `type:"string"`

	// The value of the JSONPath that is used for extracting part of the matched
	// event when passing it to the target. You must use JSON dot notation, not
	// bracket notation. For more information about JSON paths, see JSONPath (http://goessner.net/articles/JsonPath/).
	InputPath *string `type:"string"`

	// Settings to enable you to provide custom input to a target based on certain
	// event data. You can extract one or more key-value pairs from the event and
	// then use that data to send customized input to the target.
	InputTransformer *InputTransformer `type:"structure"`

	// The custom parameter that you can use to control the shard assignment when
	// the target is a Kinesis data stream. If you don't include this parameter,
	// the default is to use the eventId as the partition key.
	KinesisParameters *KinesisParameters `type:"structure"`

	// The Amazon Resource Name (ARN) of the IAM role to be used for this target
	// when the rule is triggered. If one rule triggers multiple targets, you can
	// use a different IAM role for each target.
	RoleArn *string `min:"1" type:"string"`

	// Parameters used when you are using the rule to invoke Amazon EC2 Run Command.
	RunCommandParameters *RunCommandParameters `type:"structure"`

	// Contains the message group ID to use when the target is a FIFO queue.
	//
	// If you specify an SQS FIFO queue as a target, the queue must have content-based
	// deduplication enabled.
	SqsParameters *SqsParameters `type:"structure"`
}

// String returns the string representation
func (s Target) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Target) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Target"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 1))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 1))
	}
	if s.BatchParameters != nil {
		if err := s.BatchParameters.Validate(); err != nil {
			invalidParams.AddNested("BatchParameters", err.(aws.ErrInvalidParams))
		}
	}
	if s.EcsParameters != nil {
		if err := s.EcsParameters.Validate(); err != nil {
			invalidParams.AddNested("EcsParameters", err.(aws.ErrInvalidParams))
		}
	}
	if s.InputTransformer != nil {
		if err := s.InputTransformer.Validate(); err != nil {
			invalidParams.AddNested("InputTransformer", err.(aws.ErrInvalidParams))
		}
	}
	if s.KinesisParameters != nil {
		if err := s.KinesisParameters.Validate(); err != nil {
			invalidParams.AddNested("KinesisParameters", err.(aws.ErrInvalidParams))
		}
	}
	if s.RunCommandParameters != nil {
		if err := s.RunCommandParameters.Validate(); err != nil {
			invalidParams.AddNested("RunCommandParameters", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
