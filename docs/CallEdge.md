# CallEdge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalleeFunctionId** | Pointer to **int64** |  | [optional] 
**CalleeName** | Pointer to **string** |  | [optional] 
**CalleeVaddr** | **int64** |  | 
**CallerFunctionId** | **int64** |  | 
**CallerName** | Pointer to **string** |  | [optional] 
**CallerVaddr** | **int64** | Entry vaddr of the caller function (joined from function_t). | 
**IsExternal** | **bool** |  | 
**ThunkedVaddr** | Pointer to **int64** |  | [optional] 

## Methods

### NewCallEdge

`func NewCallEdge(calleeVaddr int64, callerFunctionId int64, callerVaddr int64, isExternal bool, ) *CallEdge`

NewCallEdge instantiates a new CallEdge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallEdgeWithDefaults

`func NewCallEdgeWithDefaults() *CallEdge`

NewCallEdgeWithDefaults instantiates a new CallEdge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalleeFunctionId

`func (o *CallEdge) GetCalleeFunctionId() int64`

GetCalleeFunctionId returns the CalleeFunctionId field if non-nil, zero value otherwise.

### GetCalleeFunctionIdOk

`func (o *CallEdge) GetCalleeFunctionIdOk() (*int64, bool)`

GetCalleeFunctionIdOk returns a tuple with the CalleeFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalleeFunctionId

`func (o *CallEdge) SetCalleeFunctionId(v int64)`

SetCalleeFunctionId sets CalleeFunctionId field to given value.

### HasCalleeFunctionId

`func (o *CallEdge) HasCalleeFunctionId() bool`

HasCalleeFunctionId returns a boolean if a field has been set.

### GetCalleeName

`func (o *CallEdge) GetCalleeName() string`

GetCalleeName returns the CalleeName field if non-nil, zero value otherwise.

### GetCalleeNameOk

`func (o *CallEdge) GetCalleeNameOk() (*string, bool)`

GetCalleeNameOk returns a tuple with the CalleeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalleeName

`func (o *CallEdge) SetCalleeName(v string)`

SetCalleeName sets CalleeName field to given value.

### HasCalleeName

`func (o *CallEdge) HasCalleeName() bool`

HasCalleeName returns a boolean if a field has been set.

### GetCalleeVaddr

`func (o *CallEdge) GetCalleeVaddr() int64`

GetCalleeVaddr returns the CalleeVaddr field if non-nil, zero value otherwise.

### GetCalleeVaddrOk

`func (o *CallEdge) GetCalleeVaddrOk() (*int64, bool)`

GetCalleeVaddrOk returns a tuple with the CalleeVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalleeVaddr

`func (o *CallEdge) SetCalleeVaddr(v int64)`

SetCalleeVaddr sets CalleeVaddr field to given value.


### GetCallerFunctionId

`func (o *CallEdge) GetCallerFunctionId() int64`

GetCallerFunctionId returns the CallerFunctionId field if non-nil, zero value otherwise.

### GetCallerFunctionIdOk

`func (o *CallEdge) GetCallerFunctionIdOk() (*int64, bool)`

GetCallerFunctionIdOk returns a tuple with the CallerFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerFunctionId

`func (o *CallEdge) SetCallerFunctionId(v int64)`

SetCallerFunctionId sets CallerFunctionId field to given value.


### GetCallerName

`func (o *CallEdge) GetCallerName() string`

GetCallerName returns the CallerName field if non-nil, zero value otherwise.

### GetCallerNameOk

`func (o *CallEdge) GetCallerNameOk() (*string, bool)`

GetCallerNameOk returns a tuple with the CallerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerName

`func (o *CallEdge) SetCallerName(v string)`

SetCallerName sets CallerName field to given value.

### HasCallerName

`func (o *CallEdge) HasCallerName() bool`

HasCallerName returns a boolean if a field has been set.

### GetCallerVaddr

`func (o *CallEdge) GetCallerVaddr() int64`

GetCallerVaddr returns the CallerVaddr field if non-nil, zero value otherwise.

### GetCallerVaddrOk

`func (o *CallEdge) GetCallerVaddrOk() (*int64, bool)`

GetCallerVaddrOk returns a tuple with the CallerVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerVaddr

`func (o *CallEdge) SetCallerVaddr(v int64)`

SetCallerVaddr sets CallerVaddr field to given value.


### GetIsExternal

`func (o *CallEdge) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *CallEdge) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *CallEdge) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.


### GetThunkedVaddr

`func (o *CallEdge) GetThunkedVaddr() int64`

GetThunkedVaddr returns the ThunkedVaddr field if non-nil, zero value otherwise.

### GetThunkedVaddrOk

`func (o *CallEdge) GetThunkedVaddrOk() (*int64, bool)`

GetThunkedVaddrOk returns a tuple with the ThunkedVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThunkedVaddr

`func (o *CallEdge) SetThunkedVaddr(v int64)`

SetThunkedVaddr sets ThunkedVaddr field to given value.

### HasThunkedVaddr

`func (o *CallEdge) HasThunkedVaddr() bool`

HasThunkedVaddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


