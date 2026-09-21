# SuspiciousString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Subject** | Pointer to [**NullableSubject**](Subject.md) |  | [optional] 

## Methods

### NewSuspiciousString

`func NewSuspiciousString(value string, ) *SuspiciousString`

NewSuspiciousString instantiates a new SuspiciousString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuspiciousStringWithDefaults

`func NewSuspiciousStringWithDefaults() *SuspiciousString`

NewSuspiciousStringWithDefaults instantiates a new SuspiciousString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *SuspiciousString) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SuspiciousString) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SuspiciousString) SetValue(v string)`

SetValue sets Value field to given value.


### GetSubject

`func (o *SuspiciousString) GetSubject() Subject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SuspiciousString) GetSubjectOk() (*Subject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SuspiciousString) SetSubject(v Subject)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SuspiciousString) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *SuspiciousString) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *SuspiciousString) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


