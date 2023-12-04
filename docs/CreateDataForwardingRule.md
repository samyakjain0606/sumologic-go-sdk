# CreateDataForwardingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexId** | **string** | The &#x60;id&#x60; of the Partition or Scheduled View the rule applies to. | 
**DestinationId** | **string** | The data forwarding destination id. | 
**Enabled** | Pointer to **bool** | True when the data forwarding rule is enabled. | [optional] 
**FileFormat** | Pointer to **string** | Specify the path prefix to a directory in the S3 bucket and how to format the file name. | [optional] 
**Format** | Pointer to **string** | Format of the payload. | [optional] [default to "csv"]

## Methods

### NewCreateDataForwardingRule

`func NewCreateDataForwardingRule(indexId string, destinationId string, ) *CreateDataForwardingRule`

NewCreateDataForwardingRule instantiates a new CreateDataForwardingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDataForwardingRuleWithDefaults

`func NewCreateDataForwardingRuleWithDefaults() *CreateDataForwardingRule`

NewCreateDataForwardingRuleWithDefaults instantiates a new CreateDataForwardingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexId

`func (o *CreateDataForwardingRule) GetIndexId() string`

GetIndexId returns the IndexId field if non-nil, zero value otherwise.

### GetIndexIdOk

`func (o *CreateDataForwardingRule) GetIndexIdOk() (*string, bool)`

GetIndexIdOk returns a tuple with the IndexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexId

`func (o *CreateDataForwardingRule) SetIndexId(v string)`

SetIndexId sets IndexId field to given value.


### GetDestinationId

`func (o *CreateDataForwardingRule) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *CreateDataForwardingRule) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *CreateDataForwardingRule) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetEnabled

`func (o *CreateDataForwardingRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateDataForwardingRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateDataForwardingRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateDataForwardingRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFileFormat

`func (o *CreateDataForwardingRule) GetFileFormat() string`

GetFileFormat returns the FileFormat field if non-nil, zero value otherwise.

### GetFileFormatOk

`func (o *CreateDataForwardingRule) GetFileFormatOk() (*string, bool)`

GetFileFormatOk returns a tuple with the FileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormat

`func (o *CreateDataForwardingRule) SetFileFormat(v string)`

SetFileFormat sets FileFormat field to given value.

### HasFileFormat

`func (o *CreateDataForwardingRule) HasFileFormat() bool`

HasFileFormat returns a boolean if a field has been set.

### GetFormat

`func (o *CreateDataForwardingRule) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateDataForwardingRule) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateDataForwardingRule) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateDataForwardingRule) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


