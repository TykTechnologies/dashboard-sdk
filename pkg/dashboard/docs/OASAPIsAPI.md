# \OASAPIsAPI

All URIs are relative to *https://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiOAS**](OASAPIsAPI.md#CreateApiOAS) | **Post** /api/apis/oas | Create OAS API
[**DeleteOASApi**](OASAPIsAPI.md#DeleteOASApi) | **Delete** /api/apis/oas/{apiId} | Delete OAS API
[**DownloadApiOASPublic**](OASAPIsAPI.md#DownloadApiOASPublic) | **Get** /api/apis/oas/{apiId}/export | Export OAS API
[**GetApiCategories**](OASAPIsAPI.md#GetApiCategories) | **Get** /api/apis/oas/{apiId}/categories | Get OAS API&#39;s Categories
[**GetOASAPIDetails**](OASAPIsAPI.md#GetOASAPIDetails) | **Get** /api/apis/oas/{apiId} | Get OAS API Details
[**ImportOAS**](OASAPIsAPI.md#ImportOAS) | **Post** /api/apis/oas/import | Import OAS
[**ListOASApiVersions**](OASAPIsAPI.md#ListOASApiVersions) | **Get** /api/apis/oas/{apiId}/versions | List OAS API Versions
[**PatchApiOAS**](OASAPIsAPI.md#PatchApiOAS) | **Patch** /api/apis/oas/{apiId} | Patch a single OAS API by ID
[**UpdateApiCategories**](OASAPIsAPI.md#UpdateApiCategories) | **Put** /api/apis/oas/{apiId}/categories | Update Oas Api categories
[**UpdateApiOAS**](OASAPIsAPI.md#UpdateApiOAS) | **Put** /api/apis/oas/{apiId} | Update OAS API



## CreateApiOAS

> ApiError CreateApiOAS(ctx).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).TemplateID(templateID).Schema(schema).Execute()

Create OAS API



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
	baseApiId := "663a4ed9b6be920001b191ae" // string | The base API which the new version will be linked to. (optional)
	baseApiVersionName := "Default" // string | The version name of the base API while creating the first version. This doesn't have to be sent for the next versions but if it is set, it will override base API version name. (optional)
	newVersionName := "v2" // string | The version name of the created version. (optional)
	setDefault := true // bool | If true, the new version is set as default version. (optional)
	templateID := "my-unique-template-id" // string | The asset ID of template to apply while creating the OAS API (optional)
	schema := *openapiclient.NewSchema("Openapi_example", *openapiclient.NewInfo("Title_example", "Version_example"), map[string]interface{}(123)) // Schema |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.CreateApiOAS(context.Background()).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).TemplateID(templateID).Schema(schema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.CreateApiOAS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApiOAS`: ApiError
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.CreateApiOAS`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseApiId** | **string** | The base API which the new version will be linked to. | 
 **baseApiVersionName** | **string** | The version name of the base API while creating the first version. This doesn&#39;t have to be sent for the next versions but if it is set, it will override base API version name. | 
 **newVersionName** | **string** | The version name of the created version. | 
 **setDefault** | **bool** | If true, the new version is set as default version. | 
 **templateID** | **string** | The asset ID of template to apply while creating the OAS API | 
 **schema** | [**Schema**](Schema.md) |  | 

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


## DeleteOASApi

> ApiError DeleteOASApi(ctx, apiId).Execute()

Delete OAS API



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
	apiId := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the api you want to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.DeleteOASApi(context.Background(), apiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.DeleteOASApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOASApi`: ApiError
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.DeleteOASApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of the api you want to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOASApiRequest struct via the builder pattern


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


## DownloadApiOASPublic

> *os.File DownloadApiOASPublic(ctx, apiId).Mode(mode).Execute()

Export OAS API



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
	apiId := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the api you want to export
	mode := "public" // string | Mode of OAS export, by default mode could be empty which means to export OAS spec including OAS Tyk extension.When mode=public, OAS spec excluding Tyk extension is exported (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.DownloadApiOASPublic(context.Background(), apiId).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.DownloadApiOASPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadApiOASPublic`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.DownloadApiOASPublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of the api you want to export | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadApiOASPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mode** | **string** | Mode of OAS export, by default mode could be empty which means to export OAS spec including OAS Tyk extension.When mode&#x3D;public, OAS spec excluding Tyk extension is exported | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiCategories

> CategoriesPayload GetApiCategories(ctx, apiId).Execute()

Get OAS API's Categories



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
	apiId := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the api

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.GetApiCategories(context.Background(), apiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.GetApiCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiCategories`: CategoriesPayload
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.GetApiCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of the api | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CategoriesPayload**](CategoriesPayload.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOASAPIDetails

> Schema GetOASAPIDetails(ctx, apiId).Graph(graph).Execute()

Get OAS API Details



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
	apiId := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the api you want to fetch
	graph := "1" // string | Transform the response payload for graphql. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.GetOASAPIDetails(context.Background(), apiId).Graph(graph).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.GetOASAPIDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOASAPIDetails`: Schema
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.GetOASAPIDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of the api you want to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOASAPIDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **graph** | **string** | Transform the response payload for graphql. | 

### Return type

[**Schema**](Schema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportOAS

> ApiError ImportOAS(ctx).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).ValidateRequest(validateRequest).MockResponse(mockResponse).AllowList(allowList).Authentication(authentication).TemplateID(templateID).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).PatchApiOASRequest(patchApiOASRequest).Execute()

Import OAS



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
	upstreamURL := "https://localhost:8080" // string | Upstream URL for the API (optional)
	listenPath := "/user-test/" // string | Listen path for the API (optional)
	customDomain := "tyk.io" // string | Custom domain for the API (optional)
	validateRequest := true // bool | Enable validateRequest middleware for all endpoints having a request body with media type application/json (optional)
	mockResponse := true // bool | Enable mockResponse middleware for all endpoints having responses configured. (optional)
	allowList := true // bool | Enable allowList middleware for all endpoints (optional)
	authentication := true // bool | Enable/disable the authentication mechanism in your Tyk Gateway for your OAS API (optional)
	templateID := "my-unique-template-id" // string | The asset ID of template to apply while importing an OAS API. (optional)
	baseApiId := "663a4ed9b6be920001b191ae" // string | The base API which the new version will be linked to. (optional)
	baseApiVersionName := "Default" // string | The version name of the base API while creating the first version. This doesn't have to be sent for the next versions but if it is set, it will override base API version name. (optional)
	newVersionName := "v2" // string | The version name of the created version. (optional)
	setDefault := true // bool | If true, the new version is set as default version. (optional)
	patchApiOASRequest := openapiclient.patchApiOAS_request{ApiImportByUrlPayload: openapiclient.NewApiImportByUrlPayload()} // PatchApiOASRequest | The content of the file should be the OpenAPI document in JSON format (without x-tyk-api-gateway extension). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.ImportOAS(context.Background()).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).ValidateRequest(validateRequest).MockResponse(mockResponse).AllowList(allowList).Authentication(authentication).TemplateID(templateID).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).PatchApiOASRequest(patchApiOASRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.ImportOAS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportOAS`: ApiError
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.ImportOAS`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upstreamURL** | **string** | Upstream URL for the API | 
 **listenPath** | **string** | Listen path for the API | 
 **customDomain** | **string** | Custom domain for the API | 
 **validateRequest** | **bool** | Enable validateRequest middleware for all endpoints having a request body with media type application/json | 
 **mockResponse** | **bool** | Enable mockResponse middleware for all endpoints having responses configured. | 
 **allowList** | **bool** | Enable allowList middleware for all endpoints | 
 **authentication** | **bool** | Enable/disable the authentication mechanism in your Tyk Gateway for your OAS API | 
 **templateID** | **string** | The asset ID of template to apply while importing an OAS API. | 
 **baseApiId** | **string** | The base API which the new version will be linked to. | 
 **baseApiVersionName** | **string** | The version name of the base API while creating the first version. This doesn&#39;t have to be sent for the next versions but if it is set, it will override base API version name. | 
 **newVersionName** | **string** | The version name of the created version. | 
 **setDefault** | **bool** | If true, the new version is set as default version. | 
 **patchApiOASRequest** | [**PatchApiOASRequest**](PatchApiOASRequest.md) | The content of the file should be the OpenAPI document in JSON format (without x-tyk-api-gateway extension). | 

### Return type

[**ApiError**](ApiError.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOASApiVersions

> VersionMetas ListOASApiVersions(ctx, apiId).SearchText(searchText).AccessType(accessType).Execute()

List OAS API Versions



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
	apiId := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the api you want to fetch
	searchText := "Sample oas" // string | Search for API version name (optional)
	accessType := "internal" // string | Filter for internal or external API versions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.ListOASApiVersions(context.Background(), apiId).SearchText(searchText).AccessType(accessType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.ListOASApiVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOASApiVersions`: VersionMetas
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.ListOASApiVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of the api you want to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOASApiVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchText** | **string** | Search for API version name | 
 **accessType** | **string** | Filter for internal or external API versions | 

### Return type

[**VersionMetas**](VersionMetas.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApiOAS

> ApiError PatchApiOAS(ctx, apiId).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).ValidateRequest(validateRequest).MockResponse(mockResponse).AllowList(allowList).Authentication(authentication).PatchApiOASRequest(patchApiOASRequest).Execute()

Patch a single OAS API by ID



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
	apiId := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the api you want to patch
	upstreamURL := "https://localhost:8080" // string | Upstream URL for the API (optional)
	listenPath := "/user-test/" // string | Listen path for the API (optional)
	customDomain := "tyk.io" // string | Custom domain for the API (optional)
	validateRequest := true // bool | Enable validateRequest middleware for all endpoints having a request body with media type application/json (optional)
	mockResponse := true // bool | Enable mockResponse middleware for all endpoints having responses configured. (optional)
	allowList := true // bool | Enable allowList middleware for all endpoints (optional)
	authentication := true // bool | Enable/disable the authentication mechanism in your Tyk Gateway for your OAS API (optional)
	patchApiOASRequest := openapiclient.patchApiOAS_request{ApiImportByUrlPayload: openapiclient.NewApiImportByUrlPayload()} // PatchApiOASRequest | The content of the file should be the OpenAPI document in JSON format (without x-tyk-api-gateway extension). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.PatchApiOAS(context.Background(), apiId).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).ValidateRequest(validateRequest).MockResponse(mockResponse).AllowList(allowList).Authentication(authentication).PatchApiOASRequest(patchApiOASRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.PatchApiOAS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchApiOAS`: ApiError
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.PatchApiOAS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of the api you want to patch | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upstreamURL** | **string** | Upstream URL for the API | 
 **listenPath** | **string** | Listen path for the API | 
 **customDomain** | **string** | Custom domain for the API | 
 **validateRequest** | **bool** | Enable validateRequest middleware for all endpoints having a request body with media type application/json | 
 **mockResponse** | **bool** | Enable mockResponse middleware for all endpoints having responses configured. | 
 **allowList** | **bool** | Enable allowList middleware for all endpoints | 
 **authentication** | **bool** | Enable/disable the authentication mechanism in your Tyk Gateway for your OAS API | 
 **patchApiOASRequest** | [**PatchApiOASRequest**](PatchApiOASRequest.md) | The content of the file should be the OpenAPI document in JSON format (without x-tyk-api-gateway extension). | 

### Return type

[**ApiError**](ApiError.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiCategories

> ApiError UpdateApiCategories(ctx, apiId).CategoriesPayload(categoriesPayload).Execute()

Update Oas Api categories



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
	apiId := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the api
	categoriesPayload := *openapiclient.NewCategoriesPayload() // CategoriesPayload |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.UpdateApiCategories(context.Background(), apiId).CategoriesPayload(categoriesPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.UpdateApiCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApiCategories`: ApiError
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.UpdateApiCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of the api | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categoriesPayload** | [**CategoriesPayload**](CategoriesPayload.md) |  | 

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


## UpdateApiOAS

> ApiError UpdateApiOAS(ctx, apiId).Schema(schema).Execute()

Update OAS API



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
	apiId := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the api you want to update
	schema := *openapiclient.NewSchema("Openapi_example", *openapiclient.NewInfo("Title_example", "Version_example"), map[string]interface{}(123)) // Schema |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.UpdateApiOAS(context.Background(), apiId).Schema(schema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.UpdateApiOAS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApiOAS`: ApiError
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.UpdateApiOAS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of the api you want to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **schema** | [**Schema**](Schema.md) |  | 

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

