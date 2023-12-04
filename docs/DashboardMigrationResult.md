# DashboardMigrationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **map[string]string** | A mapping of Legacy Dashboard Content Ids to migrated Dashboard(New) Content Ids. Only successful migration are shown here, see errors field for failed migrations and the failure reason.  | 
**RichData** | Pointer to [**map[string]MigratedDashboardInfo**](MigratedDashboardInfo.md) | A mapping of Legacy Dashboard Content Ids to migrated Dashboard(New) info. Only successful migration are shown here, see errors field for failed migrations and the failure reason.  | [optional] 
**Status** | [**DashboardMigrationStatus**](DashboardMigrationStatus.md) |  | 
**Errors** | Pointer to [**map[string][]ErrorDescription**](array.md) | A mapping of Legacy Dashboards Content Identifiers that failed validation to the failure reason(s). | [optional] 
**Warnings** | Pointer to [**map[string][]ErrorDescription**](array.md) | A mapping of Legacy Dashboards Content Identifiers to warnings. | [optional] 

## Methods

### NewDashboardMigrationResult

`func NewDashboardMigrationResult(data map[string]string, status DashboardMigrationStatus, ) *DashboardMigrationResult`

NewDashboardMigrationResult instantiates a new DashboardMigrationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardMigrationResultWithDefaults

`func NewDashboardMigrationResultWithDefaults() *DashboardMigrationResult`

NewDashboardMigrationResultWithDefaults instantiates a new DashboardMigrationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DashboardMigrationResult) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DashboardMigrationResult) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DashboardMigrationResult) SetData(v map[string]string)`

SetData sets Data field to given value.


### GetRichData

`func (o *DashboardMigrationResult) GetRichData() map[string]MigratedDashboardInfo`

GetRichData returns the RichData field if non-nil, zero value otherwise.

### GetRichDataOk

`func (o *DashboardMigrationResult) GetRichDataOk() (*map[string]MigratedDashboardInfo, bool)`

GetRichDataOk returns a tuple with the RichData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichData

`func (o *DashboardMigrationResult) SetRichData(v map[string]MigratedDashboardInfo)`

SetRichData sets RichData field to given value.

### HasRichData

`func (o *DashboardMigrationResult) HasRichData() bool`

HasRichData returns a boolean if a field has been set.

### GetStatus

`func (o *DashboardMigrationResult) GetStatus() DashboardMigrationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DashboardMigrationResult) GetStatusOk() (*DashboardMigrationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DashboardMigrationResult) SetStatus(v DashboardMigrationStatus)`

SetStatus sets Status field to given value.


### GetErrors

`func (o *DashboardMigrationResult) GetErrors() map[string][]ErrorDescription`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *DashboardMigrationResult) GetErrorsOk() (*map[string][]ErrorDescription, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *DashboardMigrationResult) SetErrors(v map[string][]ErrorDescription)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *DashboardMigrationResult) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *DashboardMigrationResult) GetWarnings() map[string][]ErrorDescription`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DashboardMigrationResult) GetWarningsOk() (*map[string][]ErrorDescription, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DashboardMigrationResult) SetWarnings(v map[string][]ErrorDescription)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DashboardMigrationResult) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


