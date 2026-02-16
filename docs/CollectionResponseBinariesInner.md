# CollectionResponseBinariesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int32** | Analysis ID | 
**BinaryId** | **int32** | Binary ID | 
**BinaryName** | **string** | Binary name | 
**OwnerId** | **int32** | Binary owner | 
**Sha256Hash** | **string** | Binary SHA-256 hash | 
**CreatedAt** | **time.Time** | Binary creation date | 
**IsSystemAnalysis** | **bool** | Is the analysis owned by a RevEng.AI account | 

## Methods

### NewCollectionResponseBinariesInner

`func NewCollectionResponseBinariesInner(analysisId int32, binaryId int32, binaryName string, ownerId int32, sha256Hash string, createdAt time.Time, isSystemAnalysis bool, ) *CollectionResponseBinariesInner`

NewCollectionResponseBinariesInner instantiates a new CollectionResponseBinariesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseBinariesInnerWithDefaults

`func NewCollectionResponseBinariesInnerWithDefaults() *CollectionResponseBinariesInner`

NewCollectionResponseBinariesInnerWithDefaults instantiates a new CollectionResponseBinariesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *CollectionResponseBinariesInner) GetAnalysisId() int32`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *CollectionResponseBinariesInner) GetAnalysisIdOk() (*int32, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *CollectionResponseBinariesInner) SetAnalysisId(v int32)`

SetAnalysisId sets AnalysisId field to given value.


### GetBinaryId

`func (o *CollectionResponseBinariesInner) GetBinaryId() int32`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *CollectionResponseBinariesInner) GetBinaryIdOk() (*int32, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *CollectionResponseBinariesInner) SetBinaryId(v int32)`

SetBinaryId sets BinaryId field to given value.


### GetBinaryName

`func (o *CollectionResponseBinariesInner) GetBinaryName() string`

GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.

### GetBinaryNameOk

`func (o *CollectionResponseBinariesInner) GetBinaryNameOk() (*string, bool)`

GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryName

`func (o *CollectionResponseBinariesInner) SetBinaryName(v string)`

SetBinaryName sets BinaryName field to given value.


### GetOwnerId

`func (o *CollectionResponseBinariesInner) GetOwnerId() int32`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *CollectionResponseBinariesInner) GetOwnerIdOk() (*int32, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *CollectionResponseBinariesInner) SetOwnerId(v int32)`

SetOwnerId sets OwnerId field to given value.


### GetSha256Hash

`func (o *CollectionResponseBinariesInner) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *CollectionResponseBinariesInner) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *CollectionResponseBinariesInner) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.


### GetCreatedAt

`func (o *CollectionResponseBinariesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CollectionResponseBinariesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CollectionResponseBinariesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetIsSystemAnalysis

`func (o *CollectionResponseBinariesInner) GetIsSystemAnalysis() bool`

GetIsSystemAnalysis returns the IsSystemAnalysis field if non-nil, zero value otherwise.

### GetIsSystemAnalysisOk

`func (o *CollectionResponseBinariesInner) GetIsSystemAnalysisOk() (*bool, bool)`

GetIsSystemAnalysisOk returns a tuple with the IsSystemAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemAnalysis

`func (o *CollectionResponseBinariesInner) SetIsSystemAnalysis(v bool)`

SetIsSystemAnalysis sets IsSystemAnalysis field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


