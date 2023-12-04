# PaginatedLogSearches

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogSearches** | [**[]LogSearch**](LogSearch.md) | List of log searches. | 
**Token** | Pointer to **string** | Next continuation token. &#x60;token&#x60; is set to null when no more pages are left. | [optional] 

## Methods

### NewPaginatedLogSearches

`func NewPaginatedLogSearches(logSearches []LogSearch, ) *PaginatedLogSearches`

NewPaginatedLogSearches instantiates a new PaginatedLogSearches object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedLogSearchesWithDefaults

`func NewPaginatedLogSearchesWithDefaults() *PaginatedLogSearches`

NewPaginatedLogSearchesWithDefaults instantiates a new PaginatedLogSearches object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogSearches

`func (o *PaginatedLogSearches) GetLogSearches() []LogSearch`

GetLogSearches returns the LogSearches field if non-nil, zero value otherwise.

### GetLogSearchesOk

`func (o *PaginatedLogSearches) GetLogSearchesOk() (*[]LogSearch, bool)`

GetLogSearchesOk returns a tuple with the LogSearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSearches

`func (o *PaginatedLogSearches) SetLogSearches(v []LogSearch)`

SetLogSearches sets LogSearches field to given value.


### GetToken

`func (o *PaginatedLogSearches) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaginatedLogSearches) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaginatedLogSearches) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PaginatedLogSearches) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


