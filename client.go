/*
Package drs provides a very basic and simple API for working with V2 of the Amazon Dash Replenishment Services
For more information, check the README file at https://github.com/kevineaton/drs-sdk

Most of the functions will require a deviceToken. This is the DRS Access Token retrieved after the user signs up for DRS, most often through LWA. The token will need to be managed and refreshed. At the time of this library's creation, it needed to be refreshed at least once an hour. This library does not handle that responsibility.
*/
package drs

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty"
)

// APIError represents an error from the API and SDK. It implements Error() and contains additional data such as Code and Data. Code represents, in most cases, the HTTP status code. Data will be filled with information that depends on the context of the usage.
type APIError struct {
	Code int
	Data interface{}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%d", e.Code)
}

//This is the primary method in the package. It should NOT be called directly except through the SDK.
//I really want to look at replace the pathParams in a sane way
func makeCall(endpoint string, pathParams []interface{}, deviceAuth string, data interface{}) (statusCode int, responseData map[string]interface{}, err *APIError) {
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
	if deviceAuth == "TEST" {
		json.Unmarshal([]byte(end.MockGood), &responseData)

		return 200, responseData, nil
	}

	url := fmt.Sprintf("%s%s", Config.RootURL, end.Path)
	//some endpoints take path parameters, so we need to do a quick replace here
	//this could probably be more elegant
	if len(pathParams) > 0 && pathParams != nil {
		url = fmt.Sprintf(url, pathParams...)
	}

	var response *resty.Response

	request := resty.R().
		SetHeader("Accept", "application/json").
		SetAuthToken(deviceAuth)

	//loop over the headers and add them in
	for index := range end.Headers {
		request.SetHeader(end.Headers[index].Header, end.Headers[index].Value)
	}

	//now, do what we need to do depending on the method
	var reqErr error
	method := end.Method
	if method == methodGet {
		response, reqErr = request.SetQueryParams(data.(map[string]string)).Get(url)
	} else if method == methodDelete {
		response, reqErr = request.Delete(url)
	} else if method == methodPost {
		response, reqErr = request.SetBody(data).Post(url)
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
