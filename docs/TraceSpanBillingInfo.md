# TraceSpanBillingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BilledBytes** | **int32** | Number of bytes that were charged for the span. | 
**BilledFormat** | **string** | Billing format of the span. Number of bytes of this representation of the span is equal to &#x60;billedBytes&#x60;. | 

## Methods

### NewTraceSpanBillingInfo

`func NewTraceSpanBillingInfo(billedBytes int32, billedFormat string, ) *TraceSpanBillingInfo`

NewTraceSpanBillingInfo instantiates a new TraceSpanBillingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceSpanBillingInfoWithDefaults

`func NewTraceSpanBillingInfoWithDefaults() *TraceSpanBillingInfo`

NewTraceSpanBillingInfoWithDefaults instantiates a new TraceSpanBillingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBilledBytes

`func (o *TraceSpanBillingInfo) GetBilledBytes() int32`

GetBilledBytes returns the BilledBytes field if non-nil, zero value otherwise.

### GetBilledBytesOk

`func (o *TraceSpanBillingInfo) GetBilledBytesOk() (*int32, bool)`

GetBilledBytesOk returns a tuple with the BilledBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledBytes

`func (o *TraceSpanBillingInfo) SetBilledBytes(v int32)`

SetBilledBytes sets BilledBytes field to given value.


### GetBilledFormat

`func (o *TraceSpanBillingInfo) GetBilledFormat() string`

GetBilledFormat returns the BilledFormat field if non-nil, zero value otherwise.

### GetBilledFormatOk

`func (o *TraceSpanBillingInfo) GetBilledFormatOk() (*string, bool)`

GetBilledFormatOk returns a tuple with the BilledFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledFormat

`func (o *TraceSpanBillingInfo) SetBilledFormat(v string)`

SetBilledFormat sets BilledFormat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


