# TraceExistsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exists** | **bool** | Indicates whether the trace with the given trace id exists. | 
**Url** | Pointer to **string** | A path to the trace view page in Sumo Logic UI. | [optional] 

## Methods

### NewTraceExistsResponse

`func NewTraceExistsResponse(exists bool, ) *TraceExistsResponse`

NewTraceExistsResponse instantiates a new TraceExistsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceExistsResponseWithDefaults

`func NewTraceExistsResponseWithDefaults() *TraceExistsResponse`

NewTraceExistsResponseWithDefaults instantiates a new TraceExistsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExists

`func (o *TraceExistsResponse) GetExists() bool`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *TraceExistsResponse) GetExistsOk() (*bool, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *TraceExistsResponse) SetExists(v bool)`

SetExists sets Exists field to given value.


### GetUrl

`func (o *TraceExistsResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TraceExistsResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TraceExistsResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TraceExistsResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


