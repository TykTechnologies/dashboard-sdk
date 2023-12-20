# \OASAPIsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiOAS**](OASAPIsAPI.md#CreateApiOAS) | **Post** /api/apis/oas | 
[**DeleteOASApi**](OASAPIsAPI.md#DeleteOASApi) | **Delete** /api/apis/oas/{apiID} | 
[**DownloadApiOASPublic**](OASAPIsAPI.md#DownloadApiOASPublic) | **Get** /api/apis/oas/{apiID}/export | 
[**DownloadApisOASPublic**](OASAPIsAPI.md#DownloadApisOASPublic) | **Get** /api/apis/oas/export | 
[**ImportOAS**](OASAPIsAPI.md#ImportOAS) | **Post** /api/apis/oas/import | 
[**ListApiOAS**](OASAPIsAPI.md#ListApiOAS) | **Get** /api/apis/oas/{apiID} | 
[**ListOASApiVersions**](OASAPIsAPI.md#ListOASApiVersions) | **Get** /api/apis/oas/{apiID}/versions | 
[**PatchApiOAS**](OASAPIsAPI.md#PatchApiOAS) | **Patch** /api/apis/oas/{apiID} | Patch a single OAS API by ID
[**UpdateApiOAS**](OASAPIsAPI.md#UpdateApiOAS) | **Put** /api/apis/oas/{apiID} | 



## CreateApiOAS

> ApiModifyKeySuccess CreateApiOAS(ctx).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).Schema(schema).Execute()





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
	baseApiId := "baseApiId_example" // string | The base API which the new version will be linked to. (optional)
	baseApiVersionName := "baseApiVersionName_example" // string | The version name of the base API while creating the first version. This doesn't have to be sent for the next versions but if it is set, it will override base API version name. (optional)
	newVersionName := "newVersionName_example" // string | The version name of the created version. (optional)
	setDefault := true // bool | If true, the new version is set as default version. (optional)
	schema := *openapiclient.NewSchema("Openapi_example", *openapiclient.NewInfo("Title_example", "Version_example"), map[string]interface{}(123)) // Schema |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.CreateApiOAS(context.Background()).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).Schema(schema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.CreateApiOAS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApiOAS`: ApiModifyKeySuccess
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
 **schema** | [**Schema**](Schema.md) |  | 

### Return type

[**ApiModifyKeySuccess**](ApiModifyKeySuccess.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOASApi

> ApiStatusMessage DeleteOASApi(ctx, apiID).Execute()





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
	apiID := "apiID_example" // string | The API ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.DeleteOASApi(context.Background(), apiID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.DeleteOASApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOASApi`: ApiStatusMessage
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.DeleteOASApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOASApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiStatusMessage**](ApiStatusMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadApiOASPublic

> OASSchemaResponse DownloadApiOASPublic(ctx, apiID).Mode(mode).Execute()





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
	apiID := "apiID_example" // string | The API ID
	mode := "public" // string | Mode of OAS export, by default mode could be empty which means to export OAS spec including OAS Tyk extension. When mode=public, OAS spec excluding Tyk extension is exported (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.DownloadApiOASPublic(context.Background(), apiID).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.DownloadApiOASPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadApiOASPublic`: OASSchemaResponse
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.DownloadApiOASPublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadApiOASPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mode** | **string** | Mode of OAS export, by default mode could be empty which means to export OAS spec including OAS Tyk extension. When mode&#x3D;public, OAS spec excluding Tyk extension is exported | 

### Return type

[**OASSchemaResponse**](OASSchemaResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadApisOASPublic

> []OASSchemaResponse DownloadApisOASPublic(ctx).Mode(mode).Execute()





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
	mode := "public" // string | The mode of OAS export. By default the mode is not set which means the OAS spec is exported including the OAS Tyk extension.   If the mode is set to public, the OAS spec excluding the Tyk extension is exported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.DownloadApisOASPublic(context.Background()).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.DownloadApisOASPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadApisOASPublic`: []OASSchemaResponse
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.DownloadApisOASPublic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadApisOASPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mode** | **string** | The mode of OAS export. By default the mode is not set which means the OAS spec is exported including the OAS Tyk extension.   If the mode is set to public, the OAS spec excluding the Tyk extension is exported. | 

### Return type

[**[]OASSchemaResponse**](OASSchemaResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportOAS

> ApiModifyKeySuccess ImportOAS(ctx).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).AllowList(allowList).ValidateRequest(validateRequest).MockResponse(mockResponse).Authentication(authentication).Schema(schema).Execute()





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
	upstreamURL := "upstreamURL_example" // string | Upstream URL for the API (optional)
	listenPath := "listenPath_example" // string | Listen path for the API (optional)
	customDomain := "customDomain_example" // string | Custom domain for the API (optional)
	allowList := openapiclient.BooleanQueryParam("true") // BooleanQueryParam | Enable allowList middleware for all endpoints (optional)
	validateRequest := openapiclient.BooleanQueryParam("true") // BooleanQueryParam | Enable validateRequest middleware for all endpoints having a request body with media type application/json (optional)
	mockResponse := openapiclient.BooleanQueryParam("true") // BooleanQueryParam | Enable mockResponse middleware for all endpoints having responses configured. (optional)
	authentication := openapiclient.BooleanQueryParam("true") // BooleanQueryParam | Enable/disable the authentication mechanism in your Tyk Gateway for your OAS API (optional)
	schema := *openapiclient.NewSchema("Openapi_example", *openapiclient.NewInfo("Title_example", "Version_example"), map[string]interface{}(123)) // Schema |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.ImportOAS(context.Background()).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).AllowList(allowList).ValidateRequest(validateRequest).MockResponse(mockResponse).Authentication(authentication).Schema(schema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.ImportOAS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportOAS`: ApiModifyKeySuccess
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
 **allowList** | [**BooleanQueryParam**](BooleanQueryParam.md) | Enable allowList middleware for all endpoints | 
 **validateRequest** | [**BooleanQueryParam**](BooleanQueryParam.md) | Enable validateRequest middleware for all endpoints having a request body with media type application/json | 
 **mockResponse** | [**BooleanQueryParam**](BooleanQueryParam.md) | Enable mockResponse middleware for all endpoints having responses configured. | 
 **authentication** | [**BooleanQueryParam**](BooleanQueryParam.md) | Enable/disable the authentication mechanism in your Tyk Gateway for your OAS API | 
 **schema** | [**Schema**](Schema.md) |  | 

### Return type

[**ApiModifyKeySuccess**](ApiModifyKeySuccess.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiOAS

> OASSchemaResponse ListApiOAS(ctx, apiID).Mode(mode).Execute()





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
	apiID := "apiID_example" // string | The API ID
	mode := "public" // string | The mode of OAS export. By default the mode is not set which means the OAS spec is exported including the OAS Tyk extension.   If the mode is set to public, the OAS spec excluding the Tyk extension is exported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.ListApiOAS(context.Background(), apiID).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.ListApiOAS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApiOAS`: OASSchemaResponse
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.ListApiOAS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApiOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mode** | **string** | The mode of OAS export. By default the mode is not set which means the OAS spec is exported including the OAS Tyk extension.   If the mode is set to public, the OAS spec excluding the Tyk extension is exported. | 

### Return type

[**OASSchemaResponse**](OASSchemaResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOASApiVersions

> ListOASApiVersions200Response ListOASApiVersions(ctx, apiID).SearchText(searchText).AccessType(accessType).Execute()





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
	apiID := "apiID_example" // string | The API ID
	searchText := "searchText_example" // string | Search for API version name (optional)
	accessType := "accessType_example" // string | Filter for internal or external API versions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.ListOASApiVersions(context.Background(), apiID).SearchText(searchText).AccessType(accessType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.ListOASApiVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOASApiVersions`: ListOASApiVersions200Response
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.ListOASApiVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOASApiVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchText** | **string** | Search for API version name | 
 **accessType** | **string** | Filter for internal or external API versions | 

### Return type

[**ListOASApiVersions200Response**](ListOASApiVersions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApiOAS

> ApiModifyKeySuccess PatchApiOAS(ctx, apiID).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).ValidateRequest(validateRequest).AllowList(allowList).MockResponse(mockResponse).Authentication(authentication).Schema(schema).Execute()

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
	apiID := "apiID_example" // string | The API ID
	upstreamURL := "upstreamURL_example" // string | Upstream URL for the API (optional)
	listenPath := "listenPath_example" // string | Listen path for the API (optional)
	customDomain := "customDomain_example" // string | Custom domain for the API (optional)
	validateRequest := openapiclient.BooleanQueryParam("true") // BooleanQueryParam | Enable validateRequest middleware for all endpoints having a request body with media type application/json (optional)
	allowList := openapiclient.BooleanQueryParam("true") // BooleanQueryParam | Enable allowList middleware for all endpoints (optional)
	mockResponse := openapiclient.BooleanQueryParam("true") // BooleanQueryParam | Enable mockResponse middleware for all endpoints having responses configured. (optional)
	authentication := openapiclient.BooleanQueryParam("true") // BooleanQueryParam | Enable/disable the authentication mechanism in your Tyk Gateway for your OAS API (optional)
	schema := *openapiclient.NewSchema("Openapi_example", *openapiclient.NewInfo("Title_example", "Version_example"), map[string]interface{}(123)) // Schema |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.PatchApiOAS(context.Background(), apiID).UpstreamURL(upstreamURL).ListenPath(listenPath).CustomDomain(customDomain).ValidateRequest(validateRequest).AllowList(allowList).MockResponse(mockResponse).Authentication(authentication).Schema(schema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.PatchApiOAS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchApiOAS`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.PatchApiOAS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upstreamURL** | **string** | Upstream URL for the API | 
 **listenPath** | **string** | Listen path for the API | 
 **customDomain** | **string** | Custom domain for the API | 
 **validateRequest** | [**BooleanQueryParam**](BooleanQueryParam.md) | Enable validateRequest middleware for all endpoints having a request body with media type application/json | 
 **allowList** | [**BooleanQueryParam**](BooleanQueryParam.md) | Enable allowList middleware for all endpoints | 
 **mockResponse** | [**BooleanQueryParam**](BooleanQueryParam.md) | Enable mockResponse middleware for all endpoints having responses configured. | 
 **authentication** | [**BooleanQueryParam**](BooleanQueryParam.md) | Enable/disable the authentication mechanism in your Tyk Gateway for your OAS API | 
 **schema** | [**Schema**](Schema.md) |  | 

### Return type

[**ApiModifyKeySuccess**](ApiModifyKeySuccess.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiOAS

> ApiModifyKeySuccess UpdateApiOAS(ctx, apiID).Schema(schema).Execute()





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
	apiID := "apiID_example" // string | The API ID
	schema := *openapiclient.NewSchema("Openapi_example", *openapiclient.NewInfo("Title_example", "Version_example"), map[string]interface{}(123)) // Schema |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OASAPIsAPI.UpdateApiOAS(context.Background(), apiID).Schema(schema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OASAPIsAPI.UpdateApiOAS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApiOAS`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `OASAPIsAPI.UpdateApiOAS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiOASRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **schema** | [**Schema**](Schema.md) |  | 

### Return type

[**ApiModifyKeySuccess**](ApiModifyKeySuccess.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

