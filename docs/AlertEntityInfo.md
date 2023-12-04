# AlertEntityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **string** | Identifier of the entity. | [optional] 
**EntityName** | Pointer to **string** | Name of the entity. | [optional] 
**EntityTypeId** | Pointer to **string** | Entity type ID or empty if unknown. | [optional] [default to ""]
**IsPrimaryWithinDomain** | Pointer to **bool** | Whether entity is the most specific entity within its domain for that alert. | [optional] [default to true]
**IsPrimaryDomain** | Pointer to **bool** | Whether entity is from the most accurate domain found for this alert. | [optional] [default to true]

## Methods

### NewAlertEntityInfo

`func NewAlertEntityInfo() *AlertEntityInfo`

NewAlertEntityInfo instantiates a new AlertEntityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertEntityInfoWithDefaults

`func NewAlertEntityInfoWithDefaults() *AlertEntityInfo`

NewAlertEntityInfoWithDefaults instantiates a new AlertEntityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *AlertEntityInfo) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AlertEntityInfo) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AlertEntityInfo) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AlertEntityInfo) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityName

`func (o *AlertEntityInfo) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *AlertEntityInfo) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *AlertEntityInfo) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *AlertEntityInfo) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetEntityTypeId

`func (o *AlertEntityInfo) GetEntityTypeId() string`

GetEntityTypeId returns the EntityTypeId field if non-nil, zero value otherwise.

### GetEntityTypeIdOk

`func (o *AlertEntityInfo) GetEntityTypeIdOk() (*string, bool)`

GetEntityTypeIdOk returns a tuple with the EntityTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypeId

`func (o *AlertEntityInfo) SetEntityTypeId(v string)`

SetEntityTypeId sets EntityTypeId field to given value.

### HasEntityTypeId

`func (o *AlertEntityInfo) HasEntityTypeId() bool`

HasEntityTypeId returns a boolean if a field has been set.

### GetIsPrimaryWithinDomain

`func (o *AlertEntityInfo) GetIsPrimaryWithinDomain() bool`

GetIsPrimaryWithinDomain returns the IsPrimaryWithinDomain field if non-nil, zero value otherwise.

### GetIsPrimaryWithinDomainOk

`func (o *AlertEntityInfo) GetIsPrimaryWithinDomainOk() (*bool, bool)`

GetIsPrimaryWithinDomainOk returns a tuple with the IsPrimaryWithinDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimaryWithinDomain

`func (o *AlertEntityInfo) SetIsPrimaryWithinDomain(v bool)`

SetIsPrimaryWithinDomain sets IsPrimaryWithinDomain field to given value.

### HasIsPrimaryWithinDomain

`func (o *AlertEntityInfo) HasIsPrimaryWithinDomain() bool`

HasIsPrimaryWithinDomain returns a boolean if a field has been set.

### GetIsPrimaryDomain

`func (o *AlertEntityInfo) GetIsPrimaryDomain() bool`

GetIsPrimaryDomain returns the IsPrimaryDomain field if non-nil, zero value otherwise.

### GetIsPrimaryDomainOk

`func (o *AlertEntityInfo) GetIsPrimaryDomainOk() (*bool, bool)`

GetIsPrimaryDomainOk returns a tuple with the IsPrimaryDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimaryDomain

`func (o *AlertEntityInfo) SetIsPrimaryDomain(v bool)`

SetIsPrimaryDomain sets IsPrimaryDomain field to given value.

### HasIsPrimaryDomain

`func (o *AlertEntityInfo) HasIsPrimaryDomain() bool`

HasIsPrimaryDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


