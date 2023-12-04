# TraceSpan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the span. | 
**ParentId** | Pointer to **string** | Identifier of the parent span, if any. If the span has no parent it&#39;s considered a root span. | [optional] 
**OperationName** | **string** | The name of the operation given to the span. | 
**Resource** | Pointer to **string** | The name of the resource attached to the span. | [optional] 
**Service** | Pointer to **string** | The name of the service this span is part of. | [optional] 
**ServiceColor** | Pointer to **string** | Color hex code assigned to the service. | [optional] 
**ServiceType** | Pointer to **string** | Defines type of service. | [optional] 
**Duration** | **int64** | Number of nanoseconds the span lasted. | 
**StartedAt** | **time.Time** | Date and time the span was started in the [ISO 8601 / RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**Status** | [**TraceSpanStatus**](TraceSpanStatus.md) |  | 
**Kind** | Pointer to **string** | Span kind describes the relationship between the Span, its parents, and its children in a Trace. Possible values: &#x60;CLIENT&#x60;, &#x60;SERVER&#x60;, &#x60;PRODUCER&#x60;, &#x60;CONSUMER&#x60;, &#x60;INTERNAL&#x60;. | [optional] 
**RemoteService** | Pointer to **string** | Name of the possible remote span&#39;s service. | [optional] 
**RemoteServiceColor** | Pointer to **string** | Color hex code assigned to the remote service. | [optional] 
**RemoteServiceType** | Pointer to **string** | Defines type of service. | [optional] 
**Info** | Pointer to [**TraceSpanInfo**](TraceSpanInfo.md) |  | [optional] 
**NumberOfLinks** | Pointer to **int32** | Number of span links in this span. | [optional] 

## Methods

### NewTraceSpan

`func NewTraceSpan(id string, operationName string, duration int64, startedAt time.Time, status TraceSpanStatus, ) *TraceSpan`

NewTraceSpan instantiates a new TraceSpan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceSpanWithDefaults

`func NewTraceSpanWithDefaults() *TraceSpan`

NewTraceSpanWithDefaults instantiates a new TraceSpan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TraceSpan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TraceSpan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TraceSpan) SetId(v string)`

SetId sets Id field to given value.


### GetParentId

`func (o *TraceSpan) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *TraceSpan) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *TraceSpan) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *TraceSpan) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetOperationName

`func (o *TraceSpan) GetOperationName() string`

GetOperationName returns the OperationName field if non-nil, zero value otherwise.

### GetOperationNameOk

`func (o *TraceSpan) GetOperationNameOk() (*string, bool)`

GetOperationNameOk returns a tuple with the OperationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationName

`func (o *TraceSpan) SetOperationName(v string)`

SetOperationName sets OperationName field to given value.


### GetResource

`func (o *TraceSpan) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *TraceSpan) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *TraceSpan) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *TraceSpan) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetService

`func (o *TraceSpan) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *TraceSpan) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *TraceSpan) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *TraceSpan) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServiceColor

`func (o *TraceSpan) GetServiceColor() string`

GetServiceColor returns the ServiceColor field if non-nil, zero value otherwise.

### GetServiceColorOk

`func (o *TraceSpan) GetServiceColorOk() (*string, bool)`

GetServiceColorOk returns a tuple with the ServiceColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceColor

`func (o *TraceSpan) SetServiceColor(v string)`

SetServiceColor sets ServiceColor field to given value.

### HasServiceColor

`func (o *TraceSpan) HasServiceColor() bool`

HasServiceColor returns a boolean if a field has been set.

### GetServiceType

`func (o *TraceSpan) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *TraceSpan) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *TraceSpan) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *TraceSpan) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetDuration

`func (o *TraceSpan) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TraceSpan) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TraceSpan) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetStartedAt

`func (o *TraceSpan) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *TraceSpan) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *TraceSpan) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetStatus

`func (o *TraceSpan) GetStatus() TraceSpanStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TraceSpan) GetStatusOk() (*TraceSpanStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TraceSpan) SetStatus(v TraceSpanStatus)`

SetStatus sets Status field to given value.


### GetKind

`func (o *TraceSpan) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TraceSpan) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TraceSpan) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TraceSpan) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetRemoteService

`func (o *TraceSpan) GetRemoteService() string`

GetRemoteService returns the RemoteService field if non-nil, zero value otherwise.

### GetRemoteServiceOk

`func (o *TraceSpan) GetRemoteServiceOk() (*string, bool)`

GetRemoteServiceOk returns a tuple with the RemoteService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteService

`func (o *TraceSpan) SetRemoteService(v string)`

SetRemoteService sets RemoteService field to given value.

### HasRemoteService

`func (o *TraceSpan) HasRemoteService() bool`

HasRemoteService returns a boolean if a field has been set.

### GetRemoteServiceColor

`func (o *TraceSpan) GetRemoteServiceColor() string`

GetRemoteServiceColor returns the RemoteServiceColor field if non-nil, zero value otherwise.

### GetRemoteServiceColorOk

`func (o *TraceSpan) GetRemoteServiceColorOk() (*string, bool)`

GetRemoteServiceColorOk returns a tuple with the RemoteServiceColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteServiceColor

`func (o *TraceSpan) SetRemoteServiceColor(v string)`

SetRemoteServiceColor sets RemoteServiceColor field to given value.

### HasRemoteServiceColor

`func (o *TraceSpan) HasRemoteServiceColor() bool`

HasRemoteServiceColor returns a boolean if a field has been set.

### GetRemoteServiceType

`func (o *TraceSpan) GetRemoteServiceType() string`

GetRemoteServiceType returns the RemoteServiceType field if non-nil, zero value otherwise.

### GetRemoteServiceTypeOk

`func (o *TraceSpan) GetRemoteServiceTypeOk() (*string, bool)`

GetRemoteServiceTypeOk returns a tuple with the RemoteServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteServiceType

`func (o *TraceSpan) SetRemoteServiceType(v string)`

SetRemoteServiceType sets RemoteServiceType field to given value.

### HasRemoteServiceType

`func (o *TraceSpan) HasRemoteServiceType() bool`

HasRemoteServiceType returns a boolean if a field has been set.

### GetInfo

`func (o *TraceSpan) GetInfo() TraceSpanInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TraceSpan) GetInfoOk() (*TraceSpanInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TraceSpan) SetInfo(v TraceSpanInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TraceSpan) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetNumberOfLinks

`func (o *TraceSpan) GetNumberOfLinks() int32`

GetNumberOfLinks returns the NumberOfLinks field if non-nil, zero value otherwise.

### GetNumberOfLinksOk

`func (o *TraceSpan) GetNumberOfLinksOk() (*int32, bool)`

GetNumberOfLinksOk returns a tuple with the NumberOfLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfLinks

`func (o *TraceSpan) SetNumberOfLinks(v int32)`

SetNumberOfLinks sets NumberOfLinks field to given value.

### HasNumberOfLinks

`func (o *TraceSpan) HasNumberOfLinks() bool`

HasNumberOfLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


