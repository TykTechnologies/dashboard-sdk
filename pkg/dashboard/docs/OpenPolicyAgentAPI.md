# \OpenPolicyAgentAPI

All URIs are relative to *https://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOPARules**](OpenPolicyAgentAPI.md#GetOPARules) | **Get** /api/org/opa | List OPA Rules
[**UpdateOrgOPARules**](OpenPolicyAgentAPI.md#UpdateOrgOPARules) | **Put** /api/org/opa | Update OPA Rules



## GetOPARules

> NewOPARules GetOPARules(ctx).Execute()

List OPA Rules



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenPolicyAgentAPI.GetOPARules(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenPolicyAgentAPI.GetOPARules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOPARules`: NewOPARules
	fmt.Fprintf(os.Stdout, "Response from `OpenPolicyAgentAPI.GetOPARules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOPARulesRequest struct via the builder pattern


### Return type

[**NewOPARules**](NewOPARules.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgOPARules

> ApiError UpdateOrgOPARules(ctx).NewOPARules(newOPARules).Execute()

Update OPA Rules



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
	newOPARules := *openapiclient.NewNewOPARules() // NewOPARules | Create rule to prevent creation of keyless apis (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenPolicyAgentAPI.UpdateOrgOPARules(context.Background()).NewOPARules(newOPARules).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenPolicyAgentAPI.UpdateOrgOPARules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgOPARules`: ApiError
	fmt.Fprintf(os.Stdout, "Response from `OpenPolicyAgentAPI.UpdateOrgOPARules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgOPARulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newOPARules** | [**NewOPARules**](NewOPARules.md) | Create rule to prevent creation of keyless apis | 

### Return type

[**ApiError**](ApiError.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

