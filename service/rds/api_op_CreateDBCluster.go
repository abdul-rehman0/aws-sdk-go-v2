// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CreateDBClusterMessage
type CreateDBClusterInput struct {
	_ struct{} `type:"structure"`

	// A list of Availability Zones (AZs) where instances in the DB cluster can
	// be created. For information on AWS Regions and Availability Zones, see Choosing
	// the Regions and Availability Zones (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.RegionsAndAvailabilityZones.html)
	// in the Amazon Aurora User Guide.
	AvailabilityZones []string `locationNameList:"AvailabilityZone" type:"list"`

	// The target backtrack window, in seconds. To disable backtracking, set this
	// value to 0.
	//
	// Default: 0
	//
	// Constraints:
	//
	//    * If specified, this value must be set to a number from 0 to 259,200 (72
	//    hours).
	BacktrackWindow *int64 `type:"long"`

	// The number of days for which automated backups are retained.
	//
	// Default: 1
	//
	// Constraints:
	//
	//    * Must be a value from 1 to 35
	BackupRetentionPeriod *int64 `type:"integer"`

	// A value that indicates that the DB cluster should be associated with the
	// specified CharacterSet.
	CharacterSetName *string `type:"string"`

	// A value that indicates whether to copy all tags from the DB cluster to snapshots
	// of the DB cluster. The default is not to copy them.
	CopyTagsToSnapshot *bool `type:"boolean"`

	// The DB cluster identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//    * First character must be a letter.
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example: my-cluster1
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// The name of the DB cluster parameter group to associate with this DB cluster.
	// If this argument is omitted, default.aurora5.6 is used.
	//
	// Constraints:
	//
	//    * If supplied, must match the name of an existing DB cluster parameter
	//    group.
	DBClusterParameterGroupName *string `type:"string"`

	// A DB subnet group to associate with this DB cluster.
	//
	// Constraints: Must match the name of an existing DBSubnetGroup. Must not be
	// default.
	//
	// Example: mySubnetgroup
	DBSubnetGroupName *string `type:"string"`

	// The name for your database of up to 64 alpha-numeric characters. If you do
	// not provide a name, Amazon RDS will not create a database in the DB cluster
	// you are creating.
	DatabaseName *string `type:"string"`

	// A value that indicates whether the DB cluster has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled.
	DeletionProtection *bool `type:"boolean"`

	// DestinationRegion is used for presigning the request to a given region.
	DestinationRegion *string `type:"string"`

	// The list of log types that need to be enabled for exporting to CloudWatch
	// Logs. The values in the list depend on the DB engine being used. For more
	// information, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon Aurora User Guide.
	EnableCloudwatchLogsExports []string `type:"list"`

	// A value that indicates whether to enable mapping of AWS Identity and Access
	// Management (IAM) accounts to database accounts. By default, mapping is disabled.
	EnableIAMDatabaseAuthentication *bool `type:"boolean"`

	// The name of the database engine to be used for this DB cluster.
	//
	// Valid Values: aurora (for MySQL 5.6-compatible Aurora), aurora-mysql (for
	// MySQL 5.7-compatible Aurora), and aurora-postgresql
	//
	// Engine is a required field
	Engine *string `type:"string" required:"true"`

	// The DB engine mode of the DB cluster, either provisioned, serverless, parallelquery,
	// or global.
	EngineMode *string `type:"string"`

	// The version number of the database engine to use.
	//
	// Aurora MySQL
	//
	// Example: 5.6.10a, 5.7.12
	//
	// Aurora PostgreSQL
	//
	// Example: 9.6.3
	EngineVersion *string `type:"string"`

	// The global cluster ID of an Aurora cluster that becomes the primary cluster
	// in the new global database cluster.
	GlobalClusterIdentifier *string `type:"string"`

	// The AWS KMS key identifier for an encrypted DB cluster.
	//
	// The KMS key identifier is the Amazon Resource Name (ARN) for the KMS encryption
	// key. If you are creating a DB cluster with the same AWS account that owns
	// the KMS encryption key used to encrypt the new DB cluster, then you can use
	// the KMS key alias instead of the ARN for the KMS encryption key.
	//
	// If an encryption key is not specified in KmsKeyId:
	//
	//    * If ReplicationSourceIdentifier identifies an encrypted source, then
	//    Amazon RDS will use the encryption key used to encrypt the source. Otherwise,
	//    Amazon RDS will use your default encryption key.
	//
	//    * If the StorageEncrypted parameter is enabled and ReplicationSourceIdentifier
	//    is not specified, then Amazon RDS will use your default encryption key.
	//
	// AWS KMS creates the default encryption key for your AWS account. Your AWS
	// account has a different default encryption key for each AWS Region.
	//
	// If you create a Read Replica of an encrypted DB cluster in another AWS Region,
	// you must set KmsKeyId to a KMS key ID that is valid in the destination AWS
	// Region. This key is used to encrypt the Read Replica in that AWS Region.
	KmsKeyId *string `type:"string"`

	// The password for the master database user. This password can contain any
	// printable ASCII character except "/", """, or "@".
	//
	// Constraints: Must contain from 8 to 41 characters.
	MasterUserPassword *string `type:"string"`

	// The name of the master user for the DB cluster.
	//
	// Constraints:
	//
	//    * Must be 1 to 16 letters or numbers.
	//
	//    * First character must be a letter.
	//
	//    * Can't be a reserved word for the chosen database engine.
	MasterUsername *string `type:"string"`

	// A value that indicates that the DB cluster should be associated with the
	// specified option group.
	//
	// Permanent options can't be removed from an option group. The option group
	// can't be removed from a DB cluster once it is associated with a DB cluster.
	OptionGroupName *string `type:"string"`

	// The port number on which the instances in the DB cluster accept connections.
	//
	// Default: 3306 if engine is set as aurora or 5432 if set to aurora-postgresql.
	Port *int64 `type:"integer"`

	// A URL that contains a Signature Version 4 signed request for the CreateDBCluster
	// action to be called in the source AWS Region where the DB cluster is replicated
	// from. You only need to specify PreSignedUrl when you are performing cross-region
	// replication from an encrypted DB cluster.
	//
	// The pre-signed URL must be a valid request for the CreateDBCluster API action
	// that can be executed in the source AWS Region that contains the encrypted
	// DB cluster to be copied.
	//
	// The pre-signed URL request must contain the following parameter values:
	//
	//    * KmsKeyId - The AWS KMS key identifier for the key to use to encrypt
	//    the copy of the DB cluster in the destination AWS Region. This should
	//    refer to the same KMS key for both the CreateDBCluster action that is
	//    called in the destination AWS Region, and the action contained in the
	//    pre-signed URL.
	//
	//    * DestinationRegion - The name of the AWS Region that Aurora Read Replica
	//    will be created in.
	//
	//    * ReplicationSourceIdentifier - The DB cluster identifier for the encrypted
	//    DB cluster to be copied. This identifier must be in the Amazon Resource
	//    Name (ARN) format for the source AWS Region. For example, if you are copying
	//    an encrypted DB cluster from the us-west-2 AWS Region, then your ReplicationSourceIdentifier
	//    would look like Example: arn:aws:rds:us-west-2:123456789012:cluster:aurora-cluster1.
	//
	// To learn how to generate a Signature Version 4 signed request, see Authenticating
	// Requests: Using Query Parameters (AWS Signature Version 4) (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html)
	// and Signature Version 4 Signing Process (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
	PreSignedUrl *string `type:"string"`

	// The daily time range during which automated backups are created if automated
	// backups are enabled using the BackupRetentionPeriod parameter.
	//
	// The default is a 30-minute window selected at random from an 8-hour block
	// of time for each AWS Region. To see the time blocks available, see Adjusting
	// the Preferred DB Cluster Maintenance Window (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora)
	// in the Amazon Aurora User Guide.
	//
	// Constraints:
	//
	//    * Must be in the format hh24:mi-hh24:mi.
	//
	//    * Must be in Universal Coordinated Time (UTC).
	//
	//    * Must not conflict with the preferred maintenance window.
	//
	//    * Must be at least 30 minutes.
	PreferredBackupWindow *string `type:"string"`

	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC).
	//
	// Format: ddd:hh24:mi-ddd:hh24:mi
	//
	// The default is a 30-minute window selected at random from an 8-hour block
	// of time for each AWS Region, occurring on a random day of the week. To see
	// the time blocks available, see Adjusting the Preferred DB Cluster Maintenance
	// Window (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora)
	// in the Amazon Aurora User Guide.
	//
	// Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `type:"string"`

	// The Amazon Resource Name (ARN) of the source DB instance or DB cluster if
	// this DB cluster is created as a Read Replica.
	ReplicationSourceIdentifier *string `type:"string"`

	// For DB clusters in serverless DB engine mode, the scaling properties of the
	// DB cluster.
	ScalingConfiguration *ScalingConfiguration `type:"structure"`

	// SourceRegion is the source region where the resource exists. This is not
	// sent over the wire and is only used for presigning. This value should always
	// have the same region as the source ARN.
	SourceRegion *string `type:"string" ignore:"true"`

	// A value that indicates whether the DB cluster is encrypted.
	StorageEncrypted *bool `type:"boolean"`

	// Tags to assign to the DB cluster.
	Tags []Tag `locationNameList:"Tag" type:"list"`

	// A list of EC2 VPC security groups to associate with this DB cluster.
	VpcSecurityGroupIds []string `locationNameList:"VpcSecurityGroupId" type:"list"`
}

