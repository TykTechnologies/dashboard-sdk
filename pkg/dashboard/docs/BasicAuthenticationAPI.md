# \BasicAuthenticationAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBasicAuthUser**](BasicAuthenticationAPI.md#CreateBasicAuthUser) | **Post** /api/apis/keys/basic/{username} | Create a Basic Auth User
[**DeleteBasicAuthUser**](BasicAuthenticationAPI.md#DeleteBasicAuthUser) | **Delete** /api/apis/keys/basic/{username} | Delete a Basic Auth User
[**GetBasicAuthUser**](BasicAuthenticationAPI.md#GetBasicAuthUser) | **Get** /api/apis/keys/basic/{username} | Get a Basic Auth User
[**UpdateBasicAuthUser**](BasicAuthenticationAPI.md#UpdateBasicAuthUser) | **Put** /api/apis/keys/basic/{username} | Update a Basic Auth User



## CreateBasicAuthUser

> SessionState CreateBasicAuthUser(ctx, username).SessionState(sessionState).Execute()

Create a Basic Auth User



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
	username := "username_example" // string | Username of Basic Auth user to create, get, update, or delete.
	sessionState := *openapiclient.NewSessionState() // SessionState |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAuthenticationAPI.CreateBasicAuthUser(context.Background(), username).SessionState(sessionState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAuthenticationAPI.CreateBasicAuthUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBasicAuthUser`: SessionState
	fmt.Fprintf(os.Stdout, "Response from `BasicAuthenticationAPI.CreateBasicAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of Basic Auth user to create, get, update, or delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBasicAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sessionState** | [**SessionState**](SessionState.md) |  | 

### Return type

[**SessionState**](SessionState.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBasicAuthUser

> SessionState DeleteBasicAuthUser(ctx, username).Execute()

Delete a Basic Auth User



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
	username := "username_example" // string | Username of Basic Auth user to create, get, update, or delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAuthenticationAPI.DeleteBasicAuthUser(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAuthenticationAPI.DeleteBasicAuthUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBasicAuthUser`: SessionState
	fmt.Fprintf(os.Stdout, "Response from `BasicAuthenticationAPI.DeleteBasicAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of Basic Auth user to create, get, update, or delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBasicAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SessionState**](SessionState.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBasicAuthUser

> SessionState GetBasicAuthUser(ctx, username).Execute()

Get a Basic Auth User



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
	username := "username_example" // string | Username of Basic Auth user to create, get, update, or delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAuthenticationAPI.GetBasicAuthUser(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAuthenticationAPI.GetBasicAuthUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBasicAuthUser`: SessionState
	fmt.Fprintf(os.Stdout, "Response from `BasicAuthenticationAPI.GetBasicAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of Basic Auth user to create, get, update, or delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBasicAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SessionState**](SessionState.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBasicAuthUser

> SessionState UpdateBasicAuthUser(ctx, username).SessionState(sessionState).Execute()

Update a Basic Auth User



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
	username := "username_example" // string | Username of Basic Auth user to create, get, update, or delete.
	sessionState := *openapiclient.NewSessionState() // SessionState |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAuthenticationAPI.UpdateBasicAuthUser(context.Background(), username).SessionState(sessionState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAuthenticationAPI.UpdateBasicAuthUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBasicAuthUser`: SessionState
	fmt.Fprintf(os.Stdout, "Response from `BasicAuthenticationAPI.UpdateBasicAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of Basic Auth user to create, get, update, or delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBasicAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sessionState** | [**SessionState**](SessionState.md) |  | 

### Return type

[**SessionState**](SessionState.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

