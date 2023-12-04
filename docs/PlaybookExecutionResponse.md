# PlaybookExecutionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunningId** | Pointer to **string** | The id of the playbook which is running. | [optional] 
**PlaybookExecutedId** | Pointer to **string** | The id of the playbook when it is executed. | [optional] 

## Methods

### NewPlaybookExecutionResponse

`func NewPlaybookExecutionResponse() *PlaybookExecutionResponse`

NewPlaybookExecutionResponse instantiates a new PlaybookExecutionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaybookExecutionResponseWithDefaults

`func NewPlaybookExecutionResponseWithDefaults() *PlaybookExecutionResponse`

NewPlaybookExecutionResponseWithDefaults instantiates a new PlaybookExecutionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunningId

`func (o *PlaybookExecutionResponse) GetRunningId() string`

GetRunningId returns the RunningId field if non-nil, zero value otherwise.

### GetRunningIdOk

`func (o *PlaybookExecutionResponse) GetRunningIdOk() (*string, bool)`

GetRunningIdOk returns a tuple with the RunningId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningId

`func (o *PlaybookExecutionResponse) SetRunningId(v string)`

SetRunningId sets RunningId field to given value.

### HasRunningId

`func (o *PlaybookExecutionResponse) HasRunningId() bool`

HasRunningId returns a boolean if a field has been set.

### GetPlaybookExecutedId

`func (o *PlaybookExecutionResponse) GetPlaybookExecutedId() string`

GetPlaybookExecutedId returns the PlaybookExecutedId field if non-nil, zero value otherwise.

### GetPlaybookExecutedIdOk

`func (o *PlaybookExecutionResponse) GetPlaybookExecutedIdOk() (*string, bool)`

GetPlaybookExecutedIdOk returns a tuple with the PlaybookExecutedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookExecutedId

`func (o *PlaybookExecutionResponse) SetPlaybookExecutedId(v string)`

SetPlaybookExecutedId sets PlaybookExecutedId field to given value.

### HasPlaybookExecutedId

`func (o *PlaybookExecutionResponse) HasPlaybookExecutedId() bool`

HasPlaybookExecutedId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


