# \AuditLogsAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditLogs**](AuditLogsAPI.md#GetAuditLogs) | **Get** /api/audit-logs | List audit logs



## GetAuditLogs

> AuditLogs GetAuditLogs(ctx).P(p).User(user).Action(action).Ip(ip).Method(method).Status(status).Url(url).FromDate(fromDate).ToDate(toDate).Execute()

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
	url := "/api/apis" // string | Filters audit logs based on the specific URL path of the API endpoint that was accessed. This parameter allows you to focus on actions performed on particular resources or sections of your API. (optional)
	fromDate := "01-01-1990" // string | Specifies the start date for the audit log search. If not provided, the search will include records from the earliest available date. Format DD-MM-YYY. (optional)
	toDate := "01-01-2030" // string | Specifies the end date for the audit log search. If not provided, the search will include records up to the current date and time. Format DD-MM-YYY. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogsAPI.GetAuditLogs(context.Background()).P(p).User(user).Action(action).Ip(ip).Method(method).Status(status).Url(url).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.GetAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLogs`: AuditLogs
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
 **url** | **string** | Filters audit logs based on the specific URL path of the API endpoint that was accessed. This parameter allows you to focus on actions performed on particular resources or sections of your API. | 
 **fromDate** | **string** | Specifies the start date for the audit log search. If not provided, the search will include records from the earliest available date. Format DD-MM-YYY. | 
 **toDate** | **string** | Specifies the end date for the audit log search. If not provided, the search will include records up to the current date and time. Format DD-MM-YYY. | 

### Return type

[**AuditLogs**](AuditLogs.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

