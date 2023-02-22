/*
LaunchDarkly REST API

API version: 2.0
*/

package ldapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type ApiSemanticPatchFeatureFlagRequest struct {
	ctx                      context.Context
	ApiService               *FeatureFlagsApiService
	projectKey               string
	featureFlagKey           string
	semanticPatchWithComment *SemanticPatchWithComment
}

func (r ApiSemanticPatchFeatureFlagRequest) SemanticPatchWithComment(
	semanticPatchWithComment SemanticPatchWithComment,
) ApiSemanticPatchFeatureFlagRequest {
	r.semanticPatchWithComment = &semanticPatchWithComment
	return r
}

func (r ApiSemanticPatchFeatureFlagRequest) Execute() (*FeatureFlag, *http.Response, error) {
	return r.ApiService.SemanticPatchFeatureFlagExecute(r)
}

/*
SemanticPatchFeatureFlag Update feature flag

Perform a partial update to a feature flag. The request body must be a valid semantic patch or JSON patch.

### Using semantic patches on a feature flag

To make a semantic patch request, you must append `domain-model=launchdarkly.semanticpatch` to your `Content-Type` header. To learn more, read [Updates using semantic patch](/reference#updates-using-semantic-patch).

The body of a semantic patch request for updating feature flags requires an `environmentKey` in addition to `instructions` and an optional `comment`. The body of the request takes the following properties:

* `comment` (string): (Optional) A description of the update.
* `environmentKey` (string): (Required) The key of the LaunchDarkly environment.
* `instructions` (array): (Required) A list of actions the update should perform. Each action in the list must be an object with a `kind` property that indicates the instruction. If the action requires parameters, you must include those parameters as additional fields in the object.

### Instructions

Semantic patch requests support the following `kind` instructions for updating feature flags.

<details>
<summary>Click to expand instructions for turning flags on and off</summary>

#### turnFlagOff

Sets the flag's targeting state to **Off**.

For example, to turn a flag off, use this request body:

```json

	{
	  "environmentKey": "example-environment-key",
	  "instructions": [ { "kind": "turnFlagOff" } ]
	}

```

#### turnFlagOn

Sets the flag's targeting state to **On**.

For example, to turn a flag on, use this request body:

```json

	{
	  "environmentKey": "example-environment-key",
	  "instructions": [ { "kind": "turnFlagOn" } ]
	}

```

</details><br />

<details>
<summary>Click to expand instructions for working with targeting and variations</summary>

#### addClauses

Adds the given clauses to the rule indicated by `ruleId`.

##### Parameters

- `ruleId`: ID of a rule in the flag.
- `clauses`: Array of clause objects, with `attribute` (string), `op` (string), `negate` (boolean), and `values` (array of strings, numbers, or dates) properties.

#### addPrerequisite

Adds the flag indicated by `key` with variation `variationId` as a prerequisite to the flag in the path parameter.

##### Parameters

- `key`: Flag key of the prerequisite flag.
- `variationId`: ID of a variation of the prerequisite flag.

#### addRule

Adds a new targeting rule to the flag. The rule may contain `clauses` and serve the variation that `variationId` indicates, or serve a percentage rollout that `rolloutWeights` and `rolloutBucketBy` indicate.

If you set `beforeRuleId`, this adds the new rule before the indicated rule. Otherwise, adds the new rule to the end of the list.

##### Parameters

- `clauses`: Array of clause objects, with `attribute` (string), `op` (string), `negate` (boolean), and `values` (array of strings, numbers, or dates) properties.
- `beforeRuleId`: (Optional) ID of a flag rule.
- `variationId`: ID of a variation of the flag.
- `rolloutWeights`: Map of `variationId` to weight, in thousandths of a percent (0-100000).
- `rolloutBucketBy`: (Optional) User attribute.

#### addUserTargets

Adds user keys to the individual user targets for the variation that `variationId` specifies. Returns an error if this causes the flag to target the same user key in multiple variations.

##### Parameters

- `values`: List of user keys.
- `variationId`: ID of a variation on the flag.

#### addValuesToClause

Adds `values` to the values of the clause that `ruleId` and `clauseId` indicate.

##### Parameters

- `ruleId`: ID of a rule in the flag.
- `clauseId`: ID of a clause in that rule.
- `values`: Array of strings.

#### clearUserTargets

Removes all individual user targets from the variation that `variationId` specifies.

##### Parameters

- `variationId`: ID of a variation on the flag.

#### removeClauses

Removes the clauses specified by `clauseIds` from the rule indicated by `ruleId`.

##### Parameters

- `ruleId`: ID of a rule in the flag.
- `clauseIds`: Array of IDs of clauses in the rule.

#### removePrerequisite

Removes the prerequisite flag indicated by `key`. Does nothing if this prerequisite does not exist.

##### Parameters

- `key`: Flag key of an existing prerequisite flag.

#### removeRule

Removes the targeting rule specified by `ruleId`. Does nothing if the rule does not exist.

##### Parameters

- `ruleId`: ID of a rule in the flag.

#### removeUserTargets

Removes user keys from the individual user targets for the variation that `variationId` specifies. Does nothing if the flag does not target the user keys.

##### Parameters

- `values`: List of user keys.
- `variationId`: ID of a flag variation.

#### removeValuesFromClause

Removes `values` from the values of the clause indicated by `ruleId` and `clauseId`.

##### Parameters

- `ruleId`: ID of a rule in the flag.
- `clauseId`: ID of a clause in that rule.
- `values`: Array of strings.

#### reorderRules

Rearranges the rules to match the order given in `ruleIds`. Returns an error if `ruleIds` does not match the current set of rules on the flag.

##### Parameters

- `ruleIds`: Array of IDs of all rules in the flag.

#### replacePrerequisites

Removes all existing prerequisites and replaces them with the list you provide.

##### Parameters

- `prerequisites`: A list of prerequisites. Each item in the list must include a flag `key` and `variationId`.

For example, to replace prerequisites, use this request body:

```json

	{
	  "environmentKey": "example-environment-key",
	  "instructions": [
	    {
	      "kind": "replacePrerequisites",
	      "prerequisites": [
	        {
	          "key": "prereq-flag-key",
	          "variationId": "variation-1"
	        },
	        {
	          "key": "another-prereq-flag-key",
	          "variationId": "variation-2"
	        }
	      ]
	    }
	  ]
	}

```

#### replaceRules

Removes all targeting rules for the flag and replaces them with the list you provide.

##### Parameters

- `rules`: A list of rules.

For example, to replace rules, use this request body:

```json

	{
	  "environmentKey": "example-environment-key",
	  "instructions": [
	    {
	      "kind": "replaceRules",
	      "rules": [
	        {
	          "variationId": "variation-1",
	          "description": "myRule",
	          "clauses": [
	            {
	              "attribute": "segmentMatch",
	              "op": "segmentMatch",
	              "values": ["test"]
	            }
	          ],
	          "trackEvents": true
	        }
	      ]
	    }
	  ]
	}

```

#### replaceUserTargets

Removes all existing user targeting and replaces it with the list of targets you provide.

##### Parameters

- `targets`: A list of user targeting. Each item in the list must include a `variationId` and a list of `values`.

For example, to replace user targets, use this request body:

```json

	{
	  "environmentKey": "example-environment-key",
	  "instructions": [
	    {
	      "kind": "replaceUserTargets",
	      "targets": [
	        {
	          "variationId": "variation-1",
	          "values": ["blah", "foo", "bar"]
	        },
	        {
	          "variationId": "variation-2",
	          "values": ["abc", "def"]
	        }
	      ]
	    }
	  ]
	}

```

#### updateClause

Replaces the clause indicated by `ruleId` and `clauseId` with `clause`.

##### Parameters

- `ruleId`: ID of a rule in the flag.
- `clauseId`: ID of a clause in that rule.
- `clause`: New `clause` object, with `attribute` (string), `op` (string), `negate` (boolean), and `values` (array of strings, numbers, or dates) properties.

#### updateFallthroughVariationOrRollout

Updates the default or "fallthrough" rule for the flag, which the flag serves when a user matches none of the targeting rules. The rule can serve either the variation that `variationId` indicates, or a percent rollout that `rolloutWeights` and `rolloutBucketBy` indicate.

##### Parameters

- `variationId`: ID of a variation of the flag.
or
- `rolloutWeights`: Map of `variationId` to weight, in thousandths of a percent (0-100000).
- `rolloutBucketBy`: Optional user attribute.

#### updateOffVariation

Updates the default off variation to `variationId`. The flag serves the default off variation when the flag's targeting is **Off**.

##### Parameters

- `variationId`: ID of a variation of the flag.

#### updatePrerequisite

Changes the prerequisite flag that `key` indicates to use the variation that `variationId` indicates. Returns an error if this prerequisite does not exist.

##### Parameters

- `key`: Flag key of an existing prerequisite flag.
- `variationId`: ID of a variation of the prerequisite flag.

#### updateRuleDescription

Updates the description of the feature flag rule.

##### Parameters

- `description`: The new human-readable description for this rule.
- `ruleId`: The ID of the rule. You can retrieve this by making a GET request for the flag.

#### updateRuleTrackEvents

Updates whether or not LaunchDarkly tracks events for the feature flag associated with this rule.

##### Parameters

- `ruleId`: The ID of the rule. You can retrieve this by making a GET request for the flag.
- `trackEvents`: Whether or not events are tracked.

#### updateRuleVariationOrRollout

Updates what `ruleId` serves when its clauses evaluate to true. The rule can serve either the variation that `variationId` indicates, or a percent rollout that `rolloutWeights` and `rolloutBucketBy` indicate.

##### Parameters

- `ruleId`: ID of a rule in the flag.
- `variationId`: ID of a variation of the flag.

	or

- `rolloutWeights`: Map of `variationId` to weight, in thousandths of a percent (0-100000).
- `rolloutBucketBy`: Optional user attribute.

#### updateTrackEvents

Updates whether or not LaunchDarkly tracks events for the feature flag, for all rules.

##### Parameters

- `trackEvents`: Whether or not events are tracked.

#### updateTrackEventsFallthrough

Updates whether or not LaunchDarkly tracks events for the feature flag, for the default rule.

##### Parameters

- `trackEvents`: Whether or not events are tracked.

</details><br />

<details>
<summary>Click to expand instructions for updating flag settings</summary>

#### addTags

Adds tags to the feature flag.

##### Parameters

- `values`: A list of tags to add.

#### makeFlagPermanent

Marks the feature flag as permanent. LaunchDarkly does not prompt you to remove permanent flags, even if one variation is rolled out to all your users.

#### makeFlagTemporary

Marks the feature flag as temporary.

#### removeMaintainer

Removes the flag's maintainer. To set a new maintainer, use the flag's **Settings** tab in the LaunchDarkly user interface.

#### removeTags

Removes tags from the feature flag.

##### Parameters

- `values`: A list of tags to remove.

#### turnOffClientSideAvailability

Turns off client-side SDK availability for the flag. This is equivalent to unchecking the **SDKs using Mobile Key** and/or **SDKs using client-side ID** boxes for the flag. If you're using a client-side or mobile SDK, you must expose your feature flags in order for the client-side or mobile SDKs to evaluate them.

##### Parameters

- `value`: Use "usingMobileKey" to turn on availability for mobile SDKs. Use "usingEnvironmentId" to turn on availability for client-side SDKs.

#### turnOnClientSideAvailability

Turns on client-side SDK availability for the flag. This is equivalent to unchecking the **SDKs using Mobile Key** and/or **SDKs using client-side ID** boxes for the flag. If you're using a client-side or mobile SDK, you must expose your feature flags in order for the client-side or mobile SDKs to evaluate them.

##### Parameters

- `value`: Use "usingMobileKey" to turn on availability for mobile SDKs. Use "usingEnvironmentId" to turn on availability for client-side SDKs.

#### updateDescription

Updates the feature flag description.

##### Parameters

- `value`: The new description.

#### updateName

Updates the feature flag name.

##### Parameters

- `value`: The new name.

</details><br />

<details>
<summary>Click to expand instructions for updating the flag lifecycle</summary>

#### archiveFlag

Archives the feature flag. This retires it from LaunchDarkly without deleting it. You cannot archive a flag that is a prerequisite of other flags.

#### deleteFlag

Deletes the feature flag and its rules. You cannot restore a deleted flag. If this flag is requested again, the flag value defined in code will be returned for all users.

#### restoreFlag

Restores the feature flag if it was previously archived.

</details>

### Example

The body of a single semantic patch can contain many different instructions.

<details>
<summary>Click to expand example semantic patch request body</summary>

```json

	{
	  "environmentKey": "production",
	  "instructions": [
	    {
	      "kind": "turnFlagOn"
	    },
	    {
	      "kind": "turnFlagOff"
	    },
	    {
	      "kind": "addUserTargets",
	      "variationId": "8bfb304e-d516-47e5-8727-e7f798e8992d",
	      "values": ["userId", "userId2"]
	    },
	    {
	      "kind": "removeUserTargets",
	      "variationId": "8bfb304e-d516-47e5-8727-e7f798e8992d",
	      "values": ["userId3", "userId4"]
	    },
	    {
	      "kind": "updateFallthroughVariationOrRollout",
	      "rolloutWeights": {
	        "variationId": 50000,
	        "variationId2": 50000
	      },
	      "rolloutBucketBy": null
	    },
	    {
	      "kind": "addRule",
	      "clauses": [
	        {
	          "attribute": "segmentMatch",
	          "negate": false,
	          "values": ["test-segment"]
	        }
	      ],
	      "variationId": null,
	      "rolloutWeights": {
	        "variationId": 50000,
	        "variationId2": 50000
	      },
	      "rolloutBucketBy": "key"
	    },
	    {
	      "kind": "removeRule",
	      "ruleId": "99f12464-a429-40fc-86cc-b27612188955"
	    },
	    {
	      "kind": "reorderRules",
	      "ruleIds": ["2f72974e-de68-4243-8dd3-739582147a1f", "8bfb304e-d516-47e5-8727-e7f798e8992d"]
	    },
	    {
	      "kind": "addClauses",
	      "ruleId": "1134",
	      "clauses": [
	        {
	          "attribute": "email",
	          "op": "in",
	          "negate": false,
	          "values": ["test@test.com"]
	        }
	      ]
	    },
	    {
	      "kind": "removeClauses",
	      "ruleId": "1242529",
	      "clauseIds": ["8bfb304e-d516-47e5-8727-e7f798e8992d"]
	    },
	    {
	      "kind": "updateClause",
	      "ruleId": "2f72974e-de68-4243-8dd3-739582147a1f",
	      "clauseId": "309845",
	      "clause": {
	        "attribute": "segmentMatch",
	        "negate": false,
	        "values": ["test-segment"]
	      }
	    },
	    {
	      "kind": "updateRuleVariationOrRollout",
	      "ruleId": "2342",
	      "rolloutWeights": null,
	      "rolloutBucketBy": null
	    },
	    {
	      "kind": "updateOffVariation",
	      "variationId": "3242453"
	    },
	    {
	      "kind": "addPrerequisite",
	      "variationId": "234235",
	      "key": "flagKey2"
	    },
	    {
	      "kind": "updatePrerequisite",
	      "variationId": "234235",
	      "key": "flagKey2"
	    },
	    {
	      "kind": "removePrerequisite",
	      "key": "flagKey"
	    }
	  ]
	}

```
</details>

### Using JSON Patches on a feature flag
If you do not include the header described above, you can use [JSON patch](/reference#updates-using-json-patch).

When using the update feature flag endpoint to add individual users to a specific variation, there are two different patch documents, depending on whether users are already being individually targeted for the variation.

If a flag variation already has users individually targeted, the path for the JSON Patch operation is:

```json

	{
	  "op": "add",
	  "path": "/environments/devint/targets/0/values/-",
	  "value": "TestClient10"
	}

```

If a flag variation does not already have users individually targeted, the path for the JSON Patch operation is:

```json
[

	{
	  "op": "add",
	  "path": "/environments/devint/targets/-",
	  "value": { "variation": 0, "values": ["TestClient10"] }
	}

]
```

### Required approvals
If a request attempts to alter a flag configuration in an environment where approvals are required for the flag, the request will fail with a 405. Changes to the flag configuration in that environment will require creating an [approval request](/tag/Approvals) or a [workflow](/tag/Workflows-(beta)).

### Conflicts
If a flag configuration change made through this endpoint would cause a pending scheduled change or approval request to fail, this endpoint will return a 400. You can ignore this check by adding an `ignoreConflicts` query parameter set to `true`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey The project key
	@param featureFlagKey The feature flag key. The key identifies the flag in your code.
	@return ApiSemanticPatchFeatureFlagRequest
*/
func (a *FeatureFlagsApiService) SemanticPatchFeatureFlag(
	ctx context.Context,
	projectKey string,
	featureFlagKey string,
) ApiSemanticPatchFeatureFlagRequest {
	return ApiSemanticPatchFeatureFlagRequest{
		ApiService:     a,
		ctx:            ctx,
		projectKey:     projectKey,
		featureFlagKey: featureFlagKey,
	}
}

