# FunctionBlockResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asm** | **[]string** | The ordered assembly strings for this chunk | 
**Id** | **int32** | ID of the block | 
**MinAddr** | **int32** | The minimum vaddr of the block | 
**MaxAddr** | **int32** | The maximum vaddr of the block | 
**Destinations** | [**[]FunctionBlockDestinationResponse**](FunctionBlockDestinationResponse.md) | The potential execution flow destinations from this block | 
**Comment** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFunctionBlockResponse

`func NewFunctionBlockResponse(asm []string, id int32, minAddr int32, maxAddr int32, destinations []FunctionBlockDestinationResponse, ) *FunctionBlockResponse`

NewFunctionBlockResponse instantiates a new FunctionBlockResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionBlockResponseWithDefaults

`func NewFunctionBlockResponseWithDefaults() *FunctionBlockResponse`

NewFunctionBlockResponseWithDefaults instantiates a new FunctionBlockResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsm

`func (o *FunctionBlockResponse) GetAsm() []string`

GetAsm returns the Asm field if non-nil, zero value otherwise.

### GetAsmOk

`func (o *FunctionBlockResponse) GetAsmOk() (*[]string, bool)`

GetAsmOk returns a tuple with the Asm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsm

`func (o *FunctionBlockResponse) SetAsm(v []string)`

SetAsm sets Asm field to given value.


### GetId

`func (o *FunctionBlockResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FunctionBlockResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FunctionBlockResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetMinAddr

`func (o *FunctionBlockResponse) GetMinAddr() int32`

GetMinAddr returns the MinAddr field if non-nil, zero value otherwise.

### GetMinAddrOk

`func (o *FunctionBlockResponse) GetMinAddrOk() (*int32, bool)`

GetMinAddrOk returns a tuple with the MinAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAddr

`func (o *FunctionBlockResponse) SetMinAddr(v int32)`

SetMinAddr sets MinAddr field to given value.


### GetMaxAddr

`func (o *FunctionBlockResponse) GetMaxAddr() int32`

GetMaxAddr returns the MaxAddr field if non-nil, zero value otherwise.

### GetMaxAddrOk

`func (o *FunctionBlockResponse) GetMaxAddrOk() (*int32, bool)`

GetMaxAddrOk returns a tuple with the MaxAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAddr

`func (o *FunctionBlockResponse) SetMaxAddr(v int32)`

SetMaxAddr sets MaxAddr field to given value.


### GetDestinations

`func (o *FunctionBlockResponse) GetDestinations() []FunctionBlockDestinationResponse`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *FunctionBlockResponse) GetDestinationsOk() (*[]FunctionBlockDestinationResponse, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *FunctionBlockResponse) SetDestinations(v []FunctionBlockDestinationResponse)`

SetDestinations sets Destinations field to given value.


### GetComment

`func (o *FunctionBlockResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FunctionBlockResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FunctionBlockResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FunctionBlockResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *FunctionBlockResponse) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *FunctionBlockResponse) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


