# \OASAPIsAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiOAS**](OASAPIsAPI.md#CreateApiOAS) | **Post** /api/apis/oas | Create OAS API.
[**DeleteOASApi**](OASAPIsAPI.md#DeleteOASApi) | **Delete** /api/apis/oas/{apiId} | Delete OAS API.
[**DownloadApiOASPublic**](OASAPIsAPI.md#DownloadApiOASPublic) | **Get** /api/apis/oas/{apiId}/export | Export OAS API.
[**DryRunApiOAS**](OASAPIsAPI.md#DryRunApiOAS) | **Post** /api/apis/oas/dry-run | Dry Run OAS.
[**GetApiCategories**](OASAPIsAPI.md#GetApiCategories) | **Get** /api/apis/oas/{apiId}/categories | Get OAS API&#39;s Categories.
[**GetOASAPIDetails**](OASAPIsAPI.md#GetOASAPIDetails) | **Get** /api/apis/oas/{apiId} | Get OAS API details.
[**ImportOAS**](OASAPIsAPI.md#ImportOAS) | **Post** /api/apis/oas/import | Import OAS.
[**ListOASApiVersions**](OASAPIsAPI.md#ListOASApiVersions) | **Get** /api/apis/oas/{apiId}/versions | List OAS API versions.
[**PatchApiOAS**](OASAPIsAPI.md#PatchApiOAS) | **Patch** /api/apis/oas/{apiId} | Patch a single OAS API by ID.
[**UpdateApiCategories**](OASAPIsAPI.md#UpdateApiCategories) | **Put** /api/apis/oas/{apiId}/categories | Update OAS API categories.
[**UpdateApiOAS**](OASAPIsAPI.md#UpdateApiOAS) | **Put** /api/apis/oas/{apiId} | Update OAS API.



## CreateApiOAS

> ApiResponse CreateApiOAS(ctx).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).TemplateID(templateID).CreateApiOASRequest(createApiOASRequest).Execute()

Create OAS API.



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
	baseApiId := "663a4ed9b6be920001b191ae" // string | The base API which the new version will be linked to. (optional)
	baseApiVersionName := "Default" // string | The version name of the base API while creating the first version. This doesn't have to be sent for the next versions but if it is set, it will override base API version name. (optional)
	newVersionName := "v2" // string | The version name of the created version. (optional)
	setDefault := true // bool | If true, the new version is set as default version. (optional)
	templateID := "my-unique-template-id" // string | The Asset ID of template applied while creating or importing an OAS API. (optional)
	createApiOASRequest := *openapiclient.NewCreateApiOASRequest("Openapi_example", *openapiclient.NewInfo1("Title_example", "Version_example"), map[string]interface{}(123)) // CreateApiOASRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.CreateApiOAS(context.Background()).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).TemplateID(templateID).CreateApiOASRequest(createApiOASRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.CreateApiOAS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApiOAS`: ApiResponse
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
 **templateID** | **string** | The Asset ID of template applied while creating or importing an OAS API. | 
 **createApiOASRequest** | [**CreateApiOASRequest**](CreateApiOASRequest.md) |  | 

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


## DeleteOASApi

> ApiResponse DeleteOASApi(ctx, apiId).Execute()

Delete OAS API.



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
	resp, r, err := apiClient.OASAPIsAPI.DeleteOASApi(context.Background(), apiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.DeleteOASApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOASApi`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.DeleteOASApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of the API you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOASApiRequest struct via the builder pattern


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


## DownloadApiOASPublic

> *os.File DownloadApiOASPublic(ctx, apiId).Mode(mode).Execute()

Export OAS API.



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
	apiId := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the API you want to export.
	mode := "public" // string | Mode of OAS export, by default mode could be empty which means to export OAS spec including OAS Tyk extension. When mode=public, OAS spec excluding Tyk extension is exported. (optional)

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
**apiId** | **string** | ID of the API you want to export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadApiOASPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mode** | **string** | Mode of OAS export, by default mode could be empty which means to export OAS spec including OAS Tyk extension. When mode&#x3D;public, OAS spec excluding Tyk extension is exported. | 

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


## DryRunApiOAS

> DryRunApiOAS200Response DryRunApiOAS(ctx).TemplateID(templateID).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).AllowList(allowList).ValidateRequest(validateRequest).MockResponse(mockResponse).Authentication(authentication).DryRunRequest(dryRunRequest).Execute()

Dry Run OAS.



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
	templateID := "my-unique-template-id" // string | The Asset ID of template applied while creating or importing an OAS API. (optional)
	baseApiId := "663a4ed9b6be920001b191ae" // string | The base API which the new version will be linked to. (optional)
	baseApiVersionName := "Default" // string | The version name of the base API while creating the first version. This doesn't have to be sent for the next versions but if it is set, it will override base API version name. (optional)
	newVersionName := "v2" // string | The version name of the created version. (optional)
	setDefault := true // bool | If true, the new version is set as default version. (optional)
	upstreamURL := "https://localhost:8080" // string | Upstream URL for the API. (optional)
	listenPath := "/user-test/" // string | Listen path for the API. (optional)
	customDomain := "tyk.io" // string | Custom domain for the API. (optional)
	allowList := true // bool | Enable allowList middleware for all endpoints. (optional)
	validateRequest := true // bool | Enable validateRequest middleware for all endpoints having a request body with media type application/json. (optional)
	mockResponse := true // bool | Enable mockResponse middleware for all endpoints having responses configured. (optional)
	authentication := true // bool | Enable/disable the authentication mechanism in your Tyk Gateway for your OAS API. (optional)
	dryRunRequest := *openapiclient.NewDryRunRequest() // DryRunRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.DryRunApiOAS(context.Background()).TemplateID(templateID).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).AllowList(allowList).ValidateRequest(validateRequest).MockResponse(mockResponse).Authentication(authentication).DryRunRequest(dryRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.DryRunApiOAS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DryRunApiOAS`: DryRunApiOAS200Response
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.DryRunApiOAS`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDryRunApiOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateID** | **string** | The Asset ID of template applied while creating or importing an OAS API. | 
 **baseApiId** | **string** | The base API which the new version will be linked to. | 
 **baseApiVersionName** | **string** | The version name of the base API while creating the first version. This doesn&#39;t have to be sent for the next versions but if it is set, it will override base API version name. | 
 **newVersionName** | **string** | The version name of the created version. | 
 **setDefault** | **bool** | If true, the new version is set as default version. | 
 **upstreamURL** | **string** | Upstream URL for the API. | 
 **listenPath** | **string** | Listen path for the API. | 
 **customDomain** | **string** | Custom domain for the API. | 
 **allowList** | **bool** | Enable allowList middleware for all endpoints. | 
 **validateRequest** | **bool** | Enable validateRequest middleware for all endpoints having a request body with media type application/json. | 
 **mockResponse** | **bool** | Enable mockResponse middleware for all endpoints having responses configured. | 
 **authentication** | **bool** | Enable/disable the authentication mechanism in your Tyk Gateway for your OAS API. | 
 **dryRunRequest** | [**DryRunRequest**](DryRunRequest.md) |  | 

### Return type

[**DryRunApiOAS200Response**](DryRunApiOAS200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiCategories

> CategoriesPayload GetApiCategories(ctx, apiId).Execute()

Get OAS API's Categories.



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
	apiId := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the API.

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
**apiId** | **string** | ID of the API. | 

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

> CreateApiOASRequest GetOASAPIDetails(ctx, apiId).Execute()

Get OAS API details.



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
	resp, r, err := apiClient.OASAPIsAPI.GetOASAPIDetails(context.Background(), apiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.GetOASAPIDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOASAPIDetails`: CreateApiOASRequest
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.GetOASAPIDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of the API you want to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOASAPIDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateApiOASRequest**](CreateApiOASRequest.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportOAS

> ApiResponse ImportOAS(ctx).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).AllowList(allowList).ValidateRequest(validateRequest).MockResponse(mockResponse).Authentication(authentication).TemplateID(templateID).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).ImportOASRequest(importOASRequest).Execute()

Import OAS.



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
	upstreamURL := "https://localhost:8080" // string | Upstream URL for the API. (optional)
	listenPath := "/user-test/" // string | Listen path for the API. (optional)
	customDomain := "tyk.io" // string | Custom domain for the API. (optional)
	allowList := true // bool | Enable allowList middleware for all endpoints. (optional)
	validateRequest := true // bool | Enable validateRequest middleware for all endpoints having a request body with media type application/json. (optional)
	mockResponse := true // bool | Enable mockResponse middleware for all endpoints having responses configured. (optional)
	authentication := true // bool | Enable/disable the authentication mechanism in your Tyk Gateway for your OAS API. (optional)
	templateID := "my-unique-template-id" // string | The Asset ID of template applied while creating or importing an OAS API. (optional)
	baseApiId := "663a4ed9b6be920001b191ae" // string | The base API which the new version will be linked to. (optional)
	baseApiVersionName := "Default" // string | The version name of the base API while creating the first version. This doesn't have to be sent for the next versions but if it is set, it will override base API version name. (optional)
	newVersionName := "v2" // string | The version name of the created version. (optional)
	setDefault := true // bool | If true, the new version is set as default version. (optional)
	importOASRequest := openapiclient.importOAS_request{ApiImportByUrlPayload: openapiclient.NewApiImportByUrlPayload()} // ImportOASRequest | The content of the file should be the OpenAPI document in JSON format (without x-tyk-api-gateway extension). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.ImportOAS(context.Background()).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).AllowList(allowList).ValidateRequest(validateRequest).MockResponse(mockResponse).Authentication(authentication).TemplateID(templateID).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).ImportOASRequest(importOASRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.ImportOAS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportOAS`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.ImportOAS`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upstreamURL** | **string** | Upstream URL for the API. | 
 **listenPath** | **string** | Listen path for the API. | 
 **customDomain** | **string** | Custom domain for the API. | 
 **allowList** | **bool** | Enable allowList middleware for all endpoints. | 
 **validateRequest** | **bool** | Enable validateRequest middleware for all endpoints having a request body with media type application/json. | 
 **mockResponse** | **bool** | Enable mockResponse middleware for all endpoints having responses configured. | 
 **authentication** | **bool** | Enable/disable the authentication mechanism in your Tyk Gateway for your OAS API. | 
 **templateID** | **string** | The Asset ID of template applied while creating or importing an OAS API. | 
 **baseApiId** | **string** | The base API which the new version will be linked to. | 
 **baseApiVersionName** | **string** | The version name of the base API while creating the first version. This doesn&#39;t have to be sent for the next versions but if it is set, it will override base API version name. | 
 **newVersionName** | **string** | The version name of the created version. | 
 **setDefault** | **bool** | If true, the new version is set as default version. | 
 **importOASRequest** | [**ImportOASRequest**](ImportOASRequest.md) | The content of the file should be the OpenAPI document in JSON format (without x-tyk-api-gateway extension). | 

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


## ListOASApiVersions

> VersionMetas ListOASApiVersions(ctx, apiId).SearchText(searchText).AccessType(accessType).Execute()

List OAS API versions.



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
	searchText := "Sample oas" // string | Search for API version name. (optional)
	accessType := "internal" // string | Filter for internal or external API versions. (optional)

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
**apiId** | **string** | ID of the API you want to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOASApiVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchText** | **string** | Search for API version name. | 
 **accessType** | **string** | Filter for internal or external API versions. | 

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

> ApiResponse PatchApiOAS(ctx, apiId).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).AllowList(allowList).ValidateRequest(validateRequest).MockResponse(mockResponse).Authentication(authentication).PatchApiOASRequest(patchApiOASRequest).Execute()

Patch a single OAS API by ID.



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
	apiId := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the API you want to patch.
	upstreamURL := "https://localhost:8080" // string | Upstream URL for the API. (optional)
	listenPath := "/user-test/" // string | Listen path for the API. (optional)
	customDomain := "tyk.io" // string | Custom domain for the API. (optional)
	allowList := true // bool | Enable allowList middleware for all endpoints. (optional)
	validateRequest := true // bool | Enable validateRequest middleware for all endpoints having a request body with media type application/json. (optional)
	mockResponse := true // bool | Enable mockResponse middleware for all endpoints having responses configured. (optional)
	authentication := true // bool | Enable/disable the authentication mechanism in your Tyk Gateway for your OAS API. (optional)
	patchApiOASRequest := openapiclient.patchApiOAS_request{ApiImportByUrlPayload: openapiclient.NewApiImportByUrlPayload()} // PatchApiOASRequest | The content of the file should be the OpenAPI document in JSON format. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.PatchApiOAS(context.Background(), apiId).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).AllowList(allowList).ValidateRequest(validateRequest).MockResponse(mockResponse).Authentication(authentication).PatchApiOASRequest(patchApiOASRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.PatchApiOAS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchApiOAS`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.PatchApiOAS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of the API you want to patch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upstreamURL** | **string** | Upstream URL for the API. | 
 **listenPath** | **string** | Listen path for the API. | 
 **customDomain** | **string** | Custom domain for the API. | 
 **allowList** | **bool** | Enable allowList middleware for all endpoints. | 
 **validateRequest** | **bool** | Enable validateRequest middleware for all endpoints having a request body with media type application/json. | 
 **mockResponse** | **bool** | Enable mockResponse middleware for all endpoints having responses configured. | 
 **authentication** | **bool** | Enable/disable the authentication mechanism in your Tyk Gateway for your OAS API. | 
 **patchApiOASRequest** | [**PatchApiOASRequest**](PatchApiOASRequest.md) | The content of the file should be the OpenAPI document in JSON format. | 

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


## UpdateApiCategories

> ApiResponse UpdateApiCategories(ctx, apiId).CategoriesPayload(categoriesPayload).Execute()

Update OAS API categories.



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
	apiId := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the API.
	categoriesPayload := *openapiclient.NewCategoriesPayload() // CategoriesPayload |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.UpdateApiCategories(context.Background(), apiId).CategoriesPayload(categoriesPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.UpdateApiCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApiCategories`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.UpdateApiCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of the API. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categoriesPayload** | [**CategoriesPayload**](CategoriesPayload.md) |  | 

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


## UpdateApiOAS

> ApiResponse UpdateApiOAS(ctx, apiId).CreateApiOASRequest(createApiOASRequest).Execute()

Update OAS API.



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
	apiId := "4c1c0d8fc885401053ddac4e39ef676b" // string | ID of the API you want to update.
	createApiOASRequest := *openapiclient.NewCreateApiOASRequest("Openapi_example", *openapiclient.NewInfo1("Title_example", "Version_example"), map[string]interface{}(123)) // CreateApiOASRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.UpdateApiOAS(context.Background(), apiId).CreateApiOASRequest(createApiOASRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.UpdateApiOAS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApiOAS`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.UpdateApiOAS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of the API you want to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createApiOASRequest** | [**CreateApiOASRequest**](CreateApiOASRequest.md) |  | 

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

