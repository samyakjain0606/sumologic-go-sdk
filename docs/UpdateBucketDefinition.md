# UpdateBucketDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationName** | Pointer to **string** | Name of the S3 data forwarding destination. | [optional] 
**Description** | Pointer to **string** | Description of the S3 data forwarding destination. | [optional] 
**AuthenticationMode** | **string** | AWS IAM authentication method used for access. Possible values are: 1. &#x60;AccessKey&#x60; 2. &#x60;RoleBased&#x60; | 
**AccessKeyId** | Pointer to **string** | The AWS Access ID to access the S3 bucket. | [optional] 
**SecretAccessKey** | Pointer to **string** | The AWS Secret Key to access the S3 bucket. | [optional] 
**RoleArn** | Pointer to **string** | The AWS Role ARN to access the S3 bucket. | [optional] 
**Region** | Pointer to **string** | The region where the S3 bucket is located. | [optional] 
**Encrypted** | Pointer to **bool** | Enable S3 server-side encryption. | [optional] 
**Enabled** | Pointer to **bool** | True if the destination is Active. | [optional] 

## Methods

### NewUpdateBucketDefinition

`func NewUpdateBucketDefinition(authenticationMode string, ) *UpdateBucketDefinition`

NewUpdateBucketDefinition instantiates a new UpdateBucketDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBucketDefinitionWithDefaults

`func NewUpdateBucketDefinitionWithDefaults() *UpdateBucketDefinition`

NewUpdateBucketDefinitionWithDefaults instantiates a new UpdateBucketDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationName

`func (o *UpdateBucketDefinition) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *UpdateBucketDefinition) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *UpdateBucketDefinition) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.

### HasDestinationName

`func (o *UpdateBucketDefinition) HasDestinationName() bool`

HasDestinationName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateBucketDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateBucketDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateBucketDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateBucketDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthenticationMode

`func (o *UpdateBucketDefinition) GetAuthenticationMode() string`

GetAuthenticationMode returns the AuthenticationMode field if non-nil, zero value otherwise.

### GetAuthenticationModeOk

`func (o *UpdateBucketDefinition) GetAuthenticationModeOk() (*string, bool)`

GetAuthenticationModeOk returns a tuple with the AuthenticationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMode

`func (o *UpdateBucketDefinition) SetAuthenticationMode(v string)`

SetAuthenticationMode sets AuthenticationMode field to given value.


### GetAccessKeyId

`func (o *UpdateBucketDefinition) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *UpdateBucketDefinition) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *UpdateBucketDefinition) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *UpdateBucketDefinition) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *UpdateBucketDefinition) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *UpdateBucketDefinition) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *UpdateBucketDefinition) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *UpdateBucketDefinition) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### GetRoleArn

`func (o *UpdateBucketDefinition) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *UpdateBucketDefinition) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *UpdateBucketDefinition) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *UpdateBucketDefinition) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetRegion

`func (o *UpdateBucketDefinition) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *UpdateBucketDefinition) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *UpdateBucketDefinition) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *UpdateBucketDefinition) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetEncrypted

`func (o *UpdateBucketDefinition) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *UpdateBucketDefinition) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *UpdateBucketDefinition) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *UpdateBucketDefinition) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateBucketDefinition) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateBucketDefinition) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateBucketDefinition) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateBucketDefinition) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


