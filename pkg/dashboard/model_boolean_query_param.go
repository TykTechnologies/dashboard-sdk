/*
Tyk Dashboard API

## <a name=\"introduction\"></a> Introduction  The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.  A common question around using a database-backed configuration is how to programatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.  The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).  ## <a name=\"security-hierarchy\"></a> Security Hierarchy  The Dashboard API provides a more structured security layer to managing Tyk nodes.  ### Organisations, APIs and Users  With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming an creating definitions), the following security model is applied to the management of Upstream APIs:  * **Organisations**: All APIs are *owned* by an organisation, this is designated by the `OrgID` parameter in the API Definition. * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access). * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations. * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation. * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.  In order to use the Dashboard API, you'll need to get the `Tyk Dashboard API Access Credentials` secret from your user profile on the Dashboard UI.  The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:  ``` authorization: <your-secret> ```

API version: 5.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboard

import (
	"encoding/json"
	"fmt"
)

// BooleanQueryParam the model 'BooleanQueryParam'
type BooleanQueryParam string

// List of BooleanQueryParam
const (
	TRUE  BooleanQueryParam = "true"
	FALSE BooleanQueryParam = "false"
)

// All allowed values of BooleanQueryParam enum
var AllowedBooleanQueryParamEnumValues = []BooleanQueryParam{
	"true",
	"false",
}

func (v *BooleanQueryParam) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BooleanQueryParam(value)
	for _, existing := range AllowedBooleanQueryParamEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BooleanQueryParam", value)
}

// NewBooleanQueryParamFromValue returns a pointer to a valid BooleanQueryParam
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBooleanQueryParamFromValue(v string) (*BooleanQueryParam, error) {
	ev := BooleanQueryParam(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BooleanQueryParam: valid values are %v", v, AllowedBooleanQueryParamEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BooleanQueryParam) IsValid() bool {
	for _, existing := range AllowedBooleanQueryParamEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BooleanQueryParam value
func (v BooleanQueryParam) Ptr() *BooleanQueryParam {
	return &v
}

type NullableBooleanQueryParam struct {
	value *BooleanQueryParam
	isSet bool
}

func (v NullableBooleanQueryParam) Get() *BooleanQueryParam {
	return v.value
}

func (v *NullableBooleanQueryParam) Set(val *BooleanQueryParam) {
	v.value = val
	v.isSet = true
}

func (v NullableBooleanQueryParam) IsSet() bool {
	return v.isSet
}

func (v *NullableBooleanQueryParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBooleanQueryParam(val *BooleanQueryParam) *NullableBooleanQueryParam {
	return &NullableBooleanQueryParam{value: val, isSet: true}
}

func (v NullableBooleanQueryParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBooleanQueryParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
