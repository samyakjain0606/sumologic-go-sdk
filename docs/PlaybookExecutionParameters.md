# PlaybookExecutionParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlaybookId** | **string** | The id of the playbook which needs to run. | 
**AlertId** | **string** | The alert id which needs to run the playbook. | 

## Methods

### NewPlaybookExecutionParameters

`func NewPlaybookExecutionParameters(playbookId string, alertId string, ) *PlaybookExecutionParameters`

NewPlaybookExecutionParameters instantiates a new PlaybookExecutionParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaybookExecutionParametersWithDefaults

`func NewPlaybookExecutionParametersWithDefaults() *PlaybookExecutionParameters`

NewPlaybookExecutionParametersWithDefaults instantiates a new PlaybookExecutionParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaybookId

`func (o *PlaybookExecutionParameters) GetPlaybookId() string`

GetPlaybookId returns the PlaybookId field if non-nil, zero value otherwise.

### GetPlaybookIdOk

`func (o *PlaybookExecutionParameters) GetPlaybookIdOk() (*string, bool)`

GetPlaybookIdOk returns a tuple with the PlaybookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookId

`func (o *PlaybookExecutionParameters) SetPlaybookId(v string)`

SetPlaybookId sets PlaybookId field to given value.


### GetAlertId

`func (o *PlaybookExecutionParameters) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *PlaybookExecutionParameters) GetAlertIdOk() (*string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *PlaybookExecutionParameters) SetAlertId(v string)`

SetAlertId sets AlertId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


