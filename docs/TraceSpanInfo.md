# TraceSpanInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of this span. Possible values: &#x60;TraceHttpSpanInfo&#x60;, &#x60;TraceDbSpanInfo&#x60;, &#x60;TraceMessageBusSpanInfo&#x60;. | 

## Methods

### NewTraceSpanInfo

`func NewTraceSpanInfo(type_ string, ) *TraceSpanInfo`

NewTraceSpanInfo instantiates a new TraceSpanInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceSpanInfoWithDefaults

`func NewTraceSpanInfoWithDefaults() *TraceSpanInfo`

NewTraceSpanInfoWithDefaults instantiates a new TraceSpanInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TraceSpanInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TraceSpanInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TraceSpanInfo) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


