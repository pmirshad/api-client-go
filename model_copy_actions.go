/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.32
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type CopyActions string

// List of CopyActions
const (
	UPDATE_ON_CopyActions CopyActions = "updateOn"
	UPDATE_PREREQUISITES_CopyActions CopyActions = "updatePrerequisites"
	UPDATE_TARGETS_CopyActions CopyActions = "updateTargets"
	UPDATE_RULES_CopyActions CopyActions = "updateRules"
	UPDATE_FALLTHROUGH_CopyActions CopyActions = "updateFallthrough"
	UPDATE_OFF_VARIATION_CopyActions CopyActions = "updateOffVariation"
)
