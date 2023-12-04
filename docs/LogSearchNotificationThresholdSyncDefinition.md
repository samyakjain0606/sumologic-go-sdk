# LogSearchNotificationThresholdSyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThresholdType** | Pointer to **string** | This property is deprecated. The system will automatically infer the value of this field from the query going forward, so the user-specified value will no longer be honored. Threshold type. Possible values are:  1. &#x60;message&#x60;  2. &#x60;group&#x60;  Use &#x60;group&#x60; as threshold type if the search query is of aggregate type. For non-aggregate queries, set it to &#x60;message&#x60;. | [optional] 
**Operator** | **string** | Criterion to be applied when comparing actual result count with expected count. Possible values are:  1. &#x60;eq&#x60;  2. &#x60;gt&#x60;  3. &#x60;ge&#x60;  4. &#x60;lt&#x60;  5. &#x60;le&#x60; | 
**Count** | **int32** | Expected result count. | 

## Methods

### NewLogSearchNotificationThresholdSyncDefinition

`func NewLogSearchNotificationThresholdSyncDefinition(operator string, count int32, ) *LogSearchNotificationThresholdSyncDefinition`

NewLogSearchNotificationThresholdSyncDefinition instantiates a new LogSearchNotificationThresholdSyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogSearchNotificationThresholdSyncDefinitionWithDefaults

`func NewLogSearchNotificationThresholdSyncDefinitionWithDefaults() *LogSearchNotificationThresholdSyncDefinition`

NewLogSearchNotificationThresholdSyncDefinitionWithDefaults instantiates a new LogSearchNotificationThresholdSyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThresholdType

`func (o *LogSearchNotificationThresholdSyncDefinition) GetThresholdType() string`

GetThresholdType returns the ThresholdType field if non-nil, zero value otherwise.

### GetThresholdTypeOk

`func (o *LogSearchNotificationThresholdSyncDefinition) GetThresholdTypeOk() (*string, bool)`

GetThresholdTypeOk returns a tuple with the ThresholdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdType

`func (o *LogSearchNotificationThresholdSyncDefinition) SetThresholdType(v string)`

SetThresholdType sets ThresholdType field to given value.

### HasThresholdType

`func (o *LogSearchNotificationThresholdSyncDefinition) HasThresholdType() bool`

HasThresholdType returns a boolean if a field has been set.

### GetOperator

`func (o *LogSearchNotificationThresholdSyncDefinition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *LogSearchNotificationThresholdSyncDefinition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *LogSearchNotificationThresholdSyncDefinition) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetCount

`func (o *LogSearchNotificationThresholdSyncDefinition) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *LogSearchNotificationThresholdSyncDefinition) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *LogSearchNotificationThresholdSyncDefinition) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


