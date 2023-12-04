# TraceFieldValuesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldValues** | **[]string** | List of filter field values. | 
**TotalCount** | **int64** | Total number of values for a field matching the query. Can be approximated when it&#39;s above 3000. | 
**Next** | Pointer to **string** | Next continuation token. | [optional] 

## Methods

### NewTraceFieldValuesResponse

`func NewTraceFieldValuesResponse(fieldValues []string, totalCount int64, ) *TraceFieldValuesResponse`

NewTraceFieldValuesResponse instantiates a new TraceFieldValuesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceFieldValuesResponseWithDefaults

`func NewTraceFieldValuesResponseWithDefaults() *TraceFieldValuesResponse`

NewTraceFieldValuesResponseWithDefaults instantiates a new TraceFieldValuesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldValues

`func (o *TraceFieldValuesResponse) GetFieldValues() []string`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *TraceFieldValuesResponse) GetFieldValuesOk() (*[]string, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *TraceFieldValuesResponse) SetFieldValues(v []string)`

SetFieldValues sets FieldValues field to given value.


### GetTotalCount

`func (o *TraceFieldValuesResponse) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TraceFieldValuesResponse) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TraceFieldValuesResponse) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetNext

`func (o *TraceFieldValuesResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *TraceFieldValuesResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *TraceFieldValuesResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *TraceFieldValuesResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


