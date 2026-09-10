# SuggestedTypeView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTypeId** | Pointer to **int64** | Existing data type the members were accessed through. Null when nothing resolved to a row; never minted for a suggestion. | [optional] 
**Holes** | [**[]SuggestedHole**](SuggestedHole.md) | Gaps between consecutive placed members, in offset order. | 
**ImpliedSize** | Pointer to **int64** | Highest byte_offset+byte_size across the members. A lower bound on the type&#39;s size, not its size. | [optional] 
**Key** | **string** | Identity of the suggestion: index:&lt;data_type_id&gt; where the access named a row, else token:&lt;type_token&gt;. | 
**Members** | [**[]SuggestedMemberView**](SuggestedMemberView.md) | Members in offset order, unplaced ones last. | 
**Name** | **string** | Name the type renders as: a database or frozen name where one exists, else the suggested one. | 
**TypeToken** | Pointer to **string** | Placeholder the type renders as, when it appears in this function&#39;s source. | [optional] 
**UnderlyingType** | Pointer to **string** | Set only for a type with no observed members, where the suggestion is a name and a scalar type rather than a layout. | [optional] 

## Methods

### NewSuggestedTypeView

`func NewSuggestedTypeView(holes []SuggestedHole, key string, members []SuggestedMemberView, name string, ) *SuggestedTypeView`

NewSuggestedTypeView instantiates a new SuggestedTypeView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestedTypeViewWithDefaults

`func NewSuggestedTypeViewWithDefaults() *SuggestedTypeView`

NewSuggestedTypeViewWithDefaults instantiates a new SuggestedTypeView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTypeId

`func (o *SuggestedTypeView) GetDataTypeId() int64`

GetDataTypeId returns the DataTypeId field if non-nil, zero value otherwise.

### GetDataTypeIdOk

`func (o *SuggestedTypeView) GetDataTypeIdOk() (*int64, bool)`

GetDataTypeIdOk returns a tuple with the DataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeId

`func (o *SuggestedTypeView) SetDataTypeId(v int64)`

SetDataTypeId sets DataTypeId field to given value.

### HasDataTypeId

`func (o *SuggestedTypeView) HasDataTypeId() bool`

HasDataTypeId returns a boolean if a field has been set.

### GetHoles

`func (o *SuggestedTypeView) GetHoles() []SuggestedHole`

GetHoles returns the Holes field if non-nil, zero value otherwise.

### GetHolesOk

`func (o *SuggestedTypeView) GetHolesOk() (*[]SuggestedHole, bool)`

GetHolesOk returns a tuple with the Holes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoles

`func (o *SuggestedTypeView) SetHoles(v []SuggestedHole)`

SetHoles sets Holes field to given value.


### SetHolesNil

`func (o *SuggestedTypeView) SetHolesNil(b bool)`

 SetHolesNil sets the value for Holes to be an explicit nil

### UnsetHoles
`func (o *SuggestedTypeView) UnsetHoles()`

UnsetHoles ensures that no value is present for Holes, not even an explicit nil
### GetImpliedSize

`func (o *SuggestedTypeView) GetImpliedSize() int64`

GetImpliedSize returns the ImpliedSize field if non-nil, zero value otherwise.

### GetImpliedSizeOk

`func (o *SuggestedTypeView) GetImpliedSizeOk() (*int64, bool)`

GetImpliedSizeOk returns a tuple with the ImpliedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpliedSize

`func (o *SuggestedTypeView) SetImpliedSize(v int64)`

SetImpliedSize sets ImpliedSize field to given value.

### HasImpliedSize

`func (o *SuggestedTypeView) HasImpliedSize() bool`

HasImpliedSize returns a boolean if a field has been set.

### GetKey

`func (o *SuggestedTypeView) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SuggestedTypeView) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SuggestedTypeView) SetKey(v string)`

SetKey sets Key field to given value.


### GetMembers

`func (o *SuggestedTypeView) GetMembers() []SuggestedMemberView`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *SuggestedTypeView) GetMembersOk() (*[]SuggestedMemberView, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *SuggestedTypeView) SetMembers(v []SuggestedMemberView)`

SetMembers sets Members field to given value.


### SetMembersNil

`func (o *SuggestedTypeView) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *SuggestedTypeView) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil
### GetName

`func (o *SuggestedTypeView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SuggestedTypeView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SuggestedTypeView) SetName(v string)`

SetName sets Name field to given value.


### GetTypeToken

`func (o *SuggestedTypeView) GetTypeToken() string`

GetTypeToken returns the TypeToken field if non-nil, zero value otherwise.

### GetTypeTokenOk

`func (o *SuggestedTypeView) GetTypeTokenOk() (*string, bool)`

GetTypeTokenOk returns a tuple with the TypeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeToken

`func (o *SuggestedTypeView) SetTypeToken(v string)`

SetTypeToken sets TypeToken field to given value.

### HasTypeToken

`func (o *SuggestedTypeView) HasTypeToken() bool`

HasTypeToken returns a boolean if a field has been set.

### GetUnderlyingType

`func (o *SuggestedTypeView) GetUnderlyingType() string`

GetUnderlyingType returns the UnderlyingType field if non-nil, zero value otherwise.

### GetUnderlyingTypeOk

`func (o *SuggestedTypeView) GetUnderlyingTypeOk() (*string, bool)`

GetUnderlyingTypeOk returns a tuple with the UnderlyingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingType

`func (o *SuggestedTypeView) SetUnderlyingType(v string)`

SetUnderlyingType sets UnderlyingType field to given value.

### HasUnderlyingType

`func (o *SuggestedTypeView) HasUnderlyingType() bool`

HasUnderlyingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


