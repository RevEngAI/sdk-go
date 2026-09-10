# CryptoCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalleeName** | **string** | Name of the called function | 
**Category** | **string** | Crypto category of the match | 
**How** | **string** | Detection tier that produced the match | 
**Library** | **string** | Crypto library the match belongs to | 
**MatchedName** | **string** | Name or token that matched | 

## Methods

### NewCryptoCall

`func NewCryptoCall(calleeName string, category string, how string, library string, matchedName string, ) *CryptoCall`

NewCryptoCall instantiates a new CryptoCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoCallWithDefaults

`func NewCryptoCallWithDefaults() *CryptoCall`

NewCryptoCallWithDefaults instantiates a new CryptoCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalleeName

`func (o *CryptoCall) GetCalleeName() string`

GetCalleeName returns the CalleeName field if non-nil, zero value otherwise.

### GetCalleeNameOk

`func (o *CryptoCall) GetCalleeNameOk() (*string, bool)`

GetCalleeNameOk returns a tuple with the CalleeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalleeName

`func (o *CryptoCall) SetCalleeName(v string)`

SetCalleeName sets CalleeName field to given value.


### GetCategory

`func (o *CryptoCall) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CryptoCall) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CryptoCall) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetHow

`func (o *CryptoCall) GetHow() string`

GetHow returns the How field if non-nil, zero value otherwise.

### GetHowOk

`func (o *CryptoCall) GetHowOk() (*string, bool)`

GetHowOk returns a tuple with the How field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHow

`func (o *CryptoCall) SetHow(v string)`

SetHow sets How field to given value.


### GetLibrary

`func (o *CryptoCall) GetLibrary() string`

GetLibrary returns the Library field if non-nil, zero value otherwise.

### GetLibraryOk

`func (o *CryptoCall) GetLibraryOk() (*string, bool)`

GetLibraryOk returns a tuple with the Library field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibrary

`func (o *CryptoCall) SetLibrary(v string)`

SetLibrary sets Library field to given value.


### GetMatchedName

`func (o *CryptoCall) GetMatchedName() string`

GetMatchedName returns the MatchedName field if non-nil, zero value otherwise.

### GetMatchedNameOk

`func (o *CryptoCall) GetMatchedNameOk() (*string, bool)`

GetMatchedNameOk returns a tuple with the MatchedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedName

`func (o *CryptoCall) SetMatchedName(v string)`

SetMatchedName sets MatchedName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


