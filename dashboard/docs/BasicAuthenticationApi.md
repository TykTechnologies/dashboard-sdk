# {{classname}}

All URIs are relative to *http://localhost:3000/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBasicAuthUser**](BasicAuthenticationApi.md#CreateBasicAuthUser) | **Post** /api/apis/keys/basic/{username} | Create a Basic Auth User
[**DeleteBasicAuthUser**](BasicAuthenticationApi.md#DeleteBasicAuthUser) | **Delete** /api/apis/keys/basic/{username} | Delete a Basic Auth User
[**GetBasicAuthUser**](BasicAuthenticationApi.md#GetBasicAuthUser) | **Get** /api/apis/keys/basic/{username} | Get a Basic Auth User
[**UpdateBasicAuthUser**](BasicAuthenticationApi.md#UpdateBasicAuthUser) | **Put** /api/apis/keys/basic/{username} | Update a Basic Auth User

# **CreateBasicAuthUser**
> SessionState CreateBasicAuthUser(ctx, username, optional)
Create a Basic Auth User

Create a Basic Auth user with the username specified in the path {username}.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| Username of Basic Auth user to create, get, update, or delete. | 
 **optional** | ***BasicAuthenticationApiCreateBasicAuthUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BasicAuthenticationApiCreateBasicAuthUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SessionState**](SessionState.md)|  | 

### Return type

[**SessionState**](SessionState.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBasicAuthUser**
> SessionState DeleteBasicAuthUser(ctx, username)
Delete a Basic Auth User

Delete the Basic Auth user specified in the path {username}.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| Username of Basic Auth user to create, get, update, or delete. | 

### Return type

[**SessionState**](SessionState.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBasicAuthUser**
> SessionState GetBasicAuthUser(ctx, username)
Get a Basic Auth User

Get the Basic Auth user specified in the path {username}.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| Username of Basic Auth user to create, get, update, or delete. | 

### Return type

[**SessionState**](SessionState.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBasicAuthUser**
> SessionState UpdateBasicAuthUser(ctx, username, optional)
Update a Basic Auth User

Update the Basic Auth user specified in the path {username}.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| Username of Basic Auth user to create, get, update, or delete. | 
 **optional** | ***BasicAuthenticationApiUpdateBasicAuthUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BasicAuthenticationApiUpdateBasicAuthUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SessionState**](SessionState.md)|  | 

### Return type

[**SessionState**](SessionState.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

