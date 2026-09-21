# Subject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Function** | **int32** | Function entry address. | 
**CallSite** | **int32** |  | 
**Instruction** | **int32** |  | 
**Data** | **int32** |  | 

## Methods

### NewSubject

`func NewSubject(function int32, callSite int32, instruction int32, data int32, ) *Subject`

NewSubject instantiates a new Subject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectWithDefaults

`func NewSubjectWithDefaults() *Subject`

NewSubjectWithDefaults instantiates a new Subject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunction

`func (o *Subject) GetFunction() int32`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *Subject) GetFunctionOk() (*int32, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *Subject) SetFunction(v int32)`

SetFunction sets Function field to given value.


### GetCallSite

`func (o *Subject) GetCallSite() int32`

GetCallSite returns the CallSite field if non-nil, zero value otherwise.

### GetCallSiteOk

`func (o *Subject) GetCallSiteOk() (*int32, bool)`

GetCallSiteOk returns a tuple with the CallSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSite

`func (o *Subject) SetCallSite(v int32)`

SetCallSite sets CallSite field to given value.


### GetInstruction

`func (o *Subject) GetInstruction() int32`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *Subject) GetInstructionOk() (*int32, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *Subject) SetInstruction(v int32)`

SetInstruction sets Instruction field to given value.


### GetData

`func (o *Subject) GetData() int32`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Subject) GetDataOk() (*int32, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Subject) SetData(v int32)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


