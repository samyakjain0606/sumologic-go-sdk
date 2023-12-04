# GetRulesAndBucketsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RuleAndBucketDetail**](RuleAndBucketDetail.md) | List of S3 data forwarding rules. | [optional] 
**NextToken** | Pointer to **string** | Next continuation token. | [optional] 

## Methods

### NewGetRulesAndBucketsResult

`func NewGetRulesAndBucketsResult() *GetRulesAndBucketsResult`

NewGetRulesAndBucketsResult instantiates a new GetRulesAndBucketsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRulesAndBucketsResultWithDefaults

`func NewGetRulesAndBucketsResultWithDefaults() *GetRulesAndBucketsResult`

NewGetRulesAndBucketsResultWithDefaults instantiates a new GetRulesAndBucketsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetRulesAndBucketsResult) GetData() []RuleAndBucketDetail`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRulesAndBucketsResult) GetDataOk() (*[]RuleAndBucketDetail, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRulesAndBucketsResult) SetData(v []RuleAndBucketDetail)`

SetData sets Data field to given value.

### HasData

`func (o *GetRulesAndBucketsResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNextToken

`func (o *GetRulesAndBucketsResult) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *GetRulesAndBucketsResult) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *GetRulesAndBucketsResult) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *GetRulesAndBucketsResult) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


