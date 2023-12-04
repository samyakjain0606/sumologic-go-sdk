# DashboardMigrationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuccessCount** | **int32** | A successful migration to Dashboard(New). | 
**FailedCount** | **int32** | A failed migration to Dashboard(New). | 
**TotalCount** | **int32** | The total number of Legacy Dashboards to migrate. | 

## Methods

### NewDashboardMigrationStatus

`func NewDashboardMigrationStatus(successCount int32, failedCount int32, totalCount int32, ) *DashboardMigrationStatus`

NewDashboardMigrationStatus instantiates a new DashboardMigrationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardMigrationStatusWithDefaults

`func NewDashboardMigrationStatusWithDefaults() *DashboardMigrationStatus`

NewDashboardMigrationStatusWithDefaults instantiates a new DashboardMigrationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessCount

`func (o *DashboardMigrationStatus) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *DashboardMigrationStatus) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *DashboardMigrationStatus) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.


### GetFailedCount

`func (o *DashboardMigrationStatus) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *DashboardMigrationStatus) GetFailedCountOk() (*int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *DashboardMigrationStatus) SetFailedCount(v int32)`

SetFailedCount sets FailedCount field to given value.


### GetTotalCount

`func (o *DashboardMigrationStatus) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DashboardMigrationStatus) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DashboardMigrationStatus) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


