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

type ClientSideAvailability struct {
	// When set to true, this flag will be available to SDKs using the client-side id.
	UsingEnvironmentId bool `json:"usingEnvironmentId,omitempty"`
	// When set to true, this flag will be available to SDKS using a mobile key.
	UsingMobileKey bool `json:"usingMobileKey,omitempty"`
}
