# \AnalyticsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAnalyticsOfApiKey**](AnalyticsAPI.md#GetAnalyticsOfApiKey) | **Get** /api/activity/keys/{keyHash}/{startDay}/{startMonth}/{startYear}/{EndDay}/{EndMonth}/{EndYear} | Analytics of API Key
[**GetAnalyticsOfOauthClientId**](AnalyticsAPI.md#GetAnalyticsOfOauthClientId) | **Get** /api/activity/oauthid/{OAuthClientID}/{startDay}/{startMonth}/{startYear}/{EndDay}/{EndMonth}/{EndYear} | Analytics of Oauth Client ID



## GetAnalyticsOfApiKey

> GetAnalyticsOfApiKey(ctx, keyHash, startDay, startMonth, startYear, endDay, endMonth, endYear).Execute()

Analytics of API Key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TykTechnologies/dashboard-sdk/pkg/dashboard"
)

func main() {
	keyHash := "keyHash_example" // string | Hash of your API key.
	startDay := "startDay_example" // string | Day of analytics to query.
	startMonth := "startMonth_example" // string | Month of analytics to query.
	startYear := "startYear_example" // string | Start year of analytics to query.
	endDay := "endDay_example" // string | End date of analytics to query.
	endMonth := "endMonth_example" // string | End month of analytics to query.
	endYear := "endYear_example" // string | End year of analytics to query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnalyticsAPI.GetAnalyticsOfApiKey(context.Background(), keyHash, startDay, startMonth, startYear, endDay, endMonth, endYear).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetAnalyticsOfApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyHash** | **string** | Hash of your API key. | 
**startDay** | **string** | Day of analytics to query. | 
**startMonth** | **string** | Month of analytics to query. | 
**startYear** | **string** | Start year of analytics to query. | 
**endDay** | **string** | End date of analytics to query. | 
**endMonth** | **string** | End month of analytics to query. | 
**endYear** | **string** | End year of analytics to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticsOfApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------








### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyticsOfOauthClientId

> GetAnalyticsOfOauthClientId(ctx, oAuthClientID, startDay, startMonth, startYear, endDay, endMonth, endYear).Execute()

Analytics of Oauth Client ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TykTechnologies/dashboard-sdk/pkg/dashboard"
)

func main() {
	oAuthClientID := "oAuthClientID_example" // string | OAuthClientID
	startDay := "startDay_example" // string | Day of analytics to query.
	startMonth := "startMonth_example" // string | Month of analytics to query.
	startYear := "startYear_example" // string | Start year of analytics to query.
	endDay := "endDay_example" // string | End date of analytics to query.
	endMonth := "endMonth_example" // string | End month of analytics to query.
	endYear := "endYear_example" // string | End year of analytics to query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnalyticsAPI.GetAnalyticsOfOauthClientId(context.Background(), oAuthClientID, startDay, startMonth, startYear, endDay, endMonth, endYear).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetAnalyticsOfOauthClientId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oAuthClientID** | **string** | OAuthClientID | 
**startDay** | **string** | Day of analytics to query. | 
**startMonth** | **string** | Month of analytics to query. | 
**startYear** | **string** | Start year of analytics to query. | 
**endDay** | **string** | End date of analytics to query. | 
**endMonth** | **string** | End month of analytics to query. | 
**endYear** | **string** | End year of analytics to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticsOfOauthClientIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------








### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

