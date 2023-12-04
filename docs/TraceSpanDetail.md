# TraceSpanDetail

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
**ErrorMessage** | Pointer to **string** | Produced error message (could be a stack trace, database error code, ..) | [optional] 
**Fields** | Pointer to [**map[string]TracingValue**](TracingValue.md) | Fields attached to this span. | [optional] 
**CriticalPathContribution** | Pointer to [**TraceSpanCriticalPathContribution**](TraceSpanCriticalPathContribution.md) |  | [optional] 
**Logs** | Pointer to **[]string** | Logs attached to this span. | [optional] 
**Events** | Pointer to [**[]SpanEvent**](SpanEvent.md) | Events attached to this span. | [optional] 
**Links** | Pointer to [**[]SpanLink**](SpanLink.md) | List of casually related spans. | [optional] 

## Methods

### NewTraceSpanDetail

`func NewTraceSpanDetail(id string, operationName string, duration int64, startedAt time.Time, status TraceSpanStatus, ) *TraceSpanDetail`

NewTraceSpanDetail instantiates a new TraceSpanDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceSpanDetailWithDefaults

`func NewTraceSpanDetailWithDefaults() *TraceSpanDetail`

NewTraceSpanDetailWithDefaults instantiates a new TraceSpanDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TraceSpanDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TraceSpanDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TraceSpanDetail) SetId(v string)`

SetId sets Id field to given value.


### GetParentId

`func (o *TraceSpanDetail) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *TraceSpanDetail) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *TraceSpanDetail) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *TraceSpanDetail) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetOperationName

`func (o *TraceSpanDetail) GetOperationName() string`

GetOperationName returns the OperationName field if non-nil, zero value otherwise.

### GetOperationNameOk

`func (o *TraceSpanDetail) GetOperationNameOk() (*string, bool)`

GetOperationNameOk returns a tuple with the OperationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationName

`func (o *TraceSpanDetail) SetOperationName(v string)`

SetOperationName sets OperationName field to given value.


### GetResource

`func (o *TraceSpanDetail) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *TraceSpanDetail) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *TraceSpanDetail) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *TraceSpanDetail) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetService

`func (o *TraceSpanDetail) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *TraceSpanDetail) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *TraceSpanDetail) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *TraceSpanDetail) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServiceColor

`func (o *TraceSpanDetail) GetServiceColor() string`

GetServiceColor returns the ServiceColor field if non-nil, zero value otherwise.

### GetServiceColorOk

`func (o *TraceSpanDetail) GetServiceColorOk() (*string, bool)`

GetServiceColorOk returns a tuple with the ServiceColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceColor

`func (o *TraceSpanDetail) SetServiceColor(v string)`

SetServiceColor sets ServiceColor field to given value.

### HasServiceColor

`func (o *TraceSpanDetail) HasServiceColor() bool`

HasServiceColor returns a boolean if a field has been set.

### GetServiceType

`func (o *TraceSpanDetail) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *TraceSpanDetail) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *TraceSpanDetail) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *TraceSpanDetail) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetDuration

`func (o *TraceSpanDetail) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TraceSpanDetail) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TraceSpanDetail) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetStartedAt

`func (o *TraceSpanDetail) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *TraceSpanDetail) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *TraceSpanDetail) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetStatus

`func (o *TraceSpanDetail) GetStatus() TraceSpanStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TraceSpanDetail) GetStatusOk() (*TraceSpanStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TraceSpanDetail) SetStatus(v TraceSpanStatus)`

SetStatus sets Status field to given value.


### GetKind

`func (o *TraceSpanDetail) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TraceSpanDetail) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TraceSpanDetail) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TraceSpanDetail) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetRemoteService

`func (o *TraceSpanDetail) GetRemoteService() string`

GetRemoteService returns the RemoteService field if non-nil, zero value otherwise.

### GetRemoteServiceOk

`func (o *TraceSpanDetail) GetRemoteServiceOk() (*string, bool)`

GetRemoteServiceOk returns a tuple with the RemoteService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteService

`func (o *TraceSpanDetail) SetRemoteService(v string)`

SetRemoteService sets RemoteService field to given value.

### HasRemoteService

`func (o *TraceSpanDetail) HasRemoteService() bool`

HasRemoteService returns a boolean if a field has been set.

### GetRemoteServiceColor

`func (o *TraceSpanDetail) GetRemoteServiceColor() string`

