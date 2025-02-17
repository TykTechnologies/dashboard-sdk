# \ProxyAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProxyRequest**](ProxyAPI.md#ProxyRequest) | **Post** /api/proxy | Proxy API request



## ProxyRequest

> ProxyResponse ProxyRequest(ctx).ProxyRequest(proxyRequest).Execute()

Proxy API request



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
	proxyRequest := *openapiclient.NewProxyRequest("Method_example", "Url_example") // ProxyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProxyAPI.ProxyRequest(context.Background()).ProxyRequest(proxyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyAPI.ProxyRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProxyRequest`: ProxyResponse
	fmt.Fprintf(os.Stdout, "Response from `ProxyAPI.ProxyRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProxyRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proxyRequest** | [**ProxyRequest**](ProxyRequest.md) |  | 

### Return type

[**ProxyResponse**](ProxyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

