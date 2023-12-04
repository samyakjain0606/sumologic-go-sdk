# MutingSchedulesLibraryMutingScheduleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | [**ScheduleDefinition**](ScheduleDefinition.md) |  | 
**Monitor** | Pointer to [**MonitorScope**](MonitorScope.md) |  | [optional] 

## Methods

### NewMutingSchedulesLibraryMutingScheduleResponse

`func NewMutingSchedulesLibraryMutingScheduleResponse(schedule ScheduleDefinition, ) *MutingSchedulesLibraryMutingScheduleResponse`

NewMutingSchedulesLibraryMutingScheduleResponse instantiates a new MutingSchedulesLibraryMutingScheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutingSchedulesLibraryMutingScheduleResponseWithDefaults

`func NewMutingSchedulesLibraryMutingScheduleResponseWithDefaults() *MutingSchedulesLibraryMutingScheduleResponse`

NewMutingSchedulesLibraryMutingScheduleResponseWithDefaults instantiates a new MutingSchedulesLibraryMutingScheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *MutingSchedulesLibraryMutingScheduleResponse) GetSchedule() ScheduleDefinition`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *MutingSchedulesLibraryMutingScheduleResponse) GetScheduleOk() (*ScheduleDefinition, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *MutingSchedulesLibraryMutingScheduleResponse) SetSchedule(v ScheduleDefinition)`

SetSchedule sets Schedule field to given value.


### GetMonitor

`func (o *MutingSchedulesLibraryMutingScheduleResponse) GetMonitor() MonitorScope`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *MutingSchedulesLibraryMutingScheduleResponse) GetMonitorOk() (*MonitorScope, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *MutingSchedulesLibraryMutingScheduleResponse) SetMonitor(v MonitorScope)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *MutingSchedulesLibraryMutingScheduleResponse) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


