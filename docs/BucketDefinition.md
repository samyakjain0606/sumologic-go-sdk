# BucketDefinition

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
**CreatedAt** | **time.Time** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**CreatedBy** | **string** | Identifier of the user who created the resource. | 
**ModifiedAt** | **time.Time** | Last modification timestamp in UTC. | 
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | 
**Id** | **string** | The unique identifier of the data forwarding destination. | 
**InvalidatedBySystem** | Pointer to **bool** | True if invalidated by the system. | [optional] 

## Methods

### NewBucketDefinition

`func NewBucketDefinition(destinationName string, authenticationMode string, bucketName string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, id string, ) *BucketDefinition`

NewBucketDefinition instantiates a new BucketDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketDefinitionWithDefaults

`func NewBucketDefinitionWithDefaults() *BucketDefinition`

NewBucketDefinitionWithDefaults instantiates a new BucketDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationName

`func (o *BucketDefinition) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *BucketDefinition) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *BucketDefinition) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.


### GetDescription

`func (o *BucketDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BucketDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BucketDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BucketDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthenticationMode

`func (o *BucketDefinition) GetAuthenticationMode() string`

GetAuthenticationMode returns the AuthenticationMode field if non-nil, zero value otherwise.

### GetAuthenticationModeOk

`func (o *BucketDefinition) GetAuthenticationModeOk() (*string, bool)`

GetAuthenticationModeOk returns a tuple with the AuthenticationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMode

`func (o *BucketDefinition) SetAuthenticationMode(v string)`

SetAuthenticationMode sets AuthenticationMode field to given value.


### GetAccessKeyId

`func (o *BucketDefinition) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *BucketDefinition) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *BucketDefinition) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *BucketDefinition) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *BucketDefinition) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *BucketDefinition) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *BucketDefinition) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *BucketDefinition) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### GetRoleArn

`func (o *BucketDefinition) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *BucketDefinition) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *BucketDefinition) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *BucketDefinition) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetRegion

`func (o *BucketDefinition) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *BucketDefinition) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *BucketDefinition) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *BucketDefinition) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetEncrypted

`func (o *BucketDefinition) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *BucketDefinition) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *BucketDefinition) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *BucketDefinition) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetEnabled

`func (o *BucketDefinition) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BucketDefinition) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BucketDefinition) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BucketDefinition) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetBucketName

`func (o *BucketDefinition) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *BucketDefinition) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *BucketDefinition) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetCreatedAt

`func (o *BucketDefinition) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BucketDefinition) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BucketDefinition) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *BucketDefinition) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BucketDefinition) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BucketDefinition) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *BucketDefinition) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BucketDefinition) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BucketDefinition) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *BucketDefinition) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BucketDefinition) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BucketDefinition) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetId

`func (o *BucketDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BucketDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BucketDefinition) SetId(v string)`

SetId sets Id field to given value.


### GetInvalidatedBySystem

`func (o *BucketDefinition) GetInvalidatedBySystem() bool`

GetInvalidatedBySystem returns the InvalidatedBySystem field if non-nil, zero value otherwise.

### GetInvalidatedBySystemOk

`func (o *BucketDefinition) GetInvalidatedBySystemOk() (*bool, bool)`

GetInvalidatedBySystemOk returns a tuple with the InvalidatedBySystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidatedBySystem

`func (o *BucketDefinition) SetInvalidatedBySystem(v bool)`

SetInvalidatedBySystem sets InvalidatedBySystem field to given value.

### HasInvalidatedBySystem

`func (o *BucketDefinition) HasInvalidatedBySystem() bool`

HasInvalidatedBySystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


