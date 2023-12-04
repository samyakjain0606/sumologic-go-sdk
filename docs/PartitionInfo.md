# PartitionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the partition. | 
**AnalyticsTier** | Pointer to **string** | The Data Tier where the data in the partition will reside. Possible values are:               1. &#x60;continuous&#x60;               2. &#x60;frequent&#x60;               3. &#x60;infrequent&#x60; Note: The \&quot;infrequent\&quot; and \&quot;frequent\&quot; tiers are only available to Cloud Flex Credits Enterprise Suite accounts. | [optional] 
**DataFilterGroup** | Pointer to **string** | The Data Filter Group to which this parition belongs to. Possible values are :               1. &#x60;LOGS&#x60;               2. &#x60;SECURITY&#x60;               3. &#x60;AUDIT&#x60; | [optional] 

## Methods

### NewPartitionInfo

`func NewPartitionInfo(name string, ) *PartitionInfo`

NewPartitionInfo instantiates a new PartitionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionInfoWithDefaults

`func NewPartitionInfoWithDefaults() *PartitionInfo`

NewPartitionInfoWithDefaults instantiates a new PartitionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PartitionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartitionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartitionInfo) SetName(v string)`

SetName sets Name field to given value.


### GetAnalyticsTier

`func (o *PartitionInfo) GetAnalyticsTier() string`

GetAnalyticsTier returns the AnalyticsTier field if non-nil, zero value otherwise.

### GetAnalyticsTierOk

`func (o *PartitionInfo) GetAnalyticsTierOk() (*string, bool)`

GetAnalyticsTierOk returns a tuple with the AnalyticsTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsTier

`func (o *PartitionInfo) SetAnalyticsTier(v string)`

SetAnalyticsTier sets AnalyticsTier field to given value.

### HasAnalyticsTier

`func (o *PartitionInfo) HasAnalyticsTier() bool`

HasAnalyticsTier returns a boolean if a field has been set.

### GetDataFilterGroup

`func (o *PartitionInfo) GetDataFilterGroup() string`

GetDataFilterGroup returns the DataFilterGroup field if non-nil, zero value otherwise.

### GetDataFilterGroupOk

`func (o *PartitionInfo) GetDataFilterGroupOk() (*string, bool)`

GetDataFilterGroupOk returns a tuple with the DataFilterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFilterGroup

`func (o *PartitionInfo) SetDataFilterGroup(v string)`

SetDataFilterGroup sets DataFilterGroup field to given value.

### HasDataFilterGroup

`func (o *PartitionInfo) HasDataFilterGroup() bool`

HasDataFilterGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


