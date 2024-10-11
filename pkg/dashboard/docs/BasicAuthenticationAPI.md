# \BasicAuthenticationAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBasicAuthUser**](BasicAuthenticationAPI.md#CreateBasicAuthUser) | **Post** /api/apis/keys/basic/{username} | Create a Basic Auth User.



## CreateBasicAuthUser

> KeyData CreateBasicAuthUser(ctx, username).SessionState(sessionState).Execute()

Create a Basic Auth User.



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
	username := "itachi" // string |  Username of Basic Auth user to create or update.
	sessionState := *openapiclient.NewSessionState() // SessionState |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAuthenticationAPI.CreateBasicAuthUser(context.Background(), username).SessionState(sessionState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAuthenticationAPI.CreateBasicAuthUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBasicAuthUser`: KeyData
	fmt.Fprintf(os.Stdout, "Response from `BasicAuthenticationAPI.CreateBasicAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  Username of Basic Auth user to create or update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBasicAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sessionState** | [**SessionState**](SessionState.md) |  | 

### Return type

[**KeyData**](KeyData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

