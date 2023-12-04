# PaginatedOTCollectorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]OTCollector**](OTCollector.md) | paginated list of OT Collectors. | 
**Next** | Pointer to **string** | next page token. | [optional] 
**Count** | Pointer to **int32** | count of otCollectors in response. | [optional] 

## Methods

### NewPaginatedOTCollectorsResponse

`func NewPaginatedOTCollectorsResponse(data []OTCollector, ) *PaginatedOTCollectorsResponse`

NewPaginatedOTCollectorsResponse instantiates a new PaginatedOTCollectorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedOTCollectorsResponseWithDefaults

`func NewPaginatedOTCollectorsResponseWithDefaults() *PaginatedOTCollectorsResponse`

NewPaginatedOTCollectorsResponseWithDefaults instantiates a new PaginatedOTCollectorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaginatedOTCollectorsResponse) GetData() []OTCollector`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginatedOTCollectorsResponse) GetDataOk() (*[]OTCollector, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginatedOTCollectorsResponse) SetData(v []OTCollector)`

SetData sets Data field to given value.


### GetNext

`func (o *PaginatedOTCollectorsResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedOTCollectorsResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedOTCollectorsResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedOTCollectorsResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetCount

`func (o *PaginatedOTCollectorsResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedOTCollectorsResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedOTCollectorsResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedOTCollectorsResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


