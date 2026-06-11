# Binary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int64** |  | 
**BinaryId** | **int64** |  | 
**BinaryName** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**IsSystemAnalysis** | **bool** |  | 
**OwnerId** | **int64** |  | 
**Sha256Hash** | **string** |  | 

## Methods

### NewBinary

`func NewBinary(analysisId int64, binaryId int64, binaryName string, createdAt time.Time, isSystemAnalysis bool, ownerId int64, sha256Hash string, ) *Binary`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


