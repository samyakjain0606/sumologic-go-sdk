# EntitlementUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitlementType** | **string** | String value denoting the type of entitlement. - &#x60;continuous&#x60; for Continuous Analytics, - &#x60;frequent&#x60; for Frequent Analytics, - &#x60;storage&#x60; for Total Storage, - &#x60;metrics&#x60; for Metrics, - &#x60;inFrequentIngest&#x60; for  Infrequent Ingest, - &#x60;inFrequentsStorage&#x60; for Infrequent Storage, - &#x60;inFrequentScannedBytes&#x60; for Infrequent Scan, - &#x60;cloudSIEMContinuous&#x60; for CSE Ingest, - &#x60;tracing&#x60; for Tracing Ingest, - &#x60;soarCount&#x60; for Soar Count, - &#x60;dataForwarding&#x60; for Data Forwarding | 
**Datapoints** | Pointer to [**[]DataPoints**](DataPoints.md) | Array of data points of the entitlement with their respective date range. | [optional] 
**Operators** | [**[]Operator**](Operator.md) | Operators used on the data. Available operators are &#x60;sum&#x60;, &#x60;average&#x60;, &#x60;usagePercentage&#x60;, &#x60;forecastValue&#x60;, &#x60;forecastPercentage&#x60;, and &#x60;forecastRemainingDays&#x60;. sum - Returns the sum of the usages. average - Returns the average of the usages. usagePercentage - Returns percentage of total capacity used for the startDate and endDate. forecastValue - Returns expected usage value assuming current usage behavior continues. forecastPercentage - Returns expected percentage usage by the endDate assuming current usage behavior continues. forecastRemainingDays- Returns the number of expected days, from today, that consumption will last assuming current usage behavior continues. | 
**Tier** | **string** | Tier defines the priority in which the usage for an entitlement is calculated. For example &#x60;promotional&#x60;  for promotional tier. | 
**Label** | **string** | The label for the entitlement. | 

## Methods

### NewEntitlementUsage

`func NewEntitlementUsage(entitlementType string, operators []Operator, tier string, label string, ) *EntitlementUsage`

NewEntitlementUsage instantiates a new EntitlementUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementUsageWithDefaults

`func NewEntitlementUsageWithDefaults() *EntitlementUsage`

NewEntitlementUsageWithDefaults instantiates a new EntitlementUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlementType

`func (o *EntitlementUsage) GetEntitlementType() string`

GetEntitlementType returns the EntitlementType field if non-nil, zero value otherwise.

### GetEntitlementTypeOk

`func (o *EntitlementUsage) GetEntitlementTypeOk() (*string, bool)`

GetEntitlementTypeOk returns a tuple with the EntitlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementType

`func (o *EntitlementUsage) SetEntitlementType(v string)`

SetEntitlementType sets EntitlementType field to given value.


### GetDatapoints

`func (o *EntitlementUsage) GetDatapoints() []DataPoints`

GetDatapoints returns the Datapoints field if non-nil, zero value otherwise.

### GetDatapointsOk

`func (o *EntitlementUsage) GetDatapointsOk() (*[]DataPoints, bool)`

GetDatapointsOk returns a tuple with the Datapoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatapoints

`func (o *EntitlementUsage) SetDatapoints(v []DataPoints)`

SetDatapoints sets Datapoints field to given value.

### HasDatapoints

`func (o *EntitlementUsage) HasDatapoints() bool`

HasDatapoints returns a boolean if a field has been set.

### GetOperators

`func (o *EntitlementUsage) GetOperators() []Operator`

GetOperators returns the Operators field if non-nil, zero value otherwise.

### GetOperatorsOk

`func (o *EntitlementUsage) GetOperatorsOk() (*[]Operator, bool)`

GetOperatorsOk returns a tuple with the Operators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperators

`func (o *EntitlementUsage) SetOperators(v []Operator)`

SetOperators sets Operators field to given value.


### GetTier

`func (o *EntitlementUsage) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *EntitlementUsage) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *EntitlementUsage) SetTier(v string)`

SetTier sets Tier field to given value.


### GetLabel

`func (o *EntitlementUsage) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EntitlementUsage) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EntitlementUsage) SetLabel(v string)`

SetLabel sets Label field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