// Execute executes the request
//
//	@return FeatureFlag
func (a *FeatureFlagsApiService) SemanticPatchFeatureFlagExecute(
	r ApiSemanticPatchFeatureFlagRequest,
) (*FeatureFlag, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FeatureFlag
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagsApiService.SemanticPatchFeatureFlag")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/flags/{projectKey}/{featureFlagKey}"
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"projectKey"+"}",
		url.PathEscape(parameterToString(r.projectKey, "")),
		-1,
	)
	localVarPath = strings.Replace(
		localVarPath,
		"{"+"featureFlagKey"+"}",
		url.PathEscape(parameterToString(r.featureFlagKey, "")),
		-1,
	)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.semanticPatchWithComment == nil {
		return localVarReturnValue, nil, reportError("semanticPatchWithComment is required and must be specified")
	}

	// set Content-Type header
	localVarHeaderParams["Content-Type"] = "application/json; domain-model=launchdarkly.semanticpatch"

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.semanticPatchWithComment
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarPostBody,
		localVarHeaderParams,
		localVarQueryParams,
		localVarFormParams,
		formFiles,
	)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InvalidRequestErrorRep
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UnauthorizedErrorRep
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v NotFoundErrorRep
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v MethodNotAllowedErrorRep
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v StatusConflictErrorRep
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RateLimitedErrorRep
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
