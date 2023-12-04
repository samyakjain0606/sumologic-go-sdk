# OTCollector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the OT Collector. | 
**Name** | **string** | Name of the OT Collector. | 
**Version** | [**OTCollectorVersion**](OTCollectorVersion.md) |  | 
**Category** | Pointer to **string** | Category of the OT Collector. | [optional] 
**Description** | Pointer to **string** | Description of the OT Collector. | [optional] 
**Tags** | Pointer to **map[string]string** | Tags associated with the OT Collector. | [optional] 
**HealthIncidentsTracker** | Pointer to [**OTCollectorHealthIncidentsTracker**](OTCollectorHealthIncidentsTracker.md) |  | [optional] 
**Ephemeral** | Pointer to **bool** | Ephemeral Status of the OT Collector. | [optional] 
**Alive** | Pointer to **bool** | Alive Status of the OT Collector based on heartbeat. | [optional] 
**IsRemotelyManaged** | Pointer to **bool** | Management Status of the OT Collector based on if it is remotely or locally managed. | [optional] 
**EffectiveConfig** | Pointer to **map[string]string** | Config map that includes Base 64 Encoded Effective Configuration Yaml of the Remotely managed OT Collector. | [optional] 
**SystemInfo** | [**OTCollectorSystemInfo**](OTCollectorSystemInfo.md) |  | 
**TimeZone** | Pointer to **string** | timezone of the collector | [optional] 
**CreatedAt** | **time.Time** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**CreatedBy** | **string** | Identifier of the user who created the resource. | 
**ModifiedAt** | **time.Time** | Last modification timestamp in UTC. | 
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | 
**SourceTemplateLinkedCount** | Pointer to **int32** | Count of the source templates linked to a collector | [optional] 

## Methods

### NewOTCollector

`func NewOTCollector(id string, name string, version OTCollectorVersion, systemInfo OTCollectorSystemInfo, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, ) *OTCollector`

NewOTCollector instantiates a new OTCollector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOTCollectorWithDefaults

`func NewOTCollectorWithDefaults() *OTCollector`

NewOTCollectorWithDefaults instantiates a new OTCollector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OTCollector) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OTCollector) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OTCollector) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *OTCollector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OTCollector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OTCollector) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *OTCollector) GetVersion() OTCollectorVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OTCollector) GetVersionOk() (*OTCollectorVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OTCollector) SetVersion(v OTCollectorVersion)`

SetVersion sets Version field to given value.


### GetCategory

`func (o *OTCollector) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *OTCollector) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *OTCollector) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *OTCollector) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *OTCollector) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OTCollector) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OTCollector) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OTCollector) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *OTCollector) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OTCollector) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OTCollector) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OTCollector) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetHealthIncidentsTracker

`func (o *OTCollector) GetHealthIncidentsTracker() OTCollectorHealthIncidentsTracker`

GetHealthIncidentsTracker returns the HealthIncidentsTracker field if non-nil, zero value otherwise.

### GetHealthIncidentsTrackerOk

`func (o *OTCollector) GetHealthIncidentsTrackerOk() (*OTCollectorHealthIncidentsTracker, bool)`

GetHealthIncidentsTrackerOk returns a tuple with the HealthIncidentsTracker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthIncidentsTracker

`func (o *OTCollector) SetHealthIncidentsTracker(v OTCollectorHealthIncidentsTracker)`

SetHealthIncidentsTracker sets HealthIncidentsTracker field to given value.

### HasHealthIncidentsTracker

`func (o *OTCollector) HasHealthIncidentsTracker() bool`

HasHealthIncidentsTracker returns a boolean if a field has been set.

### GetEphemeral

`func (o *OTCollector) GetEphemeral() bool`

GetEphemeral returns the Ephemeral field if non-nil, zero value otherwise.

### GetEphemeralOk

`func (o *OTCollector) GetEphemeralOk() (*bool, bool)`

GetEphemeralOk returns a tuple with the Ephemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeral

