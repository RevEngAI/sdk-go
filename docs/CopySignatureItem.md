# CopySignatureItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceFunctionId** | **int64** | Function to copy the signature from. May belong to another analysis the caller can read. | 
**TargetFunctionId** | **int64** | Function to copy the signature to. Must belong to the analysis in the URL. | 

## Methods

### NewCopySignatureItem

`func NewCopySignatureItem(sourceFunctionId int64, targetFunctionId int64, ) *CopySignatureItem`

NewCopySignatureItem instantiates a new CopySignatureItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopySignatureItemWithDefaults

`func NewCopySignatureItemWithDefaults() *CopySignatureItem`

NewCopySignatureItemWithDefaults instantiates a new CopySignatureItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceFunctionId

`func (o *CopySignatureItem) GetSourceFunctionId() int64`

GetSourceFunctionId returns the SourceFunctionId field if non-nil, zero value otherwise.

### GetSourceFunctionIdOk

`func (o *CopySignatureItem) GetSourceFunctionIdOk() (*int64, bool)`

GetSourceFunctionIdOk returns a tuple with the SourceFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFunctionId

`func (o *CopySignatureItem) SetSourceFunctionId(v int64)`

SetSourceFunctionId sets SourceFunctionId field to given value.


### GetTargetFunctionId

`func (o *CopySignatureItem) GetTargetFunctionId() int64`

GetTargetFunctionId returns the TargetFunctionId field if non-nil, zero value otherwise.

### GetTargetFunctionIdOk

`func (o *CopySignatureItem) GetTargetFunctionIdOk() (*int64, bool)`

GetTargetFunctionIdOk returns a tuple with the TargetFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFunctionId

`func (o *CopySignatureItem) SetTargetFunctionId(v int64)`

SetTargetFunctionId sets TargetFunctionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


