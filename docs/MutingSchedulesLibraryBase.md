# MutingSchedulesLibraryBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the mutingschedule or folder. | 
**Description** | Pointer to **string** | Description of the mutingschedule or folder. | [optional] [default to ""]
**Type** | **string** | Type of the object model. Valid values:   1) MutingSchedulesLibraryMutingschedule   2) MutingSchedulesLibraryFolder | 

## Methods

### NewMutingSchedulesLibraryBase

`func NewMutingSchedulesLibraryBase(name string, type_ string, ) *MutingSchedulesLibraryBase`

NewMutingSchedulesLibraryBase instantiates a new MutingSchedulesLibraryBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutingSchedulesLibraryBaseWithDefaults

`func NewMutingSchedulesLibraryBaseWithDefaults() *MutingSchedulesLibraryBase`

NewMutingSchedulesLibraryBaseWithDefaults instantiates a new MutingSchedulesLibraryBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MutingSchedulesLibraryBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MutingSchedulesLibraryBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MutingSchedulesLibraryBase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MutingSchedulesLibraryBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MutingSchedulesLibraryBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MutingSchedulesLibraryBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MutingSchedulesLibraryBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *MutingSchedulesLibraryBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MutingSchedulesLibraryBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MutingSchedulesLibraryBase) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


