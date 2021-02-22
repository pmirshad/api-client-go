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

type MembersBody struct {
	Email string `json:"email"`
	FirstName string `json:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty"`
	Role *Role `json:"role,omitempty"`
	CustomRoles []string `json:"customRoles,omitempty"`
	InlineRole []Statement `json:"inlineRole,omitempty"`
}
