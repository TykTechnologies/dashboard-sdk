# \CertificatesAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCertificates**](CertificatesAPI.md#CreateCertificates) | **Post** /api/certs | Create a certificate in Tyk org cert store.
[**DeleteCertificateDependencies**](CertificatesAPI.md#DeleteCertificateDependencies) | **Delete** /api/certs/dependencies/{certId} | Delete Certificate dependencies.
[**DeleteCertificates**](CertificatesAPI.md#DeleteCertificates) | **Delete** /api/certs/{certId} | Delete Certificate.
[**GetCertificate**](CertificatesAPI.md#GetCertificate) | **Get** /api/certs/{certId} | Get single certificate with ID.
[**GetCertificateDependencies**](CertificatesAPI.md#GetCertificateDependencies) | **Get** /api/certs/dependencies/{certId} | Get Certificate Dependencies.
[**ListCertificates**](CertificatesAPI.md#ListCertificates) | **Get** /api/certs | List Certificates.
[**ListDetailedCertificates**](CertificatesAPI.md#ListDetailedCertificates) | **Get** /api/certs/details | List all certificates details.



## CreateCertificates

> APICertificateStatusMessage CreateCertificates(ctx).File(file).Execute()

Create a certificate in Tyk org cert store.



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
	file := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.CreateCertificates(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.CreateCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCertificates`: APICertificateStatusMessage
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.CreateCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

[**APICertificateStatusMessage**](APICertificateStatusMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCertificateDependencies

> ApiResponse DeleteCertificateDependencies(ctx, certId).Execute()

Delete Certificate dependencies.



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
	certId := "5e9d9544a1dcd60001d0ed208edce514c2d0a866063550c64d6c90be99d01561ac5aa7e82b8610b7e273d37d" // string | ID of the certificate you want to delete dependencies for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.DeleteCertificateDependencies(context.Background(), certId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.DeleteCertificateDependencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCertificateDependencies`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.DeleteCertificateDependencies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certId** | **string** | ID of the certificate you want to delete dependencies for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificateDependenciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCertificates

> ApiStatusMessage DeleteCertificates(ctx, certId).Execute()

Delete Certificate.



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
	certId := "5e9d9544a1dcd60001d0ed208edce514c2d0a866063550c64d6c90be99d01561ac5aa7e82b8610b7e273d37d" // string | ID of the certificate you want to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.DeleteCertificates(context.Background(), certId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.DeleteCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCertificates`: ApiStatusMessage
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.DeleteCertificates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certId** | **string** | ID of the certificate you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiStatusMessage**](ApiStatusMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificate

> CertsCertificateMeta GetCertificate(ctx, certId).Execute()

Get single certificate with ID.



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
	certId := "5e9d9544a1dcd60001d0ed208edce514c2d0a866063550c64d6c90be99d01561ac5aa7e82b8610b7e273d37d" // string | ID of the certificate you want to fetch.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.GetCertificate(context.Background(), certId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.GetCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCertificate`: CertsCertificateMeta
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.GetCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certId** | **string** | ID of the certificate you want to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CertsCertificateMeta**](CertsCertificateMeta.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateDependencies

> CertificateDependencies GetCertificateDependencies(ctx, certId).Execute()

Get Certificate Dependencies.



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
	certId := "5e9d9544a1dcd60001d0ed208edce514c2d0a866063550c64d6c90be99d01561ac5aa7e82b8610b7e273d37d" // string | ID of the certificate you want to fetch dependencies for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.GetCertificateDependencies(context.Background(), certId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.GetCertificateDependencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCertificateDependencies`: CertificateDependencies
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.GetCertificateDependencies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certId** | **string** | ID of the certificate you want to fetch dependencies for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateDependenciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CertificateDependencies**](CertificateDependencies.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCertificates

> ListCertificates200Response ListCertificates(ctx).P(p).Mode(mode).Execute()

List Certificates.



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
	p := int32(1) // int32 | Use p query parameter to say which page you want returned. Send number less than 0 to return all items. (optional)
	mode := "detailed" // string | Set to detailed to get certificates that are more with more details (Will contains certs basic details). To retrieve a list of certificates with all the certificate details use [this endpoint](#operation/listDetailedCertificates). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.ListCertificates(context.Background()).P(p).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.ListCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCertificates`: ListCertificates200Response
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.ListCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **p** | **int32** | Use p query parameter to say which page you want returned. Send number less than 0 to return all items. | 
 **mode** | **string** | Set to detailed to get certificates that are more with more details (Will contains certs basic details). To retrieve a list of certificates with all the certificate details use [this endpoint](#operation/listDetailedCertificates). | 

### Return type

[**ListCertificates200Response**](ListCertificates200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDetailedCertificates

> CertificateDetailedList ListDetailedCertificates(ctx).P(p).Execute()

List all certificates details.



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
	p := int32(1) // int32 | Use p query parameter to say which page you want returned. Send number less than 0 to return all items. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.ListDetailedCertificates(context.Background()).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.ListDetailedCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDetailedCertificates`: CertificateDetailedList
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.ListDetailedCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDetailedCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **p** | **int32** | Use p query parameter to say which page you want returned. Send number less than 0 to return all items. | 

### Return type

[**CertificateDetailedList**](CertificateDetailedList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

