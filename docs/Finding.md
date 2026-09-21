# Finding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyId** | Pointer to **string** |  | [optional] [default to ""]
**Reachability** | Pointer to [**ReportReachabilityStatus**](ReportReachabilityStatus.md) |  | [optional] [default to REPORTREACHABILITYSTATUS_UNDETERMINED]
**Evidence** | Pointer to [**[]EvidenceInner**](EvidenceInner.md) |  | [optional] 

## Methods

### NewFinding

`func NewFinding() *Finding`

NewFinding instantiates a new Finding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingWithDefaults

`func NewFindingWithDefaults() *Finding`

NewFindingWithDefaults instantiates a new Finding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyId

`func (o *Finding) GetPropertyId() string`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *Finding) GetPropertyIdOk() (*string, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *Finding) SetPropertyId(v string)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *Finding) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetReachability

`func (o *Finding) GetReachability() ReportReachabilityStatus`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *Finding) GetReachabilityOk() (*ReportReachabilityStatus, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *Finding) SetReachability(v ReportReachabilityStatus)`

SetReachability sets Reachability field to given value.

### HasReachability

`func (o *Finding) HasReachability() bool`

HasReachability returns a boolean if a field has been set.

### GetEvidence

`func (o *Finding) GetEvidence() []EvidenceInner`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *Finding) GetEvidenceOk() (*[]EvidenceInner, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *Finding) SetEvidence(v []EvidenceInner)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *Finding) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


