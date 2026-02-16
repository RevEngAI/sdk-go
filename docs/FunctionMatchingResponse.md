# FunctionMatchingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Progress** | Pointer to **int32** | Progress of the matching operation, represented as a percentage | [optional] [default to 0]
**Status** | Pointer to **NullableString** |  | [optional] 
**TotalTime** | Pointer to **NullableInt32** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**CurrentPage** | Pointer to **NullableInt32** |  | [optional] 
**TotalPages** | Pointer to **NullableInt32** |  | [optional] 
**Matches** | Pointer to [**[]FunctionMatch**](FunctionMatch.md) |  | [optional] 
**NumMatches** | Pointer to **NullableInt32** |  | [optional] 
**NumDebugMatches** | Pointer to **NullableInt32** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFunctionMatchingResponse

`func NewFunctionMatchingResponse() *FunctionMatchingResponse`

NewFunctionMatchingResponse instantiates a new FunctionMatchingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionMatchingResponseWithDefaults

`func NewFunctionMatchingResponseWithDefaults() *FunctionMatchingResponse`

NewFunctionMatchingResponseWithDefaults instantiates a new FunctionMatchingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgress

`func (o *FunctionMatchingResponse) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *FunctionMatchingResponse) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *FunctionMatchingResponse) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *FunctionMatchingResponse) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetStatus

`func (o *FunctionMatchingResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FunctionMatchingResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FunctionMatchingResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FunctionMatchingResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *FunctionMatchingResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *FunctionMatchingResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTotalTime

`func (o *FunctionMatchingResponse) GetTotalTime() int32`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *FunctionMatchingResponse) GetTotalTimeOk() (*int32, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *FunctionMatchingResponse) SetTotalTime(v int32)`

SetTotalTime sets TotalTime field to given value.

### HasTotalTime

`func (o *FunctionMatchingResponse) HasTotalTime() bool`

HasTotalTime returns a boolean if a field has been set.

### SetTotalTimeNil

`func (o *FunctionMatchingResponse) SetTotalTimeNil(b bool)`

 SetTotalTimeNil sets the value for TotalTime to be an explicit nil

### UnsetTotalTime
`func (o *FunctionMatchingResponse) UnsetTotalTime()`

UnsetTotalTime ensures that no value is present for TotalTime, not even an explicit nil
### GetErrorMessage

`func (o *FunctionMatchingResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *FunctionMatchingResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *FunctionMatchingResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *FunctionMatchingResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *FunctionMatchingResponse) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *FunctionMatchingResponse) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetCurrentPage

`func (o *FunctionMatchingResponse) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *FunctionMatchingResponse) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *FunctionMatchingResponse) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *FunctionMatchingResponse) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### SetCurrentPageNil

`func (o *FunctionMatchingResponse) SetCurrentPageNil(b bool)`

 SetCurrentPageNil sets the value for CurrentPage to be an explicit nil

### UnsetCurrentPage
`func (o *FunctionMatchingResponse) UnsetCurrentPage()`

UnsetCurrentPage ensures that no value is present for CurrentPage, not even an explicit nil
### GetTotalPages

`func (o *FunctionMatchingResponse) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *FunctionMatchingResponse) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *FunctionMatchingResponse) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *FunctionMatchingResponse) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### SetTotalPagesNil

`func (o *FunctionMatchingResponse) SetTotalPagesNil(b bool)`

 SetTotalPagesNil sets the value for TotalPages to be an explicit nil

### UnsetTotalPages
`func (o *FunctionMatchingResponse) UnsetTotalPages()`

UnsetTotalPages ensures that no value is present for TotalPages, not even an explicit nil
### GetMatches

`func (o *FunctionMatchingResponse) GetMatches() []FunctionMatch`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *FunctionMatchingResponse) GetMatchesOk() (*[]FunctionMatch, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *FunctionMatchingResponse) SetMatches(v []FunctionMatch)`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *FunctionMatchingResponse) HasMatches() bool`

HasMatches returns a boolean if a field has been set.

### SetMatchesNil

`func (o *FunctionMatchingResponse) SetMatchesNil(b bool)`

 SetMatchesNil sets the value for Matches to be an explicit nil

### UnsetMatches
`func (o *FunctionMatchingResponse) UnsetMatches()`

UnsetMatches ensures that no value is present for Matches, not even an explicit nil
### GetNumMatches

`func (o *FunctionMatchingResponse) GetNumMatches() int32`

GetNumMatches returns the NumMatches field if non-nil, zero value otherwise.

### GetNumMatchesOk

`func (o *FunctionMatchingResponse) GetNumMatchesOk() (*int32, bool)`

GetNumMatchesOk returns a tuple with the NumMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMatches

`func (o *FunctionMatchingResponse) SetNumMatches(v int32)`

SetNumMatches sets NumMatches field to given value.

### HasNumMatches

`func (o *FunctionMatchingResponse) HasNumMatches() bool`

HasNumMatches returns a boolean if a field has been set.

### SetNumMatchesNil

`func (o *FunctionMatchingResponse) SetNumMatchesNil(b bool)`

 SetNumMatchesNil sets the value for NumMatches to be an explicit nil

### UnsetNumMatches
`func (o *FunctionMatchingResponse) UnsetNumMatches()`

UnsetNumMatches ensures that no value is present for NumMatches, not even an explicit nil
### GetNumDebugMatches

`func (o *FunctionMatchingResponse) GetNumDebugMatches() int32`

GetNumDebugMatches returns the NumDebugMatches field if non-nil, zero value otherwise.

### GetNumDebugMatchesOk

`func (o *FunctionMatchingResponse) GetNumDebugMatchesOk() (*int32, bool)`

GetNumDebugMatchesOk returns a tuple with the NumDebugMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDebugMatches

`func (o *FunctionMatchingResponse) SetNumDebugMatches(v int32)`

SetNumDebugMatches sets NumDebugMatches field to given value.

### HasNumDebugMatches

`func (o *FunctionMatchingResponse) HasNumDebugMatches() bool`

HasNumDebugMatches returns a boolean if a field has been set.

### SetNumDebugMatchesNil

`func (o *FunctionMatchingResponse) SetNumDebugMatchesNil(b bool)`

 SetNumDebugMatchesNil sets the value for NumDebugMatches to be an explicit nil

### UnsetNumDebugMatches
`func (o *FunctionMatchingResponse) UnsetNumDebugMatches()`

UnsetNumDebugMatches ensures that no value is present for NumDebugMatches, not even an explicit nil
### GetUpdatedAt

`func (o *FunctionMatchingResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FunctionMatchingResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FunctionMatchingResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FunctionMatchingResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *FunctionMatchingResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *FunctionMatchingResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


