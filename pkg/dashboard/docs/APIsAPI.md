# \APIsAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApi**](APIsAPI.md#DeleteApi) | **Delete** /api/apis/{apiId} | Delete a single API by ID
[**GetAPIExample**](APIsAPI.md#GetAPIExample) | **Get** /api/examples/{location} | Get details of a single example API definition.
[**GetAPIExamples**](APIsAPI.md#GetAPIExamples) | **Get** /api/examples | Get a list of example API definitions.
[**GetAPIGroups**](APIsAPI.md#GetAPIGroups) | **Get** /api/apis/groups | Get API Groups.
[**GetAllApiCategories**](APIsAPI.md#GetAllApiCategories) | **Get** /api/apis/categories | Get API Categories
[**GetApi**](APIsAPI.md#GetApi) | **Get** /api/apis/{apiId} | Get a single API by ID
[**GetApiAccessRights**](APIsAPI.md#GetApiAccessRights) | **Get** /api/apis/{apiID}/access | Get API access rights (users and userGroups)
[**GetApiUrl**](APIsAPI.md#GetApiUrl) | **Get** /api/apis/{apiId}/url | Get API URLs
[**GetApis**](APIsAPI.md#GetApis) | **Get** /api/apis | Get List of APIs
[**PostApis**](APIsAPI.md#PostApis) | **Post** /api/apis | Create API Definition
[**PutApi**](APIsAPI.md#PutApi) | **Put** /api/apis/{apiId} | Update API Definition
[**SearchApis**](APIsAPI.md#SearchApis) | **Get** /api/apis/search | Search List of APIs
[**UpdateApiAccessRights**](APIsAPI.md#UpdateApiAccessRights) | **Put** /api/apis/{apiID}/access | Update API access rights (users and userGroups)



## DeleteApi

> ApiResponse DeleteApi(ctx, apiId).Execute()

Delete a single API by ID



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
	apiId := "b84fe1a04e5648927971c0557971565c" // string | ID of API to delete. Can either be internal or public API ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.DeleteApi(context.Background(), apiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.DeleteApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApi`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.DeleteApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of API to delete. Can either be internal or public API ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiRequest struct via the builder pattern


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


## GetAPIExample

> DetailedExampleAPIMetadata GetAPIExample(ctx, location).XTykExamplesIndex(xTykExamplesIndex).XTykRepoUrl(xTykRepoUrl).Readme(readme).ApiDef(apiDef).Execute()

Get details of a single example API definition.



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
	location := "udg/vat-checker" // string | Location of the tyk example in the repository defined in `x-tyk-repo-url`.
	xTykExamplesIndex := "https://raw.githubusercontent.com/TykTechnologies/tyk-examples/main/repository.json" // string | URL that points to the index file (repository.json) or root of a repository housing examples. (optional) (default to "https://raw.githubusercontent.com/TykTechnologies/tyk-examples/main/repository.json")
	xTykRepoUrl := "https://github.com/TykTechnologies/tyk-examples.git" // string | Repository URL to fetch example API definitions from. (optional) (default to "https://github.com/TykTechnologies/tyk-examples.git")
	readme := false // bool | If set to true, the response will contain the content of the Readme file for the specified example, the readme content will be present as the `readme` field. (optional) (default to false)
	apiDef := false // bool | If set to true, the response will contain the APIDefinition of the specified example, the definition will be present as the `apiDefinition` field. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.GetAPIExample(context.Background(), location).XTykExamplesIndex(xTykExamplesIndex).XTykRepoUrl(xTykRepoUrl).Readme(readme).ApiDef(apiDef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.GetAPIExample``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAPIExample`: DetailedExampleAPIMetadata
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.GetAPIExample`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | Location of the tyk example in the repository defined in &#x60;x-tyk-repo-url&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAPIExampleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTykExamplesIndex** | **string** | URL that points to the index file (repository.json) or root of a repository housing examples. | [default to &quot;https://raw.githubusercontent.com/TykTechnologies/tyk-examples/main/repository.json&quot;]
 **xTykRepoUrl** | **string** | Repository URL to fetch example API definitions from. | [default to &quot;https://github.com/TykTechnologies/tyk-examples.git&quot;]
 **readme** | **bool** | If set to true, the response will contain the content of the Readme file for the specified example, the readme content will be present as the &#x60;readme&#x60; field. | [default to false]
 **apiDef** | **bool** | If set to true, the response will contain the APIDefinition of the specified example, the definition will be present as the &#x60;apiDefinition&#x60; field. | [default to false]

### Return type

[**DetailedExampleAPIMetadata**](DetailedExampleAPIMetadata.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAPIExamples

> []ExampleAPIMetadata GetAPIExamples(ctx).XTykExamplesIndex(xTykExamplesIndex).Execute()

Get a list of example API definitions.



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
	xTykExamplesIndex := "https://raw.githubusercontent.com/TykTechnologies/tyk-examples/main/repository.json" // string | URL path to the repository.json in the repository defined in `x-tyk-repo-url. (optional) (default to "https://raw.githubusercontent.com/TykTechnologies/tyk-examples/main/repository.json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.GetAPIExamples(context.Background()).XTykExamplesIndex(xTykExamplesIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.GetAPIExamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAPIExamples`: []ExampleAPIMetadata
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.GetAPIExamples`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAPIExamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTykExamplesIndex** | **string** | URL path to the repository.json in the repository defined in &#x60;x-tyk-repo-url. | [default to &quot;https://raw.githubusercontent.com/TykTechnologies/tyk-examples/main/repository.json&quot;]

### Return type

[**[]ExampleAPIMetadata**](ExampleAPIMetadata.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAPIGroups

> map[string][]string GetAPIGroups(ctx).Execute()

Get API Groups.



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
	resp, r, err := apiClient.APIsAPI.GetAPIGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.GetAPIGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAPIGroups`: map[string][]string
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.GetAPIGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAPIGroupsRequest struct via the builder pattern


### Return type

[**map[string][]string**](array.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllApiCategories

> AllCategoriesResponse GetAllApiCategories(ctx).Execute()

Get API Categories



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
	resp, r, err := apiClient.APIsAPI.GetAllApiCategories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.GetAllApiCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllApiCategories`: AllCategoriesResponse
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.GetAllApiCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllApiCategoriesRequest struct via the builder pattern


### Return type

[**AllCategoriesResponse**](AllCategoriesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApi

> ApiDefinitionWrapper GetApi(ctx, apiId).Execute()

Get a single API by ID



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
	apiId := "b84fe1a04e5648927971c0557971565c" // string | ID of API to get. Can either be internal or public API ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.GetApi(context.Background(), apiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.GetApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApi`: ApiDefinitionWrapper
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.GetApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of API to get. Can either be internal or public API ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiDefinitionWrapper**](ApiDefinitionWrapper.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiAccessRights

> AccessManagementPayload GetApiAccessRights(ctx, apiID).Execute()

Get API access rights (users and userGroups)



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
	apiID := "b84fe1a04e5648927971c0557971565c" // string | The API ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.GetApiAccessRights(context.Background(), apiID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.GetApiAccessRights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiAccessRights`: AccessManagementPayload
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.GetApiAccessRights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiAccessRightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessManagementPayload**](AccessManagementPayload.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiUrl

> URLVals GetApiUrl(ctx, apiId).Execute()

Get API URLs



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
	apiId := "b84fe1a04e5648927971c0557971565c" // string | The API ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.GetApiUrl(context.Background(), apiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.GetApiUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiUrl`: URLVals
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.GetApiUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | The API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**URLVals**](URLVals.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApis

> ApiDefinitionsResponse GetApis(ctx).Compressed(compressed).ApiType(apiType).Category(category).AuthType(authType).BaseApis(baseApis).Q(q).P(p).Sort(sort).Execute()

Get List of APIs



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
	compressed := "1" // string | Return smaller API list payload. (optional)
	apiType := "rest" // string | API Type, internal or external. (optional)
	category := "category_example" // string | Comma separated list of categories you want to filter Apis by. (optional)
	authType := "keyless,authToken" // string | Comma separated list of authentication type you want to filter apis by. (optional)
	baseApis := "1" // string | For versioned APIs, return only the base versions.If any value is sent in this query parameter only the base version will be returned (optional)
	q := "Rate Limit Path API 1" // string | Query string for search/filtering.This will return all apis whose names matches the given pattern (optional)
	p := int32(1) // int32 | Use p query parameter to say which page you want returned. Send number less than 0 to return all items. (optional)
	sort := "name" // string | * 'name' - Sort by name in ascending order. * '-name' - Sort by name in descending order. * 'updated_at' - Sort by updated_at in ascending order. * '-updated_at' - Sort by updated_at in descending order. * 'created_at' - Sort by created_at in ascending order. * '-created_at' - Sort by created_at in descending order. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.GetApis(context.Background()).Compressed(compressed).ApiType(apiType).Category(category).AuthType(authType).BaseApis(baseApis).Q(q).P(p).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.GetApis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApis`: ApiDefinitionsResponse
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.GetApis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compressed** | **string** | Return smaller API list payload. | 
 **apiType** | **string** | API Type, internal or external. | 
 **category** | **string** | Comma separated list of categories you want to filter Apis by. | 
 **authType** | **string** | Comma separated list of authentication type you want to filter apis by. | 
 **baseApis** | **string** | For versioned APIs, return only the base versions.If any value is sent in this query parameter only the base version will be returned | 
 **q** | **string** | Query string for search/filtering.This will return all apis whose names matches the given pattern | 
 **p** | **int32** | Use p query parameter to say which page you want returned. Send number less than 0 to return all items. | 
 **sort** | **string** | * &#39;name&#39; - Sort by name in ascending order. * &#39;-name&#39; - Sort by name in descending order. * &#39;updated_at&#39; - Sort by updated_at in ascending order. * &#39;-updated_at&#39; - Sort by updated_at in descending order. * &#39;created_at&#39; - Sort by created_at in ascending order. * &#39;-created_at&#39; - Sort by created_at in descending order. | 

### Return type

[**ApiDefinitionsResponse**](ApiDefinitionsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApis

> ApiResponse PostApis(ctx).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).ApiDefinitionWrapper(apiDefinitionWrapper).Execute()

Create API Definition



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
	apiDefinitionWrapper := *openapiclient.NewApiDefinitionWrapper() // ApiDefinitionWrapper |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.PostApis(context.Background()).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).ApiDefinitionWrapper(apiDefinitionWrapper).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.PostApis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApis`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.PostApis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseApiId** | **string** | The base API which the new version will be linked to. | 
 **baseApiVersionName** | **string** | The version name of the base API while creating the first version. This doesn&#39;t have to be sent for the next versions but if it is set, it will override base API version name. | 
 **newVersionName** | **string** | The version name of the created version. | 
 **setDefault** | **bool** | If true, the new version is set as default version. | 
 **apiDefinitionWrapper** | [**ApiDefinitionWrapper**](ApiDefinitionWrapper.md) |  | 

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


## PutApi

> ApiResponse PutApi(ctx, apiId).ApiDefinitionWrapper(apiDefinitionWrapper).Execute()

Update API Definition



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
	apiId := "b84fe1a04e5648927971c0557971565c" // string | ID of API to get. Can either be internal or public API ID.
	apiDefinitionWrapper := *openapiclient.NewApiDefinitionWrapper() // ApiDefinitionWrapper |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.PutApi(context.Background(), apiId).ApiDefinitionWrapper(apiDefinitionWrapper).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.PutApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutApi`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.PutApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of API to get. Can either be internal or public API ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiDefinitionWrapper** | [**ApiDefinitionWrapper**](ApiDefinitionWrapper.md) |  | 

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


## SearchApis

> ApiDefinitionsResponse SearchApis(ctx).P(p).Q(q).Execute()

Search List of APIs



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
	p := int32(1) // int32 | Use p query parameter to say which page you want returned. Send number less than 0 to return all items. (optional)
	q := "Rate Limit Path API 1" // string | The name of the APIs you want to search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.SearchApis(context.Background()).P(p).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.SearchApis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchApis`: ApiDefinitionsResponse
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.SearchApis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchApisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **p** | **int32** | Use p query parameter to say which page you want returned. Send number less than 0 to return all items. | 
 **q** | **string** | The name of the APIs you want to search | 

### Return type

[**ApiDefinitionsResponse**](ApiDefinitionsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiAccessRights

> ApiResponse UpdateApiAccessRights(ctx, apiID).AccessManagementPayload(accessManagementPayload).Execute()

Update API access rights (users and userGroups)



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
	apiID := "b84fe1a04e5648927971c0557971565c" // string | The API ID
	accessManagementPayload := *openapiclient.NewAccessManagementPayload() // AccessManagementPayload |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.UpdateApiAccessRights(context.Background(), apiID).AccessManagementPayload(accessManagementPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.UpdateApiAccessRights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApiAccessRights`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.UpdateApiAccessRights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | The API ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiAccessRightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessManagementPayload** | [**AccessManagementPayload**](AccessManagementPayload.md) |  | 

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

