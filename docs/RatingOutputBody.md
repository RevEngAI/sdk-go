# RatingOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rating** | **NullableString** | The rating the caller recorded, or null when they have not rated it yet | 
**Reason** | **NullableString** | The reason the caller gave, or null | 

## Methods

### NewRatingOutputBody

`func NewRatingOutputBody(rating NullableString, reason NullableString, ) *RatingOutputBody`

NewRatingOutputBody instantiates a new RatingOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatingOutputBodyWithDefaults

`func NewRatingOutputBodyWithDefaults() *RatingOutputBody`

NewRatingOutputBodyWithDefaults instantiates a new RatingOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRating

`func (o *RatingOutputBody) GetRating() string`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *RatingOutputBody) GetRatingOk() (*string, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *RatingOutputBody) SetRating(v string)`

SetRating sets Rating field to given value.


### SetRatingNil

`func (o *RatingOutputBody) SetRatingNil(b bool)`

 SetRatingNil sets the value for Rating to be an explicit nil

### UnsetRating
`func (o *RatingOutputBody) UnsetRating()`

UnsetRating ensures that no value is present for Rating, not even an explicit nil
### GetReason

`func (o *RatingOutputBody) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RatingOutputBody) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RatingOutputBody) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *RatingOutputBody) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *RatingOutputBody) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


