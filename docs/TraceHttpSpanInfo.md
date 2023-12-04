# TraceHttpSpanInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** | HTTP method of the request for the associated span. | [optional] 
**Url** | Pointer to **string** | URL of the request being handled in this span, in the standard URI format. | [optional] 
**StatusCode** | Pointer to **int32** | HTTP response status code for the associated span. | [optional] 

## Methods

### NewTraceHttpSpanInfo

`func NewTraceHttpSpanInfo() *TraceHttpSpanInfo`

NewTraceHttpSpanInfo instantiates a new TraceHttpSpanInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceHttpSpanInfoWithDefaults

`func NewTraceHttpSpanInfoWithDefaults() *TraceHttpSpanInfo`

NewTraceHttpSpanInfoWithDefaults instantiates a new TraceHttpSpanInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *TraceHttpSpanInfo) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TraceHttpSpanInfo) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TraceHttpSpanInfo) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *TraceHttpSpanInfo) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetUrl

`func (o *TraceHttpSpanInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TraceHttpSpanInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TraceHttpSpanInfo) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TraceHttpSpanInfo) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatusCode

`func (o *TraceHttpSpanInfo) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *TraceHttpSpanInfo) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *TraceHttpSpanInfo) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *TraceHttpSpanInfo) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


