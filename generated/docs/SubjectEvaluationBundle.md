# SubjectEvaluationBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectId** | **string** |  | 
**EvaluationRun** | Pointer to [**EvaluationRun**](EvaluationRun.md) |  | [optional] 
**SignalEvents** | [**[]SignalEvent**](SignalEvent.md) |  | 
**Snapshot** | [**FeatureSnapshot**](FeatureSnapshot.md) |  | 
**ScoreResult** | [**ScoreResult**](ScoreResult.md) |  | 
**Recommendation** | [**Recommendation**](Recommendation.md) |  | 

## Methods

### NewSubjectEvaluationBundle

`func NewSubjectEvaluationBundle(subjectId string, signalEvents []SignalEvent, snapshot FeatureSnapshot, scoreResult ScoreResult, recommendation Recommendation, ) *SubjectEvaluationBundle`

NewSubjectEvaluationBundle instantiates a new SubjectEvaluationBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectEvaluationBundleWithDefaults

`func NewSubjectEvaluationBundleWithDefaults() *SubjectEvaluationBundle`

NewSubjectEvaluationBundleWithDefaults instantiates a new SubjectEvaluationBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectId

`func (o *SubjectEvaluationBundle) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *SubjectEvaluationBundle) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *SubjectEvaluationBundle) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetEvaluationRun

`func (o *SubjectEvaluationBundle) GetEvaluationRun() EvaluationRun`

GetEvaluationRun returns the EvaluationRun field if non-nil, zero value otherwise.

### GetEvaluationRunOk

`func (o *SubjectEvaluationBundle) GetEvaluationRunOk() (*EvaluationRun, bool)`

GetEvaluationRunOk returns a tuple with the EvaluationRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationRun

`func (o *SubjectEvaluationBundle) SetEvaluationRun(v EvaluationRun)`

SetEvaluationRun sets EvaluationRun field to given value.

### HasEvaluationRun

`func (o *SubjectEvaluationBundle) HasEvaluationRun() bool`

HasEvaluationRun returns a boolean if a field has been set.

### GetSignalEvents

`func (o *SubjectEvaluationBundle) GetSignalEvents() []SignalEvent`

GetSignalEvents returns the SignalEvents field if non-nil, zero value otherwise.

### GetSignalEventsOk

`func (o *SubjectEvaluationBundle) GetSignalEventsOk() (*[]SignalEvent, bool)`

GetSignalEventsOk returns a tuple with the SignalEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalEvents

`func (o *SubjectEvaluationBundle) SetSignalEvents(v []SignalEvent)`

SetSignalEvents sets SignalEvents field to given value.


### GetSnapshot

`func (o *SubjectEvaluationBundle) GetSnapshot() FeatureSnapshot`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *SubjectEvaluationBundle) GetSnapshotOk() (*FeatureSnapshot, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *SubjectEvaluationBundle) SetSnapshot(v FeatureSnapshot)`

SetSnapshot sets Snapshot field to given value.


### GetScoreResult

`func (o *SubjectEvaluationBundle) GetScoreResult() ScoreResult`

GetScoreResult returns the ScoreResult field if non-nil, zero value otherwise.

### GetScoreResultOk

`func (o *SubjectEvaluationBundle) GetScoreResultOk() (*ScoreResult, bool)`

GetScoreResultOk returns a tuple with the ScoreResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreResult

`func (o *SubjectEvaluationBundle) SetScoreResult(v ScoreResult)`

SetScoreResult sets ScoreResult field to given value.


### GetRecommendation

`func (o *SubjectEvaluationBundle) GetRecommendation() Recommendation`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *SubjectEvaluationBundle) GetRecommendationOk() (*Recommendation, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *SubjectEvaluationBundle) SetRecommendation(v Recommendation)`

SetRecommendation sets Recommendation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


