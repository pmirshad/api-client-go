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

type AuditLogEntries struct {
	Links *Links `json:"_links,omitempty"`
	Items []AuditLogEntry `json:"items,omitempty"`
}
