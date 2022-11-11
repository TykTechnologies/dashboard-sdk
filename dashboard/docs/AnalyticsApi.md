# {{classname}}

All URIs are relative to *http://localhost:3000/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAnalyticsOfApiKey**](AnalyticsApi.md#GetAnalyticsOfApiKey) | **Get** /api/activity/keys/{keyHash}/{startDay}/{startMonth}/{startYear}/{EndDay}/{EndMonth}/{EndYear} | Analytics of API Key
[**GetAnalyticsOfOauthClientId**](AnalyticsApi.md#GetAnalyticsOfOauthClientId) | **Get** /api/activity/oauthid/{OAuthClientID}/{startDay}/{startMonth}/{startYear}/{EndDay}/{EndMonth}/{EndYear} | Analytics of Oauth Client ID

# **GetAnalyticsOfApiKey**
> GetAnalyticsOfApiKey(ctx, keyHash, startDay, startMonth, startYear, endDay, endMonth, endYear)
Analytics of API Key

It returns analytics of the endpoints of all APIs called using KEY between start and end date.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyHash** | **string**| Hash of your API key. | 
  **startDay** | **string**| Day of analytics to query. | 
  **startMonth** | **string**| Month of analytics to query. | 
  **startYear** | **string**| Start year of analytics to query. | 
  **endDay** | **string**| End date of analytics to query. | 
  **endMonth** | **string**| End month of analytics to query. | 
  **endYear** | **string**| End year of analytics to query. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAnalyticsOfOauthClientId**
> GetAnalyticsOfOauthClientId(ctx, oAuthClientID, startDay, startMonth, startYear, endDay, endMonth, endYear)
Analytics of Oauth Client ID

Returns activity of all endpoints which used OAuth client `27b35a9ed46e429eb2361e440cc4005c` between October 13th 2020 and October 14th 2020.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oAuthClientID** | **string**| OAuthClientID | 
  **startDay** | **string**| Day of analytics to query. | 
  **startMonth** | **string**| Month of analytics to query. | 
  **startYear** | **string**| Start year of analytics to query. | 
  **endDay** | **string**| End date of analytics to query. | 
  **endMonth** | **string**| End month of analytics to query. | 
  **endYear** | **string**| End year of analytics to query. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