GetRemoteServiceColor returns the RemoteServiceColor field if non-nil, zero value otherwise.

### GetRemoteServiceColorOk

`func (o *TraceSpanDetail) GetRemoteServiceColorOk() (*string, bool)`

GetRemoteServiceColorOk returns a tuple with the RemoteServiceColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteServiceColor

`func (o *TraceSpanDetail) SetRemoteServiceColor(v string)`

SetRemoteServiceColor sets RemoteServiceColor field to given value.

### HasRemoteServiceColor

`func (o *TraceSpanDetail) HasRemoteServiceColor() bool`

HasRemoteServiceColor returns a boolean if a field has been set.

### GetRemoteServiceType

`func (o *TraceSpanDetail) GetRemoteServiceType() string`

GetRemoteServiceType returns the RemoteServiceType field if non-nil, zero value otherwise.

### GetRemoteServiceTypeOk

`func (o *TraceSpanDetail) GetRemoteServiceTypeOk() (*string, bool)`

GetRemoteServiceTypeOk returns a tuple with the RemoteServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteServiceType

`func (o *TraceSpanDetail) SetRemoteServiceType(v string)`

SetRemoteServiceType sets RemoteServiceType field to given value.

### HasRemoteServiceType

`func (o *TraceSpanDetail) HasRemoteServiceType() bool`

HasRemoteServiceType returns a boolean if a field has been set.

### GetInfo

`func (o *TraceSpanDetail) GetInfo() TraceSpanInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TraceSpanDetail) GetInfoOk() (*TraceSpanInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TraceSpanDetail) SetInfo(v TraceSpanInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TraceSpanDetail) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetNumberOfLinks

`func (o *TraceSpanDetail) GetNumberOfLinks() int32`

GetNumberOfLinks returns the NumberOfLinks field if non-nil, zero value otherwise.

### GetNumberOfLinksOk

`func (o *TraceSpanDetail) GetNumberOfLinksOk() (*int32, bool)`

GetNumberOfLinksOk returns a tuple with the NumberOfLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfLinks

`func (o *TraceSpanDetail) SetNumberOfLinks(v int32)`

SetNumberOfLinks sets NumberOfLinks field to given value.

### HasNumberOfLinks

`func (o *TraceSpanDetail) HasNumberOfLinks() bool`

HasNumberOfLinks returns a boolean if a field has been set.

### GetErrorMessage

`func (o *TraceSpanDetail) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TraceSpanDetail) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TraceSpanDetail) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TraceSpanDetail) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetFields

`func (o *TraceSpanDetail) GetFields() map[string]TracingValue`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TraceSpanDetail) GetFieldsOk() (*map[string]TracingValue, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TraceSpanDetail) SetFields(v map[string]TracingValue)`

SetFields sets Fields field to given value.

### HasFields

`func (o *TraceSpanDetail) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetCriticalPathContribution

`func (o *TraceSpanDetail) GetCriticalPathContribution() TraceSpanCriticalPathContribution`

GetCriticalPathContribution returns the CriticalPathContribution field if non-nil, zero value otherwise.

### GetCriticalPathContributionOk

`func (o *TraceSpanDetail) GetCriticalPathContributionOk() (*TraceSpanCriticalPathContribution, bool)`

GetCriticalPathContributionOk returns a tuple with the CriticalPathContribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalPathContribution

`func (o *TraceSpanDetail) SetCriticalPathContribution(v TraceSpanCriticalPathContribution)`

SetCriticalPathContribution sets CriticalPathContribution field to given value.

### HasCriticalPathContribution

`func (o *TraceSpanDetail) HasCriticalPathContribution() bool`

HasCriticalPathContribution returns a boolean if a field has been set.

### GetLogs

`func (o *TraceSpanDetail) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *TraceSpanDetail) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *TraceSpanDetail) SetLogs(v []string)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *TraceSpanDetail) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetEvents

`func (o *TraceSpanDetail) GetEvents() []SpanEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *TraceSpanDetail) GetEventsOk() (*[]SpanEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *TraceSpanDetail) SetEvents(v []SpanEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *TraceSpanDetail) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetLinks

`func (o *TraceSpanDetail) GetLinks() []SpanLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TraceSpanDetail) GetLinksOk() (*[]SpanLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TraceSpanDetail) SetLinks(v []SpanLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TraceSpanDetail) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


