# \DataGraphAPIsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDataGraphsDataSourcesImportPost**](DataGraphAPIsAPI.md#ApiDataGraphsDataSourcesImportPost) | **Post** /api/data-graphs/data-sources/import | Import a new data source



## ApiDataGraphsDataSourcesImportPost

> DataSourceImported ApiDataGraphsDataSourcesImportPost(ctx).NewDataSource(newDataSource).Execute()

Import a new data source

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
	newDataSource := *openapiclient.NewNewDataSource() // NewDataSource |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataGraphAPIsAPI.ApiDataGraphsDataSourcesImportPost(context.Background()).NewDataSource(newDataSource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataGraphAPIsAPI.ApiDataGraphsDataSourcesImportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiDataGraphsDataSourcesImportPost`: DataSourceImported
	fmt.Fprintf(os.Stdout, "Response from `DataGraphAPIsAPI.ApiDataGraphsDataSourcesImportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiDataGraphsDataSourcesImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newDataSource** | [**NewDataSource**](NewDataSource.md) |  | 

### Return type

[**DataSourceImported**](DataSourceImported.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

