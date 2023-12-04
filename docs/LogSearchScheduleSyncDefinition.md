# LogSearchScheduleSyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpression** | Pointer to **string** | Cron-like expression specifying the search&#39;s schedule. Field scheduleType must be set to \&quot;Custom\&quot;, otherwise, scheduleType takes precedence over cronExpression. | [optional] 
**DisplayableTimeRange** | Pointer to **string** | A human-friendly text describing the query time range. For e.g. \&quot;-2h\&quot;, \&quot;last three days\&quot;, \&quot;team default time\&quot;. This value can not be set via API. | [optional] 
**ParseableTimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 
**TimeZone** | **string** | Time zone identifier for time specification. Either an abbreviation such as \&quot;PST\&quot;, a full name such as \&quot;America/Los_Angeles\&quot;, or a custom ID such as \&quot;GMT-8:00\&quot;. Note that the support of abbreviations is for JDK 1.1.x compatibility only and full names should be used. The GMT time zone is chosen if the given time zone cannot be identified. | 
**Threshold** | Pointer to [**LogSearchNotificationThresholdSyncDefinition**](LogSearchNotificationThresholdSyncDefinition.md) |  | [optional] 
**Notification** | [**ScheduleNotificationSyncDefinition**](ScheduleNotificationSyncDefinition.md) |  | 
**ScheduleType** | **string** | Run schedule of the scheduled search. Set to \&quot;Custom\&quot; to specify the schedule with a CRON expression.Please note that with Custom, 1Day and 1Week schedule types you need to provide the corresponding cron expression to determine when to actually run the search. e.g. Sample Valid Cron for 1Day is \&quot;0 0 16 ? * 2-6 *\&quot;. Possible schedule types are:   - &#x60;RealTime&#x60;   - &#x60;15Minutes&#x60;   - &#x60;1Hour&#x60;   - &#x60;2Hours&#x60;   - &#x60;4Hours&#x60;   - &#x60;6Hours&#x60;   - &#x60;8Hours&#x60;   - &#x60;12Hours&#x60;   - &#x60;1Day&#x60;   - &#x60;1Week&#x60;   - &#x60;Custom&#x60; | 
**MuteErrorEmails** | Pointer to **bool** | If enabled, emails are not sent out in case of errors with the search. | [optional] 
**Parameters** | Pointer to [**[]ScheduleSearchParameterSyncDefinition**](ScheduleSearchParameterSyncDefinition.md) | A list of scheduled search template parameters to be used while executing the query. This is different from the queryParameters field in parent object as this field will be  used for execution as  per the schedule. The parent object field is for search itself, not part of execution.  Learn more about the search templates here :  https://help.sumologic.com/docs/search/get-started-with-search/build-search/search-templates/ | [optional] 

## Methods

### NewLogSearchScheduleSyncDefinition

`func NewLogSearchScheduleSyncDefinition(parseableTimeRange ResolvableTimeRange, timeZone string, notification ScheduleNotificationSyncDefinition, scheduleType string, ) *LogSearchScheduleSyncDefinition`

NewLogSearchScheduleSyncDefinition instantiates a new LogSearchScheduleSyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogSearchScheduleSyncDefinitionWithDefaults

`func NewLogSearchScheduleSyncDefinitionWithDefaults() *LogSearchScheduleSyncDefinition`

NewLogSearchScheduleSyncDefinitionWithDefaults instantiates a new LogSearchScheduleSyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpression

`func (o *LogSearchScheduleSyncDefinition) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *LogSearchScheduleSyncDefinition) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *LogSearchScheduleSyncDefinition) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *LogSearchScheduleSyncDefinition) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetDisplayableTimeRange

`func (o *LogSearchScheduleSyncDefinition) GetDisplayableTimeRange() string`

GetDisplayableTimeRange returns the DisplayableTimeRange field if non-nil, zero value otherwise.

### GetDisplayableTimeRangeOk

`func (o *LogSearchScheduleSyncDefinition) GetDisplayableTimeRangeOk() (*string, bool)`

GetDisplayableTimeRangeOk returns a tuple with the DisplayableTimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayableTimeRange

`func (o *LogSearchScheduleSyncDefinition) SetDisplayableTimeRange(v string)`

SetDisplayableTimeRange sets DisplayableTimeRange field to given value.

### HasDisplayableTimeRange

`func (o *LogSearchScheduleSyncDefinition) HasDisplayableTimeRange() bool`

HasDisplayableTimeRange returns a boolean if a field has been set.

### GetParseableTimeRange

`func (o *LogSearchScheduleSyncDefinition) GetParseableTimeRange() ResolvableTimeRange`

GetParseableTimeRange returns the ParseableTimeRange field if non-nil, zero value otherwise.

### GetParseableTimeRangeOk

`func (o *LogSearchScheduleSyncDefinition) GetParseableTimeRangeOk() (*ResolvableTimeRange, bool)`

GetParseableTimeRangeOk returns a tuple with the ParseableTimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseableTimeRange

`func (o *LogSearchScheduleSyncDefinition) SetParseableTimeRange(v ResolvableTimeRange)`

SetParseableTimeRange sets ParseableTimeRange field to given value.


### GetTimeZone

`func (o *LogSearchScheduleSyncDefinition) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *LogSearchScheduleSyncDefinition) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *LogSearchScheduleSyncDefinition) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetThreshold

`func (o *LogSearchScheduleSyncDefinition) GetThreshold() LogSearchNotificationThresholdSyncDefinition`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *LogSearchScheduleSyncDefinition) GetThresholdOk() (*LogSearchNotificationThresholdSyncDefinition, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *LogSearchScheduleSyncDefinition) SetThreshold(v LogSearchNotificationThresholdSyncDefinition)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *LogSearchScheduleSyncDefinition) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetNotification

`func (o *LogSearchScheduleSyncDefinition) GetNotification() ScheduleNotificationSyncDefinition`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *LogSearchScheduleSyncDefinition) GetNotificationOk() (*ScheduleNotificationSyncDefinition, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *LogSearchScheduleSyncDefinition) SetNotification(v ScheduleNotificationSyncDefinition)`

SetNotification sets Notification field to given value.


### GetScheduleType

`func (o *LogSearchScheduleSyncDefinition) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *LogSearchScheduleSyncDefinition) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *LogSearchScheduleSyncDefinition) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.


### GetMuteErrorEmails

`func (o *LogSearchScheduleSyncDefinition) GetMuteErrorEmails() bool`

GetMuteErrorEmails returns the MuteErrorEmails field if non-nil, zero value otherwise.

### GetMuteErrorEmailsOk

`func (o *LogSearchScheduleSyncDefinition) GetMuteErrorEmailsOk() (*bool, bool)`

GetMuteErrorEmailsOk returns a tuple with the MuteErrorEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteErrorEmails

`func (o *LogSearchScheduleSyncDefinition) SetMuteErrorEmails(v bool)`

SetMuteErrorEmails sets MuteErrorEmails field to given value.

### HasMuteErrorEmails

`func (o *LogSearchScheduleSyncDefinition) HasMuteErrorEmails() bool`

HasMuteErrorEmails returns a boolean if a field has been set.

### GetParameters

`func (o *LogSearchScheduleSyncDefinition) GetParameters() []ScheduleSearchParameterSyncDefinition`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *LogSearchScheduleSyncDefinition) GetParametersOk() (*[]ScheduleSearchParameterSyncDefinition, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *LogSearchScheduleSyncDefinition) SetParameters(v []ScheduleSearchParameterSyncDefinition)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *LogSearchScheduleSyncDefinition) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


