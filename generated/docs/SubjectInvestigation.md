# SubjectInvestigation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectId** | **string** |  | 
**LatestEvaluation** | Pointer to [**SubjectEvaluationBundle**](SubjectEvaluationBundle.md) |  | [optional] 
**SignalEvents** | Pointer to [**[]SignalEvent**](SignalEvent.md) |  | [optional] 
**FeatureSnapshots** | Pointer to [**[]FeatureSnapshot**](FeatureSnapshot.md) |  | [optional] 
**ScoreResults** | Pointer to [**[]ScoreResult**](ScoreResult.md) |  | [optional] 
**Recommendations** | Pointer to [**[]Recommendation**](Recommendation.md) |  | [optional] 
**EvaluationRuns** | Pointer to [**[]EvaluationRun**](EvaluationRun.md) |  | [optional] 
**ReviewOutcomes** | Pointer to [**[]ReviewOutcome**](ReviewOutcome.md) |  | [optional] 
**OutcomeComparisons** | Pointer to [**[]OutcomeComparison**](OutcomeComparison.md) |  | [optional] 

## Methods

### NewSubjectInvestigation

`func NewSubjectInvestigation(subjectId string, ) *SubjectInvestigation`

NewSubjectInvestigation instantiates a new SubjectInvestigation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectInvestigationWithDefaults

`func NewSubjectInvestigationWithDefaults() *SubjectInvestigation`

NewSubjectInvestigationWithDefaults instantiates a new SubjectInvestigation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectId

`func (o *SubjectInvestigation) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *SubjectInvestigation) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *SubjectInvestigation) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetLatestEvaluation

`func (o *SubjectInvestigation) GetLatestEvaluation() SubjectEvaluationBundle`

GetLatestEvaluation returns the LatestEvaluation field if non-nil, zero value otherwise.

### GetLatestEvaluationOk

`func (o *SubjectInvestigation) GetLatestEvaluationOk() (*SubjectEvaluationBundle, bool)`

GetLatestEvaluationOk returns a tuple with the LatestEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestEvaluation

`func (o *SubjectInvestigation) SetLatestEvaluation(v SubjectEvaluationBundle)`

SetLatestEvaluation sets LatestEvaluation field to given value.

### HasLatestEvaluation

`func (o *SubjectInvestigation) HasLatestEvaluation() bool`

HasLatestEvaluation returns a boolean if a field has been set.

### GetSignalEvents

`func (o *SubjectInvestigation) GetSignalEvents() []SignalEvent`

GetSignalEvents returns the SignalEvents field if non-nil, zero value otherwise.

### GetSignalEventsOk

`func (o *SubjectInvestigation) GetSignalEventsOk() (*[]SignalEvent, bool)`

GetSignalEventsOk returns a tuple with the SignalEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalEvents

`func (o *SubjectInvestigation) SetSignalEvents(v []SignalEvent)`

SetSignalEvents sets SignalEvents field to given value.

### HasSignalEvents

`func (o *SubjectInvestigation) HasSignalEvents() bool`

HasSignalEvents returns a boolean if a field has been set.

### GetFeatureSnapshots

`func (o *SubjectInvestigation) GetFeatureSnapshots() []FeatureSnapshot`

GetFeatureSnapshots returns the FeatureSnapshots field if non-nil, zero value otherwise.

### GetFeatureSnapshotsOk

`func (o *SubjectInvestigation) GetFeatureSnapshotsOk() (*[]FeatureSnapshot, bool)`

GetFeatureSnapshotsOk returns a tuple with the FeatureSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSnapshots

`func (o *SubjectInvestigation) SetFeatureSnapshots(v []FeatureSnapshot)`

SetFeatureSnapshots sets FeatureSnapshots field to given value.

### HasFeatureSnapshots

`func (o *SubjectInvestigation) HasFeatureSnapshots() bool`

HasFeatureSnapshots returns a boolean if a field has been set.

### GetScoreResults

`func (o *SubjectInvestigation) GetScoreResults() []ScoreResult`

GetScoreResults returns the ScoreResults field if non-nil, zero value otherwise.

### GetScoreResultsOk

`func (o *SubjectInvestigation) GetScoreResultsOk() (*[]ScoreResult, bool)`

GetScoreResultsOk returns a tuple with the ScoreResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreResults

`func (o *SubjectInvestigation) SetScoreResults(v []ScoreResult)`

SetScoreResults sets ScoreResults field to given value.

### HasScoreResults

`func (o *SubjectInvestigation) HasScoreResults() bool`

HasScoreResults returns a boolean if a field has been set.

### GetRecommendations

`func (o *SubjectInvestigation) GetRecommendations() []Recommendation`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *SubjectInvestigation) GetRecommendationsOk() (*[]Recommendation, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *SubjectInvestigation) SetRecommendations(v []Recommendation)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *SubjectInvestigation) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.

### GetEvaluationRuns

`func (o *SubjectInvestigation) GetEvaluationRuns() []EvaluationRun`

GetEvaluationRuns returns the EvaluationRuns field if non-nil, zero value otherwise.

### GetEvaluationRunsOk

`func (o *SubjectInvestigation) GetEvaluationRunsOk() (*[]EvaluationRun, bool)`

GetEvaluationRunsOk returns a tuple with the EvaluationRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationRuns

`func (o *SubjectInvestigation) SetEvaluationRuns(v []EvaluationRun)`

SetEvaluationRuns sets EvaluationRuns field to given value.

### HasEvaluationRuns

`func (o *SubjectInvestigation) HasEvaluationRuns() bool`

HasEvaluationRuns returns a boolean if a field has been set.

### GetReviewOutcomes

`func (o *SubjectInvestigation) GetReviewOutcomes() []ReviewOutcome`

GetReviewOutcomes returns the ReviewOutcomes field if non-nil, zero value otherwise.

### GetReviewOutcomesOk

`func (o *SubjectInvestigation) GetReviewOutcomesOk() (*[]ReviewOutcome, bool)`

GetReviewOutcomesOk returns a tuple with the ReviewOutcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewOutcomes

`func (o *SubjectInvestigation) SetReviewOutcomes(v []ReviewOutcome)`

SetReviewOutcomes sets ReviewOutcomes field to given value.

### HasReviewOutcomes

`func (o *SubjectInvestigation) HasReviewOutcomes() bool`

HasReviewOutcomes returns a boolean if a field has been set.

### GetOutcomeComparisons

`func (o *SubjectInvestigation) GetOutcomeComparisons() []OutcomeComparison`

GetOutcomeComparisons returns the OutcomeComparisons field if non-nil, zero value otherwise.

### GetOutcomeComparisonsOk

`func (o *SubjectInvestigation) GetOutcomeComparisonsOk() (*[]OutcomeComparison, bool)`

GetOutcomeComparisonsOk returns a tuple with the OutcomeComparisons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeComparisons

`func (o *SubjectInvestigation) SetOutcomeComparisons(v []OutcomeComparison)`

SetOutcomeComparisons sets OutcomeComparisons field to given value.

### HasOutcomeComparisons

`func (o *SubjectInvestigation) HasOutcomeComparisons() bool`

HasOutcomeComparisons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


