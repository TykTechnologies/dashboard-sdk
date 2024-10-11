# \KeysAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddKey**](KeysAPI.md#AddKey) | **Post** /api/keys | Create a key.
[**CreateCustomKey**](KeysAPI.md#CreateCustomKey) | **Post** /api/keys/{keyId} | Create custom key.
[**DeleteApiKeyDetail**](KeysAPI.md#DeleteApiKeyDetail) | **Delete** /api/apis/{apiID}/keys/{keyID} | Delete key with API ID and key ID.
[**DeleteKey**](KeysAPI.md#DeleteKey) | **Delete** /api/keys/{keyId} | Delete key.
[**GetApiKeyDetail**](KeysAPI.md#GetApiKeyDetail) | **Get** /api/apis/{apiID}/keys/{keyID} | Get key details with API ID and key ID.
[**GetKeyDetail**](KeysAPI.md#GetKeyDetail) | **Get** /api/keys/{keyId} | Get key Details.
[**GetKeysDetailed**](KeysAPI.md#GetKeysDetailed) | **Get** /api/keys/detailed | List All the Keys info.
[**ListApiKeys**](KeysAPI.md#ListApiKeys) | **Get** /api/apis/{apiID}/keys | List keys by API.
[**ListKeys**](KeysAPI.md#ListKeys) | **Get** /api/apis/keys | List All the keys.
[**SearchKeys**](KeysAPI.md#SearchKeys) | **Get** /api/apis/{apiId}/keys/search | Search keys by API.
[**UpdateApiKeyDetail**](KeysAPI.md#UpdateApiKeyDetail) | **Put** /api/apis/{apiID}/keys/{keyID} | With API ID and key ID.
[**UpdateKeyDetail**](KeysAPI.md#UpdateKeyDetail) | **Put** /api/keys/{keyId} | Update key.
[**ValidateAKeyDefinition**](KeysAPI.md#ValidateAKeyDefinition) | **Post** /api/keys/preview | This will validate a key definition.



## AddKey

> KeyData AddKey(ctx).BasicAuth(basicAuth).SessionState(sessionState).Execute()

Create a key.



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
	basicAuth := true // bool | Set this to true to create a basic user. Note you have to send basic_auth_data(user and password) in the request body if this value is set to true. (optional) (default to false)
	sessionState := *openapiclient.NewSessionState() // SessionState |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.AddKey(context.Background()).BasicAuth(basicAuth).SessionState(sessionState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.AddKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddKey`: KeyData
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.AddKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **basicAuth** | **bool** | Set this to true to create a basic user. Note you have to send basic_auth_data(user and password) in the request body if this value is set to true. | [default to false]
 **sessionState** | [**SessionState**](SessionState.md) |  | 

### Return type

[**KeyData**](KeyData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomKey

> KeyData CreateCustomKey(ctx, keyId).SessionState(sessionState).Execute()

Create custom key.



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
	keyId := "my-custom-key" // string | The ID to give the key.
	sessionState := *openapiclient.NewSessionState() // SessionState |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.CreateCustomKey(context.Background(), keyId).SessionState(sessionState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.CreateCustomKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomKey`: KeyData
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.CreateCustomKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | The ID to give the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sessionState** | [**SessionState**](SessionState.md) |  | 

### Return type

[**KeyData**](KeyData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiKeyDetail

> ApiResponse DeleteApiKeyDetail(ctx, apiID, keyID).AutoGuess(autoGuess).Hashed(hashed).Username(username).Execute()

Delete key with API ID and key ID.



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
	apiID := "546e885199c947147e7f39b7d6e8e984" // string | ID of API the keys grant access to. Can either be the internal or external API ID.
	keyID := "5e9d9544a1dcd60001d0ed20e7f75f9e03534825b7aef9df749582e5" // string | The key ID
	autoGuess := true // bool | If you are not sure if a key is hashed you can send this as true. (optional) (default to false)
	hashed := "1" // string | Use the hash of the key as input instead of the full key. Any none empty string will be interpreted as to say you want to use hash input. (optional)
	username := true // bool | Set to true if the passed key is a username. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.DeleteApiKeyDetail(context.Background(), apiID, keyID).AutoGuess(autoGuess).Hashed(hashed).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.DeleteApiKeyDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApiKeyDetail`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.DeleteApiKeyDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | ID of API the keys grant access to. Can either be the internal or external API ID. | 
**keyID** | **string** | The key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiKeyDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **autoGuess** | **bool** | If you are not sure if a key is hashed you can send this as true. | [default to false]
 **hashed** | **string** | Use the hash of the key as input instead of the full key. Any none empty string will be interpreted as to say you want to use hash input. | 
 **username** | **bool** | Set to true if the passed key is a username. | 

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


## DeleteKey

> ApiResponse DeleteKey(ctx, keyId).AutoGuess(autoGuess).Hashed(hashed).Username(username).Execute()

Delete key.



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
	keyId := "5e9d9544a1dcd60001d0ed20e7f75f9e03534825b7aef9df749582e5" // string | The ID of the key.
	autoGuess := true // bool | If you are not sure if a key is hashed you can send this as true. (optional) (default to false)
	hashed := "1" // string | Use the hash of the key as input instead of the full key.Any none empty string will be interpreted as to say you want to use hash input. (optional)
	username := true // bool | Set to true if the passed key is a username (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.DeleteKey(context.Background(), keyId).AutoGuess(autoGuess).Hashed(hashed).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.DeleteKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteKey`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.DeleteKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | The ID of the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoGuess** | **bool** | If you are not sure if a key is hashed you can send this as true. | [default to false]
 **hashed** | **string** | Use the hash of the key as input instead of the full key.Any none empty string will be interpreted as to say you want to use hash input. | 
 **username** | **bool** | Set to true if the passed key is a username | 

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


## GetApiKeyDetail

> KeyData GetApiKeyDetail(ctx, apiID, keyID).AutoGuess(autoGuess).Hashed(hashed).Username(username).Execute()

Get key details with API ID and key ID.



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
	apiID := "546e885199c947147e7f39b7d6e8e984" // string | ID of API the keys grant access to. Can either be the internal or external API ID.
	keyID := "5e9d9544a1dcd60001d0ed20e7f75f9e03534825b7aef9df749582e5" // string | The Key ID.
	autoGuess := true // bool | If you are not sure if a key is hashed you can send this as true. (optional) (default to false)
	hashed := "1" // string | Use the hash of the key as input instead of the full key.Any none empty string will be interpreted as to say you want to use hash input. (optional)
	username := true // bool | Set to true if the passed key ID is a username. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.GetApiKeyDetail(context.Background(), apiID, keyID).AutoGuess(autoGuess).Hashed(hashed).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.GetApiKeyDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiKeyDetail`: KeyData
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.GetApiKeyDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | ID of API the keys grant access to. Can either be the internal or external API ID. | 
**keyID** | **string** | The Key ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeyDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **autoGuess** | **bool** | If you are not sure if a key is hashed you can send this as true. | [default to false]
 **hashed** | **string** | Use the hash of the key as input instead of the full key.Any none empty string will be interpreted as to say you want to use hash input. | 
 **username** | **bool** | Set to true if the passed key ID is a username. | 

### Return type

[**KeyData**](KeyData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeyDetail

> KeyData GetKeyDetail(ctx, keyId).AutoGuess(autoGuess).Hashed(hashed).Username(username).Execute()

Get key Details.



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
	keyId := "5e9d9544a1dcd60001d0ed20e7f75f9e03534825b7aef9df749582e5" // string | The ID of the key.
	autoGuess := true // bool | If you are not sure if a key is hashed you can send this as true. (optional) (default to false)
	hashed := "1" // string | Use the hash of the key as input instead of the full key.Any none empty string will be interpreted as to say you want to use hash input. (optional)
	username := true // bool | Set to true if the passed key ID is a username. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.GetKeyDetail(context.Background(), keyId).AutoGuess(autoGuess).Hashed(hashed).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.GetKeyDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKeyDetail`: KeyData
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.GetKeyDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | The ID of the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoGuess** | **bool** | If you are not sure if a key is hashed you can send this as true. | [default to false]
 **hashed** | **string** | Use the hash of the key as input instead of the full key.Any none empty string will be interpreted as to say you want to use hash input. | 
 **username** | **bool** | Set to true if the passed key ID is a username. | 

### Return type

[**KeyData**](KeyData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeysDetailed

> KeysDetailed GetKeysDetailed(ctx).Q(q).Execute()

List All the Keys info.



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
	q := "itachi" // string | Filter and return all keys that contain this text in there key ID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.GetKeysDetailed(context.Background()).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.GetKeysDetailed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKeysDetailed`: KeysDetailed
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.GetKeysDetailed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKeysDetailedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Filter and return all keys that contain this text in there key ID. | 

### Return type

[**KeysDetailed**](KeysDetailed.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiKeys

> Keys ListApiKeys(ctx, apiID).P(p).Execute()

List keys by API.



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
	apiID := "663a4ed9b6be920001b191ae" // string | ID of the API.
	p := int32(1) // int32 | Use p query parameter to say which page you want returned. Send number less than 0 to return all items. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.ListApiKeys(context.Background(), apiID).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.ListApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApiKeys`: Keys
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.ListApiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | ID of the API. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **p** | **int32** | Use p query parameter to say which page you want returned. Send number less than 0 to return all items. | 

### Return type

[**Keys**](Keys.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKeys

> Keys ListKeys(ctx).P(p).Execute()

List All the keys.



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
	resp, r, err := apiClient.KeysAPI.ListKeys(context.Background()).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.ListKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKeys`: Keys
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.ListKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **p** | **int32** | Use p query parameter to say which page you want returned. Send number less than 0 to return all items. | 

### Return type

[**Keys**](Keys.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchKeys

> Keys SearchKeys(ctx, apiId).Q(q).P(p).Execute()

Search keys by API.



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
	apiId := "663a4ed9b6be920001b191ae" // string | ID of the API.
	q := "itachi" // string | Filter and return all keys that contain this text in there key ID. (optional)
	p := int32(1) // int32 | Use p query parameter to say which page you want returned. Send number less than 0 to return all items. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.SearchKeys(context.Background(), apiId).Q(q).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.SearchKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchKeys`: Keys
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.SearchKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiId** | **string** | ID of the API. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | Filter and return all keys that contain this text in there key ID. | 
 **p** | **int32** | Use p query parameter to say which page you want returned. Send number less than 0 to return all items. | 

### Return type

[**Keys**](Keys.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiKeyDetail

> KeyData UpdateApiKeyDetail(ctx, apiID, keyID).AutoGuess(autoGuess).Hashed(hashed).Username(username).SuppressReset(suppressReset).SessionState(sessionState).Execute()

With API ID and key ID.



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
	apiID := "546e885199c947147e7f39b7d6e8e984" // string | ID of API the keys grant access to. Can either be the internal or external API ID.
	keyID := "5e9d9544a1dcd60001d0ed20e7f75f9e03534825b7aef9df749582e5" // string | The Key ID.
	autoGuess := true // bool | If you are not sure if a key is hashed you can send this as true. (optional) (default to false)
	hashed := "1" // string | Use the hash of the key as input instead of the full key. Any none empty string will be interpreted as to say you want to use hash input. (optional)
	username := true // bool | Set to true if the passed key ID is a username. (optional)
	suppressReset := "1" // string | Adding the suppress_reset parameter and setting it to 1, will cause Tyk not to reset the quota limit that is in the current live quota manager. By default Tyk will reset the quota in the live quota manager (initialising it) when adding a key. Adding the `suppress_reset` flag to the URL parameters will avoid this behaviour. (optional)
	sessionState := *openapiclient.NewSessionState() // SessionState |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.UpdateApiKeyDetail(context.Background(), apiID, keyID).AutoGuess(autoGuess).Hashed(hashed).Username(username).SuppressReset(suppressReset).SessionState(sessionState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.UpdateApiKeyDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApiKeyDetail`: KeyData
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.UpdateApiKeyDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiID** | **string** | ID of API the keys grant access to. Can either be the internal or external API ID. | 
**keyID** | **string** | The Key ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiKeyDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **autoGuess** | **bool** | If you are not sure if a key is hashed you can send this as true. | [default to false]
 **hashed** | **string** | Use the hash of the key as input instead of the full key. Any none empty string will be interpreted as to say you want to use hash input. | 
 **username** | **bool** | Set to true if the passed key ID is a username. | 
 **suppressReset** | **string** | Adding the suppress_reset parameter and setting it to 1, will cause Tyk not to reset the quota limit that is in the current live quota manager. By default Tyk will reset the quota in the live quota manager (initialising it) when adding a key. Adding the &#x60;suppress_reset&#x60; flag to the URL parameters will avoid this behaviour. | 
 **sessionState** | [**SessionState**](SessionState.md) |  | 

### Return type

[**KeyData**](KeyData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKeyDetail

> KeyData UpdateKeyDetail(ctx, keyId).Hashed(hashed).AutoGuess(autoGuess).Username(username).SuppressReset(suppressReset).SessionState(sessionState).Execute()

Update key.



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
	keyId := "5e9d9544a1dcd60001d0ed20e7f75f9e03534825b7aef9df749582e5" // string | The ID of the key.
	hashed := "1" // string | Use the hash of the key as input instead of the full key. Any none empty string will be interpreted as to say you want to use hash input. (optional)
	autoGuess := true // bool | If you are not sure if a key is hashed you can send this as true. (optional) (default to false)
	username := true // bool | Set to true if the passed key ID is a username. (optional)
	suppressReset := "suppressReset_example" // string | Adding the suppress_reset parameter and setting it to 1, will cause Tyk not to reset the quota limit that is in the current live quota manager. By default Tyk will reset the quota in the live quota manager (initialising it) when adding a key. Adding the `suppress_reset` flag to the URL parameters will avoid this behaviour. (optional)
	sessionState := *openapiclient.NewSessionState() // SessionState |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.UpdateKeyDetail(context.Background(), keyId).Hashed(hashed).AutoGuess(autoGuess).Username(username).SuppressReset(suppressReset).SessionState(sessionState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.UpdateKeyDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKeyDetail`: KeyData
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.UpdateKeyDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | The ID of the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKeyDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hashed** | **string** | Use the hash of the key as input instead of the full key. Any none empty string will be interpreted as to say you want to use hash input. | 
 **autoGuess** | **bool** | If you are not sure if a key is hashed you can send this as true. | [default to false]
 **username** | **bool** | Set to true if the passed key ID is a username. | 
 **suppressReset** | **string** | Adding the suppress_reset parameter and setting it to 1, will cause Tyk not to reset the quota limit that is in the current live quota manager. By default Tyk will reset the quota in the live quota manager (initialising it) when adding a key. Adding the &#x60;suppress_reset&#x60; flag to the URL parameters will avoid this behaviour. | 
 **sessionState** | [**SessionState**](SessionState.md) |  | 

### Return type

[**KeyData**](KeyData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAKeyDefinition

> KeyData ValidateAKeyDefinition(ctx).SessionState(sessionState).Execute()

This will validate a key definition.



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
	sessionState := *openapiclient.NewSessionState() // SessionState |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.ValidateAKeyDefinition(context.Background()).SessionState(sessionState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.ValidateAKeyDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateAKeyDefinition`: KeyData
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.ValidateAKeyDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAKeyDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionState** | [**SessionState**](SessionState.md) |  | 

### Return type

[**KeyData**](KeyData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

