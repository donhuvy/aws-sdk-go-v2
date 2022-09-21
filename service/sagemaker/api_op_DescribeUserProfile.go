// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a user profile. For more information, see CreateUserProfile.
func (c *Client) DescribeUserProfile(ctx context.Context, params *DescribeUserProfileInput, optFns ...func(*Options)) (*DescribeUserProfileOutput, error) {
	if params == nil {
		params = &DescribeUserProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeUserProfile", params, optFns, c.addOperationDescribeUserProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeUserProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUserProfileInput struct {

	// The domain ID.
	//
	// This member is required.
	DomainId *string

	// The user profile name. This value is not case sensitive.
	//
	// This member is required.
	UserProfileName *string

	noSmithyDocumentSerde
}

type DescribeUserProfileOutput struct {

	// The creation time.
	CreationTime *time.Time

	// The ID of the domain that contains the profile.
	DomainId *string

	// The failure reason.
	FailureReason *string

	// The ID of the user's profile in the Amazon Elastic File System (EFS) volume.
	HomeEfsFileSystemUid *string

	// The last modified time.
	LastModifiedTime *time.Time

	// The IAM Identity Center user identifier.
	SingleSignOnUserIdentifier *string

	// The IAM Identity Center user value.
	SingleSignOnUserValue *string

	// The status.
	Status types.UserProfileStatus

	// The user profile Amazon Resource Name (ARN).
	UserProfileArn *string

	// The user profile name.
	UserProfileName *string

	// A collection of settings.
	UserSettings *types.UserSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeUserProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeUserProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeUserProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeUserProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUserProfile(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeUserProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeUserProfile",
	}
}
