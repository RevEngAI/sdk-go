# AnalysisRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int32** | ID to identify analysis | 
**AnalysisScope** | **string** | Scope of the analysis | 
**BinaryId** | **int32** | ID to identify the binary analyse | 
**ModelId** | **int32** | ID to identify the model used for analysis | 
**ModelName** | **string** | Name of the model used for analysis | 
**Status** | **string** | The current status of analysis | 
**Creation** | **time.Time** | The current status of analysis | 
**IsOwner** | **bool** | Whether the current user is the owner of a binary | 
**BinaryName** | **string** | The name of the file uploaded | 
**Sha256Hash** | **string** | The hash of the binary | 
**FunctionBoundariesHash** | **string** | The hash of the function boundaries | 
**BinarySize** | **int32** | The size of the binary | 
**Username** | **string** | The username of the analysis owner | 
**DynamicExecutionStatus** | Pointer to [**NullableAppApiRestV2AnalysesEnumsDynamicExecutionStatus**](AppApiRestV2AnalysesEnumsDynamicExecutionStatus.md) |  | [optional] 
**DynamicExecutionTaskId** | Pointer to **NullableInt32** |  | [optional] 
**BaseAddress** | **int32** | The base address of the binary | 

## Methods

### NewAnalysisRecord

`func NewAnalysisRecord(analysisId int32, analysisScope string, binaryId int32, modelId int32, modelName string, status string, creation time.Time, isOwner bool, binaryName string, sha256Hash string, functionBoundariesHash string, binarySize int32, username string, baseAddress int32, ) *AnalysisRecord`

NewAnalysisRecord instantiates a new AnalysisRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisRecordWithDefaults

`func NewAnalysisRecordWithDefaults() *AnalysisRecord`

NewAnalysisRecordWithDefaults instantiates a new AnalysisRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *AnalysisRecord) GetAnalysisId() int32`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *AnalysisRecord) GetAnalysisIdOk() (*int32, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *AnalysisRecord) SetAnalysisId(v int32)`

SetAnalysisId sets AnalysisId field to given value.


### GetAnalysisScope

`func (o *AnalysisRecord) GetAnalysisScope() string`

GetAnalysisScope returns the AnalysisScope field if non-nil, zero value otherwise.

### GetAnalysisScopeOk

`func (o *AnalysisRecord) GetAnalysisScopeOk() (*string, bool)`

GetAnalysisScopeOk returns a tuple with the AnalysisScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisScope

`func (o *AnalysisRecord) SetAnalysisScope(v string)`

SetAnalysisScope sets AnalysisScope field to given value.


### GetBinaryId

`func (o *AnalysisRecord) GetBinaryId() int32`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *AnalysisRecord) GetBinaryIdOk() (*int32, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *AnalysisRecord) SetBinaryId(v int32)`

SetBinaryId sets BinaryId field to given value.


### GetModelId

`func (o *AnalysisRecord) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *AnalysisRecord) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *AnalysisRecord) SetModelId(v int32)`

SetModelId sets ModelId field to given value.


### GetModelName

`func (o *AnalysisRecord) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *AnalysisRecord) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *AnalysisRecord) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetStatus

`func (o *AnalysisRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnalysisRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnalysisRecord) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreation

`func (o *AnalysisRecord) GetCreation() time.Time`

GetCreation returns the Creation field if non-nil, zero value otherwise.

### GetCreationOk

`func (o *AnalysisRecord) GetCreationOk() (*time.Time, bool)`

GetCreationOk returns a tuple with the Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreation

`func (o *AnalysisRecord) SetCreation(v time.Time)`

SetCreation sets Creation field to given value.


### GetIsOwner

`func (o *AnalysisRecord) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *AnalysisRecord) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *AnalysisRecord) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.


### GetBinaryName

`func (o *AnalysisRecord) GetBinaryName() string`

GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.

### GetBinaryNameOk

`func (o *AnalysisRecord) GetBinaryNameOk() (*string, bool)`

GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryName

`func (o *AnalysisRecord) SetBinaryName(v string)`

SetBinaryName sets BinaryName field to given value.


### GetSha256Hash

