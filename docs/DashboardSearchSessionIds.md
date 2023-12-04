# DashboardSearchSessionIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **map[string]string** | Map of search keys to session ids. | [default to {}]
**Errors** | Pointer to [**map[string]ErrorResponse**](ErrorResponse.md) | Error description for the session keys that failed validation. | [optional] 

## Methods

### NewDashboardSearchSessionIds

`func NewDashboardSearchSessionIds(data map[string]string, ) *DashboardSearchSessionIds`

NewDashboardSearchSessionIds instantiates a new DashboardSearchSessionIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardSearchSessionIdsWithDefaults

`func NewDashboardSearchSessionIdsWithDefaults() *DashboardSearchSessionIds`

NewDashboardSearchSessionIdsWithDefaults instantiates a new DashboardSearchSessionIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DashboardSearchSessionIds) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DashboardSearchSessionIds) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DashboardSearchSessionIds) SetData(v map[string]string)`

SetData sets Data field to given value.


### GetErrors

`func (o *DashboardSearchSessionIds) GetErrors() map[string]ErrorResponse`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *DashboardSearchSessionIds) GetErrorsOk() (*map[string]ErrorResponse, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *DashboardSearchSessionIds) SetErrors(v map[string]ErrorResponse)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *DashboardSearchSessionIds) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


