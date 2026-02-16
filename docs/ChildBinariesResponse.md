# ChildBinariesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]RelativeBinaryResponse**](RelativeBinaryResponse.md) | List of child binaries associated with the parent binary | 
**Parent** | Pointer to [**NullableRelativeBinaryResponse**](RelativeBinaryResponse.md) |  | [optional] 

## Methods

### NewChildBinariesResponse

`func NewChildBinariesResponse(children []RelativeBinaryResponse, ) *ChildBinariesResponse`

NewChildBinariesResponse instantiates a new ChildBinariesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChildBinariesResponseWithDefaults

`func NewChildBinariesResponseWithDefaults() *ChildBinariesResponse`

NewChildBinariesResponseWithDefaults instantiates a new ChildBinariesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *ChildBinariesResponse) GetChildren() []RelativeBinaryResponse`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ChildBinariesResponse) GetChildrenOk() (*[]RelativeBinaryResponse, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ChildBinariesResponse) SetChildren(v []RelativeBinaryResponse)`

SetChildren sets Children field to given value.


### GetParent

`func (o *ChildBinariesResponse) GetParent() RelativeBinaryResponse`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ChildBinariesResponse) GetParentOk() (*RelativeBinaryResponse, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ChildBinariesResponse) SetParent(v RelativeBinaryResponse)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ChildBinariesResponse) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *ChildBinariesResponse) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *ChildBinariesResponse) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


