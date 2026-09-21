# Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogHistory** | **[][]interface{}** | Progress messages the run recorded, oldest first, empty until the run logs anything. Each entry is a [timestamp, message] pair. | 
**Status** | **string** | Run status. UNINITIALISED means this agent has never been triggered for the analysis, so it is safe to start one. | 

## Methods

### NewMetadata

`func NewMetadata(logHistory [][]interface{}, status string, ) *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogHistory

`func (o *Metadata) GetLogHistory() [][]interface{}`

GetLogHistory returns the LogHistory field if non-nil, zero value otherwise.

### GetLogHistoryOk

`func (o *Metadata) GetLogHistoryOk() (*[][]interface{}, bool)`

GetLogHistoryOk returns a tuple with the LogHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogHistory

`func (o *Metadata) SetLogHistory(v [][]interface{})`

SetLogHistory sets LogHistory field to given value.


### SetLogHistoryNil

`func (o *Metadata) SetLogHistoryNil(b bool)`

 SetLogHistoryNil sets the value for LogHistory to be an explicit nil

### UnsetLogHistory
`func (o *Metadata) UnsetLogHistory()`

UnsetLogHistory ensures that no value is present for LogHistory, not even an explicit nil
### GetStatus

`func (o *Metadata) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Metadata) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Metadata) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


