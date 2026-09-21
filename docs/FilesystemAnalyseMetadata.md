# FilesystemAnalyseMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogHistory** | Pointer to **[][]interface{}** | Progress messages the run recorded, oldest first. | [optional] 
**Status** | **string** | Run status. UNINITIALISED means the agent has never been triggered for this function. | 

## Methods

### NewFilesystemAnalyseMetadata

`func NewFilesystemAnalyseMetadata(status string, ) *FilesystemAnalyseMetadata`

NewFilesystemAnalyseMetadata instantiates a new FilesystemAnalyseMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemAnalyseMetadataWithDefaults

`func NewFilesystemAnalyseMetadataWithDefaults() *FilesystemAnalyseMetadata`

NewFilesystemAnalyseMetadataWithDefaults instantiates a new FilesystemAnalyseMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogHistory

`func (o *FilesystemAnalyseMetadata) GetLogHistory() [][]interface{}`

GetLogHistory returns the LogHistory field if non-nil, zero value otherwise.

### GetLogHistoryOk

`func (o *FilesystemAnalyseMetadata) GetLogHistoryOk() (*[][]interface{}, bool)`

GetLogHistoryOk returns a tuple with the LogHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogHistory

`func (o *FilesystemAnalyseMetadata) SetLogHistory(v [][]interface{})`

SetLogHistory sets LogHistory field to given value.

### HasLogHistory

`func (o *FilesystemAnalyseMetadata) HasLogHistory() bool`

HasLogHistory returns a boolean if a field has been set.

### SetLogHistoryNil

`func (o *FilesystemAnalyseMetadata) SetLogHistoryNil(b bool)`

 SetLogHistoryNil sets the value for LogHistory to be an explicit nil

### UnsetLogHistory
`func (o *FilesystemAnalyseMetadata) UnsetLogHistory()`

UnsetLogHistory ensures that no value is present for LogHistory, not even an explicit nil
### GetStatus

`func (o *FilesystemAnalyseMetadata) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FilesystemAnalyseMetadata) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FilesystemAnalyseMetadata) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


