# NetworkingCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalleeName** | **string** | Name of the called function | 
**Category** | **string** | Networking category of the match | 
**How** | **string** | Detection tier that produced the match | 
**MatchedName** | **string** | Name or token that matched | 
**Source** | **string** | Networking source the match belongs to | 

## Methods

### NewNetworkingCall

`func NewNetworkingCall(calleeName string, category string, how string, matchedName string, source string, ) *NetworkingCall`

NewNetworkingCall instantiates a new NetworkingCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingCallWithDefaults

`func NewNetworkingCallWithDefaults() *NetworkingCall`

NewNetworkingCallWithDefaults instantiates a new NetworkingCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalleeName

`func (o *NetworkingCall) GetCalleeName() string`

GetCalleeName returns the CalleeName field if non-nil, zero value otherwise.

### GetCalleeNameOk

`func (o *NetworkingCall) GetCalleeNameOk() (*string, bool)`

GetCalleeNameOk returns a tuple with the CalleeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalleeName

`func (o *NetworkingCall) SetCalleeName(v string)`

SetCalleeName sets CalleeName field to given value.


### GetCategory

`func (o *NetworkingCall) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NetworkingCall) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NetworkingCall) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetHow

`func (o *NetworkingCall) GetHow() string`

GetHow returns the How field if non-nil, zero value otherwise.

### GetHowOk

`func (o *NetworkingCall) GetHowOk() (*string, bool)`

GetHowOk returns a tuple with the How field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHow

`func (o *NetworkingCall) SetHow(v string)`

SetHow sets How field to given value.


### GetMatchedName

`func (o *NetworkingCall) GetMatchedName() string`

GetMatchedName returns the MatchedName field if non-nil, zero value otherwise.

### GetMatchedNameOk

`func (o *NetworkingCall) GetMatchedNameOk() (*string, bool)`

GetMatchedNameOk returns a tuple with the MatchedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedName

`func (o *NetworkingCall) SetMatchedName(v string)`

SetMatchedName sets MatchedName field to given value.


### GetSource

`func (o *NetworkingCall) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NetworkingCall) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NetworkingCall) SetSource(v string)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


