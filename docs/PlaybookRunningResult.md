# PlaybookRunningResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **time.Time** | The running start date time of the playbook. | [optional] 
**EndDate** | Pointer to **time.Time** | The running end date time of the playbook. | [optional] 
**Id** | **string** | The id of the playbook running. | 
**PlaybookId** | **string** | The id of the playbook. | 
**IsChild** | **bool** | The isChild of other playbook. | [default to false]
**Name** | **string** | The name of the playbook running. | 
**Status** | **string** | The status of the playbook running. | 
**StatusCode** | **int32** | The status code of the playbook running. | 

## Methods

### NewPlaybookRunningResult

`func NewPlaybookRunningResult(id string, playbookId string, isChild bool, name string, status string, statusCode int32, ) *PlaybookRunningResult`

NewPlaybookRunningResult instantiates a new PlaybookRunningResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaybookRunningResultWithDefaults

`func NewPlaybookRunningResultWithDefaults() *PlaybookRunningResult`

NewPlaybookRunningResultWithDefaults instantiates a new PlaybookRunningResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *PlaybookRunningResult) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PlaybookRunningResult) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PlaybookRunningResult) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PlaybookRunningResult) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *PlaybookRunningResult) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PlaybookRunningResult) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PlaybookRunningResult) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PlaybookRunningResult) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetId

`func (o *PlaybookRunningResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlaybookRunningResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlaybookRunningResult) SetId(v string)`

SetId sets Id field to given value.


### GetPlaybookId

`func (o *PlaybookRunningResult) GetPlaybookId() string`

GetPlaybookId returns the PlaybookId field if non-nil, zero value otherwise.

### GetPlaybookIdOk

`func (o *PlaybookRunningResult) GetPlaybookIdOk() (*string, bool)`

GetPlaybookIdOk returns a tuple with the PlaybookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookId

`func (o *PlaybookRunningResult) SetPlaybookId(v string)`

SetPlaybookId sets PlaybookId field to given value.


### GetIsChild

`func (o *PlaybookRunningResult) GetIsChild() bool`

GetIsChild returns the IsChild field if non-nil, zero value otherwise.

### GetIsChildOk

`func (o *PlaybookRunningResult) GetIsChildOk() (*bool, bool)`

GetIsChildOk returns a tuple with the IsChild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChild

`func (o *PlaybookRunningResult) SetIsChild(v bool)`

SetIsChild sets IsChild field to given value.


### GetName

`func (o *PlaybookRunningResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlaybookRunningResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlaybookRunningResult) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *PlaybookRunningResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlaybookRunningResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlaybookRunningResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusCode

`func (o *PlaybookRunningResult) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *PlaybookRunningResult) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *PlaybookRunningResult) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


