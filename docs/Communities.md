# Communities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalFunctions** | **int32** | The total number of matched community functions | 
**TotalMatchedFunctions** | **int32** | The total number of functions in the binary | 
**DirectCommunityMatchPercentages** | [**[]CommunityMatchPercentages**](CommunityMatchPercentages.md) | The list of directly matched communities | 
**TopComponents** | **[]map[string]interface{}** | The top components of the binary | 

## Methods

### NewCommunities

`func NewCommunities(totalFunctions int32, totalMatchedFunctions int32, directCommunityMatchPercentages []CommunityMatchPercentages, topComponents []map[string]interface{}, ) *Communities`

NewCommunities instantiates a new Communities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunitiesWithDefaults

`func NewCommunitiesWithDefaults() *Communities`

NewCommunitiesWithDefaults instantiates a new Communities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalFunctions

`func (o *Communities) GetTotalFunctions() int32`

GetTotalFunctions returns the TotalFunctions field if non-nil, zero value otherwise.

### GetTotalFunctionsOk

`func (o *Communities) GetTotalFunctionsOk() (*int32, bool)`

GetTotalFunctionsOk returns a tuple with the TotalFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFunctions

`func (o *Communities) SetTotalFunctions(v int32)`

SetTotalFunctions sets TotalFunctions field to given value.


### GetTotalMatchedFunctions

`func (o *Communities) GetTotalMatchedFunctions() int32`

GetTotalMatchedFunctions returns the TotalMatchedFunctions field if non-nil, zero value otherwise.

### GetTotalMatchedFunctionsOk

`func (o *Communities) GetTotalMatchedFunctionsOk() (*int32, bool)`

GetTotalMatchedFunctionsOk returns a tuple with the TotalMatchedFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMatchedFunctions

`func (o *Communities) SetTotalMatchedFunctions(v int32)`

SetTotalMatchedFunctions sets TotalMatchedFunctions field to given value.


### GetDirectCommunityMatchPercentages

`func (o *Communities) GetDirectCommunityMatchPercentages() []CommunityMatchPercentages`

GetDirectCommunityMatchPercentages returns the DirectCommunityMatchPercentages field if non-nil, zero value otherwise.

### GetDirectCommunityMatchPercentagesOk

`func (o *Communities) GetDirectCommunityMatchPercentagesOk() (*[]CommunityMatchPercentages, bool)`

GetDirectCommunityMatchPercentagesOk returns a tuple with the DirectCommunityMatchPercentages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectCommunityMatchPercentages

`func (o *Communities) SetDirectCommunityMatchPercentages(v []CommunityMatchPercentages)`

SetDirectCommunityMatchPercentages sets DirectCommunityMatchPercentages field to given value.


### GetTopComponents

`func (o *Communities) GetTopComponents() []map[string]interface{}`

GetTopComponents returns the TopComponents field if non-nil, zero value otherwise.

### GetTopComponentsOk

`func (o *Communities) GetTopComponentsOk() (*[]map[string]interface{}, bool)`

GetTopComponentsOk returns a tuple with the TopComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopComponents

`func (o *Communities) SetTopComponents(v []map[string]interface{})`

SetTopComponents sets TopComponents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


