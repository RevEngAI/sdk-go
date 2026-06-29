# FunctionHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addr** | **int64** |  | 
**Args** | [**map[string]FunctionArgument**](FunctionArgument.md) | Argument map keyed by ordinal hex (e.g. \&quot;0x0\&quot;, \&quot;0x1\&quot;). | 
**LastChange** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Scope** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewFunctionHeader

`func NewFunctionHeader(addr int64, args map[string]FunctionArgument, name string, type_ string, ) *FunctionHeader`

NewFunctionHeader instantiates a new FunctionHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionHeaderWithDefaults

`func NewFunctionHeaderWithDefaults() *FunctionHeader`

NewFunctionHeaderWithDefaults instantiates a new FunctionHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddr

`func (o *FunctionHeader) GetAddr() int64`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *FunctionHeader) GetAddrOk() (*int64, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *FunctionHeader) SetAddr(v int64)`

SetAddr sets Addr field to given value.


### GetArgs

`func (o *FunctionHeader) GetArgs() map[string]FunctionArgument`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *FunctionHeader) GetArgsOk() (*map[string]FunctionArgument, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *FunctionHeader) SetArgs(v map[string]FunctionArgument)`

SetArgs sets Args field to given value.


### GetLastChange

`func (o *FunctionHeader) GetLastChange() string`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *FunctionHeader) GetLastChangeOk() (*string, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *FunctionHeader) SetLastChange(v string)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *FunctionHeader) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### GetName

`func (o *FunctionHeader) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionHeader) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionHeader) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *FunctionHeader) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *FunctionHeader) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *FunctionHeader) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *FunctionHeader) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetType

`func (o *FunctionHeader) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionHeader) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionHeader) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


