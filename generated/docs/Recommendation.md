# Recommendation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SubjectId** | **string** |  | 
**RecommendedAt** | **time.Time** |  | 
**PolicyVersion** | **string** |  | 
**Type** | **string** |  | 
**Confidence** | **float32** |  | 
**Rationale** | **string** |  | 
**PrimaryReason** | [**RecommendationReasonCode**](RecommendationReasonCode.md) |  | 
**AdditionalReasons** | Pointer to [**[]RecommendationReasonCode**](RecommendationReasonCode.md) |  | [optional] 
**Explanation** | Pointer to [**RecommendationExplanation**](RecommendationExplanation.md) |  | [optional] 
**ScoreResult** | [**ArtifactRef**](ArtifactRef.md) |  | 
**Guardrails** | Pointer to [**[]Guardrail**](Guardrail.md) |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**ReviewBy** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRecommendation

`func NewRecommendation(id string, subjectId string, recommendedAt time.Time, policyVersion string, type_ string, confidence float32, rationale string, primaryReason RecommendationReasonCode, scoreResult ArtifactRef, ) *Recommendation`

NewRecommendation instantiates a new Recommendation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationWithDefaults

`func NewRecommendationWithDefaults() *Recommendation`

NewRecommendationWithDefaults instantiates a new Recommendation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Recommendation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Recommendation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Recommendation) SetId(v string)`

SetId sets Id field to given value.


### GetSubjectId

`func (o *Recommendation) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *Recommendation) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *Recommendation) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetRecommendedAt

`func (o *Recommendation) GetRecommendedAt() time.Time`

GetRecommendedAt returns the RecommendedAt field if non-nil, zero value otherwise.

### GetRecommendedAtOk

`func (o *Recommendation) GetRecommendedAtOk() (*time.Time, bool)`

GetRecommendedAtOk returns a tuple with the RecommendedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedAt

`func (o *Recommendation) SetRecommendedAt(v time.Time)`

SetRecommendedAt sets RecommendedAt field to given value.


### GetPolicyVersion

`func (o *Recommendation) GetPolicyVersion() string`

GetPolicyVersion returns the PolicyVersion field if non-nil, zero value otherwise.

### GetPolicyVersionOk

`func (o *Recommendation) GetPolicyVersionOk() (*string, bool)`

GetPolicyVersionOk returns a tuple with the PolicyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyVersion

`func (o *Recommendation) SetPolicyVersion(v string)`

SetPolicyVersion sets PolicyVersion field to given value.


### GetType

`func (o *Recommendation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Recommendation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Recommendation) SetType(v string)`

SetType sets Type field to given value.


### GetConfidence

`func (o *Recommendation) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *Recommendation) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *Recommendation) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.


### GetRationale

`func (o *Recommendation) GetRationale() string`

GetRationale returns the Rationale field if non-nil, zero value otherwise.

### GetRationaleOk

`func (o *Recommendation) GetRationaleOk() (*string, bool)`

GetRationaleOk returns a tuple with the Rationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRationale

`func (o *Recommendation) SetRationale(v string)`

SetRationale sets Rationale field to given value.


### GetPrimaryReason

`func (o *Recommendation) GetPrimaryReason() RecommendationReasonCode`

GetPrimaryReason returns the PrimaryReason field if non-nil, zero value otherwise.

### GetPrimaryReasonOk

`func (o *Recommendation) GetPrimaryReasonOk() (*RecommendationReasonCode, bool)`

GetPrimaryReasonOk returns a tuple with the PrimaryReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryReason

`func (o *Recommendation) SetPrimaryReason(v RecommendationReasonCode)`

SetPrimaryReason sets PrimaryReason field to given value.


### GetAdditionalReasons

`func (o *Recommendation) GetAdditionalReasons() []RecommendationReasonCode`

GetAdditionalReasons returns the AdditionalReasons field if non-nil, zero value otherwise.

### GetAdditionalReasonsOk

`func (o *Recommendation) GetAdditionalReasonsOk() (*[]RecommendationReasonCode, bool)`

GetAdditionalReasonsOk returns a tuple with the AdditionalReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalReasons

`func (o *Recommendation) SetAdditionalReasons(v []RecommendationReasonCode)`

SetAdditionalReasons sets AdditionalReasons field to given value.

### HasAdditionalReasons

`func (o *Recommendation) HasAdditionalReasons() bool`

HasAdditionalReasons returns a boolean if a field has been set.

### GetExplanation

`func (o *Recommendation) GetExplanation() RecommendationExplanation`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *Recommendation) GetExplanationOk() (*RecommendationExplanation, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *Recommendation) SetExplanation(v RecommendationExplanation)`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *Recommendation) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.

### GetScoreResult

`func (o *Recommendation) GetScoreResult() ArtifactRef`

GetScoreResult returns the ScoreResult field if non-nil, zero value otherwise.

### GetScoreResultOk

`func (o *Recommendation) GetScoreResultOk() (*ArtifactRef, bool)`

GetScoreResultOk returns a tuple with the ScoreResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreResult

`func (o *Recommendation) SetScoreResult(v ArtifactRef)`

SetScoreResult sets ScoreResult field to given value.


### GetGuardrails

`func (o *Recommendation) GetGuardrails() []Guardrail`

GetGuardrails returns the Guardrails field if non-nil, zero value otherwise.

### GetGuardrailsOk

`func (o *Recommendation) GetGuardrailsOk() (*[]Guardrail, bool)`

GetGuardrailsOk returns a tuple with the Guardrails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrails

`func (o *Recommendation) SetGuardrails(v []Guardrail)`

SetGuardrails sets Guardrails field to given value.

### HasGuardrails

`func (o *Recommendation) HasGuardrails() bool`

HasGuardrails returns a boolean if a field has been set.

### GetExpiresAt

`func (o *Recommendation) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Recommendation) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Recommendation) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Recommendation) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetReviewBy

`func (o *Recommendation) GetReviewBy() time.Time`

GetReviewBy returns the ReviewBy field if non-nil, zero value otherwise.

### GetReviewByOk

`func (o *Recommendation) GetReviewByOk() (*time.Time, bool)`

GetReviewByOk returns a tuple with the ReviewBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewBy

`func (o *Recommendation) SetReviewBy(v time.Time)`

SetReviewBy sets ReviewBy field to given value.

### HasReviewBy

`func (o *Recommendation) HasReviewBy() bool`

HasReviewBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


