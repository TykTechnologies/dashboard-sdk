# \AssetsAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAsset**](AssetsAPI.md#AddAsset) | **Post** /api/assets | Create an asset
[**DeleteAsset**](AssetsAPI.md#DeleteAsset) | **Delete** /api/assets/{assetID} | Delete an asset by ID.
[**GetAsset**](AssetsAPI.md#GetAsset) | **Get** /api/assets/{assetID} | Retrieve an asset.
[**ListAssets**](AssetsAPI.md#ListAssets) | **Get** /api/assets | Retrieve list of assets
[**UpdateAsset**](AssetsAPI.md#UpdateAsset) | **Put** /api/assets/{assetID} | Update an asset by ID.



## AddAsset

> ApiError AddAsset(ctx).AddAssetRequest(addAssetRequest).Execute()

Create an asset



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
	addAssetRequest := *openapiclient.NewAddAssetRequest() // AddAssetRequest | Sample asset (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.AddAsset(context.Background()).AddAssetRequest(addAssetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.AddAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAsset`: ApiError
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.AddAsset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addAssetRequest** | [**AddAssetRequest**](AddAssetRequest.md) | Sample asset | 

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


## DeleteAsset

> ApiError DeleteAsset(ctx, assetID).Execute()

Delete an asset by ID.



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
	assetID := "my-unique-template-id" // string | ID of the asset to delete - this value can be the database ID of the asset or the custom ID provided during creation/update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.DeleteAsset(context.Background(), assetID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.DeleteAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAsset`: ApiError
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.DeleteAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetID** | **string** | ID of the asset to delete - this value can be the database ID of the asset or the custom ID provided during creation/update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiError**](ApiError.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAsset

> Asset GetAsset(ctx, assetID).Execute()

Retrieve an asset.



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
	assetID := "my-unique-template-id" // string | ID of the asset to retrieve - this value can be the database ID of the asset or the custom ID provided during creation/update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.GetAsset(context.Background(), assetID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.GetAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAsset`: Asset
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.GetAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetID** | **string** | ID of the asset to retrieve - this value can be the database ID of the asset or the custom ID provided during creation/update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Asset**](Asset.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssets

> []Asset ListAssets(ctx).Kind(kind).Execute()

Retrieve list of assets



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
	kind := "oas-template" // string | Filter assets by kind (optional). (optional) (default to "oas-template")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.ListAssets(context.Background()).Kind(kind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.ListAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssets`: []Asset
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.ListAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** | Filter assets by kind (optional). | [default to &quot;oas-template&quot;]

### Return type

[**[]Asset**](Asset.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAsset

> ApiError UpdateAsset(ctx, assetID).AddAssetRequest(addAssetRequest).Execute()

Update an asset by ID.



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
	assetID := "my-unique-template-id" // string | ID of the asset to update - this value can be the database ID of the asset or the custom ID provided during creation/update.
	addAssetRequest := *openapiclient.NewAddAssetRequest() // AddAssetRequest | update name example (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.UpdateAsset(context.Background(), assetID).AddAssetRequest(addAssetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.UpdateAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAsset`: ApiError
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.UpdateAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetID** | **string** | ID of the asset to update - this value can be the database ID of the asset or the custom ID provided during creation/update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addAssetRequest** | [**AddAssetRequest**](AddAssetRequest.md) | update name example | 

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

