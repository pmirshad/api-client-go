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

type Rule struct {
	Id string `json:"_id,omitempty"`
	Variation int32 `json:"variation,omitempty"`
	TrackEvents bool `json:"trackEvents,omitempty"`
	Rollout *Rollout `json:"rollout,omitempty"`
	Clauses []Clause `json:"clauses,omitempty"`
}
