# MonitorTemplatesLibraryBaseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the monitortemplate or folder. | 
**Name** | **string** | Identifier of the monitortemplate or folder. | 
**Description** | **string** | Description of the monitortemplate or folder. | 
**Version** | **int64** | Version of the monitortemplate or folder. | 
**CreatedAt** | **time.Time** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**CreatedBy** | **string** | Identifier of the user who created the resource. | 
**ModifiedAt** | **time.Time** | Last modification timestamp in UTC. | 
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | 
**ParentId** | **string** | Identifier of the parent folder. | 
**ContentType** | **string** | Type of the content. Valid values:   1) Monitortemplate   2) Folder | 
**Type** | **string** | Type of the object model. | 
**IsSystem** | **bool** | System objects are objects provided by Sumo Logic. System objects can only be localized. Non-local fields can&#39;t be updated. | 
**IsMutable** | **bool** | Immutable objects are \&quot;READ-ONLY\&quot;. | 
**Permissions** | Pointer to **[]string** | Aggregated permission summary for the calling user. If detailed permission statements are required, please call list permissions endpoint. | [optional] 

## Methods

### NewMonitorTemplatesLibraryBaseResponse

`func NewMonitorTemplatesLibraryBaseResponse(id string, name string, description string, version int64, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, parentId string, contentType string, type_ string, isSystem bool, isMutable bool, ) *MonitorTemplatesLibraryBaseResponse`

NewMonitorTemplatesLibraryBaseResponse instantiates a new MonitorTemplatesLibraryBaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorTemplatesLibraryBaseResponseWithDefaults

`func NewMonitorTemplatesLibraryBaseResponseWithDefaults() *MonitorTemplatesLibraryBaseResponse`

NewMonitorTemplatesLibraryBaseResponseWithDefaults instantiates a new MonitorTemplatesLibraryBaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MonitorTemplatesLibraryBaseResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MonitorTemplatesLibraryBaseResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MonitorTemplatesLibraryBaseResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MonitorTemplatesLibraryBaseResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitorTemplatesLibraryBaseResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitorTemplatesLibraryBaseResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MonitorTemplatesLibraryBaseResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MonitorTemplatesLibraryBaseResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MonitorTemplatesLibraryBaseResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetVersion

`func (o *MonitorTemplatesLibraryBaseResponse) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MonitorTemplatesLibraryBaseResponse) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MonitorTemplatesLibraryBaseResponse) SetVersion(v int64)`

SetVersion sets Version field to given value.


### GetCreatedAt

`func (o *MonitorTemplatesLibraryBaseResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MonitorTemplatesLibraryBaseResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MonitorTemplatesLibraryBaseResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *MonitorTemplatesLibraryBaseResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MonitorTemplatesLibraryBaseResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MonitorTemplatesLibraryBaseResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *MonitorTemplatesLibraryBaseResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *MonitorTemplatesLibraryBaseResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *MonitorTemplatesLibraryBaseResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *MonitorTemplatesLibraryBaseResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *MonitorTemplatesLibraryBaseResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *MonitorTemplatesLibraryBaseResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetParentId

`func (o *MonitorTemplatesLibraryBaseResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *MonitorTemplatesLibraryBaseResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *MonitorTemplatesLibraryBaseResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.


### GetContentType

`func (o *MonitorTemplatesLibraryBaseResponse) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MonitorTemplatesLibraryBaseResponse) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *MonitorTemplatesLibraryBaseResponse) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetType

`func (o *MonitorTemplatesLibraryBaseResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorTemplatesLibraryBaseResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorTemplatesLibraryBaseResponse) SetType(v string)`

SetType sets Type field to given value.


### GetIsSystem

`func (o *MonitorTemplatesLibraryBaseResponse) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *MonitorTemplatesLibraryBaseResponse) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *MonitorTemplatesLibraryBaseResponse) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.


### GetIsMutable

`func (o *MonitorTemplatesLibraryBaseResponse) GetIsMutable() bool`

GetIsMutable returns the IsMutable field if non-nil, zero value otherwise.

### GetIsMutableOk

`func (o *MonitorTemplatesLibraryBaseResponse) GetIsMutableOk() (*bool, bool)`

GetIsMutableOk returns a tuple with the IsMutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMutable

`func (o *MonitorTemplatesLibraryBaseResponse) SetIsMutable(v bool)`

SetIsMutable sets IsMutable field to given value.


### GetPermissions

`func (o *MonitorTemplatesLibraryBaseResponse) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *MonitorTemplatesLibraryBaseResponse) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *MonitorTemplatesLibraryBaseResponse) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *MonitorTemplatesLibraryBaseResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


