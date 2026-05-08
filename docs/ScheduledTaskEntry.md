# ScheduledTaskEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** |  | [optional] 
**Day** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**Events** | Pointer to [**[]ReportEvent**](ReportEvent.md) |  | [optional] 
**Executable** | Pointer to **string** |  | [optional] 
**Modifier** | Pointer to **string** |  | [optional] 
**RunAs** | Pointer to **string** |  | [optional] 
**ScheduleType** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**TaskName** | Pointer to **string** |  | [optional] 

## Methods

### NewScheduledTaskEntry

`func NewScheduledTaskEntry() *ScheduledTaskEntry`

NewScheduledTaskEntry instantiates a new ScheduledTaskEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledTaskEntryWithDefaults

`func NewScheduledTaskEntryWithDefaults() *ScheduledTaskEntry`

NewScheduledTaskEntryWithDefaults instantiates a new ScheduledTaskEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *ScheduledTaskEntry) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ScheduledTaskEntry) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ScheduledTaskEntry) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *ScheduledTaskEntry) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetDay

`func (o *ScheduledTaskEntry) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *ScheduledTaskEntry) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *ScheduledTaskEntry) SetDay(v string)`

SetDay sets Day field to given value.

### HasDay

`func (o *ScheduledTaskEntry) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetEndDate

`func (o *ScheduledTaskEntry) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ScheduledTaskEntry) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ScheduledTaskEntry) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ScheduledTaskEntry) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEvents

`func (o *ScheduledTaskEntry) GetEvents() []ReportEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ScheduledTaskEntry) GetEventsOk() (*[]ReportEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ScheduledTaskEntry) SetEvents(v []ReportEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ScheduledTaskEntry) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *ScheduledTaskEntry) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *ScheduledTaskEntry) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetExecutable

`func (o *ScheduledTaskEntry) GetExecutable() string`

GetExecutable returns the Executable field if non-nil, zero value otherwise.

### GetExecutableOk

`func (o *ScheduledTaskEntry) GetExecutableOk() (*string, bool)`

GetExecutableOk returns a tuple with the Executable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutable

`func (o *ScheduledTaskEntry) SetExecutable(v string)`

SetExecutable sets Executable field to given value.

### HasExecutable

`func (o *ScheduledTaskEntry) HasExecutable() bool`

HasExecutable returns a boolean if a field has been set.

### GetModifier

`func (o *ScheduledTaskEntry) GetModifier() string`

GetModifier returns the Modifier field if non-nil, zero value otherwise.

### GetModifierOk

`func (o *ScheduledTaskEntry) GetModifierOk() (*string, bool)`

GetModifierOk returns a tuple with the Modifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifier

`func (o *ScheduledTaskEntry) SetModifier(v string)`

SetModifier sets Modifier field to given value.

### HasModifier

`func (o *ScheduledTaskEntry) HasModifier() bool`

HasModifier returns a boolean if a field has been set.

### GetRunAs

`func (o *ScheduledTaskEntry) GetRunAs() string`

GetRunAs returns the RunAs field if non-nil, zero value otherwise.

### GetRunAsOk

`func (o *ScheduledTaskEntry) GetRunAsOk() (*string, bool)`

GetRunAsOk returns a tuple with the RunAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAs

`func (o *ScheduledTaskEntry) SetRunAs(v string)`

SetRunAs sets RunAs field to given value.

### HasRunAs

`func (o *ScheduledTaskEntry) HasRunAs() bool`

HasRunAs returns a boolean if a field has been set.

### GetScheduleType

`func (o *ScheduledTaskEntry) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *ScheduledTaskEntry) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *ScheduledTaskEntry) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *ScheduledTaskEntry) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetStartDate

`func (o *ScheduledTaskEntry) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ScheduledTaskEntry) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ScheduledTaskEntry) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ScheduledTaskEntry) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartTime

`func (o *ScheduledTaskEntry) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ScheduledTaskEntry) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ScheduledTaskEntry) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ScheduledTaskEntry) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTaskName

`func (o *ScheduledTaskEntry) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *ScheduledTaskEntry) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *ScheduledTaskEntry) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *ScheduledTaskEntry) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


