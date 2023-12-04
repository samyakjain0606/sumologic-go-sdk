# GroupFieldsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupFields** | **[]string** | List of group fields | 
**IsQueryAggregate** | **bool** | Whether or not the queries are aggregate. | [default to false]

## Methods

### NewGroupFieldsResponse

`func NewGroupFieldsResponse(groupFields []string, isQueryAggregate bool, ) *GroupFieldsResponse`

NewGroupFieldsResponse instantiates a new GroupFieldsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupFieldsResponseWithDefaults

`func NewGroupFieldsResponseWithDefaults() *GroupFieldsResponse`

NewGroupFieldsResponseWithDefaults instantiates a new GroupFieldsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupFields

`func (o *GroupFieldsResponse) GetGroupFields() []string`

GetGroupFields returns the GroupFields field if non-nil, zero value otherwise.

### GetGroupFieldsOk

`func (o *GroupFieldsResponse) GetGroupFieldsOk() (*[]string, bool)`

GetGroupFieldsOk returns a tuple with the GroupFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupFields

`func (o *GroupFieldsResponse) SetGroupFields(v []string)`

SetGroupFields sets GroupFields field to given value.


### GetIsQueryAggregate

`func (o *GroupFieldsResponse) GetIsQueryAggregate() bool`

GetIsQueryAggregate returns the IsQueryAggregate field if non-nil, zero value otherwise.

### GetIsQueryAggregateOk

`func (o *GroupFieldsResponse) GetIsQueryAggregateOk() (*bool, bool)`

GetIsQueryAggregateOk returns a tuple with the IsQueryAggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueryAggregate

`func (o *GroupFieldsResponse) SetIsQueryAggregate(v bool)`

SetIsQueryAggregate sets IsQueryAggregate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


