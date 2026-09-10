# CryptoExplainMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogHistory** | Pointer to **[][]interface{}** | Progress messages the run recorded, oldest first. | [optional] 
**Status** | **string** | Run status. UNINITIALISED means the agent has never been triggered for this function. | 

## Methods

### NewCryptoExplainMetadata

`func NewCryptoExplainMetadata(status string, ) *CryptoExplainMetadata`

NewCryptoExplainMetadata instantiates a new CryptoExplainMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoExplainMetadataWithDefaults

`func NewCryptoExplainMetadataWithDefaults() *CryptoExplainMetadata`

NewCryptoExplainMetadataWithDefaults instantiates a new CryptoExplainMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogHistory

`func (o *CryptoExplainMetadata) GetLogHistory() [][]interface{}`

GetLogHistory returns the LogHistory field if non-nil, zero value otherwise.

### GetLogHistoryOk

`func (o *CryptoExplainMetadata) GetLogHistoryOk() (*[][]interface{}, bool)`

GetLogHistoryOk returns a tuple with the LogHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogHistory

`func (o *CryptoExplainMetadata) SetLogHistory(v [][]interface{})`

SetLogHistory sets LogHistory field to given value.

### HasLogHistory

`func (o *CryptoExplainMetadata) HasLogHistory() bool`

HasLogHistory returns a boolean if a field has been set.

### SetLogHistoryNil

`func (o *CryptoExplainMetadata) SetLogHistoryNil(b bool)`

 SetLogHistoryNil sets the value for LogHistory to be an explicit nil

### UnsetLogHistory
`func (o *CryptoExplainMetadata) UnsetLogHistory()`

UnsetLogHistory ensures that no value is present for LogHistory, not even an explicit nil
### GetStatus

`func (o *CryptoExplainMetadata) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CryptoExplainMetadata) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CryptoExplainMetadata) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


