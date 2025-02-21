# \AnalyticsAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAnalyticsOfApiKey**](AnalyticsAPI.md#GetAnalyticsOfApiKey) | **Get** /api/activity/keys/{keyHash}/{startDay}/{startMonth}/{startYear}/{EndDay}/{EndMonth}/{EndYear} | Analytics of API Key.
[**GetAnalyticsOfOauthClientId**](AnalyticsAPI.md#GetAnalyticsOfOauthClientId) | **Get** /api/activity/oauthid/{OAuthClientID}/{startDay}/{startMonth}/{startYear}/{EndDay}/{EndMonth}/{EndYear} | Analytics of Oauth Client ID.



## GetAnalyticsOfApiKey

> AggregateAnalyticsData GetAnalyticsOfApiKey(ctx, startDay, startMonth, startYear, endDay, endMonth, endYear, keyHash).Execute()

Analytics of API Key.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TykTechnologies/dashboard-sdk"
)

func main() {
	startDay := "15" // string | Day to start querying the analytics from.
	startMonth := "1" // string | Month to start querying the analytics from.
	startYear := "2024" // string | Year to start querying the analytics from.
	endDay := "20" // string | End date of analytics to query.
	endMonth := "6" // string | End month of analytics to query.
	endYear := "2025" // string | End year of analytics to query.
	keyHash := "keyHash_example" // string | Hash of your API key.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetAnalyticsOfApiKey(context.Background(), startDay, startMonth, startYear, endDay, endMonth, endYear, keyHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetAnalyticsOfApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalyticsOfApiKey`: AggregateAnalyticsData
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetAnalyticsOfApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**startDay** | **string** | Day to start querying the analytics from. | 
**startMonth** | **string** | Month to start querying the analytics from. | 
**startYear** | **string** | Year to start querying the analytics from. | 
**endDay** | **string** | End date of analytics to query. | 
**endMonth** | **string** | End month of analytics to query. | 
**endYear** | **string** | End year of analytics to query. | 
**keyHash** | **string** | Hash of your API key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticsOfApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------








### Return type

[**AggregateAnalyticsData**](AggregateAnalyticsData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyticsOfOauthClientId

> AggregateAnalyticsData GetAnalyticsOfOauthClientId(ctx, startDay, startMonth, startYear, endDay, endMonth, endYear, oAuthClientID).Execute()

Analytics of Oauth Client ID.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TykTechnologies/dashboard-sdk"
)

func main() {
	startDay := "15" // string | Day to start querying the analytics from.
	startMonth := "1" // string | Month to start querying the analytics from.
	startYear := "2024" // string | Year to start querying the analytics from.
	endDay := "20" // string | End date of analytics to query.
	endMonth := "6" // string | End month of analytics to query.
	endYear := "2025" // string | End year of analytics to query.
	oAuthClientID := "oAuthClientID_example" // string | OAuthClientID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetAnalyticsOfOauthClientId(context.Background(), startDay, startMonth, startYear, endDay, endMonth, endYear, oAuthClientID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetAnalyticsOfOauthClientId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalyticsOfOauthClientId`: AggregateAnalyticsData
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetAnalyticsOfOauthClientId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**startDay** | **string** | Day to start querying the analytics from. | 
**startMonth** | **string** | Month to start querying the analytics from. | 
**startYear** | **string** | Year to start querying the analytics from. | 
**endDay** | **string** | End date of analytics to query. | 
**endMonth** | **string** | End month of analytics to query. | 
**endYear** | **string** | End year of analytics to query. | 
**oAuthClientID** | **string** | OAuthClientID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticsOfOauthClientIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------








### Return type

[**AggregateAnalyticsData**](AggregateAnalyticsData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

