# \SingleSignOnAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateAuthToken**](SingleSignOnAPI.md#GenerateAuthToken) | **Post** /api/sso/ | Generate authentication token



## GenerateAuthToken

> ApiStatusMessage GenerateAuthToken(ctx).GenerateAuthTokenRequest(generateAuthTokenRequest).Execute()

Generate authentication token



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
	generateAuthTokenRequest := *openapiclient.NewGenerateAuthTokenRequest() // GenerateAuthTokenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SingleSignOnAPI.GenerateAuthToken(context.Background()).GenerateAuthTokenRequest(generateAuthTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SingleSignOnAPI.GenerateAuthToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateAuthToken`: ApiStatusMessage
	fmt.Fprintf(os.Stdout, "Response from `SingleSignOnAPI.GenerateAuthToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateAuthTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateAuthTokenRequest** | [**GenerateAuthTokenRequest**](GenerateAuthTokenRequest.md) |  | 

### Return type

[**ApiStatusMessage**](ApiStatusMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

