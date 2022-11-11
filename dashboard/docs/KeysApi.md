# {{classname}}

All URIs are relative to *http://localhost:3000/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddKey**](KeysApi.md#AddKey) | **Post** /api/keys | Create a key
[**CreateCustomKey**](KeysApi.md#CreateCustomKey) | **Post** /api/keys/{keyID} | Create custom key
[**DeleteKey**](KeysApi.md#DeleteKey) | **Delete** /api/apis/{apiID}/keys/{keyID} | Delete key
[**GetKey**](KeysApi.md#GetKey) | **Get** /api/apis/{apiID}/keys/{keyID} | Get key
[**ListKeys**](KeysApi.md#ListKeys) | **Get** /api/apis/{apiID}/keys | List keys
[**UpdateKey**](KeysApi.md#UpdateKey) | **Put** /api/apis/{apiID}/keys/{keyID} | Update key

# **AddKey**
> SessionState AddKey(ctx, optional)
Create a key

Tyk will generate the access token based on the OrgID specified in the API Definition and a random UUID. This ensures that keys can be \"owned\" by different API Owners should segmentation be needed at an organisational level. <br/><br/> API keys without access_rights data will be written to all APIs on the system (this also means that they will be created across all SessionHandlers and StorageHandlers, it is recommended to always embed access_rights data in a key to ensure that only targeted APIs and their back-ends are written to.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***KeysApiAddKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KeysApiAddKeyOpts struct
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

# **CreateCustomKey**
> SessionState CreateCustomKey(ctx, keyID, optional)
Create custom key

Creates a key with a custom key ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyID** | **string**| The Key ID. | 
 **optional** | ***KeysApiCreateCustomKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KeysApiCreateCustomKeyOpts struct
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

# **DeleteKey**
> ApiStatusMessage DeleteKey(ctx, apiID, keyID)
Delete key

Deleting a key will remove it permanently from the system, however analytics relating to that key will still be available.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiID** | **string**| ID of API the keys grant access to. Can either be the internal or external API ID. | 
  **keyID** | **string**| The Key ID. | 

### Return type

[**ApiStatusMessage**](apiStatusMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKey**
> SessionState GetKey(ctx, apiID, keyID)
Get key

Fetches the key that grant access to the API with the ID {apiID} and key ID {keyID}.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiID** | **string**| ID of API the keys grant access to. Can either be the internal or external API ID. | 
  **keyID** | **string**| The Key ID. | 

### Return type

[**SessionState**](SessionState.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListKeys**
> ApiAllKeys ListKeys(ctx, apiID)
List keys

Lists keys that grant access to the API with the ID {apiID}.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiID** | **string**| ID of API the keys grant access to. Can either be the internal or external API ID. | 

### Return type

[**ApiAllKeys**](apiAllKeys.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateKey**
> SessionState UpdateKey(ctx, apiID, keyID, optional)
Update key

You can also manually add keys to Tyk using your own key-generation algorithm. It is recommended if using this approach to ensure that the OrgID being used in the API Definition and the key data is blank so that Tyk does not try to prepend or manage the key in any way.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiID** | **string**| ID of API the keys grant access to. Can either be the internal or external API ID. | 
  **keyID** | **string**| The Key ID. | 
 **optional** | ***KeysApiUpdateKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KeysApiUpdateKeyOpts struct
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

