# \APIsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApi**](APIsAPI.md#DeleteApi) | **Delete** /api/apis/{id} | Delete a single API by ID
[**GetAPIExample**](APIsAPI.md#GetAPIExample) | **Get** /api/examples/{location} | Get details of a single example API definition
[**GetAPIExamples**](APIsAPI.md#GetAPIExamples) | **Get** /api/examples | Get a list of example API definitions
[**GetApi**](APIsAPI.md#GetApi) | **Get** /api/apis/{id} | Get a single API by ID
[**GetApis**](APIsAPI.md#GetApis) | **Get** /api/apis | Get List of APIs
[**PostApis**](APIsAPI.md#PostApis) | **Post** /api/apis | Create API Definition
[**PutApi**](APIsAPI.md#PutApi) | **Put** /api/apis/{id} | Update API Definition
[**SearchApis**](APIsAPI.md#SearchApis) | **Get** /api/apis/search | Search List of APIs



## DeleteApi

> ApiStatusMessage DeleteApi(ctx, id).Execute()

Delete a single API by ID



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
	id := "id_example" // string | ID of API to get. Can either be internal or public API id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.DeleteApi(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.DeleteApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApi`: ApiStatusMessage
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.DeleteApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of API to get. Can either be internal or public API id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiRequest struct via the builder pattern


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


## GetAPIExample

> APIExampleDetailed GetAPIExample(ctx, location).XTykRepoUrl(xTykRepoUrl).XTykExamplesIndex(xTykExamplesIndex).ApiDef(apiDef).Execute()

Get details of a single example API definition



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
	location := "udg/simple-app" // string | Location of the tyk example in the repository defined in `x-tyk-repo-url`
	xTykRepoUrl := "xTykRepoUrl_example" // string | Repository URL to fetch example API definitions from (optional) (default to "https://github.com/TykTechnologies/tyk-examples.git")
	xTykExamplesIndex := "xTykExamplesIndex_example" // string | URL path to the \"repository.json\" in the repository defined in `x-tyk-repo-url` (optional) (default to "https://raw.githubusercontent.com/TykTechnologies/tyk-examples/main/repository.json")
	apiDef := true // bool | If set to true, the response will contain the APIDefinition of the specified example, the definition will be present as the `apiDefinition` field. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.GetAPIExample(context.Background(), location).XTykRepoUrl(xTykRepoUrl).XTykExamplesIndex(xTykExamplesIndex).ApiDef(apiDef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.GetAPIExample``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAPIExample`: APIExampleDetailed
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.GetAPIExample`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | Location of the tyk example in the repository defined in &#x60;x-tyk-repo-url&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAPIExampleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTykRepoUrl** | **string** | Repository URL to fetch example API definitions from | [default to &quot;https://github.com/TykTechnologies/tyk-examples.git&quot;]
 **xTykExamplesIndex** | **string** | URL path to the \&quot;repository.json\&quot; in the repository defined in &#x60;x-tyk-repo-url&#x60; | [default to &quot;https://raw.githubusercontent.com/TykTechnologies/tyk-examples/main/repository.json&quot;]
 **apiDef** | **bool** | If set to true, the response will contain the APIDefinition of the specified example, the definition will be present as the &#x60;apiDefinition&#x60; field. | [default to false]

### Return type

[**APIExampleDetailed**](APIExampleDetailed.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAPIExamples

> GetAPIExamples200Response GetAPIExamples(ctx).P(p).XTykExamplesIndex(xTykExamplesIndex).Execute()

Get a list of example API definitions



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
	p := int32(56) // int32 | Page number. If not provided, the first page will be returned (optional)
	xTykExamplesIndex := "xTykExamplesIndex_example" // string | URL that points to the index file (repository.json) or root of a repository housing examples. (optional) (default to "https://raw.githubusercontent.com/TykTechnologies/tyk-examples/main/repository.json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.GetAPIExamples(context.Background()).P(p).XTykExamplesIndex(xTykExamplesIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.GetAPIExamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAPIExamples`: GetAPIExamples200Response
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.GetAPIExamples`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAPIExamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **p** | **int32** | Page number. If not provided, the first page will be returned | 
 **xTykExamplesIndex** | **string** | URL that points to the index file (repository.json) or root of a repository housing examples. | [default to &quot;https://raw.githubusercontent.com/TykTechnologies/tyk-examples/main/repository.json&quot;]

### Return type

[**GetAPIExamples200Response**](GetAPIExamples200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApi

> APIDefinition GetApi(ctx, id).Execute()

Get a single API by ID



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
	id := "id_example" // string | ID of API to get. Can either be internal or public API id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.GetApi(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.GetApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApi`: APIDefinition
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.GetApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of API to get. Can either be internal or public API id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**APIDefinition**](APIDefinition.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApis

> []APIDefinition GetApis(ctx).Compressed(compressed).Graph(graph).ApiType(apiType).Q(q).Sort(sort).Category(category).AuthType(authType).BaseApis(baseApis).Execute()

Get List of APIs



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
	compressed := "compressed_example" // string | Return smaller API list payload. (optional)
	graph := "graph_example" // string | Transform the response payload for graphql. (optional)
	apiType := "apiType_example" // string | API Type, internal or external. (optional)
	q := "q_example" // string | Query string for search/filtering. (optional)
	sort := "sort_example" // string | Sorting criteria. (optional)
	category := "category_example" // string | Filter APIs by category (CSV). (optional)
	authType := "authType_example" // string | Filter APIs by authentication type (CSV). (optional)
	baseApis := "baseApis_example" // string | For versioned APIs, return only the base versions. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.GetApis(context.Background()).Compressed(compressed).Graph(graph).ApiType(apiType).Q(q).Sort(sort).Category(category).AuthType(authType).BaseApis(baseApis).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.GetApis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApis`: []APIDefinition
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.GetApis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compressed** | **string** | Return smaller API list payload. | 
 **graph** | **string** | Transform the response payload for graphql. | 
 **apiType** | **string** | API Type, internal or external. | 
 **q** | **string** | Query string for search/filtering. | 
 **sort** | **string** | Sorting criteria. | 
 **category** | **string** | Filter APIs by category (CSV). | 
 **authType** | **string** | Filter APIs by authentication type (CSV). | 
 **baseApis** | **string** | For versioned APIs, return only the base versions. | 

### Return type

[**[]APIDefinition**](APIDefinition.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApis

> ApiModifyKeySuccess PostApis(ctx).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).APIDefinition(aPIDefinition).Execute()

Create API Definition



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
	aPIDefinition := *openapiclient.NewAPIDefinition() // APIDefinition |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.PostApis(context.Background()).BaseApiId(baseApiId).BaseApiVersionName(baseApiVersionName).NewVersionName(newVersionName).SetDefault(setDefault).APIDefinition(aPIDefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.PostApis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApis`: ApiModifyKeySuccess
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
 **aPIDefinition** | [**APIDefinition**](APIDefinition.md) |  | 

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


## PutApi

> ApiModifyKeySuccess PutApi(ctx, id).APIDefinition(aPIDefinition).Execute()

Update API Definition



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
	id := "id_example" // string | ID of API to update. Can either be internal or public API id.
	aPIDefinition := *openapiclient.NewAPIDefinition() // APIDefinition |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.PutApi(context.Background(), id).APIDefinition(aPIDefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.PutApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutApi`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.PutApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of API to update. Can either be internal or public API id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIDefinition** | [**APIDefinition**](APIDefinition.md) |  | 

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


## SearchApis

> []APIDefinition SearchApis(ctx).Q(q).Execute()

Search List of APIs



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
	q := "q_example" // string | Search query (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsAPI.SearchApis(context.Background()).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsAPI.SearchApis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchApis`: []APIDefinition
	fmt.Fprintf(os.Stdout, "Response from `APIsAPI.SearchApis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchApisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Search query | 

### Return type

[**[]APIDefinition**](APIDefinition.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

