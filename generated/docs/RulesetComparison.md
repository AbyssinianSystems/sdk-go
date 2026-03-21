# RulesetComparison

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectId** | **string** |  | 
**CandidateRulesetVersion** | **string** |  | 
**Baseline** | [**SubjectEvaluationBundle**](SubjectEvaluationBundle.md) |  | 
**Candidate** | [**SubjectEvaluationBundle**](SubjectEvaluationBundle.md) |  | 
**ScoreDelta** | **float32** |  | 
**RecommendationChanged** | **bool** |  | 

## Methods

### NewRulesetComparison

`func NewRulesetComparison(subjectId string, candidateRulesetVersion string, baseline SubjectEvaluationBundle, candidate SubjectEvaluationBundle, scoreDelta float32, recommendationChanged bool, ) *RulesetComparison`

NewRulesetComparison instantiates a new RulesetComparison object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesetComparisonWithDefaults

`func NewRulesetComparisonWithDefaults() *RulesetComparison`

NewRulesetComparisonWithDefaults instantiates a new RulesetComparison object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectId

`func (o *RulesetComparison) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *RulesetComparison) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *RulesetComparison) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetCandidateRulesetVersion

`func (o *RulesetComparison) GetCandidateRulesetVersion() string`

GetCandidateRulesetVersion returns the CandidateRulesetVersion field if non-nil, zero value otherwise.

### GetCandidateRulesetVersionOk

`func (o *RulesetComparison) GetCandidateRulesetVersionOk() (*string, bool)`

GetCandidateRulesetVersionOk returns a tuple with the CandidateRulesetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidateRulesetVersion

`func (o *RulesetComparison) SetCandidateRulesetVersion(v string)`

SetCandidateRulesetVersion sets CandidateRulesetVersion field to given value.


### GetBaseline

`func (o *RulesetComparison) GetBaseline() SubjectEvaluationBundle`

GetBaseline returns the Baseline field if non-nil, zero value otherwise.

### GetBaselineOk

`func (o *RulesetComparison) GetBaselineOk() (*SubjectEvaluationBundle, bool)`

GetBaselineOk returns a tuple with the Baseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseline

`func (o *RulesetComparison) SetBaseline(v SubjectEvaluationBundle)`

SetBaseline sets Baseline field to given value.


### GetCandidate

`func (o *RulesetComparison) GetCandidate() SubjectEvaluationBundle`

GetCandidate returns the Candidate field if non-nil, zero value otherwise.

### GetCandidateOk

`func (o *RulesetComparison) GetCandidateOk() (*SubjectEvaluationBundle, bool)`

GetCandidateOk returns a tuple with the Candidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidate

`func (o *RulesetComparison) SetCandidate(v SubjectEvaluationBundle)`

SetCandidate sets Candidate field to given value.


### GetScoreDelta

`func (o *RulesetComparison) GetScoreDelta() float32`

GetScoreDelta returns the ScoreDelta field if non-nil, zero value otherwise.

### GetScoreDeltaOk

`func (o *RulesetComparison) GetScoreDeltaOk() (*float32, bool)`

GetScoreDeltaOk returns a tuple with the ScoreDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreDelta

`func (o *RulesetComparison) SetScoreDelta(v float32)`

SetScoreDelta sets ScoreDelta field to given value.


### GetRecommendationChanged

`func (o *RulesetComparison) GetRecommendationChanged() bool`

GetRecommendationChanged returns the RecommendationChanged field if non-nil, zero value otherwise.

### GetRecommendationChangedOk

`func (o *RulesetComparison) GetRecommendationChangedOk() (*bool, bool)`

GetRecommendationChangedOk returns a tuple with the RecommendationChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationChanged

`func (o *RulesetComparison) SetRecommendationChanged(v bool)`

SetRecommendationChanged sets RecommendationChanged field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


