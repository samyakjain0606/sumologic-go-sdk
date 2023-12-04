# AdhocMutingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | **int64** | Start time of adhoc muting period in Epoch. | 
**EndTime** | Pointer to **int64** | End time of the adhoc muting period in Epoch.If muting indefinitely, this will be empty. | [optional] 

## Methods

### NewAdhocMutingResponse

`func NewAdhocMutingResponse(startTime int64, ) *AdhocMutingResponse`

NewAdhocMutingResponse instantiates a new AdhocMutingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdhocMutingResponseWithDefaults

`func NewAdhocMutingResponseWithDefaults() *AdhocMutingResponse`

NewAdhocMutingResponseWithDefaults instantiates a new AdhocMutingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *AdhocMutingResponse) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AdhocMutingResponse) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AdhocMutingResponse) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *AdhocMutingResponse) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *AdhocMutingResponse) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *AdhocMutingResponse) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *AdhocMutingResponse) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