`func (o *AnalysisRecord) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *AnalysisRecord) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *AnalysisRecord) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetFunctionBoundariesHash

`func (o *AnalysisRecord) GetFunctionBoundariesHash() string`

GetFunctionBoundariesHash returns the FunctionBoundariesHash field if non-nil, zero value otherwise.

### GetFunctionBoundariesHashOk

`func (o *AnalysisRecord) GetFunctionBoundariesHashOk() (*string, bool)`

GetFunctionBoundariesHashOk returns a tuple with the FunctionBoundariesHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionBoundariesHash

`func (o *AnalysisRecord) SetFunctionBoundariesHash(v string)`

SetFunctionBoundariesHash sets FunctionBoundariesHash field to given value.


### GetBinarySize

`func (o *AnalysisRecord) GetBinarySize() int32`

GetBinarySize returns the BinarySize field if non-nil, zero value otherwise.

### GetBinarySizeOk

`func (o *AnalysisRecord) GetBinarySizeOk() (*int32, bool)`

GetBinarySizeOk returns a tuple with the BinarySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinarySize

`func (o *AnalysisRecord) SetBinarySize(v int32)`

SetBinarySize sets BinarySize field to given value.


### GetUsername

`func (o *AnalysisRecord) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AnalysisRecord) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AnalysisRecord) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDynamicExecutionStatus

`func (o *AnalysisRecord) GetDynamicExecutionStatus() AppApiRestV2AnalysesEnumsDynamicExecutionStatus`

GetDynamicExecutionStatus returns the DynamicExecutionStatus field if non-nil, zero value otherwise.

### GetDynamicExecutionStatusOk

`func (o *AnalysisRecord) GetDynamicExecutionStatusOk() (*AppApiRestV2AnalysesEnumsDynamicExecutionStatus, bool)`

GetDynamicExecutionStatusOk returns a tuple with the DynamicExecutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicExecutionStatus

`func (o *AnalysisRecord) SetDynamicExecutionStatus(v AppApiRestV2AnalysesEnumsDynamicExecutionStatus)`

SetDynamicExecutionStatus sets DynamicExecutionStatus field to given value.

### HasDynamicExecutionStatus

`func (o *AnalysisRecord) HasDynamicExecutionStatus() bool`

HasDynamicExecutionStatus returns a boolean if a field has been set.

### SetDynamicExecutionStatusNil

`func (o *AnalysisRecord) SetDynamicExecutionStatusNil(b bool)`

 SetDynamicExecutionStatusNil sets the value for DynamicExecutionStatus to be an explicit nil

### UnsetDynamicExecutionStatus
`func (o *AnalysisRecord) UnsetDynamicExecutionStatus()`

UnsetDynamicExecutionStatus ensures that no value is present for DynamicExecutionStatus, not even an explicit nil
### GetDynamicExecutionTaskId

`func (o *AnalysisRecord) GetDynamicExecutionTaskId() int32`

GetDynamicExecutionTaskId returns the DynamicExecutionTaskId field if non-nil, zero value otherwise.

### GetDynamicExecutionTaskIdOk

`func (o *AnalysisRecord) GetDynamicExecutionTaskIdOk() (*int32, bool)`

GetDynamicExecutionTaskIdOk returns a tuple with the DynamicExecutionTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicExecutionTaskId

`func (o *AnalysisRecord) SetDynamicExecutionTaskId(v int32)`

SetDynamicExecutionTaskId sets DynamicExecutionTaskId field to given value.

### HasDynamicExecutionTaskId

`func (o *AnalysisRecord) HasDynamicExecutionTaskId() bool`

HasDynamicExecutionTaskId returns a boolean if a field has been set.

### SetDynamicExecutionTaskIdNil

`func (o *AnalysisRecord) SetDynamicExecutionTaskIdNil(b bool)`

 SetDynamicExecutionTaskIdNil sets the value for DynamicExecutionTaskId to be an explicit nil

### UnsetDynamicExecutionTaskId
`func (o *AnalysisRecord) UnsetDynamicExecutionTaskId()`

UnsetDynamicExecutionTaskId ensures that no value is present for DynamicExecutionTaskId, not even an explicit nil
### GetBaseAddress

`func (o *AnalysisRecord) GetBaseAddress() int32`

GetBaseAddress returns the BaseAddress field if non-nil, zero value otherwise.

### GetBaseAddressOk

`func (o *AnalysisRecord) GetBaseAddressOk() (*int32, bool)`

GetBaseAddressOk returns a tuple with the BaseAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAddress

`func (o *AnalysisRecord) SetBaseAddress(v int32)`

SetBaseAddress sets BaseAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


