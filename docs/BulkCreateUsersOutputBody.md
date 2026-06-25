# BulkCreateUsersOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failed** | **int64** |  | 
**Results** | [**[]BulkCreateUserResult**](BulkCreateUserResult.md) |  | 
**Succeeded** | **int64** |  | 
**Total** | **int64** |  | 

## Methods

### NewBulkCreateUsersOutputBody

`func NewBulkCreateUsersOutputBody(failed int64, results []BulkCreateUserResult, succeeded int64, total int64, ) *BulkCreateUsersOutputBody`

NewBulkCreateUsersOutputBody instantiates a new BulkCreateUsersOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCreateUsersOutputBodyWithDefaults

`func NewBulkCreateUsersOutputBodyWithDefaults() *BulkCreateUsersOutputBody`

NewBulkCreateUsersOutputBodyWithDefaults instantiates a new BulkCreateUsersOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailed

`func (o *BulkCreateUsersOutputBody) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *BulkCreateUsersOutputBody) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *BulkCreateUsersOutputBody) SetFailed(v int64)`

SetFailed sets Failed field to given value.


### GetResults

`func (o *BulkCreateUsersOutputBody) GetResults() []BulkCreateUserResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BulkCreateUsersOutputBody) GetResultsOk() (*[]BulkCreateUserResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BulkCreateUsersOutputBody) SetResults(v []BulkCreateUserResult)`

SetResults sets Results field to given value.


### SetResultsNil

`func (o *BulkCreateUsersOutputBody) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *BulkCreateUsersOutputBody) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetSucceeded

`func (o *BulkCreateUsersOutputBody) GetSucceeded() int64`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *BulkCreateUsersOutputBody) GetSucceededOk() (*int64, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *BulkCreateUsersOutputBody) SetSucceeded(v int64)`

SetSucceeded sets Succeeded field to given value.


### GetTotal

`func (o *BulkCreateUsersOutputBody) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BulkCreateUsersOutputBody) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BulkCreateUsersOutputBody) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


