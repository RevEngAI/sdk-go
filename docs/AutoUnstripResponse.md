# AutoUnstripResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Progress** | Pointer to **int32** | Progress of the auto-unstrip operation, represented as a percentage | [optional] [default to 0]
**Status** | Pointer to **NullableString** |  | [optional] 
**TotalTime** | Pointer to **NullableInt32** |  | [optional] 
**Matches** | Pointer to [**[]MatchedFunctionSuggestion**](MatchedFunctionSuggestion.md) |  | [optional] 
**Applied** | Pointer to **NullableBool** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAutoUnstripResponse

`func NewAutoUnstripResponse() *AutoUnstripResponse`

NewAutoUnstripResponse instantiates a new AutoUnstripResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoUnstripResponseWithDefaults

`func NewAutoUnstripResponseWithDefaults() *AutoUnstripResponse`

NewAutoUnstripResponseWithDefaults instantiates a new AutoUnstripResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgress

`func (o *AutoUnstripResponse) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *AutoUnstripResponse) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *AutoUnstripResponse) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *AutoUnstripResponse) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetStatus

`func (o *AutoUnstripResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutoUnstripResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutoUnstripResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AutoUnstripResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AutoUnstripResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AutoUnstripResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTotalTime

`func (o *AutoUnstripResponse) GetTotalTime() int32`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *AutoUnstripResponse) GetTotalTimeOk() (*int32, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *AutoUnstripResponse) SetTotalTime(v int32)`

SetTotalTime sets TotalTime field to given value.

### HasTotalTime

`func (o *AutoUnstripResponse) HasTotalTime() bool`

HasTotalTime returns a boolean if a field has been set.

### SetTotalTimeNil

`func (o *AutoUnstripResponse) SetTotalTimeNil(b bool)`

 SetTotalTimeNil sets the value for TotalTime to be an explicit nil

### UnsetTotalTime
`func (o *AutoUnstripResponse) UnsetTotalTime()`

UnsetTotalTime ensures that no value is present for TotalTime, not even an explicit nil
### GetMatches

`func (o *AutoUnstripResponse) GetMatches() []MatchedFunctionSuggestion`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *AutoUnstripResponse) GetMatchesOk() (*[]MatchedFunctionSuggestion, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *AutoUnstripResponse) SetMatches(v []MatchedFunctionSuggestion)`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *AutoUnstripResponse) HasMatches() bool`

HasMatches returns a boolean if a field has been set.

### SetMatchesNil

`func (o *AutoUnstripResponse) SetMatchesNil(b bool)`

 SetMatchesNil sets the value for Matches to be an explicit nil

### UnsetMatches
`func (o *AutoUnstripResponse) UnsetMatches()`

UnsetMatches ensures that no value is present for Matches, not even an explicit nil
### GetApplied

`func (o *AutoUnstripResponse) GetApplied() bool`

GetApplied returns the Applied field if non-nil, zero value otherwise.

### GetAppliedOk

`func (o *AutoUnstripResponse) GetAppliedOk() (*bool, bool)`

GetAppliedOk returns a tuple with the Applied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplied

`func (o *AutoUnstripResponse) SetApplied(v bool)`

SetApplied sets Applied field to given value.

### HasApplied

`func (o *AutoUnstripResponse) HasApplied() bool`

HasApplied returns a boolean if a field has been set.

### SetAppliedNil

`func (o *AutoUnstripResponse) SetAppliedNil(b bool)`

 SetAppliedNil sets the value for Applied to be an explicit nil

### UnsetApplied
`func (o *AutoUnstripResponse) UnsetApplied()`

UnsetApplied ensures that no value is present for Applied, not even an explicit nil
### GetErrorMessage

`func (o *AutoUnstripResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AutoUnstripResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AutoUnstripResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AutoUnstripResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *AutoUnstripResponse) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *AutoUnstripResponse) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


