/*
Tyk Dashboard API

 ## <a name=\"introduction\"></a> Introduction  The Tyk Dashboard API offers granular, programmatic access to a centralised database of resources that your Tyk nodes can pull from. This API has a dynamic user administrative structure which means the secret key that is used to communicate with your Tyk nodes can be kept secret and access to the wider management functions can be handled on a user-by-user and organisation-by-organisation basis.  A common question around using a database-backed configuration is how to programmatically add API definitions to your Tyk nodes, the Dashboard API allows much more fine-grained, secure and multi-user access to your Tyk cluster, and should be used to manage a database-backed Tyk node.  The Tyk Dashboard API works seamlessly with the Tyk Dashboard (and the two come bundled together).  ## <a name=\"security-hierarchy\"></a> Security Hierarchy  The Dashboard API provides a more structured security layer to managing Tyk nodes.  ### Organisations, APIs and Users  With the Dashboard API and a database-backed Tyk setup, (and to an extent with file-based API setups - if diligence is used in naming and creating definitions), the following security model is applied to the management of Upstream APIs:  * **Organisations**: All APIs are *owned* by an organisation, this is designated by the 'OrgID' parameter in the API Definition. * **Users**: All users created in the Dashboard belong to an organisation (unless an exception is made for super-administrative access). * **APIs**: All APIs belong to an Organisation and only Users that belong to that organisation can see the analytics for those APIs and manage their configurations. * **API Keys**: API Keys are designated by organisation, this means an API key that has full access rights will not be allowed to access the APIs of another organisation on the same system, but can have full access to all APIs within the organisation. * **Access Rights**: Access rights are stored with the key, this enables a key to give access to multiple APIs, this is defined by the session object in the core Tyk API.  In order to use the Dashboard API, you'll need to get the 'Tyk Dashboard API Access Credentials' secret from your user profile on the Dashboard UI.  The secret you set should then be sent along as a header with each Dashboard API Request in order for it to be successful:   authorization: <your-secret>

API version: 5.7.1
Contact: support@tyk.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboard

import (
	"encoding/json"
)

// checks if the Server type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Server{}

// Server struct for Server
type Server struct {
	Authentication       NullableAuthentication       `json:"authentication,omitempty"`
	ClientCertificates   NullableClientCertificates   `json:"clientCertificates,omitempty"`
	CustomDomain         NullableDomain               `json:"customDomain,omitempty"`
	DetailedActivityLogs NullableDetailedActivityLogs `json:"detailedActivityLogs,omitempty"`
	DetailedTracing      NullableDetailedTracing      `json:"detailedTracing,omitempty"`
	EventHandlers        []EventHandler               `json:"eventHandlers,omitempty"`
	GatewayTags          NullableGatewayTags          `json:"gatewayTags,omitempty"`
	ListenPath           *ListenPath                  `json:"listenPath,omitempty"`
}

// NewServer instantiates a new Server object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServer() *Server {
	this := Server{}
	return &this
}

// NewServerWithDefaults instantiates a new Server object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerWithDefaults() *Server {
	this := Server{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Server) GetAuthentication() Authentication {
	if o == nil || IsNil(o.Authentication.Get()) {
		var ret Authentication
		return ret
	}
	return *o.Authentication.Get()
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Server) GetAuthenticationOk() (*Authentication, bool) {
	if o == nil {
		return nil, false
	}
	return o.Authentication.Get(), o.Authentication.IsSet()
}

// HasAuthentication returns a boolean if a field has been set.
func (o *Server) HasAuthentication() bool {
	if o != nil && o.Authentication.IsSet() {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given NullableAuthentication and assigns it to the Authentication field.
func (o *Server) SetAuthentication(v Authentication) {
	o.Authentication.Set(&v)
}

// SetAuthenticationNil sets the value for Authentication to be an explicit nil
func (o *Server) SetAuthenticationNil() {
	o.Authentication.Set(nil)
}

// UnsetAuthentication ensures that no value is present for Authentication, not even an explicit nil
func (o *Server) UnsetAuthentication() {
	o.Authentication.Unset()
}

// GetClientCertificates returns the ClientCertificates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Server) GetClientCertificates() ClientCertificates {
	if o == nil || IsNil(o.ClientCertificates.Get()) {
		var ret ClientCertificates
		return ret
	}
	return *o.ClientCertificates.Get()
}

// GetClientCertificatesOk returns a tuple with the ClientCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Server) GetClientCertificatesOk() (*ClientCertificates, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientCertificates.Get(), o.ClientCertificates.IsSet()
}

// HasClientCertificates returns a boolean if a field has been set.
func (o *Server) HasClientCertificates() bool {
	if o != nil && o.ClientCertificates.IsSet() {
		return true
	}

	return false
}

// SetClientCertificates gets a reference to the given NullableClientCertificates and assigns it to the ClientCertificates field.
func (o *Server) SetClientCertificates(v ClientCertificates) {
	o.ClientCertificates.Set(&v)
}

// SetClientCertificatesNil sets the value for ClientCertificates to be an explicit nil
func (o *Server) SetClientCertificatesNil() {
	o.ClientCertificates.Set(nil)
}

// UnsetClientCertificates ensures that no value is present for ClientCertificates, not even an explicit nil
func (o *Server) UnsetClientCertificates() {
	o.ClientCertificates.Unset()
}

// GetCustomDomain returns the CustomDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Server) GetCustomDomain() Domain {
	if o == nil || IsNil(o.CustomDomain.Get()) {
		var ret Domain
		return ret
	}
	return *o.CustomDomain.Get()
}

// GetCustomDomainOk returns a tuple with the CustomDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Server) GetCustomDomainOk() (*Domain, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomDomain.Get(), o.CustomDomain.IsSet()
}

// HasCustomDomain returns a boolean if a field has been set.
func (o *Server) HasCustomDomain() bool {
	if o != nil && o.CustomDomain.IsSet() {
		return true
	}

	return false
}

// SetCustomDomain gets a reference to the given NullableDomain and assigns it to the CustomDomain field.
func (o *Server) SetCustomDomain(v Domain) {
	o.CustomDomain.Set(&v)
}

// SetCustomDomainNil sets the value for CustomDomain to be an explicit nil
func (o *Server) SetCustomDomainNil() {
	o.CustomDomain.Set(nil)
}

// UnsetCustomDomain ensures that no value is present for CustomDomain, not even an explicit nil
func (o *Server) UnsetCustomDomain() {
	o.CustomDomain.Unset()
}

// GetDetailedActivityLogs returns the DetailedActivityLogs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Server) GetDetailedActivityLogs() DetailedActivityLogs {
	if o == nil || IsNil(o.DetailedActivityLogs.Get()) {
		var ret DetailedActivityLogs
		return ret
	}
	return *o.DetailedActivityLogs.Get()
}

// GetDetailedActivityLogsOk returns a tuple with the DetailedActivityLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Server) GetDetailedActivityLogsOk() (*DetailedActivityLogs, bool) {
	if o == nil {
		return nil, false
	}
	return o.DetailedActivityLogs.Get(), o.DetailedActivityLogs.IsSet()
}

// HasDetailedActivityLogs returns a boolean if a field has been set.
func (o *Server) HasDetailedActivityLogs() bool {
	if o != nil && o.DetailedActivityLogs.IsSet() {
		return true
	}

	return false
}

// SetDetailedActivityLogs gets a reference to the given NullableDetailedActivityLogs and assigns it to the DetailedActivityLogs field.
func (o *Server) SetDetailedActivityLogs(v DetailedActivityLogs) {
	o.DetailedActivityLogs.Set(&v)
}

// SetDetailedActivityLogsNil sets the value for DetailedActivityLogs to be an explicit nil
func (o *Server) SetDetailedActivityLogsNil() {
	o.DetailedActivityLogs.Set(nil)
}

// UnsetDetailedActivityLogs ensures that no value is present for DetailedActivityLogs, not even an explicit nil
func (o *Server) UnsetDetailedActivityLogs() {
	o.DetailedActivityLogs.Unset()
}

// GetDetailedTracing returns the DetailedTracing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Server) GetDetailedTracing() DetailedTracing {
	if o == nil || IsNil(o.DetailedTracing.Get()) {
		var ret DetailedTracing
		return ret
	}
	return *o.DetailedTracing.Get()
}

// GetDetailedTracingOk returns a tuple with the DetailedTracing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Server) GetDetailedTracingOk() (*DetailedTracing, bool) {
	if o == nil {
		return nil, false
	}
	return o.DetailedTracing.Get(), o.DetailedTracing.IsSet()
}

// HasDetailedTracing returns a boolean if a field has been set.
func (o *Server) HasDetailedTracing() bool {
	if o != nil && o.DetailedTracing.IsSet() {
		return true
	}

	return false
}

// SetDetailedTracing gets a reference to the given NullableDetailedTracing and assigns it to the DetailedTracing field.
func (o *Server) SetDetailedTracing(v DetailedTracing) {
	o.DetailedTracing.Set(&v)
}

// SetDetailedTracingNil sets the value for DetailedTracing to be an explicit nil
func (o *Server) SetDetailedTracingNil() {
	o.DetailedTracing.Set(nil)
}

// UnsetDetailedTracing ensures that no value is present for DetailedTracing, not even an explicit nil
func (o *Server) UnsetDetailedTracing() {
	o.DetailedTracing.Unset()
}

// GetEventHandlers returns the EventHandlers field value if set, zero value otherwise.
func (o *Server) GetEventHandlers() []EventHandler {
	if o == nil || IsNil(o.EventHandlers) {
		var ret []EventHandler
		return ret
	}
	return o.EventHandlers
}

// GetEventHandlersOk returns a tuple with the EventHandlers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetEventHandlersOk() ([]EventHandler, bool) {
	if o == nil || IsNil(o.EventHandlers) {
		return nil, false
	}
	return o.EventHandlers, true
}

// HasEventHandlers returns a boolean if a field has been set.
func (o *Server) HasEventHandlers() bool {
	if o != nil && !IsNil(o.EventHandlers) {
		return true
	}

	return false
}

// SetEventHandlers gets a reference to the given []EventHandler and assigns it to the EventHandlers field.
func (o *Server) SetEventHandlers(v []EventHandler) {
	o.EventHandlers = v
}

// GetGatewayTags returns the GatewayTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Server) GetGatewayTags() GatewayTags {
	if o == nil || IsNil(o.GatewayTags.Get()) {
		var ret GatewayTags
		return ret
	}
	return *o.GatewayTags.Get()
}

// GetGatewayTagsOk returns a tuple with the GatewayTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Server) GetGatewayTagsOk() (*GatewayTags, bool) {
	if o == nil {
		return nil, false
	}
	return o.GatewayTags.Get(), o.GatewayTags.IsSet()
}

// HasGatewayTags returns a boolean if a field has been set.
func (o *Server) HasGatewayTags() bool {
	if o != nil && o.GatewayTags.IsSet() {
		return true
	}

	return false
}

// SetGatewayTags gets a reference to the given NullableGatewayTags and assigns it to the GatewayTags field.
func (o *Server) SetGatewayTags(v GatewayTags) {
	o.GatewayTags.Set(&v)
}

// SetGatewayTagsNil sets the value for GatewayTags to be an explicit nil
func (o *Server) SetGatewayTagsNil() {
	o.GatewayTags.Set(nil)
}

// UnsetGatewayTags ensures that no value is present for GatewayTags, not even an explicit nil
func (o *Server) UnsetGatewayTags() {
	o.GatewayTags.Unset()
}

// GetListenPath returns the ListenPath field value if set, zero value otherwise.
func (o *Server) GetListenPath() ListenPath {
	if o == nil || IsNil(o.ListenPath) {
		var ret ListenPath
		return ret
	}
	return *o.ListenPath
}

// GetListenPathOk returns a tuple with the ListenPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetListenPathOk() (*ListenPath, bool) {
	if o == nil || IsNil(o.ListenPath) {
		return nil, false
	}
	return o.ListenPath, true
}

// HasListenPath returns a boolean if a field has been set.
func (o *Server) HasListenPath() bool {
	if o != nil && !IsNil(o.ListenPath) {
		return true
	}

	return false
}

// SetListenPath gets a reference to the given ListenPath and assigns it to the ListenPath field.
func (o *Server) SetListenPath(v ListenPath) {
	o.ListenPath = &v
}

func (o Server) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Server) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Authentication.IsSet() {
		toSerialize["authentication"] = o.Authentication.Get()
	}
	if o.ClientCertificates.IsSet() {
		toSerialize["clientCertificates"] = o.ClientCertificates.Get()
	}
	if o.CustomDomain.IsSet() {
		toSerialize["customDomain"] = o.CustomDomain.Get()
	}
	if o.DetailedActivityLogs.IsSet() {
		toSerialize["detailedActivityLogs"] = o.DetailedActivityLogs.Get()
	}
	if o.DetailedTracing.IsSet() {
		toSerialize["detailedTracing"] = o.DetailedTracing.Get()
	}
	if !IsNil(o.EventHandlers) {
		toSerialize["eventHandlers"] = o.EventHandlers
	}
	if o.GatewayTags.IsSet() {
		toSerialize["gatewayTags"] = o.GatewayTags.Get()
	}
	if !IsNil(o.ListenPath) {
		toSerialize["listenPath"] = o.ListenPath
	}
	return toSerialize, nil
}

type NullableServer struct {
	value *Server
	isSet bool
}

func (v NullableServer) Get() *Server {
	return v.value
}

func (v *NullableServer) Set(val *Server) {
	v.value = val
	v.isSet = true
}

func (v NullableServer) IsSet() bool {
	return v.isSet
}

func (v *NullableServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServer(val *Server) *NullableServer {
	return &NullableServer{value: val, isSet: true}
}

func (v NullableServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
