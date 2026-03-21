# RecommendationExplanation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | **string** |  | 
**RecommendationReason** | **string** |  | 
**PrimaryReason** | [**RecommendationReasonCode**](RecommendationReasonCode.md) |  | 
**AdditionalReasons** | Pointer to [**[]RecommendationReasonCode**](RecommendationReasonCode.md) |  | [optional] 
**Score** | **float32** |  | 
**Band** | **string** |  | 
**Rules** | Pointer to [**[]ExplanationRule**](ExplanationRule.md) |  | [optional] 
**Features** | Pointer to [**[]ExplanationFeature**](ExplanationFeature.md) |  | [optional] 
**Guardrails** | Pointer to [**[]Guardrail**](Guardrail.md) |  | [optional] 

## Methods

### NewRecommendationExplanation

`func NewRecommendationExplanation(summary string, recommendationReason string, primaryReason RecommendationReasonCode, score float32, band string, ) *RecommendationExplanation`

NewRecommendationExplanation instantiates a new RecommendationExplanation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationExplanationWithDefaults

`func NewRecommendationExplanationWithDefaults() *RecommendationExplanation`

NewRecommendationExplanationWithDefaults instantiates a new RecommendationExplanation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *RecommendationExplanation) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *RecommendationExplanation) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *RecommendationExplanation) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetRecommendationReason

`func (o *RecommendationExplanation) GetRecommendationReason() string`

GetRecommendationReason returns the RecommendationReason field if non-nil, zero value otherwise.

### GetRecommendationReasonOk

`func (o *RecommendationExplanation) GetRecommendationReasonOk() (*string, bool)`

GetRecommendationReasonOk returns a tuple with the RecommendationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationReason

`func (o *RecommendationExplanation) SetRecommendationReason(v string)`

SetRecommendationReason sets RecommendationReason field to given value.


### GetPrimaryReason

`func (o *RecommendationExplanation) GetPrimaryReason() RecommendationReasonCode`

GetPrimaryReason returns the PrimaryReason field if non-nil, zero value otherwise.

### GetPrimaryReasonOk

`func (o *RecommendationExplanation) GetPrimaryReasonOk() (*RecommendationReasonCode, bool)`

GetPrimaryReasonOk returns a tuple with the PrimaryReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryReason

`func (o *RecommendationExplanation) SetPrimaryReason(v RecommendationReasonCode)`

SetPrimaryReason sets PrimaryReason field to given value.


### GetAdditionalReasons

`func (o *RecommendationExplanation) GetAdditionalReasons() []RecommendationReasonCode`

GetAdditionalReasons returns the AdditionalReasons field if non-nil, zero value otherwise.

### GetAdditionalReasonsOk

`func (o *RecommendationExplanation) GetAdditionalReasonsOk() (*[]RecommendationReasonCode, bool)`

GetAdditionalReasonsOk returns a tuple with the AdditionalReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalReasons

`func (o *RecommendationExplanation) SetAdditionalReasons(v []RecommendationReasonCode)`

SetAdditionalReasons sets AdditionalReasons field to given value.

### HasAdditionalReasons

`func (o *RecommendationExplanation) HasAdditionalReasons() bool`

HasAdditionalReasons returns a boolean if a field has been set.

### GetScore

`func (o *RecommendationExplanation) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *RecommendationExplanation) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *RecommendationExplanation) SetScore(v float32)`

SetScore sets Score field to given value.


### GetBand

`func (o *RecommendationExplanation) GetBand() string`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *RecommendationExplanation) GetBandOk() (*string, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *RecommendationExplanation) SetBand(v string)`

SetBand sets Band field to given value.


### GetRules

`func (o *RecommendationExplanation) GetRules() []ExplanationRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *RecommendationExplanation) GetRulesOk() (*[]ExplanationRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *RecommendationExplanation) SetRules(v []ExplanationRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *RecommendationExplanation) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetFeatures

`func (o *RecommendationExplanation) GetFeatures() []ExplanationFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *RecommendationExplanation) GetFeaturesOk() (*[]ExplanationFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *RecommendationExplanation) SetFeatures(v []ExplanationFeature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *RecommendationExplanation) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetGuardrails

`func (o *RecommendationExplanation) GetGuardrails() []Guardrail`

GetGuardrails returns the Guardrails field if non-nil, zero value otherwise.

### GetGuardrailsOk

`func (o *RecommendationExplanation) GetGuardrailsOk() (*[]Guardrail, bool)`

GetGuardrailsOk returns a tuple with the Guardrails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrails

`func (o *RecommendationExplanation) SetGuardrails(v []Guardrail)`

SetGuardrails sets Guardrails field to given value.

### HasGuardrails

`func (o *RecommendationExplanation) HasGuardrails() bool`

HasGuardrails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


