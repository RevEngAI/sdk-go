# ReportSubject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallSite** | Pointer to **int64** | Call-site address. | [optional] 
**Data** | Pointer to **int64** | Data address. | [optional] 
**Function** | Pointer to **int64** | Function entry address. | [optional] 
**Instruction** | Pointer to **int64** | Instruction address. | [optional] 

## Methods

### NewReportSubject

`func NewReportSubject() *ReportSubject`

NewReportSubject instantiates a new ReportSubject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportSubjectWithDefaults

`func NewReportSubjectWithDefaults() *ReportSubject`

NewReportSubjectWithDefaults instantiates a new ReportSubject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallSite

`func (o *ReportSubject) GetCallSite() int64`

GetCallSite returns the CallSite field if non-nil, zero value otherwise.

### GetCallSiteOk

`func (o *ReportSubject) GetCallSiteOk() (*int64, bool)`

GetCallSiteOk returns a tuple with the CallSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSite

`func (o *ReportSubject) SetCallSite(v int64)`

SetCallSite sets CallSite field to given value.

### HasCallSite

`func (o *ReportSubject) HasCallSite() bool`

HasCallSite returns a boolean if a field has been set.

### GetData

`func (o *ReportSubject) GetData() int64`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ReportSubject) GetDataOk() (*int64, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ReportSubject) SetData(v int64)`

SetData sets Data field to given value.

### HasData

`func (o *ReportSubject) HasData() bool`

HasData returns a boolean if a field has been set.

### GetFunction

`func (o *ReportSubject) GetFunction() int64`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *ReportSubject) GetFunctionOk() (*int64, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *ReportSubject) SetFunction(v int64)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *ReportSubject) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetInstruction

`func (o *ReportSubject) GetInstruction() int64`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *ReportSubject) GetInstructionOk() (*int64, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *ReportSubject) SetInstruction(v int64)`

SetInstruction sets Instruction field to given value.

### HasInstruction

`func (o *ReportSubject) HasInstruction() bool`

HasInstruction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


