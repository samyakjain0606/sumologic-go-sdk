# SpansQueryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**[]SpansFilter**](SpansFilter.md) | A list of filters for the spans query. | 
**Visualizations** | [**[]SpansVisualization**](SpansVisualization.md) | A list of used visualization methods for the spans query. | 
**GroupBy** | [**[]SpansGroupBy**](SpansGroupBy.md) | A list of group-by clauses for the spans query. | 
**Limit** | [**[]SpansLimitItem**](SpansLimitItem.md) | A list of limits that will be applied to the spans query. | 

## Methods

### NewSpansQueryData

`func NewSpansQueryData(filters []SpansFilter, visualizations []SpansVisualization, groupBy []SpansGroupBy, limit []SpansLimitItem, ) *SpansQueryData`

NewSpansQueryData instantiates a new SpansQueryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpansQueryDataWithDefaults

`func NewSpansQueryDataWithDefaults() *SpansQueryData`

NewSpansQueryDataWithDefaults instantiates a new SpansQueryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *SpansQueryData) GetFilters() []SpansFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SpansQueryData) GetFiltersOk() (*[]SpansFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SpansQueryData) SetFilters(v []SpansFilter)`

SetFilters sets Filters field to given value.


### GetVisualizations

`func (o *SpansQueryData) GetVisualizations() []SpansVisualization`

GetVisualizations returns the Visualizations field if non-nil, zero value otherwise.

### GetVisualizationsOk

`func (o *SpansQueryData) GetVisualizationsOk() (*[]SpansVisualization, bool)`

GetVisualizationsOk returns a tuple with the Visualizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualizations

`func (o *SpansQueryData) SetVisualizations(v []SpansVisualization)`

SetVisualizations sets Visualizations field to given value.


### GetGroupBy

`func (o *SpansQueryData) GetGroupBy() []SpansGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *SpansQueryData) GetGroupByOk() (*[]SpansGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *SpansQueryData) SetGroupBy(v []SpansGroupBy)`

SetGroupBy sets GroupBy field to given value.


### GetLimit

`func (o *SpansQueryData) GetLimit() []SpansLimitItem`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SpansQueryData) GetLimitOk() (*[]SpansLimitItem, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SpansQueryData) SetLimit(v []SpansLimitItem)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


