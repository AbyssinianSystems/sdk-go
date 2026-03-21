# ScoreExplanation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | **string** |  | 
**Rules** | Pointer to [**[]ExplanationRule**](ExplanationRule.md) |  | [optional] 
**Features** | Pointer to [**[]ExplanationFeature**](ExplanationFeature.md) |  | [optional] 

## Methods

### NewScoreExplanation

`func NewScoreExplanation(summary string, ) *ScoreExplanation`

NewScoreExplanation instantiates a new ScoreExplanation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScoreExplanationWithDefaults

`func NewScoreExplanationWithDefaults() *ScoreExplanation`

NewScoreExplanationWithDefaults instantiates a new ScoreExplanation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *ScoreExplanation) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ScoreExplanation) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ScoreExplanation) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetRules

`func (o *ScoreExplanation) GetRules() []ExplanationRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ScoreExplanation) GetRulesOk() (*[]ExplanationRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ScoreExplanation) SetRules(v []ExplanationRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *ScoreExplanation) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetFeatures

`func (o *ScoreExplanation) GetFeatures() []ExplanationFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ScoreExplanation) GetFeaturesOk() (*[]ExplanationFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ScoreExplanation) SetFeatures(v []ExplanationFeature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ScoreExplanation) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


