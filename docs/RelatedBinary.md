# RelatedBinary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **NullableInt64** | Most recent analysis of the related binary, null when it has never been analysed | 
**BinaryId** | **int64** | ID of the related binary | 
**Name** | **string** | Name of the related binary | 
**Sha256** | **string** | SHA-256 of the related binary | 

## Methods

### NewRelatedBinary

`func NewRelatedBinary(analysisId NullableInt64, binaryId int64, name string, sha256 string, ) *RelatedBinary`

NewRelatedBinary instantiates a new RelatedBinary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedBinaryWithDefaults

`func NewRelatedBinaryWithDefaults() *RelatedBinary`

NewRelatedBinaryWithDefaults instantiates a new RelatedBinary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *RelatedBinary) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *RelatedBinary) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *RelatedBinary) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### SetAnalysisIdNil

`func (o *RelatedBinary) SetAnalysisIdNil(b bool)`

 SetAnalysisIdNil sets the value for AnalysisId to be an explicit nil

### UnsetAnalysisId
`func (o *RelatedBinary) UnsetAnalysisId()`

UnsetAnalysisId ensures that no value is present for AnalysisId, not even an explicit nil
### GetBinaryId

`func (o *RelatedBinary) GetBinaryId() int64`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *RelatedBinary) GetBinaryIdOk() (*int64, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *RelatedBinary) SetBinaryId(v int64)`

SetBinaryId sets BinaryId field to given value.


### GetName

`func (o *RelatedBinary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RelatedBinary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RelatedBinary) SetName(v string)`

SetName sets Name field to given value.


### GetSha256

`func (o *RelatedBinary) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *RelatedBinary) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *RelatedBinary) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


