# Created

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisId** | **int32** | The ID corresponding to the newly created analysis | 
**BinaryId** | **int32** | The ID corresponding to the binary that was created | 
**Reference** | **string** | Deprecated will always be empty string | 

## Methods

### NewCreated

`func NewCreated(analysisId int32, binaryId int32, reference string, ) *Created`

NewCreated instantiates a new Created object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedWithDefaults

`func NewCreatedWithDefaults() *Created`

NewCreatedWithDefaults instantiates a new Created object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisId

`func (o *Created) GetAnalysisId() int32`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *Created) GetAnalysisIdOk() (*int32, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *Created) SetAnalysisId(v int32)`

SetAnalysisId sets AnalysisId field to given value.


### GetBinaryId

`func (o *Created) GetBinaryId() int32`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *Created) GetBinaryIdOk() (*int32, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *Created) SetBinaryId(v int32)`

SetBinaryId sets BinaryId field to given value.


### GetReference

`func (o *Created) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Created) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Created) SetReference(v string)`

SetReference sets Reference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


