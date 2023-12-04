# OrderBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | **string** | Field based on which results should be sorted. When not provided, the default behavior is to sort by timestamp descending. Sortable fields values: &#x60;trace_id&#x60;, &#x60;start_timestamp&#x60;, &#x60;duration&#x60;, &#x60;spans_number&#x60;, &#x60;errors&#x60;, &#x60;status_code&#x60;. | 
**Order** | **string** | Type of sorting values - descending or ascending. | [default to "Desc"]

## Methods

### NewOrderBy

`func NewOrderBy(fieldName string, order string, ) *OrderBy`

NewOrderBy instantiates a new OrderBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderByWithDefaults

`func NewOrderByWithDefaults() *OrderBy`

NewOrderByWithDefaults instantiates a new OrderBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *OrderBy) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *OrderBy) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *OrderBy) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetOrder

`func (o *OrderBy) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OrderBy) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *OrderBy) SetOrder(v string)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


