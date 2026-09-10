# AnalysisRecordBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int64** | Analysis ID | 
**AnalysisScope** | **string** | Scope of the analysis | 
**BaseAddress** | **int64** | Binary base address | 
**BinaryId** | **int64** | Binary ID | 
**BinaryName** | **string** | Binary filename | 
**BinarySize** | **int64** | Binary size in bytes | 
**Creation** | **time.Time** | When the analysis was created | 
**DetectedArchitecture** | **string** | Detected instruction-set architecture; empty when unavailable | 
**DetectedBinaryFormat** | **string** | Detected binary container format; empty when unavailable | 
**DetectedBinaryType** | **string** | Detected operating-system platform; empty when unavailable | 
**FunctionBoundariesHash** | **string** | Hash of the binary&#39;s provided function boundaries | 
**IsOwner** | **bool** | True when the caller owns the analysis | 
**ModelId** | **int64** | Model ID | 
**ModelName** | **string** | Model name | 
**ModelUpgradeAvailable** | **bool** | True when the analysis ran on a model older than the current one, so its owner can re-analyse it on the latest. Describes the analysis, not the caller&#39;s rights — only the owner may act on it | 
**Sha256Hash** | **string** | SHA-256 hash of the binary | 
**Status** | **string** | Analysis status | 
**SuppliedArchitecture** | **string** | User-supplied instruction-set architecture; \&quot;AUTO\&quot; when not overridden | 
**SuppliedBinaryFormat** | **string** | User-supplied binary container format; \&quot;AUTO\&quot; when not overridden | 
**SuppliedBinaryType** | **string** | User-supplied operating-system platform; \&quot;AUTO\&quot; when not overridden | 
**Tags** | [**[]AnalysisTagBody**](AnalysisTagBody.md) | Tags associated with the binary | 
**Username** | **string** | Username of the analysis owner | 

## Methods

### NewAnalysisRecordBody

`func NewAnalysisRecordBody(analysisId int64, analysisScope string, baseAddress int64, binaryId int64, binaryName string, binarySize int64, creation time.Time, detectedArchitecture string, detectedBinaryFormat string, detectedBinaryType string, functionBoundariesHash string, isOwner bool, modelId int64, modelName string, modelUpgradeAvailable bool, sha256Hash string, status string, suppliedArchitecture string, suppliedBinaryFormat string, suppliedBinaryType string, tags []AnalysisTagBody, username string, ) *AnalysisRecordBody`

NewAnalysisRecordBody instantiates a new AnalysisRecordBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisRecordBodyWithDefaults

`func NewAnalysisRecordBodyWithDefaults() *AnalysisRecordBody`

NewAnalysisRecordBodyWithDefaults instantiates a new AnalysisRecordBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *AnalysisRecordBody) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *AnalysisRecordBody) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *AnalysisRecordBody) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### GetAnalysisScope

`func (o *AnalysisRecordBody) GetAnalysisScope() string`

GetAnalysisScope returns the AnalysisScope field if non-nil, zero value otherwise.

### GetAnalysisScopeOk

`func (o *AnalysisRecordBody) GetAnalysisScopeOk() (*string, bool)`

GetAnalysisScopeOk returns a tuple with the AnalysisScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisScope

`func (o *AnalysisRecordBody) SetAnalysisScope(v string)`

SetAnalysisScope sets AnalysisScope field to given value.


### GetBaseAddress

`func (o *AnalysisRecordBody) GetBaseAddress() int64`

GetBaseAddress returns the BaseAddress field if non-nil, zero value otherwise.

### GetBaseAddressOk

`func (o *AnalysisRecordBody) GetBaseAddressOk() (*int64, bool)`

GetBaseAddressOk returns a tuple with the BaseAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAddress

`func (o *AnalysisRecordBody) SetBaseAddress(v int64)`

SetBaseAddress sets BaseAddress field to given value.


### GetBinaryId

