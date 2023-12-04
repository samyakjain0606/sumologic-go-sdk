# GetDataForwardingDestinations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** | Next continuation token. | [optional] 
**Data** | Pointer to [**[]BucketDefinition**](BucketDefinition.md) | List of data forwarding destinations. | [optional] 

## Methods

### NewGetDataForwardingDestinations

`func NewGetDataForwardingDestinations() *GetDataForwardingDestinations`

NewGetDataForwardingDestinations instantiates a new GetDataForwardingDestinations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDataForwardingDestinationsWithDefaults

`func NewGetDataForwardingDestinationsWithDefaults() *GetDataForwardingDestinations`

NewGetDataForwardingDestinationsWithDefaults instantiates a new GetDataForwardingDestinations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *GetDataForwardingDestinations) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *GetDataForwardingDestinations) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *GetDataForwardingDestinations) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *GetDataForwardingDestinations) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetData

`func (o *GetDataForwardingDestinations) GetData() []BucketDefinition`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetDataForwardingDestinations) GetDataOk() (*[]BucketDefinition, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetDataForwardingDestinations) SetData(v []BucketDefinition)`

SetData sets Data field to given value.

### HasData

`func (o *GetDataForwardingDestinations) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


