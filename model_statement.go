/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 3.0.0
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type Statement struct {
	Resources []string `json:"resources,omitempty"`
	// Targeted resource will be those resources NOT in this list. The \"resources`\" field must be empty to use this field.
	NotResources []string `json:"notResources,omitempty"`
	Actions []string `json:"actions,omitempty"`
	// Targeted actions will be those actions NOT in this list. The \"actions\" field must be empty to use this field.
	NotActions []string `json:"notActions,omitempty"`
	Effect string `json:"effect,omitempty"`
}
