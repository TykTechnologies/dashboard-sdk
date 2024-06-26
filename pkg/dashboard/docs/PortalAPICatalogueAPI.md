# \PortalAPICatalogueAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePortalCatalogue**](PortalAPICatalogueAPI.md#CreatePortalCatalogue) | **Post** /api/portal/catalogue | Create Portal Catalogue
[**DeleteDocumentation**](PortalAPICatalogueAPI.md#DeleteDocumentation) | **Delete** /api/portal/documentation/{id} | Delete documentation
[**GetPortalCatalogue**](PortalAPICatalogueAPI.md#GetPortalCatalogue) | **Get** /api/portal/catalogue | Get Portal Catalogue
[**UpdatePortalCatalogue**](PortalAPICatalogueAPI.md#UpdatePortalCatalogue) | **Put** /api/portal/catalogue | Update Organization Portal Catalogue
[**UploadDocumentation**](PortalAPICatalogueAPI.md#UploadDocumentation) | **Post** /api/portal/documentation | Create Documentation



## CreatePortalCatalogue

> ApiError CreatePortalCatalogue(ctx).ApiCatalogue(apiCatalogue).Execute()

Create Portal Catalogue



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
	apiCatalogue := *openapiclient.NewApiCatalogue([]openapiclient.APIDescription{*openapiclient.NewAPIDescription()}) // ApiCatalogue |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalAPICatalogueAPI.CreatePortalCatalogue(context.Background()).ApiCatalogue(apiCatalogue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalAPICatalogueAPI.CreatePortalCatalogue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePortalCatalogue`: ApiError
	fmt.Fprintf(os.Stdout, "Response from `PortalAPICatalogueAPI.CreatePortalCatalogue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortalCatalogueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiCatalogue** | [**ApiCatalogue**](ApiCatalogue.md) |  | 

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


## DeleteDocumentation

> ApiError DeleteDocumentation(ctx, id).Execute()

Delete documentation



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
	id := "665453c35715ec051dc8fc02" // string | ID of the documentation you want to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalAPICatalogueAPI.DeleteDocumentation(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalAPICatalogueAPI.DeleteDocumentation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDocumentation`: ApiError
	fmt.Fprintf(os.Stdout, "Response from `PortalAPICatalogueAPI.DeleteDocumentation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the documentation you want to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDocumentationRequest struct via the builder pattern


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


## GetPortalCatalogue

> []ApiCatalogue GetPortalCatalogue(ctx).Execute()

Get Portal Catalogue



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalAPICatalogueAPI.GetPortalCatalogue(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalAPICatalogueAPI.GetPortalCatalogue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortalCatalogue`: []ApiCatalogue
	fmt.Fprintf(os.Stdout, "Response from `PortalAPICatalogueAPI.GetPortalCatalogue`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortalCatalogueRequest struct via the builder pattern


### Return type

[**[]ApiCatalogue**](ApiCatalogue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePortalCatalogue

> ApiError UpdatePortalCatalogue(ctx).ApiCatalogue(apiCatalogue).Execute()

Update Organization Portal Catalogue



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
	apiCatalogue := *openapiclient.NewApiCatalogue([]openapiclient.APIDescription{*openapiclient.NewAPIDescription()}) // ApiCatalogue |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalAPICatalogueAPI.UpdatePortalCatalogue(context.Background()).ApiCatalogue(apiCatalogue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalAPICatalogueAPI.UpdatePortalCatalogue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePortalCatalogue`: ApiError
	fmt.Fprintf(os.Stdout, "Response from `PortalAPICatalogueAPI.UpdatePortalCatalogue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePortalCatalogueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiCatalogue** | [**ApiCatalogue**](ApiCatalogue.md) |  | 

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


## UploadDocumentation

> ApiError UploadDocumentation(ctx).APIDocumentation(aPIDocumentation).Execute()

Create Documentation



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
	aPIDocumentation := *openapiclient.NewAPIDocumentation() // APIDocumentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalAPICatalogueAPI.UploadDocumentation(context.Background()).APIDocumentation(aPIDocumentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalAPICatalogueAPI.UploadDocumentation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadDocumentation`: ApiError
	fmt.Fprintf(os.Stdout, "Response from `PortalAPICatalogueAPI.UploadDocumentation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadDocumentationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aPIDocumentation** | [**APIDocumentation**](APIDocumentation.md) |  | 

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

