# PrincipalBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The user or team&#39;s ID | 
**Name** | **string** | The user&#39;s username or the team&#39;s name | 
**Type** | **string** | Whether this result is a user or a team | 

## Methods

### NewPrincipalBody

`func NewPrincipalBody(id int64, name string, type_ string, ) *PrincipalBody`

NewPrincipalBody instantiates a new PrincipalBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalBodyWithDefaults

`func NewPrincipalBodyWithDefaults() *PrincipalBody`

NewPrincipalBodyWithDefaults instantiates a new PrincipalBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrincipalBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrincipalBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrincipalBody) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *PrincipalBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrincipalBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrincipalBody) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *PrincipalBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrincipalBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrincipalBody) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


