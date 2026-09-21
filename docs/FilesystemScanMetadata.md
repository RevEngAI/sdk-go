# FilesystemScanMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogHistory** | Pointer to **[][]interface{}** | Progress messages the run recorded, oldest first. | [optional] 
**Status** | **string** | Run status. UNINITIALISED means the agent has never been triggered for this analysis. | 

## Methods

### NewFilesystemScanMetadata

`func NewFilesystemScanMetadata(status string, ) *FilesystemScanMetadata`

NewFilesystemScanMetadata instantiates a new FilesystemScanMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemScanMetadataWithDefaults

`func NewFilesystemScanMetadataWithDefaults() *FilesystemScanMetadata`

NewFilesystemScanMetadataWithDefaults instantiates a new FilesystemScanMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogHistory

`func (o *FilesystemScanMetadata) GetLogHistory() [][]interface{}`

GetLogHistory returns the LogHistory field if non-nil, zero value otherwise.

### GetLogHistoryOk

`func (o *FilesystemScanMetadata) GetLogHistoryOk() (*[][]interface{}, bool)`

GetLogHistoryOk returns a tuple with the LogHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogHistory

`func (o *FilesystemScanMetadata) SetLogHistory(v [][]interface{})`

SetLogHistory sets LogHistory field to given value.

### HasLogHistory

`func (o *FilesystemScanMetadata) HasLogHistory() bool`

HasLogHistory returns a boolean if a field has been set.

### SetLogHistoryNil

`func (o *FilesystemScanMetadata) SetLogHistoryNil(b bool)`

 SetLogHistoryNil sets the value for LogHistory to be an explicit nil

### UnsetLogHistory
`func (o *FilesystemScanMetadata) UnsetLogHistory()`

UnsetLogHistory ensures that no value is present for LogHistory, not even an explicit nil
### GetStatus

`func (o *FilesystemScanMetadata) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FilesystemScanMetadata) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FilesystemScanMetadata) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


