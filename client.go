// Package drs provides a very basic and simple API for working with V2 of the
// Amazon Dash Replenishment Services
// For more information, check the README file at https://github.com/getwagz/drs-sdk
//
// Most of the functions will require a deviceToken. This is the DRS Access Token
// retrieved after the user signs up for DRS, most often through LWA. The token
// will need to be managed and refreshed. At the time of this library's creation,
// it needed to be refreshed at least once an hour. This library does not handle
// that responsibility.
//
package drs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty"
)

// APIError represents an error from the API and SDK. It implements Error() and
// contains additional data such as Code and Data. Code represents, in most
// cases, the HTTP status code. Data will be filled with information that
// depends on the context of the usage.
type APIError struct {
	Code int
	Data interface{}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%d", e.Code)
}

// This is the primary method in the package. It should NOT be called directly
// except through the SDK.
//
// TODO: I really want to look at replace the pathParams in a sane way.
func makeCall(endpoint string, pathParams []interface{}, deviceAuth string, data interface{}) (statusCode int, responseData map[string]interface{}, err error) {
	// Clean up the url and endpoint
	err = &APIError{}
	if strings.HasPrefix(endpoint, "/") {
		endpoint = endpoint[1:]
	}
	// Make sure that endpoint exists
	end, endpointFound := endpoints[endpoint]
	if !endpointFound {
		return http.StatusNotFound, responseData, &APIError{
			Code: http.StatusNotFound,
			Data: map[string]string{
				"message": "Invalid endpoint",
			},
		}
	}

	// If the userAuth is passed in as TEST, we just send mock data back
	if deviceAuth == "TEST" {
		json.Unmarshal([]byte(end.MockGood), &responseData)

		return http.StatusOK, responseData, nil
	}

	url := fmt.Sprintf("%s%s", Config.RootURL, end.Path)
	// Some endpoints take path parameters, so we need to do a quick replace
	// here this could probably be more elegant
	if len(pathParams) > 0 && pathParams != nil {
		url = fmt.Sprintf(url, pathParams...)
	}

	var response *resty.Response

	request := resty.R().
		SetHeader("Accept", "application/json").
		SetAuthToken(deviceAuth)

	// Loop over the headers and add them in
	for index := range end.Headers {
		request.SetHeader(end.Headers[index].Header, end.Headers[index].Value)
	}

	// Now, do what we need to do depending on the method
	var reqErr error

	switch end.Method {
	case http.MethodGet:
		response, reqErr = request.SetQueryParams(data.(map[string]string)).Get(url)
	case http.MethodDelete:
		response, reqErr = request.Delete(url)
	case http.MethodPost:
		response, reqErr = request.SetBody(data).Post(url)
	}

	if reqErr != nil {
		return http.StatusInternalServerError, responseData, &APIError{
			Code: http.StatusInternalServerError,
			Data: reqErr.Error(),
		}
	}

	statusCode = response.StatusCode()
	if statusCode >= http.StatusMultipleChoices {
		apiError := map[string]interface{}{}
		json.Unmarshal(response.Body(), &apiError)
		return statusCode, responseData, &APIError{
			Code: statusCode,
			Data: apiError,
		}
	}

	json.Unmarshal(response.Body(), &responseData)

	return statusCode, responseData, nil
}
