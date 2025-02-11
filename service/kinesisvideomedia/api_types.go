// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideomedia

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// Identifies the chunk on the Kinesis video stream where you want the GetMedia
// API to start returning media data. You have the following options to identify
// the starting chunk:
//
//    * Choose the latest (or oldest) chunk.
//
//    * Identify a specific chunk. You can identify a specific chunk either
//    by providing a fragment number or timestamp (server or producer).
//
//    * Each chunk's metadata includes a continuation token as a Matroska (MKV)
//    tag (AWS_KINESISVIDEO_CONTINUATION_TOKEN). If your previous GetMedia request
//    terminated, you can use this tag value in your next GetMedia request.
//    The API then starts returning chunks starting where the last API ended.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-media-2017-09-30/StartSelector
type StartSelector struct {
	_ struct{} `type:"structure"`

	// Specifies the fragment number from where you want the GetMedia API to start
	// returning the fragments.
	AfterFragmentNumber *string `min:"1" type:"string"`

	// Continuation token that Kinesis Video Streams returned in the previous GetMedia
	// response. The GetMedia API then starts with the chunk identified by the continuation
	// token.
	ContinuationToken *string `min:"1" type:"string"`

	// Identifies the fragment on the Kinesis video stream where you want to start
	// getting the data from.
	//
	//    * NOW - Start with the latest chunk on the stream.
	//
	//    * EARLIEST - Start with earliest available chunk on the stream.
	//
	//    * FRAGMENT_NUMBER - Start with the chunk after a specific fragment. You
	//    must also specify the AfterFragmentNumber parameter.
	//
	//    * PRODUCER_TIMESTAMP or SERVER_TIMESTAMP - Start with the chunk containing
	//    a fragment with the specified producer or server timestamp. You specify
	//    the timestamp by adding StartTimestamp.
	//
	//    * CONTINUATION_TOKEN - Read using the specified continuation token.
	//
	// If you choose the NOW, EARLIEST, or CONTINUATION_TOKEN as the startSelectorType,
	// you don't provide any additional information in the startSelector.
	//
	// StartSelectorType is a required field
	StartSelectorType StartSelectorType `type:"string" required:"true" enum:"true"`

	// A timestamp value. This value is required if you choose the PRODUCER_TIMESTAMP
	// or the SERVER_TIMESTAMP as the startSelectorType. The GetMedia API then starts
	// with the chunk containing the fragment that has the specified timestamp.
	StartTimestamp *time.Time `type:"timestamp" timestampFormat:"unix"`
}

// String returns the string representation
func (s StartSelector) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartSelector) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartSelector"}
	if s.AfterFragmentNumber != nil && len(*s.AfterFragmentNumber) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AfterFragmentNumber", 1))
	}
	if s.ContinuationToken != nil && len(*s.ContinuationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContinuationToken", 1))
	}
	if len(s.StartSelectorType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("StartSelectorType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartSelector) MarshalFields(e protocol.FieldEncoder) error {
	if s.AfterFragmentNumber != nil {
		v := *s.AfterFragmentNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AfterFragmentNumber", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ContinuationToken != nil {
		v := *s.ContinuationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ContinuationToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.StartSelectorType) > 0 {
		v := s.StartSelectorType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StartSelectorType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StartTimestamp != nil {
		v := *s.StartTimestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StartTimestamp", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	return nil
}
