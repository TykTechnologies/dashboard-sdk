# \SystemAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemStats**](SystemAPI.md#GetSystemStats) | **Get** /api/system/stats | Get system usage info.



## GetSystemStats

> SystemStatsResp GetSystemStats(ctx).StartDay(startDay).EndDay(endDay).Resolution(resolution).Entity(entity).Execute()

Get system usage info.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/TykTechnologies/dashboard-sdk"
)

func main() {
	startDay := time.Now() // string | start date (optional)
	endDay := time.Now() // string | end date (optional)
	resolution := "day" // string | Resolve daily. (optional)
	entity := "apis" // string | The entity for which stats should be retrieved. (optional) (default to "apis")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.GetSystemStats(context.Background()).StartDay(startDay).EndDay(endDay).Resolution(resolution).Entity(entity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetSystemStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemStats`: SystemStatsResp
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetSystemStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDay** | **string** | start date | 
 **endDay** | **string** | end date | 
 **resolution** | **string** | Resolve daily. | 
 **entity** | **string** | The entity for which stats should be retrieved. | [default to &quot;apis&quot;]

### Return type

[**SystemStatsResp**](SystemStatsResp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

