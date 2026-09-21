# FilesystemAnalyseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cancelled** | **bool** | Whether the run was cancelled | 
**DataDescription** | Pointer to **string** | Description of the data read or written | [optional] 
**DataDestination** | Pointer to **string** | Where the data ends up | [optional] 
**DataOrigin** | Pointer to **string** | Where the data originates from | [optional] 
**FunctionId** | **int64** | ID of the explained function | 
**FunctionName** | Pointer to **string** | Name of the explained function | [optional] 
**FunctionsInvolved** | Pointer to [**[]FilesystemExplainedFunction**](FilesystemExplainedFunction.md) | Other functions involved in the filesystem access | [optional] 
**Summary** | Pointer to **string** | Explanation of the filesystem access the function performs | [optional] 
**Targets** | Pointer to **[]string** | Concrete filesystem targets identified -- paths, registry keys, or environment variables | [optional] 

## Methods

### NewFilesystemAnalyseResult

`func NewFilesystemAnalyseResult(cancelled bool, functionId int64, ) *FilesystemAnalyseResult`

NewFilesystemAnalyseResult instantiates a new FilesystemAnalyseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemAnalyseResultWithDefaults

`func NewFilesystemAnalyseResultWithDefaults() *FilesystemAnalyseResult`

NewFilesystemAnalyseResultWithDefaults instantiates a new FilesystemAnalyseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelled

`func (o *FilesystemAnalyseResult) GetCancelled() bool`

GetCancelled returns the Cancelled field if non-nil, zero value otherwise.

### GetCancelledOk

`func (o *FilesystemAnalyseResult) GetCancelledOk() (*bool, bool)`

GetCancelledOk returns a tuple with the Cancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelled

`func (o *FilesystemAnalyseResult) SetCancelled(v bool)`

SetCancelled sets Cancelled field to given value.


### GetDataDescription

`func (o *FilesystemAnalyseResult) GetDataDescription() string`

GetDataDescription returns the DataDescription field if non-nil, zero value otherwise.

### GetDataDescriptionOk

`func (o *FilesystemAnalyseResult) GetDataDescriptionOk() (*string, bool)`

GetDataDescriptionOk returns a tuple with the DataDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDescription

`func (o *FilesystemAnalyseResult) SetDataDescription(v string)`

SetDataDescription sets DataDescription field to given value.

### HasDataDescription

`func (o *FilesystemAnalyseResult) HasDataDescription() bool`

HasDataDescription returns a boolean if a field has been set.

### GetDataDestination

`func (o *FilesystemAnalyseResult) GetDataDestination() string`

GetDataDestination returns the DataDestination field if non-nil, zero value otherwise.

### GetDataDestinationOk

`func (o *FilesystemAnalyseResult) GetDataDestinationOk() (*string, bool)`

GetDataDestinationOk returns a tuple with the DataDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDestination

`func (o *FilesystemAnalyseResult) SetDataDestination(v string)`

SetDataDestination sets DataDestination field to given value.

### HasDataDestination

`func (o *FilesystemAnalyseResult) HasDataDestination() bool`

HasDataDestination returns a boolean if a field has been set.

### GetDataOrigin

`func (o *FilesystemAnalyseResult) GetDataOrigin() string`

GetDataOrigin returns the DataOrigin field if non-nil, zero value otherwise.

### GetDataOriginOk

`func (o *FilesystemAnalyseResult) GetDataOriginOk() (*string, bool)`

GetDataOriginOk returns a tuple with the DataOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataOrigin

`func (o *FilesystemAnalyseResult) SetDataOrigin(v string)`

SetDataOrigin sets DataOrigin field to given value.

### HasDataOrigin

`func (o *FilesystemAnalyseResult) HasDataOrigin() bool`

HasDataOrigin returns a boolean if a field has been set.

### GetFunctionId

`func (o *FilesystemAnalyseResult) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *FilesystemAnalyseResult) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *FilesystemAnalyseResult) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *FilesystemAnalyseResult) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *FilesystemAnalyseResult) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *FilesystemAnalyseResult) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.

### HasFunctionName

`func (o *FilesystemAnalyseResult) HasFunctionName() bool`

HasFunctionName returns a boolean if a field has been set.

### GetFunctionsInvolved

`func (o *FilesystemAnalyseResult) GetFunctionsInvolved() []FilesystemExplainedFunction`

GetFunctionsInvolved returns the FunctionsInvolved field if non-nil, zero value otherwise.

### GetFunctionsInvolvedOk

`func (o *FilesystemAnalyseResult) GetFunctionsInvolvedOk() (*[]FilesystemExplainedFunction, bool)`

GetFunctionsInvolvedOk returns a tuple with the FunctionsInvolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionsInvolved

`func (o *FilesystemAnalyseResult) SetFunctionsInvolved(v []FilesystemExplainedFunction)`

SetFunctionsInvolved sets FunctionsInvolved field to given value.

### HasFunctionsInvolved

`func (o *FilesystemAnalyseResult) HasFunctionsInvolved() bool`

HasFunctionsInvolved returns a boolean if a field has been set.

### SetFunctionsInvolvedNil

`func (o *FilesystemAnalyseResult) SetFunctionsInvolvedNil(b bool)`

 SetFunctionsInvolvedNil sets the value for FunctionsInvolved to be an explicit nil

### UnsetFunctionsInvolved
`func (o *FilesystemAnalyseResult) UnsetFunctionsInvolved()`

UnsetFunctionsInvolved ensures that no value is present for FunctionsInvolved, not even an explicit nil
### GetSummary

`func (o *FilesystemAnalyseResult) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *FilesystemAnalyseResult) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *FilesystemAnalyseResult) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *FilesystemAnalyseResult) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTargets

`func (o *FilesystemAnalyseResult) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *FilesystemAnalyseResult) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *FilesystemAnalyseResult) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *FilesystemAnalyseResult) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### SetTargetsNil

`func (o *FilesystemAnalyseResult) SetTargetsNil(b bool)`

 SetTargetsNil sets the value for Targets to be an explicit nil

### UnsetTargets
`func (o *FilesystemAnalyseResult) UnsetTargets()`

UnsetTargets ensures that no value is present for Targets, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


