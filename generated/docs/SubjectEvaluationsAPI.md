# \SubjectEvaluationsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompareSubjectRuleset**](SubjectEvaluationsAPI.md#CompareSubjectRuleset) | **Post** /v1/subjects/{subject_id}/ruleset-comparisons | Compare the latest evaluation with a candidate ruleset replay
[**GetLatestSubjectEvaluation**](SubjectEvaluationsAPI.md#GetLatestSubjectEvaluation) | **Get** /v1/subjects/{subject_id}/latest-evaluation | Get the latest persisted evaluation for a subject
[**GetOutcomeAnalysis**](SubjectEvaluationsAPI.md#GetOutcomeAnalysis) | **Get** /v1/outcome-analysis | Get deterministic recommendation outcome analysis
[**GetSubjectInvestigation**](SubjectEvaluationsAPI.md#GetSubjectInvestigation) | **Get** /v1/subjects/{subject_id}/investigation | Get the investigation read model for a subject
[**ListSubjectEvaluations**](SubjectEvaluationsAPI.md#ListSubjectEvaluations) | **Get** /v1/subjects/{subject_id}/evaluations | List persisted evaluations for a subject
[**PostSubjectReviewOutcome**](SubjectEvaluationsAPI.md#PostSubjectReviewOutcome) | **Post** /v1/subjects/{subject_id}/review-outcomes | Record a review outcome for a subject artifact
[**RecomputeSubject**](SubjectEvaluationsAPI.md#RecomputeSubject) | **Post** /v1/subjects/{subject_id}/recompute | Recompute and persist a subject evaluation



## CompareSubjectRuleset

> RulesetComparison CompareSubjectRuleset(ctx, subjectId).RulesetComparisonRequest(rulesetComparisonRequest).Execute()

Compare the latest evaluation with a candidate ruleset replay



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/AbyssForge/sdk-go"
)

func main() {
	subjectId := "subjectId_example" // string | Canonical account identifier for the scored subject
	rulesetComparisonRequest := *openapiclient.NewRulesetComparisonRequest("CandidateScoreRulesetVersion_example") // RulesetComparisonRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubjectEvaluationsAPI.CompareSubjectRuleset(context.Background(), subjectId).RulesetComparisonRequest(rulesetComparisonRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubjectEvaluationsAPI.CompareSubjectRuleset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompareSubjectRuleset`: RulesetComparison
	fmt.Fprintf(os.Stdout, "Response from `SubjectEvaluationsAPI.CompareSubjectRuleset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectId** | **string** | Canonical account identifier for the scored subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompareSubjectRulesetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rulesetComparisonRequest** | [**RulesetComparisonRequest**](RulesetComparisonRequest.md) |  | 

### Return type

[**RulesetComparison**](RulesetComparison.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestSubjectEvaluation

> SubjectEvaluationBundle GetLatestSubjectEvaluation(ctx, subjectId).Execute()

Get the latest persisted evaluation for a subject



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/AbyssForge/sdk-go"
)

func main() {
	subjectId := "subjectId_example" // string | Canonical account identifier for the scored subject

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubjectEvaluationsAPI.GetLatestSubjectEvaluation(context.Background(), subjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubjectEvaluationsAPI.GetLatestSubjectEvaluation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestSubjectEvaluation`: SubjectEvaluationBundle
	fmt.Fprintf(os.Stdout, "Response from `SubjectEvaluationsAPI.GetLatestSubjectEvaluation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectId** | **string** | Canonical account identifier for the scored subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestSubjectEvaluationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubjectEvaluationBundle**](SubjectEvaluationBundle.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutcomeAnalysis

> OutcomeAnalysis GetOutcomeAnalysis(ctx).SubjectId(subjectId).RecommendationType(recommendationType).ReviewLabel(reviewLabel).Execute()

Get deterministic recommendation outcome analysis



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/AbyssForge/sdk-go"
)

func main() {
	subjectId := "subjectId_example" // string |  (optional)
	recommendationType := "recommendationType_example" // string |  (optional)
	reviewLabel := "reviewLabel_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubjectEvaluationsAPI.GetOutcomeAnalysis(context.Background()).SubjectId(subjectId).RecommendationType(recommendationType).ReviewLabel(reviewLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubjectEvaluationsAPI.GetOutcomeAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutcomeAnalysis`: OutcomeAnalysis
	fmt.Fprintf(os.Stdout, "Response from `SubjectEvaluationsAPI.GetOutcomeAnalysis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOutcomeAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subjectId** | **string** |  | 
 **recommendationType** | **string** |  | 
 **reviewLabel** | **string** |  | 

### Return type

[**OutcomeAnalysis**](OutcomeAnalysis.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubjectInvestigation

> SubjectInvestigation GetSubjectInvestigation(ctx, subjectId).Execute()

Get the investigation read model for a subject



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/AbyssForge/sdk-go"
)

func main() {
	subjectId := "subjectId_example" // string | Canonical account identifier for the scored subject

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubjectEvaluationsAPI.GetSubjectInvestigation(context.Background(), subjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubjectEvaluationsAPI.GetSubjectInvestigation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubjectInvestigation`: SubjectInvestigation
	fmt.Fprintf(os.Stdout, "Response from `SubjectEvaluationsAPI.GetSubjectInvestigation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectId** | **string** | Canonical account identifier for the scored subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubjectInvestigationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubjectInvestigation**](SubjectInvestigation.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubjectEvaluations

> SubjectEvaluationHistory ListSubjectEvaluations(ctx, subjectId).Execute()

List persisted evaluations for a subject



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/AbyssForge/sdk-go"
)

func main() {
	subjectId := "subjectId_example" // string | Canonical account identifier for the scored subject

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubjectEvaluationsAPI.ListSubjectEvaluations(context.Background(), subjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubjectEvaluationsAPI.ListSubjectEvaluations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubjectEvaluations`: SubjectEvaluationHistory
	fmt.Fprintf(os.Stdout, "Response from `SubjectEvaluationsAPI.ListSubjectEvaluations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectId** | **string** | Canonical account identifier for the scored subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubjectEvaluationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubjectEvaluationHistory**](SubjectEvaluationHistory.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSubjectReviewOutcome

> ReviewOutcome PostSubjectReviewOutcome(ctx, subjectId).ReviewOutcomeWriteRequest(reviewOutcomeWriteRequest).Execute()

Record a review outcome for a subject artifact



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/AbyssForge/sdk-go"
)

func main() {
	subjectId := "subjectId_example" // string | Canonical account identifier for the scored subject
	reviewOutcomeWriteRequest := *openapiclient.NewReviewOutcomeWriteRequest(*openapiclient.NewReviewTargetRef("Kind_example", "Id_example"), "Reviewer_example", "Label_example") // ReviewOutcomeWriteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubjectEvaluationsAPI.PostSubjectReviewOutcome(context.Background(), subjectId).ReviewOutcomeWriteRequest(reviewOutcomeWriteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubjectEvaluationsAPI.PostSubjectReviewOutcome``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSubjectReviewOutcome`: ReviewOutcome
	fmt.Fprintf(os.Stdout, "Response from `SubjectEvaluationsAPI.PostSubjectReviewOutcome`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectId** | **string** | Canonical account identifier for the scored subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSubjectReviewOutcomeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reviewOutcomeWriteRequest** | [**ReviewOutcomeWriteRequest**](ReviewOutcomeWriteRequest.md) |  | 

### Return type

[**ReviewOutcome**](ReviewOutcome.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecomputeSubject

> SubjectEvaluationBundle RecomputeSubject(ctx, subjectId).SubjectRecomputeRequest(subjectRecomputeRequest).Execute()

Recompute and persist a subject evaluation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/AbyssForge/sdk-go"
)

func main() {
	subjectId := "subjectId_example" // string | Canonical account identifier for the scored subject
	subjectRecomputeRequest := *openapiclient.NewSubjectRecomputeRequest() // SubjectRecomputeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubjectEvaluationsAPI.RecomputeSubject(context.Background(), subjectId).SubjectRecomputeRequest(subjectRecomputeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubjectEvaluationsAPI.RecomputeSubject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecomputeSubject`: SubjectEvaluationBundle
	fmt.Fprintf(os.Stdout, "Response from `SubjectEvaluationsAPI.RecomputeSubject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectId** | **string** | Canonical account identifier for the scored subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecomputeSubjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subjectRecomputeRequest** | [**SubjectRecomputeRequest**](SubjectRecomputeRequest.md) |  | 

### Return type

[**SubjectEvaluationBundle**](SubjectEvaluationBundle.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

