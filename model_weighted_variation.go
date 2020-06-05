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

type WeightedVariation struct {
	Variation int32 `json:"variation,omitempty"`
	Weight int32 `json:"weight,omitempty"`
}
