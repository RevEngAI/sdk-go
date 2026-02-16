# InverseFunctionMapItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Addr** | [**NullableAddr**](Addr.md) |  | 
**IsExternal** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewInverseFunctionMapItem

`func NewInverseFunctionMapItem(name string, addr NullableAddr, ) *InverseFunctionMapItem`

NewInverseFunctionMapItem instantiates a new InverseFunctionMapItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInverseFunctionMapItemWithDefaults

`func NewInverseFunctionMapItemWithDefaults() *InverseFunctionMapItem`

NewInverseFunctionMapItemWithDefaults instantiates a new InverseFunctionMapItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InverseFunctionMapItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InverseFunctionMapItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InverseFunctionMapItem) SetName(v string)`

SetName sets Name field to given value.


### GetAddr

`func (o *InverseFunctionMapItem) GetAddr() Addr`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *InverseFunctionMapItem) GetAddrOk() (*Addr, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *InverseFunctionMapItem) SetAddr(v Addr)`

SetAddr sets Addr field to given value.


### SetAddrNil

`func (o *InverseFunctionMapItem) SetAddrNil(b bool)`

 SetAddrNil sets the value for Addr to be an explicit nil

### UnsetAddr
`func (o *InverseFunctionMapItem) UnsetAddr()`

UnsetAddr ensures that no value is present for Addr, not even an explicit nil
### GetIsExternal

`func (o *InverseFunctionMapItem) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *InverseFunctionMapItem) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *InverseFunctionMapItem) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *InverseFunctionMapItem) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


