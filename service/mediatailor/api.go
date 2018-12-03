// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediatailor

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

const opDeletePlaybackConfiguration = "DeletePlaybackConfiguration"

// DeletePlaybackConfigurationRequest is a API request type for the DeletePlaybackConfiguration API operation.
type DeletePlaybackConfigurationRequest struct {
	*aws.Request
	Input *DeletePlaybackConfigurationInput
	Copy  func(*DeletePlaybackConfigurationInput) DeletePlaybackConfigurationRequest
}

// Send marshals and sends the DeletePlaybackConfiguration API request.
func (r DeletePlaybackConfigurationRequest) Send() (*DeletePlaybackConfigurationOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*DeletePlaybackConfigurationOutput), nil
}

// DeletePlaybackConfigurationRequest returns a request value for making API operation for
// AWS MediaTailor.
//
// Deletes the configuration for the specified name.
//
//    // Example sending a request using the DeletePlaybackConfigurationRequest method.
//    req := client.DeletePlaybackConfigurationRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/DeletePlaybackConfiguration
func (c *MediaTailor) DeletePlaybackConfigurationRequest(input *DeletePlaybackConfigurationInput) DeletePlaybackConfigurationRequest {
	op := &aws.Operation{
		Name:       opDeletePlaybackConfiguration,
		HTTPMethod: "DELETE",
		HTTPPath:   "/playbackConfiguration/{Name}",
	}

	if input == nil {
		input = &DeletePlaybackConfigurationInput{}
	}

	output := &DeletePlaybackConfigurationOutput{}
	req := c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	output.responseMetadata = aws.Response{Request: req}

	return DeletePlaybackConfigurationRequest{Request: req, Input: input, Copy: c.DeletePlaybackConfigurationRequest}
}

const opGetPlaybackConfiguration = "GetPlaybackConfiguration"

// GetPlaybackConfigurationRequest is a API request type for the GetPlaybackConfiguration API operation.
type GetPlaybackConfigurationRequest struct {
	*aws.Request
	Input *GetPlaybackConfigurationInput
	Copy  func(*GetPlaybackConfigurationInput) GetPlaybackConfigurationRequest
}

// Send marshals and sends the GetPlaybackConfiguration API request.
func (r GetPlaybackConfigurationRequest) Send() (*GetPlaybackConfigurationOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*GetPlaybackConfigurationOutput), nil
}

// GetPlaybackConfigurationRequest returns a request value for making API operation for
// AWS MediaTailor.
//
// Returns the configuration for the specified name.
//
//    // Example sending a request using the GetPlaybackConfigurationRequest method.
//    req := client.GetPlaybackConfigurationRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/GetPlaybackConfiguration
func (c *MediaTailor) GetPlaybackConfigurationRequest(input *GetPlaybackConfigurationInput) GetPlaybackConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetPlaybackConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/playbackConfiguration/{Name}",
	}

	if input == nil {
		input = &GetPlaybackConfigurationInput{}
	}

	output := &GetPlaybackConfigurationOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return GetPlaybackConfigurationRequest{Request: req, Input: input, Copy: c.GetPlaybackConfigurationRequest}
}

const opListPlaybackConfigurations = "ListPlaybackConfigurations"

// ListPlaybackConfigurationsRequest is a API request type for the ListPlaybackConfigurations API operation.
type ListPlaybackConfigurationsRequest struct {
	*aws.Request
	Input *ListPlaybackConfigurationsInput
	Copy  func(*ListPlaybackConfigurationsInput) ListPlaybackConfigurationsRequest
}

// Send marshals and sends the ListPlaybackConfigurations API request.
func (r ListPlaybackConfigurationsRequest) Send() (*ListPlaybackConfigurationsOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*ListPlaybackConfigurationsOutput), nil
}

// ListPlaybackConfigurationsRequest returns a request value for making API operation for
// AWS MediaTailor.
//
// Returns a list of the configurations defined in AWS Elemental MediaTailor.
// You can specify a max number of configurations to return at a time. The default
// max is 50. Results are returned in pagefuls. If AWS Elemental MediaTailor
// has more configurations than the specified max, it provides parameters in
// the response that you can use to retrieve the next pageful.
//
//    // Example sending a request using the ListPlaybackConfigurationsRequest method.
//    req := client.ListPlaybackConfigurationsRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/ListPlaybackConfigurations
func (c *MediaTailor) ListPlaybackConfigurationsRequest(input *ListPlaybackConfigurationsInput) ListPlaybackConfigurationsRequest {
	op := &aws.Operation{
		Name:       opListPlaybackConfigurations,
		HTTPMethod: "GET",
		HTTPPath:   "/playbackConfigurations",
	}

	if input == nil {
		input = &ListPlaybackConfigurationsInput{}
	}

	output := &ListPlaybackConfigurationsOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return ListPlaybackConfigurationsRequest{Request: req, Input: input, Copy: c.ListPlaybackConfigurationsRequest}
}

