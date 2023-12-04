# MonitorTemplatesLibraryBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the monitortemplate or folder. | 
**Description** | Pointer to **string** | Description of the monitortemplate or folder. | [optional] [default to ""]
**Type** | **string** | Type of the object model. Valid values:   1) MonitorTemplatesLibraryMonitortemplate   2) MonitorTemplatesLibraryFolder | 

## Methods

### NewMonitorTemplatesLibraryBase

`func NewMonitorTemplatesLibraryBase(name string, type_ string, ) *MonitorTemplatesLibraryBase`

NewMonitorTemplatesLibraryBase instantiates a new MonitorTemplatesLibraryBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorTemplatesLibraryBaseWithDefaults

`func NewMonitorTemplatesLibraryBaseWithDefaults() *MonitorTemplatesLibraryBase`

NewMonitorTemplatesLibraryBaseWithDefaults instantiates a new MonitorTemplatesLibraryBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MonitorTemplatesLibraryBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitorTemplatesLibraryBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitorTemplatesLibraryBase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MonitorTemplatesLibraryBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MonitorTemplatesLibraryBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MonitorTemplatesLibraryBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MonitorTemplatesLibraryBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *MonitorTemplatesLibraryBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorTemplatesLibraryBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorTemplatesLibraryBase) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


