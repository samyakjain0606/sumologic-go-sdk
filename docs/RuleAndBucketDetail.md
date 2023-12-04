# RuleAndBucketDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexId** | **string** | The &#x60;id&#x60; of the Partition or Scheduled View the rule applies to. | 
**DestinationId** | **string** | The data forwarding destination id. | 
**Enabled** | Pointer to **bool** | True when the data forwarding rule is enabled. | [optional] 
**FileFormat** | Pointer to **string** | Specify the path prefix to a directory in the S3 bucket and how to format the file name. | [optional] 
**Format** | Pointer to **string** | Format of the payload. | [optional] [default to "csv"]
**CreatedAt** | **time.Time** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**CreatedBy** | **string** | Identifier of the user who created the resource. | 
**ModifiedAt** | **time.Time** | Last modification timestamp in UTC. | 
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | 
**Id** | Pointer to **string** | The unique identifier of the data forwarding rule. | [optional] 
**Bucket** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRuleAndBucketDetail

`func NewRuleAndBucketDetail(indexId string, destinationId string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, ) *RuleAndBucketDetail`

NewRuleAndBucketDetail instantiates a new RuleAndBucketDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleAndBucketDetailWithDefaults

`func NewRuleAndBucketDetailWithDefaults() *RuleAndBucketDetail`

NewRuleAndBucketDetailWithDefaults instantiates a new RuleAndBucketDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexId

`func (o *RuleAndBucketDetail) GetIndexId() string`

GetIndexId returns the IndexId field if non-nil, zero value otherwise.

### GetIndexIdOk

`func (o *RuleAndBucketDetail) GetIndexIdOk() (*string, bool)`

GetIndexIdOk returns a tuple with the IndexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexId

`func (o *RuleAndBucketDetail) SetIndexId(v string)`

SetIndexId sets IndexId field to given value.


### GetDestinationId

`func (o *RuleAndBucketDetail) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *RuleAndBucketDetail) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *RuleAndBucketDetail) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetEnabled

`func (o *RuleAndBucketDetail) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RuleAndBucketDetail) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RuleAndBucketDetail) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RuleAndBucketDetail) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFileFormat

`func (o *RuleAndBucketDetail) GetFileFormat() string`

GetFileFormat returns the FileFormat field if non-nil, zero value otherwise.

### GetFileFormatOk

`func (o *RuleAndBucketDetail) GetFileFormatOk() (*string, bool)`

GetFileFormatOk returns a tuple with the FileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormat

`func (o *RuleAndBucketDetail) SetFileFormat(v string)`

SetFileFormat sets FileFormat field to given value.

### HasFileFormat

`func (o *RuleAndBucketDetail) HasFileFormat() bool`

HasFileFormat returns a boolean if a field has been set.

### GetFormat

`func (o *RuleAndBucketDetail) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *RuleAndBucketDetail) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *RuleAndBucketDetail) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *RuleAndBucketDetail) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RuleAndBucketDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RuleAndBucketDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RuleAndBucketDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *RuleAndBucketDetail) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RuleAndBucketDetail) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RuleAndBucketDetail) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *RuleAndBucketDetail) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *RuleAndBucketDetail) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *RuleAndBucketDetail) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *RuleAndBucketDetail) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *RuleAndBucketDetail) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *RuleAndBucketDetail) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetId

`func (o *RuleAndBucketDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleAndBucketDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleAndBucketDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RuleAndBucketDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBucket

`func (o *RuleAndBucketDetail) GetBucket() map[string]interface{}`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *RuleAndBucketDetail) GetBucketOk() (*map[string]interface{}, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *RuleAndBucketDetail) SetBucket(v map[string]interface{})`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *RuleAndBucketDetail) HasBucket() bool`

HasBucket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


