package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	baseUrl := "http://localhost:3000"
	policyUrl := fmt.Sprintf("%s/api/portal/policies", baseUrl)
	deletePolicyUrl := fmt.Sprintf("%s/api/portal/policies/%s", baseUrl, "66d9121d5715ec0715608640")
	deleteApiUrl := fmt.Sprintf("%s/api/apis/%s", baseUrl, "a12b34c56d78e90f1234567890abcdef")
	createApiurl := fmt.Sprintf("%s/api/apis", baseUrl)
	bearerToken := "a32e0ae5408940527b2543fa85d7dc00"
	err := sendDeleteRequest(deletePolicyUrl, bearerToken)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
		return
	}
	err = sendDeleteRequest(deleteApiUrl, bearerToken)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
		return
	}
	err = sendPostRequest(createApiurl, miniApiData, bearerToken)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
		return
	}
	err = sendPostRequest(policyUrl, miniPolicy, bearerToken)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
		return
	}
}

func sendPostRequest(url string, jsonData string, bearerToken string) error {
	// Convert string to byte array
	jsonBytes := []byte(jsonData)

	// Create a new POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	// Set content type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Set the Authorization header with the Bearer token
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

	// Send the request using the default client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check for a successful response
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error response: %s", string(body))
	}

	// Print response (if needed)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response: ", string(body))

	return nil
}

func sendDeleteRequest(url string, bearerToken string) error {
	// Create a new DELETE request
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	// Set the Authorization header with the Bearer token
	if bearerToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))
	}

	// Send the request using the default client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check for a successful response
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotFound {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error response: %s", string(body))
	}

	// Print response (if needed)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response: ", string(body))

	return nil
}

var miniPolicy = `{
  "access_rights": {
    "Tyk Test API": {
      "allowed_urls": [
        {
          "methods": [
            "GET"
          ],
          "url": "/users"
        }
      ],
      "api_id": "a12b34c56d78e90f1234567890abcdef",
      "api_name": "Tyk Test API",
      "disable_introspection": false,
      "versions": [
        "Default"
      ]
    }
  },
  "id":"66d9121d5715ec0715608640",
  "_id":"66d9121d5715ec0715608640",
  "active": true,
  "hmac_enabled": false,
  "is_inactive": false,
  "key_expires_in": 2592000,
  "max_query_depth": -1,
  "meta_data": {
    "email": "itachi@tyk.io",
    "user_type": "mobile_user"
  },
  "name": "Itachi sasuke testing policy items",
  "partitions": {
    "acl": true,
    "complexity": false,
    "per_api": false,
    "quota": true,
    "rate_limit": true
  },
  "per": 60,
  "quota_max": 10000,
  "quota_renewal_rate": 3600,
  "rate": 1000,
  "tags": [
    "security"
  ],
  "throttle_interval": 10,
  "throttle_retry_limit": 10
}`

var miniApiData = `{
  "api_definition": {
    "active": true,
    "api_id": "a12b34c56d78e90f1234567890abcdef",
    "auth": {
      "auth_header_name": "authorization"
    },
    "definition": {
      "key": "version",
      "location": "header"
    },
    "name": "Tyk Test API",
    "proxy": {
      "listen_path": "/tyk-api-test/",
      "strip_listen_path": true,
      "target_url": "https://httpbin.org"
    },
    "use_oauth2": true,
    "version_data": {
      "not_versioned": true,
      "versions": {
        "Default": {
          "name": "Default"
        }
      }
    }
  }
}`
