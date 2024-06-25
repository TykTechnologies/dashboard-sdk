# \AdditionalPermissionsAPI

All URIs are relative to *https://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAdditionalPermissions**](AdditionalPermissionsAPI.md#ListAdditionalPermissions) | **Get** /api/org/permissions | List additional permissions
[**UpdateAdditionalPermissions**](AdditionalPermissionsAPI.md#UpdateAdditionalPermissions) | **Put** /api/org/permissions | Add/Delete/Update Additional Permission



## ListAdditionalPermissions

> NewAdditionalPermissions ListAdditionalPermissions(ctx).Execute()

List additional permissions



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
	resp, r, err := apiClient.AdditionalPermissionsAPI.ListAdditionalPermissions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdditionalPermissionsAPI.ListAdditionalPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAdditionalPermissions`: NewAdditionalPermissions
	fmt.Fprintf(os.Stdout, "Response from `AdditionalPermissionsAPI.ListAdditionalPermissions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAdditionalPermissionsRequest struct via the builder pattern


### Return type

[**NewAdditionalPermissions**](NewAdditionalPermissions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdditionalPermissions

> ApiError UpdateAdditionalPermissions(ctx).NewAdditionalPermissions(newAdditionalPermissions).Execute()

Add/Delete/Update Additional Permission



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
	newAdditionalPermissions := *openapiclient.NewNewAdditionalPermissions() // NewAdditionalPermissions |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdditionalPermissionsAPI.UpdateAdditionalPermissions(context.Background()).NewAdditionalPermissions(newAdditionalPermissions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdditionalPermissionsAPI.UpdateAdditionalPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAdditionalPermissions`: ApiError
	fmt.Fprintf(os.Stdout, "Response from `AdditionalPermissionsAPI.UpdateAdditionalPermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdditionalPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newAdditionalPermissions** | [**NewAdditionalPermissions**](NewAdditionalPermissions.md) |  | 

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

