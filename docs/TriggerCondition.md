# TriggerCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectionMethod** | Pointer to **string** | Detection method of the trigger condition. Valid values:   1. &#x60;StaticCondition&#x60;: A condition that triggers based off of a static threshold. This &#x60;detectionMethod&#x60; is deprecated, it is recommended to use other ones instead.   2. &#x60;LogsStaticCondition&#x60;: A logs condition that triggers based off of a static threshold.   3. &#x60;MetricsStaticCondition&#x60;: A metrics condition that triggers based off of a static threshold.   4. &#x60;LogsOutlierCondition&#x60;: A logs condition that triggers based off of a dynamic outlier threshold.   5. &#x60;MetricsOutlierCondition&#x60;: A metrics condition that triggers based off of a dynamic outlier threshold.   6. &#x60;LogsMissingDataCondition&#x60;: A logs missing data condition that triggers based off of no data available.   7. &#x60;MetricsMissingDataCondition&#x60;: A metrics missing data condition that triggers based off of no data available.   8. &#x60;SloSliCondition&#x60;: An SLO condition that triggers based off of current SLI value.   9. &#x60;SloBurnRateCondition&#x60;: An SLO condition that triggers based off of error budget burn rate. | [optional] [default to "StaticCondition"]
**TriggerType** | **string** | The type of trigger condition. Valid values:   1. &#x60;Critical&#x60;: A critical condition to trigger on.   2. &#x60;Warning&#x60;: A warning condition to trigger on.   3. &#x60;MissingData&#x60;: A condition that indicates data is missing.   4. &#x60;ResolvedCritical&#x60;: A condition to resolve a Critical trigger on.   5. &#x60;ResolvedWarning&#x60;: A condition to resolve a Warning trigger on.   6. &#x60;ResolvedMissingData&#x60;: A condition to resolve a MissingData trigger. | 
**ResolutionWindow** | Pointer to **NullableString** | The resolution window that the recovery condition must be met in each evaluation that happens within this entire duration before the alert is recovered (resolved). If not specified, the time range of your trigger will be used. Valid values are: &#x60;0m&#x60;, &#x60;-5m&#x60;, &#x60;-10m&#x60;, &#x60;-15m&#x60;, &#x60;-30m&#x60;, &#x60;-1h&#x60;, &#x60;-3h&#x60;, &#x60;-6h&#x60;, &#x60;-12h&#x60;, or &#x60;-24h&#x60; | [optional] 

## Methods

### NewTriggerCondition

`func NewTriggerCondition(triggerType string, ) *TriggerCondition`

NewTriggerCondition instantiates a new TriggerCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerConditionWithDefaults

`func NewTriggerConditionWithDefaults() *TriggerCondition`

NewTriggerConditionWithDefaults instantiates a new TriggerCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetectionMethod

`func (o *TriggerCondition) GetDetectionMethod() string`

GetDetectionMethod returns the DetectionMethod field if non-nil, zero value otherwise.

### GetDetectionMethodOk

`func (o *TriggerCondition) GetDetectionMethodOk() (*string, bool)`

GetDetectionMethodOk returns a tuple with the DetectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionMethod

`func (o *TriggerCondition) SetDetectionMethod(v string)`

SetDetectionMethod sets DetectionMethod field to given value.

### HasDetectionMethod

`func (o *TriggerCondition) HasDetectionMethod() bool`

HasDetectionMethod returns a boolean if a field has been set.

### GetTriggerType

`func (o *TriggerCondition) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *TriggerCondition) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *TriggerCondition) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.


### GetResolutionWindow

`func (o *TriggerCondition) GetResolutionWindow() string`

GetResolutionWindow returns the ResolutionWindow field if non-nil, zero value otherwise.

### GetResolutionWindowOk

`func (o *TriggerCondition) GetResolutionWindowOk() (*string, bool)`

GetResolutionWindowOk returns a tuple with the ResolutionWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionWindow

`func (o *TriggerCondition) SetResolutionWindow(v string)`

SetResolutionWindow sets ResolutionWindow field to given value.

### HasResolutionWindow

`func (o *TriggerCondition) HasResolutionWindow() bool`

HasResolutionWindow returns a boolean if a field has been set.

### SetResolutionWindowNil

`func (o *TriggerCondition) SetResolutionWindowNil(b bool)`

 SetResolutionWindowNil sets the value for ResolutionWindow to be an explicit nil

### UnsetResolutionWindow
`func (o *TriggerCondition) UnsetResolutionWindow()`

UnsetResolutionWindow ensures that no value is present for ResolutionWindow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


