# SliQueryGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryGroupType** | **string** | Type of Query (successful/unsuccessful/total/threshold). | 
**QueryGroup** | [**[]SliQuery**](SliQuery.md) | Group of queries to allow for query arithmetic. | 

## Methods

### NewSliQueryGroup

`func NewSliQueryGroup(queryGroupType string, queryGroup []SliQuery, ) *SliQueryGroup`

NewSliQueryGroup instantiates a new SliQueryGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliQueryGroupWithDefaults

`func NewSliQueryGroupWithDefaults() *SliQueryGroup`

NewSliQueryGroupWithDefaults instantiates a new SliQueryGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryGroupType

`func (o *SliQueryGroup) GetQueryGroupType() string`

GetQueryGroupType returns the QueryGroupType field if non-nil, zero value otherwise.

### GetQueryGroupTypeOk

`func (o *SliQueryGroup) GetQueryGroupTypeOk() (*string, bool)`

GetQueryGroupTypeOk returns a tuple with the QueryGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryGroupType

`func (o *SliQueryGroup) SetQueryGroupType(v string)`

SetQueryGroupType sets QueryGroupType field to given value.


### GetQueryGroup

`func (o *SliQueryGroup) GetQueryGroup() []SliQuery`

GetQueryGroup returns the QueryGroup field if non-nil, zero value otherwise.

### GetQueryGroupOk

`func (o *SliQueryGroup) GetQueryGroupOk() (*[]SliQuery, bool)`

GetQueryGroupOk returns a tuple with the QueryGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryGroup

`func (o *SliQueryGroup) SetQueryGroup(v []SliQuery)`

SetQueryGroup sets QueryGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


