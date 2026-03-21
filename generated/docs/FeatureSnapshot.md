# FeatureSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SubjectId** | **string** |  | 
**FeatureSetVersion** | **string** |  | 
**EvaluatedAt** | **time.Time** |  | 
**Windows** | Pointer to [**[]FeatureWindow**](FeatureWindow.md) |  | [optional] 
**Values** | Pointer to [**[]FeatureValue**](FeatureValue.md) |  | [optional] 
**Summaries** | Pointer to [**[]FeatureValue**](FeatureValue.md) |  | [optional] 
**Provenance** | Pointer to [**[]ArtifactRef**](ArtifactRef.md) |  | [optional] 

## Methods

### NewFeatureSnapshot

`func NewFeatureSnapshot(id string, subjectId string, featureSetVersion string, evaluatedAt time.Time, ) *FeatureSnapshot`

NewFeatureSnapshot instantiates a new FeatureSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureSnapshotWithDefaults

`func NewFeatureSnapshotWithDefaults() *FeatureSnapshot`

NewFeatureSnapshotWithDefaults instantiates a new FeatureSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FeatureSnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeatureSnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeatureSnapshot) SetId(v string)`

SetId sets Id field to given value.


### GetSubjectId

`func (o *FeatureSnapshot) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *FeatureSnapshot) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *FeatureSnapshot) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetFeatureSetVersion

`func (o *FeatureSnapshot) GetFeatureSetVersion() string`

GetFeatureSetVersion returns the FeatureSetVersion field if non-nil, zero value otherwise.

### GetFeatureSetVersionOk

`func (o *FeatureSnapshot) GetFeatureSetVersionOk() (*string, bool)`

GetFeatureSetVersionOk returns a tuple with the FeatureSetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSetVersion

`func (o *FeatureSnapshot) SetFeatureSetVersion(v string)`

SetFeatureSetVersion sets FeatureSetVersion field to given value.


### GetEvaluatedAt

`func (o *FeatureSnapshot) GetEvaluatedAt() time.Time`

GetEvaluatedAt returns the EvaluatedAt field if non-nil, zero value otherwise.

### GetEvaluatedAtOk

`func (o *FeatureSnapshot) GetEvaluatedAtOk() (*time.Time, bool)`

GetEvaluatedAtOk returns a tuple with the EvaluatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedAt

`func (o *FeatureSnapshot) SetEvaluatedAt(v time.Time)`

SetEvaluatedAt sets EvaluatedAt field to given value.


### GetWindows

`func (o *FeatureSnapshot) GetWindows() []FeatureWindow`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *FeatureSnapshot) GetWindowsOk() (*[]FeatureWindow, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *FeatureSnapshot) SetWindows(v []FeatureWindow)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *FeatureSnapshot) HasWindows() bool`

HasWindows returns a boolean if a field has been set.

### GetValues

`func (o *FeatureSnapshot) GetValues() []FeatureValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *FeatureSnapshot) GetValuesOk() (*[]FeatureValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *FeatureSnapshot) SetValues(v []FeatureValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *FeatureSnapshot) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetSummaries

`func (o *FeatureSnapshot) GetSummaries() []FeatureValue`

GetSummaries returns the Summaries field if non-nil, zero value otherwise.

### GetSummariesOk

`func (o *FeatureSnapshot) GetSummariesOk() (*[]FeatureValue, bool)`

GetSummariesOk returns a tuple with the Summaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaries

`func (o *FeatureSnapshot) SetSummaries(v []FeatureValue)`

SetSummaries sets Summaries field to given value.

### HasSummaries

`func (o *FeatureSnapshot) HasSummaries() bool`

HasSummaries returns a boolean if a field has been set.

### GetProvenance

`func (o *FeatureSnapshot) GetProvenance() []ArtifactRef`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *FeatureSnapshot) GetProvenanceOk() (*[]ArtifactRef, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *FeatureSnapshot) SetProvenance(v []ArtifactRef)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *FeatureSnapshot) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


