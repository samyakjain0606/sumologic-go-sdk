# PaginatedOTCollectorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Search** | Pointer to **string** | search by collector id or free text search on collector properties. | [optional] 
**Filters** | Pointer to [**PaginatedOTCollectorsRequestFilters**](PaginatedOTCollectorsRequestFilters.md) |  | [optional] 
**SortBy** | Pointer to **string** | parameter which is used for sorting. | [optional] 
**Next** | Pointer to **string** | parameter which is used for fetching next set of results. | [optional] 
**Limit** | Pointer to **int32** | parameter which is used for limiting number of otCollectors on a page. | [optional] 
**IncludeCount** | Pointer to **NullableBool** | count of filtered otCollectors. | [optional] 

## Methods

### NewPaginatedOTCollectorsRequest

`func NewPaginatedOTCollectorsRequest() *PaginatedOTCollectorsRequest`

NewPaginatedOTCollectorsRequest instantiates a new PaginatedOTCollectorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedOTCollectorsRequestWithDefaults

`func NewPaginatedOTCollectorsRequestWithDefaults() *PaginatedOTCollectorsRequest`

NewPaginatedOTCollectorsRequestWithDefaults instantiates a new PaginatedOTCollectorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearch

`func (o *PaginatedOTCollectorsRequest) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *PaginatedOTCollectorsRequest) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *PaginatedOTCollectorsRequest) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *PaginatedOTCollectorsRequest) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetFilters

`func (o *PaginatedOTCollectorsRequest) GetFilters() PaginatedOTCollectorsRequestFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PaginatedOTCollectorsRequest) GetFiltersOk() (*PaginatedOTCollectorsRequestFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PaginatedOTCollectorsRequest) SetFilters(v PaginatedOTCollectorsRequestFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PaginatedOTCollectorsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSortBy

`func (o *PaginatedOTCollectorsRequest) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *PaginatedOTCollectorsRequest) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *PaginatedOTCollectorsRequest) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *PaginatedOTCollectorsRequest) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedOTCollectorsRequest) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedOTCollectorsRequest) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedOTCollectorsRequest) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedOTCollectorsRequest) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetLimit

`func (o *PaginatedOTCollectorsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PaginatedOTCollectorsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PaginatedOTCollectorsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PaginatedOTCollectorsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetIncludeCount

`func (o *PaginatedOTCollectorsRequest) GetIncludeCount() bool`

GetIncludeCount returns the IncludeCount field if non-nil, zero value otherwise.

### GetIncludeCountOk

`func (o *PaginatedOTCollectorsRequest) GetIncludeCountOk() (*bool, bool)`

GetIncludeCountOk returns a tuple with the IncludeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCount

`func (o *PaginatedOTCollectorsRequest) SetIncludeCount(v bool)`

SetIncludeCount sets IncludeCount field to given value.

### HasIncludeCount

`func (o *PaginatedOTCollectorsRequest) HasIncludeCount() bool`

HasIncludeCount returns a boolean if a field has been set.

### SetIncludeCountNil

`func (o *PaginatedOTCollectorsRequest) SetIncludeCountNil(b bool)`

 SetIncludeCountNil sets the value for IncludeCount to be an explicit nil

### UnsetIncludeCount
`func (o *PaginatedOTCollectorsRequest) UnsetIncludeCount()`

UnsetIncludeCount ensures that no value is present for IncludeCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


