# BatchBinaryMatchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryId** | **int64** | Target binary | 
**ErrorMessage** | Pointer to **string** | Error description when status&#x3D;FAILED. | [optional] 
**MatchId** | Pointer to **string** | Opaque token for this binary&#39;s matching run. Present on dispatch and when statuses were fetched by token. | [optional] 
**MatchedFunctionCount** | **int64** | Number of source functions that received at least one candidate match. Only meaningful when status&#x3D;COMPLETED. | 
**Status** | **string** | Per-binary workflow status | 

## Methods

### NewBatchBinaryMatchResult

`func NewBatchBinaryMatchResult(binaryId int64, matchedFunctionCount int64, status string, ) *BatchBinaryMatchResult`

NewBatchBinaryMatchResult instantiates a new BatchBinaryMatchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchBinaryMatchResultWithDefaults

`func NewBatchBinaryMatchResultWithDefaults() *BatchBinaryMatchResult`

NewBatchBinaryMatchResultWithDefaults instantiates a new BatchBinaryMatchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryId

`func (o *BatchBinaryMatchResult) GetBinaryId() int64`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *BatchBinaryMatchResult) GetBinaryIdOk() (*int64, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *BatchBinaryMatchResult) SetBinaryId(v int64)`

SetBinaryId sets BinaryId field to given value.


### GetErrorMessage

`func (o *BatchBinaryMatchResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BatchBinaryMatchResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BatchBinaryMatchResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BatchBinaryMatchResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetMatchId

`func (o *BatchBinaryMatchResult) GetMatchId() string`

GetMatchId returns the MatchId field if non-nil, zero value otherwise.

### GetMatchIdOk

`func (o *BatchBinaryMatchResult) GetMatchIdOk() (*string, bool)`

GetMatchIdOk returns a tuple with the MatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchId

`func (o *BatchBinaryMatchResult) SetMatchId(v string)`

SetMatchId sets MatchId field to given value.

### HasMatchId

`func (o *BatchBinaryMatchResult) HasMatchId() bool`

HasMatchId returns a boolean if a field has been set.

### GetMatchedFunctionCount

`func (o *BatchBinaryMatchResult) GetMatchedFunctionCount() int64`

GetMatchedFunctionCount returns the MatchedFunctionCount field if non-nil, zero value otherwise.

### GetMatchedFunctionCountOk

`func (o *BatchBinaryMatchResult) GetMatchedFunctionCountOk() (*int64, bool)`

GetMatchedFunctionCountOk returns a tuple with the MatchedFunctionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedFunctionCount

`func (o *BatchBinaryMatchResult) SetMatchedFunctionCount(v int64)`

SetMatchedFunctionCount sets MatchedFunctionCount field to given value.


### GetStatus

`func (o *BatchBinaryMatchResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchBinaryMatchResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchBinaryMatchResult) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


