// Package drs provides a very basic and simple API for working with V2 of the Amazon Dash Replenishment Services
package drs

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
	"strings"
)

// APIError represents an error from the API and SDK. It implements Error() and contains additional data such as Code and Data. Code represents, in most cases, the HTTP status code. Data will be filled with information that depends on the context of the usage.
type APIError struct {
	Code int
	Data interface{}
}

func (e *APIError) Error() string {
	if e.Data != nil {
		return fmt.Sprintf("%+v", e.Data)
	}
	return fmt.Sprintf("%d", e.Code)
}

type endpoint struct {
	Path     string
	Headers  []endpointHeader
	MockGood string
}

type endpointHeader struct {
	Header string
	Value  string
}

// endpoints holds all of the endpoints the SDK currently supports, including the mockdata needed for tests
var endpoints = map[string]endpoint{
	"subscriptionInfo": endpoint{
		Path: "/subscriptionInfo",
		Headers: []endpointHeader{
			endpointHeader{
				Header: "x-amzn-accept-type",
				Value:  "com.amazon.dash.replenishment.DrsSubscriptionInfoResult@2.0",
			},
			endpointHeader{
				Header: "x-amzn-type-version",
				Value:  "com.amazon.dash.replenishment.DrsSubscriptionInfoInput@1.0",
			},
		},
		MockGood: `{
			"slotsSubscriptionStatus": {
				"slot1": {
					"productInfoList": [{
						"asin": "string",
						"quantity": 1,
						"unit": "count"
					}],
					"subscribed": true
				}
			}
		 }`,
	},
}

//This is the primary method in the package. It should NOT be called directly except through the SDK.
func makeCall(method, endpoint string, userAuth string, data interface{}) (statusCode int, responseData map[string]interface{}, err *APIError) {
	//clean up the url and endpoint
	err = &APIError{}
	if strings.HasPrefix(endpoint, "/") {
		endpoint = endpoint[1:]
	}
	//make sure that endpoint exists
	end, endpointFound := endpoints[endpoint]
	if !endpointFound {
		err.Code = 404
		err.Data = map[string]string{
			"message": "Invalid endpoint",
		}
		return 404, responseData, err
	}

	//if the userAuth is passed in as TEST, we just send mock data back
	if userAuth == "TEST" {
		json.Unmarshal([]byte(end.MockGood), &responseData)

		return 200, responseData, nil
	}

	url := fmt.Sprintf("%s%s", Config.RootURL, end.Path)
	method = strings.ToLower(method)
	if method != "get" && method != "post" && method != "delete" && method != "put" && method != "patch" {
		err.Code = 400
		err.Data = map[string]string{
			"message": fmt.Sprintf("method must be either get, patch, put, post, or delete; received: %s", method),
		}
		return 500, responseData, err
	}

	var response *resty.Response

	request := resty.R().
		SetHeader("Accept", "application/json").
		SetAuthToken(userAuth)

	//loop over the headers and add them in
	for index := range end.Headers {
		request.SetHeader(end.Headers[index].Header, end.Headers[index].Value)
	}

	//now, do what we need to do depending on the method
	var reqErr error
	if method == "get" {
		response, reqErr = request.SetQueryParams(data.(map[string]string)).Get(url)
	} else if method == "delete" {
		response, reqErr = request.Delete(url)
	} else if method == "post" {
		response, reqErr = request.SetBody(data).Post(url)
	} else if method == "put" {
		response, reqErr = request.SetBody(data).Put(url)
	} else if method == "patch" {
		response, reqErr = request.SetBody(data).Patch(url)
	}

	if reqErr != nil {
		err.Code = 500
		err.Data = reqErr.Error()
		return 500, responseData, err
	}

	statusCode = response.StatusCode()
	if statusCode >= 300 {
		apiError := map[string]interface{}{}
		json.Unmarshal(response.Body(), &apiError)
		err.Code = statusCode
		err.Data = apiError
		return statusCode, responseData, err
	}

	json.Unmarshal(response.Body(), &responseData)

	return statusCode, responseData, nil
}
