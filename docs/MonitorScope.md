# MonitorScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | List of monitor Ids in hex. Must be empty if &#x60;all&#x60; is true. | [optional] 
**All** | Pointer to **bool** | true if the schedule applies to all monitors | [optional] [default to false]

## Methods

### NewMonitorScope

`func NewMonitorScope() *MonitorScope`

NewMonitorScope instantiates a new MonitorScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorScopeWithDefaults

`func NewMonitorScopeWithDefaults() *MonitorScope`

NewMonitorScopeWithDefaults instantiates a new MonitorScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *MonitorScope) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *MonitorScope) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *MonitorScope) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *MonitorScope) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetAll

`func (o *MonitorScope) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *MonitorScope) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *MonitorScope) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *MonitorScope) HasAll() bool`

HasAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


