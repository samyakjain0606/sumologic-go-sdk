# AlertsListPageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AlertsListPageObject**](AlertsListPageObject.md) | List of alerts summaries. | [optional] 

## Methods

### NewAlertsListPageResponse

`func NewAlertsListPageResponse() *AlertsListPageResponse`

NewAlertsListPageResponse instantiates a new AlertsListPageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsListPageResponseWithDefaults

`func NewAlertsListPageResponseWithDefaults() *AlertsListPageResponse`

NewAlertsListPageResponseWithDefaults instantiates a new AlertsListPageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AlertsListPageResponse) GetData() []AlertsListPageObject`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AlertsListPageResponse) GetDataOk() (*[]AlertsListPageObject, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AlertsListPageResponse) SetData(v []AlertsListPageObject)`

SetData sets Data field to given value.

### HasData

`func (o *AlertsListPageResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


