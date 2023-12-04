# MutingInformationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifier of the monitor. | [optional] 
**IsMuted** | **bool** | Flag to indicate the monitor muted or not. | [default to false]
**MutingEndTime** | Pointer to **int64** | Timestamp in Epoch that this monitor is currently muted until. | [optional] 
**MutingSchedules** | Pointer to [**[]MutingScheduleResponse**](MutingScheduleResponse.md) | Array of muting schedules that this monitor is associated with. | [optional] 
**AdhocMuting** | Pointer to [**AdhocMutingResponse**](AdhocMutingResponse.md) |  | [optional] 

## Methods

### NewMutingInformationResponse

`func NewMutingInformationResponse(isMuted bool, ) *MutingInformationResponse`

NewMutingInformationResponse instantiates a new MutingInformationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutingInformationResponseWithDefaults

`func NewMutingInformationResponseWithDefaults() *MutingInformationResponse`

NewMutingInformationResponseWithDefaults instantiates a new MutingInformationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MutingInformationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MutingInformationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MutingInformationResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MutingInformationResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsMuted

`func (o *MutingInformationResponse) GetIsMuted() bool`

GetIsMuted returns the IsMuted field if non-nil, zero value otherwise.

### GetIsMutedOk

`func (o *MutingInformationResponse) GetIsMutedOk() (*bool, bool)`

GetIsMutedOk returns a tuple with the IsMuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMuted

`func (o *MutingInformationResponse) SetIsMuted(v bool)`

SetIsMuted sets IsMuted field to given value.


### GetMutingEndTime

`func (o *MutingInformationResponse) GetMutingEndTime() int64`

GetMutingEndTime returns the MutingEndTime field if non-nil, zero value otherwise.

### GetMutingEndTimeOk

`func (o *MutingInformationResponse) GetMutingEndTimeOk() (*int64, bool)`

GetMutingEndTimeOk returns a tuple with the MutingEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutingEndTime

`func (o *MutingInformationResponse) SetMutingEndTime(v int64)`

SetMutingEndTime sets MutingEndTime field to given value.

### HasMutingEndTime

`func (o *MutingInformationResponse) HasMutingEndTime() bool`

HasMutingEndTime returns a boolean if a field has been set.

### GetMutingSchedules

`func (o *MutingInformationResponse) GetMutingSchedules() []MutingScheduleResponse`

GetMutingSchedules returns the MutingSchedules field if non-nil, zero value otherwise.

### GetMutingSchedulesOk

`func (o *MutingInformationResponse) GetMutingSchedulesOk() (*[]MutingScheduleResponse, bool)`

GetMutingSchedulesOk returns a tuple with the MutingSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutingSchedules

`func (o *MutingInformationResponse) SetMutingSchedules(v []MutingScheduleResponse)`

SetMutingSchedules sets MutingSchedules field to given value.

### HasMutingSchedules

`func (o *MutingInformationResponse) HasMutingSchedules() bool`

HasMutingSchedules returns a boolean if a field has been set.

### GetAdhocMuting

`func (o *MutingInformationResponse) GetAdhocMuting() AdhocMutingResponse`

GetAdhocMuting returns the AdhocMuting field if non-nil, zero value otherwise.

### GetAdhocMutingOk

`func (o *MutingInformationResponse) GetAdhocMutingOk() (*AdhocMutingResponse, bool)`

GetAdhocMutingOk returns a tuple with the AdhocMuting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdhocMuting

`func (o *MutingInformationResponse) SetAdhocMuting(v AdhocMutingResponse)`

SetAdhocMuting sets AdhocMuting field to given value.

### HasAdhocMuting

`func (o *MutingInformationResponse) HasAdhocMuting() bool`

HasAdhocMuting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


