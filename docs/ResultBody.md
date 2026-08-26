# ResultBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binaries** | [**[]ExtractedBinary**](ExtractedBinary.md) | Child binaries recovered from the extraction. | 
**ExtractionDepth** | **int64** | Number of nested-archive extraction passes taken. | 
**FilenameToExtractionFailure** | **map[string]string** | Per-file extraction failures, keyed by filename. | 
**SkippedFiles** | **int64** | Files skipped because they were not recognised as binaries. | 

## Methods

### NewResultBody

`func NewResultBody(binaries []ExtractedBinary, extractionDepth int64, filenameToExtractionFailure map[string]string, skippedFiles int64, ) *ResultBody`

NewResultBody instantiates a new ResultBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultBodyWithDefaults

`func NewResultBodyWithDefaults() *ResultBody`

NewResultBodyWithDefaults instantiates a new ResultBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaries

`func (o *ResultBody) GetBinaries() []ExtractedBinary`

GetBinaries returns the Binaries field if non-nil, zero value otherwise.

### GetBinariesOk

`func (o *ResultBody) GetBinariesOk() (*[]ExtractedBinary, bool)`

GetBinariesOk returns a tuple with the Binaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaries

`func (o *ResultBody) SetBinaries(v []ExtractedBinary)`

SetBinaries sets Binaries field to given value.


### SetBinariesNil

`func (o *ResultBody) SetBinariesNil(b bool)`

 SetBinariesNil sets the value for Binaries to be an explicit nil

### UnsetBinaries
`func (o *ResultBody) UnsetBinaries()`

UnsetBinaries ensures that no value is present for Binaries, not even an explicit nil
### GetExtractionDepth

`func (o *ResultBody) GetExtractionDepth() int64`

GetExtractionDepth returns the ExtractionDepth field if non-nil, zero value otherwise.

### GetExtractionDepthOk

`func (o *ResultBody) GetExtractionDepthOk() (*int64, bool)`

GetExtractionDepthOk returns a tuple with the ExtractionDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionDepth

`func (o *ResultBody) SetExtractionDepth(v int64)`

SetExtractionDepth sets ExtractionDepth field to given value.


### GetFilenameToExtractionFailure

`func (o *ResultBody) GetFilenameToExtractionFailure() map[string]string`

GetFilenameToExtractionFailure returns the FilenameToExtractionFailure field if non-nil, zero value otherwise.

### GetFilenameToExtractionFailureOk

`func (o *ResultBody) GetFilenameToExtractionFailureOk() (*map[string]string, bool)`

GetFilenameToExtractionFailureOk returns a tuple with the FilenameToExtractionFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilenameToExtractionFailure

`func (o *ResultBody) SetFilenameToExtractionFailure(v map[string]string)`

SetFilenameToExtractionFailure sets FilenameToExtractionFailure field to given value.


### GetSkippedFiles

`func (o *ResultBody) GetSkippedFiles() int64`

GetSkippedFiles returns the SkippedFiles field if non-nil, zero value otherwise.

### GetSkippedFilesOk

`func (o *ResultBody) GetSkippedFilesOk() (*int64, bool)`

GetSkippedFilesOk returns a tuple with the SkippedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedFiles

`func (o *ResultBody) SetSkippedFiles(v int64)`

SetSkippedFiles sets SkippedFiles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


