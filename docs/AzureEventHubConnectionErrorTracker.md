# AzureEventHubConnectionErrorTracker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullyQualifiedNamespace** | Pointer to **string** | The namespace of the associated source. | [optional] 
**EventHub** | Pointer to **string** | The event hub name of the associated source. | [optional] 

## Methods

### NewAzureEventHubConnectionErrorTracker

`func NewAzureEventHubConnectionErrorTracker() *AzureEventHubConnectionErrorTracker`

NewAzureEventHubConnectionErrorTracker instantiates a new AzureEventHubConnectionErrorTracker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureEventHubConnectionErrorTrackerWithDefaults

`func NewAzureEventHubConnectionErrorTrackerWithDefaults() *AzureEventHubConnectionErrorTracker`

NewAzureEventHubConnectionErrorTrackerWithDefaults instantiates a new AzureEventHubConnectionErrorTracker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullyQualifiedNamespace

`func (o *AzureEventHubConnectionErrorTracker) GetFullyQualifiedNamespace() string`

GetFullyQualifiedNamespace returns the FullyQualifiedNamespace field if non-nil, zero value otherwise.

### GetFullyQualifiedNamespaceOk

`func (o *AzureEventHubConnectionErrorTracker) GetFullyQualifiedNamespaceOk() (*string, bool)`

GetFullyQualifiedNamespaceOk returns a tuple with the FullyQualifiedNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedNamespace

`func (o *AzureEventHubConnectionErrorTracker) SetFullyQualifiedNamespace(v string)`

SetFullyQualifiedNamespace sets FullyQualifiedNamespace field to given value.

### HasFullyQualifiedNamespace

`func (o *AzureEventHubConnectionErrorTracker) HasFullyQualifiedNamespace() bool`

HasFullyQualifiedNamespace returns a boolean if a field has been set.

### GetEventHub

`func (o *AzureEventHubConnectionErrorTracker) GetEventHub() string`

GetEventHub returns the EventHub field if non-nil, zero value otherwise.

### GetEventHubOk

`func (o *AzureEventHubConnectionErrorTracker) GetEventHubOk() (*string, bool)`

GetEventHubOk returns a tuple with the EventHub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHub

`func (o *AzureEventHubConnectionErrorTracker) SetEventHub(v string)`

SetEventHub sets EventHub field to given value.

### HasEventHub

`func (o *AzureEventHubConnectionErrorTracker) HasEventHub() bool`

HasEventHub returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


