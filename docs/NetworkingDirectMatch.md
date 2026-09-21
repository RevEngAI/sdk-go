# NetworkingDirectMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | Networking category of the match | 
**How** | **string** | Detection tier that produced the match | 
**MatchedName** | **string** | Name or token that matched | 
**Source** | **string** | Networking source the match belongs to | 

## Methods

### NewNetworkingDirectMatch

`func NewNetworkingDirectMatch(category string, how string, matchedName string, source string, ) *NetworkingDirectMatch`

NewNetworkingDirectMatch instantiates a new NetworkingDirectMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingDirectMatchWithDefaults

`func NewNetworkingDirectMatchWithDefaults() *NetworkingDirectMatch`

NewNetworkingDirectMatchWithDefaults instantiates a new NetworkingDirectMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *NetworkingDirectMatch) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NetworkingDirectMatch) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NetworkingDirectMatch) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetHow

`func (o *NetworkingDirectMatch) GetHow() string`

GetHow returns the How field if non-nil, zero value otherwise.

### GetHowOk

`func (o *NetworkingDirectMatch) GetHowOk() (*string, bool)`

GetHowOk returns a tuple with the How field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHow

`func (o *NetworkingDirectMatch) SetHow(v string)`

SetHow sets How field to given value.


### GetMatchedName

`func (o *NetworkingDirectMatch) GetMatchedName() string`

GetMatchedName returns the MatchedName field if non-nil, zero value otherwise.

### GetMatchedNameOk

`func (o *NetworkingDirectMatch) GetMatchedNameOk() (*string, bool)`

GetMatchedNameOk returns a tuple with the MatchedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedName

`func (o *NetworkingDirectMatch) SetMatchedName(v string)`

SetMatchedName sets MatchedName field to given value.


### GetSource

`func (o *NetworkingDirectMatch) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NetworkingDirectMatch) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NetworkingDirectMatch) SetSource(v string)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


