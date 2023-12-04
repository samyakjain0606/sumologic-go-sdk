# UsageReportStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | An error code describing the type of error. | 
**Message** | **string** | A short English-language description of the error. | 
**Detail** | Pointer to **string** | An optional fuller English-language description of the error. | [optional] 
**Meta** | Pointer to **map[string]interface{}** | An optional list of metadata about the error. | [optional] 
**Status** | Pointer to **string** | Status export | [optional] 
**StatusMessage** | Pointer to **string** | Status message export | [optional] 
**ReportDownloadURL** | Pointer to **string** | S3 presigned download URL for the report. It is valid for 10 minutes. | [optional] 

## Methods

### NewUsageReportStatusResponse

`func NewUsageReportStatusResponse(code string, message string, ) *UsageReportStatusResponse`

NewUsageReportStatusResponse instantiates a new UsageReportStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageReportStatusResponseWithDefaults

`func NewUsageReportStatusResponseWithDefaults() *UsageReportStatusResponse`

NewUsageReportStatusResponseWithDefaults instantiates a new UsageReportStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UsageReportStatusResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UsageReportStatusResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UsageReportStatusResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *UsageReportStatusResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UsageReportStatusResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UsageReportStatusResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetail

`func (o *UsageReportStatusResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *UsageReportStatusResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *UsageReportStatusResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *UsageReportStatusResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetMeta

`func (o *UsageReportStatusResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UsageReportStatusResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UsageReportStatusResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UsageReportStatusResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetStatus

`func (o *UsageReportStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UsageReportStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UsageReportStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UsageReportStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *UsageReportStatusResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *UsageReportStatusResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *UsageReportStatusResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *UsageReportStatusResponse) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetReportDownloadURL

`func (o *UsageReportStatusResponse) GetReportDownloadURL() string`

GetReportDownloadURL returns the ReportDownloadURL field if non-nil, zero value otherwise.

### GetReportDownloadURLOk

`func (o *UsageReportStatusResponse) GetReportDownloadURLOk() (*string, bool)`

GetReportDownloadURLOk returns a tuple with the ReportDownloadURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDownloadURL

`func (o *UsageReportStatusResponse) SetReportDownloadURL(v string)`

SetReportDownloadURL sets ReportDownloadURL field to given value.

### HasReportDownloadURL

`func (o *UsageReportStatusResponse) HasReportDownloadURL() bool`

HasReportDownloadURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


