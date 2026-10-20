# GlobalHistoryEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Previous** | [**GlobalVariable**](GlobalVariable.md) | The global as it stood before that edit. | 
**ReplacedAt** | **time.Time** | When the value below was replaced. | 
**ReplacedByUserId** | **int64** | Who replaced it. | 

## Methods

### NewGlobalHistoryEntry

`func NewGlobalHistoryEntry(previous GlobalVariable, replacedAt time.Time, replacedByUserId int64, ) *GlobalHistoryEntry`

NewGlobalHistoryEntry instantiates a new GlobalHistoryEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalHistoryEntryWithDefaults

`func NewGlobalHistoryEntryWithDefaults() *GlobalHistoryEntry`

NewGlobalHistoryEntryWithDefaults instantiates a new GlobalHistoryEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrevious

`func (o *GlobalHistoryEntry) GetPrevious() GlobalVariable`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *GlobalHistoryEntry) GetPreviousOk() (*GlobalVariable, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *GlobalHistoryEntry) SetPrevious(v GlobalVariable)`

SetPrevious sets Previous field to given value.


### GetReplacedAt

`func (o *GlobalHistoryEntry) GetReplacedAt() time.Time`

GetReplacedAt returns the ReplacedAt field if non-nil, zero value otherwise.

### GetReplacedAtOk

`func (o *GlobalHistoryEntry) GetReplacedAtOk() (*time.Time, bool)`

GetReplacedAtOk returns a tuple with the ReplacedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedAt

`func (o *GlobalHistoryEntry) SetReplacedAt(v time.Time)`

SetReplacedAt sets ReplacedAt field to given value.


### GetReplacedByUserId

`func (o *GlobalHistoryEntry) GetReplacedByUserId() int64`

GetReplacedByUserId returns the ReplacedByUserId field if non-nil, zero value otherwise.

### GetReplacedByUserIdOk

`func (o *GlobalHistoryEntry) GetReplacedByUserIdOk() (*int64, bool)`

GetReplacedByUserIdOk returns a tuple with the ReplacedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedByUserId

`func (o *GlobalHistoryEntry) SetReplacedByUserId(v int64)`

SetReplacedByUserId sets ReplacedByUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


