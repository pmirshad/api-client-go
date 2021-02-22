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

type Users struct {
	Links *Links `json:"_links,omitempty"`
	TotalCount float32 `json:"totalCount,omitempty"`
	Items []UserRecord `json:"items,omitempty"`
}
