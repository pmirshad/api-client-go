/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.9
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type Statement struct {

	Resources *Resources `json:"resources,omitempty"`

	// Targeted resource will be those resources NOT in this list. The \"resources`\" field must be empty to use this field.
	Notresources *Resources `json:"notresources,omitempty"`

	Actions *Actions `json:"actions,omitempty"`

	// Targeted actions will be those actions NOT in this list. The \"actions\" field must be empty to use this field.
	Notactions *Actions `json:"notactions,omitempty"`

	Effect string `json:"effect,omitempty"`
}
