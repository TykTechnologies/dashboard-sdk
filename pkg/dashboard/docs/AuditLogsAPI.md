# \AuditLogsAPI

All URIs are relative to *https://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditLog**](AuditLogsAPI.md#GetAuditLog) | **Get** /api/audit-logs/{audit-log-id} | Retrieve single audit log
[**GetAuditLogs**](AuditLogsAPI.md#GetAuditLogs) | **Get** /api/audit-logs | List audit logs



## GetAuditLog

> AuditLog GetAuditLog(ctx, auditLogId).Execute()

Retrieve single audit log



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TykTechnologies/dashboard-sdk"
)

func main() {
	auditLogId := int32(1) // int32 | ID of the audit log record to fetch.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogsAPI.GetAuditLog(context.Background(), auditLogId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.GetAuditLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLog`: AuditLog
	fmt.Fprintf(os.Stdout, "Response from `AuditLogsAPI.GetAuditLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditLogId** | **int32** | ID of the audit log record to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuditLog**](AuditLog.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditLogs

> *os.File GetAuditLogs(ctx).P(p).User(user).Action(action).Ip(ip).Method(method).Status(status).Url(url).FromDate(fromDate).ToDate(toDate).Download(download).Type_(type_).Execute()

List audit logs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TykTechnologies/dashboard-sdk"
)

func main() {
	p := int32(1) // int32 | Use p query parameter to say which page you want returned. The size of the page is determined by the configuration option page_size of dashboard. (optional)
	user := "jhon@mail.com" // string | Filters audit logs to show only actions performed by the specified user. This parameter allows you to focus on the activity of a particular user across the system. (optional)
	action := "List APIs" // string | Filters audit logs based on the specific action performed by users. This parameter allows you to focus on particular types of activities within the system. (optional)
	ip := "127.0.0.1" // string | Filters audit logs based on the IP address from which the action originated. This parameter allows you to focus on activities from specific network locations or to investigate actions from particular IP addresses. (optional)
	method := "POST" // string | Filters audit logs based on the HTTP method used in the API request. This parameter allows you to focus on specific types of operations performed on the API. (optional)
	status := int32(200) // int32 | Filters audit logs based on the HTTP status code returned by the API in response to the request. This parameter allows you to focus on specific outcomes of API interactions. (optional)
	url := "/api/apis" // string | This parameter filters audit logs based on partially matching the accessed API endpoint's URL path. It allows searching for actions performed on related resources or sections of the API by matching any portion of the URL. The match is case-sensitive and ignores additional path segments or query parameters beyond the matched portion.   For example, if the database contains URLs like `/tib/create`, `/tib/get/1?schema=json`,  `/api/schema`, and `/schema1` searching with `url=schema` would return  `/api/schema` and `/schema1`.  (optional)
	fromDate := "01-01-1990" // string | Specifies the start date for the audit log search. If not provided, the search will include records from the earliest available date. Format DD-MM-YYY. (optional)
	toDate := "01-01-2030" // string | Specifies the end date for the audit log search. If not provided, the search will include records up to the current date and time. Format DD-MM-YYY. (optional)
	download := true // bool | Determines whether the response should be a downloadable file containing the records. If set to `true`, the API returns a file instead of a JSON list of records. When enabled, pagination is not applied, and the file will include all records that match the search criteria. (optional)
	type_ := "csv" // string | Specifies the format of the downloadable file. This parameter is only applied when `download` is set to `true`. If set to `csv`, the file content will be in CSV format; otherwise, JSON format will be used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogsAPI.GetAuditLogs(context.Background()).P(p).User(user).Action(action).Ip(ip).Method(method).Status(status).Url(url).FromDate(fromDate).ToDate(toDate).Download(download).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.GetAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLogs`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AuditLogsAPI.GetAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **p** | **int32** | Use p query parameter to say which page you want returned. The size of the page is determined by the configuration option page_size of dashboard. | 
 **user** | **string** | Filters audit logs to show only actions performed by the specified user. This parameter allows you to focus on the activity of a particular user across the system. | 
 **action** | **string** | Filters audit logs based on the specific action performed by users. This parameter allows you to focus on particular types of activities within the system. | 
 **ip** | **string** | Filters audit logs based on the IP address from which the action originated. This parameter allows you to focus on activities from specific network locations or to investigate actions from particular IP addresses. | 
 **method** | **string** | Filters audit logs based on the HTTP method used in the API request. This parameter allows you to focus on specific types of operations performed on the API. | 
 **status** | **int32** | Filters audit logs based on the HTTP status code returned by the API in response to the request. This parameter allows you to focus on specific outcomes of API interactions. | 
 **url** | **string** | This parameter filters audit logs based on partially matching the accessed API endpoint&#39;s URL path. It allows searching for actions performed on related resources or sections of the API by matching any portion of the URL. The match is case-sensitive and ignores additional path segments or query parameters beyond the matched portion.   For example, if the database contains URLs like &#x60;/tib/create&#x60;, &#x60;/tib/get/1?schema&#x3D;json&#x60;,  &#x60;/api/schema&#x60;, and &#x60;/schema1&#x60; searching with &#x60;url&#x3D;schema&#x60; would return  &#x60;/api/schema&#x60; and &#x60;/schema1&#x60;.  | 
 **fromDate** | **string** | Specifies the start date for the audit log search. If not provided, the search will include records from the earliest available date. Format DD-MM-YYY. | 
 **toDate** | **string** | Specifies the end date for the audit log search. If not provided, the search will include records up to the current date and time. Format DD-MM-YYY. | 
 **download** | **bool** | Determines whether the response should be a downloadable file containing the records. If set to &#x60;true&#x60;, the API returns a file instead of a JSON list of records. When enabled, pagination is not applied, and the file will include all records that match the search criteria. | 
 **type_** | **string** | Specifies the format of the downloadable file. This parameter is only applied when &#x60;download&#x60; is set to &#x60;true&#x60;. If set to &#x60;csv&#x60;, the file content will be in CSV format; otherwise, JSON format will be used. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

