# DataForwardingRule

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

## Methods

### NewDataForwardingRule

`func NewDataForwardingRule(indexId string, destinationId string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, ) *DataForwardingRule`

NewDataForwardingRule instantiates a new DataForwardingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataForwardingRuleWithDefaults

`func NewDataForwardingRuleWithDefaults() *DataForwardingRule`

NewDataForwardingRuleWithDefaults instantiates a new DataForwardingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexId

`func (o *DataForwardingRule) GetIndexId() string`

GetIndexId returns the IndexId field if non-nil, zero value otherwise.

### GetIndexIdOk

`func (o *DataForwardingRule) GetIndexIdOk() (*string, bool)`

GetIndexIdOk returns a tuple with the IndexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexId

`func (o *DataForwardingRule) SetIndexId(v string)`

SetIndexId sets IndexId field to given value.


### GetDestinationId

`func (o *DataForwardingRule) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *DataForwardingRule) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *DataForwardingRule) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetEnabled

`func (o *DataForwardingRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DataForwardingRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DataForwardingRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DataForwardingRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFileFormat

`func (o *DataForwardingRule) GetFileFormat() string`

GetFileFormat returns the FileFormat field if non-nil, zero value otherwise.

### GetFileFormatOk

`func (o *DataForwardingRule) GetFileFormatOk() (*string, bool)`

GetFileFormatOk returns a tuple with the FileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormat

`func (o *DataForwardingRule) SetFileFormat(v string)`

SetFileFormat sets FileFormat field to given value.

### HasFileFormat

`func (o *DataForwardingRule) HasFileFormat() bool`

HasFileFormat returns a boolean if a field has been set.

### GetFormat

`func (o *DataForwardingRule) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *DataForwardingRule) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *DataForwardingRule) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *DataForwardingRule) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DataForwardingRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataForwardingRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataForwardingRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *DataForwardingRule) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DataForwardingRule) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DataForwardingRule) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *DataForwardingRule) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *DataForwardingRule) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *DataForwardingRule) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *DataForwardingRule) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *DataForwardingRule) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *DataForwardingRule) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetId

`func (o *DataForwardingRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataForwardingRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataForwardingRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataForwardingRule) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


