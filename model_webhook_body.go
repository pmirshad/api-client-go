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

type WebhookBody struct {
	// The URL of the remote webhook.
	Url string `json:"url"`
	// If sign is true, and the secret attribute is omitted, LaunchDarkly will automatically generate a secret for you.
	Secret string `json:"secret,omitempty"`
	// If sign is false, the webhook will not include a signature header, and the secret can be omitted.
	Sign bool `json:"sign"`
	// Whether this webhook is enabled or not.
	On bool `json:"on"`
	// The name of the webhook.
	Name string `json:"name,omitempty"`
	Statements []Statement `json:"statements,omitempty"`
	// Tags for the webhook.
	Tags []string `json:"tags,omitempty"`
}
