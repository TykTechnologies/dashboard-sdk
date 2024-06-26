# \AnalyticsAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAverageUsagePerAPI**](AnalyticsAPI.md#GetAverageUsagePerAPI) | **Get** /api/usage/apis/{startDay}/{startMonth}/{startYear}/{endDay}/{endMonth}/{endYear} | Average Usage Per API



## GetAverageUsagePerAPI

> AggregateAnalyticsData GetAverageUsagePerAPI(ctx, startDay, startMonth, startYear, endDay, endMonth, endYear).Sort(sort).Execute()

Average Usage Per API



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
	startDay := "15" // string | Day of analytics to start querying from
	startMonth := "1" // string | Month of analytics to start querying from
	startYear := "2024" // string | Start year of analytics to query.
	endDay := "20" // string | End date of analytics to query.
	endMonth := "6" // string | End month of analytics to query.
	endYear := "2025" // string | End year of analytics to query.
	sort := "1" // string | Sort the activities (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetAverageUsagePerAPI(context.Background(), startDay, startMonth, startYear, endDay, endMonth, endYear).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetAverageUsagePerAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAverageUsagePerAPI`: AggregateAnalyticsData
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetAverageUsagePerAPI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**startDay** | **string** | Day of analytics to start querying from | 
**startMonth** | **string** | Month of analytics to start querying from | 
**startYear** | **string** | Start year of analytics to query. | 
**endDay** | **string** | End date of analytics to query. | 
**endMonth** | **string** | End month of analytics to query. | 
**endYear** | **string** | End year of analytics to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAverageUsagePerAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **sort** | **string** | Sort the activities | 

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

