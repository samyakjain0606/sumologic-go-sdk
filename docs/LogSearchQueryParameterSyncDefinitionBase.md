# LogSearchQueryParameterSyncDefinitionBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the parameter. | 
**Description** | Pointer to **string** | A description of the parameter. | [optional] 
**DataType** | **string** | The data type of the parameter. Supported values are:   1. &#x60;NUMBER&#x60;   2. &#x60;STRING&#x60;   3. &#x60;ANY&#x60;   4. &#x60;KEYWORD&#x60; | 
**Value** | **string** | A value for the parameter. Should be compatible with the type set in dataType field. | 

## Methods

### NewLogSearchQueryParameterSyncDefinitionBase

`func NewLogSearchQueryParameterSyncDefinitionBase(name string, dataType string, value string, ) *LogSearchQueryParameterSyncDefinitionBase`

NewLogSearchQueryParameterSyncDefinitionBase instantiates a new LogSearchQueryParameterSyncDefinitionBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogSearchQueryParameterSyncDefinitionBaseWithDefaults

`func NewLogSearchQueryParameterSyncDefinitionBaseWithDefaults() *LogSearchQueryParameterSyncDefinitionBase`

NewLogSearchQueryParameterSyncDefinitionBaseWithDefaults instantiates a new LogSearchQueryParameterSyncDefinitionBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LogSearchQueryParameterSyncDefinitionBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogSearchQueryParameterSyncDefinitionBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogSearchQueryParameterSyncDefinitionBase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LogSearchQueryParameterSyncDefinitionBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LogSearchQueryParameterSyncDefinitionBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LogSearchQueryParameterSyncDefinitionBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LogSearchQueryParameterSyncDefinitionBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDataType

`func (o *LogSearchQueryParameterSyncDefinitionBase) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *LogSearchQueryParameterSyncDefinitionBase) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *LogSearchQueryParameterSyncDefinitionBase) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetValue

`func (o *LogSearchQueryParameterSyncDefinitionBase) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LogSearchQueryParameterSyncDefinitionBase) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LogSearchQueryParameterSyncDefinitionBase) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


