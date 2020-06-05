/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 3.3.2
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type UserTargetingExpirationForSegment struct {
	// Unix epoch time in milliseconds specifying the expiration date
	ExpirationDate int64 `json:"expirationDate,omitempty"`
	// either the included or excluded variation that the user is targeted on a segment
	TargetType string `json:"targetType,omitempty"`
	// Unique identifier for the user
	UserKey string `json:"userKey,omitempty"`
	Id string `json:"_id,omitempty"`
	ResourceId *UserTargetingExpirationResourceIdForFlag `json:"_resourceId,omitempty"`
	Links *Links `json:"_links,omitempty"`
	Version int32 `json:"_version,omitempty"`
}
