# Block

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualAddress** | **int32** |  | 

## Methods

### NewBlock

`func NewBlock(virtualAddress int32, ) *Block`

NewBlock instantiates a new Block object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockWithDefaults

`func NewBlockWithDefaults() *Block`

NewBlockWithDefaults instantiates a new Block object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualAddress

`func (o *Block) GetVirtualAddress() int32`

GetVirtualAddress returns the VirtualAddress field if non-nil, zero value otherwise.

### GetVirtualAddressOk

`func (o *Block) GetVirtualAddressOk() (*int32, bool)`

GetVirtualAddressOk returns a tuple with the VirtualAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAddress

`func (o *Block) SetVirtualAddress(v int32)`

SetVirtualAddress sets VirtualAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


