# FilesystemCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalleeName** | **string** | Name of the called function | 
**Category** | **string** | Filesystem category of the match | 
**How** | **string** | Detection tier that produced the match | 
**MatchedName** | **string** | Name or token that matched | 
**Source** | **string** | Filesystem source the match belongs to | 

## Methods

### NewFilesystemCall

`func NewFilesystemCall(calleeName string, category string, how string, matchedName string, source string, ) *FilesystemCall`

NewFilesystemCall instantiates a new FilesystemCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemCallWithDefaults

`func NewFilesystemCallWithDefaults() *FilesystemCall`

NewFilesystemCallWithDefaults instantiates a new FilesystemCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalleeName

`func (o *FilesystemCall) GetCalleeName() string`

GetCalleeName returns the CalleeName field if non-nil, zero value otherwise.

### GetCalleeNameOk

`func (o *FilesystemCall) GetCalleeNameOk() (*string, bool)`

GetCalleeNameOk returns a tuple with the CalleeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalleeName

`func (o *FilesystemCall) SetCalleeName(v string)`

SetCalleeName sets CalleeName field to given value.


### GetCategory

`func (o *FilesystemCall) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *FilesystemCall) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *FilesystemCall) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetHow

`func (o *FilesystemCall) GetHow() string`

GetHow returns the How field if non-nil, zero value otherwise.

### GetHowOk

`func (o *FilesystemCall) GetHowOk() (*string, bool)`

GetHowOk returns a tuple with the How field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHow

`func (o *FilesystemCall) SetHow(v string)`

SetHow sets How field to given value.


### GetMatchedName

`func (o *FilesystemCall) GetMatchedName() string`

GetMatchedName returns the MatchedName field if non-nil, zero value otherwise.

### GetMatchedNameOk

`func (o *FilesystemCall) GetMatchedNameOk() (*string, bool)`

GetMatchedNameOk returns a tuple with the MatchedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedName

`func (o *FilesystemCall) SetMatchedName(v string)`

SetMatchedName sets MatchedName field to given value.


### GetSource

`func (o *FilesystemCall) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FilesystemCall) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FilesystemCall) SetSource(v string)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


