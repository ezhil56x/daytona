/*
Daytona Server API

Daytona Server API

API version: v0.0.0-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the ContainerRegistry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerRegistry{}

// ContainerRegistry struct for ContainerRegistry
type ContainerRegistry struct {
	Password *string `json:"password,omitempty"`
	Server   *string `json:"server,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewContainerRegistry instantiates a new ContainerRegistry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerRegistry() *ContainerRegistry {
	this := ContainerRegistry{}
	return &this
}

// NewContainerRegistryWithDefaults instantiates a new ContainerRegistry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerRegistryWithDefaults() *ContainerRegistry {
	this := ContainerRegistry{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ContainerRegistry) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistry) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ContainerRegistry) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ContainerRegistry) SetPassword(v string) {
	o.Password = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *ContainerRegistry) GetServer() string {
	if o == nil || IsNil(o.Server) {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistry) GetServerOk() (*string, bool) {
	if o == nil || IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *ContainerRegistry) HasServer() bool {
	if o != nil && !IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *ContainerRegistry) SetServer(v string) {
	o.Server = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ContainerRegistry) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistry) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ContainerRegistry) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ContainerRegistry) SetUsername(v string) {
	o.Username = &v
}

func (o ContainerRegistry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerRegistry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Server) {
		toSerialize["server"] = o.Server
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableContainerRegistry struct {
	value *ContainerRegistry
	isSet bool
}

func (v NullableContainerRegistry) Get() *ContainerRegistry {
	return v.value
}

func (v *NullableContainerRegistry) Set(val *ContainerRegistry) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerRegistry) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerRegistry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerRegistry(val *ContainerRegistry) *NullableContainerRegistry {
	return &NullableContainerRegistry{value: val, isSet: true}
}

func (v NullableContainerRegistry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerRegistry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
