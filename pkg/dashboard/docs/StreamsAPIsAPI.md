# \StreamsAPIsAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStreamsApi**](StreamsAPIsAPI.md#CreateStreamsApi) | **Post** /api/apis/streams | Create Streams API.
[**DeleteStreamsApi**](StreamsAPIsAPI.md#DeleteStreamsApi) | **Delete** /api/apis/streams/{apiId} | Delete Streams API.
[**GetStreamsAPIDetails**](StreamsAPIsAPI.md#GetStreamsAPIDetails) | **Get** /api/apis/streams/{apiId} | Get Streams API details.
[**PatchApiStreams**](StreamsAPIsAPI.md#PatchApiStreams) | **Patch** /api/apis/streams/{apiId} | Patch a single Streams API by ID.
[**UpdateStreamsApi**](StreamsAPIsAPI.md#UpdateStreamsApi) | **Put** /api/apis/streams/{apiId} | Update Streams API.



## CreateStreamsApi

> ApiResponse CreateStreamsApi(ctx).ContentType(contentType).CreateStreamsApiRequest(createStreamsApiRequest).Execute()

Create Streams API.



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
	contentType := "contentType_example" // string | Content type for streams endpoints should be `application/vnd.tyk.streams.oas`
	createStreamsApiRequest := *openapiclient.NewCreateStreamsApiRequest("Openapi_example", *openapiclient.NewInfo1("Title_example", "Version_example"), map[string]interface{}(123)) // CreateStreamsApiRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPIsAPI.CreateStreamsApi(context.Background()).ContentType(contentType).CreateStreamsApiRequest(createStreamsApiRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPIsAPI.CreateStreamsApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStreamsApi`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPIsAPI.CreateStreamsApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStreamsApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** | Content type for streams endpoints should be &#x60;application/vnd.tyk.streams.oas&#x60; | 
 **createStreamsApiRequest** | [**CreateStreamsApiRequest**](CreateStreamsApiRequest.md) |  | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStreamsApi

> ApiResponse DeleteStreamsApi(ctx, apiId).Execute()

Delete Streams API.



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
	apiId := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the API you want to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPIsAPI.DeleteStreamsApi(context.Background(), apiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPIsAPI.DeleteStreamsApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStreamsApi`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPIsAPI.DeleteStreamsApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of the API you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStreamsApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamsAPIDetails

> CreateStreamsApiRequest GetStreamsAPIDetails(ctx, apiId).Execute()

Get Streams API details.



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
	apiId := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the API you want to fetch.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPIsAPI.GetStreamsAPIDetails(context.Background(), apiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPIsAPI.GetStreamsAPIDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreamsAPIDetails`: CreateStreamsApiRequest
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPIsAPI.GetStreamsAPIDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of the API you want to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamsAPIDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateStreamsApiRequest**](CreateStreamsApiRequest.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApiStreams

> ApiResponse PatchApiStreams(ctx, apiId).ContentType(contentType).ListenPath(listenPath).CustomDomain(customDomain).AllowList(allowList).ValidateRequest(validateRequest).MockResponse(mockResponse).Authentication(authentication).PatchApiStreamsRequest(patchApiStreamsRequest).Execute()

Patch a single Streams API by ID.



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
	contentType := "contentType_example" // string | Content type for streams endpoints should be `application/vnd.tyk.streams.oas`
	apiId := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the API you want to patch.
	listenPath := "/user-test/" // string | Listen path for the API. (optional)
	customDomain := "tyk.io" // string | Custom domain for the API. (optional)
	allowList := true // bool | Enable allowList middleware for all endpoints. (optional)
	validateRequest := true // bool | Enable validateRequest middleware for all endpoints having a request body with media type application/json. (optional)
	mockResponse := true // bool | Enable mockResponse middleware for all endpoints having responses configured. (optional)
	authentication := true // bool | Enable/disable the authentication mechanism in your Tyk Gateway for your OAS API. (optional)
	patchApiStreamsRequest := openapiclient.patchApiStreams_request{CreateStreamsApiRequest: openapiclient.NewCreateStreamsApiRequest("Openapi_example", *openapiclient.NewInfo1("Title_example", "Version_example"), map[string]interface{}(123))} // PatchApiStreamsRequest | The content of the file should be the OpenAPI document in JSON format. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPIsAPI.PatchApiStreams(context.Background(), apiId).ContentType(contentType).ListenPath(listenPath).CustomDomain(customDomain).AllowList(allowList).ValidateRequest(validateRequest).MockResponse(mockResponse).Authentication(authentication).PatchApiStreamsRequest(patchApiStreamsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPIsAPI.PatchApiStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchApiStreams`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPIsAPI.PatchApiStreams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of the API you want to patch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** | Content type for streams endpoints should be &#x60;application/vnd.tyk.streams.oas&#x60; | 

 **listenPath** | **string** | Listen path for the API. | 
 **customDomain** | **string** | Custom domain for the API. | 
 **allowList** | **bool** | Enable allowList middleware for all endpoints. | 
 **validateRequest** | **bool** | Enable validateRequest middleware for all endpoints having a request body with media type application/json. | 
 **mockResponse** | **bool** | Enable mockResponse middleware for all endpoints having responses configured. | 
 **authentication** | **bool** | Enable/disable the authentication mechanism in your Tyk Gateway for your OAS API. | 
 **patchApiStreamsRequest** | [**PatchApiStreamsRequest**](PatchApiStreamsRequest.md) | The content of the file should be the OpenAPI document in JSON format. | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStreamsApi

> ApiResponse UpdateStreamsApi(ctx, apiId).ContentType(contentType).CreateStreamsApiRequest(createStreamsApiRequest).Execute()

Update Streams API.



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
	contentType := "contentType_example" // string | Content type for streams endpoints should be `application/vnd.tyk.streams.oas`
	apiId := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the API you want to update.
	createStreamsApiRequest := *openapiclient.NewCreateStreamsApiRequest("Openapi_example", *openapiclient.NewInfo1("Title_example", "Version_example"), map[string]interface{}(123)) // CreateStreamsApiRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPIsAPI.UpdateStreamsApi(context.Background(), apiId).ContentType(contentType).CreateStreamsApiRequest(createStreamsApiRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPIsAPI.UpdateStreamsApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStreamsApi`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPIsAPI.UpdateStreamsApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of the API you want to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStreamsApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** | Content type for streams endpoints should be &#x60;application/vnd.tyk.streams.oas&#x60; | 

 **createStreamsApiRequest** | [**CreateStreamsApiRequest**](CreateStreamsApiRequest.md) |  | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

