/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.9
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type UserSegmentsApiService service


/* UserSegmentsApiService Delete a user segment.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param projectKey The project key, used to tie the flags together under one project so they can be managed together.
 @param environmentKey The environment key, used to tie together flag configuration and users under one environment so they can be managed together.
 @param userSegmentKey The user segment&#39;s key. The key identifies the user segment in your code.
 @return */
func (a *UserSegmentsApiService) DeleteUserSegment(ctx context.Context, projectKey string, environmentKey string, userSegmentKey string) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/segments/{projectKey}/{environmentKey}/{userSegmentKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", fmt.Sprintf("%v", projectKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentKey"+"}", fmt.Sprintf("%v", environmentKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userSegmentKey"+"}", fmt.Sprintf("%v", userSegmentKey), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* UserSegmentsApiService Get a single user segment by key.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param projectKey The project key, used to tie the flags together under one project so they can be managed together.
 @param environmentKey The environment key, used to tie together flag configuration and users under one environment so they can be managed together.
 @param userSegmentKey The user segment&#39;s key. The key identifies the user segment in your code.
 @return UserSegment*/
func (a *UserSegmentsApiService) GetUserSegment(ctx context.Context, projectKey string, environmentKey string, userSegmentKey string) (UserSegment,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  UserSegment
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/segments/{projectKey}/{environmentKey}/{userSegmentKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", fmt.Sprintf("%v", projectKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentKey"+"}", fmt.Sprintf("%v", environmentKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userSegmentKey"+"}", fmt.Sprintf("%v", userSegmentKey), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* UserSegmentsApiService Get a list of all user segments in the given project.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param projectKey The project key, used to tie the flags together under one project so they can be managed together.
 @param environmentKey The environment key, used to tie together flag configuration and users under one environment so they can be managed together.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "tag" (string) Filter by tag. A tag can be used to group flags across projects.
 @return UserSegments*/
func (a *UserSegmentsApiService) GetUserSegments(ctx context.Context, projectKey string, environmentKey string, localVarOptionals map[string]interface{}) (UserSegments,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  UserSegments
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/segments/{projectKey}/{environmentKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", fmt.Sprintf("%v", projectKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentKey"+"}", fmt.Sprintf("%v", environmentKey), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["tag"], "string", "tag"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["tag"].(string); localVarOk {
		localVarQueryParams.Add("tag", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* UserSegmentsApiService Perform a partial update to a user segment.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param projectKey The project key, used to tie the flags together under one project so they can be managed together.
 @param environmentKey The environment key, used to tie together flag configuration and users under one environment so they can be managed together.
 @param userSegmentKey The user segment&#39;s key. The key identifies the user segment in your code.
 @param patchOnly Requires a JSON Patch representation of the desired changes to the project. &#39;http://jsonpatch.com/&#39; Feature flag patches also support JSON Merge Patch format. &#39;https://tools.ietf.org/html/rfc7386&#39; The addition of comments is also supported.
 @return UserSegment*/
func (a *UserSegmentsApiService) PatchUserSegment(ctx context.Context, projectKey string, environmentKey string, userSegmentKey string, patchOnly []PatchOperation) (UserSegment,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Patch")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  UserSegment
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/segments/{projectKey}/{environmentKey}/{userSegmentKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", fmt.Sprintf("%v", projectKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentKey"+"}", fmt.Sprintf("%v", environmentKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userSegmentKey"+"}", fmt.Sprintf("%v", userSegmentKey), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &patchOnly
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* UserSegmentsApiService Creates a new user segment.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param projectKey The project key, used to tie the flags together under one project so they can be managed together.
 @param environmentKey The environment key, used to tie together flag configuration and users under one environment so they can be managed together.
 @param userSegmentBody Create a new user segment.
 @return */
func (a *UserSegmentsApiService) PostUserSegment(ctx context.Context, projectKey string, environmentKey string, userSegmentBody UserSegmentBody) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/segments/{projectKey}/{environmentKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", fmt.Sprintf("%v", projectKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentKey"+"}", fmt.Sprintf("%v", environmentKey), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &userSegmentBody
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

