# {{classname}}

All URIs are relative to *http://localhost:3000/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAdditionalPermissions**](AdditionalPermissionsApi.md#ListAdditionalPermissions) | **Get** /api/org/permissions | List additional permissions
[**UpdateAdditionalPermissions**](AdditionalPermissionsApi.md#UpdateAdditionalPermissions) | **Put** /api/org/permissions | Modify additional permissions

# **ListAdditionalPermissions**
> InlineResponse2001 ListAdditionalPermissions(ctx, )
List additional permissions

This API helps you to add and delete (CRUD) a list of additional (custom) permissions for your Dashboard users. Once created, a custom permission will be added to standard list of user permissions. Only Admin Dashboard users will be authorised to use this API.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAdditionalPermissions**
> ApiStatusMessage UpdateAdditionalPermissions(ctx, optional)
Modify additional permissions

Whenever you want to add/update/delete an additional permission, just send back the updated list of permissions, through a PUT request to the API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AdditionalPermissionsApiUpdateAdditionalPermissionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdditionalPermissionsApiUpdateAdditionalPermissionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OrgPermissionsBody**](OrgPermissionsBody.md)|  | 

### Return type

[**ApiStatusMessage**](apiStatusMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

