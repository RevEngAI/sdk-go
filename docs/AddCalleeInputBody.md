# AddCalleeInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalleeFunctionId** | Pointer to **int64** | Internal callee function ID; 0 means not provided | [optional] 
**CalleeName** | Pointer to **string** | Callee name (for external calls) | [optional] 
**CalleeVaddr** | **int64** | Virtual address of the callee | 
**IsExternal** | **bool** | Whether the callee is outside the binary | 
**ThunkedVaddr** | Pointer to **int64** | Thunked virtual address | [optional] 

## Methods

### NewAddCalleeInputBody

`func NewAddCalleeInputBody(calleeVaddr int64, isExternal bool, ) *AddCalleeInputBody`

NewAddCalleeInputBody instantiates a new AddCalleeInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCalleeInputBodyWithDefaults

`func NewAddCalleeInputBodyWithDefaults() *AddCalleeInputBody`

NewAddCalleeInputBodyWithDefaults instantiates a new AddCalleeInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalleeFunctionId

`func (o *AddCalleeInputBody) GetCalleeFunctionId() int64`

GetCalleeFunctionId returns the CalleeFunctionId field if non-nil, zero value otherwise.

### GetCalleeFunctionIdOk

`func (o *AddCalleeInputBody) GetCalleeFunctionIdOk() (*int64, bool)`

GetCalleeFunctionIdOk returns a tuple with the CalleeFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalleeFunctionId

`func (o *AddCalleeInputBody) SetCalleeFunctionId(v int64)`

SetCalleeFunctionId sets CalleeFunctionId field to given value.

### HasCalleeFunctionId

`func (o *AddCalleeInputBody) HasCalleeFunctionId() bool`

HasCalleeFunctionId returns a boolean if a field has been set.

### GetCalleeName

`func (o *AddCalleeInputBody) GetCalleeName() string`

GetCalleeName returns the CalleeName field if non-nil, zero value otherwise.

### GetCalleeNameOk

`func (o *AddCalleeInputBody) GetCalleeNameOk() (*string, bool)`

GetCalleeNameOk returns a tuple with the CalleeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalleeName

`func (o *AddCalleeInputBody) SetCalleeName(v string)`

SetCalleeName sets CalleeName field to given value.

### HasCalleeName

`func (o *AddCalleeInputBody) HasCalleeName() bool`

HasCalleeName returns a boolean if a field has been set.

### GetCalleeVaddr

`func (o *AddCalleeInputBody) GetCalleeVaddr() int64`

GetCalleeVaddr returns the CalleeVaddr field if non-nil, zero value otherwise.

### GetCalleeVaddrOk

`func (o *AddCalleeInputBody) GetCalleeVaddrOk() (*int64, bool)`

GetCalleeVaddrOk returns a tuple with the CalleeVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalleeVaddr

`func (o *AddCalleeInputBody) SetCalleeVaddr(v int64)`

SetCalleeVaddr sets CalleeVaddr field to given value.


### GetIsExternal

`func (o *AddCalleeInputBody) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *AddCalleeInputBody) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *AddCalleeInputBody) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.


### GetThunkedVaddr

`func (o *AddCalleeInputBody) GetThunkedVaddr() int64`

GetThunkedVaddr returns the ThunkedVaddr field if non-nil, zero value otherwise.

### GetThunkedVaddrOk

`func (o *AddCalleeInputBody) GetThunkedVaddrOk() (*int64, bool)`

GetThunkedVaddrOk returns a tuple with the ThunkedVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThunkedVaddr

`func (o *AddCalleeInputBody) SetThunkedVaddr(v int64)`

SetThunkedVaddr sets ThunkedVaddr field to given value.

### HasThunkedVaddr

`func (o *AddCalleeInputBody) HasThunkedVaddr() bool`

HasThunkedVaddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


