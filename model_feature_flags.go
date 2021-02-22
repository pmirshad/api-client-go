/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 5.0.2
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type FeatureFlags struct {
	Links *Links `json:"_links,omitempty"`
	Items []FeatureFlag `json:"items,omitempty"`
	TotalCount float32 `json:"totalCount,omitempty"`
}
