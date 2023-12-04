# SpansTimeGroupBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldValue** | **string** | A fixed interval grouping in the following format &lt;#&gt;&lt;time_period&gt;,  supported &lt;time_period&gt; values are weeks (w), days (d), hours (h), minutes (m), and seconds (s).  | 

## Methods

### NewSpansTimeGroupBy

`func NewSpansTimeGroupBy(fieldValue string, ) *SpansTimeGroupBy`

NewSpansTimeGroupBy instantiates a new SpansTimeGroupBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpansTimeGroupByWithDefaults

`func NewSpansTimeGroupByWithDefaults() *SpansTimeGroupBy`

NewSpansTimeGroupByWithDefaults instantiates a new SpansTimeGroupBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldValue

`func (o *SpansTimeGroupBy) GetFieldValue() string`

GetFieldValue returns the FieldValue field if non-nil, zero value otherwise.

### GetFieldValueOk

`func (o *SpansTimeGroupBy) GetFieldValueOk() (*string, bool)`

GetFieldValueOk returns a tuple with the FieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValue

`func (o *SpansTimeGroupBy) SetFieldValue(v string)`

SetFieldValue sets FieldValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


