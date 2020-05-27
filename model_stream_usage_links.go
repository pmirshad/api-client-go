/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 3.3.0
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type StreamUsageLinks struct {
	Parent *Link `json:"parent,omitempty"`
	Self *Link `json:"self,omitempty"`
	// The following links that are in the response.
	Subseries []Link `json:"subseries,omitempty"`
}
