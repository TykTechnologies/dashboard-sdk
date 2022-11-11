# {{classname}}

All URIs are relative to *http://localhost:3000/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePolicy**](PoliciesApi.md#DeletePolicy) | **Delete** /api/portal/policies/{id} | Delete a single Policy by ID
[**GetPolicies**](PoliciesApi.md#GetPolicies) | **Get** /api/portal/policies/ | Get List of Policies
[**GetPolicy**](PoliciesApi.md#GetPolicy) | **Get** /api/portal/policies/{id} | Get a single Policy by ID
[**PostPolicies**](PoliciesApi.md#PostPolicies) | **Post** /api/portal/policies/ | Create Policy Definition
[**PutPolicies**](PoliciesApi.md#PutPolicies) | **Put** /api/portal/policies/{id} | Update Policy Definition
[**SearchPolicies**](PoliciesApi.md#SearchPolicies) | **Get** /api/portal/policies/search | Search List of Policies

# **DeletePolicy**
> ApiModifyKeySuccess DeletePolicy(ctx, id)
Delete a single Policy by ID

Delete a Policy by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of Policy to delete. | 

### Return type

[**ApiModifyKeySuccess**](apiModifyKeySuccess.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicies**
> []Policy GetPolicies(ctx, )
Get List of Policies

Get List of Policies

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Policy**](Policy.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicy**
> Policy GetPolicy(ctx, id)
Get a single Policy by ID

Get a Policy by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of Policy to get. | 

### Return type

[**Policy**](Policy.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPolicies**
> ApiModifyKeySuccess PostPolicies(ctx, optional)
Create Policy Definition

Creating policy definitions is slightly different to the core API, API definitions are wrapped inside an api_definition field and event handlers, such as webhooks are not embedded in the main api_defintion object (though they can be), webhooks are instead appended as references into the hook_references field, the API will embed the correct webhook data into the event handler interface.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PoliciesApiPostPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PoliciesApiPostPoliciesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Policy**](Policy.md)|  | 

### Return type

[**ApiModifyKeySuccess**](apiModifyKeySuccess.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPolicies**
> ApiModifyKeySuccess PutPolicies(ctx, id, optional)
Update Policy Definition

Update a Policy Definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of Policy to update. | 
 **optional** | ***PoliciesApiPutPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PoliciesApiPutPoliciesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Policy**](Policy.md)|  | 

### Return type

[**ApiModifyKeySuccess**](apiModifyKeySuccess.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchPolicies**
> SearchPolicies(ctx, optional)
Search List of Policies

Search List of Policies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PoliciesApiSearchPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PoliciesApiSearchPoliciesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Search query | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

