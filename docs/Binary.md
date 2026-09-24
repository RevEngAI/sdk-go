# Binary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int64** |  | 
**BinaryId** | **int64** |  | 
**BinaryName** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**DetectedArchitecture** | **string** | Detected instruction-set architecture; empty when unavailable | 
**DetectedBinaryType** | **string** | Detected operating-system platform; empty when unavailable | 
**IsSystemAnalysis** | **bool** |  | 
**ModelName** | **string** | Name of the model the analysis ran on | 
**OwnerId** | **int64** |  | 
**Sha256Hash** | **string** |  | 
**SuppliedArchitecture** | **string** | User-supplied instruction-set architecture; \&quot;AUTO\&quot; when not overridden | 
**SuppliedBinaryType** | **string** | User-supplied operating-system platform; \&quot;AUTO\&quot; when not overridden | 

## Methods

### NewBinary

`func NewBinary(analysisId int64, binaryId int64, binaryName string, createdAt time.Time, detectedArchitecture string, detectedBinaryType string, isSystemAnalysis bool, modelName string, ownerId int64, sha256Hash string, suppliedArchitecture string, suppliedBinaryType string, ) *Binary`

NewBinary instantiates a new Binary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinaryWithDefaults

`func NewBinaryWithDefaults() *Binary`

NewBinaryWithDefaults instantiates a new Binary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *Binary) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *Binary) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *Binary) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### GetBinaryId

`func (o *Binary) GetBinaryId() int64`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *Binary) GetBinaryIdOk() (*int64, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *Binary) SetBinaryId(v int64)`

SetBinaryId sets BinaryId field to given value.


### GetBinaryName

`func (o *Binary) GetBinaryName() string`

GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.

### GetBinaryNameOk

`func (o *Binary) GetBinaryNameOk() (*string, bool)`

GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryName

`func (o *Binary) SetBinaryName(v string)`

SetBinaryName sets BinaryName field to given value.


### GetCreatedAt

`func (o *Binary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Binary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Binary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDetectedArchitecture

`func (o *Binary) GetDetectedArchitecture() string`

GetDetectedArchitecture returns the DetectedArchitecture field if non-nil, zero value otherwise.

### GetDetectedArchitectureOk

`func (o *Binary) GetDetectedArchitectureOk() (*string, bool)`

GetDetectedArchitectureOk returns a tuple with the DetectedArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedArchitecture

`func (o *Binary) SetDetectedArchitecture(v string)`

SetDetectedArchitecture sets DetectedArchitecture field to given value.


### GetDetectedBinaryType

`func (o *Binary) GetDetectedBinaryType() string`

GetDetectedBinaryType returns the DetectedBinaryType field if non-nil, zero value otherwise.

### GetDetectedBinaryTypeOk

`func (o *Binary) GetDetectedBinaryTypeOk() (*string, bool)`

GetDetectedBinaryTypeOk returns a tuple with the DetectedBinaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedBinaryType

`func (o *Binary) SetDetectedBinaryType(v string)`

SetDetectedBinaryType sets DetectedBinaryType field to given value.


### GetIsSystemAnalysis

`func (o *Binary) GetIsSystemAnalysis() bool`

GetIsSystemAnalysis returns the IsSystemAnalysis field if non-nil, zero value otherwise.

### GetIsSystemAnalysisOk

`func (o *Binary) GetIsSystemAnalysisOk() (*bool, bool)`

GetIsSystemAnalysisOk returns a tuple with the IsSystemAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemAnalysis

`func (o *Binary) SetIsSystemAnalysis(v bool)`

SetIsSystemAnalysis sets IsSystemAnalysis field to given value.


### GetModelName

`func (o *Binary) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *Binary) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *Binary) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetOwnerId

`func (o *Binary) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Binary) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Binary) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.


### GetSha256Hash

`func (o *Binary) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *Binary) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *Binary) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetSuppliedArchitecture

`func (o *Binary) GetSuppliedArchitecture() string`

GetSuppliedArchitecture returns the SuppliedArchitecture field if non-nil, zero value otherwise.

### GetSuppliedArchitectureOk

`func (o *Binary) GetSuppliedArchitectureOk() (*string, bool)`

GetSuppliedArchitectureOk returns a tuple with the SuppliedArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppliedArchitecture

`func (o *Binary) SetSuppliedArchitecture(v string)`

SetSuppliedArchitecture sets SuppliedArchitecture field to given value.


### GetSuppliedBinaryType

`func (o *Binary) GetSuppliedBinaryType() string`

GetSuppliedBinaryType returns the SuppliedBinaryType field if non-nil, zero value otherwise.

### GetSuppliedBinaryTypeOk

`func (o *Binary) GetSuppliedBinaryTypeOk() (*string, bool)`

GetSuppliedBinaryTypeOk returns a tuple with the SuppliedBinaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppliedBinaryType

`func (o *Binary) SetSuppliedBinaryType(v string)`

SetSuppliedBinaryType sets SuppliedBinaryType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


