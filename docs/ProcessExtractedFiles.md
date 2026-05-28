# ProcessExtractedFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]ExtractedFileEntry**](ExtractedFileEntry.md) |  | [optional] 
**ProcessSeqid** | **int64** |  | 

## Methods

### NewProcessExtractedFiles

`func NewProcessExtractedFiles(processSeqid int64, ) *ProcessExtractedFiles`

NewProcessExtractedFiles instantiates a new ProcessExtractedFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessExtractedFilesWithDefaults

`func NewProcessExtractedFilesWithDefaults() *ProcessExtractedFiles`

NewProcessExtractedFilesWithDefaults instantiates a new ProcessExtractedFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *ProcessExtractedFiles) GetFiles() []ExtractedFileEntry`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ProcessExtractedFiles) GetFilesOk() (*[]ExtractedFileEntry, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ProcessExtractedFiles) SetFiles(v []ExtractedFileEntry)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ProcessExtractedFiles) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *ProcessExtractedFiles) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *ProcessExtractedFiles) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetProcessSeqid

`func (o *ProcessExtractedFiles) GetProcessSeqid() int64`

GetProcessSeqid returns the ProcessSeqid field if non-nil, zero value otherwise.

### GetProcessSeqidOk

`func (o *ProcessExtractedFiles) GetProcessSeqidOk() (*int64, bool)`

GetProcessSeqidOk returns a tuple with the ProcessSeqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessSeqid

`func (o *ProcessExtractedFiles) SetProcessSeqid(v int64)`

SetProcessSeqid sets ProcessSeqid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


