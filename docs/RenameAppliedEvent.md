# RenameAppliedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddrHex** | Pointer to **string** |  | [optional] 
**Attempt** | **int32** |  | 
**NewName** | **string** |  | 
**OldName** | **string** |  | 
**Seq** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewRenameAppliedEvent

`func NewRenameAppliedEvent(attempt int32, newName string, oldName string, seq int32, type_ string, ) *RenameAppliedEvent`

NewRenameAppliedEvent instantiates a new RenameAppliedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenameAppliedEventWithDefaults

`func NewRenameAppliedEventWithDefaults() *RenameAppliedEvent`

NewRenameAppliedEventWithDefaults instantiates a new RenameAppliedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddrHex

`func (o *RenameAppliedEvent) GetAddrHex() string`

GetAddrHex returns the AddrHex field if non-nil, zero value otherwise.

### GetAddrHexOk

`func (o *RenameAppliedEvent) GetAddrHexOk() (*string, bool)`

GetAddrHexOk returns a tuple with the AddrHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrHex

`func (o *RenameAppliedEvent) SetAddrHex(v string)`

SetAddrHex sets AddrHex field to given value.

### HasAddrHex

`func (o *RenameAppliedEvent) HasAddrHex() bool`

HasAddrHex returns a boolean if a field has been set.

### GetAttempt

`func (o *RenameAppliedEvent) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *RenameAppliedEvent) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *RenameAppliedEvent) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.


### GetNewName

`func (o *RenameAppliedEvent) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *RenameAppliedEvent) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *RenameAppliedEvent) SetNewName(v string)`

SetNewName sets NewName field to given value.


### GetOldName

`func (o *RenameAppliedEvent) GetOldName() string`

GetOldName returns the OldName field if non-nil, zero value otherwise.

### GetOldNameOk

`func (o *RenameAppliedEvent) GetOldNameOk() (*string, bool)`

GetOldNameOk returns a tuple with the OldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldName

`func (o *RenameAppliedEvent) SetOldName(v string)`

SetOldName sets OldName field to given value.


### GetSeq

`func (o *RenameAppliedEvent) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *RenameAppliedEvent) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *RenameAppliedEvent) SetSeq(v int32)`

SetSeq sets Seq field to given value.


### GetType

`func (o *RenameAppliedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RenameAppliedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RenameAppliedEvent) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


