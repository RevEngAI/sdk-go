# DieMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Display** | **string** | Human-readable description from DIE; suitable for display, not parsing | 
**Name** | **string** | Canonical name of the matched signature or technology | 
**Type** | **string** | Category DIE assigns the match, such as compiler, packer or file | 
**Version** | **string** | Version DIE extracted, empty when it could not determine one | 

## Methods

### NewDieMatch

`func NewDieMatch(display string, name string, type_ string, version string, ) *DieMatch`

NewDieMatch instantiates a new DieMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDieMatchWithDefaults

`func NewDieMatchWithDefaults() *DieMatch`

NewDieMatchWithDefaults instantiates a new DieMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplay

`func (o *DieMatch) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *DieMatch) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *DieMatch) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *DieMatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DieMatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DieMatch) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *DieMatch) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DieMatch) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DieMatch) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *DieMatch) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DieMatch) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DieMatch) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


