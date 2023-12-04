# TraceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Trace identifier. | 
**RootService** | Pointer to **string** | Root service which started the trace. Examples: &#x60;user-service&#x60;, &#x60;authentication-service&#x60;, &#x60;payment-service&#x60;, &#x60;/shopping-cart&#x60; | [optional] 
**RootResource** | Pointer to **string** | Root resource on which the trace was started. Examples: &#x60;db.query&#x60;, &#x60;http.request&#x60;, &#x60;rpc.call&#x60;, &#x60;container&#x60; | [optional] 
**RootStatus** | Pointer to [**TraceSpanStatus**](TraceSpanStatus.md) |  | [optional] 
**RootOperationName** | Pointer to **string** | The name of the operation given to the root span. | [optional] 
**Metrics** | Pointer to [**map[string]DoubleTracingValue**](DoubleTracingValue.md) | Calculated trace metrics. | [optional] 
**StartedAt** | Pointer to **time.Time** | Date and time the trace was started in [ISO 8601 / RFC3339](https://tools.ietf.org/html/rfc3339) format. | [optional] 
**CriticalPathServiceBreakdownSummary** | Pointer to [**CriticalPathServiceBreakdownSummary**](CriticalPathServiceBreakdownSummary.md) |  | [optional] 

## Methods

### NewTraceDetail

`func NewTraceDetail(id string, ) *TraceDetail`

NewTraceDetail instantiates a new TraceDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceDetailWithDefaults

`func NewTraceDetailWithDefaults() *TraceDetail`

NewTraceDetailWithDefaults instantiates a new TraceDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TraceDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TraceDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TraceDetail) SetId(v string)`

SetId sets Id field to given value.


### GetRootService

`func (o *TraceDetail) GetRootService() string`

GetRootService returns the RootService field if non-nil, zero value otherwise.

### GetRootServiceOk

`func (o *TraceDetail) GetRootServiceOk() (*string, bool)`

GetRootServiceOk returns a tuple with the RootService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootService

`func (o *TraceDetail) SetRootService(v string)`

SetRootService sets RootService field to given value.

### HasRootService

`func (o *TraceDetail) HasRootService() bool`

HasRootService returns a boolean if a field has been set.

### GetRootResource

`func (o *TraceDetail) GetRootResource() string`

GetRootResource returns the RootResource field if non-nil, zero value otherwise.

### GetRootResourceOk

`func (o *TraceDetail) GetRootResourceOk() (*string, bool)`

GetRootResourceOk returns a tuple with the RootResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootResource

`func (o *TraceDetail) SetRootResource(v string)`

SetRootResource sets RootResource field to given value.

### HasRootResource

`func (o *TraceDetail) HasRootResource() bool`

HasRootResource returns a boolean if a field has been set.

### GetRootStatus

`func (o *TraceDetail) GetRootStatus() TraceSpanStatus`

GetRootStatus returns the RootStatus field if non-nil, zero value otherwise.

### GetRootStatusOk

`func (o *TraceDetail) GetRootStatusOk() (*TraceSpanStatus, bool)`

GetRootStatusOk returns a tuple with the RootStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootStatus

`func (o *TraceDetail) SetRootStatus(v TraceSpanStatus)`

SetRootStatus sets RootStatus field to given value.

### HasRootStatus

`func (o *TraceDetail) HasRootStatus() bool`

HasRootStatus returns a boolean if a field has been set.

### GetRootOperationName

`func (o *TraceDetail) GetRootOperationName() string`

GetRootOperationName returns the RootOperationName field if non-nil, zero value otherwise.

### GetRootOperationNameOk

`func (o *TraceDetail) GetRootOperationNameOk() (*string, bool)`

GetRootOperationNameOk returns a tuple with the RootOperationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootOperationName

`func (o *TraceDetail) SetRootOperationName(v string)`

SetRootOperationName sets RootOperationName field to given value.

### HasRootOperationName

`func (o *TraceDetail) HasRootOperationName() bool`

HasRootOperationName returns a boolean if a field has been set.

### GetMetrics

`func (o *TraceDetail) GetMetrics() map[string]DoubleTracingValue`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *TraceDetail) GetMetricsOk() (*map[string]DoubleTracingValue, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *TraceDetail) SetMetrics(v map[string]DoubleTracingValue)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *TraceDetail) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetStartedAt

`func (o *TraceDetail) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *TraceDetail) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *TraceDetail) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *TraceDetail) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCriticalPathServiceBreakdownSummary

`func (o *TraceDetail) GetCriticalPathServiceBreakdownSummary() CriticalPathServiceBreakdownSummary`

GetCriticalPathServiceBreakdownSummary returns the CriticalPathServiceBreakdownSummary field if non-nil, zero value otherwise.

### GetCriticalPathServiceBreakdownSummaryOk

`func (o *TraceDetail) GetCriticalPathServiceBreakdownSummaryOk() (*CriticalPathServiceBreakdownSummary, bool)`

GetCriticalPathServiceBreakdownSummaryOk returns a tuple with the CriticalPathServiceBreakdownSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalPathServiceBreakdownSummary

`func (o *TraceDetail) SetCriticalPathServiceBreakdownSummary(v CriticalPathServiceBreakdownSummary)`

SetCriticalPathServiceBreakdownSummary sets CriticalPathServiceBreakdownSummary field to given value.

### HasCriticalPathServiceBreakdownSummary

`func (o *TraceDetail) HasCriticalPathServiceBreakdownSummary() bool`

HasCriticalPathServiceBreakdownSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