const opPutPlaybackConfiguration = "PutPlaybackConfiguration"

// PutPlaybackConfigurationRequest is a API request type for the PutPlaybackConfiguration API operation.
type PutPlaybackConfigurationRequest struct {
	*aws.Request
	Input *PutPlaybackConfigurationInput
	Copy  func(*PutPlaybackConfigurationInput) PutPlaybackConfigurationRequest
}

// Send marshals and sends the PutPlaybackConfiguration API request.
func (r PutPlaybackConfigurationRequest) Send() (*PutPlaybackConfigurationOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*PutPlaybackConfigurationOutput), nil
}

// PutPlaybackConfigurationRequest returns a request value for making API operation for
// AWS MediaTailor.
//
// Adds a new configuration to AWS Elemental MediaTailor.
//
//    // Example sending a request using the PutPlaybackConfigurationRequest method.
//    req := client.PutPlaybackConfigurationRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/PutPlaybackConfiguration
func (c *MediaTailor) PutPlaybackConfigurationRequest(input *PutPlaybackConfigurationInput) PutPlaybackConfigurationRequest {
	op := &aws.Operation{
		Name:       opPutPlaybackConfiguration,
		HTTPMethod: "PUT",
		HTTPPath:   "/playbackConfiguration",
	}

	if input == nil {
		input = &PutPlaybackConfigurationInput{}
	}

	output := &PutPlaybackConfigurationOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return PutPlaybackConfigurationRequest{Request: req, Input: input, Copy: c.PutPlaybackConfigurationRequest}
}

// The configuration for using a content delivery network (CDN), like Amazon
// CloudFront, for content and ad segment management.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/CdnConfiguration
type CdnConfiguration struct {
	_ struct{} `type:"structure"`

	// A non-default content delivery network (CDN) to serve ad segments. By default,
	// AWS Elemental MediaTailor uses Amazon CloudFront with default cache settings
	// as its CDN for ad segments. To set up an alternate CDN, create a rule in
	// your CDN for the following origin: ads.mediatailor.<region>.amazonaws.com.
	// Then specify the rule's name in this AdSegmentUrlPrefix. When AWS Elemental
	// MediaTailor serves a manifest, it reports your CDN as the source for ad segments.
	AdSegmentUrlPrefix *string `type:"string"`

	// A content delivery network (CDN) to cache content segments, so that content
	// requests don’t always have to go to the origin server. First, create a rule
	// in your CDN for the content segment origin server. Then specify the rule's
	// name in this ContentSegmentUrlPrefix. When AWS Elemental MediaTailor serves
	// a manifest, it reports your CDN as the source for content segments.
	ContentSegmentUrlPrefix *string `type:"string"`
}