`func (o *OTCollector) SetEphemeral(v bool)`

SetEphemeral sets Ephemeral field to given value.

### HasEphemeral

`func (o *OTCollector) HasEphemeral() bool`

HasEphemeral returns a boolean if a field has been set.

### GetAlive

`func (o *OTCollector) GetAlive() bool`

GetAlive returns the Alive field if non-nil, zero value otherwise.

### GetAliveOk

`func (o *OTCollector) GetAliveOk() (*bool, bool)`

GetAliveOk returns a tuple with the Alive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlive

`func (o *OTCollector) SetAlive(v bool)`

SetAlive sets Alive field to given value.

### HasAlive

`func (o *OTCollector) HasAlive() bool`

HasAlive returns a boolean if a field has been set.

### GetIsRemotelyManaged

`func (o *OTCollector) GetIsRemotelyManaged() bool`

GetIsRemotelyManaged returns the IsRemotelyManaged field if non-nil, zero value otherwise.

### GetIsRemotelyManagedOk

`func (o *OTCollector) GetIsRemotelyManagedOk() (*bool, bool)`

GetIsRemotelyManagedOk returns a tuple with the IsRemotelyManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemotelyManaged

`func (o *OTCollector) SetIsRemotelyManaged(v bool)`

SetIsRemotelyManaged sets IsRemotelyManaged field to given value.

### HasIsRemotelyManaged

`func (o *OTCollector) HasIsRemotelyManaged() bool`

HasIsRemotelyManaged returns a boolean if a field has been set.

### GetEffectiveConfig

`func (o *OTCollector) GetEffectiveConfig() map[string]string`

GetEffectiveConfig returns the EffectiveConfig field if non-nil, zero value otherwise.

### GetEffectiveConfigOk

`func (o *OTCollector) GetEffectiveConfigOk() (*map[string]string, bool)`

GetEffectiveConfigOk returns a tuple with the EffectiveConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveConfig

`func (o *OTCollector) SetEffectiveConfig(v map[string]string)`

SetEffectiveConfig sets EffectiveConfig field to given value.

### HasEffectiveConfig

`func (o *OTCollector) HasEffectiveConfig() bool`

HasEffectiveConfig returns a boolean if a field has been set.

### GetSystemInfo

`func (o *OTCollector) GetSystemInfo() OTCollectorSystemInfo`

GetSystemInfo returns the SystemInfo field if non-nil, zero value otherwise.

### GetSystemInfoOk

`func (o *OTCollector) GetSystemInfoOk() (*OTCollectorSystemInfo, bool)`

GetSystemInfoOk returns a tuple with the SystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemInfo

`func (o *OTCollector) SetSystemInfo(v OTCollectorSystemInfo)`

SetSystemInfo sets SystemInfo field to given value.


### GetTimeZone

`func (o *OTCollector) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *OTCollector) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *OTCollector) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *OTCollector) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OTCollector) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OTCollector) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OTCollector) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *OTCollector) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OTCollector) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OTCollector) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *OTCollector) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *OTCollector) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *OTCollector) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *OTCollector) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *OTCollector) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *OTCollector) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetSourceTemplateLinkedCount

`func (o *OTCollector) GetSourceTemplateLinkedCount() int32`

GetSourceTemplateLinkedCount returns the SourceTemplateLinkedCount field if non-nil, zero value otherwise.

### GetSourceTemplateLinkedCountOk

`func (o *OTCollector) GetSourceTemplateLinkedCountOk() (*int32, bool)`

GetSourceTemplateLinkedCountOk returns a tuple with the SourceTemplateLinkedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTemplateLinkedCount

`func (o *OTCollector) SetSourceTemplateLinkedCount(v int32)`

SetSourceTemplateLinkedCount sets SourceTemplateLinkedCount field to given value.

### HasSourceTemplateLinkedCount

`func (o *OTCollector) HasSourceTemplateLinkedCount() bool`

HasSourceTemplateLinkedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


