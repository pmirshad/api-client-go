/*
LaunchDarkly REST API

API version: 2.0
*/

package ldapi

import (
	"encoding/json"
)

// SemanticPatchWithComment struct for SemanticPatchWithComment
type SemanticPatchWithComment struct {
	// Optional comment describing the update
	Comment        *string                  `json:"comment,omitempty"`
	EnvironmentKey string                   `json:"environmentKey"`
	Instructions   []map[string]interface{} `json:"instructions"`
}

// NewSemanticPatchWithComment instantiates a new SemanticPatchWithComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSemanticPatchWithComment(
	instructions []map[string]interface{},
	environmentKey string,
) *SemanticPatchWithComment {
	this := SemanticPatchWithComment{}
	this.Instructions = instructions
	this.EnvironmentKey = environmentKey
	return &this
}

// NewSemanticPatchWithCommentWithDefaults instantiates a new SemanticPatchWithComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSemanticPatchWithCommentWithDefaults() *SemanticPatchWithComment {
	this := SemanticPatchWithComment{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *SemanticPatchWithComment) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SemanticPatchWithComment) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *SemanticPatchWithComment) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *SemanticPatchWithComment) SetComment(v string) {
	o.Comment = &v
}

// GetEnvironmentKey returns the EnvironmentKey field value if set, zero value otherwise.
func (o *SemanticPatchWithComment) GetEnvironmentKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EnvironmentKey
}

// SetEnvironmentKey gets a reference to the given string and assigns it to the
// EnvironmentKey field.
func (o *SemanticPatchWithComment) SetEnvironmentKey(v string) {
	o.EnvironmentKey = v
}

// GetInstructions returns the Instructions field value
func (o *SemanticPatchWithComment) GetInstructions() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value
// and a boolean to check if the value has been set.
func (o *SemanticPatchWithComment) GetInstructionsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instructions, true
}

// SetInstructions sets field value
func (o *SemanticPatchWithComment) SetInstructions(v []map[string]interface{}) {
	o.Instructions = v
}

func (o SemanticPatchWithComment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	toSerialize["environmentKey"] = o.EnvironmentKey
	if true {
		toSerialize["instructions"] = o.Instructions
	}
	j, _ := json.Marshal(toSerialize)
	return json.Marshal(toSerialize)
}

type NullableSemanticPatchWithComment struct {
	value *SemanticPatchWithComment
	isSet bool
}

func (v NullableSemanticPatchWithComment) Get() *SemanticPatchWithComment {
	return v.value
}

func (v *NullableSemanticPatchWithComment) Set(val *SemanticPatchWithComment) {
	v.value = val
	v.isSet = true
}

func (v NullableSemanticPatchWithComment) IsSet() bool {
	return v.isSet
}

func (v *NullableSemanticPatchWithComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSemanticPatchWithComment(val *SemanticPatchWithComment) *NullableSemanticPatchWithComment {
	return &NullableSemanticPatchWithComment{value: val, isSet: true}
}

func (v NullableSemanticPatchWithComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSemanticPatchWithComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
