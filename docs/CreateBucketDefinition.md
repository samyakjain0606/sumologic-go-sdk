# CreateBucketDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationName** | **string** | Name of the S3 data forwarding destination. | 
**Description** | Pointer to **string** | Description of the S3 data forwarding destination. | [optional] 
**AuthenticationMode** | **string** | AWS IAM authentication method used for access. Possible values are: 1. &#x60;AccessKey&#x60; 2. &#x60;RoleBased&#x60; | 
**AccessKeyId** | Pointer to **string** | The AWS Access ID to access the S3 bucket. | [optional] 
**SecretAccessKey** | Pointer to **string** | The AWS Secret Key to access the S3 bucket. | [optional] 
**RoleArn** | Pointer to **string** | The AWS Role ARN to access the S3 bucket. | [optional] 
**Region** | Pointer to **string** | The region where the S3 bucket is located. | [optional] 
**Encrypted** | Pointer to **bool** | Enable S3 server-side encryption. | [optional] 
**Enabled** | Pointer to **bool** | True if the destination is Active. | [optional] 
**BucketName** | **string** | The name of the Amazon S3 bucket. | 

## Methods

### NewCreateBucketDefinition

`func NewCreateBucketDefinition(destinationName string, authenticationMode string, bucketName string, ) *CreateBucketDefinition`

NewCreateBucketDefinition instantiates a new CreateBucketDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBucketDefinitionWithDefaults

`func NewCreateBucketDefinitionWithDefaults() *CreateBucketDefinition`

NewCreateBucketDefinitionWithDefaults instantiates a new CreateBucketDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationName

`func (o *CreateBucketDefinition) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *CreateBucketDefinition) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *CreateBucketDefinition) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.


### GetDescription

`func (o *CreateBucketDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateBucketDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateBucketDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateBucketDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthenticationMode

`func (o *CreateBucketDefinition) GetAuthenticationMode() string`

GetAuthenticationMode returns the AuthenticationMode field if non-nil, zero value otherwise.

### GetAuthenticationModeOk

`func (o *CreateBucketDefinition) GetAuthenticationModeOk() (*string, bool)`

GetAuthenticationModeOk returns a tuple with the AuthenticationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMode

`func (o *CreateBucketDefinition) SetAuthenticationMode(v string)`

SetAuthenticationMode sets AuthenticationMode field to given value.


### GetAccessKeyId

`func (o *CreateBucketDefinition) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *CreateBucketDefinition) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *CreateBucketDefinition) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *CreateBucketDefinition) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *CreateBucketDefinition) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *CreateBucketDefinition) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *CreateBucketDefinition) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *CreateBucketDefinition) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### GetRoleArn

`func (o *CreateBucketDefinition) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *CreateBucketDefinition) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *CreateBucketDefinition) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *CreateBucketDefinition) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetRegion

`func (o *CreateBucketDefinition) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateBucketDefinition) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateBucketDefinition) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateBucketDefinition) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetEncrypted

`func (o *CreateBucketDefinition) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *CreateBucketDefinition) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *CreateBucketDefinition) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *CreateBucketDefinition) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateBucketDefinition) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateBucketDefinition) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateBucketDefinition) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateBucketDefinition) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetBucketName

`func (o *CreateBucketDefinition) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *CreateBucketDefinition) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *CreateBucketDefinition) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