// String returns the string representation
func (s CdnConfiguration) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CdnConfiguration) GoString() string {
	return s.String()
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CdnConfiguration) MarshalFields(e protocol.FieldEncoder) error {
	if s.AdSegmentUrlPrefix != nil {
		v := *s.AdSegmentUrlPrefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AdSegmentUrlPrefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ContentSegmentUrlPrefix != nil {
		v := *s.ContentSegmentUrlPrefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ContentSegmentUrlPrefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The configuration object for dash content.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/DashConfiguration
type DashConfiguration struct {
	_ struct{} `type:"structure"`

	// The URL that is used to initiate a playback session for devices that support
	// DASH.
	ManifestEndpointPrefix *string `type:"string"`
}

// String returns the string representation
func (s DashConfiguration) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DashConfiguration) GoString() string {
	return s.String()
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DashConfiguration) MarshalFields(e protocol.FieldEncoder) error {
	if s.ManifestEndpointPrefix != nil {
		v := *s.ManifestEndpointPrefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ManifestEndpointPrefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/DeletePlaybackConfigurationRequest
type DeletePlaybackConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Name is a required field
	Name *string `location:"uri" locationName:"Name" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePlaybackConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeletePlaybackConfigurationInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePlaybackConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePlaybackConfigurationInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeletePlaybackConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/DeletePlaybackConfigurationOutput
type DeletePlaybackConfigurationOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response
}

// String returns the string representation
func (s DeletePlaybackConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeletePlaybackConfigurationOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s DeletePlaybackConfigurationOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeletePlaybackConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/GetPlaybackConfigurationRequest
type GetPlaybackConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Name is a required field
	Name *string `location:"uri" locationName:"Name" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPlaybackConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetPlaybackConfigurationInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPlaybackConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPlaybackConfigurationInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPlaybackConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/GetPlaybackConfigurationResponse
type GetPlaybackConfigurationOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// The URL for the ad decision server (ADS). This includes the specification
	// of static parameters and placeholders for dynamic parameters. AWS Elemental
	// MediaTailor substitutes player-specific and session-specific parameters as
	// needed when calling the ADS. Alternately, for testing, you can provide a
	// static VAST URL. The maximum length is 25000 characters.
	AdDecisionServerUrl *string `type:"string"`

	// The configuration for using a content delivery network (CDN), like Amazon
	// CloudFront, for content and ad segment management.
	CdnConfiguration *CdnConfiguration `type:"structure"`

	// The configuration object for dash content.
	DashConfiguration *DashConfiguration `type:"structure"`

	// The configuration for HLS content.
	HlsConfiguration *HlsConfiguration `type:"structure"`

	// The identifier for the configuration.
	Name *string `type:"string"`

	// The URL that the player accesses to get a manifest from AWS Elemental MediaTailor.
	// This session will use server-side reporting.
	PlaybackEndpointPrefix *string `type:"string"`

	// The URL that the player uses to initialize a session that uses client-side
	// reporting.
	SessionInitializationEndpointPrefix *string `type:"string"`

	// URL for a high-quality video asset to transcode and use to fill in time that's
	// not used by ads. AWS Elemental MediaTailor shows the slate to fill in gaps
	// in media content. Configuring the slate is optional for non-VPAID configurations.
	// For VPAID, the slate is required because AWS Elemental MediaTailor provides
	// it in the slots designated for dynamic ad content. The slate must be a high-quality
	// asset that contains both audio and video.
	SlateAdUrl *string `type:"string"`

	// Associate this playbackConfiguration with a custom transcode profile, overriding
	// MediaTailor's dynamic transcoding defaults. Do not include this field if
	// you have not setup custom profiles with the MediaTailor service team.
	TranscodeProfileName *string `type:"string"`

	// The URL prefix for the master playlist for the stream, minus the asset ID.
	// The maximum length is 512 characters.
	VideoContentSourceUrl *string `type:"string"`
}

// String returns the string representation
func (s GetPlaybackConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetPlaybackConfigurationOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s GetPlaybackConfigurationOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPlaybackConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AdDecisionServerUrl != nil {
		v := *s.AdDecisionServerUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AdDecisionServerUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CdnConfiguration != nil {
		v := s.CdnConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "CdnConfiguration", v, metadata)
	}
	if s.DashConfiguration != nil {
		v := s.DashConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DashConfiguration", v, metadata)
	}
	if s.HlsConfiguration != nil {
		v := s.HlsConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "HlsConfiguration", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PlaybackEndpointPrefix != nil {
		v := *s.PlaybackEndpointPrefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PlaybackEndpointPrefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SessionInitializationEndpointPrefix != nil {
		v := *s.SessionInitializationEndpointPrefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SessionInitializationEndpointPrefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SlateAdUrl != nil {
		v := *s.SlateAdUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SlateAdUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TranscodeProfileName != nil {
		v := *s.TranscodeProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TranscodeProfileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VideoContentSourceUrl != nil {
		v := *s.VideoContentSourceUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VideoContentSourceUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The configuration for HLS content.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/HlsConfiguration
type HlsConfiguration struct {
	_ struct{} `type:"structure"`

	// The URL that is used to initiate a playback session for devices that support
	// Apple HLS. The session uses server-side reporting.
	ManifestEndpointPrefix *string `type:"string"`
}

// String returns the string representation
func (s HlsConfiguration) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s HlsConfiguration) GoString() string {
	return s.String()
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s HlsConfiguration) MarshalFields(e protocol.FieldEncoder) error {
	if s.ManifestEndpointPrefix != nil {
		v := *s.ManifestEndpointPrefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ManifestEndpointPrefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/ListPlaybackConfigurationsRequest
type ListPlaybackConfigurationsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `location:"querystring" locationName:"MaxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s ListPlaybackConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListPlaybackConfigurationsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPlaybackConfigurationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPlaybackConfigurationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPlaybackConfigurationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/ListPlaybackConfigurationsResponse
type ListPlaybackConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// Array of playback configurations. This may be all of the available configurations
	// or a subset, depending on the settings you provide and on the total number
	// of configurations stored.
	Items []PlaybackConfiguration `type:"list"`

	// Pagination token returned by the GET list request when results overrun the
	// meximum allowed. Use the token to fetch the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListPlaybackConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListPlaybackConfigurationsOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s ListPlaybackConfigurationsOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPlaybackConfigurationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Items) > 0 {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Items", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/PlaybackConfiguration
type PlaybackConfiguration struct {
	_ struct{} `type:"structure"`

	AdDecisionServerUrl *string `type:"string"`

	// The configuration for using a content delivery network (CDN), like Amazon
	// CloudFront, for content and ad segment management.
	CdnConfiguration *CdnConfiguration `type:"structure"`

	Name *string `type:"string"`

	SlateAdUrl *string `type:"string"`

	VideoContentSourceUrl *string `type:"string"`
}

// String returns the string representation
func (s PlaybackConfiguration) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PlaybackConfiguration) GoString() string {
	return s.String()
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PlaybackConfiguration) MarshalFields(e protocol.FieldEncoder) error {
	if s.AdDecisionServerUrl != nil {
		v := *s.AdDecisionServerUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AdDecisionServerUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CdnConfiguration != nil {
		v := s.CdnConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "CdnConfiguration", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SlateAdUrl != nil {
		v := *s.SlateAdUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SlateAdUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VideoContentSourceUrl != nil {
		v := *s.VideoContentSourceUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VideoContentSourceUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/PutPlaybackConfigurationRequest
type PutPlaybackConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The URL for the ad decision server (ADS). This includes the specification
	// of static parameters and placeholders for dynamic parameters. AWS Elemental
	// MediaTailor substitutes player-specific and session-specific parameters as
	// needed when calling the ADS. Alternately, for testing you can provide a static
	// VAST URL. The maximum length is 25000 characters.
	AdDecisionServerUrl *string `type:"string"`

	// The configuration for using a content delivery network (CDN), like Amazon
	// CloudFront, for content and ad segment management.
	CdnConfiguration *CdnConfiguration `type:"structure"`

	// The identifier for the configuration.
	Name *string `type:"string"`

	// The URL for a high-quality video asset to transcode and use to fill in time
	// that's not used by ads. AWS Elemental MediaTailor shows the slate to fill
	// in gaps in media content. Configuring the slate is optional for non-VPAID
	// configurations. For VPAID, the slate is required because AWS Elemental MediaTailor
	// provides it in the slots that are designated for dynamic ad content. The
	// slate must be a high-quality asset that contains both audio and video.
	SlateAdUrl *string `type:"string"`

	// Associate this playbackConfiguration with a custom transcode profile, overriding
	// MediaTailor's dynamic transcoding defaults. Do not include this field if
	// you have not setup custom profiles with the MediaTailor service team.
	TranscodeProfileName *string `type:"string"`

	// The URL prefix for the master playlist for the stream, minus the asset ID.
	// The maximum length is 512 characters.
	VideoContentSourceUrl *string `type:"string"`
}

// String returns the string representation
func (s PutPlaybackConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutPlaybackConfigurationInput) GoString() string {
	return s.String()
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutPlaybackConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.AdDecisionServerUrl != nil {
		v := *s.AdDecisionServerUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AdDecisionServerUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CdnConfiguration != nil {
		v := s.CdnConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "CdnConfiguration", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SlateAdUrl != nil {
		v := *s.SlateAdUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SlateAdUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TranscodeProfileName != nil {
		v := *s.TranscodeProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TranscodeProfileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VideoContentSourceUrl != nil {
		v := *s.VideoContentSourceUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VideoContentSourceUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/PutPlaybackConfigurationResponse
type PutPlaybackConfigurationOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	AdDecisionServerUrl *string `type:"string"`

	// The configuration for using a content delivery network (CDN), like Amazon
	// CloudFront, for content and ad segment management.
	CdnConfiguration *CdnConfiguration `type:"structure"`

	// The configuration object for dash content.
	DashConfiguration *DashConfiguration `type:"structure"`

	// The configuration for HLS content.
	HlsConfiguration *HlsConfiguration `type:"structure"`

	Name *string `type:"string"`

	PlaybackEndpointPrefix *string `type:"string"`

	SessionInitializationEndpointPrefix *string `type:"string"`

	SlateAdUrl *string `type:"string"`

	TranscodeProfileName *string `type:"string"`

	VideoContentSourceUrl *string `type:"string"`
}

// String returns the string representation
func (s PutPlaybackConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutPlaybackConfigurationOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s PutPlaybackConfigurationOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutPlaybackConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AdDecisionServerUrl != nil {
		v := *s.AdDecisionServerUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AdDecisionServerUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CdnConfiguration != nil {
		v := s.CdnConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "CdnConfiguration", v, metadata)
	}
	if s.DashConfiguration != nil {
		v := s.DashConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DashConfiguration", v, metadata)
	}
	if s.HlsConfiguration != nil {
		v := s.HlsConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "HlsConfiguration", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PlaybackEndpointPrefix != nil {
		v := *s.PlaybackEndpointPrefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PlaybackEndpointPrefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SessionInitializationEndpointPrefix != nil {
		v := *s.SessionInitializationEndpointPrefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SessionInitializationEndpointPrefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SlateAdUrl != nil {
		v := *s.SlateAdUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SlateAdUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TranscodeProfileName != nil {
		v := *s.TranscodeProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TranscodeProfileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VideoContentSourceUrl != nil {
		v := *s.VideoContentSourceUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VideoContentSourceUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}
