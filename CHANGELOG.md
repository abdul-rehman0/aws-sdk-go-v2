Release v0.10.0 (2019-07-24)
===

### Services
* Synced the V2 SDK with latest AWS service API definitions.
* Fixes [#341](https://github.com/aws/aws-sdk-go-v2/issues/341)
* Fixes [#342](https://github.com/aws/aws-sdk-go-v2/issues/342)

### SDK Breaking Changes
* `aws`: Add default HTTP client instead of http.DefaultClient/Transport ([#315](https://github.com/aws/aws-sdk-go-v2/pull/315))
  * Adds a new BuildableHTTPClient type to the SDK's aws package. The type uses the builder pattern with immutable changes. Modifications to the buildable client create copies of the client.  Adds a HTTPClient interface to the aws package that the SDK will use as an abstraction over the specific HTTP client implementation. The SDK will default to the BuildableHTTPClient, but a *http.Client can be also provided for custom configuration.  When the SDK's aws.Config.HTTPClient value is a BuildableHTTPClient the SDK will be able to use API client specific request timeout options.
  * Fixes [#279](https://github.com/aws/aws-sdk-go-v2/issues/279)
  * Fixes [#269](https://github.com/aws/aws-sdk-go-v2/issues/269)

### SDK Enhancements
* `service/s3/s3manager`: Update S3 Upload Multipart location ([#324](https://github.com/aws/aws-sdk-go-v2/pull/324))
  * Updates the Location returned value of S3 Upload's Multipart UploadOutput type to be consistent with single part upload URL. This update also brings the multipart upload Location inline with the S3 object URLs created by the SDK.
  * Fixes [#323](https://github.com/aws/aws-sdk-go-v2/issues/323)
  * V2 Port [aws/aws-sdk-go#2453](https://github.com/aws/aws-sdk-go/issues/2453)

### SDK Bugs
* `private/model`: Handles empty map vs unset map behavior in send request ([#337](https://github.com/aws/aws-sdk-go-v2/pull/337))
  * Updated shape marshal model to handle the empty map vs nil map behavior. Adding a test case to assert behavior when a user sends an empty map vs nil map.
  * Fix [#332](https://github.com/aws/aws-sdk-go-v2/issues/332)
* `service/rds`: Fix presign URL for same region ([#331](https://github.com/aws/aws-sdk-go-v2/pull/331)) 
  * Fixes RDS no-autopresign URL for same region issue for aws-sdk-go-v2. Solves the issue by making sure that the presigned URLs are not created, when the source and destination regions are the same. Added and updated the tests accordingly.
  * Fix [#271](https://github.com/aws/aws-sdk-go-v2/issues/271)
* `private/protocola/json/jsonutil`: Fix Unmarshal map[string]bool ([#320](https://github.com/aws/aws-sdk-go-v2/pull/320))
  * Fixes the JSON unmarshaling of maps of bools. The unmarshal case was missing the condition for bool value, in addition the bool pointer.
  * Fix [#319](https://github.com/aws/aws-sdk-go-v2/issues/319)

Release v0.9.0 (2019-05-28)
===

### Services
* Synced the V2 SDK with latest AWS service API definitions.
* Fixes [#304](https://github.com/aws/aws-sdk-go-v2/issues/304)
* Fixes [#295](https://github.com/aws/aws-sdk-go-v2/issues/295)

### SDK Breaking changes
This update includes multiple breaking changes to the SDK. These updates improve the SDK's usability, consistency.

#### Client type name
The API client type is renamed to `Client` for consistency, and remove stutter between package and client type name. Using Amazon S3 API client as an example, the `s3.S3` type is renamed to `s3.Client`.

#### New API operation response type
API operations' `Request.Send` method now returns a Response type for the specific operation. The Response type wraps the operation's Output parameter, and includes a method for the response's metadata such as RequestID. The Output type is an anonymous embedded field within the Output type. If your application was passing the Output value around you'll need to extract it directly, or pass the Response type instead.

#### New API operation paginator utility
This change removes the `Paginate` method from API operation Request types, (e.g. ListObjectsRequest). A new Paginator constructor is added that can be used to page these operations. To update your application to use the new pattern, where `Paginate` was being called, replace this with the Paginator type's constructor. The usage of the returned Paginator type is unchanged.

```go
req := svc.ListObjectsRequest(params)
p := req.Paginate()
```

Is updated to to use the Paginator constructor instead of Paginate method.

```go
req := svc.ListObjectsRequest(params)
p := s3.NewListObjectsPaginator(req)
```

#### Other changes
  * Standardizes API client package name to be based on the API model's `ServiceID`.
  * Standardizes API client operation input and output type names.
  * Removes `endpoints` package's service identifier constants. These values were unstable. Each API client package contains an `EndpointsID` constant that can be used for service specific endpoint lookup.
  * Fix API endpoint lookups to use the API's modeled `EndpointsID` (aka `enpdointPrefix`). Searching for API endpoints in the `endpoints` package should use the API client package's, `EndpointsID`.

### SDK Enhancements
*  Update CI tests to ensure all codegen changes are accounted for in PR ([#183](https://github.com/aws/aws-sdk-go-v2/issues/183))
  * Updates the CI tests to ensure that any code generation changes are accounted for in the PR, and that there were no mistaken changes made without also running code generation. This change should also help ensure that code generation order is stable, and there are no ordering issues with the SDK's codegen.
  * Related [aws/aws-sdk-go#1966](https://github.com/aws/aws-sdk-go/issues/1966)

### SDK Bugs
* `service/dynamodb/expression`: Fix Builder with KeyCondition example ([#306](https://github.com/aws/aws-sdk-go-v2/issues/306))
  * Fixes the ExampleBuilder_WithKeyCondition example to include the ExpressionAttributeNames member being set.
  * Fixes [#285](https://github.com/aws/aws-sdk-go-v2/issues/285)
* `aws/defaults`: Fix UserAgent execution environment key ([#307](https://github.com/aws/aws-sdk-go-v2/issues/307))
  * Fixes the SDK's UserAgent key for the execution environment.
  * Fixes [#276](https://github.com/aws/aws-sdk-go-v2/issues/276)
* `private/model/api`: Improve SDK API reference doc generation ([#309](https://github.com/aws/aws-sdk-go-v2/issues/309))
  * Improves the SDK's generated documentation for API client, operation, and types. This fixes several bugs in the doc generation causing poor formatting, an difficult to read reference documentation.
  * Fix [#308](https://github.com/aws/aws-sdk-go-v2/issues/308)
  * Related [aws/aws-sdk-go#2617](https://github.com/aws/aws-sdk-go/issues/2617)

Release v0.8.0 (2019-04-25)
===

### Services
* Synced the V2 SDK with latests AWS service API definitions.

### SDK Breaking changes
* Update SDK API operation request send to required Context ([#265](https://github.com/aws/aws-sdk-go-v2/pull/265))
  * Updates the SDK's API operation request Send method to require a context.Context value when called. This is done to encourage best applications to use Context for cancellation and request tracing.  Standardizing on this pattern will also help reduce code paths which accidentally do not have the Context causing the cancellation and tracing chain to be lost. Leading to difficult to trace down losses of cancellation and tracing within an application.
  * Fixes [#264](https://github.com/aws/aws-sdk-go-v2/pull/264)

### SDK Enhancements
* Update README.md for getting SDK without Go Modules
  * Updates the README.md with instructions how to get the SDK without Go Modules enabled, or using the SDK within a GOPATH with Go 1.11, and Go 1.12.
* Refactor SDK's integration tests to be code generated ([#283](https://github.com/aws/aws-sdk-go-v2/pull/283))
* `aws`: Add RequestThrottledException to set of throttled exceptions ([#292](https://github.com/aws/aws-sdk-go-v2/pull/292))
* `private/model/api`: Backfill authtype, STS and Cognito Identity ([#293](https://github.com/aws/aws-sdk-go-v2/pull/293))
  * Backfills the authtype=none modeled trait for STS and Cognito Identity services. This removes the in code customization for these two services' APIs that should not be signed.

### SDK Bugs
* Fix HTTP endpoint credential provider test for unresolved hosts ([#262](https://github.com/aws/aws-sdk-go-v2/pull/262))
  * Fixes the HTTP endpoint credential provider's tests to check for a host that resolves to no addresses.
* `example/service/s3/mockPaginator`: Update example to not use internal pkg ([#278](https://github.com/aws/aws-sdk-go-v2/pull/278))
  * Updates the SDK's S3 Mock Paginator example to not use internal SDK packages and instead use the SDK's provided defaults package for default configuration.
  * Fixes [#116](https://github.com/aws/aws-sdk-go-v2/issues/116)
* Cleanup go mod unused dependencies ([#284](https://github.com/aws/aws-sdk-go-v2/pull/284))
* `service/s3/s3manager`: Fix brittle Upload unit test ([#288](https://github.com/aws/aws-sdk-go-v2/pull/288))
* `aws/ec2metadata`: Fix EC2 Metadata client panic with debug logging ([#290](https://github.com/aws/aws-sdk-go-v2/pull/290))
  * Fixes a panic that could occur within the EC2 Metadata client when both AWS_EC2_METADATA_DISABLED env var is set and log level is LogDebugWithHTTPBody. The SDK's client response body debug functionality would panic because the Request.HTTPResponse value was not specified.
* `aws`: Fix RequestUserAgent test to be stable ([#289](https://github.com/aws/aws-sdk-go-v2/pull/289))
* `private/protocol/rest`: Trim space in header key and value ([#291](https://github.com/aws/aws-sdk-go-v2/pull/291))
  * Fixes a bug when using S3 metadata where metadata values with leading spaces would trigger request signature validation errors when the request is received by the service.


Release v0.7.0 (2019-01-03)
===

### Services
* Synced the V2 SDK with latests AWS service API definitions.

### SDK Enhancements
* deps: Update SDK to latest go-jmespath ([#254](https://github.com/aws/aws-sdk-go-v2/pull/254))

### SDK Bugs
* `internal/ini`: Fix bug on trimming rhs spaces closes ([#260](https://github.com/aws/aws-sdk-go-v2/pull/260))
  * Fixes a bug trimming RHS spaces not being read correctly from the ini file.
  * Fix [#259](https://github.com/aws/aws-sdk-go-v2/pull/259)

Release v0.6.0 (2018-12-03)
===

### Services
* Synced the V2 SDK with latests AWS service API definitions.

### SDK Bugs
* Updates the SDK's release tagging scheme to use `v0` until the v2 SDK reaches
	* General Availability (GA). This allows the SDK to be used with Go 1.11 modules. Post GA, v2 SDK's release tagging version will most likely follow a `v1.<x>.<y>` patter.
	* Fixes [#221](https://github.com/aws/aws-sdk-go-v2/issues/221)

Release v0.5.0 (2018-09-19)
===

### Services
* Synced the V2 SDK with latests AWS service API definitions.

### SDK Bugs
* Fix SDK Go 1.11 connection reset handling (#207)
	* Fixes how the SDK checks for connection reset errors when making API calls to be compatiable with Go 1.11.
* `aws/signer/v4`: Fix X-Amz-Content-Sha256 being in to query for presign (#188)
	* Fixes the bug which would allow the X-Amz-Content-Sha256 header to be promoted to the query string when presigning a S3 request.  This bug also was preventing users from setting their own sha256 value for a presigned URL. Presigned requests generated with the custom sha256 would of always failed with invalid signature.
	* Related to aws/aws-sdk-go#1974

### SDK Enhancements
* Cleanup SDK README and example documenation.
* `service/s3/s3manager`: Add doc for sequential download (#201)
	Adds documentation for downloading object sequentially with the S3 download manager.
* `aws/credentials`: Update Credentials cache to have less overhead (#184)
	* Updates the Credentials type's cache of the CredentialsValue to be synchronized with an atomic value in addition to the Mutex. This reduces the overhead applications will encounter when many concurrent API requests are being made.
	* Related to: aws/aws-sdk-go#1973
* `service/dynamodb/dynamodbattribute`: Add support for custom struct tag keys (#203)
	* Adds support for (un)marshaling Go types using custom struct tag keys. The new `MarshalOptions.TagKey` allows the user to specify the tag key to use when (un)marshaling struct fields.  Adds support for struct tags such as `yaml`, `toml`, etc. Support for these keys are in name only, and require the tag value format and values to be supported by the package's Marshalers.
* `internal/ini`: Add custom INI parser for shared config/credentials file (#209)
	* Related to: aws/aws-sdk-go#2024

Release v0.4.0 (2018-05-25)
===

### Services
* Synced the V2 SDK with latests AWS service API definitions.

### SDK Bugs
* `private/protocol/xml/xmlutil`: Fix XML unmarshaler not correctly unmarshaling list of timestamp values ([#166](https://github.com/aws/aws-sdk-go-v2/pull/166))
	* Fixes a bug in the XML unmarshaler that would incorrectly try to unmarshal "time.Time" parameters that did not have the struct tag type on them.
	* Related to [aws/aws-sdk-go#1894](https://github.com/aws/aws-sdk-go/pull/1894)
* `service/s3`: Fix typos for migrated S3 specific config options ([#173](https://github.com/aws/aws-sdk-go-v2/pull/173))
	* Updates the S3 specific config error messages to the correct fields.
* `aws/endpoints`: Fix SDK endpoint signing name resolution ([#181](https://github.com/aws/aws-sdk-go-v2/pull/181))
	* Fixes how the SDK derives service signing names. If the signing name is not modeled in the endpoints package the service will fallback to the signing name modeled in the service model.
	* Fix [#163](https://github.com/aws/aws-sdk-go-v2/pull/163)
	* Fix [#153](https://github.com/aws/aws-sdk-go-v2/pull/153)
	* Related to [aws/aws-sdk-go#1854](https://github.com/aws/aws-sdk-go/pull/1854)
* `service/s3`: remove SelectContent until EventStream supported ([#175](https://github.com/aws/aws-sdk-go-v2/pull/175])
	* S3's SelectContent API is not functional in the SDK yet, and was not supposed to be generated until EventStream support is available.
	* Related to [aws/aws-sdk-go#1941](https://github.com/aws/aws-sdk-go/pull/1941)

### SDK Enhancements
* `service/s3/s3manager/s3manageriface`: Add WithIterator to mock interface ([#156](https://github.com/aws/aws-sdk-go-v2/pull/156))
	* Updates the `DownloaderAPI` and `UploaderAPI` mocking interfaces to have parity with the concrete types.
	* Fixes [#155](https://github.com/aws/aws-sdk-go-v2/issues/155)


Release v0.3.0 (2018-03-08)
===

### Services
* Synced the V2 SDK with latests AWS service API definitions.

### Breaking Changes
* `private/mode/api`: Refactor service paginator helpers to iterator pattern ([#119](https://github.com/aws/aws-sdk-go-v2/pull/119))
	* Refactors the generated service client paginators to be an iterator pattern. This pattern improves usability while removing the need for callbacks.
	* See the linked PR for an example.
* `private/model/api`: Removes setter helpers from service API types ([#101](https://github.com/aws/aws-sdk-go-v2/pull/101))
	* Removes the setter helper methods from service API types. Removing clutter and noise from the API type's signature.
	* Based on feedback [#81][https://github.com/aws/aws-sdk-go-v2/issues/81]
* `aws`: Rename CanceledErrorCode to ErrCodeRequestCanceled ([#131](https://github.com/aws/aws-sdk-go-v2/pull/131))
	* Renames CanceledErrorCode to correct naming scheme.

### SDK Bugs
* `internal/awsutil`: Fix DeepEqual to consider string alias type equal to string ([#102](https://github.com/aws/aws-sdk-go-v2/pull/102))
	* Fixes SDK waiters not detecting the correct condition is met. [#92](https://github.com/aws/aws-sdk-go-v2/issues/92)
* `aws/external`: Fix EnvConfig misspelled container endpoint path getter ([#106](https://github.com/aws/aws-sdk-go-v2/pull/106))
	* This caused the type to not satisfy the ContainerCredentialsEndpointPathProvider interface.
	* Fixes [#105](https://github.com/aws/aws-sdk-go-v2/issues/105)
* `service/s3/s3crypto`: Fix S3Crypto's handling of TagLen ([#107](https://github.com/aws/aws-sdk-go-v2/pull/107))
	* Fixes the S3Crypto's handling of TagLen to only be set if present.
	* V2 Fix for [aws/aws-sdk-go#1742](https://github.com/aws/aws-sdk-go/issues/1742)
* `private/model/api`: Update SDK service client initialization documentation ([#141](https://github.com/aws/aws-sdk-go-v2/pull/141))
	* Updates the SDK's service initialization doc template to reflect the v2 SDK's configuration update change from v1.
	* Related to [#136](https://github.com/aws/aws-sdk-go-v2/issues/136)

### SDK Enhancements
* `aws`: Improve pagination unit tests ([#97](https://github.com/aws/aws-sdk-go-v2/pull/97))
	* V2 port of [aws/aws-sdk-go#1733](https://github.com/aws/aws-sdk-go/pull/1733)
* `aws/external`: Add example for shared config and static credential helper ([#109](https://github.com/aws/aws-sdk-go-v2/pull/109))
	* Adds examples for the  config helpers; WithSharedConfigProfile, WithCredentialsValue, WithMFATokenFunc.
* `private/model/api`: Add validation to prevent collision of api defintions ([#112](https://github.com/aws/aws-sdk-go-v2/pull/112)) 
	* V2 port of [aws/aws-sdk-go#1758](https://github.com/aws/aws-sdk-go/pull/1758)
* `aws/ec2metadata`: Add support for AWS_EC2_METADATA_DISABLED env var ([#128](https://github.com/aws/aws-sdk-go-v2/pull/128))
	* When this environment variable is set. The SDK's EC2 Metadata Client will not attempt to make requests. All requests made with the EC2 Metadata Client will fail.
	* V2 port of [aws/aws-sdk-go#1799](https://github.com/aws/aws-sdk-go/pull/1799)
* Add code of conduct ([#138](https://github.com/aws/aws-sdk-go-v2/pull/138))
* Update SDK README dep usage ([#140](https://github.com/aws/aws-sdk-go-v2/pull/140))

Release v0.2.0 (2018-01-15)
===

### Services
* Synced the V2 SDK with latests AWS service API definitions.

### SDK Bugs
* `service/s3/s3manager`: Fix Upload Manger's UploadInput fields ([#89](https://github.com/aws/aws-sdk-go-v2/pull/89))
	* Fixes [#88](https://github.com/aws/aws-sdk-go-v2/issues/88)
* `aws`: Fix Pagination handling of empty string NextToken ([#94](https://github.com/aws/aws-sdk-go-v2/pull/94))
	* Fixes [#84](https://github.com/aws/aws-sdk-go-v2/issues/84)


Release v0.1.0 (2017-12-21)
===

## What has changed?

Our focus for the 2.0 SDK is to improve the SDK’s development experience and performance, make the SDK easy to extend, and add new features. The changes made in the Developer Preview target the major pain points of configuring the SDK and using AWS service API calls. Check out the SDK for details on pending changes that are in development and designs we’re discussing.

The following are some of the larger changes included in the AWS SDK for Go 2.0 Developer Preview.

### SDK configuration

The 2.0 SDK simplifies how you configure the SDK's service clients by using a single `Config` type. This simplifies the `Session` and `Config` type interaction from the 1.x SDK. In addition, we’ve moved the service-specific configuration flags to the individual service client types. This reduces confusion about where service clients will use configuration members.

We added the external package to provide the utilities for you to use external configuration sources to populate the SDK's `Config` type. External sources include environmental variables, shared credentials file (`~/.aws/credentials`), and shared config file (`~/.aws/config`). By default, the 2.0 SDK will now automatically load configuration values from the shared config file. The external package also provides you with the tools to customize how the external sources are loaded and used to populate the `Config` type.

You can even customize external configuration sources to include your own custom sources, for example, JSON files or a custom file format.

Use `LoadDefaultAWSConfig` in the external package to create the default `Config` value, and load configuration values from the external configuration sources.

```go
cfg, err := external.LoadDefaultAWSConfig()
```

To specify the shared configuration profile load used in code, use the `WithSharedConfigProfile` helper passed into `LoadDefaultAWSConfig` with the profile name to use.

```go
cfg, err := external.LoadDefaultAWSConfig(
	external.WithSharedConfigProfile("gopher")
)
```

Once a `Config` value is returned by `LoadDefaultAWSConfig`, you can set or override configuration values by setting the fields on the `Config` struct, such as `Region`.

```go
cfg.Region = endpoints.UsWest2RegionID
```

Use the `cfg` value to provide the loaded configuration to new AWS service clients.

```go
svc := dynamodb.New(cfg)
```

### API request methods

The 2.0 SDK consolidates several ways to make an API call, providing a single request constructor for each API call. This means that your application will create an API request from input parameters, then send it. This enables you to optionally modify and configure how the request will be sent. This includes, but isn’t limited to, modifications such as setting the `Context` per request, adding request handlers, and enabling logging.

Each API request method is suffixed with `Request` and returns a typed value for the specific API request.

As an example, to use the Amazon Simple Storage Service GetObject API, the signature of the method is:

```go
func (c *S3) GetObjectRequest(*s3.GetObjectInput) *s3.GetObjectRequest
```

To use the GetObject API, we pass in the input parameters to the method, just like we would with the 1.x SDK. The 2.0 SDK's method will initialize a `GetObjectRequest` value that we can then use to send our request to Amazon S3.

```go
req := svc.GetObjectRequest(params)

// Optionally, set the context or other configuration for the request to use
req.SetContext(ctx)

// Send the request and get the response
resp, err := req.Send()
```

### API enumerations

The 2.0 SDK uses typed enumeration values for all API enumeration fields. This change provides you with the type safety that you and your IDE can leverage to discover which enumeration values are valid for particular fields. Typed enumeration values also provide a stronger type safety story for your application than using string literals directly. The 2.0 SDK uses string aliases for each enumeration type. The SDK also generates typed constants for each enumeration value. This change removes the need for enumeration API fields to be pointers, as a zero-value enumeration always means the field isn’t set.

For example, the Amazon Simple Storage Service PutObject API has a field, `ACL ObjectCannedACL`. An `ObjectCannedACL` string alias is defined within the s3 package, and each enumeration value is also defined as a typed constant. In this example, we want to use the typed enumeration values to set an uploaded object to have an ACL of `public-read`. The constant that the SDK provides for this enumeration value is `ObjectCannedACLPublicRead`.

```go
svc.PutObjectRequest(&s3.PutObjectInput{
	Bucket: aws.String("myBucket"),
	Key:    aws.String("myKey"),
	ACL:    s3.ObjectCannedACLPublicRead,
	Body:   body,
})
```

### API slice and map elements

The 2.0 SDK removes the need to convert slice and map elements from values to pointers for API calls. This will reduce the overhead of needing to use fields with a type, such as `[]string`, in API calls. The 1.x SDK's pattern of using pointer types for all slice and map elements was a significant pain point for users, requiring them to convert between the types. The 2.0 SDK does away with the pointer types for slices and maps, using value types instead.

The following example shows how value types for the Amazon Simple Queue Service AddPermission API's `AWSAccountIds` and `Actions` member slices are set.

```go
svc := sqs.New(cfg)

svc.AddPermission(&sqs.AddPermissionInput{
	AWSAcountIds: []string{
		"123456789",
	},
	Actions: []string{
		"SendMessage",
	},

	Label:    aws.String("MessageSender"),
	QueueUrl: aws.String(queueURL)
})
```

