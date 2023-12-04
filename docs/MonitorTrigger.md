# MonitorTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorId** | **string** | Hex-id of the monitor on which the SLI is based. | 
**TriggerTypes** | **[]string** | The types of trigger conditions (such as Critical, Warning, MissingData etc). | 

## Methods

### NewMonitorTrigger

`func NewMonitorTrigger(monitorId string, triggerTypes []string, ) *MonitorTrigger`

NewMonitorTrigger instantiates a new MonitorTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorTriggerWithDefaults

`func NewMonitorTriggerWithDefaults() *MonitorTrigger`

NewMonitorTriggerWithDefaults instantiates a new MonitorTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorId

`func (o *MonitorTrigger) GetMonitorId() string`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *MonitorTrigger) GetMonitorIdOk() (*string, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *MonitorTrigger) SetMonitorId(v string)`

SetMonitorId sets MonitorId field to given value.


### GetTriggerTypes

`func (o *MonitorTrigger) GetTriggerTypes() []string`

GetTriggerTypes returns the TriggerTypes field if non-nil, zero value otherwise.

### GetTriggerTypesOk

`func (o *MonitorTrigger) GetTriggerTypesOk() (*[]string, bool)`

GetTriggerTypesOk returns a tuple with the TriggerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTypes

`func (o *MonitorTrigger) SetTriggerTypes(v []string)`

SetTriggerTypes sets TriggerTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


