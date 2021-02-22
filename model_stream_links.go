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

type StreamLinks struct {
	Parent *Link `json:"parent,omitempty"`
	Self *Link `json:"self,omitempty"`
	// Links to endpoints that are in the request path.
	Subseries []Link `json:"subseries,omitempty"`
}
