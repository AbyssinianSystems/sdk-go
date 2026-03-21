# \SignalEventsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSubjectSignalEvents**](SignalEventsAPI.md#ListSubjectSignalEvents) | **Get** /v1/subjects/{subject_id}/signal-events | List canonical signal events for a subject
[**PostSignalEvent**](SignalEventsAPI.md#PostSignalEvent) | **Post** /v1/signal-events | Ingest a raw detector event



## ListSubjectSignalEvents

> SubjectSignalEventList ListSubjectSignalEvents(ctx, subjectId).Execute()

List canonical signal events for a subject



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/abyssforge/sdk-go"
)

func main() {
	subjectId := "subjectId_example" // string | Canonical account identifier for the scored subject

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalEventsAPI.ListSubjectSignalEvents(context.Background(), subjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalEventsAPI.ListSubjectSignalEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubjectSignalEvents`: SubjectSignalEventList
	fmt.Fprintf(os.Stdout, "Response from `SignalEventsAPI.ListSubjectSignalEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectId** | **string** | Canonical account identifier for the scored subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubjectSignalEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubjectSignalEventList**](SubjectSignalEventList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSignalEvent

> IngestResult PostSignalEvent(ctx).RawSignalEventPayload(rawSignalEventPayload).Execute()

Ingest a raw detector event



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/abyssforge/sdk-go"
)

func main() {
	rawSignalEventPayload := *openapiclient.NewRawSignalEventPayload("SubjectId_example", "Producer_example", "SignalType_example", time.Now()) // RawSignalEventPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalEventsAPI.PostSignalEvent(context.Background()).RawSignalEventPayload(rawSignalEventPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalEventsAPI.PostSignalEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSignalEvent`: IngestResult
	fmt.Fprintf(os.Stdout, "Response from `SignalEventsAPI.PostSignalEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSignalEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rawSignalEventPayload** | [**RawSignalEventPayload**](RawSignalEventPayload.md) |  | 

### Return type

[**IngestResult**](IngestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

