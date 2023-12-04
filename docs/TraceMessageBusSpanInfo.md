# TraceMessageBusSpanInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** | An address at which messages can be exchanged e.g. a Kafka record has an associated \&quot;topic name\&quot; that can be stored using this tag. | [optional] 

## Methods

### NewTraceMessageBusSpanInfo

`func NewTraceMessageBusSpanInfo() *TraceMessageBusSpanInfo`

NewTraceMessageBusSpanInfo instantiates a new TraceMessageBusSpanInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceMessageBusSpanInfoWithDefaults

`func NewTraceMessageBusSpanInfoWithDefaults() *TraceMessageBusSpanInfo`

NewTraceMessageBusSpanInfoWithDefaults instantiates a new TraceMessageBusSpanInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *TraceMessageBusSpanInfo) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *TraceMessageBusSpanInfo) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *TraceMessageBusSpanInfo) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *TraceMessageBusSpanInfo) HasDestination() bool`

HasDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


