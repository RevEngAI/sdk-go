# StartMatchingOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchId** | **string** | Opaque token for this matching run. Pass it to the GET/status endpoints&#39; match_id query parameter to fetch this exact run. | 
**Messages** | [**[]ProgressMessage**](ProgressMessage.md) | Log messages emitted during execution | 
**Status** | **string** | Current workflow status | 
**Step** | **string** | Name of the current step | 
**StepIndex** | **int64** | Zero-based index of the current step | 
**StepsTotal** | **int64** | Total number of steps in the workflow | 

## Methods

### NewStartMatchingOutputBody

`func NewStartMatchingOutputBody(matchId string, messages []ProgressMessage, status string, step string, stepIndex int64, stepsTotal int64, ) *StartMatchingOutputBody`

NewStartMatchingOutputBody instantiates a new StartMatchingOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartMatchingOutputBodyWithDefaults

`func NewStartMatchingOutputBodyWithDefaults() *StartMatchingOutputBody`

NewStartMatchingOutputBodyWithDefaults instantiates a new StartMatchingOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchId

`func (o *StartMatchingOutputBody) GetMatchId() string`

GetMatchId returns the MatchId field if non-nil, zero value otherwise.

### GetMatchIdOk

`func (o *StartMatchingOutputBody) GetMatchIdOk() (*string, bool)`

GetMatchIdOk returns a tuple with the MatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchId

`func (o *StartMatchingOutputBody) SetMatchId(v string)`

SetMatchId sets MatchId field to given value.


### GetMessages

`func (o *StartMatchingOutputBody) GetMessages() []ProgressMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *StartMatchingOutputBody) GetMessagesOk() (*[]ProgressMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *StartMatchingOutputBody) SetMessages(v []ProgressMessage)`

SetMessages sets Messages field to given value.


### SetMessagesNil

`func (o *StartMatchingOutputBody) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *StartMatchingOutputBody) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetStatus

`func (o *StartMatchingOutputBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StartMatchingOutputBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StartMatchingOutputBody) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStep

`func (o *StartMatchingOutputBody) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *StartMatchingOutputBody) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *StartMatchingOutputBody) SetStep(v string)`

SetStep sets Step field to given value.


### GetStepIndex

`func (o *StartMatchingOutputBody) GetStepIndex() int64`

GetStepIndex returns the StepIndex field if non-nil, zero value otherwise.

### GetStepIndexOk

`func (o *StartMatchingOutputBody) GetStepIndexOk() (*int64, bool)`

GetStepIndexOk returns a tuple with the StepIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepIndex

`func (o *StartMatchingOutputBody) SetStepIndex(v int64)`

SetStepIndex sets StepIndex field to given value.


### GetStepsTotal

`func (o *StartMatchingOutputBody) GetStepsTotal() int64`

GetStepsTotal returns the StepsTotal field if non-nil, zero value otherwise.

### GetStepsTotalOk

`func (o *StartMatchingOutputBody) GetStepsTotalOk() (*int64, bool)`

GetStepsTotalOk returns a tuple with the StepsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsTotal

`func (o *StartMatchingOutputBody) SetStepsTotal(v int64)`

SetStepsTotal sets StepsTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


