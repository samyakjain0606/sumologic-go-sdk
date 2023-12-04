# SlosLibraryBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the slo or folder. | 
**Description** | Pointer to **string** | Description of the slo or folder. | [optional] [default to ""]
**Type** | **string** | Type of the object model. Valid values:   1) SlosLibrarySlo   2) SlosLibraryFolder | 

## Methods

### NewSlosLibraryBase

`func NewSlosLibraryBase(name string, type_ string, ) *SlosLibraryBase`

NewSlosLibraryBase instantiates a new SlosLibraryBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlosLibraryBaseWithDefaults

`func NewSlosLibraryBaseWithDefaults() *SlosLibraryBase`

NewSlosLibraryBaseWithDefaults instantiates a new SlosLibraryBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SlosLibraryBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlosLibraryBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlosLibraryBase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SlosLibraryBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SlosLibraryBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SlosLibraryBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SlosLibraryBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *SlosLibraryBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SlosLibraryBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SlosLibraryBase) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


