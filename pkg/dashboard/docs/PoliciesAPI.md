# \PoliciesAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePolicy**](PoliciesAPI.md#DeletePolicy) | **Delete** /api/portal/policies/{id} | Delete a single Policy by ID
[**GetPolicies**](PoliciesAPI.md#GetPolicies) | **Get** /api/portal/policies/ | Get List of Policies
[**GetPolicy**](PoliciesAPI.md#GetPolicy) | **Get** /api/portal/policies/{id} | Get a single Policy by ID
[**PostPolicies**](PoliciesAPI.md#PostPolicies) | **Post** /api/portal/policies/ | Create Policy Definition
[**PutPolicies**](PoliciesAPI.md#PutPolicies) | **Put** /api/portal/policies/{id} | Update Policy Definition
[**SearchPolicies**](PoliciesAPI.md#SearchPolicies) | **Get** /api/portal/policies/search | Search List of Policies



## DeletePolicy

> ApiModifyKeySuccess DeletePolicy(ctx, id).Execute()

Delete a single Policy by ID



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
	id := "id_example" // string | ID of Policy to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.DeletePolicy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.DeletePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePolicy`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.DeletePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Policy to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiModifyKeySuccess**](ApiModifyKeySuccess.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicies

> []Policy GetPolicies(ctx).Execute()

Get List of Policies



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
	resp, r, err := apiClient.PoliciesAPI.GetPolicies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.GetPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicies`: []Policy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.GetPolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoliciesRequest struct via the builder pattern


### Return type

[**[]Policy**](Policy.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicy

> Policy GetPolicy(ctx, id).Execute()

Get a single Policy by ID



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
	id := "id_example" // string | ID of Policy to get.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.GetPolicy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.GetPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicy`: Policy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.GetPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Policy to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Policy**](Policy.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPolicies

> ApiModifyKeySuccess PostPolicies(ctx).Policy(policy).Execute()

Create Policy Definition



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
	policy := *openapiclient.NewPolicy() // Policy |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PostPolicies(context.Background()).Policy(policy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PostPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPolicies`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PostPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policy** | [**Policy**](Policy.md) |  | 

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


## PutPolicies

> ApiModifyKeySuccess PutPolicies(ctx, id).Policy(policy).Execute()

Update Policy Definition



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
	id := "id_example" // string | ID of Policy to update.
	policy := *openapiclient.NewPolicy() // Policy |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PutPolicies(context.Background(), id).Policy(policy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PutPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPolicies`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PutPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Policy to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policy** | [**Policy**](Policy.md) |  | 

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


## SearchPolicies

> SearchPolicies(ctx).Q(q).Execute()

Search List of Policies



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
	r, err := apiClient.PoliciesAPI.SearchPolicies(context.Background()).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.SearchPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Search query | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