`func (o *AnalysisRecordBody) GetBinaryId() int64`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *AnalysisRecordBody) GetBinaryIdOk() (*int64, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *AnalysisRecordBody) SetBinaryId(v int64)`

SetBinaryId sets BinaryId field to given value.


### GetBinaryName

`func (o *AnalysisRecordBody) GetBinaryName() string`

GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.

### GetBinaryNameOk

`func (o *AnalysisRecordBody) GetBinaryNameOk() (*string, bool)`

GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryName

`func (o *AnalysisRecordBody) SetBinaryName(v string)`

SetBinaryName sets BinaryName field to given value.


### GetBinarySize

`func (o *AnalysisRecordBody) GetBinarySize() int64`

GetBinarySize returns the BinarySize field if non-nil, zero value otherwise.

### GetBinarySizeOk

`func (o *AnalysisRecordBody) GetBinarySizeOk() (*int64, bool)`

GetBinarySizeOk returns a tuple with the BinarySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinarySize

`func (o *AnalysisRecordBody) SetBinarySize(v int64)`

SetBinarySize sets BinarySize field to given value.


### GetCreation

`func (o *AnalysisRecordBody) GetCreation() time.Time`

GetCreation returns the Creation field if non-nil, zero value otherwise.

### GetCreationOk

`func (o *AnalysisRecordBody) GetCreationOk() (*time.Time, bool)`

GetCreationOk returns a tuple with the Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreation

`func (o *AnalysisRecordBody) SetCreation(v time.Time)`

SetCreation sets Creation field to given value.


### GetDetectedArchitecture

`func (o *AnalysisRecordBody) GetDetectedArchitecture() string`

GetDetectedArchitecture returns the DetectedArchitecture field if non-nil, zero value otherwise.

### GetDetectedArchitectureOk

`func (o *AnalysisRecordBody) GetDetectedArchitectureOk() (*string, bool)`

GetDetectedArchitectureOk returns a tuple with the DetectedArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedArchitecture

`func (o *AnalysisRecordBody) SetDetectedArchitecture(v string)`

SetDetectedArchitecture sets DetectedArchitecture field to given value.


### GetDetectedBinaryFormat

`func (o *AnalysisRecordBody) GetDetectedBinaryFormat() string`

GetDetectedBinaryFormat returns the DetectedBinaryFormat field if non-nil, zero value otherwise.

### GetDetectedBinaryFormatOk

`func (o *AnalysisRecordBody) GetDetectedBinaryFormatOk() (*string, bool)`

GetDetectedBinaryFormatOk returns a tuple with the DetectedBinaryFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedBinaryFormat

`func (o *AnalysisRecordBody) SetDetectedBinaryFormat(v string)`

SetDetectedBinaryFormat sets DetectedBinaryFormat field to given value.


### GetDetectedBinaryType

`func (o *AnalysisRecordBody) GetDetectedBinaryType() string`

GetDetectedBinaryType returns the DetectedBinaryType field if non-nil, zero value otherwise.

### GetDetectedBinaryTypeOk

`func (o *AnalysisRecordBody) GetDetectedBinaryTypeOk() (*string, bool)`

GetDetectedBinaryTypeOk returns a tuple with the DetectedBinaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedBinaryType

`func (o *AnalysisRecordBody) SetDetectedBinaryType(v string)`

SetDetectedBinaryType sets DetectedBinaryType field to given value.


### GetFunctionBoundariesHash

`func (o *AnalysisRecordBody) GetFunctionBoundariesHash() string`

GetFunctionBoundariesHash returns the FunctionBoundariesHash field if non-nil, zero value otherwise.

### GetFunctionBoundariesHashOk

`func (o *AnalysisRecordBody) GetFunctionBoundariesHashOk() (*string, bool)`

GetFunctionBoundariesHashOk returns a tuple with the FunctionBoundariesHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionBoundariesHash

`func (o *AnalysisRecordBody) SetFunctionBoundariesHash(v string)`

SetFunctionBoundariesHash sets FunctionBoundariesHash field to given value.


### GetIsOwner

`func (o *AnalysisRecordBody) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *AnalysisRecordBody) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *AnalysisRecordBody) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.


