# SoftwareTypeCountsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Benign** | **int64** | Reports classifying the software as benign | 
**Legitimate** | **int64** | Reports classifying the software as legitimate | 
**LegitimateSoftwareBackdoor** | **int64** | Reports classifying the software as a legitimate backdoor | 
**Malicious** | **int64** | Reports classifying the software as malicious | 
**PotentiallyUnwantedApplication** | **int64** | Reports classifying the software as a potentially unwanted application | 

## Methods

### NewSoftwareTypeCountsBody

`func NewSoftwareTypeCountsBody(benign int64, legitimate int64, legitimateSoftwareBackdoor int64, malicious int64, potentiallyUnwantedApplication int64, ) *SoftwareTypeCountsBody`

NewSoftwareTypeCountsBody instantiates a new SoftwareTypeCountsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareTypeCountsBodyWithDefaults

`func NewSoftwareTypeCountsBodyWithDefaults() *SoftwareTypeCountsBody`

NewSoftwareTypeCountsBodyWithDefaults instantiates a new SoftwareTypeCountsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBenign

`func (o *SoftwareTypeCountsBody) GetBenign() int64`

GetBenign returns the Benign field if non-nil, zero value otherwise.

### GetBenignOk

`func (o *SoftwareTypeCountsBody) GetBenignOk() (*int64, bool)`

GetBenignOk returns a tuple with the Benign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenign

`func (o *SoftwareTypeCountsBody) SetBenign(v int64)`

SetBenign sets Benign field to given value.


### GetLegitimate

`func (o *SoftwareTypeCountsBody) GetLegitimate() int64`

GetLegitimate returns the Legitimate field if non-nil, zero value otherwise.

### GetLegitimateOk

`func (o *SoftwareTypeCountsBody) GetLegitimateOk() (*int64, bool)`

GetLegitimateOk returns a tuple with the Legitimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegitimate

`func (o *SoftwareTypeCountsBody) SetLegitimate(v int64)`

SetLegitimate sets Legitimate field to given value.


### GetLegitimateSoftwareBackdoor

`func (o *SoftwareTypeCountsBody) GetLegitimateSoftwareBackdoor() int64`

GetLegitimateSoftwareBackdoor returns the LegitimateSoftwareBackdoor field if non-nil, zero value otherwise.

### GetLegitimateSoftwareBackdoorOk

`func (o *SoftwareTypeCountsBody) GetLegitimateSoftwareBackdoorOk() (*int64, bool)`

GetLegitimateSoftwareBackdoorOk returns a tuple with the LegitimateSoftwareBackdoor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegitimateSoftwareBackdoor

`func (o *SoftwareTypeCountsBody) SetLegitimateSoftwareBackdoor(v int64)`

SetLegitimateSoftwareBackdoor sets LegitimateSoftwareBackdoor field to given value.


### GetMalicious

`func (o *SoftwareTypeCountsBody) GetMalicious() int64`

GetMalicious returns the Malicious field if non-nil, zero value otherwise.

### GetMaliciousOk

`func (o *SoftwareTypeCountsBody) GetMaliciousOk() (*int64, bool)`

GetMaliciousOk returns a tuple with the Malicious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalicious

`func (o *SoftwareTypeCountsBody) SetMalicious(v int64)`

SetMalicious sets Malicious field to given value.


### GetPotentiallyUnwantedApplication

`func (o *SoftwareTypeCountsBody) GetPotentiallyUnwantedApplication() int64`

GetPotentiallyUnwantedApplication returns the PotentiallyUnwantedApplication field if non-nil, zero value otherwise.

### GetPotentiallyUnwantedApplicationOk

`func (o *SoftwareTypeCountsBody) GetPotentiallyUnwantedApplicationOk() (*int64, bool)`

GetPotentiallyUnwantedApplicationOk returns a tuple with the PotentiallyUnwantedApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPotentiallyUnwantedApplication

`func (o *SoftwareTypeCountsBody) SetPotentiallyUnwantedApplication(v int64)`

SetPotentiallyUnwantedApplication sets PotentiallyUnwantedApplication field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


