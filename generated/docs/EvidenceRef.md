# EvidenceRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifact** | [**ArtifactRef**](ArtifactRef.md) |  | 
**Field** | **string** |  | 

## Methods

### NewEvidenceRef

`func NewEvidenceRef(artifact ArtifactRef, field string, ) *EvidenceRef`

NewEvidenceRef instantiates a new EvidenceRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvidenceRefWithDefaults

`func NewEvidenceRefWithDefaults() *EvidenceRef`

NewEvidenceRefWithDefaults instantiates a new EvidenceRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifact

`func (o *EvidenceRef) GetArtifact() ArtifactRef`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *EvidenceRef) GetArtifactOk() (*ArtifactRef, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *EvidenceRef) SetArtifact(v ArtifactRef)`

SetArtifact sets Artifact field to given value.


### GetField

`func (o *EvidenceRef) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *EvidenceRef) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *EvidenceRef) SetField(v string)`

SetField sets Field field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


