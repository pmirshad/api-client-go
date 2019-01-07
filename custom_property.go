/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.12
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

// A name and value describing a custom property.
type CustomProperty struct {

	// The name of the property.
	Name string `json:"name"`

	Value *CustomPropertyValues `json:"value,omitempty"`
}
