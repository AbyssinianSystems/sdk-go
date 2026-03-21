# EvaluationRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SubjectId** | **string** |  | 
**TriggerType** | **string** |  | 
**TriggerRef** | Pointer to **string** |  | [optional] 
**RulesetVersion** | **string** |  | 
**FeatureSetVersion** | **string** |  | 
**StartedAt** | **time.Time** |  | 
**CompletedAt** | **time.Time** |  | 
**Status** | **string** |  | 
**SnapshotRef** | Pointer to [**ArtifactRef**](ArtifactRef.md) |  | [optional] 
**ScoreResultRef** | Pointer to [**ArtifactRef**](ArtifactRef.md) |  | [optional] 
**RecommendationRef** | Pointer to [**ArtifactRef**](ArtifactRef.md) |  | [optional] 
**PreviousEvaluationRef** | Pointer to [**ArtifactRef**](ArtifactRef.md) |  | [optional] 

## Methods

### NewEvaluationRun

`func NewEvaluationRun(id string, subjectId string, triggerType string, rulesetVersion string, featureSetVersion string, startedAt time.Time, completedAt time.Time, status string, ) *EvaluationRun`

NewEvaluationRun instantiates a new EvaluationRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluationRunWithDefaults

`func NewEvaluationRunWithDefaults() *EvaluationRun`

NewEvaluationRunWithDefaults instantiates a new EvaluationRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EvaluationRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EvaluationRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EvaluationRun) SetId(v string)`

SetId sets Id field to given value.


### GetSubjectId

`func (o *EvaluationRun) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *EvaluationRun) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *EvaluationRun) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetTriggerType

`func (o *EvaluationRun) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *EvaluationRun) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *EvaluationRun) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.


### GetTriggerRef

`func (o *EvaluationRun) GetTriggerRef() string`

GetTriggerRef returns the TriggerRef field if non-nil, zero value otherwise.

### GetTriggerRefOk

`func (o *EvaluationRun) GetTriggerRefOk() (*string, bool)`

GetTriggerRefOk returns a tuple with the TriggerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerRef

`func (o *EvaluationRun) SetTriggerRef(v string)`

SetTriggerRef sets TriggerRef field to given value.

### HasTriggerRef

`func (o *EvaluationRun) HasTriggerRef() bool`

HasTriggerRef returns a boolean if a field has been set.

### GetRulesetVersion

`func (o *EvaluationRun) GetRulesetVersion() string`

GetRulesetVersion returns the RulesetVersion field if non-nil, zero value otherwise.

### GetRulesetVersionOk

`func (o *EvaluationRun) GetRulesetVersionOk() (*string, bool)`

GetRulesetVersionOk returns a tuple with the RulesetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetVersion

`func (o *EvaluationRun) SetRulesetVersion(v string)`

SetRulesetVersion sets RulesetVersion field to given value.


### GetFeatureSetVersion

`func (o *EvaluationRun) GetFeatureSetVersion() string`

GetFeatureSetVersion returns the FeatureSetVersion field if non-nil, zero value otherwise.

### GetFeatureSetVersionOk

`func (o *EvaluationRun) GetFeatureSetVersionOk() (*string, bool)`

GetFeatureSetVersionOk returns a tuple with the FeatureSetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSetVersion

`func (o *EvaluationRun) SetFeatureSetVersion(v string)`

SetFeatureSetVersion sets FeatureSetVersion field to given value.


### GetStartedAt

`func (o *EvaluationRun) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *EvaluationRun) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *EvaluationRun) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetCompletedAt

`func (o *EvaluationRun) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *EvaluationRun) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *EvaluationRun) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetStatus

`func (o *EvaluationRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EvaluationRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EvaluationRun) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSnapshotRef

`func (o *EvaluationRun) GetSnapshotRef() ArtifactRef`

GetSnapshotRef returns the SnapshotRef field if non-nil, zero value otherwise.

### GetSnapshotRefOk

`func (o *EvaluationRun) GetSnapshotRefOk() (*ArtifactRef, bool)`

GetSnapshotRefOk returns a tuple with the SnapshotRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRef

`func (o *EvaluationRun) SetSnapshotRef(v ArtifactRef)`

SetSnapshotRef sets SnapshotRef field to given value.

### HasSnapshotRef

`func (o *EvaluationRun) HasSnapshotRef() bool`

HasSnapshotRef returns a boolean if a field has been set.

### GetScoreResultRef

`func (o *EvaluationRun) GetScoreResultRef() ArtifactRef`

GetScoreResultRef returns the ScoreResultRef field if non-nil, zero value otherwise.

### GetScoreResultRefOk

`func (o *EvaluationRun) GetScoreResultRefOk() (*ArtifactRef, bool)`

GetScoreResultRefOk returns a tuple with the ScoreResultRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreResultRef

`func (o *EvaluationRun) SetScoreResultRef(v ArtifactRef)`

SetScoreResultRef sets ScoreResultRef field to given value.

### HasScoreResultRef

`func (o *EvaluationRun) HasScoreResultRef() bool`

HasScoreResultRef returns a boolean if a field has been set.

### GetRecommendationRef

`func (o *EvaluationRun) GetRecommendationRef() ArtifactRef`

GetRecommendationRef returns the RecommendationRef field if non-nil, zero value otherwise.

### GetRecommendationRefOk

`func (o *EvaluationRun) GetRecommendationRefOk() (*ArtifactRef, bool)`

GetRecommendationRefOk returns a tuple with the RecommendationRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationRef

`func (o *EvaluationRun) SetRecommendationRef(v ArtifactRef)`

SetRecommendationRef sets RecommendationRef field to given value.

### HasRecommendationRef

`func (o *EvaluationRun) HasRecommendationRef() bool`

HasRecommendationRef returns a boolean if a field has been set.

### GetPreviousEvaluationRef

`func (o *EvaluationRun) GetPreviousEvaluationRef() ArtifactRef`

GetPreviousEvaluationRef returns the PreviousEvaluationRef field if non-nil, zero value otherwise.

### GetPreviousEvaluationRefOk

`func (o *EvaluationRun) GetPreviousEvaluationRefOk() (*ArtifactRef, bool)`

GetPreviousEvaluationRefOk returns a tuple with the PreviousEvaluationRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousEvaluationRef

`func (o *EvaluationRun) SetPreviousEvaluationRef(v ArtifactRef)`

SetPreviousEvaluationRef sets PreviousEvaluationRef field to given value.

### HasPreviousEvaluationRef

`func (o *EvaluationRun) HasPreviousEvaluationRef() bool`

HasPreviousEvaluationRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


