# \PoliciesAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePolicy**](PoliciesAPI.md#DeletePolicy) | **Delete** /api/portal/policies/{id} | Delete a single Policy by ID
[**GetPolicies**](PoliciesAPI.md#GetPolicies) | **Get** /api/portal/policies | List Portal Policies
[**GetPolicy**](PoliciesAPI.md#GetPolicy) | **Get** /api/portal/policies/{id} | Get a single Policy by ID
[**PostPolicies**](PoliciesAPI.md#PostPolicies) | **Post** /api/portal/policies | Create Policy Definition
[**PutPolicies**](PoliciesAPI.md#PutPolicies) | **Put** /api/portal/policies/{id} | Update Policy Definition
[**SearchPolicies**](PoliciesAPI.md#SearchPolicies) | **Get** /api/portal/policies/search | Search List of Policies



## DeletePolicy

> ApiError DeletePolicy(ctx, id).Execute()

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
	id := "66570989d98dd00001da17f1" // string | ID of Policy to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.DeletePolicy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.DeletePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePolicy`: ApiError
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.DeletePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Policy to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRequest struct via the builder pattern


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


## GetPolicies

> ReturnDataStruct GetPolicies(ctx).P(p).Active(active).Execute()

List Portal Policies



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
	p := int32(1) // int32 | Use p query parameter to say which page you want returned.Send number less than 0 to return all items (optional)
	active := "true" // string | Send any value in this query parameter to return only the active policies (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.GetPolicies(context.Background()).P(p).Active(active).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.GetPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicies`: ReturnDataStruct
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.GetPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **p** | **int32** | Use p query parameter to say which page you want returned.Send number less than 0 to return all items | 
 **active** | **string** | Send any value in this query parameter to return only the active policies | 

### Return type

[**ReturnDataStruct**](ReturnDataStruct.md)

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
	id := "66570989d98dd00001da17f1" // string | ID of Policy to get

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
**id** | **string** | ID of Policy to get | 

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

> ApiError PostPolicies(ctx).Policy(policy).Execute()

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
	// response from `PostPolicies`: ApiError
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

[**ApiError**](ApiError.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPolicies

> ApiError PutPolicies(ctx, id).Policy(policy).Execute()

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
	id := "66570989d98dd00001da17f1" // string | ID of Policy to update
	policy := *openapiclient.NewPolicy() // Policy |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PutPolicies(context.Background(), id).Policy(policy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PutPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPolicies`: ApiError
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PutPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Policy to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policy** | [**Policy**](Policy.md) |  | 

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


## SearchPolicies

> ReturnDataStruct SearchPolicies(ctx).Q(q).PolicyIds(policyIds).Active(active).State(state).Sort(sort).ApiId(apiId).AuthType(authType).P(p).Execute()

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
	q := "5eb6349543f0440001373f5c" // string | Search for policy using policy ID or name. (optional)
	policyIds := "5eb6349543f0440001373f5c,5ead7120575961000181867e" // string | A list of comma separated policy ids that you want to search for (optional)
	active := "true" // string | Send any value in this query parameter to return only the active policies (optional)
	state := "deny" // string | Return policies whose state field matches the sent value.e.g if you send state as deny policies returned are those whose state filed value is set as deny (optional)
	sort := "name" // string | Field you want to use to sort the returned policies (optional)
	apiId := "5963f8fdedee405143f5858ea17de422" // string | Comma separete list of Api IDs.Return only policy that that have the given api ids in their access_right (optional)
	authType := "authToken" // string | Return policy whose auth_type field has the given value (optional)
	p := int32(1) // int32 | Use p query parameter to say which page you want returned.Send number less than 0 to return all items (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.SearchPolicies(context.Background()).Q(q).PolicyIds(policyIds).Active(active).State(state).Sort(sort).ApiId(apiId).AuthType(authType).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.SearchPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPolicies`: ReturnDataStruct
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.SearchPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Search for policy using policy ID or name. | 
 **policyIds** | **string** | A list of comma separated policy ids that you want to search for | 
 **active** | **string** | Send any value in this query parameter to return only the active policies | 
 **state** | **string** | Return policies whose state field matches the sent value.e.g if you send state as deny policies returned are those whose state filed value is set as deny | 
 **sort** | **string** | Field you want to use to sort the returned policies | 
 **apiId** | **string** | Comma separete list of Api IDs.Return only policy that that have the given api ids in their access_right | 
 **authType** | **string** | Return policy whose auth_type field has the given value | 
 **p** | **int32** | Use p query parameter to say which page you want returned.Send number less than 0 to return all items | 

### Return type

[**ReturnDataStruct**](ReturnDataStruct.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

