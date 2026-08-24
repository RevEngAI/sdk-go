# UpsertOverridesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlaceholderToUserOverride** | [**map[string]Token**](Token.md) | Every override on the function after applying this request, keyed by placeholder token. | 

## Methods

### NewUpsertOverridesData

`func NewUpsertOverridesData(placeholderToUserOverride map[string]Token, ) *UpsertOverridesData`

NewUpsertOverridesData instantiates a new UpsertOverridesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertOverridesDataWithDefaults

`func NewUpsertOverridesDataWithDefaults() *UpsertOverridesData`

NewUpsertOverridesDataWithDefaults instantiates a new UpsertOverridesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaceholderToUserOverride

`func (o *UpsertOverridesData) GetPlaceholderToUserOverride() map[string]Token`

GetPlaceholderToUserOverride returns the PlaceholderToUserOverride field if non-nil, zero value otherwise.

### GetPlaceholderToUserOverrideOk

`func (o *UpsertOverridesData) GetPlaceholderToUserOverrideOk() (*map[string]Token, bool)`

GetPlaceholderToUserOverrideOk returns a tuple with the PlaceholderToUserOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholderToUserOverride

`func (o *UpsertOverridesData) SetPlaceholderToUserOverride(v map[string]Token)`

SetPlaceholderToUserOverride sets PlaceholderToUserOverride field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


