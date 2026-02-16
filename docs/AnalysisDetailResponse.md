# AnalysisDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | [**AnalysisAccessInfo**](AnalysisAccessInfo.md) |  | 
**AnalysisId** | **int32** |  | 
**AnalysisScope** | **string** |  | 
**Architecture** | **string** |  | 
**BinaryDynamic** | **bool** |  | 
**BinaryFormat** | **string** |  | 
**BinaryName** | **string** |  | 
**BinarySize** | **int32** |  | 
**BinaryType** | **string** |  | 
**Creation** | **string** |  | 
**Debug** | **bool** |  | 
**ModelName** | **string** |  | 
**Sbom** | Pointer to **map[string]interface{}** |  | [optional] 
**Sha256Hash** | **string** |  | 

## Methods

### NewAnalysisDetailResponse

`func NewAnalysisDetailResponse(access AnalysisAccessInfo, analysisId int32, analysisScope string, architecture string, binaryDynamic bool, binaryFormat string, binaryName string, binarySize int32, binaryType string, creation string, debug bool, modelName string, sha256Hash string, ) *AnalysisDetailResponse`

NewAnalysisDetailResponse instantiates a new AnalysisDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisDetailResponseWithDefaults

`func NewAnalysisDetailResponseWithDefaults() *AnalysisDetailResponse`

NewAnalysisDetailResponseWithDefaults instantiates a new AnalysisDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *AnalysisDetailResponse) GetAccess() AnalysisAccessInfo`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AnalysisDetailResponse) GetAccessOk() (*AnalysisAccessInfo, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AnalysisDetailResponse) SetAccess(v AnalysisAccessInfo)`

SetAccess sets Access field to given value.


### GetAnalysisId

`func (o *AnalysisDetailResponse) GetAnalysisId() int32`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *AnalysisDetailResponse) GetAnalysisIdOk() (*int32, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *AnalysisDetailResponse) SetAnalysisId(v int32)`

SetAnalysisId sets AnalysisId field to given value.


### GetAnalysisScope

`func (o *AnalysisDetailResponse) GetAnalysisScope() string`

GetAnalysisScope returns the AnalysisScope field if non-nil, zero value otherwise.

### GetAnalysisScopeOk

`func (o *AnalysisDetailResponse) GetAnalysisScopeOk() (*string, bool)`

GetAnalysisScopeOk returns a tuple with the AnalysisScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisScope

`func (o *AnalysisDetailResponse) SetAnalysisScope(v string)`

SetAnalysisScope sets AnalysisScope field to given value.


### GetArchitecture

`func (o *AnalysisDetailResponse) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *AnalysisDetailResponse) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *AnalysisDetailResponse) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetBinaryDynamic

`func (o *AnalysisDetailResponse) GetBinaryDynamic() bool`

GetBinaryDynamic returns the BinaryDynamic field if non-nil, zero value otherwise.

### GetBinaryDynamicOk

`func (o *AnalysisDetailResponse) GetBinaryDynamicOk() (*bool, bool)`

GetBinaryDynamicOk returns a tuple with the BinaryDynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDynamic

`func (o *AnalysisDetailResponse) SetBinaryDynamic(v bool)`

SetBinaryDynamic sets BinaryDynamic field to given value.


### GetBinaryFormat

`func (o *AnalysisDetailResponse) GetBinaryFormat() string`

GetBinaryFormat returns the BinaryFormat field if non-nil, zero value otherwise.

### GetBinaryFormatOk

`func (o *AnalysisDetailResponse) GetBinaryFormatOk() (*string, bool)`

GetBinaryFormatOk returns a tuple with the BinaryFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryFormat

`func (o *AnalysisDetailResponse) SetBinaryFormat(v string)`

SetBinaryFormat sets BinaryFormat field to given value.


### GetBinaryName

`func (o *AnalysisDetailResponse) GetBinaryName() string`

GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.

### GetBinaryNameOk

`func (o *AnalysisDetailResponse) GetBinaryNameOk() (*string, bool)`

GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryName

`func (o *AnalysisDetailResponse) SetBinaryName(v string)`

SetBinaryName sets BinaryName field to given value.


### GetBinarySize

`func (o *AnalysisDetailResponse) GetBinarySize() int32`

GetBinarySize returns the BinarySize field if non-nil, zero value otherwise.

### GetBinarySizeOk

`func (o *AnalysisDetailResponse) GetBinarySizeOk() (*int32, bool)`

GetBinarySizeOk returns a tuple with the BinarySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinarySize

`func (o *AnalysisDetailResponse) SetBinarySize(v int32)`

SetBinarySize sets BinarySize field to given value.


### GetBinaryType

`func (o *AnalysisDetailResponse) GetBinaryType() string`

GetBinaryType returns the BinaryType field if non-nil, zero value otherwise.

### GetBinaryTypeOk

`func (o *AnalysisDetailResponse) GetBinaryTypeOk() (*string, bool)`

GetBinaryTypeOk returns a tuple with the BinaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryType

`func (o *AnalysisDetailResponse) SetBinaryType(v string)`

SetBinaryType sets BinaryType field to given value.


### GetCreation

`func (o *AnalysisDetailResponse) GetCreation() string`

GetCreation returns the Creation field if non-nil, zero value otherwise.

### GetCreationOk

`func (o *AnalysisDetailResponse) GetCreationOk() (*string, bool)`

GetCreationOk returns a tuple with the Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreation

`func (o *AnalysisDetailResponse) SetCreation(v string)`

SetCreation sets Creation field to given value.


### GetDebug

`func (o *AnalysisDetailResponse) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *AnalysisDetailResponse) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *AnalysisDetailResponse) SetDebug(v bool)`

SetDebug sets Debug field to given value.


### GetModelName

`func (o *AnalysisDetailResponse) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *AnalysisDetailResponse) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *AnalysisDetailResponse) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetSbom

`func (o *AnalysisDetailResponse) GetSbom() map[string]interface{}`

GetSbom returns the Sbom field if non-nil, zero value otherwise.

### GetSbomOk

`func (o *AnalysisDetailResponse) GetSbomOk() (*map[string]interface{}, bool)`

GetSbomOk returns a tuple with the Sbom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbom

`func (o *AnalysisDetailResponse) SetSbom(v map[string]interface{})`

SetSbom sets Sbom field to given value.

### HasSbom

`func (o *AnalysisDetailResponse) HasSbom() bool`

HasSbom returns a boolean if a field has been set.

### SetSbomNil

`func (o *AnalysisDetailResponse) SetSbomNil(b bool)`

 SetSbomNil sets the value for Sbom to be an explicit nil

### UnsetSbom
`func (o *AnalysisDetailResponse) UnsetSbom()`

UnsetSbom ensures that no value is present for Sbom, not even an explicit nil
### GetSha256Hash

`func (o *AnalysisDetailResponse) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *AnalysisDetailResponse) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *AnalysisDetailResponse) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


