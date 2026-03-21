# ScoreResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SubjectId** | **string** |  | 
**RulesetVersion** | **string** |  | 
**EvaluatedAt** | **time.Time** |  | 
**Snapshot** | [**ArtifactRef**](ArtifactRef.md) |  | 
**Score** | **float32** |  | 
**Band** | **string** |  | 
**Contributions** | Pointer to [**[]ScoreContribution**](ScoreContribution.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Explanation** | Pointer to [**ScoreExplanation**](ScoreExplanation.md) |  | [optional] 
**Confidence** | Pointer to **float32** |  | [optional] 

## Methods

### NewScoreResult

`func NewScoreResult(id string, subjectId string, rulesetVersion string, evaluatedAt time.Time, snapshot ArtifactRef, score float32, band string, ) *ScoreResult`

NewScoreResult instantiates a new ScoreResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScoreResultWithDefaults

`func NewScoreResultWithDefaults() *ScoreResult`

NewScoreResultWithDefaults instantiates a new ScoreResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScoreResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScoreResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScoreResult) SetId(v string)`

SetId sets Id field to given value.


### GetSubjectId

`func (o *ScoreResult) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *ScoreResult) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *ScoreResult) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetRulesetVersion

`func (o *ScoreResult) GetRulesetVersion() string`

GetRulesetVersion returns the RulesetVersion field if non-nil, zero value otherwise.

### GetRulesetVersionOk

`func (o *ScoreResult) GetRulesetVersionOk() (*string, bool)`

GetRulesetVersionOk returns a tuple with the RulesetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetVersion

`func (o *ScoreResult) SetRulesetVersion(v string)`

SetRulesetVersion sets RulesetVersion field to given value.


### GetEvaluatedAt

`func (o *ScoreResult) GetEvaluatedAt() time.Time`

GetEvaluatedAt returns the EvaluatedAt field if non-nil, zero value otherwise.

### GetEvaluatedAtOk

`func (o *ScoreResult) GetEvaluatedAtOk() (*time.Time, bool)`

GetEvaluatedAtOk returns a tuple with the EvaluatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedAt

`func (o *ScoreResult) SetEvaluatedAt(v time.Time)`

SetEvaluatedAt sets EvaluatedAt field to given value.


### GetSnapshot

`func (o *ScoreResult) GetSnapshot() ArtifactRef`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *ScoreResult) GetSnapshotOk() (*ArtifactRef, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *ScoreResult) SetSnapshot(v ArtifactRef)`

SetSnapshot sets Snapshot field to given value.


### GetScore

`func (o *ScoreResult) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ScoreResult) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ScoreResult) SetScore(v float32)`

SetScore sets Score field to given value.


### GetBand

`func (o *ScoreResult) GetBand() string`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *ScoreResult) GetBandOk() (*string, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *ScoreResult) SetBand(v string)`

SetBand sets Band field to given value.


### GetContributions

`func (o *ScoreResult) GetContributions() []ScoreContribution`

GetContributions returns the Contributions field if non-nil, zero value otherwise.

### GetContributionsOk

`func (o *ScoreResult) GetContributionsOk() (*[]ScoreContribution, bool)`

GetContributionsOk returns a tuple with the Contributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributions

`func (o *ScoreResult) SetContributions(v []ScoreContribution)`

SetContributions sets Contributions field to given value.

### HasContributions

`func (o *ScoreResult) HasContributions() bool`

HasContributions returns a boolean if a field has been set.

### GetSummary

`func (o *ScoreResult) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ScoreResult) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ScoreResult) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ScoreResult) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetExplanation

`func (o *ScoreResult) GetExplanation() ScoreExplanation`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *ScoreResult) GetExplanationOk() (*ScoreExplanation, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *ScoreResult) SetExplanation(v ScoreExplanation)`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *ScoreResult) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.

### GetConfidence

`func (o *ScoreResult) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *ScoreResult) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *ScoreResult) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *ScoreResult) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


