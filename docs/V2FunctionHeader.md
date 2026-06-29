# V2FunctionHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastChange** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** | Name of the function | 
**Addr** | **int32** | Memory address of the function | 
**Type** | **string** | Return type of the function | 
**Args** | [**map[string]Argument**](Argument.md) | Dictionary of function arguments | 

## Methods

### NewV2FunctionHeader

`func NewV2FunctionHeader(name string, addr int32, type_ string, args map[string]Argument, ) *V2FunctionHeader`

NewV2FunctionHeader instantiates a new V2FunctionHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2FunctionHeaderWithDefaults

`func NewV2FunctionHeaderWithDefaults() *V2FunctionHeader`

NewV2FunctionHeaderWithDefaults instantiates a new V2FunctionHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastChange

`func (o *V2FunctionHeader) GetLastChange() string`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *V2FunctionHeader) GetLastChangeOk() (*string, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *V2FunctionHeader) SetLastChange(v string)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *V2FunctionHeader) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### SetLastChangeNil

`func (o *V2FunctionHeader) SetLastChangeNil(b bool)`

 SetLastChangeNil sets the value for LastChange to be an explicit nil

### UnsetLastChange
`func (o *V2FunctionHeader) UnsetLastChange()`

UnsetLastChange ensures that no value is present for LastChange, not even an explicit nil
### GetName

`func (o *V2FunctionHeader) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2FunctionHeader) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2FunctionHeader) SetName(v string)`

SetName sets Name field to given value.


### GetAddr

`func (o *V2FunctionHeader) GetAddr() int32`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *V2FunctionHeader) GetAddrOk() (*int32, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *V2FunctionHeader) SetAddr(v int32)`

SetAddr sets Addr field to given value.


### GetType

`func (o *V2FunctionHeader) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V2FunctionHeader) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V2FunctionHeader) SetType(v string)`

SetType sets Type field to given value.


### GetArgs

`func (o *V2FunctionHeader) GetArgs() map[string]Argument`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *V2FunctionHeader) GetArgsOk() (*map[string]Argument, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *V2FunctionHeader) SetArgs(v map[string]Argument)`

SetArgs sets Args field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


