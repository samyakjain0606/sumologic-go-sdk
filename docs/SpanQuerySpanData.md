# SpanQuerySpanData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpanId** | Pointer to **string** | Identifier of the span. | [optional] 
**TraceId** | Pointer to **string** | Identifier of the trace. | [optional] 
**ParentSpanId** | Pointer to **string** | Identifier of the parent span, if any. If the span has no parent it&#39;s considered a root span. | [optional] 
**OperationName** | Pointer to **string** | The name of the operation given to the span. | [optional] 
**Service** | Pointer to **string** | The name of the service this span is part of. | [optional] 
**RemoteService** | Pointer to **string** | Name of the possible remote span&#39;s service. | [optional] 
**Duration** | **int64** | Number of nanoseconds the span lasted. | 
**StartedAt** | **time.Time** | Date and time the span was started in [ISO 8601 / RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**Status** | Pointer to [**TraceSpanStatus**](TraceSpanStatus.md) |  | [optional] 
**Kind** | Pointer to **string** | Span kind describes the relationship between the Span, its parents, and its children in a Trace. Possible values: &#x60;CLIENT&#x60;, &#x60;SERVER&#x60;, &#x60;PRODUCER&#x60;, &#x60;CONSUMER&#x60;, &#x60;INTERNAL&#x60;. | [optional] 
**TagsJSON** | Pointer to **string** | Tags attached to this span as JSON. | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata attached to the span. | [optional] 

## Methods

### NewSpanQuerySpanData

`func NewSpanQuerySpanData(duration int64, startedAt time.Time, ) *SpanQuerySpanData`

NewSpanQuerySpanData instantiates a new SpanQuerySpanData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanQuerySpanDataWithDefaults

`func NewSpanQuerySpanDataWithDefaults() *SpanQuerySpanData`

NewSpanQuerySpanDataWithDefaults instantiates a new SpanQuerySpanData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpanId

`func (o *SpanQuerySpanData) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *SpanQuerySpanData) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *SpanQuerySpanData) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.

### HasSpanId

`func (o *SpanQuerySpanData) HasSpanId() bool`

HasSpanId returns a boolean if a field has been set.

### GetTraceId

`func (o *SpanQuerySpanData) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *SpanQuerySpanData) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *SpanQuerySpanData) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *SpanQuerySpanData) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetParentSpanId

`func (o *SpanQuerySpanData) GetParentSpanId() string`

GetParentSpanId returns the ParentSpanId field if non-nil, zero value otherwise.

### GetParentSpanIdOk

`func (o *SpanQuerySpanData) GetParentSpanIdOk() (*string, bool)`

GetParentSpanIdOk returns a tuple with the ParentSpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpanId

`func (o *SpanQuerySpanData) SetParentSpanId(v string)`

SetParentSpanId sets ParentSpanId field to given value.

### HasParentSpanId

`func (o *SpanQuerySpanData) HasParentSpanId() bool`

HasParentSpanId returns a boolean if a field has been set.

### GetOperationName

`func (o *SpanQuerySpanData) GetOperationName() string`

GetOperationName returns the OperationName field if non-nil, zero value otherwise.

### GetOperationNameOk

`func (o *SpanQuerySpanData) GetOperationNameOk() (*string, bool)`

GetOperationNameOk returns a tuple with the OperationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationName

`func (o *SpanQuerySpanData) SetOperationName(v string)`

SetOperationName sets OperationName field to given value.

### HasOperationName

`func (o *SpanQuerySpanData) HasOperationName() bool`

HasOperationName returns a boolean if a field has been set.

### GetService

`func (o *SpanQuerySpanData) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *SpanQuerySpanData) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *SpanQuerySpanData) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *SpanQuerySpanData) HasService() bool`

HasService returns a boolean if a field has been set.

### GetRemoteService

`func (o *SpanQuerySpanData) GetRemoteService() string`

GetRemoteService returns the RemoteService field if non-nil, zero value otherwise.

### GetRemoteServiceOk

`func (o *SpanQuerySpanData) GetRemoteServiceOk() (*string, bool)`

GetRemoteServiceOk returns a tuple with the RemoteService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteService

`func (o *SpanQuerySpanData) SetRemoteService(v string)`

SetRemoteService sets RemoteService field to given value.

### HasRemoteService

`func (o *SpanQuerySpanData) HasRemoteService() bool`

HasRemoteService returns a boolean if a field has been set.

### GetDuration

`func (o *SpanQuerySpanData) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SpanQuerySpanData) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SpanQuerySpanData) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetStartedAt

`func (o *SpanQuerySpanData) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *SpanQuerySpanData) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *SpanQuerySpanData) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetStatus

`func (o *SpanQuerySpanData) GetStatus() TraceSpanStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpanQuerySpanData) GetStatusOk() (*TraceSpanStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SpanQuerySpanData) SetStatus(v TraceSpanStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SpanQuerySpanData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetKind

`func (o *SpanQuerySpanData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SpanQuerySpanData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SpanQuerySpanData) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SpanQuerySpanData) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetTagsJSON

`func (o *SpanQuerySpanData) GetTagsJSON() string`

GetTagsJSON returns the TagsJSON field if non-nil, zero value otherwise.

### GetTagsJSONOk

`func (o *SpanQuerySpanData) GetTagsJSONOk() (*string, bool)`

GetTagsJSONOk returns a tuple with the TagsJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsJSON

`func (o *SpanQuerySpanData) SetTagsJSON(v string)`

SetTagsJSON sets TagsJSON field to given value.

### HasTagsJSON

`func (o *SpanQuerySpanData) HasTagsJSON() bool`

HasTagsJSON returns a boolean if a field has been set.

### GetMetadata

`func (o *SpanQuerySpanData) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SpanQuerySpanData) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SpanQuerySpanData) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SpanQuerySpanData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


