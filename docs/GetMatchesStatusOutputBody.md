# GetMatchesStatusOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]ProgressMessage**](ProgressMessage.md) | Log messages emitted during execution | 
**Status** | **string** | Current workflow status | 
**Step** | **string** | Name of the current step | 
**StepIndex** | **int64** | Zero-based index of the current step | 
**StepsTotal** | **int64** | Total number of steps in the workflow | 

## Methods

### NewGetMatchesStatusOutputBody

`func NewGetMatchesStatusOutputBody(messages []ProgressMessage, status string, step string, stepIndex int64, stepsTotal int64, ) *GetMatchesStatusOutputBody`

NewGetMatchesStatusOutputBody instantiates a new GetMatchesStatusOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMatchesStatusOutputBodyWithDefaults

`func NewGetMatchesStatusOutputBodyWithDefaults() *GetMatchesStatusOutputBody`

NewGetMatchesStatusOutputBodyWithDefaults instantiates a new GetMatchesStatusOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *GetMatchesStatusOutputBody) GetMessages() []ProgressMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *GetMatchesStatusOutputBody) GetMessagesOk() (*[]ProgressMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *GetMatchesStatusOutputBody) SetMessages(v []ProgressMessage)`

SetMessages sets Messages field to given value.


### SetMessagesNil

`func (o *GetMatchesStatusOutputBody) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *GetMatchesStatusOutputBody) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetStatus

`func (o *GetMatchesStatusOutputBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMatchesStatusOutputBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMatchesStatusOutputBody) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStep

`func (o *GetMatchesStatusOutputBody) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *GetMatchesStatusOutputBody) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *GetMatchesStatusOutputBody) SetStep(v string)`

SetStep sets Step field to given value.


### GetStepIndex

`func (o *GetMatchesStatusOutputBody) GetStepIndex() int64`

GetStepIndex returns the StepIndex field if non-nil, zero value otherwise.

### GetStepIndexOk

`func (o *GetMatchesStatusOutputBody) GetStepIndexOk() (*int64, bool)`

GetStepIndexOk returns a tuple with the StepIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepIndex

`func (o *GetMatchesStatusOutputBody) SetStepIndex(v int64)`

SetStepIndex sets StepIndex field to given value.


### GetStepsTotal

`func (o *GetMatchesStatusOutputBody) GetStepsTotal() int64`

GetStepsTotal returns the StepsTotal field if non-nil, zero value otherwise.

### GetStepsTotalOk

`func (o *GetMatchesStatusOutputBody) GetStepsTotalOk() (*int64, bool)`

GetStepsTotalOk returns a tuple with the StepsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsTotal

`func (o *GetMatchesStatusOutputBody) SetStepsTotal(v int64)`

SetStepsTotal sets StepsTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


