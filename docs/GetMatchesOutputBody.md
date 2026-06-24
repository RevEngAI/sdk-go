# GetMatchesOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Matches** | Pointer to [**[]FunctionMatch**](FunctionMatch.md) | Per-source-function matches. Populated when status&#x3D;COMPLETED; empty otherwise. | [optional] 
**Status** | **string** | Current workflow status | 

## Methods

### NewGetMatchesOutputBody

`func NewGetMatchesOutputBody(status string, ) *GetMatchesOutputBody`

NewGetMatchesOutputBody instantiates a new GetMatchesOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMatchesOutputBodyWithDefaults

`func NewGetMatchesOutputBodyWithDefaults() *GetMatchesOutputBody`

NewGetMatchesOutputBodyWithDefaults instantiates a new GetMatchesOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatches

`func (o *GetMatchesOutputBody) GetMatches() []FunctionMatch`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *GetMatchesOutputBody) GetMatchesOk() (*[]FunctionMatch, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *GetMatchesOutputBody) SetMatches(v []FunctionMatch)`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *GetMatchesOutputBody) HasMatches() bool`

HasMatches returns a boolean if a field has been set.

### SetMatchesNil

`func (o *GetMatchesOutputBody) SetMatchesNil(b bool)`

 SetMatchesNil sets the value for Matches to be an explicit nil

### UnsetMatches
`func (o *GetMatchesOutputBody) UnsetMatches()`

UnsetMatches ensures that no value is present for Matches, not even an explicit nil
### GetStatus

`func (o *GetMatchesOutputBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMatchesOutputBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMatchesOutputBody) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


