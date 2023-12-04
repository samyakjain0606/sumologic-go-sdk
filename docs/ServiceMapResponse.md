# ServiceMapResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | [**[]ServiceMapNode**](ServiceMapNode.md) | List of service map nodes. | 
**Edges** | [**[]ServiceMapEdge**](ServiceMapEdge.md) | List of service map edges. | 

## Methods

### NewServiceMapResponse

`func NewServiceMapResponse(nodes []ServiceMapNode, edges []ServiceMapEdge, ) *ServiceMapResponse`

NewServiceMapResponse instantiates a new ServiceMapResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMapResponseWithDefaults

`func NewServiceMapResponseWithDefaults() *ServiceMapResponse`

NewServiceMapResponseWithDefaults instantiates a new ServiceMapResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *ServiceMapResponse) GetNodes() []ServiceMapNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ServiceMapResponse) GetNodesOk() (*[]ServiceMapNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ServiceMapResponse) SetNodes(v []ServiceMapNode)`

SetNodes sets Nodes field to given value.


### GetEdges

`func (o *ServiceMapResponse) GetEdges() []ServiceMapEdge`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *ServiceMapResponse) GetEdgesOk() (*[]ServiceMapEdge, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *ServiceMapResponse) SetEdges(v []ServiceMapEdge)`

SetEdges sets Edges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


