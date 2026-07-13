# IndirectCallSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalleeFunctionId** | Pointer to **int64** |  | [optional] 
**InstVaddr** | **int64** | Vaddr of the indirect call instruction. | 
**IsExternal** | **bool** |  | 
**TargetName** | Pointer to **string** |  | [optional] 
**TargetVaddr** | **int64** | Resolved call target vaddr. | 

## Methods

### NewIndirectCallSite

`func NewIndirectCallSite(instVaddr int64, isExternal bool, targetVaddr int64, ) *IndirectCallSite`

NewIndirectCallSite instantiates a new IndirectCallSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndirectCallSiteWithDefaults

`func NewIndirectCallSiteWithDefaults() *IndirectCallSite`

NewIndirectCallSiteWithDefaults instantiates a new IndirectCallSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalleeFunctionId

`func (o *IndirectCallSite) GetCalleeFunctionId() int64`

GetCalleeFunctionId returns the CalleeFunctionId field if non-nil, zero value otherwise.

### GetCalleeFunctionIdOk

`func (o *IndirectCallSite) GetCalleeFunctionIdOk() (*int64, bool)`

GetCalleeFunctionIdOk returns a tuple with the CalleeFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalleeFunctionId

`func (o *IndirectCallSite) SetCalleeFunctionId(v int64)`

SetCalleeFunctionId sets CalleeFunctionId field to given value.

### HasCalleeFunctionId

`func (o *IndirectCallSite) HasCalleeFunctionId() bool`

HasCalleeFunctionId returns a boolean if a field has been set.

### GetInstVaddr

`func (o *IndirectCallSite) GetInstVaddr() int64`

GetInstVaddr returns the InstVaddr field if non-nil, zero value otherwise.

### GetInstVaddrOk

`func (o *IndirectCallSite) GetInstVaddrOk() (*int64, bool)`

GetInstVaddrOk returns a tuple with the InstVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstVaddr

`func (o *IndirectCallSite) SetInstVaddr(v int64)`

SetInstVaddr sets InstVaddr field to given value.


### GetIsExternal

`func (o *IndirectCallSite) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *IndirectCallSite) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *IndirectCallSite) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.


### GetTargetName

`func (o *IndirectCallSite) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *IndirectCallSite) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *IndirectCallSite) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *IndirectCallSite) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetTargetVaddr

`func (o *IndirectCallSite) GetTargetVaddr() int64`

GetTargetVaddr returns the TargetVaddr field if non-nil, zero value otherwise.

### GetTargetVaddrOk

`func (o *IndirectCallSite) GetTargetVaddrOk() (*int64, bool)`

GetTargetVaddrOk returns a tuple with the TargetVaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVaddr

`func (o *IndirectCallSite) SetTargetVaddr(v int64)`

SetTargetVaddr sets TargetVaddr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


