# SBOM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Packages** | [**[]SBOMPackage**](SBOMPackage.md) | The packages found | 
**ImportedLibs** | **[]string** | The import libraries found | 

## Methods

### NewSBOM

`func NewSBOM(packages []SBOMPackage, importedLibs []string, ) *SBOM`

NewSBOM instantiates a new SBOM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSBOMWithDefaults

`func NewSBOMWithDefaults() *SBOM`

NewSBOMWithDefaults instantiates a new SBOM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackages

`func (o *SBOM) GetPackages() []SBOMPackage`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *SBOM) GetPackagesOk() (*[]SBOMPackage, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *SBOM) SetPackages(v []SBOMPackage)`

SetPackages sets Packages field to given value.


### GetImportedLibs

`func (o *SBOM) GetImportedLibs() []string`

GetImportedLibs returns the ImportedLibs field if non-nil, zero value otherwise.

### GetImportedLibsOk

`func (o *SBOM) GetImportedLibsOk() (*[]string, bool)`

GetImportedLibsOk returns a tuple with the ImportedLibs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedLibs

`func (o *SBOM) SetImportedLibs(v []string)`

SetImportedLibs sets ImportedLibs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


