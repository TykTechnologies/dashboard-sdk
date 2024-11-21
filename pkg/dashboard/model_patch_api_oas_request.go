/*
Tyk Dashboard API

 ## <a name=\"introduction\"></a> Introduction  The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.  A common question around using a database-backed configuration is how to programmatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.  The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).  ## <a name=\"security-hierarchy\"></a> Security Hierarchy  The Dashboard API provides a more structured security layer to managing Tyk nodes.  ### Organisations, APIs and Users  With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming and creating definitions), the following security model is applied to the management of Upstream APIs:  * **Organisations**: All APIs are *owned* by an organisation, this is designated by the 'OrgID' parameter in the API Definition. * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access). * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations. * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation. * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.  In order to use the Dashboard API, you'll need to get the 'Tyk Dashboard API Access Credentials' secret from your user profile on the Dashboard UI.  The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:   authorization: <your-secret>

API version: 5.7.0
Contact: support@tyk.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboard

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// PatchApiOASRequest - struct for PatchApiOASRequest
type PatchApiOASRequest struct {
	ApiImportByUrlPayload *ApiImportByUrlPayload
	CreateApiOASRequest   *CreateApiOASRequest
}

// ApiImportByUrlPayloadAsPatchApiOASRequest is a convenience function that returns ApiImportByUrlPayload wrapped in PatchApiOASRequest
func ApiImportByUrlPayloadAsPatchApiOASRequest(v *ApiImportByUrlPayload) PatchApiOASRequest {
	return PatchApiOASRequest{
		ApiImportByUrlPayload: v,
	}
}

// CreateApiOASRequestAsPatchApiOASRequest is a convenience function that returns CreateApiOASRequest wrapped in PatchApiOASRequest
func CreateApiOASRequestAsPatchApiOASRequest(v *CreateApiOASRequest) PatchApiOASRequest {
	return PatchApiOASRequest{
		CreateApiOASRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PatchApiOASRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ApiImportByUrlPayload
	err = newStrictDecoder(data).Decode(&dst.ApiImportByUrlPayload)
	if err == nil {
		jsonApiImportByUrlPayload, _ := json.Marshal(dst.ApiImportByUrlPayload)
		if string(jsonApiImportByUrlPayload) == "{}" { // empty struct
			dst.ApiImportByUrlPayload = nil
		} else {
			if err = validator.Validate(dst.ApiImportByUrlPayload); err != nil {
				dst.ApiImportByUrlPayload = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApiImportByUrlPayload = nil
	}

	// try to unmarshal data into CreateApiOASRequest
	err = newStrictDecoder(data).Decode(&dst.CreateApiOASRequest)
	if err == nil {
		jsonCreateApiOASRequest, _ := json.Marshal(dst.CreateApiOASRequest)
		if string(jsonCreateApiOASRequest) == "{}" { // empty struct
			dst.CreateApiOASRequest = nil
		} else {
			if err = validator.Validate(dst.CreateApiOASRequest); err != nil {
				dst.CreateApiOASRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.CreateApiOASRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ApiImportByUrlPayload = nil
		dst.CreateApiOASRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PatchApiOASRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PatchApiOASRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PatchApiOASRequest) MarshalJSON() ([]byte, error) {
	if src.ApiImportByUrlPayload != nil {
		return json.Marshal(&src.ApiImportByUrlPayload)
	}

	if src.CreateApiOASRequest != nil {
		return json.Marshal(&src.CreateApiOASRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PatchApiOASRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ApiImportByUrlPayload != nil {
		return obj.ApiImportByUrlPayload
	}

	if obj.CreateApiOASRequest != nil {
		return obj.CreateApiOASRequest
	}

	// all schemas are nil
	return nil
}

type NullablePatchApiOASRequest struct {
	value *PatchApiOASRequest
	isSet bool
}

func (v NullablePatchApiOASRequest) Get() *PatchApiOASRequest {
	return v.value
}

func (v *NullablePatchApiOASRequest) Set(val *PatchApiOASRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchApiOASRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchApiOASRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchApiOASRequest(val *PatchApiOASRequest) *NullablePatchApiOASRequest {
	return &NullablePatchApiOASRequest{value: val, isSet: true}
}

func (v NullablePatchApiOASRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchApiOASRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
