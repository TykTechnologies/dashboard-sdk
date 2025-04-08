# \OauthAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewClientApp**](OauthAPI.md#CreateNewClientApp) | **Post** /api/apis/oauth/{apiId} | Create a new OAuth2.0 client.
[**DeleteOathClient**](OauthAPI.md#DeleteOathClient) | **Delete** /api/apis/oauth/{apiId}/{clientId} | Delete OAuth client.
[**GetClientTokens**](OauthAPI.md#GetClientTokens) | **Get** /api/apis/oauth/{apiId}/{clientId}/tokens | List OAuth client tokens.
[**GetOAuthClientDetail**](OauthAPI.md#GetOAuthClientDetail) | **Get** /api/apis/oauth/{apiId}/{clientId} | Get single OAuth client details.
[**GetOathClientsList**](OauthAPI.md#GetOathClientsList) | **Get** /api/apis/oauth/{apiId} | List OAuth clients.



## CreateNewClientApp

> OAuthClient CreateNewClientApp(ctx, apiId).NewClientRequest(newClientRequest).Execute()

Create a new OAuth2.0 client.



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
	apiId := "e95400eba23c4a2d4622a722be06fe95" // string | The API’s ID.
	newClientRequest := *openapiclient.NewNewClientRequest() // NewClientRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAPI.CreateNewClientApp(context.Background(), apiId).NewClientRequest(newClientRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAPI.CreateNewClientApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewClientApp`: OAuthClient
	fmt.Fprintf(os.Stdout, "Response from `OauthAPI.CreateNewClientApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | The API’s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewClientAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newClientRequest** | [**NewClientRequest**](NewClientRequest.md) |  | 

### Return type

[**OAuthClient**](OAuthClient.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOathClient

> ApiResponse DeleteOathClient(ctx, apiId, clientId).Execute()

Delete OAuth client.



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
	apiId := "e95400eba23c4a2d4622a722be06fe95" // string | The API’s ID.
	clientId := "2a06b398c17f46908de3dffcb71ef87" // string | The client ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAPI.DeleteOathClient(context.Background(), apiId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAPI.DeleteOathClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOathClient`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `OauthAPI.DeleteOathClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | The API’s ID. | 
**clientId** | **string** | The client ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOathClientRequest struct via the builder pattern


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


## GetClientTokens

> []OAuthClientToken GetClientTokens(ctx, apiId, clientId).Execute()

List OAuth client tokens.



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
	apiId := "e95400eba23c4a2d4622a722be06fe95" // string | The API’s ID.
	clientId := "2a06b398c17f46908de3dffcb71ef87" // string | The client ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAPI.GetClientTokens(context.Background(), apiId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAPI.GetClientTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientTokens`: []OAuthClientToken
	fmt.Fprintf(os.Stdout, "Response from `OauthAPI.GetClientTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | The API’s ID. | 
**clientId** | **string** | The client ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]OAuthClientToken**](OAuthClientToken.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOAuthClientDetail

> OAuthClient GetOAuthClientDetail(ctx, apiId, clientId).Execute()

Get single OAuth client details.



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
	apiId := "e95400eba23c4a2d4622a722be06fe95" // string | The API’s ID.
	clientId := "2a06b398c17f46908de3dffcb71ef87" // string | The client ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAPI.GetOAuthClientDetail(context.Background(), apiId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAPI.GetOAuthClientDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOAuthClientDetail`: OAuthClient
	fmt.Fprintf(os.Stdout, "Response from `OauthAPI.GetOAuthClientDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | The API’s ID. | 
**clientId** | **string** | The client ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuthClientDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OAuthClient**](OAuthClient.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOathClientsList

> OAuthApps GetOathClientsList(ctx, apiId).Execute()

List OAuth clients.



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
	apiId := "e95400eba23c4a2d4622a722be06fe95" // string | The API’s ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAPI.GetOathClientsList(context.Background(), apiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAPI.GetOathClientsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOathClientsList`: OAuthApps
	fmt.Fprintf(os.Stdout, "Response from `OauthAPI.GetOathClientsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | The API’s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOathClientsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OAuthApps**](OAuthApps.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

