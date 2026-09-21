# ReferencedConstant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | [**Subject**](Subject.md) |  | 
**Value** | [**BytesConstant**](BytesConstant.md) |  | 
**Display** | Pointer to [**NullableDisplay**](Display.md) |  | [optional] 

## Methods

### NewReferencedConstant

`func NewReferencedConstant(subject Subject, value BytesConstant, ) *ReferencedConstant`

NewReferencedConstant instantiates a new ReferencedConstant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferencedConstantWithDefaults

`func NewReferencedConstantWithDefaults() *ReferencedConstant`

NewReferencedConstantWithDefaults instantiates a new ReferencedConstant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *ReferencedConstant) GetSubject() Subject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ReferencedConstant) GetSubjectOk() (*Subject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ReferencedConstant) SetSubject(v Subject)`

SetSubject sets Subject field to given value.


### GetValue

`func (o *ReferencedConstant) GetValue() BytesConstant`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ReferencedConstant) GetValueOk() (*BytesConstant, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ReferencedConstant) SetValue(v BytesConstant)`

SetValue sets Value field to given value.


### GetDisplay

`func (o *ReferencedConstant) GetDisplay() Display`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ReferencedConstant) GetDisplayOk() (*Display, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ReferencedConstant) SetDisplay(v Display)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *ReferencedConstant) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### SetDisplayNil

`func (o *ReferencedConstant) SetDisplayNil(b bool)`

 SetDisplayNil sets the value for Display to be an explicit nil

### UnsetDisplay
`func (o *ReferencedConstant) UnsetDisplay()`

UnsetDisplay ensures that no value is present for Display, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


