# PaginatedDashboards

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dashboards** | [**[]Dashboard**](Dashboard.md) | List of dashboards. | 
**Next** | Pointer to **string** | Next continuation token. &#x60;token&#x60; is set to null when no more pages are left. | [optional] 

## Methods

### NewPaginatedDashboards

`func NewPaginatedDashboards(dashboards []Dashboard, ) *PaginatedDashboards`

NewPaginatedDashboards instantiates a new PaginatedDashboards object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedDashboardsWithDefaults

`func NewPaginatedDashboardsWithDefaults() *PaginatedDashboards`

NewPaginatedDashboardsWithDefaults instantiates a new PaginatedDashboards object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboards

`func (o *PaginatedDashboards) GetDashboards() []Dashboard`

GetDashboards returns the Dashboards field if non-nil, zero value otherwise.

### GetDashboardsOk

`func (o *PaginatedDashboards) GetDashboardsOk() (*[]Dashboard, bool)`

GetDashboardsOk returns a tuple with the Dashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboards

`func (o *PaginatedDashboards) SetDashboards(v []Dashboard)`

SetDashboards sets Dashboards field to given value.


### GetNext

`func (o *PaginatedDashboards) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedDashboards) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedDashboards) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedDashboards) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


