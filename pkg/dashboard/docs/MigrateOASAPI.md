# \MigrateOASAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MigrateAPI**](MigrateOASAPI.md#MigrateAPI) | **Post** /api/apis/migrate | Migrate APIs between Classic and OAS formats



## MigrateAPI

> MigrateAPIResponse MigrateAPI(ctx).MigrateAPIRequest(migrateAPIRequest).Execute()

Migrate APIs between Classic and OAS formats



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
	migrateAPIRequest := *openapiclient.NewMigrateAPIRequest("Mode_example") // MigrateAPIRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MigrateOASAPI.MigrateAPI(context.Background()).MigrateAPIRequest(migrateAPIRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrateOASAPI.MigrateAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MigrateAPI`: MigrateAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `MigrateOASAPI.MigrateAPI`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMigrateAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **migrateAPIRequest** | [**MigrateAPIRequest**](MigrateAPIRequest.md) |  | 

### Return type

[**MigrateAPIResponse**](MigrateAPIResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

