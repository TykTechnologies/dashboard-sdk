# {{classname}}

All URIs are relative to *http://localhost:3000/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateAuthToken**](SingleSignOnApi.md#GenerateAuthToken) | **Post** /api/sso/ | Generate authentication token

# **GenerateAuthToken**
> ApiStatusMessage GenerateAuthToken(ctx, optional)
Generate authentication token

The Dashboard exposes the /api/sso Dashboard API which allows you to generate a temporary authentication token, valid for 60 seconds.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SingleSignOnApiGenerateAuthTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SingleSignOnApiGenerateAuthTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiSsoBody**](ApiSsoBody.md)|  | 

### Return type

[**ApiStatusMessage**](apiStatusMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