// String returns the string representation
func (s CreateDBClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDBClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDBClusterInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if s.Engine == nil {
		invalidParams.Add(aws.NewErrParamRequired("Engine"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CreateDBClusterResult
type CreateDBClusterOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon Aurora DB cluster.
	//
	// This data type is used as a response element in the DescribeDBClusters, StopDBCluster,
	// and StartDBCluster actions.
	DBCluster *DBCluster `type:"structure"`
}

// String returns the string representation
func (s CreateDBClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDBCluster = "CreateDBCluster"

// CreateDBClusterRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Creates a new Amazon Aurora DB cluster.
//
// You can use the ReplicationSourceIdentifier parameter to create the DB cluster
// as a Read Replica of another DB cluster or Amazon RDS MySQL DB instance.
// For cross-region replication where the DB cluster identified by ReplicationSourceIdentifier
// is encrypted, you must also specify the PreSignedUrl parameter.
//
// For more information on Amazon Aurora, see What Is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using CreateDBClusterRequest.
//    req := client.CreateDBClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CreateDBCluster
func (c *Client) CreateDBClusterRequest(input *CreateDBClusterInput) CreateDBClusterRequest {
	op := &aws.Operation{
		Name:       opCreateDBCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDBClusterInput{}
	}

	req := c.newRequest(op, input, &CreateDBClusterOutput{})
	return CreateDBClusterRequest{Request: req, Input: input, Copy: c.CreateDBClusterRequest}
}

// CreateDBClusterRequest is the request type for the
// CreateDBCluster API operation.
type CreateDBClusterRequest struct {
	*aws.Request
	Input *CreateDBClusterInput
	Copy  func(*CreateDBClusterInput) CreateDBClusterRequest
}

// Send marshals and sends the CreateDBCluster API request.
func (r CreateDBClusterRequest) Send(ctx context.Context) (*CreateDBClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDBClusterResponse{
		CreateDBClusterOutput: r.Request.Data.(*CreateDBClusterOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDBClusterResponse is the response type for the
// CreateDBCluster API operation.
type CreateDBClusterResponse struct {
	*CreateDBClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDBCluster request.
func (r *CreateDBClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