### GetModelId

`func (o *AnalysisRecordBody) GetModelId() int64`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *AnalysisRecordBody) GetModelIdOk() (*int64, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *AnalysisRecordBody) SetModelId(v int64)`

SetModelId sets ModelId field to given value.


### GetModelName

`func (o *AnalysisRecordBody) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *AnalysisRecordBody) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *AnalysisRecordBody) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetModelUpgradeAvailable

`func (o *AnalysisRecordBody) GetModelUpgradeAvailable() bool`

GetModelUpgradeAvailable returns the ModelUpgradeAvailable field if non-nil, zero value otherwise.

### GetModelUpgradeAvailableOk

`func (o *AnalysisRecordBody) GetModelUpgradeAvailableOk() (*bool, bool)`

GetModelUpgradeAvailableOk returns a tuple with the ModelUpgradeAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelUpgradeAvailable

`func (o *AnalysisRecordBody) SetModelUpgradeAvailable(v bool)`

SetModelUpgradeAvailable sets ModelUpgradeAvailable field to given value.


### GetSha256Hash

`func (o *AnalysisRecordBody) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *AnalysisRecordBody) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *AnalysisRecordBody) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetStatus

`func (o *AnalysisRecordBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnalysisRecordBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnalysisRecordBody) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSuppliedArchitecture

`func (o *AnalysisRecordBody) GetSuppliedArchitecture() string`

GetSuppliedArchitecture returns the SuppliedArchitecture field if non-nil, zero value otherwise.

### GetSuppliedArchitectureOk

`func (o *AnalysisRecordBody) GetSuppliedArchitectureOk() (*string, bool)`

GetSuppliedArchitectureOk returns a tuple with the SuppliedArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppliedArchitecture

`func (o *AnalysisRecordBody) SetSuppliedArchitecture(v string)`

SetSuppliedArchitecture sets SuppliedArchitecture field to given value.


### GetSuppliedBinaryFormat

`func (o *AnalysisRecordBody) GetSuppliedBinaryFormat() string`

GetSuppliedBinaryFormat returns the SuppliedBinaryFormat field if non-nil, zero value otherwise.

### GetSuppliedBinaryFormatOk

`func (o *AnalysisRecordBody) GetSuppliedBinaryFormatOk() (*string, bool)`

GetSuppliedBinaryFormatOk returns a tuple with the SuppliedBinaryFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppliedBinaryFormat

`func (o *AnalysisRecordBody) SetSuppliedBinaryFormat(v string)`

SetSuppliedBinaryFormat sets SuppliedBinaryFormat field to given value.


### GetSuppliedBinaryType

`func (o *AnalysisRecordBody) GetSuppliedBinaryType() string`

GetSuppliedBinaryType returns the SuppliedBinaryType field if non-nil, zero value otherwise.

### GetSuppliedBinaryTypeOk

`func (o *AnalysisRecordBody) GetSuppliedBinaryTypeOk() (*string, bool)`

GetSuppliedBinaryTypeOk returns a tuple with the SuppliedBinaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppliedBinaryType

`func (o *AnalysisRecordBody) SetSuppliedBinaryType(v string)`

SetSuppliedBinaryType sets SuppliedBinaryType field to given value.


### GetTags

`func (o *AnalysisRecordBody) GetTags() []AnalysisTagBody`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AnalysisRecordBody) GetTagsOk() (*[]AnalysisTagBody, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AnalysisRecordBody) SetTags(v []AnalysisTagBody)`

SetTags sets Tags field to given value.


### SetTagsNil

`func (o *AnalysisRecordBody) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AnalysisRecordBody) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetUsername

`func (o *AnalysisRecordBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AnalysisRecordBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AnalysisRecordBody) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


